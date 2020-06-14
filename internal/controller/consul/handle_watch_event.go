package consul

import (
	"context"
	"fmt"

	reactorcontroller "github.com/yagehu/reactor-controller/internal/controller/reactor"
	reagentinstancecontroller "github.com/yagehu/reactor-controller/internal/controller/reagentinstance"
	"github.com/yagehu/reactor-controller/internal/entity"
	"github.com/yagehu/reactor-controller/internal/errs"
)

type HandleWatchEventParams struct {
	Sources []entity.Source
}

type HandleWatchEventResult struct {
}

func (c *controller) HandleWatchEvent(
	ctx context.Context, p *HandleWatchEventParams,
) (*HandleWatchEventResult, error) {
	var (
		op       errs.Op = "controller/consul.HandleWatchEvent"
		reactors         = make(map[entity.Reagent][]entity.Reactor)
	)

	res1, err := c.reactorController.GetAllReactors(
		ctx, &reactorcontroller.GetAllReactorsParams{},
	)
	if err != nil {
		return nil, errs.E(op, err)
	}

	for _, reactor := range res1.Reactors {
		reactors[reactor.Reagent] = append(
			reactors[reactor.Reagent], reactor,
		)
	}

	res2 := c.reagentInstanceController.FilterSources(
		ctx,
		&reagentinstancecontroller.FilterSourcesParams{
			Sources:     p.Sources,
			ReactorsMap: reactors,
		},
	)

	fmt.Println(res2.ReagentInstances)

	return &HandleWatchEventResult{}, nil
}
