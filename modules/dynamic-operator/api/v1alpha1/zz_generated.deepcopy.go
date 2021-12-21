//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KuberLogicService) DeepCopyInto(out *KuberLogicService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KuberLogicService.
func (in *KuberLogicService) DeepCopy() *KuberLogicService {
	if in == nil {
		return nil
	}
	out := new(KuberLogicService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KuberLogicService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KuberLogicServiceList) DeepCopyInto(out *KuberLogicServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KuberLogicService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KuberLogicServiceList.
func (in *KuberLogicServiceList) DeepCopy() *KuberLogicServiceList {
	if in == nil {
		return nil
	}
	out := new(KuberLogicServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KuberLogicServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KuberLogicServiceStatus) DeepCopyInto(out *KuberLogicServiceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KuberLogicServiceStatus.
func (in *KuberLogicServiceStatus) DeepCopy() *KuberLogicServiceStatus {
	if in == nil {
		return nil
	}
	out := new(KuberLogicServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KuberLogicServiceType) DeepCopyInto(out *KuberLogicServiceType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KuberLogicServiceType.
func (in *KuberLogicServiceType) DeepCopy() *KuberLogicServiceType {
	if in == nil {
		return nil
	}
	out := new(KuberLogicServiceType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KuberLogicServiceType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KuberLogicServiceTypeConditions) DeepCopyInto(out *KuberLogicServiceTypeConditions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KuberLogicServiceTypeConditions.
func (in *KuberLogicServiceTypeConditions) DeepCopy() *KuberLogicServiceTypeConditions {
	if in == nil {
		return nil
	}
	out := new(KuberLogicServiceTypeConditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KuberLogicServiceTypeList) DeepCopyInto(out *KuberLogicServiceTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KuberLogicServiceType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KuberLogicServiceTypeList.
func (in *KuberLogicServiceTypeList) DeepCopy() *KuberLogicServiceTypeList {
	if in == nil {
		return nil
	}
	out := new(KuberLogicServiceTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KuberLogicServiceTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KuberLogicServiceTypeSpec) DeepCopyInto(out *KuberLogicServiceTypeSpec) {
	*out = *in
	out.Api = in.Api
	if in.SpecRef != nil {
		in, out := &in.SpecRef, &out.SpecRef
		*out = make(map[string]KuberlogicServiceTypeParam, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.StatusRef.DeepCopyInto(&out.StatusRef)
	in.DefaultSpec.DeepCopyInto(&out.DefaultSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KuberLogicServiceTypeSpec.
func (in *KuberLogicServiceTypeSpec) DeepCopy() *KuberLogicServiceTypeSpec {
	if in == nil {
		return nil
	}
	out := new(KuberLogicServiceTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KuberLogicServiceTypeStatus) DeepCopyInto(out *KuberLogicServiceTypeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KuberLogicServiceTypeStatus.
func (in *KuberLogicServiceTypeStatus) DeepCopy() *KuberLogicServiceTypeStatus {
	if in == nil {
		return nil
	}
	out := new(KuberLogicServiceTypeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KuberlogicServiceTypeParam) DeepCopyInto(out *KuberlogicServiceTypeParam) {
	*out = *in
	in.DefaultValue.DeepCopyInto(&out.DefaultValue)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KuberlogicServiceTypeParam.
func (in *KuberlogicServiceTypeParam) DeepCopy() *KuberlogicServiceTypeParam {
	if in == nil {
		return nil
	}
	out := new(KuberlogicServiceTypeParam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KuberlogicServiceTypeStatusRef) DeepCopyInto(out *KuberlogicServiceTypeStatusRef) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = new(KuberLogicServiceTypeConditions)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KuberlogicServiceTypeStatusRef.
func (in *KuberlogicServiceTypeStatusRef) DeepCopy() *KuberlogicServiceTypeStatusRef {
	if in == nil {
		return nil
	}
	out := new(KuberlogicServiceTypeStatusRef)
	in.DeepCopyInto(out)
	return out
}