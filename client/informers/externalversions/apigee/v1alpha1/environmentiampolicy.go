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

	apigeev1alpha1 "kubeform.dev/provider-google-api/apis/apigee/v1alpha1"
	versioned "kubeform.dev/provider-google-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-google-api/client/listers/apigee/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EnvironmentIamPolicyInformer provides access to a shared informer and lister for
// EnvironmentIamPolicies.
type EnvironmentIamPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.EnvironmentIamPolicyLister
}

type environmentIamPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEnvironmentIamPolicyInformer constructs a new informer for EnvironmentIamPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEnvironmentIamPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEnvironmentIamPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEnvironmentIamPolicyInformer constructs a new informer for EnvironmentIamPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEnvironmentIamPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApigeeV1alpha1().EnvironmentIamPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApigeeV1alpha1().EnvironmentIamPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&apigeev1alpha1.EnvironmentIamPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *environmentIamPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEnvironmentIamPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *environmentIamPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apigeev1alpha1.EnvironmentIamPolicy{}, f.defaultInformer)
}

func (f *environmentIamPolicyInformer) Lister() v1alpha1.EnvironmentIamPolicyLister {
	return v1alpha1.NewEnvironmentIamPolicyLister(f.Informer().GetIndexer())
}
