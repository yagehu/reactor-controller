package reactor

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"

	reactorrepository "github.com/yagehu/reactor-controller/internal/repository/reactor"
)

type Controller interface {
	ProcessEvent(
		context.Context, *ProcessEventParams,
	) (*ProcessEventResult, error)
}

type Params struct {
	fx.In

	Logger            *zap.Logger
	KubernetesClient  kubernetes.Interface
	ReactorRepository reactorrepository.Repository
}

func New(p Params) (Controller, error) {
	return &controller{
		logger:            p.Logger,
		kubernetesClient:  p.KubernetesClient,
		reactorRepository: p.ReactorRepository,
	}, nil
}

type controller struct {
	logger            *zap.Logger
	kubernetesClient  kubernetes.Interface
	reactorRepository reactorrepository.Repository
}
