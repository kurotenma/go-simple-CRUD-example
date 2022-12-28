package gameController

import (
	"github.com/labstack/echo/v4"
	"golang-ecommerce-example/internal/advanced/DTO/common"
	"golang-ecommerce-example/internal/advanced/DTO/game"
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
	if err := gameValidator.ValidateGetGamesFilter(req); err != nil {
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

func (c Controller) VerifyGame(ctx echo.Context) error {
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
	c.Game = gameMapper.InsertGameRequestToGame(req)
	c.Game.CreatedData()
	c.Game, err = c.GameService.InsertGame(c.Game)
	if err != nil {
		c.Response.Message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, c.Response)
	}

	return ctx.JSON(http.StatusOK, "Game Successfully Registered !")
}
