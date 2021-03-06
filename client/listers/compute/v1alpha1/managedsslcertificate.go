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

// ManagedSslCertificateLister helps list ManagedSslCertificates.
// All objects returned here must be treated as read-only.
type ManagedSslCertificateLister interface {
	// List lists all ManagedSslCertificates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagedSslCertificate, err error)
	// ManagedSslCertificates returns an object that can list and get ManagedSslCertificates.
	ManagedSslCertificates(namespace string) ManagedSslCertificateNamespaceLister
	ManagedSslCertificateListerExpansion
}

// managedSslCertificateLister implements the ManagedSslCertificateLister interface.
type managedSslCertificateLister struct {
	indexer cache.Indexer
}

// NewManagedSslCertificateLister returns a new ManagedSslCertificateLister.
func NewManagedSslCertificateLister(indexer cache.Indexer) ManagedSslCertificateLister {
	return &managedSslCertificateLister{indexer: indexer}
}

// List lists all ManagedSslCertificates in the indexer.
func (s *managedSslCertificateLister) List(selector labels.Selector) (ret []*v1alpha1.ManagedSslCertificate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagedSslCertificate))
	})
	return ret, err
}

// ManagedSslCertificates returns an object that can list and get ManagedSslCertificates.
func (s *managedSslCertificateLister) ManagedSslCertificates(namespace string) ManagedSslCertificateNamespaceLister {
	return managedSslCertificateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagedSslCertificateNamespaceLister helps list and get ManagedSslCertificates.
// All objects returned here must be treated as read-only.
type ManagedSslCertificateNamespaceLister interface {
	// List lists all ManagedSslCertificates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagedSslCertificate, err error)
	// Get retrieves the ManagedSslCertificate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ManagedSslCertificate, error)
	ManagedSslCertificateNamespaceListerExpansion
}

// managedSslCertificateNamespaceLister implements the ManagedSslCertificateNamespaceLister
// interface.
type managedSslCertificateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagedSslCertificates in the indexer for a given namespace.
func (s managedSslCertificateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagedSslCertificate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagedSslCertificate))
	})
	return ret, err
}

// Get retrieves the ManagedSslCertificate from the indexer for a given namespace and name.
func (s managedSslCertificateNamespaceLister) Get(name string) (*v1alpha1.ManagedSslCertificate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managedsslcertificate"), name)
	}
	return obj.(*v1alpha1.ManagedSslCertificate), nil
}
