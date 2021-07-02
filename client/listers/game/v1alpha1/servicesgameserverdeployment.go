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
	v1alpha1 "kubeform.dev/provider-google-api/apis/game/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServicesGameServerDeploymentLister helps list ServicesGameServerDeployments.
// All objects returned here must be treated as read-only.
type ServicesGameServerDeploymentLister interface {
	// List lists all ServicesGameServerDeployments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServicesGameServerDeployment, err error)
	// ServicesGameServerDeployments returns an object that can list and get ServicesGameServerDeployments.
	ServicesGameServerDeployments(namespace string) ServicesGameServerDeploymentNamespaceLister
	ServicesGameServerDeploymentListerExpansion
}

// servicesGameServerDeploymentLister implements the ServicesGameServerDeploymentLister interface.
type servicesGameServerDeploymentLister struct {
	indexer cache.Indexer
}

// NewServicesGameServerDeploymentLister returns a new ServicesGameServerDeploymentLister.
func NewServicesGameServerDeploymentLister(indexer cache.Indexer) ServicesGameServerDeploymentLister {
	return &servicesGameServerDeploymentLister{indexer: indexer}
}

// List lists all ServicesGameServerDeployments in the indexer.
func (s *servicesGameServerDeploymentLister) List(selector labels.Selector) (ret []*v1alpha1.ServicesGameServerDeployment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServicesGameServerDeployment))
	})
	return ret, err
}

// ServicesGameServerDeployments returns an object that can list and get ServicesGameServerDeployments.
func (s *servicesGameServerDeploymentLister) ServicesGameServerDeployments(namespace string) ServicesGameServerDeploymentNamespaceLister {
	return servicesGameServerDeploymentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServicesGameServerDeploymentNamespaceLister helps list and get ServicesGameServerDeployments.
// All objects returned here must be treated as read-only.
type ServicesGameServerDeploymentNamespaceLister interface {
	// List lists all ServicesGameServerDeployments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServicesGameServerDeployment, err error)
	// Get retrieves the ServicesGameServerDeployment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServicesGameServerDeployment, error)
	ServicesGameServerDeploymentNamespaceListerExpansion
}

// servicesGameServerDeploymentNamespaceLister implements the ServicesGameServerDeploymentNamespaceLister
// interface.
type servicesGameServerDeploymentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServicesGameServerDeployments in the indexer for a given namespace.
func (s servicesGameServerDeploymentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServicesGameServerDeployment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServicesGameServerDeployment))
	})
	return ret, err
}

// Get retrieves the ServicesGameServerDeployment from the indexer for a given namespace and name.
func (s servicesGameServerDeploymentNamespaceLister) Get(name string) (*v1alpha1.ServicesGameServerDeployment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicesgameserverdeployment"), name)
	}
	return obj.(*v1alpha1.ServicesGameServerDeployment), nil
}
