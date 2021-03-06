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
	v1alpha1 "kubeform.dev/provider-google-api/apis/storage/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ObjectAccessControlLister helps list ObjectAccessControls.
// All objects returned here must be treated as read-only.
type ObjectAccessControlLister interface {
	// List lists all ObjectAccessControls in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ObjectAccessControl, err error)
	// ObjectAccessControls returns an object that can list and get ObjectAccessControls.
	ObjectAccessControls(namespace string) ObjectAccessControlNamespaceLister
	ObjectAccessControlListerExpansion
}

// objectAccessControlLister implements the ObjectAccessControlLister interface.
type objectAccessControlLister struct {
	indexer cache.Indexer
}

// NewObjectAccessControlLister returns a new ObjectAccessControlLister.
func NewObjectAccessControlLister(indexer cache.Indexer) ObjectAccessControlLister {
	return &objectAccessControlLister{indexer: indexer}
}

// List lists all ObjectAccessControls in the indexer.
func (s *objectAccessControlLister) List(selector labels.Selector) (ret []*v1alpha1.ObjectAccessControl, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ObjectAccessControl))
	})
	return ret, err
}

// ObjectAccessControls returns an object that can list and get ObjectAccessControls.
func (s *objectAccessControlLister) ObjectAccessControls(namespace string) ObjectAccessControlNamespaceLister {
	return objectAccessControlNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ObjectAccessControlNamespaceLister helps list and get ObjectAccessControls.
// All objects returned here must be treated as read-only.
type ObjectAccessControlNamespaceLister interface {
	// List lists all ObjectAccessControls in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ObjectAccessControl, err error)
	// Get retrieves the ObjectAccessControl from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ObjectAccessControl, error)
	ObjectAccessControlNamespaceListerExpansion
}

// objectAccessControlNamespaceLister implements the ObjectAccessControlNamespaceLister
// interface.
type objectAccessControlNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ObjectAccessControls in the indexer for a given namespace.
func (s objectAccessControlNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ObjectAccessControl, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ObjectAccessControl))
	})
	return ret, err
}

// Get retrieves the ObjectAccessControl from the indexer for a given namespace and name.
func (s objectAccessControlNamespaceLister) Get(name string) (*v1alpha1.ObjectAccessControl, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("objectaccesscontrol"), name)
	}
	return obj.(*v1alpha1.ObjectAccessControl), nil
}
