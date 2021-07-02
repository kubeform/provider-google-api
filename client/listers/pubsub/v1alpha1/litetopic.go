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
	v1alpha1 "kubeform.dev/provider-google-api/apis/pubsub/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LiteTopicLister helps list LiteTopics.
// All objects returned here must be treated as read-only.
type LiteTopicLister interface {
	// List lists all LiteTopics in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LiteTopic, err error)
	// LiteTopics returns an object that can list and get LiteTopics.
	LiteTopics(namespace string) LiteTopicNamespaceLister
	LiteTopicListerExpansion
}

// liteTopicLister implements the LiteTopicLister interface.
type liteTopicLister struct {
	indexer cache.Indexer
}

// NewLiteTopicLister returns a new LiteTopicLister.
func NewLiteTopicLister(indexer cache.Indexer) LiteTopicLister {
	return &liteTopicLister{indexer: indexer}
}

// List lists all LiteTopics in the indexer.
func (s *liteTopicLister) List(selector labels.Selector) (ret []*v1alpha1.LiteTopic, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LiteTopic))
	})
	return ret, err
}

// LiteTopics returns an object that can list and get LiteTopics.
func (s *liteTopicLister) LiteTopics(namespace string) LiteTopicNamespaceLister {
	return liteTopicNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LiteTopicNamespaceLister helps list and get LiteTopics.
// All objects returned here must be treated as read-only.
type LiteTopicNamespaceLister interface {
	// List lists all LiteTopics in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LiteTopic, err error)
	// Get retrieves the LiteTopic from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LiteTopic, error)
	LiteTopicNamespaceListerExpansion
}

// liteTopicNamespaceLister implements the LiteTopicNamespaceLister
// interface.
type liteTopicNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LiteTopics in the indexer for a given namespace.
func (s liteTopicNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LiteTopic, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LiteTopic))
	})
	return ret, err
}

// Get retrieves the LiteTopic from the indexer for a given namespace and name.
func (s liteTopicNamespaceLister) Get(name string) (*v1alpha1.LiteTopic, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("litetopic"), name)
	}
	return obj.(*v1alpha1.LiteTopic), nil
}
