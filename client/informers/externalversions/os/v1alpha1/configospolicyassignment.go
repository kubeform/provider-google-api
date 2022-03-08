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

	osv1alpha1 "kubeform.dev/provider-google-api/apis/os/v1alpha1"
	versioned "kubeform.dev/provider-google-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-google-api/client/listers/os/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ConfigOsPolicyAssignmentInformer provides access to a shared informer and lister for
// ConfigOsPolicyAssignments.
type ConfigOsPolicyAssignmentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ConfigOsPolicyAssignmentLister
}

type configOsPolicyAssignmentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewConfigOsPolicyAssignmentInformer constructs a new informer for ConfigOsPolicyAssignment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewConfigOsPolicyAssignmentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredConfigOsPolicyAssignmentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredConfigOsPolicyAssignmentInformer constructs a new informer for ConfigOsPolicyAssignment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredConfigOsPolicyAssignmentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OsV1alpha1().ConfigOsPolicyAssignments(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OsV1alpha1().ConfigOsPolicyAssignments(namespace).Watch(context.TODO(), options)
			},
		},
		&osv1alpha1.ConfigOsPolicyAssignment{},
		resyncPeriod,
		indexers,
	)
}

func (f *configOsPolicyAssignmentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredConfigOsPolicyAssignmentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *configOsPolicyAssignmentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&osv1alpha1.ConfigOsPolicyAssignment{}, f.defaultInformer)
}

func (f *configOsPolicyAssignmentInformer) Lister() v1alpha1.ConfigOsPolicyAssignmentLister {
	return v1alpha1.NewConfigOsPolicyAssignmentLister(f.Informer().GetIndexer())
}