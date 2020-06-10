// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	versioned "pkg/generated/clientset/versioned"
	internalinterfaces "pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "pkg/generated/listers/reactor/v1alpha1"
	time "time"

	reactorv1alpha1 "github.com/yagehu/reactor-controller/pkg/apis/reactor/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ReactorInformer provides access to a shared informer and lister for
// Reactors.
type ReactorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ReactorLister
}

type reactorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewReactorInformer constructs a new informer for Reactor type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewReactorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredReactorInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredReactorInformer constructs a new informer for Reactor type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredReactorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HuyageV1alpha1().Reactors(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HuyageV1alpha1().Reactors(namespace).Watch(context.TODO(), options)
			},
		},
		&reactorv1alpha1.Reactor{},
		resyncPeriod,
		indexers,
	)
}

func (f *reactorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredReactorInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *reactorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&reactorv1alpha1.Reactor{}, f.defaultInformer)
}

func (f *reactorInformer) Lister() v1alpha1.ReactorLister {
	return v1alpha1.NewReactorLister(f.Informer().GetIndexer())
}
