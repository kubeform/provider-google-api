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
	v1alpha1 "kubeform.dev/provider-google-api/apis/game/v1alpha1"
)

// ServicesGameServerDeploymentRolloutLister helps list ServicesGameServerDeploymentRollouts.
// All objects returned here must be treated as read-only.
type ServicesGameServerDeploymentRolloutLister interface {
	// List lists all ServicesGameServerDeploymentRollouts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServicesGameServerDeploymentRollout, err error)
	// ServicesGameServerDeploymentRollouts returns an object that can list and get ServicesGameServerDeploymentRollouts.
	ServicesGameServerDeploymentRollouts(namespace string) ServicesGameServerDeploymentRolloutNamespaceLister
	ServicesGameServerDeploymentRolloutListerExpansion
}

// servicesGameServerDeploymentRolloutLister implements the ServicesGameServerDeploymentRolloutLister interface.
type servicesGameServerDeploymentRolloutLister struct {
	indexer cache.Indexer
}

// NewServicesGameServerDeploymentRolloutLister returns a new ServicesGameServerDeploymentRolloutLister.
func NewServicesGameServerDeploymentRolloutLister(indexer cache.Indexer) ServicesGameServerDeploymentRolloutLister {
	return &servicesGameServerDeploymentRolloutLister{indexer: indexer}
}

// List lists all ServicesGameServerDeploymentRollouts in the indexer.
func (s *servicesGameServerDeploymentRolloutLister) List(selector labels.Selector) (ret []*v1alpha1.ServicesGameServerDeploymentRollout, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServicesGameServerDeploymentRollout))
	})
	return ret, err
}

// ServicesGameServerDeploymentRollouts returns an object that can list and get ServicesGameServerDeploymentRollouts.
func (s *servicesGameServerDeploymentRolloutLister) ServicesGameServerDeploymentRollouts(namespace string) ServicesGameServerDeploymentRolloutNamespaceLister {
	return servicesGameServerDeploymentRolloutNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServicesGameServerDeploymentRolloutNamespaceLister helps list and get ServicesGameServerDeploymentRollouts.
// All objects returned here must be treated as read-only.
type ServicesGameServerDeploymentRolloutNamespaceLister interface {
	// List lists all ServicesGameServerDeploymentRollouts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServicesGameServerDeploymentRollout, err error)
	// Get retrieves the ServicesGameServerDeploymentRollout from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServicesGameServerDeploymentRollout, error)
	ServicesGameServerDeploymentRolloutNamespaceListerExpansion
}

// servicesGameServerDeploymentRolloutNamespaceLister implements the ServicesGameServerDeploymentRolloutNamespaceLister
// interface.
type servicesGameServerDeploymentRolloutNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServicesGameServerDeploymentRollouts in the indexer for a given namespace.
func (s servicesGameServerDeploymentRolloutNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServicesGameServerDeploymentRollout, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServicesGameServerDeploymentRollout))
	})
	return ret, err
}

// Get retrieves the ServicesGameServerDeploymentRollout from the indexer for a given namespace and name.
func (s servicesGameServerDeploymentRolloutNamespaceLister) Get(name string) (*v1alpha1.ServicesGameServerDeploymentRollout, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicesgameserverdeploymentrollout"), name)
	}
	return obj.(*v1alpha1.ServicesGameServerDeploymentRollout), nil
}
