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

package externalversions

import (
	reflect "reflect"
	sync "sync"
	time "time"

	versioned "kubeform.dev/provider-google-api/client/clientset/versioned"
	accesscontext "kubeform.dev/provider-google-api/client/informers/externalversions/accesscontext"
	active "kubeform.dev/provider-google-api/client/informers/externalversions/active"
	apigee "kubeform.dev/provider-google-api/client/informers/externalversions/apigee"
	appengine "kubeform.dev/provider-google-api/client/informers/externalversions/appengine"
	assured "kubeform.dev/provider-google-api/client/informers/externalversions/assured"
	bigquery "kubeform.dev/provider-google-api/client/informers/externalversions/bigquery"
	bigtable "kubeform.dev/provider-google-api/client/informers/externalversions/bigtable"
	billingaccount "kubeform.dev/provider-google-api/client/informers/externalversions/billingaccount"
	billingbudget "kubeform.dev/provider-google-api/client/informers/externalversions/billingbudget"
	billingsubaccount "kubeform.dev/provider-google-api/client/informers/externalversions/billingsubaccount"
	binaryauthorization "kubeform.dev/provider-google-api/client/informers/externalversions/binaryauthorization"
	cloud "kubeform.dev/provider-google-api/client/informers/externalversions/cloud"
	cloudbuild "kubeform.dev/provider-google-api/client/informers/externalversions/cloudbuild"
	cloudfunctionsfunction "kubeform.dev/provider-google-api/client/informers/externalversions/cloudfunctionsfunction"
	cloudiot "kubeform.dev/provider-google-api/client/informers/externalversions/cloudiot"
	composer "kubeform.dev/provider-google-api/client/informers/externalversions/composer"
	compute "kubeform.dev/provider-google-api/client/informers/externalversions/compute"
	container "kubeform.dev/provider-google-api/client/informers/externalversions/container"
	data "kubeform.dev/provider-google-api/client/informers/externalversions/data"
	dataflow "kubeform.dev/provider-google-api/client/informers/externalversions/dataflow"
	dataproc "kubeform.dev/provider-google-api/client/informers/externalversions/dataproc"
	datastore "kubeform.dev/provider-google-api/client/informers/externalversions/datastore"
	deployment "kubeform.dev/provider-google-api/client/informers/externalversions/deployment"
	dialogflow "kubeform.dev/provider-google-api/client/informers/externalversions/dialogflow"
	dns "kubeform.dev/provider-google-api/client/informers/externalversions/dns"
	endpoints "kubeform.dev/provider-google-api/client/informers/externalversions/endpoints"
	essential "kubeform.dev/provider-google-api/client/informers/externalversions/essential"
	eventarc "kubeform.dev/provider-google-api/client/informers/externalversions/eventarc"
	filestore "kubeform.dev/provider-google-api/client/informers/externalversions/filestore"
	firestore "kubeform.dev/provider-google-api/client/informers/externalversions/firestore"
	folder "kubeform.dev/provider-google-api/client/informers/externalversions/folder"
	game "kubeform.dev/provider-google-api/client/informers/externalversions/game"
	gke "kubeform.dev/provider-google-api/client/informers/externalversions/gke"
	healthcare "kubeform.dev/provider-google-api/client/informers/externalversions/healthcare"
	iap "kubeform.dev/provider-google-api/client/informers/externalversions/iap"
	identity "kubeform.dev/provider-google-api/client/informers/externalversions/identity"
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
	kms "kubeform.dev/provider-google-api/client/informers/externalversions/kms"
	logging "kubeform.dev/provider-google-api/client/informers/externalversions/logging"
	memcache "kubeform.dev/provider-google-api/client/informers/externalversions/memcache"
	ml "kubeform.dev/provider-google-api/client/informers/externalversions/ml"
	monitoring "kubeform.dev/provider-google-api/client/informers/externalversions/monitoring"
	network "kubeform.dev/provider-google-api/client/informers/externalversions/network"
	notebooks "kubeform.dev/provider-google-api/client/informers/externalversions/notebooks"
	org "kubeform.dev/provider-google-api/client/informers/externalversions/org"
	organization "kubeform.dev/provider-google-api/client/informers/externalversions/organization"
	os "kubeform.dev/provider-google-api/client/informers/externalversions/os"
	privateca "kubeform.dev/provider-google-api/client/informers/externalversions/privateca"
	project "kubeform.dev/provider-google-api/client/informers/externalversions/project"
	pubsub "kubeform.dev/provider-google-api/client/informers/externalversions/pubsub"
	redis "kubeform.dev/provider-google-api/client/informers/externalversions/redis"
	resource "kubeform.dev/provider-google-api/client/informers/externalversions/resource"
	runtimeconfig "kubeform.dev/provider-google-api/client/informers/externalversions/runtimeconfig"
	scc "kubeform.dev/provider-google-api/client/informers/externalversions/scc"
	secret "kubeform.dev/provider-google-api/client/informers/externalversions/secret"
	service "kubeform.dev/provider-google-api/client/informers/externalversions/service"
	sourcereporepository "kubeform.dev/provider-google-api/client/informers/externalversions/sourcereporepository"
	spanner "kubeform.dev/provider-google-api/client/informers/externalversions/spanner"
	sql "kubeform.dev/provider-google-api/client/informers/externalversions/sql"
	storage "kubeform.dev/provider-google-api/client/informers/externalversions/storage"
	tags "kubeform.dev/provider-google-api/client/informers/externalversions/tags"
	tpu "kubeform.dev/provider-google-api/client/informers/externalversions/tpu"
	vertex "kubeform.dev/provider-google-api/client/informers/externalversions/vertex"
	vpc "kubeform.dev/provider-google-api/client/informers/externalversions/vpc"
	workflows "kubeform.dev/provider-google-api/client/informers/externalversions/workflows"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// SharedInformerOption defines the functional option type for SharedInformerFactory.
type SharedInformerOption func(*sharedInformerFactory) *sharedInformerFactory

type sharedInformerFactory struct {
	client           versioned.Interface
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	lock             sync.Mutex
	defaultResync    time.Duration
	customResync     map[reflect.Type]time.Duration

	informers map[reflect.Type]cache.SharedIndexInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[reflect.Type]bool
}

// WithCustomResyncConfig sets a custom resync period for the specified informer types.
func WithCustomResyncConfig(resyncConfig map[v1.Object]time.Duration) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		for k, v := range resyncConfig {
			factory.customResync[reflect.TypeOf(k)] = v
		}
		return factory
	}
}

