package reactor

import (
	"context"

	reactorrepository "github.com/yagehu/reactor-controller/internal/repository/reactor"
)

type DeleteReactorParams struct {
	Name string
}

type DeleteReactorResult struct {
}

func (c *controller) DeleteReactor(
	ctx context.Context, p *DeleteReactorParams,
) (*DeleteReactorResult, error) {
	res, err := c.reactorRepository.GetReactorByName(
		ctx,
		&reactorrepository.GetReactorByNameParams{Name: p.Name},
	)
	if err != nil {
		return nil, err
	}

	_, err = c.reactorRepository.DeleteReactor(
		ctx,
		&reactorrepository.DeleteReactorParams{ID: res.Record.ID},
	)
	if err != nil {
		return nil, err
	}

	return &DeleteReactorResult{}, nil
}
