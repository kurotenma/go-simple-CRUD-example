package gameQuery

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"golang-ecommerce-example/internal/simple/model/game"
	queryTools "golang-ecommerce-example/pkg/query-simple"
)

var (
	gameTable = "games"
	db        = queryTools.Dialect().From(gameTable)
)

func InsertGameQuery(g gameModel.Game) *goqu.InsertDataset {
	r := goqu.Record{
		"created_at":  g.CreatedAt,
		"title":       g.Title,
		"url":         g.Url,
		"platform":    g.Platform,
		"description": g.Description,
		"status":      g.Status,
		"is_deleted":  g.IsDeleted,
	}

	fmt.Println("db : ", db)
	b := db.Insert().Rows(r).Returning("id")
	return b.Prepared(true)
}
func GetGames(perPage uint, page uint) *goqu.SelectDataset {
	b := db.Select(
		"created_at",
		"id",
		"title",
		"url",
		"platform",
		"description",
		"status",
		"is_deleted",
	).Limit(perPage).Offset(perPage * page)
	return b.Prepared(true)
}
func FilterByTitle(
	b *goqu.SelectDataset,
	t string,
) *goqu.SelectDataset {
	if t != "" {
		b = b.Where(goqu.Ex{"title": goqu.Op{"ilike": t + "%"}})
	}
	return b.Prepared(true)
}
func FilterByStatus(
	b *goqu.SelectDataset,
	s []string,
) *goqu.SelectDataset {
	if len(s) > 0 {
		b = b.Where(goqu.Ex{"status": s})
	}
	if !b.IsPrepared() {
		b = b.Prepared(true)
	}
	return b
}
func FilterByPlatform(
	b *goqu.SelectDataset,
	p []string,
) *goqu.SelectDataset {
	if len(p) > 0 {
		b = b.Where(goqu.Ex{"platform": p})
	}
	if !b.IsPrepared() {
		b = b.Prepared(true)
	}
	return b
}
func FilterByDeletedStatus(
	b *goqu.SelectDataset,
	d []string,
) *goqu.SelectDataset {
	if len(d) > 0 {
		b = b.Where(goqu.Ex{"is_deleted": d})
	}
	if !b.IsPrepared() {
		b = b.Prepared(true)
	}
	return b
}

func GetGameCount(deletedStatus []string) *goqu.SelectDataset {
	b := db.Select(goqu.COUNT("*"))
	if len(deletedStatus) < 2 {
		for i := range deletedStatus {
			if deletedStatus[i] == "true" {
				b = queryTools.SelectIsDeleted(b, gameTable)
			} else {
				b = queryTools.SelectIsNotDeleted(b, gameTable)
			}
		}
	}
	return b.Prepared(true)
}
