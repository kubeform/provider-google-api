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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
	clientset "kubeform.dev/provider-google-api/client/clientset/versioned"
	accesscontextv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/accesscontext/v1alpha1"
	fakeaccesscontextv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/accesscontext/v1alpha1/fake"
	activev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/active/v1alpha1"
	fakeactivev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/active/v1alpha1/fake"
	apigeev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/apigee/v1alpha1"
	fakeapigeev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/apigee/v1alpha1/fake"
	appenginev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/appengine/v1alpha1"
	fakeappenginev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/appengine/v1alpha1/fake"
	bigqueryv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/bigquery/v1alpha1"
	fakebigqueryv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/bigquery/v1alpha1/fake"
	bigtablev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/bigtable/v1alpha1"
	fakebigtablev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/bigtable/v1alpha1/fake"
	billingaccountv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/billingaccount/v1alpha1"
	fakebillingaccountv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/billingaccount/v1alpha1/fake"
	billingbudgetv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/billingbudget/v1alpha1"
	fakebillingbudgetv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/billingbudget/v1alpha1/fake"
	billingsubaccountv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/billingsubaccount/v1alpha1"
	fakebillingsubaccountv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/billingsubaccount/v1alpha1/fake"
	binaryauthorizationv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/binaryauthorization/v1alpha1"
	fakebinaryauthorizationv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/binaryauthorization/v1alpha1/fake"
	cloudv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/cloud/v1alpha1"
	fakecloudv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/cloud/v1alpha1/fake"
	cloudbuildv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/cloudbuild/v1alpha1"
	fakecloudbuildv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/cloudbuild/v1alpha1/fake"
	cloudfunctionsfunctionv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/cloudfunctionsfunction/v1alpha1"
	fakecloudfunctionsfunctionv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/cloudfunctionsfunction/v1alpha1/fake"
	cloudiotv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/cloudiot/v1alpha1"
	fakecloudiotv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/cloudiot/v1alpha1/fake"
	composerv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/composer/v1alpha1"
	fakecomposerv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/composer/v1alpha1/fake"
	computev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/compute/v1alpha1"
	fakecomputev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/compute/v1alpha1/fake"
	containerv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/container/v1alpha1"
	fakecontainerv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/container/v1alpha1/fake"
	datav1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/data/v1alpha1"
	fakedatav1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/data/v1alpha1/fake"
	dataflowv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/dataflow/v1alpha1"
	fakedataflowv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/dataflow/v1alpha1/fake"
	dataprocv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/dataproc/v1alpha1"
	fakedataprocv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/dataproc/v1alpha1/fake"
	datastorev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/datastore/v1alpha1"
	fakedatastorev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/datastore/v1alpha1/fake"
	deploymentv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/deployment/v1alpha1"
	fakedeploymentv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/deployment/v1alpha1/fake"
	dialogflowv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/dialogflow/v1alpha1"
	fakedialogflowv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/dialogflow/v1alpha1/fake"
	dnsv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/dns/v1alpha1"
	fakednsv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/dns/v1alpha1/fake"
	endpointsv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/endpoints/v1alpha1"
	fakeendpointsv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/endpoints/v1alpha1/fake"
	eventarcv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/eventarc/v1alpha1"
	fakeeventarcv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/eventarc/v1alpha1/fake"
	filestorev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/filestore/v1alpha1"
	fakefilestorev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/filestore/v1alpha1/fake"
	firestorev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/firestore/v1alpha1"
	fakefirestorev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/firestore/v1alpha1/fake"
	folderv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/folder/v1alpha1"
	fakefolderv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/folder/v1alpha1/fake"
	gamev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/game/v1alpha1"
	fakegamev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/game/v1alpha1/fake"
	healthcarev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/healthcare/v1alpha1"
	fakehealthcarev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/healthcare/v1alpha1/fake"
	iapv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/iap/v1alpha1"
	fakeiapv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/iap/v1alpha1/fake"
	identityv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/identity/v1alpha1"
	fakeidentityv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/identity/v1alpha1/fake"
	kmsv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/kms/v1alpha1"
	fakekmsv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/kms/v1alpha1/fake"
	loggingv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/logging/v1alpha1"
	fakeloggingv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/logging/v1alpha1/fake"
	memcachev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/memcache/v1alpha1"
	fakememcachev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/memcache/v1alpha1/fake"
	mlv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/ml/v1alpha1"
	fakemlv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/ml/v1alpha1/fake"
	monitoringv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/monitoring/v1alpha1"
	fakemonitoringv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/monitoring/v1alpha1/fake"
	networkv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/network/v1alpha1"
	fakenetworkv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/network/v1alpha1/fake"
	notebooksv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/notebooks/v1alpha1"
	fakenotebooksv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/notebooks/v1alpha1/fake"
	organizationv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/organization/v1alpha1"
	fakeorganizationv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/organization/v1alpha1/fake"
	osv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/os/v1alpha1"
	fakeosv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/os/v1alpha1/fake"
	projectv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/project/v1alpha1"
	fakeprojectv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/project/v1alpha1/fake"
	pubsubv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/pubsub/v1alpha1"
	fakepubsubv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/pubsub/v1alpha1/fake"
	redisv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/redis/v1alpha1"
	fakeredisv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/redis/v1alpha1/fake"
	resourcev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/resource/v1alpha1"
	fakeresourcev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/resource/v1alpha1/fake"
	runtimeconfigv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/runtimeconfig/v1alpha1"
	fakeruntimeconfigv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/runtimeconfig/v1alpha1/fake"
	sccv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/scc/v1alpha1"
	fakesccv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/scc/v1alpha1/fake"
	secretv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/secret/v1alpha1"
	fakesecretv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/secret/v1alpha1/fake"
	servicev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/service/v1alpha1"
	fakeservicev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/service/v1alpha1/fake"
	sourcereporepositoryv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/sourcereporepository/v1alpha1"
	fakesourcereporepositoryv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/sourcereporepository/v1alpha1/fake"
	spannerv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/spanner/v1alpha1"
	fakespannerv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/spanner/v1alpha1/fake"
	sqlv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/sql/v1alpha1"
	fakesqlv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/sql/v1alpha1/fake"
	storagev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/storage/v1alpha1"
	fakestoragev1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/storage/v1alpha1/fake"
	tagsv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/tags/v1alpha1"
	faketagsv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/tags/v1alpha1/fake"
	tpuv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/tpu/v1alpha1"
	faketpuv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/tpu/v1alpha1/fake"
	vertexv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/vertex/v1alpha1"
	fakevertexv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/vertex/v1alpha1/fake"
	vpcv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/vpc/v1alpha1"
	fakevpcv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/vpc/v1alpha1/fake"
	workflowsv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/workflows/v1alpha1"
	fakeworkflowsv1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/workflows/v1alpha1/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// AccesscontextV1alpha1 retrieves the AccesscontextV1alpha1Client
