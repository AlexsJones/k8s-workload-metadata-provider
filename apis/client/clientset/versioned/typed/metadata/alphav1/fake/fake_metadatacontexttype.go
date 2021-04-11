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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	alphav1 "github.com/AlexsJones/k8s-workload-metadata-provider/apis/metadata/alphav1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMetaDataContextTypes implements MetaDataContextTypeInterface
type FakeMetaDataContextTypes struct {
	Fake *FakeMetadataAlphav1
	ns   string
}

var metadatacontexttypesResource = schema.GroupVersionResource{Group: "metadata", Version: "alphav1", Resource: "metadatacontexttypes"}

var metadatacontexttypesKind = schema.GroupVersionKind{Group: "metadata", Version: "alphav1", Kind: "MetaDataContextType"}

// Get takes name of the metaDataContextType, and returns the corresponding metaDataContextType object, and an error if there is any.
func (c *FakeMetaDataContextTypes) Get(ctx context.Context, name string, options v1.GetOptions) (result *alphav1.MetaDataContextType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(metadatacontexttypesResource, c.ns, name), &alphav1.MetaDataContextType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*alphav1.MetaDataContextType), err
}

// List takes label and field selectors, and returns the list of MetaDataContextTypes that match those selectors.
func (c *FakeMetaDataContextTypes) List(ctx context.Context, opts v1.ListOptions) (result *alphav1.MetaDataContextTypeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(metadatacontexttypesResource, metadatacontexttypesKind, c.ns, opts), &alphav1.MetaDataContextTypeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &alphav1.MetaDataContextTypeList{ListMeta: obj.(*alphav1.MetaDataContextTypeList).ListMeta}
	for _, item := range obj.(*alphav1.MetaDataContextTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested metaDataContextTypes.
func (c *FakeMetaDataContextTypes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(metadatacontexttypesResource, c.ns, opts))

}

// Create takes the representation of a metaDataContextType and creates it.  Returns the server's representation of the metaDataContextType, and an error, if there is any.
func (c *FakeMetaDataContextTypes) Create(ctx context.Context, metaDataContextType *alphav1.MetaDataContextType, opts v1.CreateOptions) (result *alphav1.MetaDataContextType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(metadatacontexttypesResource, c.ns, metaDataContextType), &alphav1.MetaDataContextType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*alphav1.MetaDataContextType), err
}

// Update takes the representation of a metaDataContextType and updates it. Returns the server's representation of the metaDataContextType, and an error, if there is any.
func (c *FakeMetaDataContextTypes) Update(ctx context.Context, metaDataContextType *alphav1.MetaDataContextType, opts v1.UpdateOptions) (result *alphav1.MetaDataContextType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(metadatacontexttypesResource, c.ns, metaDataContextType), &alphav1.MetaDataContextType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*alphav1.MetaDataContextType), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMetaDataContextTypes) UpdateStatus(ctx context.Context, metaDataContextType *alphav1.MetaDataContextType, opts v1.UpdateOptions) (*alphav1.MetaDataContextType, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(metadatacontexttypesResource, "status", c.ns, metaDataContextType), &alphav1.MetaDataContextType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*alphav1.MetaDataContextType), err
}

// Delete takes name of the metaDataContextType and deletes it. Returns an error if one occurs.
func (c *FakeMetaDataContextTypes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(metadatacontexttypesResource, c.ns, name), &alphav1.MetaDataContextType{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMetaDataContextTypes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(metadatacontexttypesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &alphav1.MetaDataContextTypeList{})
	return err
}

// Patch applies the patch and returns the patched metaDataContextType.
func (c *FakeMetaDataContextTypes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *alphav1.MetaDataContextType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(metadatacontexttypesResource, c.ns, name, pt, data, subresources...), &alphav1.MetaDataContextType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*alphav1.MetaDataContextType), err
}
