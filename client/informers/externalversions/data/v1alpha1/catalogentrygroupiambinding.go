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
	datav1alpha1 "kubeform.dev/provider-google-api/apis/data/v1alpha1"
	versioned "kubeform.dev/provider-google-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-google-api/client/listers/data/v1alpha1"
)

// CatalogEntryGroupIamBindingInformer provides access to a shared informer and lister for
// CatalogEntryGroupIamBindings.
type CatalogEntryGroupIamBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CatalogEntryGroupIamBindingLister
}

type catalogEntryGroupIamBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCatalogEntryGroupIamBindingInformer constructs a new informer for CatalogEntryGroupIamBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCatalogEntryGroupIamBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCatalogEntryGroupIamBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCatalogEntryGroupIamBindingInformer constructs a new informer for CatalogEntryGroupIamBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCatalogEntryGroupIamBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DataV1alpha1().CatalogEntryGroupIamBindings(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DataV1alpha1().CatalogEntryGroupIamBindings(namespace).Watch(context.TODO(), options)
			},
		},
		&datav1alpha1.CatalogEntryGroupIamBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *catalogEntryGroupIamBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCatalogEntryGroupIamBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *catalogEntryGroupIamBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&datav1alpha1.CatalogEntryGroupIamBinding{}, f.defaultInformer)
}

func (f *catalogEntryGroupIamBindingInformer) Lister() v1alpha1.CatalogEntryGroupIamBindingLister {
	return v1alpha1.NewCatalogEntryGroupIamBindingLister(f.Informer().GetIndexer())
}
