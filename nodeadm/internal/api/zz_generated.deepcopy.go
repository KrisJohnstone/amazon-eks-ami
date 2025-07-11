//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package api

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDetails) DeepCopyInto(out *ClusterDetails) {
	*out = *in
	if in.CertificateAuthority != nil {
		in, out := &in.CertificateAuthority, &out.CertificateAuthority
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.EnableOutpost != nil {
		in, out := &in.EnableOutpost, &out.EnableOutpost
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDetails.
func (in *ClusterDetails) DeepCopy() *ClusterDetails {
	if in == nil {
		return nil
	}
	out := new(ClusterDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerdOptions) DeepCopyInto(out *ContainerdOptions) {
	*out = *in
	if in.BaseRuntimeSpec != nil {
		in, out := &in.BaseRuntimeSpec, &out.BaseRuntimeSpec
		*out = make(InlineDocument, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerdOptions.
func (in *ContainerdOptions) DeepCopy() *ContainerdOptions {
	if in == nil {
		return nil
	}
	out := new(ContainerdOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultOptions) DeepCopyInto(out *DefaultOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultOptions.
func (in *DefaultOptions) DeepCopy() *DefaultOptions {
	if in == nil {
		return nil
	}
	out := new(DefaultOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in InlineDocument) DeepCopyInto(out *InlineDocument) {
	{
		in := &in
		*out = make(InlineDocument, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InlineDocument.
func (in InlineDocument) DeepCopy() InlineDocument {
	if in == nil {
		return nil
	}
	out := new(InlineDocument)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceDetails) DeepCopyInto(out *InstanceDetails) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceDetails.
func (in *InstanceDetails) DeepCopy() *InstanceDetails {
	if in == nil {
		return nil
	}
	out := new(InstanceDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceOptions) DeepCopyInto(out *InstanceOptions) {
	*out = *in
	in.LocalStorage.DeepCopyInto(&out.LocalStorage)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceOptions.
func (in *InstanceOptions) DeepCopy() *InstanceOptions {
	if in == nil {
		return nil
	}
	out := new(InstanceOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in KubeletFlags) DeepCopyInto(out *KubeletFlags) {
	{
		in := &in
		*out = make(KubeletFlags, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletFlags.
func (in KubeletFlags) DeepCopy() KubeletFlags {
	if in == nil {
		return nil
	}
	out := new(KubeletFlags)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletOptions) DeepCopyInto(out *KubeletOptions) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(InlineDocument, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = make(KubeletFlags, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletOptions.
func (in *KubeletOptions) DeepCopy() *KubeletOptions {
	if in == nil {
		return nil
	}
	out := new(KubeletOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageOptions) DeepCopyInto(out *LocalStorageOptions) {
	*out = *in
	if in.DisabledMounts != nil {
		in, out := &in.DisabledMounts, &out.DisabledMounts
		*out = make([]DisabledMount, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageOptions.
func (in *LocalStorageOptions) DeepCopy() *LocalStorageOptions {
	if in == nil {
		return nil
	}
	out := new(LocalStorageOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfig) DeepCopyInto(out *NodeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfig.
func (in *NodeConfig) DeepCopy() *NodeConfig {
	if in == nil {
		return nil
	}
	out := new(NodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfigList) DeepCopyInto(out *NodeConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfigList.
func (in *NodeConfigList) DeepCopy() *NodeConfigList {
	if in == nil {
		return nil
	}
	out := new(NodeConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfigSpec) DeepCopyInto(out *NodeConfigSpec) {
	*out = *in
	in.Cluster.DeepCopyInto(&out.Cluster)
	in.Containerd.DeepCopyInto(&out.Containerd)
	in.Instance.DeepCopyInto(&out.Instance)
	in.Kubelet.DeepCopyInto(&out.Kubelet)
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[Feature]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfigSpec.
func (in *NodeConfigSpec) DeepCopy() *NodeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(NodeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfigStatus) DeepCopyInto(out *NodeConfigStatus) {
	*out = *in
	out.Instance = in.Instance
	out.Defaults = in.Defaults
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfigStatus.
func (in *NodeConfigStatus) DeepCopy() *NodeConfigStatus {
	if in == nil {
		return nil
	}
	out := new(NodeConfigStatus)
	in.DeepCopyInto(out)
	return out
}
