/*
Copyright 2019 TAKAISHI Ryo.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConcourseCIPipelineSpec defines the desired state of ConcourseCIPipeline
type ConcourseCIPipelineSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Target   string `json:"target"`
	Pipeline string `json:"pipeline"`
	Public   bool   `json:"public,omitempty"`
	Manifest string `json:"manifest"`
}

// ConcourseCIPipelineStatus defines the observed state of ConcourseCIPipeline
type ConcourseCIPipelineStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConcourseCIPipeline is the Schema for the concoursecipipelines API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type ConcourseCIPipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConcourseCIPipelineSpec   `json:"spec,omitempty"`
	Status ConcourseCIPipelineStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConcourseCIPipelineList contains a list of ConcourseCIPipeline
type ConcourseCIPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConcourseCIPipeline `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConcourseCIPipeline{}, &ConcourseCIPipelineList{})
}
