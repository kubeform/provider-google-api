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

// DiskResourcePolicyAttachmentLister helps list DiskResourcePolicyAttachments.
// All objects returned here must be treated as read-only.
type DiskResourcePolicyAttachmentLister interface {
	// List lists all DiskResourcePolicyAttachments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DiskResourcePolicyAttachment, err error)
	// DiskResourcePolicyAttachments returns an object that can list and get DiskResourcePolicyAttachments.
	DiskResourcePolicyAttachments(namespace string) DiskResourcePolicyAttachmentNamespaceLister
	DiskResourcePolicyAttachmentListerExpansion
}

// diskResourcePolicyAttachmentLister implements the DiskResourcePolicyAttachmentLister interface.
type diskResourcePolicyAttachmentLister struct {
	indexer cache.Indexer
}

// NewDiskResourcePolicyAttachmentLister returns a new DiskResourcePolicyAttachmentLister.
func NewDiskResourcePolicyAttachmentLister(indexer cache.Indexer) DiskResourcePolicyAttachmentLister {
	return &diskResourcePolicyAttachmentLister{indexer: indexer}
}

// List lists all DiskResourcePolicyAttachments in the indexer.
func (s *diskResourcePolicyAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.DiskResourcePolicyAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DiskResourcePolicyAttachment))
	})
	return ret, err
}

// DiskResourcePolicyAttachments returns an object that can list and get DiskResourcePolicyAttachments.
func (s *diskResourcePolicyAttachmentLister) DiskResourcePolicyAttachments(namespace string) DiskResourcePolicyAttachmentNamespaceLister {
	return diskResourcePolicyAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DiskResourcePolicyAttachmentNamespaceLister helps list and get DiskResourcePolicyAttachments.
// All objects returned here must be treated as read-only.
type DiskResourcePolicyAttachmentNamespaceLister interface {
	// List lists all DiskResourcePolicyAttachments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DiskResourcePolicyAttachment, err error)
	// Get retrieves the DiskResourcePolicyAttachment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DiskResourcePolicyAttachment, error)
	DiskResourcePolicyAttachmentNamespaceListerExpansion
}

// diskResourcePolicyAttachmentNamespaceLister implements the DiskResourcePolicyAttachmentNamespaceLister
// interface.
type diskResourcePolicyAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DiskResourcePolicyAttachments in the indexer for a given namespace.
func (s diskResourcePolicyAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DiskResourcePolicyAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DiskResourcePolicyAttachment))
	})
	return ret, err
}

// Get retrieves the DiskResourcePolicyAttachment from the indexer for a given namespace and name.
func (s diskResourcePolicyAttachmentNamespaceLister) Get(name string) (*v1alpha1.DiskResourcePolicyAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("diskresourcepolicyattachment"), name)
	}
	return obj.(*v1alpha1.DiskResourcePolicyAttachment), nil
}
