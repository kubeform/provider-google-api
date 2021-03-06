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
	v1alpha1 "kubeform.dev/provider-google-api/apis/pubsub/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LiteReservationLister helps list LiteReservations.
// All objects returned here must be treated as read-only.
type LiteReservationLister interface {
	// List lists all LiteReservations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LiteReservation, err error)
	// LiteReservations returns an object that can list and get LiteReservations.
	LiteReservations(namespace string) LiteReservationNamespaceLister
	LiteReservationListerExpansion
}

// liteReservationLister implements the LiteReservationLister interface.
type liteReservationLister struct {
	indexer cache.Indexer
}

// NewLiteReservationLister returns a new LiteReservationLister.
func NewLiteReservationLister(indexer cache.Indexer) LiteReservationLister {
	return &liteReservationLister{indexer: indexer}
}

// List lists all LiteReservations in the indexer.
func (s *liteReservationLister) List(selector labels.Selector) (ret []*v1alpha1.LiteReservation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LiteReservation))
	})
	return ret, err
}

// LiteReservations returns an object that can list and get LiteReservations.
func (s *liteReservationLister) LiteReservations(namespace string) LiteReservationNamespaceLister {
	return liteReservationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LiteReservationNamespaceLister helps list and get LiteReservations.
// All objects returned here must be treated as read-only.
type LiteReservationNamespaceLister interface {
	// List lists all LiteReservations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LiteReservation, err error)
	// Get retrieves the LiteReservation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LiteReservation, error)
	LiteReservationNamespaceListerExpansion
}

// liteReservationNamespaceLister implements the LiteReservationNamespaceLister
// interface.
type liteReservationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LiteReservations in the indexer for a given namespace.
func (s liteReservationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LiteReservation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LiteReservation))
	})
	return ret, err
}

// Get retrieves the LiteReservation from the indexer for a given namespace and name.
func (s liteReservationNamespaceLister) Get(name string) (*v1alpha1.LiteReservation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("litereservation"), name)
	}
	return obj.(*v1alpha1.LiteReservation), nil
}
