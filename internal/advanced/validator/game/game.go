package gameValidator

import (
	"errors"
	gameDTO "golang-ecommerce-example/internal/advanced/DTO/game"
	"golang-ecommerce-example/internal/advanced/enum/game/platform"
	"golang-ecommerce-example/internal/advanced/enum/game/status"
	"golang-ecommerce-example/internal/advanced/model/game"
	errorTools "golang-ecommerce-example/pkg/error"
	"strings"
)

type Interface interface {
	ValidateInsertGame(g gameModel.Game) error
	ValidateUpdateGame(g gameModel.Game) error
	ValidateGetGamesFilter(r gameDTO.GetGamesFilterRequest) error
	ValidateVerifyGame(g gameDTO.VerifyGameRequest) error
	ValidateGameStatuses(s []string) error
	ValidateGamePlatforms(s []string) error
	ValidateGameID(id int) error
	IsDifferent(g1 gameModel.Game, g2 gameModel.Game) bool
	IsStatusDifferent(s1 string, s2 string) bool
}

type Validator struct {
	ErrMsg errorTools.Enum
}

func NewValidator() Interface {
	var v Validator
	v.ErrMsg = errorTools.ErrFieldValidation
	return v
}

func (v Validator) ValidateInsertGame(g gameModel.Game) error {
	if err := v.ValidateGameStatuses([]string{g.Status}); err != nil {
		return err
	}
	if err := v.ValidateGamePlatforms([]string{g.Platform}); err != nil {
		return err
	}
	if g.Title == "" {
		v.ErrMsg.AddMessage(errors.New("title is required"))
		return v.ErrMsg.Error
	}
	if g.Platform == "" {
		v.ErrMsg.AddMessage(errors.New("platform is required"))
		return v.ErrMsg.Error
	}
	if g.Description == "" {
		v.ErrMsg.AddMessage(errors.New("description is required"))
		return v.ErrMsg.Error
	}
	if g.Status == "" {
		v.ErrMsg.AddMessage(errors.New("status is required"))
		return v.ErrMsg.Error
	}
	return nil
}

func (v Validator) ValidateUpdateGame(g gameModel.Game) error {
	if g.ID == 0 {
		v.ErrMsg.AddMessage(errors.New("id is required"))
		return v.ErrMsg.Error
	}
	if err := v.ValidateGameStatuses([]string{g.Status}); err != nil {
		return err
	}
	if err := v.ValidateGamePlatforms([]string{g.Platform}); err != nil {
		return err
	}
	return nil
}

func (v Validator) ValidateGetGamesFilter(r gameDTO.GetGamesFilterRequest) error {
	if err := v.ValidateGameStatuses(r.Status); err != nil {
		return err
	}
	if err := v.ValidateGamePlatforms(r.Platform); err != nil {
		return err
	}
	return nil
}

func (v Validator) ValidateVerifyGame(g gameDTO.VerifyGameRequest) error {
	if g.ID == 0 {
		v.ErrMsg.AddMessage(errors.New("id is required"))
		return v.ErrMsg.Error
	}
	return nil
}

func (v Validator) ValidateGameStatuses(s []string) error {
	for i := range s {
		if gameStatus.GetEnum(strings.ToUpper(s[i])) == gameStatus.Unknown {
			v.ErrMsg.AddMessage(errors.New("invalid status"))
			return v.ErrMsg.Error
		}
	}
	return nil
}

func (v Validator) ValidateGamePlatforms(s []string) error {
	for i := range s {
		if gamePlatform.GetEnum(
			strings.ToUpper(s[i])) == gamePlatform.Unknown {
			v.ErrMsg.AddMessage(errors.New("invalid platform"))
			return v.ErrMsg.Error
		}
	}
	return nil
}

func (v Validator) ValidateGameID(id int) error {
	if id == 0 {
		v.ErrMsg.AddMessage(errors.New("id is required"))
		return v.ErrMsg.Error
	}
	return nil
}

func (v Validator) IsStatusDifferent(s1 string, s2 string) bool {
	if s1 != s2 {
		return true
	}
	return false
}

func (v Validator) IsDifferent(g1 gameModel.Game, g2 gameModel.Game) bool {
	if g1.Platform != g2.Platform {
		return true
	}
	if g1.Status != g2.Status {
		return true
	}
	if g1.Title != g2.Title {
		return true
	}
	if g1.Description != g2.Description {
		return true
	}
	if g1.Url != g2.Url {
		return true
	}
	if g1.IsDeleted != g2.IsDeleted {
		return true
	}
	return false
}
