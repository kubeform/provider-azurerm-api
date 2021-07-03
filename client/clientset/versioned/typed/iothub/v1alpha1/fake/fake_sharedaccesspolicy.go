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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/iothub/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSharedAccessPolicies implements SharedAccessPolicyInterface
type FakeSharedAccessPolicies struct {
	Fake *FakeIothubV1alpha1
	ns   string
}

var sharedaccesspoliciesResource = schema.GroupVersionResource{Group: "iothub.azurerm.kubeform.com", Version: "v1alpha1", Resource: "sharedaccesspolicies"}

var sharedaccesspoliciesKind = schema.GroupVersionKind{Group: "iothub.azurerm.kubeform.com", Version: "v1alpha1", Kind: "SharedAccessPolicy"}

// Get takes name of the sharedAccessPolicy, and returns the corresponding sharedAccessPolicy object, and an error if there is any.
func (c *FakeSharedAccessPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SharedAccessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sharedaccesspoliciesResource, c.ns, name), &v1alpha1.SharedAccessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedAccessPolicy), err
}

// List takes label and field selectors, and returns the list of SharedAccessPolicies that match those selectors.
func (c *FakeSharedAccessPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SharedAccessPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sharedaccesspoliciesResource, sharedaccesspoliciesKind, c.ns, opts), &v1alpha1.SharedAccessPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SharedAccessPolicyList{ListMeta: obj.(*v1alpha1.SharedAccessPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.SharedAccessPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sharedAccessPolicies.
func (c *FakeSharedAccessPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sharedaccesspoliciesResource, c.ns, opts))

}

// Create takes the representation of a sharedAccessPolicy and creates it.  Returns the server's representation of the sharedAccessPolicy, and an error, if there is any.
func (c *FakeSharedAccessPolicies) Create(ctx context.Context, sharedAccessPolicy *v1alpha1.SharedAccessPolicy, opts v1.CreateOptions) (result *v1alpha1.SharedAccessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sharedaccesspoliciesResource, c.ns, sharedAccessPolicy), &v1alpha1.SharedAccessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedAccessPolicy), err
}

// Update takes the representation of a sharedAccessPolicy and updates it. Returns the server's representation of the sharedAccessPolicy, and an error, if there is any.
func (c *FakeSharedAccessPolicies) Update(ctx context.Context, sharedAccessPolicy *v1alpha1.SharedAccessPolicy, opts v1.UpdateOptions) (result *v1alpha1.SharedAccessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sharedaccesspoliciesResource, c.ns, sharedAccessPolicy), &v1alpha1.SharedAccessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedAccessPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSharedAccessPolicies) UpdateStatus(ctx context.Context, sharedAccessPolicy *v1alpha1.SharedAccessPolicy, opts v1.UpdateOptions) (*v1alpha1.SharedAccessPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sharedaccesspoliciesResource, "status", c.ns, sharedAccessPolicy), &v1alpha1.SharedAccessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedAccessPolicy), err
}

// Delete takes name of the sharedAccessPolicy and deletes it. Returns an error if one occurs.
func (c *FakeSharedAccessPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sharedaccesspoliciesResource, c.ns, name), &v1alpha1.SharedAccessPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSharedAccessPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sharedaccesspoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SharedAccessPolicyList{})
	return err
}

// Patch applies the patch and returns the patched sharedAccessPolicy.
func (c *FakeSharedAccessPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SharedAccessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sharedaccesspoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SharedAccessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedAccessPolicy), err
}
