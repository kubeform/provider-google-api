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

	iapv1alpha1 "kubeform.dev/provider-google-api/apis/iap/v1alpha1"
	versioned "kubeform.dev/provider-google-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-google-api/client/listers/iap/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TunnelIamBindingInformer provides access to a shared informer and lister for
// TunnelIamBindings.
type TunnelIamBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TunnelIamBindingLister
}

type tunnelIamBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTunnelIamBindingInformer constructs a new informer for TunnelIamBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTunnelIamBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTunnelIamBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTunnelIamBindingInformer constructs a new informer for TunnelIamBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTunnelIamBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IapV1alpha1().TunnelIamBindings(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IapV1alpha1().TunnelIamBindings(namespace).Watch(context.TODO(), options)
			},
		},
		&iapv1alpha1.TunnelIamBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *tunnelIamBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTunnelIamBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tunnelIamBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&iapv1alpha1.TunnelIamBinding{}, f.defaultInformer)
}

func (f *tunnelIamBindingInformer) Lister() v1alpha1.TunnelIamBindingLister {
	return v1alpha1.NewTunnelIamBindingLister(f.Informer().GetIndexer())
}
