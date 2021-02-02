/*
Copyright 2021 Filecoin Infra Team.

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

type MinerConfig struct {
	// Configuration options used to generate the config file.
	// Could probably pull this struct out of lotus
}

// MinerSpec defines the desired state of Miner
type MinerSpec struct {
	Config      *MinerConfig `json:"Config"`
	OwnerWallet string       `json:"OwnerWallet"`
}

// Needs to be string for serializaiton
type WindowPostHour string

type WindowPostSchedule map[WindowPostHour]int

// MinerStatus defines the observed state of Miner
type MinerStatus struct {
	MinerId               string             `json:"MinerId"`
	Sectors               uint               `json:"Sectors"`
	WindowPostSchedule    WindowPostSchedule `json:"WindowPostSchedule"`
	CurrentWindowPostHour WindowPostHour     `json:"CurrentWindowPostHour"`
	// Maybe also current work going on, how many sealing, etc?
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Miner is the Schema for the miners API
type Miner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MinerSpec   `json:"spec,omitempty"`
	Status MinerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MinerList contains a list of Miner
type MinerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Miner `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Miner{}, &MinerList{})
}
