package reactor

import (
	"context"
	"time"

	"github.com/gofrs/uuid"

	"github.com/yagehu/reactor-controller/internal/entity"
	"github.com/yagehu/reactor-controller/internal/mapper"
	reactorrepository "github.com/yagehu/reactor-controller/internal/repository/reactor"
)

type CreateReactorParams struct {
	Name            string
	ReagentName     string
	ReagentIDPrefix string
}

type CreateReactorResult struct {
	Reactor entity.Reactor
}

func (c *controller) CreateReactor(
	ctx context.Context, p *CreateReactorParams,
) (*CreateReactorResult, error) {
	reactorID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	reagentID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	reactor := entity.Reactor{
		ID:   reactorID,
		Name: p.Name,
		Reagent: entity.Reagent{
			ID:       reagentID,
			Name:     p.ReagentName,
			IDPrefix: p.ReagentIDPrefix,
		},
		CreatedAt: time.Now(),
	}

	_, err = c.reactorRepository.CreateReactor(
		ctx,
		&reactorrepository.CreateReactorParams{
			Reactor: mapper.ToReactorModel(reactor),
		},
	)
	if err != nil {
		return nil, err
	}

	return &CreateReactorResult{
		Reactor: reactor,
	}, nil
}
