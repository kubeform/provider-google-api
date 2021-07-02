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

	dataprocv1alpha1 "kubeform.dev/provider-google-api/apis/dataproc/v1alpha1"
	versioned "kubeform.dev/provider-google-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-google-api/client/listers/dataproc/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterIamPolicyInformer provides access to a shared informer and lister for
// ClusterIamPolicies.
type ClusterIamPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterIamPolicyLister
}

type clusterIamPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterIamPolicyInformer constructs a new informer for ClusterIamPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterIamPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterIamPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterIamPolicyInformer constructs a new informer for ClusterIamPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterIamPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DataprocV1alpha1().ClusterIamPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DataprocV1alpha1().ClusterIamPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&dataprocv1alpha1.ClusterIamPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterIamPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterIamPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterIamPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&dataprocv1alpha1.ClusterIamPolicy{}, f.defaultInformer)
}

func (f *clusterIamPolicyInformer) Lister() v1alpha1.ClusterIamPolicyLister {
	return v1alpha1.NewClusterIamPolicyLister(f.Informer().GetIndexer())
}
