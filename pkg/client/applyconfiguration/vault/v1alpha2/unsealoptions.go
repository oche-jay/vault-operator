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

package v1alpha2

// UnsealOptionsApplyConfiguration represents an declarative configuration of the UnsealOptions type for use
// with apply.
type UnsealOptionsApplyConfiguration struct {
	PreFlightChecks *bool `json:"preFlightChecks,omitempty"`
	StoreRootToken  *bool `json:"storeRootToken,omitempty"`
	SecretThreshold *uint `json:"secretThreshold,omitempty"`
	SecretShares    *uint `json:"secretShares,omitempty"`
}

// UnsealOptionsApplyConfiguration constructs an declarative configuration of the UnsealOptions type for use with
// apply.
func UnsealOptions() *UnsealOptionsApplyConfiguration {
	return &UnsealOptionsApplyConfiguration{}
}

// WithPreFlightChecks sets the PreFlightChecks field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PreFlightChecks field is set to the value of the last call.
func (b *UnsealOptionsApplyConfiguration) WithPreFlightChecks(value bool) *UnsealOptionsApplyConfiguration {
	b.PreFlightChecks = &value
	return b
}

// WithStoreRootToken sets the StoreRootToken field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StoreRootToken field is set to the value of the last call.
func (b *UnsealOptionsApplyConfiguration) WithStoreRootToken(value bool) *UnsealOptionsApplyConfiguration {
	b.StoreRootToken = &value
	return b
}

// WithSecretThreshold sets the SecretThreshold field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretThreshold field is set to the value of the last call.
func (b *UnsealOptionsApplyConfiguration) WithSecretThreshold(value uint) *UnsealOptionsApplyConfiguration {
	b.SecretThreshold = &value
	return b
}

// WithSecretShares sets the SecretShares field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretShares field is set to the value of the last call.
func (b *UnsealOptionsApplyConfiguration) WithSecretShares(value uint) *UnsealOptionsApplyConfiguration {
	b.SecretShares = &value
	return b
}
