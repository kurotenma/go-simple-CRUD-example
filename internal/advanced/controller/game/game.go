package gameController

import (
	"github.com/labstack/echo/v4"
	"golang-ecommerce-example/internal/advanced/DTO/common"
	"golang-ecommerce-example/internal/advanced/DTO/game"
	gameStatus "golang-ecommerce-example/internal/advanced/enum/game/status"
	"golang-ecommerce-example/internal/advanced/mapper/game"
	"golang-ecommerce-example/internal/advanced/model/game"
	gameService "golang-ecommerce-example/internal/advanced/service/game"
	"golang-ecommerce-example/internal/advanced/validator/game"
	dbPkg "golang-ecommerce-example/pkg/db"
	validatorTools "golang-ecommerce-example/pkg/validator"
	"net/http"
)

type Interface interface {
	InsertGame(ctx echo.Context) error
	GetGames(ctx echo.Context) error
	GetGame(ctx echo.Context) error
	VerifyGame(ctx echo.Context) error
	UpdateGame(ctx echo.Context) error
}
type Services struct {
	GameService gameService.Interface
}
type Models struct {
	gameModel.Game
	gameModel.Games
}
type Controller struct {
	Models
	Services
	Response commonData.CommonData
}

func NewController() Interface {
	var c Controller
	c.Models.Game = gameModel.Game{}
	c.Models.Games = gameModel.Games{}
	c.Response = commonData.CommonData{Message: "Ok"}
	return c
}

func (c Controller) InsertGame(ctx echo.Context) error {
	var req gameDTO.InsertGameRequest

	if err := ctx.Bind(&req); err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}

	validator := validatorTools.NewValidator()
	if err := validator.Validate(req); err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}

	db, err := dbPkg.NewDB()
	if err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}
	c.GameService = gameService.NewService(db)
	c.Game = gameMapper.InsertGameRequestToGame(req)
	c.Game.CreatedData()
	c.Game, err = c.GameService.InsertGame(c.Game)
	if err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}

	return ctx.JSON(http.StatusOK, "Game Successfully Registered !")
}
func (c Controller) GetGames(ctx echo.Context) error {
	var req gameDTO.GetGamesFilterRequest
	req = gameMapper.QueryParamToGamesFilterRequest(ctx)

	validator := validatorTools.NewValidator()
	if err := validator.Validate(req); err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}
	gv := gameValidator.NewValidator()
	if err := gv.ValidateGetGamesFilter(req); err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}

	db, err := dbPkg.NewDB()
	if err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}
	c.GameService = gameService.NewService(db)
	c.Games, err = c.GameService.GetGames(req)
	if err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}
	count, err := c.GameService.GetGameCount([]string{"false"})
	if err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}

	gms := gameMapper.GamesToGamesResponse(c.Games)
	c.Response.Data = gms
	c.Response.DataCount = len(gms)
	c.Response.TotalData = count
	return ctx.JSON(http.StatusOK, c.Response)
}
func (c Controller) GetGame(ctx echo.Context) error {
	gameID := gameMapper.QueryParamToID(ctx)
	gv := gameValidator.NewValidator()
	if err := gv.ValidateGameID(gameID); err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}

	db, err := dbPkg.NewDB()
	if err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}
	c.GameService = gameService.NewService(db)
	c.Game, err = c.GameService.GetGame(gameID)
	if err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}
	c.Response.Data = gameMapper.GameToGameResponse(c.Game)
	c.Response.DataCount = 1
	c.Response.TotalData = 1
	return ctx.JSON(http.StatusOK, c.Response)
}
func (c Controller) VerifyGame(ctx echo.Context) error {
	var req gameDTO.VerifyGameRequest
	if err := ctx.Bind(&req); err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}

	validator := validatorTools.NewValidator()
	if err := validator.Validate(req); err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}
	gv := gameValidator.NewValidator()
	if err := gv.ValidateGameID(req.ID); err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}

	db, err := dbPkg.NewDB()
	if err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}
	c.GameService = gameService.NewService(db)
	c.Game, err = c.GameService.GetGame(req.ID)
	if err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}
	if !gv.IsStatusDifferent(c.Game.Status, gameStatus.Registered.Type) {
		c.Response.Message = "No Change"
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}

	c.Game.UpdatedData()
	c.Game.Status = gameStatus.Registered.Type
	if err := c.GameService.UpdateGame(c.Game); err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}

	return ctx.JSON(http.StatusOK, "Game Successfully Updated !")
}
func (c Controller) UpdateGame(ctx echo.Context) error {
	var req gameDTO.UpdateGameRequest
	if err := ctx.Bind(&req); err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}
	validator := validatorTools.NewValidator()
	if err := validator.Validate(req); err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}
	db, err := dbPkg.NewDB()
	if err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}
	c.GameService = gameService.NewService(db)
	c.Game, err = c.GameService.GetGame(req.ID)
	if err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}
	g := gameMapper.QueryParamUpdateToGame(ctx, req, c.Game)
	gv := gameValidator.NewValidator()
	if !gv.IsDifferent(g, c.Game) {
		c.Response.Message = "No Change"
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}
	c.Game.UpdatedData()
	if err := c.GameService.UpdateGame(g); err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}
	return ctx.JSON(http.StatusOK, "Game Successfully Updated !")
}
