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

package v1alpha1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource/resourcestrategy"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Mouse
// +k8s:openapi-gen=true
type Mouse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MouseSpec   `json:"spec,omitempty"`
	Status MouseStatus `json:"status,omitempty"`
}

// MouseList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MouseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Mouse `json:"items"`
}

// MouseSpec defines the desired state of Mouse
type MouseSpec struct {
}

var _ resource.Object = &Mouse{}
var _ resourcestrategy.Validater = &Mouse{}

func (in *Mouse) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

func (in *Mouse) NamespaceScoped() bool {
	return false
}

func (in *Mouse) New() runtime.Object {
	return &Mouse{}
}

func (in *Mouse) NewList() runtime.Object {
	return &MouseList{}
}

func (in *Mouse) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "groupie.cwxstat.dev",
		Version:  "v1alpha1",
		Resource: "mice",
	}
}

func (in *Mouse) IsStorageVersion() bool {
	return true
}

func (in *Mouse) Validate(ctx context.Context) field.ErrorList {
	// TODO(user): Modify it, adding your API validation here.
	return nil
}

var _ resource.ObjectList = &MouseList{}

func (in *MouseList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}

// MouseStatus defines the observed state of Mouse
type MouseStatus struct {
}

// Mouse implements ObjectWithStatusSubResource interface.
var _ resource.ObjectWithStatusSubResource = &Mouse{}

func (in *Mouse) GetStatus() resource.StatusSubResource {
	return in.Status
}

// MouseStatus{} implements StatusSubResource interface.
var _ resource.StatusSubResource = &MouseStatus{}

func (in MouseStatus) CopyTo(parent resource.ObjectWithStatusSubResource) {
	parent.(*Mouse).Status = in
}
