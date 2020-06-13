package reactor

import (
	"context"

	"go.uber.org/fx"
)

type Repository interface {
	CreateReactor(
		context.Context, *CreateReactorParams,
	) (*CreateReactorResult, error)
}

type Params struct {
	fx.In
}

func New(p Params) (Repository, error) {
	return &repository{}, nil
}

type repository struct {
}
