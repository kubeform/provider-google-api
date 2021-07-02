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

// RouterPeerLister helps list RouterPeers.
// All objects returned here must be treated as read-only.
type RouterPeerLister interface {
	// List lists all RouterPeers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RouterPeer, err error)
	// RouterPeers returns an object that can list and get RouterPeers.
	RouterPeers(namespace string) RouterPeerNamespaceLister
	RouterPeerListerExpansion
}

// routerPeerLister implements the RouterPeerLister interface.
type routerPeerLister struct {
	indexer cache.Indexer
}

// NewRouterPeerLister returns a new RouterPeerLister.
func NewRouterPeerLister(indexer cache.Indexer) RouterPeerLister {
	return &routerPeerLister{indexer: indexer}
}

// List lists all RouterPeers in the indexer.
func (s *routerPeerLister) List(selector labels.Selector) (ret []*v1alpha1.RouterPeer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RouterPeer))
	})
	return ret, err
}

// RouterPeers returns an object that can list and get RouterPeers.
func (s *routerPeerLister) RouterPeers(namespace string) RouterPeerNamespaceLister {
	return routerPeerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RouterPeerNamespaceLister helps list and get RouterPeers.
// All objects returned here must be treated as read-only.
type RouterPeerNamespaceLister interface {
	// List lists all RouterPeers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RouterPeer, err error)
	// Get retrieves the RouterPeer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RouterPeer, error)
	RouterPeerNamespaceListerExpansion
}

// routerPeerNamespaceLister implements the RouterPeerNamespaceLister
// interface.
type routerPeerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RouterPeers in the indexer for a given namespace.
func (s routerPeerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RouterPeer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RouterPeer))
	})
	return ret, err
}

// Get retrieves the RouterPeer from the indexer for a given namespace and name.
func (s routerPeerNamespaceLister) Get(name string) (*v1alpha1.RouterPeer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("routerpeer"), name)
	}
	return obj.(*v1alpha1.RouterPeer), nil
}
