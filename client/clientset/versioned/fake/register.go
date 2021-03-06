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
	accesscontextv1alpha1 "kubeform.dev/provider-google-api/apis/accesscontext/v1alpha1"
	activev1alpha1 "kubeform.dev/provider-google-api/apis/active/v1alpha1"
	apigeev1alpha1 "kubeform.dev/provider-google-api/apis/apigee/v1alpha1"
	appenginev1alpha1 "kubeform.dev/provider-google-api/apis/appengine/v1alpha1"
	assuredv1alpha1 "kubeform.dev/provider-google-api/apis/assured/v1alpha1"
	bigqueryv1alpha1 "kubeform.dev/provider-google-api/apis/bigquery/v1alpha1"
	bigtablev1alpha1 "kubeform.dev/provider-google-api/apis/bigtable/v1alpha1"
	billingaccountv1alpha1 "kubeform.dev/provider-google-api/apis/billingaccount/v1alpha1"
	billingbudgetv1alpha1 "kubeform.dev/provider-google-api/apis/billingbudget/v1alpha1"
	billingsubaccountv1alpha1 "kubeform.dev/provider-google-api/apis/billingsubaccount/v1alpha1"
	binaryauthorizationv1alpha1 "kubeform.dev/provider-google-api/apis/binaryauthorization/v1alpha1"
	cloudv1alpha1 "kubeform.dev/provider-google-api/apis/cloud/v1alpha1"
	cloudbuildv1alpha1 "kubeform.dev/provider-google-api/apis/cloudbuild/v1alpha1"
	cloudfunctionsfunctionv1alpha1 "kubeform.dev/provider-google-api/apis/cloudfunctionsfunction/v1alpha1"
	cloudiotv1alpha1 "kubeform.dev/provider-google-api/apis/cloudiot/v1alpha1"
	composerv1alpha1 "kubeform.dev/provider-google-api/apis/composer/v1alpha1"
	computev1alpha1 "kubeform.dev/provider-google-api/apis/compute/v1alpha1"
	containerv1alpha1 "kubeform.dev/provider-google-api/apis/container/v1alpha1"
	datav1alpha1 "kubeform.dev/provider-google-api/apis/data/v1alpha1"
	dataflowv1alpha1 "kubeform.dev/provider-google-api/apis/dataflow/v1alpha1"
	dataprocv1alpha1 "kubeform.dev/provider-google-api/apis/dataproc/v1alpha1"
	datastorev1alpha1 "kubeform.dev/provider-google-api/apis/datastore/v1alpha1"
	deploymentv1alpha1 "kubeform.dev/provider-google-api/apis/deployment/v1alpha1"
	dialogflowv1alpha1 "kubeform.dev/provider-google-api/apis/dialogflow/v1alpha1"
	dnsv1alpha1 "kubeform.dev/provider-google-api/apis/dns/v1alpha1"
	endpointsv1alpha1 "kubeform.dev/provider-google-api/apis/endpoints/v1alpha1"
	essentialv1alpha1 "kubeform.dev/provider-google-api/apis/essential/v1alpha1"
	eventarcv1alpha1 "kubeform.dev/provider-google-api/apis/eventarc/v1alpha1"
	filestorev1alpha1 "kubeform.dev/provider-google-api/apis/filestore/v1alpha1"
	firestorev1alpha1 "kubeform.dev/provider-google-api/apis/firestore/v1alpha1"
	folderv1alpha1 "kubeform.dev/provider-google-api/apis/folder/v1alpha1"
	gamev1alpha1 "kubeform.dev/provider-google-api/apis/game/v1alpha1"
	gkev1alpha1 "kubeform.dev/provider-google-api/apis/gke/v1alpha1"
	healthcarev1alpha1 "kubeform.dev/provider-google-api/apis/healthcare/v1alpha1"
	iapv1alpha1 "kubeform.dev/provider-google-api/apis/iap/v1alpha1"
	identityv1alpha1 "kubeform.dev/provider-google-api/apis/identity/v1alpha1"
	kmsv1alpha1 "kubeform.dev/provider-google-api/apis/kms/v1alpha1"
	loggingv1alpha1 "kubeform.dev/provider-google-api/apis/logging/v1alpha1"
	memcachev1alpha1 "kubeform.dev/provider-google-api/apis/memcache/v1alpha1"
	mlv1alpha1 "kubeform.dev/provider-google-api/apis/ml/v1alpha1"
	monitoringv1alpha1 "kubeform.dev/provider-google-api/apis/monitoring/v1alpha1"
	networkv1alpha1 "kubeform.dev/provider-google-api/apis/network/v1alpha1"
	notebooksv1alpha1 "kubeform.dev/provider-google-api/apis/notebooks/v1alpha1"
	orgv1alpha1 "kubeform.dev/provider-google-api/apis/org/v1alpha1"
	organizationv1alpha1 "kubeform.dev/provider-google-api/apis/organization/v1alpha1"
	osv1alpha1 "kubeform.dev/provider-google-api/apis/os/v1alpha1"
	privatecav1alpha1 "kubeform.dev/provider-google-api/apis/privateca/v1alpha1"
	projectv1alpha1 "kubeform.dev/provider-google-api/apis/project/v1alpha1"
	pubsubv1alpha1 "kubeform.dev/provider-google-api/apis/pubsub/v1alpha1"
	recaptchav1alpha1 "kubeform.dev/provider-google-api/apis/recaptcha/v1alpha1"
	redisv1alpha1 "kubeform.dev/provider-google-api/apis/redis/v1alpha1"
	resourcev1alpha1 "kubeform.dev/provider-google-api/apis/resource/v1alpha1"
	sccv1alpha1 "kubeform.dev/provider-google-api/apis/scc/v1alpha1"
	secretv1alpha1 "kubeform.dev/provider-google-api/apis/secret/v1alpha1"
	servicev1alpha1 "kubeform.dev/provider-google-api/apis/service/v1alpha1"
	sourcereporepositoryv1alpha1 "kubeform.dev/provider-google-api/apis/sourcereporepository/v1alpha1"
	spannerv1alpha1 "kubeform.dev/provider-google-api/apis/spanner/v1alpha1"
	sqlv1alpha1 "kubeform.dev/provider-google-api/apis/sql/v1alpha1"
	storagev1alpha1 "kubeform.dev/provider-google-api/apis/storage/v1alpha1"
	tagsv1alpha1 "kubeform.dev/provider-google-api/apis/tags/v1alpha1"
	tpuv1alpha1 "kubeform.dev/provider-google-api/apis/tpu/v1alpha1"
	vertexv1alpha1 "kubeform.dev/provider-google-api/apis/vertex/v1alpha1"
	vpcv1alpha1 "kubeform.dev/provider-google-api/apis/vpc/v1alpha1"
	workflowsv1alpha1 "kubeform.dev/provider-google-api/apis/workflows/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)

