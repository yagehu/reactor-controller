package reagentinstance

import (
	"context"

	"go.uber.org/fx"

	"github.com/yagehu/reactor-controller/pkg/generated/clientset/versioned"
)

type Controller interface {
	EnsureReagentInstanceIsActive(
		context.Context, *EnsureReagentInstanceIsActiveParams,
	) (*EnsureReagentInstanceIsActiveResult, error)
}

type Params struct {
	fx.In

	ReactorClient versioned.Interface
}

func New(p Params) (Controller, error) {
	return &controller{
		reactorClient: p.ReactorClient,
	}, nil
}

type controller struct {
	reactorClient versioned.Interface
}
