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
	gamev1alpha1 "kubeform.dev/provider-google-api/apis/game/v1alpha1"
	versioned "kubeform.dev/provider-google-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-google-api/client/listers/game/v1alpha1"
)

// ServicesGameServerDeploymentRolloutInformer provides access to a shared informer and lister for
// ServicesGameServerDeploymentRollouts.
type ServicesGameServerDeploymentRolloutInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ServicesGameServerDeploymentRolloutLister
}

type servicesGameServerDeploymentRolloutInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServicesGameServerDeploymentRolloutInformer constructs a new informer for ServicesGameServerDeploymentRollout type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServicesGameServerDeploymentRolloutInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServicesGameServerDeploymentRolloutInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServicesGameServerDeploymentRolloutInformer constructs a new informer for ServicesGameServerDeploymentRollout type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServicesGameServerDeploymentRolloutInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GameV1alpha1().ServicesGameServerDeploymentRollouts(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GameV1alpha1().ServicesGameServerDeploymentRollouts(namespace).Watch(context.TODO(), options)
			},
		},
		&gamev1alpha1.ServicesGameServerDeploymentRollout{},
		resyncPeriod,
		indexers,
	)
}

func (f *servicesGameServerDeploymentRolloutInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServicesGameServerDeploymentRolloutInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *servicesGameServerDeploymentRolloutInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&gamev1alpha1.ServicesGameServerDeploymentRollout{}, f.defaultInformer)
}

func (f *servicesGameServerDeploymentRolloutInformer) Lister() v1alpha1.ServicesGameServerDeploymentRolloutLister {
	return v1alpha1.NewServicesGameServerDeploymentRolloutLister(f.Informer().GetIndexer())
}
