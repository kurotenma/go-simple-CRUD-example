package userService

import (
	"golang-ecommerce-example/internal/advanced/DTO/game"
	"golang-ecommerce-example/internal/advanced/model/game"
	gameRepository "golang-ecommerce-example/internal/advanced/repository/game"
	"golang-ecommerce-example/internal/advanced/validator/game"
	dbPkg "golang-ecommerce-example/pkg/db"
	errorTools "golang-ecommerce-example/pkg/error"
)

type Interface interface {
	InsertGame(game gameModel.Game) (gameModel.Game, error)
	GetGames(filter gameDTO.GetGamesFilterRequest) (gameModel.Games, error)
	GetGame(id int) (gameModel.Game, error)
	GetGameCount(deletedStatus []string) (int, error)
	UpdateGame(g gameModel.Game) error
}

type Repositories struct {
	GameRepository gameRepository.Interface
}

type Service struct {
	DB    *dbPkg.Types
	Error errorTools.ErrorStruct
	Repositories
}

func NewService(db *dbPkg.Types) Interface {
	var s Service
	if db.IsNotEmpty() {
		s.DB = db
		s.GameRepository = gameRepository.NewRepository(s.DB)
	}
	return s
}

func (s Service) InsertGame(g gameModel.Game) (gameModel.Game, error) {
	gv := gameValidator.NewValidator()
	if err := gv.ValidateInsertGame(g); err != nil {
		return g, err
	}
	g, err := s.GameRepository.InsertGame(g)
	if err != nil {
		return g, err
	}
	return g, nil
}

func (s Service) GetGames(f gameDTO.GetGamesFilterRequest) (
	gameModel.Games, error,
) {
	var gs gameModel.Games
	gs, err := s.GameRepository.GetGames(f)
	if err != nil {
		return gs, err
	}
	return gs, err
}

func (s Service) GetGame(id int) (gameModel.Game, error) {
	var g gameModel.Game
	gv := gameValidator.NewValidator()
	if err := gv.ValidateGameID(id); err != nil {
		return g, err
	}
	g, err := s.GameRepository.GetGame(id)
	if err != nil {
		return g, err
	}
	return g, nil
}

func (s Service) GetGameCount(deletedStatus []string) (int, error) {
	var c int
	c, err := s.GameRepository.GetGameCount(deletedStatus)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (s Service) UpdateGame(g gameModel.Game) error {
	gv := gameValidator.NewValidator()
	if err := gv.ValidateUpdateGame(g); err != nil {
		return err
	}

	if err := s.GameRepository.UpdateGame(g); err != nil {
		return err
	}
	return nil
}
