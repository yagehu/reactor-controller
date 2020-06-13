package reactor

import (
	"context"

	"github.com/yagehu/reactor-controller/internal/entity"
	"github.com/yagehu/reactor-controller/internal/errs"
	"github.com/yagehu/reactor-controller/internal/mapper"
	reactorrepository "github.com/yagehu/reactor-controller/internal/repository/reactor"
)

type GetReactorByNameParams struct {
	Name string
}

type GetReactorByNameResult struct {
	Reactor entity.Reactor
}

func (c *controller) GetReactorByName(
	ctx context.Context, p *GetReactorByNameParams,
) (*GetReactorByNameResult, error) {
	var op errs.Op = "controller/reactor.GetReactorByName"

	res, err := c.reactorRepository.GetReactorByName(
		ctx, &reactorrepository.GetReactorByNameParams{Name: p.Name},
	)
	if err != nil {
		return nil, errs.E(op, err)
	}

	reactor, err := mapper.FromReactorModel(res.Record)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &GetReactorByNameResult{
		Reactor: reactor,
	}, nil
}
