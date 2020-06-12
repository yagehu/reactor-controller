package reactor

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

type Controller interface {
	ProcessEvent(
		context.Context, *ProcessEventParams,
	) (*ProcessEventResult, error)
}

type Params struct {
	fx.In

	Logger           *zap.Logger
	KubernetesClient kubernetes.Interface
}

func New(p Params) (Controller, error) {
	return &controller{
		logger:           p.Logger,
		kubernetesClient: p.KubernetesClient,
	}, nil
}

type controller struct {
	logger           *zap.Logger
	kubernetesClient kubernetes.Interface
}
