// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvPrefix) DeepCopyInto(out *EnvPrefix) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvPrefix.
func (in *EnvPrefix) DeepCopy() *EnvPrefix {
	if in == nil {
		return nil
	}
	out := new(EnvPrefix)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalReferenceByKindAndName) DeepCopyInto(out *LocalReferenceByKindAndName) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalReferenceByKindAndName.
func (in *LocalReferenceByKindAndName) DeepCopy() *LocalReferenceByKindAndName {
	if in == nil {
		return nil
	}
	out := new(LocalReferenceByKindAndName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalReferenceByName) DeepCopyInto(out *LocalReferenceByName) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalReferenceByName.
func (in *LocalReferenceByName) DeepCopy() *LocalReferenceByName {
	if in == nil {
		return nil
	}
	out := new(LocalReferenceByName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parameters) DeepCopyInto(out *Parameters) {
	*out = *in
	if in.EnvPrefix != nil {
		in, out := &in.EnvPrefix, &out.EnvPrefix
		*out = new(EnvPrefix)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parameters.
func (in *Parameters) DeepCopy() *Parameters {
	if in == nil {
		return nil
	}
	out := new(Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceReference) DeepCopyInto(out *ResourceReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceReference.
func (in *ResourceReference) DeepCopy() *ResourceReference {
	if in == nil {
		return nil
	}
	out := new(ResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingUsage) DeepCopyInto(out *ServiceBindingUsage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingUsage.
func (in *ServiceBindingUsage) DeepCopy() *ServiceBindingUsage {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceBindingUsage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingUsageCondition) DeepCopyInto(out *ServiceBindingUsageCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingUsageCondition.
func (in *ServiceBindingUsageCondition) DeepCopy() *ServiceBindingUsageCondition {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingUsageCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingUsageList) DeepCopyInto(out *ServiceBindingUsageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceBindingUsage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingUsageList.
func (in *ServiceBindingUsageList) DeepCopy() *ServiceBindingUsageList {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingUsageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceBindingUsageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingUsageSpec) DeepCopyInto(out *ServiceBindingUsageSpec) {
	*out = *in
	out.ServiceBindingRef = in.ServiceBindingRef
	out.UsedBy = in.UsedBy
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(Parameters)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingUsageSpec.
func (in *ServiceBindingUsageSpec) DeepCopy() *ServiceBindingUsageSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingUsageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingUsageStatus) DeepCopyInto(out *ServiceBindingUsageStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ServiceBindingUsageCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingUsageStatus.
func (in *ServiceBindingUsageStatus) DeepCopy() *ServiceBindingUsageStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingUsageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UsageKind) DeepCopyInto(out *UsageKind) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UsageKind.
func (in *UsageKind) DeepCopy() *UsageKind {
	if in == nil {
		return nil
	}
	out := new(UsageKind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UsageKind) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UsageKindList) DeepCopyInto(out *UsageKindList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UsageKind, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UsageKindList.
func (in *UsageKindList) DeepCopy() *UsageKindList {
	if in == nil {
		return nil
	}
	out := new(UsageKindList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UsageKindList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UsageKindSpec) DeepCopyInto(out *UsageKindSpec) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(ResourceReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UsageKindSpec.
func (in *UsageKindSpec) DeepCopy() *UsageKindSpec {
	if in == nil {
		return nil
	}
	out := new(UsageKindSpec)
	in.DeepCopyInto(out)
	return out
}
