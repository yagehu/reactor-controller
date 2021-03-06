// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/yagehu/reactor-controller/pkg/apis/reactor/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ReactorLister helps list Reactors.
type ReactorLister interface {
	// List lists all Reactors in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Reactor, err error)
	// Reactors returns an object that can list and get Reactors.
	Reactors(namespace string) ReactorNamespaceLister
	ReactorListerExpansion
}

// reactorLister implements the ReactorLister interface.
type reactorLister struct {
	indexer cache.Indexer
}

// NewReactorLister returns a new ReactorLister.
func NewReactorLister(indexer cache.Indexer) ReactorLister {
	return &reactorLister{indexer: indexer}
}

// List lists all Reactors in the indexer.
func (s *reactorLister) List(selector labels.Selector) (ret []*v1alpha1.Reactor, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Reactor))
	})
	return ret, err
}

// Reactors returns an object that can list and get Reactors.
func (s *reactorLister) Reactors(namespace string) ReactorNamespaceLister {
	return reactorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ReactorNamespaceLister helps list and get Reactors.
type ReactorNamespaceLister interface {
	// List lists all Reactors in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Reactor, err error)
	// Get retrieves the Reactor from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Reactor, error)
	ReactorNamespaceListerExpansion
}

// reactorNamespaceLister implements the ReactorNamespaceLister
// interface.
type reactorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Reactors in the indexer for a given namespace.
func (s reactorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Reactor, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Reactor))
	})
	return ret, err
}

// Get retrieves the Reactor from the indexer for a given namespace and name.
func (s reactorNamespaceLister) Get(name string) (*v1alpha1.Reactor, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("reactor"), name)
	}
	return obj.(*v1alpha1.Reactor), nil
}
