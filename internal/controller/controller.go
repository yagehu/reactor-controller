package controller

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"github.com/yagehu/reactor-controller/pkg/apis/reactor/v1alpha1"
	"github.com/yagehu/reactor-controller/pkg/generated/clientset/versioned"
	"github.com/yagehu/reactor-controller/pkg/generated/clientset/versioned/scheme"
	"github.com/yagehu/reactor-controller/pkg/generated/informers/externalversions"
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
	logger *zap.Logger,
	c Controller,
	kubernetesClient kubernetes.Interface,
	kubernetesConfig *rest.Config,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			reactorClient, err := versioned.NewForConfig(kubernetesConfig)
			if err != nil {
				return err
			}

			apiExtensionsClient, err := apiextensionsclientset.NewForConfig(
				kubernetesConfig,
			)
			if err != nil {
				return err
			}

			sharedInformerFactory := externalversions.NewSharedInformerFactory(
				reactorClient, time.Minute,
			)
			sharedInformer := sharedInformerFactory.
				Huyage().V1alpha1().Reactors()
			sharedInformer.Informer().AddEventHandler(
				cache.ResourceEventHandlerFuncs{
					AddFunc: func(obj interface{}) {
						logger.Info("Added", zap.Any("obj", obj))
					},
					UpdateFunc: func(oldObj, newObj interface{}) {
						logger.Info("Updated")
					},
					DeleteFunc: func(obj interface{}) {
						logger.Info("Deleted")
					},
				},
			)
			sharedInformerFactory.Start(wait.NeverStop)

			if err := v1alpha1.AddToScheme(scheme.Scheme); err != nil {
				return err
			}

			wq := workqueue.NewRateLimitingQueue(
				workqueue.DefaultControllerRateLimiter(),
			)

			fmt.Println(apiExtensionsClient)
			fmt.Println(wq)

			return nil
		},
		OnStop: nil,
	})
}
