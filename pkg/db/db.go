package dbPkg

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	dbConfigTools "golang-ecommerce-example/pkg/config/db"
)

type Types struct {
	DB *pgxpool.Pool
	Tx pgx.Tx
}

func (t Types) IsNotEmpty() bool {
	return t != Types{}
}

func NewDatabase(db *pgxpool.Pool, tx pgx.Tx) (*Types, error) {
	var err error
	return &Types{
		DB: db,
		Tx: tx,
	}, err
}

func NewDB() (*Types, error) {
	var t Types

	db, err := InitPgx()
	if err != nil {
		return &t, err
	}

	t.DB = db
	return &t, nil
}

func NewTx(tx pgx.Tx) *Types {
	return &Types{Tx: tx}
}

func InitPgx() (*pgxpool.Pool, error) {
	db, err := dbConfigTools.LoadPostgresDB()
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.New(context.Background(), db.Url())
	if err != nil {
		return nil, err
	}
	return pool, nil
}