func (c *Clientset) AccesscontextV1alpha1() accesscontextv1alpha1.AccesscontextV1alpha1Interface {
	return &fakeaccesscontextv1alpha1.FakeAccesscontextV1alpha1{Fake: &c.Fake}
}

// ActiveV1alpha1 retrieves the ActiveV1alpha1Client
func (c *Clientset) ActiveV1alpha1() activev1alpha1.ActiveV1alpha1Interface {
	return &fakeactivev1alpha1.FakeActiveV1alpha1{Fake: &c.Fake}
}

// ApigeeV1alpha1 retrieves the ApigeeV1alpha1Client
func (c *Clientset) ApigeeV1alpha1() apigeev1alpha1.ApigeeV1alpha1Interface {
	return &fakeapigeev1alpha1.FakeApigeeV1alpha1{Fake: &c.Fake}
}

// AppengineV1alpha1 retrieves the AppengineV1alpha1Client
func (c *Clientset) AppengineV1alpha1() appenginev1alpha1.AppengineV1alpha1Interface {
	return &fakeappenginev1alpha1.FakeAppengineV1alpha1{Fake: &c.Fake}
}

// BigqueryV1alpha1 retrieves the BigqueryV1alpha1Client
func (c *Clientset) BigqueryV1alpha1() bigqueryv1alpha1.BigqueryV1alpha1Interface {
	return &fakebigqueryv1alpha1.FakeBigqueryV1alpha1{Fake: &c.Fake}
}

