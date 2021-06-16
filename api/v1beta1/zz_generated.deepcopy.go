// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/openshift-psap/special-resource-operator/pkg/helmer"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResource) DeepCopyInto(out *SpecialResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
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
func (in *SpecialResourceBuildArgs) DeepCopyInto(out *SpecialResourceBuildArgs) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecialResourceBuildArgs.
func (in *SpecialResourceBuildArgs) DeepCopy() *SpecialResourceBuildArgs {
	if in == nil {
		return nil
	}
	out := new(SpecialResourceBuildArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResourceClaims) DeepCopyInto(out *SpecialResourceClaims) {
	*out = *in
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
func (in *SpecialResourceConfiguration) DeepCopyInto(out *SpecialResourceConfiguration) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecialResourceConfiguration.
func (in *SpecialResourceConfiguration) DeepCopy() *SpecialResourceConfiguration {
	if in == nil {
		return nil
	}
	out := new(SpecialResourceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecialResourceDriverContainer) DeepCopyInto(out *SpecialResourceDriverContainer) {
	*out = *in
	out.Source = in.Source
	in.Artifacts.DeepCopyInto(&out.Artifacts)
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
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpecialResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
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
func (in *SpecialResourcePaths) DeepCopyInto(out *SpecialResourcePaths) {
	*out = *in
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
	in.Chart.DeepCopyInto(&out.Chart)
	in.Set.DeepCopyInto(&out.Set)
	in.DriverContainer.DeepCopyInto(&out.DriverContainer)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]helmer.HelmChart, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
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
