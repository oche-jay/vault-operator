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

import (
	v1 "k8s.io/api/networking/v1"
)

// IngressApplyConfiguration represents an declarative configuration of the Ingress type for use
// with apply.
type IngressApplyConfiguration struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	Spec        *v1.IngressSpec   `json:"spec,omitempty"`
}

// IngressApplyConfiguration constructs an declarative configuration of the Ingress type for use with
// apply.
func Ingress() *IngressApplyConfiguration {
	return &IngressApplyConfiguration{}
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *IngressApplyConfiguration) WithAnnotations(entries map[string]string) *IngressApplyConfiguration {
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *IngressApplyConfiguration) WithSpec(value v1.IngressSpec) *IngressApplyConfiguration {
	b.Spec = &value
	return b
}
