package reactor

import (
	"context"

	"github.com/gofrs/uuid"

	"github.com/yagehu/reactor-controller/internal/entity"
	"github.com/yagehu/reactor-controller/internal/mapper"
	reactorrepository "github.com/yagehu/reactor-controller/internal/repository/reactor"
)

type CreateReactorParams struct {
	ReagentName string
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
		ID: reactorID,
		Reagent: entity.Reagent{
			ID:   reagentID,
			Name: p.ReagentName,
		},
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