var localSchemeBuilder = runtime.SchemeBuilder{
	accesscontextv1alpha1.AddToScheme,
	activev1alpha1.AddToScheme,
	apigeev1alpha1.AddToScheme,
	appenginev1alpha1.AddToScheme,
	assuredv1alpha1.AddToScheme,
	bigqueryv1alpha1.AddToScheme,
	bigtablev1alpha1.AddToScheme,
	billingaccountv1alpha1.AddToScheme,
	billingbudgetv1alpha1.AddToScheme,
	billingsubaccountv1alpha1.AddToScheme,
	binaryauthorizationv1alpha1.AddToScheme,
	cloudv1alpha1.AddToScheme,
	cloudbuildv1alpha1.AddToScheme,
	cloudfunctionsfunctionv1alpha1.AddToScheme,
	cloudiotv1alpha1.AddToScheme,
	composerv1alpha1.AddToScheme,
	computev1alpha1.AddToScheme,
	containerv1alpha1.AddToScheme,
	datav1alpha1.AddToScheme,
	dataflowv1alpha1.AddToScheme,
	dataprocv1alpha1.AddToScheme,
	datastorev1alpha1.AddToScheme,
	deploymentv1alpha1.AddToScheme,
	dialogflowv1alpha1.AddToScheme,
	dnsv1alpha1.AddToScheme,
	endpointsv1alpha1.AddToScheme,
	essentialv1alpha1.AddToScheme,
	eventarcv1alpha1.AddToScheme,
	filestorev1alpha1.AddToScheme,
	firestorev1alpha1.AddToScheme,
	folderv1alpha1.AddToScheme,
	gamev1alpha1.AddToScheme,
	gkev1alpha1.AddToScheme,
	healthcarev1alpha1.AddToScheme,
	iapv1alpha1.AddToScheme,
	identityv1alpha1.AddToScheme,
	kmsv1alpha1.AddToScheme,
	loggingv1alpha1.AddToScheme,
	memcachev1alpha1.AddToScheme,
	mlv1alpha1.AddToScheme,
	monitoringv1alpha1.AddToScheme,
	networkv1alpha1.AddToScheme,
	notebooksv1alpha1.AddToScheme,
	orgv1alpha1.AddToScheme,
	organizationv1alpha1.AddToScheme,
	osv1alpha1.AddToScheme,
	privatecav1alpha1.AddToScheme,
	projectv1alpha1.AddToScheme,
	pubsubv1alpha1.AddToScheme,
	recaptchav1alpha1.AddToScheme,
	redisv1alpha1.AddToScheme,
	resourcev1alpha1.AddToScheme,
	sccv1alpha1.AddToScheme,
	secretv1alpha1.AddToScheme,
	servicev1alpha1.AddToScheme,
	sourcereporepositoryv1alpha1.AddToScheme,
	spannerv1alpha1.AddToScheme,
	sqlv1alpha1.AddToScheme,
	storagev1alpha1.AddToScheme,
	tagsv1alpha1.AddToScheme,
	tpuv1alpha1.AddToScheme,
	vertexv1alpha1.AddToScheme,
	vpcv1alpha1.AddToScheme,
	workflowsv1alpha1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   _ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(scheme))
}