// BigtableV1alpha1 retrieves the BigtableV1alpha1Client
func (c *Clientset) BigtableV1alpha1() bigtablev1alpha1.BigtableV1alpha1Interface {
	return &fakebigtablev1alpha1.FakeBigtableV1alpha1{Fake: &c.Fake}
}

// BillingaccountV1alpha1 retrieves the BillingaccountV1alpha1Client
func (c *Clientset) BillingaccountV1alpha1() billingaccountv1alpha1.BillingaccountV1alpha1Interface {
	return &fakebillingaccountv1alpha1.FakeBillingaccountV1alpha1{Fake: &c.Fake}
}

// BillingbudgetV1alpha1 retrieves the BillingbudgetV1alpha1Client
func (c *Clientset) BillingbudgetV1alpha1() billingbudgetv1alpha1.BillingbudgetV1alpha1Interface {
	return &fakebillingbudgetv1alpha1.FakeBillingbudgetV1alpha1{Fake: &c.Fake}
}

// BillingsubaccountV1alpha1 retrieves the BillingsubaccountV1alpha1Client
func (c *Clientset) BillingsubaccountV1alpha1() billingsubaccountv1alpha1.BillingsubaccountV1alpha1Interface {
	return &fakebillingsubaccountv1alpha1.FakeBillingsubaccountV1alpha1{Fake: &c.Fake}
}

// BinaryauthorizationV1alpha1 retrieves the BinaryauthorizationV1alpha1Client
func (c *Clientset) BinaryauthorizationV1alpha1() binaryauthorizationv1alpha1.BinaryauthorizationV1alpha1Interface {
	return &fakebinaryauthorizationv1alpha1.FakeBinaryauthorizationV1alpha1{Fake: &c.Fake}
}

// CloudV1alpha1 retrieves the CloudV1alpha1Client
func (c *Clientset) CloudV1alpha1() cloudv1alpha1.CloudV1alpha1Interface {
	return &fakecloudv1alpha1.FakeCloudV1alpha1{Fake: &c.Fake}
}

// CloudbuildV1alpha1 retrieves the CloudbuildV1alpha1Client
func (c *Clientset) CloudbuildV1alpha1() cloudbuildv1alpha1.CloudbuildV1alpha1Interface {
	return &fakecloudbuildv1alpha1.FakeCloudbuildV1alpha1{Fake: &c.Fake}
}

// CloudfunctionsfunctionV1alpha1 retrieves the CloudfunctionsfunctionV1alpha1Client
func (c *Clientset) CloudfunctionsfunctionV1alpha1() cloudfunctionsfunctionv1alpha1.CloudfunctionsfunctionV1alpha1Interface {
	return &fakecloudfunctionsfunctionv1alpha1.FakeCloudfunctionsfunctionV1alpha1{Fake: &c.Fake}
}

// CloudiotV1alpha1 retrieves the CloudiotV1alpha1Client
func (c *Clientset) CloudiotV1alpha1() cloudiotv1alpha1.CloudiotV1alpha1Interface {
	return &fakecloudiotv1alpha1.FakeCloudiotV1alpha1{Fake: &c.Fake}
}

// ComposerV1alpha1 retrieves the ComposerV1alpha1Client
func (c *Clientset) ComposerV1alpha1() composerv1alpha1.ComposerV1alpha1Interface {
	return &fakecomposerv1alpha1.FakeComposerV1alpha1{Fake: &c.Fake}
}

// ComputeV1alpha1 retrieves the ComputeV1alpha1Client
func (c *Clientset) ComputeV1alpha1() computev1alpha1.ComputeV1alpha1Interface {
	return &fakecomputev1alpha1.FakeComputeV1alpha1{Fake: &c.Fake}
}

// ContainerV1alpha1 retrieves the ContainerV1alpha1Client
func (c *Clientset) ContainerV1alpha1() containerv1alpha1.ContainerV1alpha1Interface {
	return &fakecontainerv1alpha1.FakeContainerV1alpha1{Fake: &c.Fake}
}