// WithTweakListOptions sets a custom filter on all listers of the configured SharedInformerFactory.
func WithTweakListOptions(tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		factory.tweakListOptions = tweakListOptions
		return factory
	}
}

// WithNamespace limits the SharedInformerFactory to the specified namespace.
func WithNamespace(namespace string) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		factory.namespace = namespace
		return factory
	}
}

// NewSharedInformerFactory constructs a new instance of sharedInformerFactory for all namespaces.
func NewSharedInformerFactory(client versioned.Interface, defaultResync time.Duration) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync)
}

// NewFilteredSharedInformerFactory constructs a new instance of sharedInformerFactory.
// Listers obtained via this SharedInformerFactory will be subject to the same filters
// as specified here.
// Deprecated: Please use NewSharedInformerFactoryWithOptions instead
func NewFilteredSharedInformerFactory(client versioned.Interface, defaultResync time.Duration, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync, WithNamespace(namespace), WithTweakListOptions(tweakListOptions))
}

// NewSharedInformerFactoryWithOptions constructs a new instance of a SharedInformerFactory with additional options.
func NewSharedInformerFactoryWithOptions(client versioned.Interface, defaultResync time.Duration, options ...SharedInformerOption) SharedInformerFactory {
	factory := &sharedInformerFactory{
		client:           client,
		namespace:        v1.NamespaceAll,
		defaultResync:    defaultResync,
		informers:        make(map[reflect.Type]cache.SharedIndexInformer),
		startedInformers: make(map[reflect.Type]bool),
		customResync:     make(map[reflect.Type]time.Duration),
	}

	// Apply all options
	for _, opt := range options {
		factory = opt(factory)
	}

	return factory
}

