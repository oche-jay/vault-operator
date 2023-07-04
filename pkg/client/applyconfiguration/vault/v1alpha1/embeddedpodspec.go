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

// EmbeddedPodSpecApplyConfiguration represents an declarative configuration of the EmbeddedPodSpec type for use
// with apply.
type EmbeddedPodSpecApplyConfiguration struct {
	Volumes                       []v1.Volume                   `json:"volumes,omitempty"`
	InitContainers                []v1.Container                `json:"initContainers,omitempty"`
	Containers                    []v1.Container                `json:"containers,omitempty"`
	EphemeralContainers           []v1.EphemeralContainer       `json:"ephemeralContainers,omitempty"`
	RestartPolicy                 *v1.RestartPolicy             `json:"restartPolicy,omitempty"`
	TerminationGracePeriodSeconds *int64                        `json:"terminationGracePeriodSeconds,omitempty"`
	ActiveDeadlineSeconds         *int64                        `json:"activeDeadlineSeconds,omitempty"`
	DNSPolicy                     *v1.DNSPolicy                 `json:"dnsPolicy,omitempty"`
	NodeSelector                  map[string]string             `json:"nodeSelector,omitempty"`
	ServiceAccountName            *string                       `json:"serviceAccountName,omitempty"`
	DeprecatedServiceAccount      *string                       `json:"serviceAccount,omitempty"`
	AutomountServiceAccountToken  *bool                         `json:"automountServiceAccountToken,omitempty"`
	NodeName                      *string                       `json:"nodeName,omitempty"`
	HostNetwork                   *bool                         `json:"hostNetwork,omitempty"`
	HostPID                       *bool                         `json:"hostPID,omitempty"`
	HostIPC                       *bool                         `json:"hostIPC,omitempty"`
	ShareProcessNamespace         *bool                         `json:"shareProcessNamespace,omitempty"`
	SecurityContext               *v1.PodSecurityContext        `json:"securityContext,omitempty"`
	ImagePullSecrets              []v1.LocalObjectReference     `json:"imagePullSecrets,omitempty"`
	Hostname                      *string                       `json:"hostname,omitempty"`
	Subdomain                     *string                       `json:"subdomain,omitempty"`
	Affinity                      *v1.Affinity                  `json:"affinity,omitempty"`
	SchedulerName                 *string                       `json:"schedulerName,omitempty"`
	Tolerations                   []v1.Toleration               `json:"tolerations,omitempty"`
	HostAliases                   []v1.HostAlias                `json:"hostAliases,omitempty"`
	PriorityClassName             *string                       `json:"priorityClassName,omitempty"`
	Priority                      *int32                        `json:"priority,omitempty"`
	DNSConfig                     *v1.PodDNSConfig              `json:"dnsConfig,omitempty"`
	ReadinessGates                []v1.PodReadinessGate         `json:"readinessGates,omitempty"`
	RuntimeClassName              *string                       `json:"runtimeClassName,omitempty"`
	EnableServiceLinks            *bool                         `json:"enableServiceLinks,omitempty"`
	PreemptionPolicy              *v1.PreemptionPolicy          `json:"preemptionPolicy,omitempty"`
	Overhead                      *v1.ResourceList              `json:"overhead,omitempty"`
	TopologySpreadConstraints     []v1.TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty"`
	SetHostnameAsFQDN             *bool                         `json:"setHostnameAsFQDN,omitempty"`
	OS                            *v1.PodOS                     `json:"os,omitempty"`
	HostUsers                     *bool                         `json:"hostUsers,omitempty"`
	SchedulingGates               []v1.PodSchedulingGate        `json:"schedulingGates,omitempty"`
	ResourceClaims                []v1.PodResourceClaim         `json:"resourceClaims,omitempty"`
}

// EmbeddedPodSpecApplyConfiguration constructs an declarative configuration of the EmbeddedPodSpec type for use with
// apply.
func EmbeddedPodSpec() *EmbeddedPodSpecApplyConfiguration {
	return &EmbeddedPodSpecApplyConfiguration{}
}

// WithVolumes adds the given value to the Volumes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Volumes field.
func (b *EmbeddedPodSpecApplyConfiguration) WithVolumes(values ...v1.Volume) *EmbeddedPodSpecApplyConfiguration {
	for i := range values {
		b.Volumes = append(b.Volumes, values[i])
	}
	return b
}

// WithInitContainers adds the given value to the InitContainers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the InitContainers field.
func (b *EmbeddedPodSpecApplyConfiguration) WithInitContainers(values ...v1.Container) *EmbeddedPodSpecApplyConfiguration {
	for i := range values {
		b.InitContainers = append(b.InitContainers, values[i])
	}
	return b
}

