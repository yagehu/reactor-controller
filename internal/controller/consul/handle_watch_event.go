package consul

import (
	"context"
	"fmt"
	"strings"

	"github.com/gofrs/uuid"

	reactorcontroller "github.com/yagehu/reactor-controller/internal/controller/reactor"
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

	reagentInstances := filterSources(p.Sources, reactors)

	fmt.Println(reagentInstances)

	return &HandleWatchEventResult{}, nil
}

func filterSources(
	sources []entity.Source,
	reactorsMap map[entity.Reagent][]entity.Reactor,
) []entity.ReagentInstance {
	var reagentInstances []entity.ReagentInstance

	for _, source := range sources {
		for reagent, reactors := range reactorsMap {
			_, ok := source.Tags[reagent.Name]
			if !ok {
				continue
			}

			var id string

			for tag := range source.Tags {
				if strings.HasPrefix(tag, reagent.IDPrefix) {
					id = strings.TrimLeft(tag, reagent.IDPrefix)
					break
				}
			}

			reagentInstanceID, err := uuid.FromString(id)
			if err != nil {
				continue
			}

			reagentInstances = append(reagentInstances, entity.ReagentInstance{
				ID:       reagentInstanceID,
				Reactors: reactors,
			})

			break
		}
	}

	return reagentInstances
}
