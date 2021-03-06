// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FeatureGateLister helps list FeatureGates.
type FeatureGateLister interface {
	// List lists all FeatureGates in the indexer.
	List(selector labels.Selector) (ret []*v1.FeatureGate, err error)
	// Get retrieves the FeatureGate from the index for a given name.
	Get(name string) (*v1.FeatureGate, error)
	FeatureGateListerExpansion
}

// featureGateLister implements the FeatureGateLister interface.
type featureGateLister struct {
	indexer cache.Indexer
}

// NewFeatureGateLister returns a new FeatureGateLister.
func NewFeatureGateLister(indexer cache.Indexer) FeatureGateLister {
	return &featureGateLister{indexer: indexer}
}

// List lists all FeatureGates in the indexer.
func (s *featureGateLister) List(selector labels.Selector) (ret []*v1.FeatureGate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.FeatureGate))
	})
	return ret, err
}

// Get retrieves the FeatureGate from the index for a given name.
func (s *featureGateLister) Get(name string) (*v1.FeatureGate, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("featuregate"), name)
	}
	return obj.(*v1.FeatureGate), nil
}
