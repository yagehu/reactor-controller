package consul

import (
	"context"

	"go.uber.org/fx"

	reactorcontroller "github.com/yagehu/reactor-controller/internal/controller/reactor"
)

type Controller interface {
	HandleWatchEvent(
		context.Context, *HandleWatchEventParams,
	) (*HandleWatchEventResult, error)
}

type Params struct {
	fx.In

	ReactorController reactorcontroller.Controller
}

func New(p Params) (Controller, error) {
	return &controller{
		reactorController: p.ReactorController,
	}, nil
}

type controller struct {
	reactorController reactorcontroller.Controller
}
