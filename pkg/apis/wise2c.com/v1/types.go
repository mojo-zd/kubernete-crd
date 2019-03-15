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

package v1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/apimachinery/pkg/api/resource"
)

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Tenant Operator
type Tenant struct {
	v1.TypeMeta     `json:",inline"`
	v1.ObjectMeta   `json:"metadata,omitempty"`
	Spec TenantSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type TenantList struct {
	v1.TypeMeta    `json:",inline"`
	v1.ListMeta    `json:"metadata"`
	Items []Tenant `json:"items"`
}

// TenantSpec is the spec for a Tenant resource
type TenantSpec struct {
	Uuid     string `json:"uuid"`
	Creator  int64  `json:"creator"`
	CreateAt int64  `json:"createAt"`
	UpdateAt int64  `json:"updateAt"`
	Quota    Quota  `json:"quotas"`
}

type Quota struct {
	Hard core.ResourceList `json:"hard"`
	Used core.ResourceList `json:"used"`
}

//ResourceName is the name identifying various resources in a ResourceList.
type ResourceName string

// ResourceList is a set of (resource name, quantity) pairs.
type ResourceList map[ResourceName]resource.Quantity
