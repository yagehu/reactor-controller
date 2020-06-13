package reactor

import (
	"context"

	"go.uber.org/fx"

	reactorrepository "github.com/yagehu/reactor-controller/internal/repository/reactor"
)

type Controller interface {
	CreateReactor(
		context.Context, *CreateReactorParams,
	) (*CreateReactorResult, error)
}

type Params struct {
	fx.In

	ReactorRepository reactorrepository.Repository
}

func New(p Params) (Controller, error) {
	return &controller{
		reactorRepository: p.ReactorRepository,
	}, nil
}

type controller struct {
	reactorRepository reactorrepository.Repository
}