// DataV1alpha1 retrieves the DataV1alpha1Client
func (c *Clientset) DataV1alpha1() datav1alpha1.DataV1alpha1Interface {
	return &fakedatav1alpha1.FakeDataV1alpha1{Fake: &c.Fake}
}

// DataflowV1alpha1 retrieves the DataflowV1alpha1Client
func (c *Clientset) DataflowV1alpha1() dataflowv1alpha1.DataflowV1alpha1Interface {
	return &fakedataflowv1alpha1.FakeDataflowV1alpha1{Fake: &c.Fake}
}

// DataprocV1alpha1 retrieves the DataprocV1alpha1Client
func (c *Clientset) DataprocV1alpha1() dataprocv1alpha1.DataprocV1alpha1Interface {
	return &fakedataprocv1alpha1.FakeDataprocV1alpha1{Fake: &c.Fake}
}

// DatastoreV1alpha1 retrieves the DatastoreV1alpha1Client
func (c *Clientset) DatastoreV1alpha1() datastorev1alpha1.DatastoreV1alpha1Interface {
	return &fakedatastorev1alpha1.FakeDatastoreV1alpha1{Fake: &c.Fake}
}

// DeploymentV1alpha1 retrieves the DeploymentV1alpha1Client
func (c *Clientset) DeploymentV1alpha1() deploymentv1alpha1.DeploymentV1alpha1Interface {
	return &fakedeploymentv1alpha1.FakeDeploymentV1alpha1{Fake: &c.Fake}
}

// DialogflowV1alpha1 retrieves the DialogflowV1alpha1Client
func (c *Clientset) DialogflowV1alpha1() dialogflowv1alpha1.DialogflowV1alpha1Interface {
	return &fakedialogflowv1alpha1.FakeDialogflowV1alpha1{Fake: &c.Fake}
}

// DnsV1alpha1 retrieves the DnsV1alpha1Client
func (c *Clientset) DnsV1alpha1() dnsv1alpha1.DnsV1alpha1Interface {
	return &fakednsv1alpha1.FakeDnsV1alpha1{Fake: &c.Fake}
}

// EndpointsV1alpha1 retrieves the EndpointsV1alpha1Client
func (c *Clientset) EndpointsV1alpha1() endpointsv1alpha1.EndpointsV1alpha1Interface {
	return &fakeendpointsv1alpha1.FakeEndpointsV1alpha1{Fake: &c.Fake}
}

// EventarcV1alpha1 retrieves the EventarcV1alpha1Client
func (c *Clientset) EventarcV1alpha1() eventarcv1alpha1.EventarcV1alpha1Interface {
	return &fakeeventarcv1alpha1.FakeEventarcV1alpha1{Fake: &c.Fake}
}

// FilestoreV1alpha1 retrieves the FilestoreV1alpha1Client
func (c *Clientset) FilestoreV1alpha1() filestorev1alpha1.FilestoreV1alpha1Interface {
	return &fakefilestorev1alpha1.FakeFilestoreV1alpha1{Fake: &c.Fake}
}

// FirestoreV1alpha1 retrieves the FirestoreV1alpha1Client
func (c *Clientset) FirestoreV1alpha1() firestorev1alpha1.FirestoreV1alpha1Interface {
	return &fakefirestorev1alpha1.FakeFirestoreV1alpha1{Fake: &c.Fake}
}

// FolderV1alpha1 retrieves the FolderV1alpha1Client
func (c *Clientset) FolderV1alpha1() folderv1alpha1.FolderV1alpha1Interface {
	return &fakefolderv1alpha1.FakeFolderV1alpha1{Fake: &c.Fake}
}

// GameV1alpha1 retrieves the GameV1alpha1Client
func (c *Clientset) GameV1alpha1() gamev1alpha1.GameV1alpha1Interface {
	return &fakegamev1alpha1.FakeGameV1alpha1{Fake: &c.Fake}
}

// HealthcareV1alpha1 retrieves the HealthcareV1alpha1Client
func (c *Clientset) HealthcareV1alpha1() healthcarev1alpha1.HealthcareV1alpha1Interface {
	return &fakehealthcarev1alpha1.FakeHealthcareV1alpha1{Fake: &c.Fake}
}