// WithContainers adds the given value to the Containers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Containers field.
func (b *EmbeddedPodSpecApplyConfiguration) WithContainers(values ...v1.Container) *EmbeddedPodSpecApplyConfiguration {
	for i := range values {
		b.Containers = append(b.Containers, values[i])
	}
	return b
}

// WithEphemeralContainers adds the given value to the EphemeralContainers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the EphemeralContainers field.
func (b *EmbeddedPodSpecApplyConfiguration) WithEphemeralContainers(values ...v1.EphemeralContainer) *EmbeddedPodSpecApplyConfiguration {
	for i := range values {
		b.EphemeralContainers = append(b.EphemeralContainers, values[i])
	}
	return b
}

// WithRestartPolicy sets the RestartPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RestartPolicy field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithRestartPolicy(value v1.RestartPolicy) *EmbeddedPodSpecApplyConfiguration {
	b.RestartPolicy = &value
	return b
}

// WithTerminationGracePeriodSeconds sets the TerminationGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TerminationGracePeriodSeconds field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithTerminationGracePeriodSeconds(value int64) *EmbeddedPodSpecApplyConfiguration {
	b.TerminationGracePeriodSeconds = &value
	return b
}

// WithActiveDeadlineSeconds sets the ActiveDeadlineSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ActiveDeadlineSeconds field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithActiveDeadlineSeconds(value int64) *EmbeddedPodSpecApplyConfiguration {
	b.ActiveDeadlineSeconds = &value
	return b
}

// WithDNSPolicy sets the DNSPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DNSPolicy field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithDNSPolicy(value v1.DNSPolicy) *EmbeddedPodSpecApplyConfiguration {
	b.DNSPolicy = &value
	return b
}

// WithNodeSelector puts the entries into the NodeSelector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the NodeSelector field,
// overwriting an existing map entries in NodeSelector field with the same key.
func (b *EmbeddedPodSpecApplyConfiguration) WithNodeSelector(entries map[string]string) *EmbeddedPodSpecApplyConfiguration {
	if b.NodeSelector == nil && len(entries) > 0 {
		b.NodeSelector = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.NodeSelector[k] = v
	}
	return b
}

// WithServiceAccountName sets the ServiceAccountName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceAccountName field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithServiceAccountName(value string) *EmbeddedPodSpecApplyConfiguration {
	b.ServiceAccountName = &value
	return b
}

// WithDeprecatedServiceAccount sets the DeprecatedServiceAccount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeprecatedServiceAccount field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithDeprecatedServiceAccount(value string) *EmbeddedPodSpecApplyConfiguration {
	b.DeprecatedServiceAccount = &value
	return b
}

// WithAutomountServiceAccountToken sets the AutomountServiceAccountToken field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AutomountServiceAccountToken field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithAutomountServiceAccountToken(value bool) *EmbeddedPodSpecApplyConfiguration {
	b.AutomountServiceAccountToken = &value
	return b
}

// WithNodeName sets the NodeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeName field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithNodeName(value string) *EmbeddedPodSpecApplyConfiguration {
	b.NodeName = &value
	return b
}

// WithHostNetwork sets the HostNetwork field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostNetwork field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithHostNetwork(value bool) *EmbeddedPodSpecApplyConfiguration {
	b.HostNetwork = &value
	return b
}

// WithHostPID sets the HostPID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostPID field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithHostPID(value bool) *EmbeddedPodSpecApplyConfiguration {
	b.HostPID = &value
	return b
}

// WithHostIPC sets the HostIPC field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostIPC field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithHostIPC(value bool) *EmbeddedPodSpecApplyConfiguration {
	b.HostIPC = &value
	return b
}

// WithShareProcessNamespace sets the ShareProcessNamespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ShareProcessNamespace field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithShareProcessNamespace(value bool) *EmbeddedPodSpecApplyConfiguration {
	b.ShareProcessNamespace = &value
	return b
}

// WithSecurityContext sets the SecurityContext field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecurityContext field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithSecurityContext(value v1.PodSecurityContext) *EmbeddedPodSpecApplyConfiguration {
	b.SecurityContext = &value
	return b
}

// WithImagePullSecrets adds the given value to the ImagePullSecrets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ImagePullSecrets field.
func (b *EmbeddedPodSpecApplyConfiguration) WithImagePullSecrets(values ...v1.LocalObjectReference) *EmbeddedPodSpecApplyConfiguration {
	for i := range values {
		b.ImagePullSecrets = append(b.ImagePullSecrets, values[i])
	}
	return b
}

// WithHostname sets the Hostname field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Hostname field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithHostname(value string) *EmbeddedPodSpecApplyConfiguration {
	b.Hostname = &value
	return b
}

// WithSubdomain sets the Subdomain field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Subdomain field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithSubdomain(value string) *EmbeddedPodSpecApplyConfiguration {
	b.Subdomain = &value
	return b
}

