// Copyright © 2023 Bank-Vaults
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// VaultStatusApplyConfiguration represents an declarative configuration of the VaultStatus type for use
// with apply.
type VaultStatusApplyConfiguration struct {
	Nodes      []string                `json:"nodes,omitempty"`
	Leader     *string                 `json:"leader,omitempty"`
	Conditions []v1.ComponentCondition `json:"conditions,omitempty"`
}

// VaultStatusApplyConfiguration constructs an declarative configuration of the VaultStatus type for use with
// apply.
func VaultStatus() *VaultStatusApplyConfiguration {
	return &VaultStatusApplyConfiguration{}
}

// WithNodes adds the given value to the Nodes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Nodes field.
func (b *VaultStatusApplyConfiguration) WithNodes(values ...string) *VaultStatusApplyConfiguration {
	for i := range values {
		b.Nodes = append(b.Nodes, values[i])
	}
	return b
}

// WithLeader sets the Leader field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Leader field is set to the value of the last call.
func (b *VaultStatusApplyConfiguration) WithLeader(value string) *VaultStatusApplyConfiguration {
	b.Leader = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *VaultStatusApplyConfiguration) WithConditions(values ...v1.ComponentCondition) *VaultStatusApplyConfiguration {
	for i := range values {
		b.Conditions = append(b.Conditions, values[i])
	}
	return b
}
