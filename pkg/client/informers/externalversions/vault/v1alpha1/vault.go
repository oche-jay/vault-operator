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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	vaultv1alpha1 "github.com/bank-vaults/vault-operator/pkg/apis/vault/v1alpha1"
	versioned "github.com/bank-vaults/vault-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/bank-vaults/vault-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/bank-vaults/vault-operator/pkg/client/listers/vault/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VaultInformer provides access to a shared informer and lister for
// Vaults.
type VaultInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.VaultLister
}

type vaultInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVaultInformer constructs a new informer for Vault type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVaultInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVaultInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVaultInformer constructs a new informer for Vault type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVaultInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VaultV1alpha1().Vaults(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VaultV1alpha1().Vaults(namespace).Watch(context.TODO(), options)
			},
		},
		&vaultv1alpha1.Vault{},
		resyncPeriod,
		indexers,
	)
}

func (f *vaultInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVaultInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *vaultInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&vaultv1alpha1.Vault{}, f.defaultInformer)
}

func (f *vaultInformer) Lister() v1alpha1.VaultLister {
	return v1alpha1.NewVaultLister(f.Informer().GetIndexer())
}