// IapV1alpha1 retrieves the IapV1alpha1Client
func (c *Clientset) IapV1alpha1() iapv1alpha1.IapV1alpha1Interface {
	return &fakeiapv1alpha1.FakeIapV1alpha1{Fake: &c.Fake}
}

// IdentityV1alpha1 retrieves the IdentityV1alpha1Client
func (c *Clientset) IdentityV1alpha1() identityv1alpha1.IdentityV1alpha1Interface {
	return &fakeidentityv1alpha1.FakeIdentityV1alpha1{Fake: &c.Fake}
}

// KmsV1alpha1 retrieves the KmsV1alpha1Client
func (c *Clientset) KmsV1alpha1() kmsv1alpha1.KmsV1alpha1Interface {
	return &fakekmsv1alpha1.FakeKmsV1alpha1{Fake: &c.Fake}
}

// LoggingV1alpha1 retrieves the LoggingV1alpha1Client
func (c *Clientset) LoggingV1alpha1() loggingv1alpha1.LoggingV1alpha1Interface {
	return &fakeloggingv1alpha1.FakeLoggingV1alpha1{Fake: &c.Fake}
}

// MemcacheV1alpha1 retrieves the MemcacheV1alpha1Client
func (c *Clientset) MemcacheV1alpha1() memcachev1alpha1.MemcacheV1alpha1Interface {
	return &fakememcachev1alpha1.FakeMemcacheV1alpha1{Fake: &c.Fake}
}

// MlV1alpha1 retrieves the MlV1alpha1Client
func (c *Clientset) MlV1alpha1() mlv1alpha1.MlV1alpha1Interface {
	return &fakemlv1alpha1.FakeMlV1alpha1{Fake: &c.Fake}
}

// MonitoringV1alpha1 retrieves the MonitoringV1alpha1Client
func (c *Clientset) MonitoringV1alpha1() monitoringv1alpha1.MonitoringV1alpha1Interface {
	return &fakemonitoringv1alpha1.FakeMonitoringV1alpha1{Fake: &c.Fake}
}

// NetworkV1alpha1 retrieves the NetworkV1alpha1Client
func (c *Clientset) NetworkV1alpha1() networkv1alpha1.NetworkV1alpha1Interface {
	return &fakenetworkv1alpha1.FakeNetworkV1alpha1{Fake: &c.Fake}
}

// NotebooksV1alpha1 retrieves the NotebooksV1alpha1Client
func (c *Clientset) NotebooksV1alpha1() notebooksv1alpha1.NotebooksV1alpha1Interface {
	return &fakenotebooksv1alpha1.FakeNotebooksV1alpha1{Fake: &c.Fake}
}

// OrganizationV1alpha1 retrieves the OrganizationV1alpha1Client
func (c *Clientset) OrganizationV1alpha1() organizationv1alpha1.OrganizationV1alpha1Interface {
	return &fakeorganizationv1alpha1.FakeOrganizationV1alpha1{Fake: &c.Fake}
}

// OsV1alpha1 retrieves the OsV1alpha1Client
func (c *Clientset) OsV1alpha1() osv1alpha1.OsV1alpha1Interface {
	return &fakeosv1alpha1.FakeOsV1alpha1{Fake: &c.Fake}
}

// ProjectV1alpha1 retrieves the ProjectV1alpha1Client
func (c *Clientset) ProjectV1alpha1() projectv1alpha1.ProjectV1alpha1Interface {
	return &fakeprojectv1alpha1.FakeProjectV1alpha1{Fake: &c.Fake}
}

// PubsubV1alpha1 retrieves the PubsubV1alpha1Client
func (c *Clientset) PubsubV1alpha1() pubsubv1alpha1.PubsubV1alpha1Interface {
	return &fakepubsubv1alpha1.FakePubsubV1alpha1{Fake: &c.Fake}
}

// RedisV1alpha1 retrieves the RedisV1alpha1Client
func (c *Clientset) RedisV1alpha1() redisv1alpha1.RedisV1alpha1Interface {
	return &fakeredisv1alpha1.FakeRedisV1alpha1{Fake: &c.Fake}
}

