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
	v1alpha1 "kubeform.dev/provider-google-api/apis/compute/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ImageIamMemberLister helps list ImageIamMembers.
// All objects returned here must be treated as read-only.
type ImageIamMemberLister interface {
	// List lists all ImageIamMembers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ImageIamMember, err error)
	// ImageIamMembers returns an object that can list and get ImageIamMembers.
	ImageIamMembers(namespace string) ImageIamMemberNamespaceLister
	ImageIamMemberListerExpansion
}

// imageIamMemberLister implements the ImageIamMemberLister interface.
type imageIamMemberLister struct {
	indexer cache.Indexer
}

// NewImageIamMemberLister returns a new ImageIamMemberLister.
func NewImageIamMemberLister(indexer cache.Indexer) ImageIamMemberLister {
	return &imageIamMemberLister{indexer: indexer}
}

// List lists all ImageIamMembers in the indexer.
func (s *imageIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.ImageIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ImageIamMember))
	})
	return ret, err
}

// ImageIamMembers returns an object that can list and get ImageIamMembers.
func (s *imageIamMemberLister) ImageIamMembers(namespace string) ImageIamMemberNamespaceLister {
	return imageIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ImageIamMemberNamespaceLister helps list and get ImageIamMembers.
// All objects returned here must be treated as read-only.
type ImageIamMemberNamespaceLister interface {
	// List lists all ImageIamMembers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ImageIamMember, err error)
	// Get retrieves the ImageIamMember from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ImageIamMember, error)
	ImageIamMemberNamespaceListerExpansion
}

// imageIamMemberNamespaceLister implements the ImageIamMemberNamespaceLister
// interface.
type imageIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ImageIamMembers in the indexer for a given namespace.
func (s imageIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ImageIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ImageIamMember))
	})
	return ret, err
}

// Get retrieves the ImageIamMember from the indexer for a given namespace and name.
func (s imageIamMemberNamespaceLister) Get(name string) (*v1alpha1.ImageIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("imageiammember"), name)
	}
	return obj.(*v1alpha1.ImageIamMember), nil
}
