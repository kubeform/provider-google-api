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

// CatalogTagTemplateIamBindingInformer provides access to a shared informer and lister for
// CatalogTagTemplateIamBindings.
type CatalogTagTemplateIamBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CatalogTagTemplateIamBindingLister
}

type catalogTagTemplateIamBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCatalogTagTemplateIamBindingInformer constructs a new informer for CatalogTagTemplateIamBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCatalogTagTemplateIamBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCatalogTagTemplateIamBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCatalogTagTemplateIamBindingInformer constructs a new informer for CatalogTagTemplateIamBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCatalogTagTemplateIamBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DataV1alpha1().CatalogTagTemplateIamBindings(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DataV1alpha1().CatalogTagTemplateIamBindings(namespace).Watch(context.TODO(), options)
			},
		},
		&datav1alpha1.CatalogTagTemplateIamBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *catalogTagTemplateIamBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCatalogTagTemplateIamBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *catalogTagTemplateIamBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&datav1alpha1.CatalogTagTemplateIamBinding{}, f.defaultInformer)
}

func (f *catalogTagTemplateIamBindingInformer) Lister() v1alpha1.CatalogTagTemplateIamBindingLister {
	return v1alpha1.NewCatalogTagTemplateIamBindingLister(f.Informer().GetIndexer())
}
