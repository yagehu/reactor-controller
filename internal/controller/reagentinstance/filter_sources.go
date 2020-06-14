package reagentinstance

import (
	"context"
	"strings"

	"github.com/gofrs/uuid"

	"github.com/yagehu/reactor-controller/internal/entity"
)

type FilterSourcesParams struct {
	Sources     []entity.Source
	ReactorsMap map[entity.Reagent][]entity.Reactor
}

type FilterSourcesResult struct {
	ReagentInstances []entity.ReagentInstance
}

func (c *controller) FilterSources(
	ctx context.Context, p *FilterSourcesParams,
) *FilterSourcesResult {
	var reagentInstances []entity.ReagentInstance

	for _, source := range p.Sources {
		for reagent, reactors := range p.ReactorsMap {
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

	return &FilterSourcesResult{
		ReagentInstances: reagentInstances,
	}
}
