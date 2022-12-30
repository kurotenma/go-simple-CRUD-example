package gameControllerSimple

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	gameValidator "golang-ecommerce-example/internal/advanced/validator/game"
	commonData "golang-ecommerce-example/internal/simple/DTO/common"
	gameDTO "golang-ecommerce-example/internal/simple/DTO/game"
	gameStatus "golang-ecommerce-example/internal/simple/enum/game/status"
	gameMapper "golang-ecommerce-example/internal/simple/mapper/game"
	gameModel "golang-ecommerce-example/internal/simple/model/game"
	gameQuery "golang-ecommerce-example/internal/simple/query/game"
	dbPkg "golang-ecommerce-example/pkg/db"
	validatorTools "golang-ecommerce-example/pkg/validator"
	"net/http"
)

func InsertGame(ctx echo.Context) error {
	var req gameDTO.InsertGameRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	v := validator.New()
	if err := v.Struct(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	db, err := dbPkg.InitPgx()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	g := gameMapper.InsertGameRequestToGame(req)
	g.CreatedData()
	b := gameQuery.InsertGameQuery(g)

	q, args, err := b.ToSQL()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	go db.Exec(context.Background(), q, args...)
	return ctx.JSON(http.StatusOK, "Ok")
}
func GetGames(ctx echo.Context) error {
	var req gameDTO.GetGamesFilterRequest
	var games gameModel.Games
	req = gameMapper.QueryParamToGamesFilterRequest(ctx)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	v := validator.New()
	if err := v.Struct(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	db, err := dbPkg.InitPgx()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	perPage, page := 10, 1
	if req.PerPage != perPage && req.PerPage != 0 {
		perPage = req.PerPage
	}
	if req.Page != page && req.Page != 0 {
		page = req.Page
	}
	page -= 1

	b := gameQuery.GetGames(uint(perPage), uint(page))
	q, args, err := b.ToSQL()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	if req.Title != "" {
		b = gameQuery.FilterByTitle(b, req.Title)
	}
	if len(req.Platform) > 0 {
		b = gameQuery.FilterByPlatform(b, req.Platform)
	}
	if len(req.Status) > 0 {
		b = gameQuery.FilterByStatus(b, req.Status)
	}
	if len(req.DeletedStatus) > 0 {
		b = gameQuery.FilterByDeletedStatus(b, req.DeletedStatus)
	}

	rows, err := db.Query(context.Background(), q, args...)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var game gameModel.Game
		if err := rows.Scan(
			&game.CreatedAt,
			&game.ID,
			&game.Title,
			&game.Url,
			&game.Platform,
			&game.Description,
			&game.Status,
			&game.IsDeleted,
		); err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		games = append(games, game)
	}

	var count int
	b = gameQuery.GetGameCount([]string{"false"})
	q, args, err = b.ToSQL()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	if err := db.QueryRow(context.Background(), q).Scan(&count); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, commonData.CommonData{
		Data:      games,
		DataCount: len(games),
		TotalData: count,
		Message:   "Ok",
	})
}
func GetGame(ctx echo.Context) error {
	var g gameModel.Game
	gameID := gameMapper.QueryParamToID(ctx)
	if gameID == 0 {
		return ctx.JSON(http.StatusInternalServerError, "id is required")
	}
	db, err := dbPkg.InitPgx()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	b := gameQuery.GetGameByID(gameID)
	q, args, err := b.ToSQL()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	if err := db.QueryRow(context.Background(), q, args...).Scan(
		&g.CreatedAt,
		&g.ID,
		&g.Title,
		&g.Url,
		&g.Platform,
		&g.Description,
		&g.Status,
		&g.IsDeleted,
	); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, commonData.CommonData{
		Data:      g,
		DataCount: 1,
		TotalData: 1,
		Message:   "Ok",
	})
}
func VerifyGame(ctx echo.Context) error {
	var req gameDTO.VerifyGameRequest
	var g gameModel.Game

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	validator := validatorTools.NewValidator()
	if err := validator.Validate(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	gv := gameValidator.NewValidator()
	if err := gv.ValidateGameID(req.ID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	db, err := dbPkg.InitPgx()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	b := gameQuery.GetGameByID(req.ID)
	q, args, err := b.ToSQL()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	if err := db.QueryRow(context.Background(), q, args...).Scan(
		&g.CreatedAt,
		&g.ID,
		&g.Title,
		&g.Url,
		&g.Platform,
		&g.Description,
		&g.Status,
		&g.IsDeleted,
	); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	if g.Status == gameStatus.Registered.Type {
		return ctx.JSON(http.StatusInternalServerError, "No Change")
	}
	g.UpdatedData()
	g.Status = gameStatus.Registered.Type
	b2 := gameQuery.UpdateGame(g)
	q, args, err = b2.ToSQL()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	go db.Exec(context.Background(), q, args...)

	return ctx.JSON(http.StatusOK, "Game Successfully Updated !")
}
func UpdateGame(ctx echo.Context) error {
	var req gameDTO.UpdateGameRequest
	var g gameModel.Game

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	validator := validatorTools.NewValidator()
	if err := validator.Validate(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	gv := gameValidator.NewValidator()
	if err := gv.ValidateGameID(req.ID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	db, err := dbPkg.InitPgx()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	b := gameQuery.GetGameByID(req.ID)
	q, args, err := b.ToSQL()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	if err := db.QueryRow(context.Background(), q, args...).Scan(
		&g.CreatedAt,
		&g.ID,
		&g.Title,
		&g.Url,
		&g.Platform,
		&g.Description,
		&g.Status,
		&g.IsDeleted,
	); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	g = gameMapper.QueryParamUpdateToGame(ctx, req, g)
	g.UpdatedData()
	b2 := gameQuery.UpdateGame(g)
	q, args, err = b2.ToSQL()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	_, err = db.Exec(context.Background(), q, args...)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, "Game Successfully Updated !")
}
