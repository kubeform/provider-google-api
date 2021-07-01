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
	v1alpha1 "kubeform.dev/provider-google-api/apis/monitoring/v1alpha1"
)

// MetricDescriptorLister helps list MetricDescriptors.
// All objects returned here must be treated as read-only.
type MetricDescriptorLister interface {
	// List lists all MetricDescriptors in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MetricDescriptor, err error)
	// MetricDescriptors returns an object that can list and get MetricDescriptors.
	MetricDescriptors(namespace string) MetricDescriptorNamespaceLister
	MetricDescriptorListerExpansion
}

// metricDescriptorLister implements the MetricDescriptorLister interface.
type metricDescriptorLister struct {
	indexer cache.Indexer
}

// NewMetricDescriptorLister returns a new MetricDescriptorLister.
func NewMetricDescriptorLister(indexer cache.Indexer) MetricDescriptorLister {
	return &metricDescriptorLister{indexer: indexer}
}

// List lists all MetricDescriptors in the indexer.
func (s *metricDescriptorLister) List(selector labels.Selector) (ret []*v1alpha1.MetricDescriptor, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MetricDescriptor))
	})
	return ret, err
}

// MetricDescriptors returns an object that can list and get MetricDescriptors.
func (s *metricDescriptorLister) MetricDescriptors(namespace string) MetricDescriptorNamespaceLister {
	return metricDescriptorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MetricDescriptorNamespaceLister helps list and get MetricDescriptors.
// All objects returned here must be treated as read-only.
type MetricDescriptorNamespaceLister interface {
	// List lists all MetricDescriptors in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MetricDescriptor, err error)
	// Get retrieves the MetricDescriptor from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MetricDescriptor, error)
	MetricDescriptorNamespaceListerExpansion
}

// metricDescriptorNamespaceLister implements the MetricDescriptorNamespaceLister
// interface.
type metricDescriptorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MetricDescriptors in the indexer for a given namespace.
func (s metricDescriptorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MetricDescriptor, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MetricDescriptor))
	})
	return ret, err
}

// Get retrieves the MetricDescriptor from the indexer for a given namespace and name.
func (s metricDescriptorNamespaceLister) Get(name string) (*v1alpha1.MetricDescriptor, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("metricdescriptor"), name)
	}
	return obj.(*v1alpha1.MetricDescriptor), nil
}
