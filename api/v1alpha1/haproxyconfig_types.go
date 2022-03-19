/*
Copyright 2022.

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

// HaproxyConfigSpec defines the desired state of HaproxyConfig
type HaproxyConfigSpec struct {
	// Config version
	Version int `json:"version,omitempty"`

	// Config data
	Data HaproxyConfigData `json:"data,omitempty"`
}

type HaproxyConfigData struct {
	Backends  []HaproxyConfigBackend  `json:"backends"`
	Frontends []HaproxyConfigFrontend `json:"frontends"`
}

type HaproxyConfigBackend struct {
	Name    string                `json:"name"`
	Servers []HaproxyConfigServer `json:"servers"`
}

type HaproxyConfigServer struct {
	Name    string `json:"name"`
	Port    int    `json:"port"`
	Address string `json:"address,omitempty"`
}

type HaproxyConfigFrontend struct {
	Name    string `json:"name"`
	Backend string `json:"backend"`
	Host    string `json:"host"`
}

// HaproxyConfigStatus defines the observed state of HaproxyConfig
type HaproxyConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:JSONPath=".spec.version",name="Version",type="integer"

// HaproxyConfig is the Schema for the haproxyconfigs API
type HaproxyConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HaproxyConfigSpec   `json:"spec,omitempty"`
	Status HaproxyConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HaproxyConfigList contains a list of HaproxyConfig
type HaproxyConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HaproxyConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HaproxyConfig{}, &HaproxyConfigList{})
}
