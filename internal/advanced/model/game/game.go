package gameModel

import (
	"golang-ecommerce-example/internal/advanced/model/common/base"
)

type Game struct {
	baseModel.Base
	ID          int
	Title       string
	Url         string
	Platform    string
	Description string
	Status      string
}

type Games []Game

type Filter struct {
	Title         string
	Url           string
	Platform      []string
	Description   string
	Status        []string
	DeletedStatus []bool
}
