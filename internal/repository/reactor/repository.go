package reactor

import (
	"context"
	"database/sql"

	"go.uber.org/fx"
)

type Repository interface {
	CreateReactor(
		context.Context, *CreateReactorParams,
	) (*CreateReactorResult, error)
	DeleteReactor(
		context.Context, *DeleteReactorParams,
	) (*DeleteReactorResult, error)
	GetReactorByName(
		context.Context, *GetReactorByNameParams,
	) (*GetReactorByNameResult, error)
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
