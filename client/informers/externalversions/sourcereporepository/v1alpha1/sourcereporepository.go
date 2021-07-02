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

	sourcereporepositoryv1alpha1 "kubeform.dev/provider-google-api/apis/sourcereporepository/v1alpha1"
	versioned "kubeform.dev/provider-google-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-google-api/client/listers/sourcereporepository/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SourcerepoRepositoryInformer provides access to a shared informer and lister for
// SourcerepoRepositories.
type SourcerepoRepositoryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SourcerepoRepositoryLister
}

type sourcerepoRepositoryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSourcerepoRepositoryInformer constructs a new informer for SourcerepoRepository type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSourcerepoRepositoryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSourcerepoRepositoryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSourcerepoRepositoryInformer constructs a new informer for SourcerepoRepository type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSourcerepoRepositoryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SourcereporepositoryV1alpha1().SourcerepoRepositories(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SourcereporepositoryV1alpha1().SourcerepoRepositories(namespace).Watch(context.TODO(), options)
			},
		},
		&sourcereporepositoryv1alpha1.SourcerepoRepository{},
		resyncPeriod,
		indexers,
	)
}

func (f *sourcerepoRepositoryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSourcerepoRepositoryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sourcerepoRepositoryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sourcereporepositoryv1alpha1.SourcerepoRepository{}, f.defaultInformer)
}

func (f *sourcerepoRepositoryInformer) Lister() v1alpha1.SourcerepoRepositoryLister {
	return v1alpha1.NewSourcerepoRepositoryLister(f.Informer().GetIndexer())
}