// Start initializes all requested informers.
func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.lock.Lock()
	defer f.lock.Unlock()

	for informerType, informer := range f.informers {
		if !f.startedInformers[informerType] {
			go informer.Run(stopCh)
			f.startedInformers[informerType] = true
		}
	}
}

// WaitForCacheSync waits for all started informers' cache were synced.
func (f *sharedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	informers := func() map[reflect.Type]cache.SharedIndexInformer {
		f.lock.Lock()
		defer f.lock.Unlock()

		informers := map[reflect.Type]cache.SharedIndexInformer{}
		for informerType, informer := range f.informers {
			if f.startedInformers[informerType] {
				informers[informerType] = informer
			}
		}
		return informers
	}()

	res := map[reflect.Type]bool{}
	for informType, informer := range informers {
		res[informType] = cache.WaitForCacheSync(stopCh, informer.HasSynced)
	}
	return res
}

// InternalInformerFor returns the SharedIndexInformer for obj using an internal
// client.
func (f *sharedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(obj)
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}

	resyncPeriod, exists := f.customResync[informerType]
	if !exists {
		resyncPeriod = f.defaultResync
	}

	informer = newFunc(f.client, resyncPeriod)
	f.informers[informerType] = informer

	return informer
}

// SharedInformerFactory provides shared informers for resources in all known
// API group versions.
type SharedInformerFactory interface {
	internalinterfaces.SharedInformerFactory
	ForResource(resource schema.GroupVersionResource) (GenericInformer, error)
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool

	Accesscontext() accesscontext.Interface
	Active() active.Interface
	Apigee() apigee.Interface
	Appengine() appengine.Interface
	Assured() assured.Interface
	Bigquery() bigquery.Interface
	Bigtable() bigtable.Interface
	Billingaccount() billingaccount.Interface
	Billingbudget() billingbudget.Interface
	Billingsubaccount() billingsubaccount.Interface
	Binaryauthorization() binaryauthorization.Interface
	Cloud() cloud.Interface
	Cloudbuild() cloudbuild.Interface
	Cloudfunctionsfunction() cloudfunctionsfunction.Interface
	Cloudiot() cloudiot.Interface
	Composer() composer.Interface
	Compute() compute.Interface
	Container() container.Interface
	Data() data.Interface
	Dataflow() dataflow.Interface
	Dataproc() dataproc.Interface
	Datastore() datastore.Interface
	Deployment() deployment.Interface
	Dialogflow() dialogflow.Interface
	Dns() dns.Interface
	Endpoints() endpoints.Interface
	Essential() essential.Interface
	Eventarc() eventarc.Interface
	Filestore() filestore.Interface
	Firestore() firestore.Interface
	Folder() folder.Interface
	Game() game.Interface
	Gke() gke.Interface
	Healthcare() healthcare.Interface
	Iap() iap.Interface
	Identity() identity.Interface
	Kms() kms.Interface
	Logging() logging.Interface
	Memcache() memcache.Interface
	Ml() ml.Interface
	Monitoring() monitoring.Interface
	Network() network.Interface
	Notebooks() notebooks.Interface
	Org() org.Interface
	Organization() organization.Interface
	Os() os.Interface
	Privateca() privateca.Interface
	Project() project.Interface
	Pubsub() pubsub.Interface
	Redis() redis.Interface
	Resource() resource.Interface
	Runtimeconfig() runtimeconfig.Interface
	Scc() scc.Interface
	Secret() secret.Interface
	Service() service.Interface
	Sourcereporepository() sourcereporepository.Interface
	Spanner() spanner.Interface
	Sql() sql.Interface
	Storage() storage.Interface
	Tags() tags.Interface
	Tpu() tpu.Interface
	Vertex() vertex.Interface
	Vpc() vpc.Interface
	Workflows() workflows.Interface
}

