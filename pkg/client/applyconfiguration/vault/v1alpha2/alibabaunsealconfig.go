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

// AlibabaUnsealConfigApplyConfiguration represents an declarative configuration of the AlibabaUnsealConfig type for use
// with apply.
type AlibabaUnsealConfigApplyConfiguration struct {
	KMSRegion   *string `json:"kmsRegion,omitempty"`
	KMSKeyID    *string `json:"kmsKeyId,omitempty"`
	OSSEndpoint *string `json:"ossEndpoint,omitempty"`
	OSSBucket   *string `json:"ossBucket,omitempty"`
	OSSPrefix   *string `json:"ossPrefix,omitempty"`
}

// AlibabaUnsealConfigApplyConfiguration constructs an declarative configuration of the AlibabaUnsealConfig type for use with
// apply.
func AlibabaUnsealConfig() *AlibabaUnsealConfigApplyConfiguration {
	return &AlibabaUnsealConfigApplyConfiguration{}
}

// WithKMSRegion sets the KMSRegion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KMSRegion field is set to the value of the last call.
func (b *AlibabaUnsealConfigApplyConfiguration) WithKMSRegion(value string) *AlibabaUnsealConfigApplyConfiguration {
	b.KMSRegion = &value
	return b
}

// WithKMSKeyID sets the KMSKeyID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KMSKeyID field is set to the value of the last call.
func (b *AlibabaUnsealConfigApplyConfiguration) WithKMSKeyID(value string) *AlibabaUnsealConfigApplyConfiguration {
	b.KMSKeyID = &value
	return b
}

// WithOSSEndpoint sets the OSSEndpoint field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OSSEndpoint field is set to the value of the last call.
func (b *AlibabaUnsealConfigApplyConfiguration) WithOSSEndpoint(value string) *AlibabaUnsealConfigApplyConfiguration {
	b.OSSEndpoint = &value
	return b
}

// WithOSSBucket sets the OSSBucket field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OSSBucket field is set to the value of the last call.
func (b *AlibabaUnsealConfigApplyConfiguration) WithOSSBucket(value string) *AlibabaUnsealConfigApplyConfiguration {
	b.OSSBucket = &value
	return b
}

// WithOSSPrefix sets the OSSPrefix field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OSSPrefix field is set to the value of the last call.
func (b *AlibabaUnsealConfigApplyConfiguration) WithOSSPrefix(value string) *AlibabaUnsealConfigApplyConfiguration {
	b.OSSPrefix = &value
	return b
}
