package kubernetesfx

import (
	"net"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"github.com/yagehu/reactor-controller/config"
	"github.com/yagehu/reactor-controller/internal/entity"
	"github.com/yagehu/reactor-controller/pkg/generated/clientset/versioned"
	"github.com/yagehu/reactor-controller/pkg/generated/informers/externalversions"
	"github.com/yagehu/reactor-controller/pkg/generated/informers/externalversions/reactor/v1alpha1"
	reactorlisters "github.com/yagehu/reactor-controller/pkg/generated/listers/reactor/v1alpha1"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	Config config.Config
	Logger *zap.Logger
}

type Result struct {
	fx.Out

	Client                       kubernetes.Interface
	Config                       *rest.Config
	ReactorClient                versioned.Interface
	APIExtensionsClient          apiextensionsclientset.Interface
	RateLimitingWorkQueue        workqueue.RateLimitingInterface
	ReactorLister                reactorlisters.ReactorLister
	SharedReactorInformer        v1alpha1.ReactorInformer
	SharedReactorInformerFactory externalversions.SharedInformerFactory
}

func New(p Params) (Result, error) {
	conf := rest.Config{
		Host: "http://" + net.JoinHostPort(
			p.Config.Kubernetes.APIServer.Host,
			p.Config.Kubernetes.APIServer.Port,
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

	wq := workqueue.NewRateLimitingQueue(
		workqueue.DefaultControllerRateLimiter(),
	)

	sharedInformerFactory := externalversions.NewSharedInformerFactory(
		reactorClient, time.Minute,
	)
	sharedInformer := sharedInformerFactory.Huyage().V1alpha1().Reactors()
	sharedInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				defer func() {
					if r := recover(); r != nil {
						p.Logger.Error(
							"Error handling add event.",
							zap.Any("error", r),
						)
					}
				}()

				key, err := cache.MetaNamespaceKeyFunc(obj)
				if err != nil {
					p.Logger.Error(
						"Could not make key for API object.",
						zap.Error(err),
					)

					return
				}

				wq.Add(entity.Event{
					Key:  key,
					Type: entity.EventTypeCreate,
				})
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(newObj)
				if err != nil {
					p.Logger.Error(
						"Could not make key for API object.",
						zap.Error(err),
						zap.Any("newObject", newObj),
						zap.Any("oldObject", oldObj),
					)

					return
				}

				wq.Add(entity.Event{
					Key:  key,
					Type: entity.EventTypeUpdate,
				})
			},
			DeleteFunc: func(obj interface{}) {
				key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
				if err != nil {
					p.Logger.Error(
						"Could not make key for API object.",
						zap.Error(err),
					)

					return
				}

				wq.Add(entity.Event{
					Key:  key,
					Type: entity.EventTypeDelete,
				})
			},
		},
	)

	return Result{
		Client:                       clientset,
		Config:                       &conf,
		ReactorClient:                reactorClient,
		APIExtensionsClient:          apiExtensionsClient,
		RateLimitingWorkQueue:        wq,
		ReactorLister:                sharedInformer.Lister(),
		SharedReactorInformer:        sharedInformer,
		SharedReactorInformerFactory: sharedInformerFactory,
	}, nil
}
