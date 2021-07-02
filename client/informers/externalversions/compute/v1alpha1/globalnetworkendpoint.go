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

	computev1alpha1 "kubeform.dev/provider-google-api/apis/compute/v1alpha1"
	versioned "kubeform.dev/provider-google-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-google-api/client/listers/compute/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GlobalNetworkEndpointInformer provides access to a shared informer and lister for
// GlobalNetworkEndpoints.
type GlobalNetworkEndpointInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.GlobalNetworkEndpointLister
}

type globalNetworkEndpointInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGlobalNetworkEndpointInformer constructs a new informer for GlobalNetworkEndpoint type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGlobalNetworkEndpointInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGlobalNetworkEndpointInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGlobalNetworkEndpointInformer constructs a new informer for GlobalNetworkEndpoint type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGlobalNetworkEndpointInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ComputeV1alpha1().GlobalNetworkEndpoints(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ComputeV1alpha1().GlobalNetworkEndpoints(namespace).Watch(context.TODO(), options)
			},
		},
		&computev1alpha1.GlobalNetworkEndpoint{},
		resyncPeriod,
		indexers,
	)
}

func (f *globalNetworkEndpointInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGlobalNetworkEndpointInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *globalNetworkEndpointInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&computev1alpha1.GlobalNetworkEndpoint{}, f.defaultInformer)
}

func (f *globalNetworkEndpointInformer) Lister() v1alpha1.GlobalNetworkEndpointLister {
	return v1alpha1.NewGlobalNetworkEndpointLister(f.Informer().GetIndexer())
}
