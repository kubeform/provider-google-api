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

// SubnetworkIamBindingInformer provides access to a shared informer and lister for
// SubnetworkIamBindings.
type SubnetworkIamBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SubnetworkIamBindingLister
}

type subnetworkIamBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSubnetworkIamBindingInformer constructs a new informer for SubnetworkIamBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSubnetworkIamBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSubnetworkIamBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSubnetworkIamBindingInformer constructs a new informer for SubnetworkIamBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSubnetworkIamBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ComputeV1alpha1().SubnetworkIamBindings(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ComputeV1alpha1().SubnetworkIamBindings(namespace).Watch(context.TODO(), options)
			},
		},
		&computev1alpha1.SubnetworkIamBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *subnetworkIamBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSubnetworkIamBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *subnetworkIamBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&computev1alpha1.SubnetworkIamBinding{}, f.defaultInformer)
}

func (f *subnetworkIamBindingInformer) Lister() v1alpha1.SubnetworkIamBindingLister {
	return v1alpha1.NewSubnetworkIamBindingLister(f.Informer().GetIndexer())
}
