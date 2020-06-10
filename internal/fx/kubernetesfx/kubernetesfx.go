package kubernetesfx

import (
	"net"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/yagehu/reactor-controller/config"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	Config config.Config
	Logger *zap.Logger
}

type Result struct {
	fx.Out

	Client kubernetes.Interface
}

func New(p Params) (Result, error) {
	clientset, err := kubernetes.NewForConfig(&rest.Config{
		Host: "http://" + net.JoinHostPort(
			p.Config.Kubernetes.APIServerHost,
			p.Config.Kubernetes.APIServerPort,
		),
	})
	if err != nil {
		return Result{}, err
	}

	version, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return Result{}, err
	}

	p.Logger.Info(
		"Kubernetes server discovered.",
		zap.String("version", version.String()),
	)

	return Result{
		Client: clientset,
	}, nil
}
