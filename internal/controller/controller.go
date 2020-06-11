package controller

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

var Module = fx.Options(
	fx.Provide(New),
	fx.Invoke(run),
)

type Controller interface {
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

func run(
	lifecycle fx.Lifecycle,
	c Controller,
	kubernetesClient kubernetes.Interface,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: nil,
	})
}
