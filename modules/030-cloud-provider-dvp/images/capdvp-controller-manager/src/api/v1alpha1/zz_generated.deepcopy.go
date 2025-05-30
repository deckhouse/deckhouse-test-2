//go:build !ignore_autogenerated

/*
Copyright 2025 Flant JSC

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPU) DeepCopyInto(out *CPU) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPU.
func (in *CPU) DeepCopy() *CPU {
	if in == nil {
		return nil
	}
	out := new(CPU)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeckhouseCluster) DeepCopyInto(out *DeckhouseCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeckhouseCluster.
func (in *DeckhouseCluster) DeepCopy() *DeckhouseCluster {
	if in == nil {
		return nil
	}
	out := new(DeckhouseCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeckhouseCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeckhouseClusterList) DeepCopyInto(out *DeckhouseClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeckhouseCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeckhouseClusterList.
func (in *DeckhouseClusterList) DeepCopy() *DeckhouseClusterList {
	if in == nil {
		return nil
	}
	out := new(DeckhouseClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeckhouseClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeckhouseClusterSpec) DeepCopyInto(out *DeckhouseClusterSpec) {
	*out = *in
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeckhouseClusterSpec.
func (in *DeckhouseClusterSpec) DeepCopy() *DeckhouseClusterSpec {
	if in == nil {
		return nil
	}
	out := new(DeckhouseClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeckhouseClusterStatus) DeepCopyInto(out *DeckhouseClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeckhouseClusterStatus.
func (in *DeckhouseClusterStatus) DeepCopy() *DeckhouseClusterStatus {
	if in == nil {
		return nil
	}
	out := new(DeckhouseClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeckhouseMachine) DeepCopyInto(out *DeckhouseMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeckhouseMachine.
func (in *DeckhouseMachine) DeepCopy() *DeckhouseMachine {
	if in == nil {
		return nil
	}
	out := new(DeckhouseMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeckhouseMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeckhouseMachineList) DeepCopyInto(out *DeckhouseMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeckhouseMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeckhouseMachineList.
func (in *DeckhouseMachineList) DeepCopy() *DeckhouseMachineList {
	if in == nil {
		return nil
	}
	out := new(DeckhouseMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeckhouseMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeckhouseMachineSpec) DeepCopyInto(out *DeckhouseMachineSpec) {
	*out = *in
	out.CPU = in.CPU
	out.Memory = in.Memory.DeepCopy()
	out.RootDiskSize = in.RootDiskSize.DeepCopy()
	out.BootDiskImageRef = in.BootDiskImageRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeckhouseMachineSpec.
func (in *DeckhouseMachineSpec) DeepCopy() *DeckhouseMachineSpec {
	if in == nil {
		return nil
	}
	out := new(DeckhouseMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeckhouseMachineSpecTemplate) DeepCopyInto(out *DeckhouseMachineSpecTemplate) {
	*out = *in
	out.CPU = in.CPU
	out.Memory = in.Memory.DeepCopy()
	out.RootDiskSize = in.RootDiskSize.DeepCopy()
	out.BootDiskImageRef = in.BootDiskImageRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeckhouseMachineSpecTemplate.
func (in *DeckhouseMachineSpecTemplate) DeepCopy() *DeckhouseMachineSpecTemplate {
	if in == nil {
		return nil
	}
	out := new(DeckhouseMachineSpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeckhouseMachineStatus) DeepCopyInto(out *DeckhouseMachineStatus) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]VMAddress, len(*in))
		copy(*out, *in)
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(string)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeckhouseMachineStatus.
func (in *DeckhouseMachineStatus) DeepCopy() *DeckhouseMachineStatus {
	if in == nil {
		return nil
	}
	out := new(DeckhouseMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeckhouseMachineTemplate) DeepCopyInto(out *DeckhouseMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeckhouseMachineTemplate.
func (in *DeckhouseMachineTemplate) DeepCopy() *DeckhouseMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(DeckhouseMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeckhouseMachineTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeckhouseMachineTemplateList) DeepCopyInto(out *DeckhouseMachineTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeckhouseMachineTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeckhouseMachineTemplateList.
func (in *DeckhouseMachineTemplateList) DeepCopy() *DeckhouseMachineTemplateList {
	if in == nil {
		return nil
	}
	out := new(DeckhouseMachineTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeckhouseMachineTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeckhouseMachineTemplateSpec) DeepCopyInto(out *DeckhouseMachineTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeckhouseMachineTemplateSpec.
func (in *DeckhouseMachineTemplateSpec) DeepCopy() *DeckhouseMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(DeckhouseMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeckhouseMachineTemplateSpecTemplate) DeepCopyInto(out *DeckhouseMachineTemplateSpecTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeckhouseMachineTemplateSpecTemplate.
func (in *DeckhouseMachineTemplateSpecTemplate) DeepCopy() *DeckhouseMachineTemplateSpecTemplate {
	if in == nil {
		return nil
	}
	out := new(DeckhouseMachineTemplateSpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeckhouseMachineTemplateStatus) DeepCopyInto(out *DeckhouseMachineTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeckhouseMachineTemplateStatus.
func (in *DeckhouseMachineTemplateStatus) DeepCopy() *DeckhouseMachineTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(DeckhouseMachineTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskImageRef) DeepCopyInto(out *DiskImageRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskImageRef.
func (in *DiskImageRef) DeepCopy() *DiskImageRef {
	if in == nil {
		return nil
	}
	out := new(DiskImageRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMAddress) DeepCopyInto(out *VMAddress) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMAddress.
func (in *VMAddress) DeepCopy() *VMAddress {
	if in == nil {
		return nil
	}
	out := new(VMAddress)
	in.DeepCopyInto(out)
	return out
}
