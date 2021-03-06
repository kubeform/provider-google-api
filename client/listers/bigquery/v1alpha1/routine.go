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
	v1alpha1 "kubeform.dev/provider-google-api/apis/bigquery/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RoutineLister helps list Routines.
// All objects returned here must be treated as read-only.
type RoutineLister interface {
	// List lists all Routines in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Routine, err error)
	// Routines returns an object that can list and get Routines.
	Routines(namespace string) RoutineNamespaceLister
	RoutineListerExpansion
}

// routineLister implements the RoutineLister interface.
type routineLister struct {
	indexer cache.Indexer
}

// NewRoutineLister returns a new RoutineLister.
func NewRoutineLister(indexer cache.Indexer) RoutineLister {
	return &routineLister{indexer: indexer}
}

// List lists all Routines in the indexer.
func (s *routineLister) List(selector labels.Selector) (ret []*v1alpha1.Routine, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Routine))
	})
	return ret, err
}

// Routines returns an object that can list and get Routines.
func (s *routineLister) Routines(namespace string) RoutineNamespaceLister {
	return routineNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RoutineNamespaceLister helps list and get Routines.
// All objects returned here must be treated as read-only.
type RoutineNamespaceLister interface {
	// List lists all Routines in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Routine, err error)
	// Get retrieves the Routine from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Routine, error)
	RoutineNamespaceListerExpansion
}

// routineNamespaceLister implements the RoutineNamespaceLister
// interface.
type routineNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Routines in the indexer for a given namespace.
func (s routineNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Routine, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Routine))
	})
	return ret, err
}

// Get retrieves the Routine from the indexer for a given namespace and name.
func (s routineNamespaceLister) Get(name string) (*v1alpha1.Routine, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("routine"), name)
	}
	return obj.(*v1alpha1.Routine), nil
}
