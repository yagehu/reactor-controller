package consul

import (
	"context"

	"go.uber.org/fx"

	reactorcontroller "github.com/yagehu/reactor-controller/internal/controller/reactor"
	reagentinstancecontroller "github.com/yagehu/reactor-controller/internal/controller/reagentinstance"
)

type Controller interface {
	HandleWatchEvent(
		context.Context, *HandleWatchEventParams,
	) (*HandleWatchEventResult, error)
}

type Params struct {
	fx.In

	ReactorController         reactorcontroller.Controller
	ReagentInstanceController reagentinstancecontroller.Controller
}

func New(p Params) (Controller, error) {
	return &controller{
		reactorController:         p.ReactorController,
		reagentInstanceController: p.ReagentInstanceController,
	}, nil
}

type controller struct {
	reactorController         reactorcontroller.Controller
	reagentInstanceController reagentinstancecontroller.Controller
}
