// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Container) DeepCopyInto(out *Container) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Container.
func (in *Container) DeepCopy() *Container {
	if in == nil {
		return nil
	}
	out := new(Container)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Reactor) DeepCopyInto(out *Reactor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Reactor.
func (in *Reactor) DeepCopy() *Reactor {
	if in == nil {
		return nil
	}
	out := new(Reactor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Reactor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReactorDeploymentSpec) DeepCopyInto(out *ReactorDeploymentSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReactorDeploymentSpec.
func (in *ReactorDeploymentSpec) DeepCopy() *ReactorDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(ReactorDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReactorDeploymentTemplate) DeepCopyInto(out *ReactorDeploymentTemplate) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReactorDeploymentTemplate.
func (in *ReactorDeploymentTemplate) DeepCopy() *ReactorDeploymentTemplate {
	if in == nil {
		return nil
	}
	out := new(ReactorDeploymentTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReactorDeploymentTemplateSpec) DeepCopyInto(out *ReactorDeploymentTemplateSpec) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]Container, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReactorDeploymentTemplateSpec.
func (in *ReactorDeploymentTemplateSpec) DeepCopy() *ReactorDeploymentTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ReactorDeploymentTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReactorList) DeepCopyInto(out *ReactorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Reactor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReactorList.
func (in *ReactorList) DeepCopy() *ReactorList {
	if in == nil {
		return nil
	}
	out := new(ReactorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReactorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReactorSpec) DeepCopyInto(out *ReactorSpec) {
	*out = *in
	out.Reagent = in.Reagent
	in.Deployment.DeepCopyInto(&out.Deployment)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReactorSpec.
func (in *ReactorSpec) DeepCopy() *ReactorSpec {
	if in == nil {
		return nil
	}
	out := new(ReactorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReagentSpec) DeepCopyInto(out *ReagentSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReagentSpec.
func (in *ReagentSpec) DeepCopy() *ReagentSpec {
	if in == nil {
		return nil
	}
	out := new(ReagentSpec)
	in.DeepCopyInto(out)
	return out
}
