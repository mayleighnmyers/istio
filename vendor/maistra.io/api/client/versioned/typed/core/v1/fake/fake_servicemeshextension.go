// Copyright Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	corev1 "maistra.io/api/core/v1"
)

// FakeServiceMeshExtensions implements ServiceMeshExtensionInterface
type FakeServiceMeshExtensions struct {
	Fake *FakeCoreV1
	ns   string
}

var servicemeshextensionsResource = schema.GroupVersionResource{Group: "maistra.io", Version: "v1", Resource: "servicemeshextensions"}

var servicemeshextensionsKind = schema.GroupVersionKind{Group: "maistra.io", Version: "v1", Kind: "ServiceMeshExtension"}

// Get takes name of the serviceMeshExtension, and returns the corresponding serviceMeshExtension object, and an error if there is any.
func (c *FakeServiceMeshExtensions) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.ServiceMeshExtension, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicemeshextensionsResource, c.ns, name), &corev1.ServiceMeshExtension{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceMeshExtension), err
}

// List takes label and field selectors, and returns the list of ServiceMeshExtensions that match those selectors.
func (c *FakeServiceMeshExtensions) List(ctx context.Context, opts v1.ListOptions) (result *corev1.ServiceMeshExtensionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicemeshextensionsResource, servicemeshextensionsKind, c.ns, opts), &corev1.ServiceMeshExtensionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.ServiceMeshExtensionList{ListMeta: obj.(*corev1.ServiceMeshExtensionList).ListMeta}
	for _, item := range obj.(*corev1.ServiceMeshExtensionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceMeshExtensions.
func (c *FakeServiceMeshExtensions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicemeshextensionsResource, c.ns, opts))

}

// Create takes the representation of a serviceMeshExtension and creates it.  Returns the server's representation of the serviceMeshExtension, and an error, if there is any.
func (c *FakeServiceMeshExtensions) Create(ctx context.Context, serviceMeshExtension *corev1.ServiceMeshExtension, opts v1.CreateOptions) (result *corev1.ServiceMeshExtension, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicemeshextensionsResource, c.ns, serviceMeshExtension), &corev1.ServiceMeshExtension{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceMeshExtension), err
}

// Update takes the representation of a serviceMeshExtension and updates it. Returns the server's representation of the serviceMeshExtension, and an error, if there is any.
func (c *FakeServiceMeshExtensions) Update(ctx context.Context, serviceMeshExtension *corev1.ServiceMeshExtension, opts v1.UpdateOptions) (result *corev1.ServiceMeshExtension, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicemeshextensionsResource, c.ns, serviceMeshExtension), &corev1.ServiceMeshExtension{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceMeshExtension), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceMeshExtensions) UpdateStatus(ctx context.Context, serviceMeshExtension *corev1.ServiceMeshExtension, opts v1.UpdateOptions) (*corev1.ServiceMeshExtension, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicemeshextensionsResource, "status", c.ns, serviceMeshExtension), &corev1.ServiceMeshExtension{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceMeshExtension), err
}

// Delete takes name of the serviceMeshExtension and deletes it. Returns an error if one occurs.
func (c *FakeServiceMeshExtensions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(servicemeshextensionsResource, c.ns, name, opts), &corev1.ServiceMeshExtension{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceMeshExtensions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicemeshextensionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.ServiceMeshExtensionList{})
	return err
}

// Patch applies the patch and returns the patched serviceMeshExtension.
func (c *FakeServiceMeshExtensions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.ServiceMeshExtension, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicemeshextensionsResource, c.ns, name, pt, data, subresources...), &corev1.ServiceMeshExtension{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceMeshExtension), err
}