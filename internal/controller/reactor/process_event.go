package reactor

import (
	"context"

	"github.com/yagehu/reactor-controller/internal/entity"
)

type ProcessEventParams struct {
	Event entity.Event
}

type ProcessEventResult struct {
}

func (c *controller) ProcessEvent(
	ctx context.Context, p *ProcessEventParams,
) (*ProcessEventResult, error) {
	return &ProcessEventResult{}, nil
}
