/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package alphav1

import (
	alphav1 "github.com/AlexsJones/k8s-workload-metadata-provider/apis/metadata/alphav1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MetaDataContextTypeLister helps list MetaDataContextTypes.
// All objects returned here must be treated as read-only.
type MetaDataContextTypeLister interface {
	// List lists all MetaDataContextTypes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*alphav1.MetaDataContextType, err error)
	// MetaDataContextTypes returns an object that can list and get MetaDataContextTypes.
	MetaDataContextTypes(namespace string) MetaDataContextTypeNamespaceLister
	MetaDataContextTypeListerExpansion
}

// metaDataContextTypeLister implements the MetaDataContextTypeLister interface.
type metaDataContextTypeLister struct {
	indexer cache.Indexer
}

// NewMetaDataContextTypeLister returns a new MetaDataContextTypeLister.
func NewMetaDataContextTypeLister(indexer cache.Indexer) MetaDataContextTypeLister {
	return &metaDataContextTypeLister{indexer: indexer}
}

// List lists all MetaDataContextTypes in the indexer.
func (s *metaDataContextTypeLister) List(selector labels.Selector) (ret []*alphav1.MetaDataContextType, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*alphav1.MetaDataContextType))
	})
	return ret, err
}

// MetaDataContextTypes returns an object that can list and get MetaDataContextTypes.
func (s *metaDataContextTypeLister) MetaDataContextTypes(namespace string) MetaDataContextTypeNamespaceLister {
	return metaDataContextTypeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MetaDataContextTypeNamespaceLister helps list and get MetaDataContextTypes.
// All objects returned here must be treated as read-only.
type MetaDataContextTypeNamespaceLister interface {
	// List lists all MetaDataContextTypes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*alphav1.MetaDataContextType, err error)
	// Get retrieves the MetaDataContextType from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*alphav1.MetaDataContextType, error)
	MetaDataContextTypeNamespaceListerExpansion
}

// metaDataContextTypeNamespaceLister implements the MetaDataContextTypeNamespaceLister
// interface.
type metaDataContextTypeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MetaDataContextTypes in the indexer for a given namespace.
func (s metaDataContextTypeNamespaceLister) List(selector labels.Selector) (ret []*alphav1.MetaDataContextType, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*alphav1.MetaDataContextType))
	})
	return ret, err
}

// Get retrieves the MetaDataContextType from the indexer for a given namespace and name.
func (s metaDataContextTypeNamespaceLister) Get(name string) (*alphav1.MetaDataContextType, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(alphav1.Resource("metadatacontexttype"), name)
	}
	return obj.(*alphav1.MetaDataContextType), nil
}
