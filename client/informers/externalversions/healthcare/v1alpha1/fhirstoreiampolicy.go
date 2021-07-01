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
	healthcarev1alpha1 "kubeform.dev/provider-google-api/apis/healthcare/v1alpha1"
	versioned "kubeform.dev/provider-google-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-google-api/client/listers/healthcare/v1alpha1"
)

// FhirStoreIamPolicyInformer provides access to a shared informer and lister for
// FhirStoreIamPolicies.
type FhirStoreIamPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FhirStoreIamPolicyLister
}

type fhirStoreIamPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFhirStoreIamPolicyInformer constructs a new informer for FhirStoreIamPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFhirStoreIamPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFhirStoreIamPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFhirStoreIamPolicyInformer constructs a new informer for FhirStoreIamPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFhirStoreIamPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HealthcareV1alpha1().FhirStoreIamPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HealthcareV1alpha1().FhirStoreIamPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&healthcarev1alpha1.FhirStoreIamPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *fhirStoreIamPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFhirStoreIamPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *fhirStoreIamPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&healthcarev1alpha1.FhirStoreIamPolicy{}, f.defaultInformer)
}

func (f *fhirStoreIamPolicyInformer) Lister() v1alpha1.FhirStoreIamPolicyLister {
	return v1alpha1.NewFhirStoreIamPolicyLister(f.Informer().GetIndexer())
}
