// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	v2alpha1 "k8s.io/api/batch/v2alpha1"
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// Deprecated: GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Workflow).DeepCopyInto(out.(*Workflow))
			return nil
		}, InType: reflect.TypeOf(&Workflow{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*WorkflowCondition).DeepCopyInto(out.(*WorkflowCondition))
			return nil
		}, InType: reflect.TypeOf(&WorkflowCondition{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*WorkflowList).DeepCopyInto(out.(*WorkflowList))
			return nil
		}, InType: reflect.TypeOf(&WorkflowList{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*WorkflowSpec).DeepCopyInto(out.(*WorkflowSpec))
			return nil
		}, InType: reflect.TypeOf(&WorkflowSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*WorkflowStatus).DeepCopyInto(out.(*WorkflowStatus))
			return nil
		}, InType: reflect.TypeOf(&WorkflowStatus{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*WorkflowStep).DeepCopyInto(out.(*WorkflowStep))
			return nil
		}, InType: reflect.TypeOf(&WorkflowStep{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*WorkflowStepStatus).DeepCopyInto(out.(*WorkflowStepStatus))
			return nil
		}, InType: reflect.TypeOf(&WorkflowStepStatus{})},
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workflow) DeepCopyInto(out *Workflow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new Workflow.
func (x *Workflow) DeepCopy() *Workflow {
	if x == nil {
		return nil
	}
	out := new(Workflow)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *Workflow) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowCondition) DeepCopyInto(out *WorkflowCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowCondition.
func (x *WorkflowCondition) DeepCopy() *WorkflowCondition {
	if x == nil {
		return nil
	}
	out := new(WorkflowCondition)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowList) DeepCopyInto(out *WorkflowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workflow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowList.
func (x *WorkflowList) DeepCopy() *WorkflowList {
	if x == nil {
		return nil
	}
	out := new(WorkflowList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *WorkflowList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowSpec) DeepCopyInto(out *WorkflowSpec) {
	*out = *in
	if in.ActiveDeadlineSeconds != nil {
		in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]WorkflowStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowSpec.
func (x *WorkflowSpec) DeepCopy() *WorkflowSpec {
	if x == nil {
		return nil
	}
	out := new(WorkflowSpec)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowStatus) DeepCopyInto(out *WorkflowStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]WorkflowCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Statuses != nil {
		in, out := &in.Statuses, &out.Statuses
		*out = make([]WorkflowStepStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowStatus.
func (x *WorkflowStatus) DeepCopy() *WorkflowStatus {
	if x == nil {
		return nil
	}
	out := new(WorkflowStatus)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowStep) DeepCopyInto(out *WorkflowStep) {
	*out = *in
	if in.JobTemplate != nil {
		in, out := &in.JobTemplate, &out.JobTemplate
		if *in == nil {
			*out = nil
		} else {
			*out = new(v2alpha1.JobTemplateSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.ObjectReference)
			**out = **in
		}
	}
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowStep.
func (x *WorkflowStep) DeepCopy() *WorkflowStep {
	if x == nil {
		return nil
	}
	out := new(WorkflowStep)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowStepStatus) DeepCopyInto(out *WorkflowStepStatus) {
	*out = *in
	out.Reference = in.Reference
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowStepStatus.
func (x *WorkflowStepStatus) DeepCopy() *WorkflowStepStatus {
	if x == nil {
		return nil
	}
	out := new(WorkflowStepStatus)
	x.DeepCopyInto(out)
	return out
}
