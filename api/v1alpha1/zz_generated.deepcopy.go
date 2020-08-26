// +build !ignore_autogenerated

/*


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
	runtime "k8s.io/apimachinery/pkg/runtime"
	"math/big"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Jwt) DeepCopyInto(out *Jwt) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Jwt.
func (in *Jwt) DeepCopy() *Jwt {
	if in == nil {
		return nil
	}
	out := new(Jwt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Jwt) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JwtList) DeepCopyInto(out *JwtList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Jwt, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JwtList.
func (in *JwtList) DeepCopy() *JwtList {
	if in == nil {
		return nil
	}
	out := new(JwtList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JwtList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JwtSpec) DeepCopyInto(out *JwtSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JwtSpec.
func (in *JwtSpec) DeepCopy() *JwtSpec {
	if in == nil {
		return nil
	}
	out := new(JwtSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JwtStatus) DeepCopyInto(out *JwtStatus) {
	*out = *in
	in.ExpiresAt.DeepCopyInto(&out.ExpiresAt)
	in.RefreshAfter.DeepCopyInto(&out.RefreshAfter)
	if in.LastRefresh != nil {
		in, out := &in.LastRefresh, &out.LastRefresh
		*out = (*in).DeepCopy()
	}
	in.NextReconcile.DeepCopyInto(&out.NextReconcile)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JwtStatus.
func (in *JwtStatus) DeepCopy() *JwtStatus {
	if in == nil {
		return nil
	}
	out := new(JwtStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RotatingKey) DeepCopyInto(out *RotatingKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RotatingKey.
func (in *RotatingKey) DeepCopy() *RotatingKey {
	if in == nil {
		return nil
	}
	out := new(RotatingKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RotatingKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RotatingKeyList) DeepCopyInto(out *RotatingKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RotatingKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RotatingKeyList.
func (in *RotatingKeyList) DeepCopy() *RotatingKeyList {
	if in == nil {
		return nil
	}
	out := new(RotatingKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RotatingKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RotatingKeySpec) DeepCopyInto(out *RotatingKeySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RotatingKeySpec.
func (in *RotatingKeySpec) DeepCopy() *RotatingKeySpec {
	if in == nil {
		return nil
	}
	out := new(RotatingKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RotatingKeyStatus) DeepCopyInto(out *RotatingKeyStatus) {
	*out = *in
	in.NexRotation.DeepCopyInto(&out.NexRotation)
	if in.ValidationKeys != nil {
		in, out := &in.ValidationKeys, &out.ValidationKeys
		*out = make([]ValidationKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.SigningKey.DeepCopyInto(&out.SigningKey)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RotatingKeyStatus.
func (in *RotatingKeyStatus) DeepCopy() *RotatingKeyStatus {
	if in == nil {
		return nil
	}
	out := new(RotatingKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningKey) DeepCopyInto(out *SigningKey) {
	*out = *in
	if in.PrivateKey2 != nil {
		in, out := &in.PrivateKey2, &out.PrivateKey2
		*out = new(big.Int)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningKey.
func (in *SigningKey) DeepCopy() *SigningKey {
	if in == nil {
		return nil
	}
	out := new(SigningKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationKey) DeepCopyInto(out *ValidationKey) {
	*out = *in
	in.ExpireAt.DeepCopyInto(&out.ExpireAt)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationKey.
func (in *ValidationKey) DeepCopy() *ValidationKey {
	if in == nil {
		return nil
	}
	out := new(ValidationKey)
	in.DeepCopyInto(out)
	return out
}
