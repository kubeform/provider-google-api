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

	cloudv1alpha1 "kubeform.dev/provider-google-api/apis/cloud/v1alpha1"
	versioned "kubeform.dev/provider-google-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-google-api/client/listers/cloud/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RunServiceIamBindingInformer provides access to a shared informer and lister for
// RunServiceIamBindings.
type RunServiceIamBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.RunServiceIamBindingLister
}

type runServiceIamBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRunServiceIamBindingInformer constructs a new informer for RunServiceIamBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRunServiceIamBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRunServiceIamBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRunServiceIamBindingInformer constructs a new informer for RunServiceIamBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRunServiceIamBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudV1alpha1().RunServiceIamBindings(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudV1alpha1().RunServiceIamBindings(namespace).Watch(context.TODO(), options)
			},
		},
		&cloudv1alpha1.RunServiceIamBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *runServiceIamBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRunServiceIamBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *runServiceIamBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cloudv1alpha1.RunServiceIamBinding{}, f.defaultInformer)
}

func (f *runServiceIamBindingInformer) Lister() v1alpha1.RunServiceIamBindingLister {
	return v1alpha1.NewRunServiceIamBindingLister(f.Informer().GetIndexer())
}
