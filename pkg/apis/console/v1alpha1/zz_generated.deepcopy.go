// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Console) DeepCopyInto(out *Console) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Console.
func (in *Console) DeepCopy() *Console {
	if in == nil {
		return nil
	}
	out := new(Console)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Console) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleConfigRoute) DeepCopyInto(out *ConsoleConfigRoute) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleConfigRoute.
func (in *ConsoleConfigRoute) DeepCopy() *ConsoleConfigRoute {
	if in == nil {
		return nil
	}
	out := new(ConsoleConfigRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleList) DeepCopyInto(out *ConsoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Console, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleList.
func (in *ConsoleList) DeepCopy() *ConsoleList {
	if in == nil {
		return nil
	}
	out := new(ConsoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleSpec) DeepCopyInto(out *ConsoleSpec) {
	*out = *in
	out.OperatorSpec = in.OperatorSpec
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]ConsoleConfigRoute, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleSpec.
func (in *ConsoleSpec) DeepCopy() *ConsoleSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleStatus) DeepCopyInto(out *ConsoleStatus) {
	*out = *in
	in.OperatorStatus.DeepCopyInto(&out.OperatorStatus)
	if in.HostNames != nil {
		in, out := &in.HostNames, &out.HostNames
		*out = make([]ConsoleConfigRoute, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleStatus.
func (in *ConsoleStatus) DeepCopy() *ConsoleStatus {
	if in == nil {
		return nil
	}
	out := new(ConsoleStatus)
	in.DeepCopyInto(out)
	return out
}
