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
	v1alpha1 "kubeform.dev/provider-google-api/apis/identity/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PlatformTenantDefaultSupportedIdpConfigLister helps list PlatformTenantDefaultSupportedIdpConfigs.
// All objects returned here must be treated as read-only.
type PlatformTenantDefaultSupportedIdpConfigLister interface {
	// List lists all PlatformTenantDefaultSupportedIdpConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PlatformTenantDefaultSupportedIdpConfig, err error)
	// PlatformTenantDefaultSupportedIdpConfigs returns an object that can list and get PlatformTenantDefaultSupportedIdpConfigs.
	PlatformTenantDefaultSupportedIdpConfigs(namespace string) PlatformTenantDefaultSupportedIdpConfigNamespaceLister
	PlatformTenantDefaultSupportedIdpConfigListerExpansion
}

// platformTenantDefaultSupportedIdpConfigLister implements the PlatformTenantDefaultSupportedIdpConfigLister interface.
type platformTenantDefaultSupportedIdpConfigLister struct {
	indexer cache.Indexer
}

// NewPlatformTenantDefaultSupportedIdpConfigLister returns a new PlatformTenantDefaultSupportedIdpConfigLister.
func NewPlatformTenantDefaultSupportedIdpConfigLister(indexer cache.Indexer) PlatformTenantDefaultSupportedIdpConfigLister {
	return &platformTenantDefaultSupportedIdpConfigLister{indexer: indexer}
}

// List lists all PlatformTenantDefaultSupportedIdpConfigs in the indexer.
func (s *platformTenantDefaultSupportedIdpConfigLister) List(selector labels.Selector) (ret []*v1alpha1.PlatformTenantDefaultSupportedIdpConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PlatformTenantDefaultSupportedIdpConfig))
	})
	return ret, err
}

// PlatformTenantDefaultSupportedIdpConfigs returns an object that can list and get PlatformTenantDefaultSupportedIdpConfigs.
func (s *platformTenantDefaultSupportedIdpConfigLister) PlatformTenantDefaultSupportedIdpConfigs(namespace string) PlatformTenantDefaultSupportedIdpConfigNamespaceLister {
	return platformTenantDefaultSupportedIdpConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PlatformTenantDefaultSupportedIdpConfigNamespaceLister helps list and get PlatformTenantDefaultSupportedIdpConfigs.
// All objects returned here must be treated as read-only.
type PlatformTenantDefaultSupportedIdpConfigNamespaceLister interface {
	// List lists all PlatformTenantDefaultSupportedIdpConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PlatformTenantDefaultSupportedIdpConfig, err error)
	// Get retrieves the PlatformTenantDefaultSupportedIdpConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PlatformTenantDefaultSupportedIdpConfig, error)
	PlatformTenantDefaultSupportedIdpConfigNamespaceListerExpansion
}

// platformTenantDefaultSupportedIdpConfigNamespaceLister implements the PlatformTenantDefaultSupportedIdpConfigNamespaceLister
// interface.
type platformTenantDefaultSupportedIdpConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PlatformTenantDefaultSupportedIdpConfigs in the indexer for a given namespace.
func (s platformTenantDefaultSupportedIdpConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PlatformTenantDefaultSupportedIdpConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PlatformTenantDefaultSupportedIdpConfig))
	})
	return ret, err
}

// Get retrieves the PlatformTenantDefaultSupportedIdpConfig from the indexer for a given namespace and name.
func (s platformTenantDefaultSupportedIdpConfigNamespaceLister) Get(name string) (*v1alpha1.PlatformTenantDefaultSupportedIdpConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("platformtenantdefaultsupportedidpconfig"), name)
	}
	return obj.(*v1alpha1.PlatformTenantDefaultSupportedIdpConfig), nil
}
