package gameMapper

import (
	"github.com/labstack/echo/v4"
	"golang-ecommerce-example/internal/advanced/DTO/game"
	"golang-ecommerce-example/internal/advanced/enum/game/status"
	"golang-ecommerce-example/internal/advanced/model/game"
	"strconv"
	"strings"
)

func InsertGameRequestToGame(r gameDTO.InsertGameRequest) gameModel.Game {
	var g gameModel.Game
	g.Title = r.Title
	g.Url = r.Url
	g.Description = r.Description
	g.Platform = r.Platform
	g.Status = gameStatus.NotRegistered.Type

	return g
}

func QueryParamToGamesFilterRequest(
	c echo.Context,
) gameDTO.GetGamesFilterRequest {
	var req gameDTO.GetGamesFilterRequest

	title := c.QueryParam("title")
	if title != "" {
		req.Title = title
	}
	url := c.QueryParam("url")
	if url != "" {
		req.Url = url
	}
	platform := c.QueryParam("platform")
	if platform != "" {
		req.Platform = strings.Split(platform, ",")
	}
	description := c.QueryParam("description")
	if description != "" {
		req.Description = description
	}
	status := c.QueryParam("status")
	if status != "" {
		req.Status = strings.Split(status, ",")
	}
	deletedStatus := c.QueryParam("is_deleted")
	if deletedStatus != "" {
		req.DeletedStatus = strings.Split(deletedStatus, ",")
	}
	perPage := c.QueryParam("per_page")
	if perPage != "" {
		req.PerPage, _ = strconv.Atoi(perPage)
	}
	page := c.QueryParam("page")
	if page != "" {
		req.Page, _ = strconv.Atoi(page)
	}

	return req
}

func GameToGameResponse(game gameModel.Game) gameDTO.GetGameResponse {
	var g gameDTO.GetGameResponse
	g.CreatedAt = game.CreatedAt.Format("2006-01-02 15:04:05")
	g.Title = game.Title
	g.Url = game.Url
	g.Platform = strings.ToUpper(game.Platform)
	g.Description = game.Description
	g.Status = strings.ToUpper(game.Status)
	g.IsDeleted = game.IsDeleted
	return g
}

func GamesToGamesResponse(games gameModel.Games) gameDTO.GetGamesResponse {
	gms := gameDTO.GetGamesResponse{}
	for i := range games {
		g := GameToGameResponse(games[i])
		gms = append(gms, g)
	}
	return gms
}
