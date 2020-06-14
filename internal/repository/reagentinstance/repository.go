package reagentinstance

import (
	"database/sql"

	"go.uber.org/fx"
)

type Repository interface {
}

type Params struct {
	fx.In

	Db *sql.DB
}

func New(p Params) (Repository, error) {
	return &repository{
		db: p.Db,
	}, nil
}

type repository struct {
	db *sql.DB
}
