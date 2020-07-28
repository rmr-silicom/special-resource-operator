// +build !ignore_autogenerated

// Code generated by operator-sdk-v0.12.0-x86_64-linux-gnu. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResource) DeepCopyInto(out *SpecialResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecialResource.
func (in *SpecialResource) DeepCopy() *SpecialResource {
	if in == nil {
		return nil
	}
	out := new(SpecialResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpecialResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResourceArtifacts) DeepCopyInto(out *SpecialResourceArtifacts) {
	*out = *in
	if in.HostPaths != nil {
		in, out := &in.HostPaths, &out.HostPaths
		*out = make([]SpecialResourcePaths, len(*in))
		copy(*out, *in)
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]SpecialResourceImages, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Claims != nil {
		in, out := &in.Claims, &out.Claims
		*out = make([]SpecialResourceClaims, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecialResourceArtifacts.
func (in *SpecialResourceArtifacts) DeepCopy() *SpecialResourceArtifacts {
	if in == nil {
		return nil
	}
	out := new(SpecialResourceArtifacts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResourceBuilArgs) DeepCopyInto(out *SpecialResourceBuilArgs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecialResourceBuilArgs.
func (in *SpecialResourceBuilArgs) DeepCopy() *SpecialResourceBuilArgs {
	if in == nil {
		return nil
	}
	out := new(SpecialResourceBuilArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResourceClaims) DeepCopyInto(out *SpecialResourceClaims) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecialResourceClaims.
func (in *SpecialResourceClaims) DeepCopy() *SpecialResourceClaims {
	if in == nil {
		return nil
	}
	out := new(SpecialResourceClaims)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResourceDependsOn) DeepCopyInto(out *SpecialResourceDependsOn) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecialResourceDependsOn.
func (in *SpecialResourceDependsOn) DeepCopy() *SpecialResourceDependsOn {
	if in == nil {
		return nil
	}
	out := new(SpecialResourceDependsOn)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResourceDriverContainer) DeepCopyInto(out *SpecialResourceDriverContainer) {
	*out = *in
	out.Source = in.Source
	if in.BuildArgs != nil {
		in, out := &in.BuildArgs, &out.BuildArgs
		*out = make([]SpecialResourceBuilArgs, len(*in))
		copy(*out, *in)
	}
	in.Artifacts.DeepCopyInto(&out.Artifacts)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecialResourceDriverContainer.
func (in *SpecialResourceDriverContainer) DeepCopy() *SpecialResourceDriverContainer {
	if in == nil {
		return nil
	}
	out := new(SpecialResourceDriverContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResourceGit) DeepCopyInto(out *SpecialResourceGit) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecialResourceGit.
func (in *SpecialResourceGit) DeepCopy() *SpecialResourceGit {
	if in == nil {
		return nil
	}
	out := new(SpecialResourceGit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResourceImages) DeepCopyInto(out *SpecialResourceImages) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]SpecialResourcePaths, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecialResourceImages.
func (in *SpecialResourceImages) DeepCopy() *SpecialResourceImages {
	if in == nil {
		return nil
	}
	out := new(SpecialResourceImages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResourceList) DeepCopyInto(out *SpecialResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpecialResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecialResourceList.
func (in *SpecialResourceList) DeepCopy() *SpecialResourceList {
	if in == nil {
		return nil
	}
	out := new(SpecialResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpecialResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResourceNode) DeepCopyInto(out *SpecialResourceNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecialResourceNode.
func (in *SpecialResourceNode) DeepCopy() *SpecialResourceNode {
	if in == nil {
		return nil
	}
	out := new(SpecialResourceNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResourcePaths) DeepCopyInto(out *SpecialResourcePaths) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecialResourcePaths.
func (in *SpecialResourcePaths) DeepCopy() *SpecialResourcePaths {
	if in == nil {
		return nil
	}
	out := new(SpecialResourcePaths)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResourceSource) DeepCopyInto(out *SpecialResourceSource) {
	*out = *in
	out.Git = in.Git
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecialResourceSource.
func (in *SpecialResourceSource) DeepCopy() *SpecialResourceSource {
	if in == nil {
		return nil
	}
	out := new(SpecialResourceSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResourceSpec) DeepCopyInto(out *SpecialResourceSpec) {
	*out = *in
	in.DriverContainer.DeepCopyInto(&out.DriverContainer)
	out.Node = in.Node
	in.DependsOn.DeepCopyInto(&out.DependsOn)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecialResourceSpec.
func (in *SpecialResourceSpec) DeepCopy() *SpecialResourceSpec {
	if in == nil {
		return nil
	}
	out := new(SpecialResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResourceStatus) DeepCopyInto(out *SpecialResourceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecialResourceStatus.
func (in *SpecialResourceStatus) DeepCopy() *SpecialResourceStatus {
	if in == nil {
		return nil
	}
	out := new(SpecialResourceStatus)
	in.DeepCopyInto(out)
	return out
}
