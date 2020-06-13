package reactor

import (
	"context"
	"errors"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"github.com/yagehu/reactor-controller/config"
	reactorcontroller "github.com/yagehu/reactor-controller/internal/controller/reactorcontroller"
	"github.com/yagehu/reactor-controller/internal/entity"
	"github.com/yagehu/reactor-controller/pkg/apis/reactor/v1alpha1"
	"github.com/yagehu/reactor-controller/pkg/generated/clientset/versioned"
	"github.com/yagehu/reactor-controller/pkg/generated/clientset/versioned/scheme"
	"github.com/yagehu/reactor-controller/pkg/generated/informers/externalversions"
	reactorinformers "github.com/yagehu/reactor-controller/pkg/generated/informers/externalversions/reactor/v1alpha1"
)

var Module = fx.Invoke(Start)

type Params struct {
	fx.In

	Config                       config.Config
	Lifecycle                    fx.Lifecycle
	Logger                       *zap.Logger
	KubernetesConfig             *rest.Config
	ReactorClient                versioned.Interface
	ReactorController            reactorcontroller.Controller
	RateLimitingWorkQueue        workqueue.RateLimitingInterface
	SharedReactorInformer        reactorinformers.ReactorInformer
	SharedReactorInformerFactory externalversions.SharedInformerFactory
}

func Start(p Params) error {
	stopCh := make(chan struct{})
	c := &Controller{
		config:            p.Config,
		logger:            p.Logger,
		workQueue:         p.RateLimitingWorkQueue,
		stopCh:            stopCh,
		reactorController: p.ReactorController,
	}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			c.logger.Info("Starting reactor controller.")

			p.SharedReactorInformerFactory.Start(stopCh)
			if err := v1alpha1.AddToScheme(scheme.Scheme); err != nil {
				return err
			}

			// wait for the caches to synchronize before starting the worker
			if !cache.WaitForNamedCacheSync(
				"reactor-controller",
				c.stopCh,
				p.SharedReactorInformer.Informer().HasSynced,
			) {
				return errors.New("timed out waiting for caches to sync")
			}

			c.logger.Info("Reactor controller synced and ready.")

			go p.SharedReactorInformer.Informer().Run(c.stopCh)
			go func() {
				// Loop until "something bad" happens.
				// The .Until will then rekick the worker after one second.
				wait.Until(
					func() {
						for c.processNextItem() {
						}
					},
					time.Second,
					c.stopCh,
				)
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			c.logger.Info("Shutting down work queue.")
			c.workQueue.ShutDown()
			close(c.stopCh)

			return nil
		},
	})

	return nil
}

type Controller struct {
	config            config.Config
	logger            *zap.Logger
	workQueue         workqueue.RateLimitingInterface
	stopCh            chan struct{}
	reactorController reactorcontroller.Controller
}

func (c *Controller) processNextItem() bool {
	newEvent, quit := c.workQueue.Get()
	if quit {
		return false
	}

	defer c.workQueue.Done(newEvent)

	namespace, name, err := cache.SplitMetaNamespaceKey(
		newEvent.(entity.Event).Key,
	)
	if err != nil {
		c.logger.Error(
			"Could not get namespace and name from event.", zap.Error(err),
		)
		c.workQueue.Forget(newEvent)

		return true
	}

	_, err = c.reactorController.ProcessEvent(
		context.Background(),
		&reactorcontroller.ProcessEventParams{
			Name:      name,
			Namespace: namespace,
		},
	)
	if err != nil {
		if c.workQueue.NumRequeues(newEvent) < c.config.Controller.WorkQueueEventRetries {
			c.logger.Error(
				"Error processing event. Will retry.",
				zap.Error(err),
				zap.String("key", newEvent.(entity.Event).Key),
			)
			c.workQueue.AddRateLimited(newEvent)

			return true
		}

		c.logger.Error(
			"Error processing event. Giving up.",
			zap.Error(err),
			zap.String("key", newEvent.(entity.Event).Key),
		)
		c.workQueue.Forget(newEvent)
	}

	// No error, reset the rate limit counters.
	c.workQueue.Forget(newEvent)

	return true
}
