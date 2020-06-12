package reactor

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Controller interface {
	ProcessEvent(
		context.Context, *ProcessEventParams,
	) (*ProcessEventResult, error)
}

type Params struct {
	fx.In

	Logger *zap.Logger
}

func New(p Params) (Controller, error) {
	return &controller{
		logger: p.Logger,
	}, nil
}

type controller struct {
	logger *zap.Logger
}