func (f *sharedInformerFactory) Accesscontext() accesscontext.Interface {
	return accesscontext.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Active() active.Interface {
	return active.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Apigee() apigee.Interface {
	return apigee.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Appengine() appengine.Interface {
	return appengine.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Assured() assured.Interface {
	return assured.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Bigquery() bigquery.Interface {
	return bigquery.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Bigtable() bigtable.Interface {
	return bigtable.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Billingaccount() billingaccount.Interface {
	return billingaccount.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Billingbudget() billingbudget.Interface {
	return billingbudget.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Billingsubaccount() billingsubaccount.Interface {
	return billingsubaccount.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Binaryauthorization() binaryauthorization.Interface {
	return binaryauthorization.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cloud() cloud.Interface {
	return cloud.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cloudbuild() cloudbuild.Interface {
	return cloudbuild.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cloudfunctionsfunction() cloudfunctionsfunction.Interface {
	return cloudfunctionsfunction.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cloudiot() cloudiot.Interface {
	return cloudiot.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Composer() composer.Interface {
	return composer.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Compute() compute.Interface {
	return compute.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Container() container.Interface {
	return container.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Data() data.Interface {
	return data.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Dataflow() dataflow.Interface {
	return dataflow.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Dataproc() dataproc.Interface {
	return dataproc.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Datastore() datastore.Interface {
	return datastore.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Deployment() deployment.Interface {
	return deployment.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Dialogflow() dialogflow.Interface {
	return dialogflow.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Dns() dns.Interface {
	return dns.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Endpoints() endpoints.Interface {
	return endpoints.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Essential() essential.Interface {
	return essential.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Eventarc() eventarc.Interface {
	return eventarc.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Filestore() filestore.Interface {
	return filestore.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Firestore() firestore.Interface {
	return firestore.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Folder() folder.Interface {
	return folder.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Game() game.Interface {
	return game.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Gke() gke.Interface {
	return gke.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Healthcare() healthcare.Interface {
	return healthcare.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Iap() iap.Interface {
	return iap.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Identity() identity.Interface {
	return identity.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Kms() kms.Interface {
	return kms.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Logging() logging.Interface {
	return logging.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Memcache() memcache.Interface {
	return memcache.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ml() ml.Interface {
	return ml.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Monitoring() monitoring.Interface {
	return monitoring.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Network() network.Interface {
	return network.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Notebooks() notebooks.Interface {
	return notebooks.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Org() org.Interface {
	return org.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Organization() organization.Interface {
	return organization.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Os() os.Interface {
	return os.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Privateca() privateca.Interface {
	return privateca.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Project() project.Interface {
	return project.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Pubsub() pubsub.Interface {
	return pubsub.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Redis() redis.Interface {
	return redis.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Resource() resource.Interface {
	return resource.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Runtimeconfig() runtimeconfig.Interface {
	return runtimeconfig.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Scc() scc.Interface {
	return scc.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Secret() secret.Interface {
	return secret.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Service() service.Interface {
	return service.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Sourcereporepository() sourcereporepository.Interface {
	return sourcereporepository.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Spanner() spanner.Interface {
	return spanner.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Sql() sql.Interface {
	return sql.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Storage() storage.Interface {
	return storage.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Tags() tags.Interface {
	return tags.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Tpu() tpu.Interface {
	return tpu.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Vertex() vertex.Interface {
	return vertex.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Vpc() vpc.Interface {
	return vpc.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Workflows() workflows.Interface {
	return workflows.New(f, f.namespace, f.tweakListOptions)
}