// ResourceV1alpha1 retrieves the ResourceV1alpha1Client
func (c *Clientset) ResourceV1alpha1() resourcev1alpha1.ResourceV1alpha1Interface {
	return &fakeresourcev1alpha1.FakeResourceV1alpha1{Fake: &c.Fake}
}

// RuntimeconfigV1alpha1 retrieves the RuntimeconfigV1alpha1Client
func (c *Clientset) RuntimeconfigV1alpha1() runtimeconfigv1alpha1.RuntimeconfigV1alpha1Interface {
	return &fakeruntimeconfigv1alpha1.FakeRuntimeconfigV1alpha1{Fake: &c.Fake}
}

// SccV1alpha1 retrieves the SccV1alpha1Client
func (c *Clientset) SccV1alpha1() sccv1alpha1.SccV1alpha1Interface {
	return &fakesccv1alpha1.FakeSccV1alpha1{Fake: &c.Fake}
}

// SecretV1alpha1 retrieves the SecretV1alpha1Client
func (c *Clientset) SecretV1alpha1() secretv1alpha1.SecretV1alpha1Interface {
	return &fakesecretv1alpha1.FakeSecretV1alpha1{Fake: &c.Fake}
}

// ServiceV1alpha1 retrieves the ServiceV1alpha1Client
func (c *Clientset) ServiceV1alpha1() servicev1alpha1.ServiceV1alpha1Interface {
	return &fakeservicev1alpha1.FakeServiceV1alpha1{Fake: &c.Fake}
}

// SourcereporepositoryV1alpha1 retrieves the SourcereporepositoryV1alpha1Client
func (c *Clientset) SourcereporepositoryV1alpha1() sourcereporepositoryv1alpha1.SourcereporepositoryV1alpha1Interface {
	return &fakesourcereporepositoryv1alpha1.FakeSourcereporepositoryV1alpha1{Fake: &c.Fake}
}

// SpannerV1alpha1 retrieves the SpannerV1alpha1Client
func (c *Clientset) SpannerV1alpha1() spannerv1alpha1.SpannerV1alpha1Interface {
	return &fakespannerv1alpha1.FakeSpannerV1alpha1{Fake: &c.Fake}
}

// SqlV1alpha1 retrieves the SqlV1alpha1Client
func (c *Clientset) SqlV1alpha1() sqlv1alpha1.SqlV1alpha1Interface {
	return &fakesqlv1alpha1.FakeSqlV1alpha1{Fake: &c.Fake}
}

// StorageV1alpha1 retrieves the StorageV1alpha1Client
func (c *Clientset) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	return &fakestoragev1alpha1.FakeStorageV1alpha1{Fake: &c.Fake}
}

// TagsV1alpha1 retrieves the TagsV1alpha1Client
func (c *Clientset) TagsV1alpha1() tagsv1alpha1.TagsV1alpha1Interface {
	return &faketagsv1alpha1.FakeTagsV1alpha1{Fake: &c.Fake}
}

// TpuV1alpha1 retrieves the TpuV1alpha1Client
func (c *Clientset) TpuV1alpha1() tpuv1alpha1.TpuV1alpha1Interface {
	return &faketpuv1alpha1.FakeTpuV1alpha1{Fake: &c.Fake}
}

// VertexV1alpha1 retrieves the VertexV1alpha1Client
func (c *Clientset) VertexV1alpha1() vertexv1alpha1.VertexV1alpha1Interface {
	return &fakevertexv1alpha1.FakeVertexV1alpha1{Fake: &c.Fake}
}

// VpcV1alpha1 retrieves the VpcV1alpha1Client
func (c *Clientset) VpcV1alpha1() vpcv1alpha1.VpcV1alpha1Interface {
	return &fakevpcv1alpha1.FakeVpcV1alpha1{Fake: &c.Fake}
}

// WorkflowsV1alpha1 retrieves the WorkflowsV1alpha1Client
func (c *Clientset) WorkflowsV1alpha1() workflowsv1alpha1.WorkflowsV1alpha1Interface {
	return &fakeworkflowsv1alpha1.FakeWorkflowsV1alpha1{Fake: &c.Fake}
}
