package reactorcontroller

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
	reactorcontroller "github.com/yagehu/reactor-controller/internal/controller/reactor"
	"github.com/yagehu/reactor-controller/internal/entity"
	"github.com/yagehu/reactor-controller/pkg/apis/reactor/v1alpha1"
	"github.com/yagehu/reactor-controller/pkg/generated/clientset/versioned"
	"github.com/yagehu/reactor-controller/pkg/generated/clientset/versioned/scheme"
	"github.com/yagehu/reactor-controller/pkg/generated/informers/externalversions"
	reactorv1alpha1 "github.com/yagehu/reactor-controller/pkg/generated/informers/externalversions/reactor/v1alpha1"
)

var Module = fx.Options(fx.Invoke(Start))

type Params struct {
	fx.In

	Config            config.Config
	Lifecycle         fx.Lifecycle
	Logger            *zap.Logger
	KubernetesConfig  *rest.Config
	ReactorClient     versioned.Interface
	ReactorController reactorcontroller.Controller
}

func Start(p Params) error {
	stopCh := make(chan struct{})
	wq := workqueue.NewRateLimitingQueue(
		workqueue.DefaultControllerRateLimiter(),
	)
	sharedInformerFactory := externalversions.NewSharedInformerFactory(
		p.ReactorClient, time.Minute,
	)
	sharedInformer := sharedInformerFactory.Huyage().V1alpha1().Reactors()
	sharedInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				key, err := cache.MetaNamespaceKeyFunc(obj)
				if err != nil {
					p.Logger.Error(
						"Could not make key for API object.",
						zap.Error(err),
					)

					return
				}

				reactor, ok := obj.(*v1alpha1.Reactor)
				if !ok {
					p.Logger.Error(
						"Not a Reactor object.",
						zap.Any("object", obj),
					)

					return
				}

				wq.Add(entity.Event{
					Key:         key,
					Type:        entity.EventTypeCreate,
					ReactorSpec: reactor.Spec,
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

				reactor, ok := newObj.(*v1alpha1.Reactor)
				if !ok {
					p.Logger.Error(
						"Not a Reactor object.",
						zap.Any("newObject", newObj),
						zap.Any("oldObject", oldObj),
					)

					return
				}

				wq.Add(entity.Event{
					Key:         key,
					Type:        entity.EventTypeUpdate,
					ReactorSpec: reactor.Spec,
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

				reactor, ok := obj.(*v1alpha1.Reactor)
				if !ok {
					p.Logger.Error(
						"Not a Reactor object.",
						zap.Any("object", obj),
					)

					return
				}

				wq.Add(entity.Event{
					Key:         key,
					Type:        entity.EventTypeDelete,
					ReactorSpec: reactor.Spec,
				})
			},
		},
	)
	sharedInformerFactory.Start(stopCh)

	if err := v1alpha1.AddToScheme(scheme.Scheme); err != nil {
		return err
	}

	c := &Controller{
		config:            p.Config,
		logger:            p.Logger,
		workQueue:         wq,
		stopCh:            stopCh,
		informer:          sharedInformer,
		reactorController: p.ReactorController,
	}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			c.logger.Info("Starting reactor controller.")

			// wait for the caches to synchronize before starting the worker
			if !cache.WaitForNamedCacheSync(
				"reactor-controller", c.stopCh, c.informer.Informer().HasSynced,
			) {
				return errors.New("timed out waiting for caches to sync")
			}

			c.logger.Info("Reactor controller synced and ready.")

			go c.informer.Informer().Run(c.stopCh)
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
	informer          reactorv1alpha1.ReactorInformer
	reactorController reactorcontroller.Controller
}

func (c *Controller) processNextItem() bool {
	newEvent, quit := c.workQueue.Get()
	if quit {
		return false
	}

	defer c.workQueue.Done(newEvent)

	_, err := c.reactorController.ProcessEvent(
		context.Background(),
		&reactorcontroller.ProcessEventParams{
			Event: newEvent.(entity.Event),
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
