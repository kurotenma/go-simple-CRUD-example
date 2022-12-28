package userRepository

import (
	"context"
	"golang-ecommerce-example/internal/advanced/DTO/game"
	"golang-ecommerce-example/internal/advanced/model/game"
	"golang-ecommerce-example/internal/advanced/query/game"
	dbPkg "golang-ecommerce-example/pkg/db"
)

type Interface interface {
	InsertGame(game gameModel.Game) (gameModel.Game, error)
	GetGames(filter gameDTO.GetGamesFilterRequest) (gameModel.Games, error)
	GetGameCount(deletedStatus []string) (int, error)
}

type Repository struct {
	*dbPkg.Types
	Query gameQuery.Interface
}

func NewRepository(db *dbPkg.Types) Interface {
	var r Repository
	r.Types = db
	r.Query = gameQuery.NewQuery()

	return r
}
func (r Repository) InsertGame(g gameModel.Game) (gameModel.Game, error) {
	b := r.Query.InsertGameQuery(g)
	q, args, err := b.ToSQL()
	if err != nil {
		return g, err
	}
	err = r.DB.QueryRow(context.Background(), q, args...).Scan(&g.ID)
	if err != nil {
		return g, err
	}
	return g, nil
}
func (r Repository) GetGames(f gameDTO.GetGamesFilterRequest) (
	gameModel.Games, error,
) {
	var gs gameModel.Games

	perPage, page := 10, 1
	if f.PerPage != perPage && f.PerPage != 0 {
		perPage = f.PerPage
	}
	if f.Page != page && f.Page != 0 {
		page = f.Page
	}
	page -= 1

	b := r.Query.GetGames(uint(perPage), uint(page))
	if f.Title != "" {
		b = r.Query.FilterByTitle(b, f.Title)
	}
	if len(f.Platform) > 0 {
		b = r.Query.FilterByPlatform(b, f.Platform)
	}
	if len(f.Status) > 0 {
		b = r.Query.FilterByStatus(b, f.Status)
	}
	if len(f.DeletedStatus) > 0 {
		b = r.Query.FilterByDeletedStatus(b, f.DeletedStatus)
	}
	q, args, err := b.ToSQL()
	if err != nil {
		return gs, err
	}
	rows, err := r.DB.Query(context.Background(), q, args...)
	if err != nil {
		return gs, err
	}
	defer rows.Close()
	for rows.Next() {
		var g gameModel.Game
		if err = rows.Scan(
			&g.CreatedAt,
			&g.ID,
			&g.Title,
			&g.Url,
			&g.Platform,
			&g.Description,
			&g.Status,
			&g.IsDeleted,
		); err != nil {
			return gs, err
		}
		gs = append(gs, g)
	}
	return gs, nil
}
func (r Repository) GetGameCount(deletedStatus []string) (int, error) {
	var count int
	b := r.Query.GetGameCount(deletedStatus)
	q, args, err := b.ToSQL()
	if err != nil {
		return 0, err
	}
	err = r.DB.QueryRow(context.Background(), q, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
func (r Repository) UpdateGame(g gameModel.Game) (gameModel.Game, error) {
	b := r.Query.UpdateGame(g)
	q, args, err := b.ToSQL()
	if err != nil {
		return g, err
	}
	_, err = r.DB.Exec(context.Background(), q, args)
	if err != nil {
		return g, err
	}
	return g, nil
}
