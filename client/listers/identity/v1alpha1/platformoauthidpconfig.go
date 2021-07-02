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

// PlatformOauthIdpConfigLister helps list PlatformOauthIdpConfigs.
// All objects returned here must be treated as read-only.
type PlatformOauthIdpConfigLister interface {
	// List lists all PlatformOauthIdpConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PlatformOauthIdpConfig, err error)
	// PlatformOauthIdpConfigs returns an object that can list and get PlatformOauthIdpConfigs.
	PlatformOauthIdpConfigs(namespace string) PlatformOauthIdpConfigNamespaceLister
	PlatformOauthIdpConfigListerExpansion
}

// platformOauthIdpConfigLister implements the PlatformOauthIdpConfigLister interface.
type platformOauthIdpConfigLister struct {
	indexer cache.Indexer
}

// NewPlatformOauthIdpConfigLister returns a new PlatformOauthIdpConfigLister.
func NewPlatformOauthIdpConfigLister(indexer cache.Indexer) PlatformOauthIdpConfigLister {
	return &platformOauthIdpConfigLister{indexer: indexer}
}

// List lists all PlatformOauthIdpConfigs in the indexer.
func (s *platformOauthIdpConfigLister) List(selector labels.Selector) (ret []*v1alpha1.PlatformOauthIdpConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PlatformOauthIdpConfig))
	})
	return ret, err
}

// PlatformOauthIdpConfigs returns an object that can list and get PlatformOauthIdpConfigs.
func (s *platformOauthIdpConfigLister) PlatformOauthIdpConfigs(namespace string) PlatformOauthIdpConfigNamespaceLister {
	return platformOauthIdpConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PlatformOauthIdpConfigNamespaceLister helps list and get PlatformOauthIdpConfigs.
// All objects returned here must be treated as read-only.
type PlatformOauthIdpConfigNamespaceLister interface {
	// List lists all PlatformOauthIdpConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PlatformOauthIdpConfig, err error)
	// Get retrieves the PlatformOauthIdpConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PlatformOauthIdpConfig, error)
	PlatformOauthIdpConfigNamespaceListerExpansion
}

// platformOauthIdpConfigNamespaceLister implements the PlatformOauthIdpConfigNamespaceLister
// interface.
type platformOauthIdpConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PlatformOauthIdpConfigs in the indexer for a given namespace.
func (s platformOauthIdpConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PlatformOauthIdpConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PlatformOauthIdpConfig))
	})
	return ret, err
}

// Get retrieves the PlatformOauthIdpConfig from the indexer for a given namespace and name.
func (s platformOauthIdpConfigNamespaceLister) Get(name string) (*v1alpha1.PlatformOauthIdpConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("platformoauthidpconfig"), name)
	}
	return obj.(*v1alpha1.PlatformOauthIdpConfig), nil
}
