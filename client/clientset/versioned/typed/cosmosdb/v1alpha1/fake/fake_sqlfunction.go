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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/cosmosdb/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSqlFunctions implements SqlFunctionInterface
type FakeSqlFunctions struct {
	Fake *FakeCosmosdbV1alpha1
	ns   string
}

var sqlfunctionsResource = schema.GroupVersionResource{Group: "cosmosdb.azurerm.kubeform.com", Version: "v1alpha1", Resource: "sqlfunctions"}

var sqlfunctionsKind = schema.GroupVersionKind{Group: "cosmosdb.azurerm.kubeform.com", Version: "v1alpha1", Kind: "SqlFunction"}

// Get takes name of the sqlFunction, and returns the corresponding sqlFunction object, and an error if there is any.
func (c *FakeSqlFunctions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SqlFunction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sqlfunctionsResource, c.ns, name), &v1alpha1.SqlFunction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlFunction), err
}

// List takes label and field selectors, and returns the list of SqlFunctions that match those selectors.
func (c *FakeSqlFunctions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SqlFunctionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sqlfunctionsResource, sqlfunctionsKind, c.ns, opts), &v1alpha1.SqlFunctionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SqlFunctionList{ListMeta: obj.(*v1alpha1.SqlFunctionList).ListMeta}
	for _, item := range obj.(*v1alpha1.SqlFunctionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sqlFunctions.
func (c *FakeSqlFunctions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sqlfunctionsResource, c.ns, opts))

}

// Create takes the representation of a sqlFunction and creates it.  Returns the server's representation of the sqlFunction, and an error, if there is any.
func (c *FakeSqlFunctions) Create(ctx context.Context, sqlFunction *v1alpha1.SqlFunction, opts v1.CreateOptions) (result *v1alpha1.SqlFunction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sqlfunctionsResource, c.ns, sqlFunction), &v1alpha1.SqlFunction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlFunction), err
}

// Update takes the representation of a sqlFunction and updates it. Returns the server's representation of the sqlFunction, and an error, if there is any.
func (c *FakeSqlFunctions) Update(ctx context.Context, sqlFunction *v1alpha1.SqlFunction, opts v1.UpdateOptions) (result *v1alpha1.SqlFunction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sqlfunctionsResource, c.ns, sqlFunction), &v1alpha1.SqlFunction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlFunction), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSqlFunctions) UpdateStatus(ctx context.Context, sqlFunction *v1alpha1.SqlFunction, opts v1.UpdateOptions) (*v1alpha1.SqlFunction, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sqlfunctionsResource, "status", c.ns, sqlFunction), &v1alpha1.SqlFunction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlFunction), err
}

// Delete takes name of the sqlFunction and deletes it. Returns an error if one occurs.
func (c *FakeSqlFunctions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sqlfunctionsResource, c.ns, name), &v1alpha1.SqlFunction{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSqlFunctions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sqlfunctionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SqlFunctionList{})
	return err
}

// Patch applies the patch and returns the patched sqlFunction.
func (c *FakeSqlFunctions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SqlFunction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sqlfunctionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SqlFunction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlFunction), err
}
