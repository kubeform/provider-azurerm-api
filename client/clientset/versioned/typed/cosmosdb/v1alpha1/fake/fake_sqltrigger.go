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

// FakeSqlTriggers implements SqlTriggerInterface
type FakeSqlTriggers struct {
	Fake *FakeCosmosdbV1alpha1
	ns   string
}

var sqltriggersResource = schema.GroupVersionResource{Group: "cosmosdb.azurerm.kubeform.com", Version: "v1alpha1", Resource: "sqltriggers"}

var sqltriggersKind = schema.GroupVersionKind{Group: "cosmosdb.azurerm.kubeform.com", Version: "v1alpha1", Kind: "SqlTrigger"}

// Get takes name of the sqlTrigger, and returns the corresponding sqlTrigger object, and an error if there is any.
func (c *FakeSqlTriggers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SqlTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sqltriggersResource, c.ns, name), &v1alpha1.SqlTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlTrigger), err
}

// List takes label and field selectors, and returns the list of SqlTriggers that match those selectors.
func (c *FakeSqlTriggers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SqlTriggerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sqltriggersResource, sqltriggersKind, c.ns, opts), &v1alpha1.SqlTriggerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SqlTriggerList{ListMeta: obj.(*v1alpha1.SqlTriggerList).ListMeta}
	for _, item := range obj.(*v1alpha1.SqlTriggerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sqlTriggers.
func (c *FakeSqlTriggers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sqltriggersResource, c.ns, opts))

}

// Create takes the representation of a sqlTrigger and creates it.  Returns the server's representation of the sqlTrigger, and an error, if there is any.
func (c *FakeSqlTriggers) Create(ctx context.Context, sqlTrigger *v1alpha1.SqlTrigger, opts v1.CreateOptions) (result *v1alpha1.SqlTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sqltriggersResource, c.ns, sqlTrigger), &v1alpha1.SqlTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlTrigger), err
}

// Update takes the representation of a sqlTrigger and updates it. Returns the server's representation of the sqlTrigger, and an error, if there is any.
func (c *FakeSqlTriggers) Update(ctx context.Context, sqlTrigger *v1alpha1.SqlTrigger, opts v1.UpdateOptions) (result *v1alpha1.SqlTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sqltriggersResource, c.ns, sqlTrigger), &v1alpha1.SqlTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlTrigger), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSqlTriggers) UpdateStatus(ctx context.Context, sqlTrigger *v1alpha1.SqlTrigger, opts v1.UpdateOptions) (*v1alpha1.SqlTrigger, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sqltriggersResource, "status", c.ns, sqlTrigger), &v1alpha1.SqlTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlTrigger), err
}

// Delete takes name of the sqlTrigger and deletes it. Returns an error if one occurs.
func (c *FakeSqlTriggers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sqltriggersResource, c.ns, name), &v1alpha1.SqlTrigger{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSqlTriggers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sqltriggersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SqlTriggerList{})
	return err
}

// Patch applies the patch and returns the patched sqlTrigger.
func (c *FakeSqlTriggers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SqlTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sqltriggersResource, c.ns, name, pt, data, subresources...), &v1alpha1.SqlTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlTrigger), err
}