// WithAffinity sets the Affinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affinity field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithAffinity(value v1.Affinity) *EmbeddedPodSpecApplyConfiguration {
	b.Affinity = &value
	return b
}

// WithSchedulerName sets the SchedulerName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SchedulerName field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithSchedulerName(value string) *EmbeddedPodSpecApplyConfiguration {
	b.SchedulerName = &value
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *EmbeddedPodSpecApplyConfiguration) WithTolerations(values ...v1.Toleration) *EmbeddedPodSpecApplyConfiguration {
	for i := range values {
		b.Tolerations = append(b.Tolerations, values[i])
	}
	return b
}

// WithHostAliases adds the given value to the HostAliases field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the HostAliases field.
func (b *EmbeddedPodSpecApplyConfiguration) WithHostAliases(values ...v1.HostAlias) *EmbeddedPodSpecApplyConfiguration {
	for i := range values {
		b.HostAliases = append(b.HostAliases, values[i])
	}
	return b
}

// WithPriorityClassName sets the PriorityClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PriorityClassName field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithPriorityClassName(value string) *EmbeddedPodSpecApplyConfiguration {
	b.PriorityClassName = &value
	return b
}

// WithPriority sets the Priority field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Priority field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithPriority(value int32) *EmbeddedPodSpecApplyConfiguration {
	b.Priority = &value
	return b
}

// WithDNSConfig sets the DNSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DNSConfig field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithDNSConfig(value v1.PodDNSConfig) *EmbeddedPodSpecApplyConfiguration {
	b.DNSConfig = &value
	return b
}

// WithReadinessGates adds the given value to the ReadinessGates field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ReadinessGates field.
func (b *EmbeddedPodSpecApplyConfiguration) WithReadinessGates(values ...v1.PodReadinessGate) *EmbeddedPodSpecApplyConfiguration {
	for i := range values {
		b.ReadinessGates = append(b.ReadinessGates, values[i])
	}
	return b
}

// WithRuntimeClassName sets the RuntimeClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RuntimeClassName field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithRuntimeClassName(value string) *EmbeddedPodSpecApplyConfiguration {
	b.RuntimeClassName = &value
	return b
}

// WithEnableServiceLinks sets the EnableServiceLinks field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableServiceLinks field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithEnableServiceLinks(value bool) *EmbeddedPodSpecApplyConfiguration {
	b.EnableServiceLinks = &value
	return b
}

// WithPreemptionPolicy sets the PreemptionPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PreemptionPolicy field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithPreemptionPolicy(value v1.PreemptionPolicy) *EmbeddedPodSpecApplyConfiguration {
	b.PreemptionPolicy = &value
	return b
}

// WithOverhead sets the Overhead field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Overhead field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithOverhead(value v1.ResourceList) *EmbeddedPodSpecApplyConfiguration {
	b.Overhead = &value
	return b
}

// WithTopologySpreadConstraints adds the given value to the TopologySpreadConstraints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TopologySpreadConstraints field.
func (b *EmbeddedPodSpecApplyConfiguration) WithTopologySpreadConstraints(values ...v1.TopologySpreadConstraint) *EmbeddedPodSpecApplyConfiguration {
	for i := range values {
		b.TopologySpreadConstraints = append(b.TopologySpreadConstraints, values[i])
	}
	return b
}

// WithSetHostnameAsFQDN sets the SetHostnameAsFQDN field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SetHostnameAsFQDN field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithSetHostnameAsFQDN(value bool) *EmbeddedPodSpecApplyConfiguration {
	b.SetHostnameAsFQDN = &value
	return b
}

// WithOS sets the OS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OS field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithOS(value v1.PodOS) *EmbeddedPodSpecApplyConfiguration {
	b.OS = &value
	return b
}

// WithHostUsers sets the HostUsers field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostUsers field is set to the value of the last call.
func (b *EmbeddedPodSpecApplyConfiguration) WithHostUsers(value bool) *EmbeddedPodSpecApplyConfiguration {
	b.HostUsers = &value
	return b
}

// WithSchedulingGates adds the given value to the SchedulingGates field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the SchedulingGates field.
func (b *EmbeddedPodSpecApplyConfiguration) WithSchedulingGates(values ...v1.PodSchedulingGate) *EmbeddedPodSpecApplyConfiguration {
	for i := range values {
		b.SchedulingGates = append(b.SchedulingGates, values[i])
	}
	return b
}

// WithResourceClaims adds the given value to the ResourceClaims field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ResourceClaims field.
func (b *EmbeddedPodSpecApplyConfiguration) WithResourceClaims(values ...v1.PodResourceClaim) *EmbeddedPodSpecApplyConfiguration {
	for i := range values {
		b.ResourceClaims = append(b.ResourceClaims, values[i])
	}
	return b
}
