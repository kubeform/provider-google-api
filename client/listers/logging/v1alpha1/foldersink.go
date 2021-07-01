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
	v1alpha1 "kubeform.dev/provider-google-api/apis/logging/v1alpha1"
)

// FolderSinkLister helps list FolderSinks.
// All objects returned here must be treated as read-only.
type FolderSinkLister interface {
	// List lists all FolderSinks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FolderSink, err error)
	// FolderSinks returns an object that can list and get FolderSinks.
	FolderSinks(namespace string) FolderSinkNamespaceLister
	FolderSinkListerExpansion
}

// folderSinkLister implements the FolderSinkLister interface.
type folderSinkLister struct {
	indexer cache.Indexer
}

// NewFolderSinkLister returns a new FolderSinkLister.
func NewFolderSinkLister(indexer cache.Indexer) FolderSinkLister {
	return &folderSinkLister{indexer: indexer}
}

// List lists all FolderSinks in the indexer.
func (s *folderSinkLister) List(selector labels.Selector) (ret []*v1alpha1.FolderSink, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FolderSink))
	})
	return ret, err
}

// FolderSinks returns an object that can list and get FolderSinks.
func (s *folderSinkLister) FolderSinks(namespace string) FolderSinkNamespaceLister {
	return folderSinkNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FolderSinkNamespaceLister helps list and get FolderSinks.
// All objects returned here must be treated as read-only.
type FolderSinkNamespaceLister interface {
	// List lists all FolderSinks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FolderSink, err error)
	// Get retrieves the FolderSink from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FolderSink, error)
	FolderSinkNamespaceListerExpansion
}

// folderSinkNamespaceLister implements the FolderSinkNamespaceLister
// interface.
type folderSinkNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FolderSinks in the indexer for a given namespace.
func (s folderSinkNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FolderSink, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FolderSink))
	})
	return ret, err
}

// Get retrieves the FolderSink from the indexer for a given namespace and name.
func (s folderSinkNamespaceLister) Get(name string) (*v1alpha1.FolderSink, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("foldersink"), name)
	}
	return obj.(*v1alpha1.FolderSink), nil
}
