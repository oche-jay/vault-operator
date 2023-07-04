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

// KubernetesUnsealConfigApplyConfiguration represents an declarative configuration of the KubernetesUnsealConfig type for use
// with apply.
type KubernetesUnsealConfigApplyConfiguration struct {
	SecretNamespace *string `json:"secretNamespace,omitempty"`
	SecretName      *string `json:"secretName,omitempty"`
}

// KubernetesUnsealConfigApplyConfiguration constructs an declarative configuration of the KubernetesUnsealConfig type for use with
// apply.
func KubernetesUnsealConfig() *KubernetesUnsealConfigApplyConfiguration {
	return &KubernetesUnsealConfigApplyConfiguration{}
}

// WithSecretNamespace sets the SecretNamespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretNamespace field is set to the value of the last call.
func (b *KubernetesUnsealConfigApplyConfiguration) WithSecretNamespace(value string) *KubernetesUnsealConfigApplyConfiguration {
	b.SecretNamespace = &value
	return b
}

// WithSecretName sets the SecretName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretName field is set to the value of the last call.
func (b *KubernetesUnsealConfigApplyConfiguration) WithSecretName(value string) *KubernetesUnsealConfigApplyConfiguration {
	b.SecretName = &value
	return b
}
