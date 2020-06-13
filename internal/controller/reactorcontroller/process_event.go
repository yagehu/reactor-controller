package reactorcontroller

import (
	"context"
)

type ProcessEventParams struct {
	Name      string
	Namespace string
}

type ProcessEventResult struct {
}

func (c *controller) ProcessEvent(
	ctx context.Context, p *ProcessEventParams,
) (*ProcessEventResult, error) {
	return &ProcessEventResult{}, nil
}
