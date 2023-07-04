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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/bank-vaults/vault-operator/v2/pkg/apis/vault/v1alpha1"
	vaultv1alpha1 "github.com/bank-vaults/vault-operator/v2/pkg/client/applyconfiguration/vault/v1alpha1"
	scheme "github.com/bank-vaults/vault-operator/v2/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VaultsGetter has a method to return a VaultInterface.
// A group's client should implement this interface.
type VaultsGetter interface {
	Vaults(namespace string) VaultInterface
}

// VaultInterface has methods to work with Vault resources.
type VaultInterface interface {
	Create(ctx context.Context, vault *v1alpha1.Vault, opts v1.CreateOptions) (*v1alpha1.Vault, error)
	Update(ctx context.Context, vault *v1alpha1.Vault, opts v1.UpdateOptions) (*v1alpha1.Vault, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Vault, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VaultList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Vault, err error)
	Apply(ctx context.Context, vault *vaultv1alpha1.VaultApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Vault, err error)
	VaultExpansion
}

// vaults implements VaultInterface
type vaults struct {
	client rest.Interface
	ns     string
}

// newVaults returns a Vaults
func newVaults(c *VaultV1alpha1Client, namespace string) *vaults {
	return &vaults{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vault, and returns the corresponding vault object, and an error if there is any.
func (c *vaults) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Vault, err error) {
	result = &v1alpha1.Vault{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vaults").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Vaults that match those selectors.
func (c *vaults) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VaultList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VaultList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vaults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vaults.
func (c *vaults) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vaults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vault and creates it.  Returns the server's representation of the vault, and an error, if there is any.
func (c *vaults) Create(ctx context.Context, vault *v1alpha1.Vault, opts v1.CreateOptions) (result *v1alpha1.Vault, err error) {
	result = &v1alpha1.Vault{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vaults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vault).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vault and updates it. Returns the server's representation of the vault, and an error, if there is any.
func (c *vaults) Update(ctx context.Context, vault *v1alpha1.Vault, opts v1.UpdateOptions) (result *v1alpha1.Vault, err error) {
	result = &v1alpha1.Vault{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vaults").
		Name(vault.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vault).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vault and deletes it. Returns an error if one occurs.
func (c *vaults) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vaults").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vaults) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vaults").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vault.
func (c *vaults) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Vault, err error) {
	result = &v1alpha1.Vault{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vaults").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied vault.
func (c *vaults) Apply(ctx context.Context, vault *vaultv1alpha1.VaultApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Vault, err error) {
	if vault == nil {
		return nil, fmt.Errorf("vault provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	name := vault.Name
	if name == nil {
		return nil, fmt.Errorf("vault.Name must be provided to Apply")
	}
	result = &v1alpha1.Vault{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("vaults").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
