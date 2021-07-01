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
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/monitoring/v1alpha1"
)

type FakeMonitoringV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMonitoringV1alpha1) AlertPolicies(namespace string) v1alpha1.AlertPolicyInterface {
	return &FakeAlertPolicies{c, namespace}
}

func (c *FakeMonitoringV1alpha1) CustomServices(namespace string) v1alpha1.CustomServiceInterface {
	return &FakeCustomServices{c, namespace}
}

func (c *FakeMonitoringV1alpha1) Dashboards(namespace string) v1alpha1.DashboardInterface {
	return &FakeDashboards{c, namespace}
}

func (c *FakeMonitoringV1alpha1) Groups(namespace string) v1alpha1.GroupInterface {
	return &FakeGroups{c, namespace}
}

func (c *FakeMonitoringV1alpha1) MetricDescriptors(namespace string) v1alpha1.MetricDescriptorInterface {
	return &FakeMetricDescriptors{c, namespace}
}

func (c *FakeMonitoringV1alpha1) NotificationChannels(namespace string) v1alpha1.NotificationChannelInterface {
	return &FakeNotificationChannels{c, namespace}
}

func (c *FakeMonitoringV1alpha1) Slos(namespace string) v1alpha1.SloInterface {
	return &FakeSlos{c, namespace}
}

func (c *FakeMonitoringV1alpha1) UptimeCheckConfigs(namespace string) v1alpha1.UptimeCheckConfigInterface {
	return &FakeUptimeCheckConfigs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMonitoringV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
