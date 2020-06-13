package reactorcontroller

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"

	reactorcontroller "github.com/yagehu/reactor-controller/internal/controller/reactor"
	reactorlisters "github.com/yagehu/reactor-controller/pkg/generated/listers/reactor/v1alpha1"
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
	ReactorLister     reactorlisters.ReactorLister
	ReactorController reactorcontroller.Controller
}

func New(p Params) (Controller, error) {
	return &controller{
		logger:            p.Logger,
		kubernetesClient:  p.KubernetesClient,
		reactorLister:     p.ReactorLister,
		reactorController: p.ReactorController,
	}, nil
}

type controller struct {
	logger            *zap.Logger
	kubernetesClient  kubernetes.Interface
	reactorLister     reactorlisters.ReactorLister
	reactorController reactorcontroller.Controller
}
