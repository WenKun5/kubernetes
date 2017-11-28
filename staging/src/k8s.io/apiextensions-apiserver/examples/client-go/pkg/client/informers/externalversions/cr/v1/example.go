/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1

import (
	cr_v1 "k8s.io/apiextensions-apiserver/examples/client-go/pkg/apis/cr/v1"
	versioned "k8s.io/apiextensions-apiserver/examples/client-go/pkg/client/clientset/versioned"
	internalinterfaces "k8s.io/apiextensions-apiserver/examples/client-go/pkg/client/informers/externalversions/internalinterfaces"
	v1 "k8s.io/apiextensions-apiserver/examples/client-go/pkg/client/listers/cr/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ExampleInformer provides access to a shared informer and lister for
// Examples.
type ExampleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ExampleLister
}

type exampleInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewExampleInformer constructs a new informer for Example type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewExampleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.CrV1().Examples(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.CrV1().Examples(namespace).Watch(options)
			},
		},
		&cr_v1.Example{},
		resyncPeriod,
		indexers,
	)
}

func defaultExampleInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewExampleInformer(client, meta_v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *exampleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cr_v1.Example{}, defaultExampleInformer)
}

func (f *exampleInformer) Lister() v1.ExampleLister {
	return v1.NewExampleLister(f.Informer().GetIndexer())
}