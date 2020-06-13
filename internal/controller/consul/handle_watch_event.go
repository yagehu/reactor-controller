package consul

import (
	"context"

	reactorcontroller "github.com/yagehu/reactor-controller/internal/controller/reactor"
	"github.com/yagehu/reactor-controller/internal/entity"
	"github.com/yagehu/reactor-controller/internal/errs"
)

type HandleWatchEventParams struct {
	ReagentInstances []entity.ReagentInstance
}

type HandleWatchEventResult struct {
}

func (c *controller) HandleWatchEvent(
	ctx context.Context, p *HandleWatchEventParams,
) (*HandleWatchEventResult, error) {
	var (
		op       errs.Op = "controller/consul.HandleWatchEvent"
		reactors         = make(map[string][]entity.Reactor)
	)

	res, err := c.reactorController.GetAllReactors(
		ctx, &reactorcontroller.GetAllReactorsParams{},
	)
	if err != nil {
		return nil, errs.E(op, err)
	}

	for _, reactor := range res.Reactors {
		reactors[reactor.Reagent.Name] = append(
			reactors[reactor.Reagent.Name], reactor,
		)
	}

	return &HandleWatchEventResult{}, nil
}
