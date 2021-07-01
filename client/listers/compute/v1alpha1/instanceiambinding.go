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
	v1alpha1 "kubeform.dev/provider-google-api/apis/compute/v1alpha1"
)

// InstanceIamBindingLister helps list InstanceIamBindings.
// All objects returned here must be treated as read-only.
type InstanceIamBindingLister interface {
	// List lists all InstanceIamBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstanceIamBinding, err error)
	// InstanceIamBindings returns an object that can list and get InstanceIamBindings.
	InstanceIamBindings(namespace string) InstanceIamBindingNamespaceLister
	InstanceIamBindingListerExpansion
}

// instanceIamBindingLister implements the InstanceIamBindingLister interface.
type instanceIamBindingLister struct {
	indexer cache.Indexer
}

// NewInstanceIamBindingLister returns a new InstanceIamBindingLister.
func NewInstanceIamBindingLister(indexer cache.Indexer) InstanceIamBindingLister {
	return &instanceIamBindingLister{indexer: indexer}
}

// List lists all InstanceIamBindings in the indexer.
func (s *instanceIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.InstanceIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InstanceIamBinding))
	})
	return ret, err
}

// InstanceIamBindings returns an object that can list and get InstanceIamBindings.
func (s *instanceIamBindingLister) InstanceIamBindings(namespace string) InstanceIamBindingNamespaceLister {
	return instanceIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InstanceIamBindingNamespaceLister helps list and get InstanceIamBindings.
// All objects returned here must be treated as read-only.
type InstanceIamBindingNamespaceLister interface {
	// List lists all InstanceIamBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstanceIamBinding, err error)
	// Get retrieves the InstanceIamBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InstanceIamBinding, error)
	InstanceIamBindingNamespaceListerExpansion
}

// instanceIamBindingNamespaceLister implements the InstanceIamBindingNamespaceLister
// interface.
type instanceIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InstanceIamBindings in the indexer for a given namespace.
func (s instanceIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InstanceIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InstanceIamBinding))
	})
	return ret, err
}

// Get retrieves the InstanceIamBinding from the indexer for a given namespace and name.
func (s instanceIamBindingNamespaceLister) Get(name string) (*v1alpha1.InstanceIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("instanceiambinding"), name)
	}
	return obj.(*v1alpha1.InstanceIamBinding), nil
}
