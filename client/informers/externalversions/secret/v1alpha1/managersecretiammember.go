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
	secretv1alpha1 "kubeform.dev/provider-google-api/apis/secret/v1alpha1"
	versioned "kubeform.dev/provider-google-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-google-api/client/listers/secret/v1alpha1"
)

// ManagerSecretIamMemberInformer provides access to a shared informer and lister for
// ManagerSecretIamMembers.
type ManagerSecretIamMemberInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ManagerSecretIamMemberLister
}

type managerSecretIamMemberInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewManagerSecretIamMemberInformer constructs a new informer for ManagerSecretIamMember type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewManagerSecretIamMemberInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredManagerSecretIamMemberInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredManagerSecretIamMemberInformer constructs a new informer for ManagerSecretIamMember type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredManagerSecretIamMemberInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecretV1alpha1().ManagerSecretIamMembers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecretV1alpha1().ManagerSecretIamMembers(namespace).Watch(context.TODO(), options)
			},
		},
		&secretv1alpha1.ManagerSecretIamMember{},
		resyncPeriod,
		indexers,
	)
}

func (f *managerSecretIamMemberInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredManagerSecretIamMemberInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *managerSecretIamMemberInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&secretv1alpha1.ManagerSecretIamMember{}, f.defaultInformer)
}

func (f *managerSecretIamMemberInformer) Lister() v1alpha1.ManagerSecretIamMemberLister {
	return v1alpha1.NewManagerSecretIamMemberLister(f.Informer().GetIndexer())
}
