package reactor

import (
	"context"
)

type CreateReactorParams struct {
}

type CreateReactorResult struct {
}

func (r *repository) CreateReactor(
	ctx context.Context, p *CreateReactorParams,
) (*CreateReactorResult, error) {
	return &CreateReactorResult{}, nil
}
