package reactor

import (
	"context"

	"github.com/yagehu/reactor-controller/internal/entity"
	"github.com/yagehu/reactor-controller/internal/errs"
	"github.com/yagehu/reactor-controller/internal/mapper"
	reactorrepository "github.com/yagehu/reactor-controller/internal/repository/reactor"
)

type GetAllReactorsParams struct {
}

type GetAllReactorsResult struct {
	Reactors []entity.Reactor
}

func (c *controller) GetAllReactors(
	ctx context.Context, _ *GetAllReactorsParams,
) (*GetAllReactorsResult, error) {
	var op errs.Op = "controller/reactor.GetAllReactors"

	res, err := c.reactorRepository.GetAllReactors(
		ctx, &reactorrepository.GetAllReactorsParams{},
	)
	if err != nil {
		return nil, errs.E(op, err)
	}

	reactors, err := mapper.FromReactorModelList(res.Records)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &GetAllReactorsResult{Reactors: reactors}, nil
}
