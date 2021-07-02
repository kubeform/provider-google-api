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

	folderv1alpha1 "kubeform.dev/provider-google-api/apis/folder/v1alpha1"
	versioned "kubeform.dev/provider-google-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-google-api/client/listers/folder/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FolderInformer provides access to a shared informer and lister for
// Folders.
type FolderInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FolderLister
}

type folderInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFolderInformer constructs a new informer for Folder type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFolderInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFolderInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFolderInformer constructs a new informer for Folder type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFolderInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FolderV1alpha1().Folders(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FolderV1alpha1().Folders(namespace).Watch(context.TODO(), options)
			},
		},
		&folderv1alpha1.Folder{},
		resyncPeriod,
		indexers,
	)
}

func (f *folderInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFolderInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *folderInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&folderv1alpha1.Folder{}, f.defaultInformer)
}

func (f *folderInformer) Lister() v1alpha1.FolderLister {
	return v1alpha1.NewFolderLister(f.Informer().GetIndexer())
}
