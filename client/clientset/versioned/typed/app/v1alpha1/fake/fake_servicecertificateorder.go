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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/app/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceCertificateOrders implements ServiceCertificateOrderInterface
type FakeServiceCertificateOrders struct {
	Fake *FakeAppV1alpha1
	ns   string
}

var servicecertificateordersResource = schema.GroupVersionResource{Group: "app.azurerm.kubeform.com", Version: "v1alpha1", Resource: "servicecertificateorders"}

var servicecertificateordersKind = schema.GroupVersionKind{Group: "app.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ServiceCertificateOrder"}

// Get takes name of the serviceCertificateOrder, and returns the corresponding serviceCertificateOrder object, and an error if there is any.
func (c *FakeServiceCertificateOrders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceCertificateOrder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicecertificateordersResource, c.ns, name), &v1alpha1.ServiceCertificateOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceCertificateOrder), err
}

// List takes label and field selectors, and returns the list of ServiceCertificateOrders that match those selectors.
func (c *FakeServiceCertificateOrders) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceCertificateOrderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicecertificateordersResource, servicecertificateordersKind, c.ns, opts), &v1alpha1.ServiceCertificateOrderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServiceCertificateOrderList{ListMeta: obj.(*v1alpha1.ServiceCertificateOrderList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServiceCertificateOrderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceCertificateOrders.
func (c *FakeServiceCertificateOrders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicecertificateordersResource, c.ns, opts))

}

// Create takes the representation of a serviceCertificateOrder and creates it.  Returns the server's representation of the serviceCertificateOrder, and an error, if there is any.
func (c *FakeServiceCertificateOrders) Create(ctx context.Context, serviceCertificateOrder *v1alpha1.ServiceCertificateOrder, opts v1.CreateOptions) (result *v1alpha1.ServiceCertificateOrder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicecertificateordersResource, c.ns, serviceCertificateOrder), &v1alpha1.ServiceCertificateOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceCertificateOrder), err
}

// Update takes the representation of a serviceCertificateOrder and updates it. Returns the server's representation of the serviceCertificateOrder, and an error, if there is any.
func (c *FakeServiceCertificateOrders) Update(ctx context.Context, serviceCertificateOrder *v1alpha1.ServiceCertificateOrder, opts v1.UpdateOptions) (result *v1alpha1.ServiceCertificateOrder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicecertificateordersResource, c.ns, serviceCertificateOrder), &v1alpha1.ServiceCertificateOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceCertificateOrder), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceCertificateOrders) UpdateStatus(ctx context.Context, serviceCertificateOrder *v1alpha1.ServiceCertificateOrder, opts v1.UpdateOptions) (*v1alpha1.ServiceCertificateOrder, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicecertificateordersResource, "status", c.ns, serviceCertificateOrder), &v1alpha1.ServiceCertificateOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceCertificateOrder), err
}

// Delete takes name of the serviceCertificateOrder and deletes it. Returns an error if one occurs.
func (c *FakeServiceCertificateOrders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicecertificateordersResource, c.ns, name), &v1alpha1.ServiceCertificateOrder{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceCertificateOrders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicecertificateordersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServiceCertificateOrderList{})
	return err
}

// Patch applies the patch and returns the patched serviceCertificateOrder.
func (c *FakeServiceCertificateOrders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceCertificateOrder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicecertificateordersResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServiceCertificateOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceCertificateOrder), err
}
