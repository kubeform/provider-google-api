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
	v1alpha1 "kubeform.dev/provider-google-api/apis/runtimeconfig/v1alpha1"
)

// VariableLister helps list Variables.
// All objects returned here must be treated as read-only.
type VariableLister interface {
	// List lists all Variables in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Variable, err error)
	// Variables returns an object that can list and get Variables.
	Variables(namespace string) VariableNamespaceLister
	VariableListerExpansion
}

// variableLister implements the VariableLister interface.
type variableLister struct {
	indexer cache.Indexer
}

// NewVariableLister returns a new VariableLister.
func NewVariableLister(indexer cache.Indexer) VariableLister {
	return &variableLister{indexer: indexer}
}

// List lists all Variables in the indexer.
func (s *variableLister) List(selector labels.Selector) (ret []*v1alpha1.Variable, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Variable))
	})
	return ret, err
}

// Variables returns an object that can list and get Variables.
func (s *variableLister) Variables(namespace string) VariableNamespaceLister {
	return variableNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VariableNamespaceLister helps list and get Variables.
// All objects returned here must be treated as read-only.
type VariableNamespaceLister interface {
	// List lists all Variables in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Variable, err error)
	// Get retrieves the Variable from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Variable, error)
	VariableNamespaceListerExpansion
}

// variableNamespaceLister implements the VariableNamespaceLister
// interface.
type variableNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Variables in the indexer for a given namespace.
func (s variableNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Variable, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Variable))
	})
	return ret, err
}

// Get retrieves the Variable from the indexer for a given namespace and name.
func (s variableNamespaceLister) Get(name string) (*v1alpha1.Variable, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("variable"), name)
	}
	return obj.(*v1alpha1.Variable), nil
}
