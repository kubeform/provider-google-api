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
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Databases returns a DatabaseInformer.
	Databases() DatabaseInformer
	// DatabaseInstances returns a DatabaseInstanceInformer.
	DatabaseInstances() DatabaseInstanceInformer
	// SourceRepresentationInstances returns a SourceRepresentationInstanceInformer.
	SourceRepresentationInstances() SourceRepresentationInstanceInformer
	// SslCerts returns a SslCertInformer.
	SslCerts() SslCertInformer
	// Users returns a UserInformer.
	Users() UserInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Databases returns a DatabaseInformer.
func (v *version) Databases() DatabaseInformer {
	return &databaseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DatabaseInstances returns a DatabaseInstanceInformer.
func (v *version) DatabaseInstances() DatabaseInstanceInformer {
	return &databaseInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SourceRepresentationInstances returns a SourceRepresentationInstanceInformer.
func (v *version) SourceRepresentationInstances() SourceRepresentationInstanceInformer {
	return &sourceRepresentationInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SslCerts returns a SslCertInformer.
func (v *version) SslCerts() SslCertInformer {
	return &sslCertInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Users returns a UserInformer.
func (v *version) Users() UserInformer {
	return &userInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
