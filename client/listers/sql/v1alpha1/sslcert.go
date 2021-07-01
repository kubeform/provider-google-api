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
	v1alpha1 "kubeform.dev/provider-google-api/apis/sql/v1alpha1"
)

// SslCertLister helps list SslCerts.
// All objects returned here must be treated as read-only.
type SslCertLister interface {
	// List lists all SslCerts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SslCert, err error)
	// SslCerts returns an object that can list and get SslCerts.
	SslCerts(namespace string) SslCertNamespaceLister
	SslCertListerExpansion
}

// sslCertLister implements the SslCertLister interface.
type sslCertLister struct {
	indexer cache.Indexer
}

// NewSslCertLister returns a new SslCertLister.
func NewSslCertLister(indexer cache.Indexer) SslCertLister {
	return &sslCertLister{indexer: indexer}
}

// List lists all SslCerts in the indexer.
func (s *sslCertLister) List(selector labels.Selector) (ret []*v1alpha1.SslCert, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SslCert))
	})
	return ret, err
}

// SslCerts returns an object that can list and get SslCerts.
func (s *sslCertLister) SslCerts(namespace string) SslCertNamespaceLister {
	return sslCertNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SslCertNamespaceLister helps list and get SslCerts.
// All objects returned here must be treated as read-only.
type SslCertNamespaceLister interface {
	// List lists all SslCerts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SslCert, err error)
	// Get retrieves the SslCert from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SslCert, error)
	SslCertNamespaceListerExpansion
}

// sslCertNamespaceLister implements the SslCertNamespaceLister
// interface.
type sslCertNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SslCerts in the indexer for a given namespace.
func (s sslCertNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SslCert, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SslCert))
	})
	return ret, err
}

// Get retrieves the SslCert from the indexer for a given namespace and name.
func (s sslCertNamespaceLister) Get(name string) (*v1alpha1.SslCert, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sslcert"), name)
	}
	return obj.(*v1alpha1.SslCert), nil
}
