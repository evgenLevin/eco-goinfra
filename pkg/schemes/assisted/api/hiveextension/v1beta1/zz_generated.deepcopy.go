//go:build !ignore_autogenerated

/*
Copyright 2020.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/openshift-kni/eco-goinfra/pkg/schemes/assisted/api/common"
	"github.com/openshift-kni/eco-goinfra/pkg/schemes/assisted/hive/api/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentClusterInstall) DeepCopyInto(out *AgentClusterInstall) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentClusterInstall.
func (in *AgentClusterInstall) DeepCopy() *AgentClusterInstall {
	if in == nil {
		return nil
	}
	out := new(AgentClusterInstall)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentClusterInstall) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentClusterInstallList) DeepCopyInto(out *AgentClusterInstallList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AgentClusterInstall, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentClusterInstallList.
func (in *AgentClusterInstallList) DeepCopy() *AgentClusterInstallList {
	if in == nil {
		return nil
	}
	out := new(AgentClusterInstallList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentClusterInstallList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentClusterInstallSpec) DeepCopyInto(out *AgentClusterInstallSpec) {
	*out = *in
	if in.ImageSetRef != nil {
		in, out := &in.ImageSetRef, &out.ImageSetRef
		*out = new(v1.ClusterImageSetReference)
		**out = **in
	}
	out.ClusterDeploymentRef = in.ClusterDeploymentRef
	if in.ClusterMetadata != nil {
		in, out := &in.ClusterMetadata, &out.ClusterMetadata
		*out = new(v1.ClusterMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.ManifestsConfigMapRef != nil {
		in, out := &in.ManifestsConfigMapRef, &out.ManifestsConfigMapRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.ManifestsConfigMapRefs != nil {
		in, out := &in.ManifestsConfigMapRefs, &out.ManifestsConfigMapRefs
		*out = make([]ManifestsConfigMapReference, len(*in))
		copy(*out, *in)
	}
	in.Networking.DeepCopyInto(&out.Networking)
	out.ProvisionRequirements = in.ProvisionRequirements
	if in.ControlPlane != nil {
		in, out := &in.ControlPlane, &out.ControlPlane
		*out = new(AgentMachinePool)
		**out = **in
	}
	if in.Compute != nil {
		in, out := &in.Compute, &out.Compute
		*out = make([]AgentMachinePool, len(*in))
		copy(*out, *in)
	}
	if in.APIVIPs != nil {
		in, out := &in.APIVIPs, &out.APIVIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IngressVIPs != nil {
		in, out := &in.IngressVIPs, &out.IngressVIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IgnitionEndpoint != nil {
		in, out := &in.IgnitionEndpoint, &out.IgnitionEndpoint
		*out = new(IgnitionEndpoint)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskEncryption != nil {
		in, out := &in.DiskEncryption, &out.DiskEncryption
		*out = new(DiskEncryption)
		(*in).DeepCopyInto(*out)
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(Proxy)
		**out = **in
	}
	if in.ExternalPlatformSpec != nil {
		in, out := &in.ExternalPlatformSpec, &out.ExternalPlatformSpec
		*out = new(ExternalPlatformSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentClusterInstallSpec.
func (in *AgentClusterInstallSpec) DeepCopy() *AgentClusterInstallSpec {
	if in == nil {
		return nil
	}
	out := new(AgentClusterInstallSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentClusterInstallStatus) DeepCopyInto(out *AgentClusterInstallStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.ClusterInstallCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Progress = in.Progress
	if in.MachineNetwork != nil {
		in, out := &in.MachineNetwork, &out.MachineNetwork
		*out = make([]MachineNetworkEntry, len(*in))
		copy(*out, *in)
	}
	out.DebugInfo = in.DebugInfo
	if in.APIVIPs != nil {
		in, out := &in.APIVIPs, &out.APIVIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IngressVIPs != nil {
		in, out := &in.IngressVIPs, &out.IngressVIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UserManagedNetworking != nil {
		in, out := &in.UserManagedNetworking, &out.UserManagedNetworking
		*out = new(bool)
		**out = **in
	}
	if in.ValidationsInfo != nil {
		in, out := &in.ValidationsInfo, &out.ValidationsInfo
		*out = make(common.ValidationsStatus, len(*in))
		for key, val := range *in {
			var outVal []common.ValidationResult
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(common.ValidationResults, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentClusterInstallStatus.
func (in *AgentClusterInstallStatus) DeepCopy() *AgentClusterInstallStatus {
	if in == nil {
		return nil
	}
	out := new(AgentClusterInstallStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentMachinePool) DeepCopyInto(out *AgentMachinePool) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentMachinePool.
func (in *AgentMachinePool) DeepCopy() *AgentMachinePool {
	if in == nil {
		return nil
	}
	out := new(AgentMachinePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaCertificateReference) DeepCopyInto(out *CaCertificateReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaCertificateReference.
func (in *CaCertificateReference) DeepCopy() *CaCertificateReference {
	if in == nil {
		return nil
	}
	out := new(CaCertificateReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetworkEntry) DeepCopyInto(out *ClusterNetworkEntry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetworkEntry.
func (in *ClusterNetworkEntry) DeepCopy() *ClusterNetworkEntry {
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProgressInfo) DeepCopyInto(out *ClusterProgressInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProgressInfo.
func (in *ClusterProgressInfo) DeepCopy() *ClusterProgressInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterProgressInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DebugInfo) DeepCopyInto(out *DebugInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DebugInfo.
func (in *DebugInfo) DeepCopy() *DebugInfo {
	if in == nil {
		return nil
	}
	out := new(DebugInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskEncryption) DeepCopyInto(out *DiskEncryption) {
	*out = *in
	if in.EnableOn != nil {
		in, out := &in.EnableOn, &out.EnableOn
		*out = new(string)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskEncryption.
func (in *DiskEncryption) DeepCopy() *DiskEncryption {
	if in == nil {
		return nil
	}
	out := new(DiskEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalPlatformSpec) DeepCopyInto(out *ExternalPlatformSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalPlatformSpec.
func (in *ExternalPlatformSpec) DeepCopy() *ExternalPlatformSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalPlatformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IgnitionEndpoint) DeepCopyInto(out *IgnitionEndpoint) {
	*out = *in
	if in.CaCertificateReference != nil {
		in, out := &in.CaCertificateReference, &out.CaCertificateReference
		*out = new(CaCertificateReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IgnitionEndpoint.
func (in *IgnitionEndpoint) DeepCopy() *IgnitionEndpoint {
	if in == nil {
		return nil
	}
	out := new(IgnitionEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineNetworkEntry) DeepCopyInto(out *MachineNetworkEntry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineNetworkEntry.
func (in *MachineNetworkEntry) DeepCopy() *MachineNetworkEntry {
	if in == nil {
		return nil
	}
	out := new(MachineNetworkEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManifestsConfigMapReference) DeepCopyInto(out *ManifestsConfigMapReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManifestsConfigMapReference.
func (in *ManifestsConfigMapReference) DeepCopy() *ManifestsConfigMapReference {
	if in == nil {
		return nil
	}
	out := new(ManifestsConfigMapReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Networking) DeepCopyInto(out *Networking) {
	*out = *in
	if in.MachineNetwork != nil {
		in, out := &in.MachineNetwork, &out.MachineNetwork
		*out = make([]MachineNetworkEntry, len(*in))
		copy(*out, *in)
	}
	if in.ClusterNetwork != nil {
		in, out := &in.ClusterNetwork, &out.ClusterNetwork
		*out = make([]ClusterNetworkEntry, len(*in))
		copy(*out, *in)
	}
	if in.ServiceNetwork != nil {
		in, out := &in.ServiceNetwork, &out.ServiceNetwork
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UserManagedNetworking != nil {
		in, out := &in.UserManagedNetworking, &out.UserManagedNetworking
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Networking.
func (in *Networking) DeepCopy() *Networking {
	if in == nil {
		return nil
	}
	out := new(Networking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionRequirements) DeepCopyInto(out *ProvisionRequirements) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionRequirements.
func (in *ProvisionRequirements) DeepCopy() *ProvisionRequirements {
	if in == nil {
		return nil
	}
	out := new(ProvisionRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Proxy) DeepCopyInto(out *Proxy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Proxy.
func (in *Proxy) DeepCopy() *Proxy {
	if in == nil {
		return nil
	}
	out := new(Proxy)
	in.DeepCopyInto(out)
	return out
}
