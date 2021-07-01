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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	pubsubv1alpha1 "kubeform.dev/provider-google-api/apis/pubsub/v1alpha1"
	versioned "kubeform.dev/provider-google-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-google-api/client/listers/pubsub/v1alpha1"
)

// LiteTopicInformer provides access to a shared informer and lister for
// LiteTopics.
type LiteTopicInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.LiteTopicLister
}

type liteTopicInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLiteTopicInformer constructs a new informer for LiteTopic type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLiteTopicInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLiteTopicInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLiteTopicInformer constructs a new informer for LiteTopic type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLiteTopicInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PubsubV1alpha1().LiteTopics(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PubsubV1alpha1().LiteTopics(namespace).Watch(context.TODO(), options)
			},
		},
		&pubsubv1alpha1.LiteTopic{},
		resyncPeriod,
		indexers,
	)
}

func (f *liteTopicInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLiteTopicInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *liteTopicInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&pubsubv1alpha1.LiteTopic{}, f.defaultInformer)
}

func (f *liteTopicInformer) Lister() v1alpha1.LiteTopicLister {
	return v1alpha1.NewLiteTopicLister(f.Informer().GetIndexer())
}
