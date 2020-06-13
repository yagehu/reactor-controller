package kubernetesfx

import (
	"net"

	"go.uber.org/fx"
	"go.uber.org/zap"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/yagehu/reactor-controller/config"
	"github.com/yagehu/reactor-controller/pkg/generated/clientset/versioned"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	Config config.Config
	Logger *zap.Logger
}

type Result struct {
	fx.Out

	Client              kubernetes.Interface
	Config              *rest.Config
	ReactorClient       versioned.Interface
	ApiExtensionsClient apiextensionsclientset.Interface
}

func New(p Params) (Result, error) {
	conf := rest.Config{
		Host: "http://" + net.JoinHostPort(
			p.Config.Kubernetes.ApiServer.Host,
			p.Config.Kubernetes.ApiServer.Port,
		),
	}

	clientset, err := kubernetes.NewForConfig(&conf)
	if err != nil {
		return Result{}, err
	}

	apiExtensionsClient, err := apiextensionsclientset.NewForConfig(&conf)
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

	reactorClient, err := versioned.NewForConfig(&conf)
	if err != nil {
		return Result{}, err
	}

	return Result{
		Client:              clientset,
		Config:              &conf,
		ReactorClient:       reactorClient,
		ApiExtensionsClient: apiExtensionsClient,
	}, nil
}
