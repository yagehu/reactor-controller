package reagentinstance

import (
	"context"
)

type EnsureReagentInstanceIsActiveParams struct {
}

type EnsureReagentInstanceIsActiveResult struct {
}

func (c *controller) EnsureReagentInstanceIsActive(
	ctx context.Context,
	p *EnsureReagentInstanceIsActiveParams,
) (*EnsureReagentInstanceIsActiveResult, error) {
	return &EnsureReagentInstanceIsActiveResult{}, nil
}
