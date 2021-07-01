/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/provider-google-api/apis/dialogflow/v1alpha1"
)

// FulfillmentLister helps list Fulfillments.
// All objects returned here must be treated as read-only.
type FulfillmentLister interface {
	// List lists all Fulfillments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Fulfillment, err error)
	// Fulfillments returns an object that can list and get Fulfillments.
	Fulfillments(namespace string) FulfillmentNamespaceLister
	FulfillmentListerExpansion
}

// fulfillmentLister implements the FulfillmentLister interface.
type fulfillmentLister struct {
	indexer cache.Indexer
}

// NewFulfillmentLister returns a new FulfillmentLister.
func NewFulfillmentLister(indexer cache.Indexer) FulfillmentLister {
	return &fulfillmentLister{indexer: indexer}
}

// List lists all Fulfillments in the indexer.
func (s *fulfillmentLister) List(selector labels.Selector) (ret []*v1alpha1.Fulfillment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Fulfillment))
	})
	return ret, err
}

// Fulfillments returns an object that can list and get Fulfillments.
func (s *fulfillmentLister) Fulfillments(namespace string) FulfillmentNamespaceLister {
	return fulfillmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FulfillmentNamespaceLister helps list and get Fulfillments.
// All objects returned here must be treated as read-only.
type FulfillmentNamespaceLister interface {
	// List lists all Fulfillments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Fulfillment, err error)
	// Get retrieves the Fulfillment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Fulfillment, error)
	FulfillmentNamespaceListerExpansion
}

// fulfillmentNamespaceLister implements the FulfillmentNamespaceLister
// interface.
type fulfillmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Fulfillments in the indexer for a given namespace.
func (s fulfillmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Fulfillment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Fulfillment))
	})
	return ret, err
}

// Get retrieves the Fulfillment from the indexer for a given namespace and name.
func (s fulfillmentNamespaceLister) Get(name string) (*v1alpha1.Fulfillment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fulfillment"), name)
	}
	return obj.(*v1alpha1.Fulfillment), nil
}
