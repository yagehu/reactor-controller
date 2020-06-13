package reactor

import (
	"context"
	"database/sql"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Repository interface {
	CreateReactor(
		context.Context, *CreateReactorParams,
	) (*CreateReactorResult, error)
	DeleteReactor(
		context.Context, *DeleteReactorParams,
	) (*DeleteReactorResult, error)
	GetAllReactors(
		context.Context, *GetAllReactorsParams,
	) (*GetAllReactorsResult, error)
	GetReactorByName(
		context.Context, *GetReactorByNameParams,
	) (*GetReactorByNameResult, error)
}

type Params struct {
	fx.In

	Db     *sql.DB
	Logger *zap.Logger
}

func New(p Params) (Repository, error) {
	return &repository{
		db:     p.Db,
		logger: p.Logger,
	}, nil
}

type repository struct {
	db     *sql.DB
	logger *zap.Logger
}
