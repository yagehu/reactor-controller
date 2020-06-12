package reactor

import (
	"context"
	"fmt"

	"go.uber.org/zap"

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
	var err error

	switch p.Event.Type {
	case entity.EventTypeCreate:
		err = c.processCreateEvent(ctx, p.Event)
	case entity.EventTypeUpdate:
		err = c.processUpdateEvent(ctx, p.Event)
	case entity.EventTypeDelete:
		err = c.processDeleteEvent(ctx, p.Event)
	default:
		c.logger.Error("Unknown event type.", zap.Any("event", p.Event))
		return nil, fmt.Errorf("unknown event type %s", p.Event.Type)
	}

	if err != nil {
		return nil, err
	}

	return &ProcessEventResult{}, nil
}

func (c *controller) processCreateEvent(
	ctx context.Context, event entity.Event,
) error {
	return nil
}

func (c *controller) processUpdateEvent(
	ctx context.Context, event entity.Event,
) error {
	return nil
}

func (c *controller) processDeleteEvent(
	ctx context.Context, event entity.Event,
) error {
	return nil
}
