/*
Copyright 2020 xutao1989103.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ApplicationConfigurationSpec defines the desired state of ApplicationConfiguration
type ApplicationConfigurationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ApplicationConfiguration. Edit ApplicationConfiguration_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ApplicationConfigurationStatus defines the observed state of ApplicationConfiguration
type ApplicationConfigurationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// ApplicationConfiguration is the Schema for the applicationconfigurations API
type ApplicationConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationConfigurationSpec   `json:"spec,omitempty"`
	Status ApplicationConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationConfigurationList contains a list of ApplicationConfiguration
type ApplicationConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationConfiguration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApplicationConfiguration{}, &ApplicationConfigurationList{})
}
