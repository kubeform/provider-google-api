// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
	apiv1alpha1 "kubeform.dev/apimachinery/api/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Device) DeepCopyInto(out *Device) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Device.
func (in *Device) DeepCopy() *Device {
	if in == nil {
		return nil
	}
	out := new(Device)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Device) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceList) DeepCopyInto(out *DeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Device, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceList.
func (in *DeviceList) DeepCopy() *DeviceList {
	if in == nil {
		return nil
	}
	out := new(DeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSpec) DeepCopyInto(out *DeviceSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(DeviceSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSpec.
func (in *DeviceSpec) DeepCopy() *DeviceSpec {
	if in == nil {
		return nil
	}
	out := new(DeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSpecConfig) DeepCopyInto(out *DeviceSpecConfig) {
	*out = *in
	if in.BinaryData != nil {
		in, out := &in.BinaryData, &out.BinaryData
		*out = new(string)
		**out = **in
	}
	if in.CloudUpdateTime != nil {
		in, out := &in.CloudUpdateTime, &out.CloudUpdateTime
		*out = new(string)
		**out = **in
	}
	if in.DeviceACKTime != nil {
		in, out := &in.DeviceACKTime, &out.DeviceACKTime
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSpecConfig.
func (in *DeviceSpecConfig) DeepCopy() *DeviceSpecConfig {
	if in == nil {
		return nil
	}
	out := new(DeviceSpecConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSpecCredentials) DeepCopyInto(out *DeviceSpecCredentials) {
	*out = *in
	if in.ExpirationTime != nil {
		in, out := &in.ExpirationTime, &out.ExpirationTime
		*out = new(string)
		**out = **in
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = new(DeviceSpecCredentialsPublicKey)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSpecCredentials.
func (in *DeviceSpecCredentials) DeepCopy() *DeviceSpecCredentials {
	if in == nil {
		return nil
	}
	out := new(DeviceSpecCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSpecCredentialsPublicKey) DeepCopyInto(out *DeviceSpecCredentialsPublicKey) {
	*out = *in
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSpecCredentialsPublicKey.
func (in *DeviceSpecCredentialsPublicKey) DeepCopy() *DeviceSpecCredentialsPublicKey {
	if in == nil {
		return nil
	}
	out := new(DeviceSpecCredentialsPublicKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSpecGatewayConfig) DeepCopyInto(out *DeviceSpecGatewayConfig) {
	*out = *in
	if in.GatewayAuthMethod != nil {
		in, out := &in.GatewayAuthMethod, &out.GatewayAuthMethod
		*out = new(string)
		**out = **in
	}
	if in.GatewayType != nil {
		in, out := &in.GatewayType, &out.GatewayType
		*out = new(string)
		**out = **in
	}
	if in.LastAccessedGatewayID != nil {
		in, out := &in.LastAccessedGatewayID, &out.LastAccessedGatewayID
		*out = new(string)
		**out = **in
	}
	if in.LastAccessedGatewayTime != nil {
		in, out := &in.LastAccessedGatewayTime, &out.LastAccessedGatewayTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSpecGatewayConfig.
func (in *DeviceSpecGatewayConfig) DeepCopy() *DeviceSpecGatewayConfig {
	if in == nil {
		return nil
	}
	out := new(DeviceSpecGatewayConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSpecLastErrorStatus) DeepCopyInto(out *DeviceSpecLastErrorStatus) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Number != nil {
		in, out := &in.Number, &out.Number
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSpecLastErrorStatus.
func (in *DeviceSpecLastErrorStatus) DeepCopy() *DeviceSpecLastErrorStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceSpecLastErrorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSpecResource) DeepCopyInto(out *DeviceSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Blocked != nil {
		in, out := &in.Blocked, &out.Blocked
		*out = new(bool)
		**out = **in
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]DeviceSpecConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]DeviceSpecCredentials, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GatewayConfig != nil {
		in, out := &in.GatewayConfig, &out.GatewayConfig
		*out = new(DeviceSpecGatewayConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.LastConfigACKTime != nil {
		in, out := &in.LastConfigACKTime, &out.LastConfigACKTime
		*out = new(string)
		**out = **in
	}
	if in.LastConfigSendTime != nil {
		in, out := &in.LastConfigSendTime, &out.LastConfigSendTime
		*out = new(string)
		**out = **in
	}
	if in.LastErrorStatus != nil {
		in, out := &in.LastErrorStatus, &out.LastErrorStatus
		*out = make([]DeviceSpecLastErrorStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastErrorTime != nil {
		in, out := &in.LastErrorTime, &out.LastErrorTime
		*out = new(string)
		**out = **in
	}
	if in.LastEventTime != nil {
		in, out := &in.LastEventTime, &out.LastEventTime
		*out = new(string)
		**out = **in
	}
	if in.LastHeartbeatTime != nil {
		in, out := &in.LastHeartbeatTime, &out.LastHeartbeatTime
		*out = new(string)
		**out = **in
	}
	if in.LastStateTime != nil {
		in, out := &in.LastStateTime, &out.LastStateTime
		*out = new(string)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NumID != nil {
		in, out := &in.NumID, &out.NumID
		*out = new(string)
		**out = **in
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = make([]DeviceSpecState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSpecResource.
func (in *DeviceSpecResource) DeepCopy() *DeviceSpecResource {
	if in == nil {
		return nil
	}
	out := new(DeviceSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSpecState) DeepCopyInto(out *DeviceSpecState) {
	*out = *in
	if in.BinaryData != nil {
		in, out := &in.BinaryData, &out.BinaryData
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSpecState.
func (in *DeviceSpecState) DeepCopy() *DeviceSpecState {
	if in == nil {
		return nil
	}
	out := new(DeviceSpecState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceStatus) DeepCopyInto(out *DeviceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceStatus.
func (in *DeviceStatus) DeepCopy() *DeviceStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Registry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryList) DeepCopyInto(out *RegistryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Registry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryList.
func (in *RegistryList) DeepCopy() *RegistryList {
	if in == nil {
		return nil
	}
	out := new(RegistryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistrySpec) DeepCopyInto(out *RegistrySpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(RegistrySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistrySpec.
func (in *RegistrySpec) DeepCopy() *RegistrySpec {
	if in == nil {
		return nil
	}
	out := new(RegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistrySpecCredentials) DeepCopyInto(out *RegistrySpecCredentials) {
	*out = *in
	if in.PublicKeyCertificate != nil {
		in, out := &in.PublicKeyCertificate, &out.PublicKeyCertificate
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistrySpecCredentials.
func (in *RegistrySpecCredentials) DeepCopy() *RegistrySpecCredentials {
	if in == nil {
		return nil
	}
	out := new(RegistrySpecCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistrySpecEventNotificationConfigs) DeepCopyInto(out *RegistrySpecEventNotificationConfigs) {
	*out = *in
	if in.PubsubTopicName != nil {
		in, out := &in.PubsubTopicName, &out.PubsubTopicName
		*out = new(string)
		**out = **in
	}
	if in.SubfolderMatches != nil {
		in, out := &in.SubfolderMatches, &out.SubfolderMatches
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistrySpecEventNotificationConfigs.
func (in *RegistrySpecEventNotificationConfigs) DeepCopy() *RegistrySpecEventNotificationConfigs {
	if in == nil {
		return nil
	}
	out := new(RegistrySpecEventNotificationConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistrySpecResource) DeepCopyInto(out *RegistrySpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]RegistrySpecCredentials, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EventNotificationConfigs != nil {
		in, out := &in.EventNotificationConfigs, &out.EventNotificationConfigs
		*out = make([]RegistrySpecEventNotificationConfigs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HttpConfig != nil {
		in, out := &in.HttpConfig, &out.HttpConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.MqttConfig != nil {
		in, out := &in.MqttConfig, &out.MqttConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.StateNotificationConfig != nil {
		in, out := &in.StateNotificationConfig, &out.StateNotificationConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistrySpecResource.
func (in *RegistrySpecResource) DeepCopy() *RegistrySpecResource {
	if in == nil {
		return nil
	}
	out := new(RegistrySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryStatus) DeepCopyInto(out *RegistryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryStatus.
func (in *RegistryStatus) DeepCopy() *RegistryStatus {
	if in == nil {
		return nil
	}
	out := new(RegistryStatus)
	in.DeepCopyInto(out)
	return out
}
