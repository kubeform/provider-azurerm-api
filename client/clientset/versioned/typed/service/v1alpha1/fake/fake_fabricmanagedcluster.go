/*
Copyright AppsCode Inc. and Contributors

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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/service/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFabricManagedClusters implements FabricManagedClusterInterface
type FakeFabricManagedClusters struct {
	Fake *FakeServiceV1alpha1
	ns   string
}

var fabricmanagedclustersResource = schema.GroupVersionResource{Group: "service.azurerm.kubeform.com", Version: "v1alpha1", Resource: "fabricmanagedclusters"}

var fabricmanagedclustersKind = schema.GroupVersionKind{Group: "service.azurerm.kubeform.com", Version: "v1alpha1", Kind: "FabricManagedCluster"}

// Get takes name of the fabricManagedCluster, and returns the corresponding fabricManagedCluster object, and an error if there is any.
func (c *FakeFabricManagedClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FabricManagedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fabricmanagedclustersResource, c.ns, name), &v1alpha1.FabricManagedCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricManagedCluster), err
}

// List takes label and field selectors, and returns the list of FabricManagedClusters that match those selectors.
func (c *FakeFabricManagedClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FabricManagedClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fabricmanagedclustersResource, fabricmanagedclustersKind, c.ns, opts), &v1alpha1.FabricManagedClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FabricManagedClusterList{ListMeta: obj.(*v1alpha1.FabricManagedClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.FabricManagedClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fabricManagedClusters.
func (c *FakeFabricManagedClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fabricmanagedclustersResource, c.ns, opts))

}

// Create takes the representation of a fabricManagedCluster and creates it.  Returns the server's representation of the fabricManagedCluster, and an error, if there is any.
func (c *FakeFabricManagedClusters) Create(ctx context.Context, fabricManagedCluster *v1alpha1.FabricManagedCluster, opts v1.CreateOptions) (result *v1alpha1.FabricManagedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fabricmanagedclustersResource, c.ns, fabricManagedCluster), &v1alpha1.FabricManagedCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricManagedCluster), err
}

// Update takes the representation of a fabricManagedCluster and updates it. Returns the server's representation of the fabricManagedCluster, and an error, if there is any.
func (c *FakeFabricManagedClusters) Update(ctx context.Context, fabricManagedCluster *v1alpha1.FabricManagedCluster, opts v1.UpdateOptions) (result *v1alpha1.FabricManagedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fabricmanagedclustersResource, c.ns, fabricManagedCluster), &v1alpha1.FabricManagedCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricManagedCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFabricManagedClusters) UpdateStatus(ctx context.Context, fabricManagedCluster *v1alpha1.FabricManagedCluster, opts v1.UpdateOptions) (*v1alpha1.FabricManagedCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fabricmanagedclustersResource, "status", c.ns, fabricManagedCluster), &v1alpha1.FabricManagedCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricManagedCluster), err
}

// Delete takes name of the fabricManagedCluster and deletes it. Returns an error if one occurs.
func (c *FakeFabricManagedClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(fabricmanagedclustersResource, c.ns, name), &v1alpha1.FabricManagedCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFabricManagedClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fabricmanagedclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FabricManagedClusterList{})
	return err
}

// Patch applies the patch and returns the patched fabricManagedCluster.
func (c *FakeFabricManagedClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FabricManagedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fabricmanagedclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.FabricManagedCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricManagedCluster), err
}
