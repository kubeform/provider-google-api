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

	endpointsv1alpha1 "kubeform.dev/provider-google-api/apis/endpoints/v1alpha1"
	versioned "kubeform.dev/provider-google-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-google-api/client/listers/endpoints/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceIamPolicyInformer provides access to a shared informer and lister for
// ServiceIamPolicies.
type ServiceIamPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ServiceIamPolicyLister
}

type serviceIamPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceIamPolicyInformer constructs a new informer for ServiceIamPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceIamPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceIamPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceIamPolicyInformer constructs a new informer for ServiceIamPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceIamPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EndpointsV1alpha1().ServiceIamPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EndpointsV1alpha1().ServiceIamPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&endpointsv1alpha1.ServiceIamPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceIamPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceIamPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceIamPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&endpointsv1alpha1.ServiceIamPolicy{}, f.defaultInformer)
}

func (f *serviceIamPolicyInformer) Lister() v1alpha1.ServiceIamPolicyLister {
	return v1alpha1.NewServiceIamPolicyLister(f.Informer().GetIndexer())
}
