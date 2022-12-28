package gameValidator

import (
	"errors"
	"golang-ecommerce-example/internal/advanced/DTO/game"
	"golang-ecommerce-example/internal/advanced/enum/game/platform"
	"golang-ecommerce-example/internal/advanced/enum/game/status"
	"golang-ecommerce-example/internal/advanced/model/game"
	"strings"
)

func ValidateInsertGame(g gameModel.Game) error {
	if err := ValidateGameStatuses([]string{g.Status}); err != nil {
		return err
	}
	if err := ValidateGamePlatforms([]string{g.Platform}); err != nil {
		return err
	}
	if g.Title == "" {
		return errors.New("title is required")
	}
	if g.Platform == "" {
		return errors.New("platform is required")
	}
	if g.Description == "" {
		return errors.New("description is required")
	}
	if g.Status == "" {
		return errors.New("status is required")
	}
	return nil
}

func ValidateUpdateGame(g gameModel.Game) error {
	if g.ID == 0 {
		return errors.New("id is required")
	}
	return nil
}

func ValidateGetGamesFilter(r gameDTO.GetGamesFilterRequest) error {
	if err := ValidateGameStatuses(r.Status); err != nil {
		return err
	}
	if err := ValidateGamePlatforms(r.Platform); err != nil {
		return err
	}
	return nil
}

func ValidateGameStatuses(s []string) error {
	for i := range s {
		if gameStatus.GetEnum(strings.ToUpper(s[i])) == gameStatus.Unknown {
			return errors.New("invalid status")
		}
	}
	return nil
}

func ValidateGamePlatforms(s []string) error {
	for i := range s {
		if gamePlatform.GetEnum(
			strings.ToUpper(s[i])) == gamePlatform.Unknown {
			return errors.New("invalid platform")
		}
	}
	return nil
}
