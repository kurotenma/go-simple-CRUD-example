package gameQuery

import (
	"github.com/doug-martin/goqu/v9"
	"golang-ecommerce-example/internal/advanced/model/game"
	queryTools "golang-ecommerce-example/pkg/query"
)

type Interface interface {
	InsertGameQuery(game gameModel.Game) *goqu.InsertDataset
	GetGames(perPage uint, page uint) *goqu.SelectDataset
	FilterByStatus(b *goqu.SelectDataset, s []string) *goqu.SelectDataset
	FilterByPlatform(b *goqu.SelectDataset, p []string) *goqu.SelectDataset
	FilterByDeletedStatus(b *goqu.SelectDataset, d []string) *goqu.SelectDataset
	FilterByTitle(b *goqu.SelectDataset, t string) *goqu.SelectDataset
	GetGameCount(deletedStatus []string) *goqu.SelectDataset
	UpdateGame(g gameModel.Game) *goqu.UpdateDataset
}

type Query struct {
	Tools queryTools.Interface
}

func NewQuery() Interface {
	var q Query
	qt := queryTools.NewQueryTools(gameTable)
	q.Tools = qt
	return q
}

var (
	gameTable = "games"
)

func (q Query) InsertGameQuery(g gameModel.Game) *goqu.InsertDataset {
	r := goqu.Record{
		"created_at":  g.CreatedAt,
		"title":       g.Title,
		"url":         g.Url,
		"platform":    g.Platform,
		"description": g.Description,
		"status":      g.Status,
		"is_deleted":  g.IsDeleted,
	}

	b := q.Tools.Query().Insert().Rows(r).Returning("id")
	return b
}
func (q Query) GetGames(perPage uint, page uint) *goqu.SelectDataset {
	b := q.Tools.Query().Select(
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
func (q Query) FilterByTitle(
	b *goqu.SelectDataset,
	t string,
) *goqu.SelectDataset {
	if t != "" {
		b = b.Where(goqu.Ex{"title": goqu.Op{"ilike": t + "%"}})
	}
	return b
}
func (q Query) FilterByStatus(
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
func (q Query) FilterByPlatform(
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
func (q Query) FilterByDeletedStatus(
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
func (q Query) GetGameCount(deletedStatus []string) *goqu.SelectDataset {
	b := q.Tools.Query().Select(goqu.COUNT("*"))
	if len(deletedStatus) < 2 {
		for i := range deletedStatus {
			if deletedStatus[i] == "true" {
				b = q.Tools.SelectIsDeleted(b)
			} else {
				b = q.Tools.SelectIsNotDeleted(b)
			}
		}
	}
	return b
}
func (q Query) UpdateGame(g gameModel.Game) *goqu.UpdateDataset {
	r := goqu.Record{
		"updated_at":  g.UpdatedAt,
		"title":       g.Title,
		"url":         g.Url,
		"platform":    g.Platform,
		"description": g.Description,
		"status":      g.Status,
		"is_deleted":  g.IsDeleted,
	}
	b := q.Tools.Query().Update().Set(r)
	b = q.Tools.UpdateIsNotDeleted(b, gameTable)
	b = b.Where(goqu.Ex{"id": g.ID})
	//b = q.Tools.UpdateIsNotDeleted(b, gameTable)
	return b
}
