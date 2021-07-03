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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/sentinel/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAlertRuleMsSecurityIncidents implements AlertRuleMsSecurityIncidentInterface
type FakeAlertRuleMsSecurityIncidents struct {
	Fake *FakeSentinelV1alpha1
	ns   string
}

var alertrulemssecurityincidentsResource = schema.GroupVersionResource{Group: "sentinel.azurerm.kubeform.com", Version: "v1alpha1", Resource: "alertrulemssecurityincidents"}

var alertrulemssecurityincidentsKind = schema.GroupVersionKind{Group: "sentinel.azurerm.kubeform.com", Version: "v1alpha1", Kind: "AlertRuleMsSecurityIncident"}

// Get takes name of the alertRuleMsSecurityIncident, and returns the corresponding alertRuleMsSecurityIncident object, and an error if there is any.
func (c *FakeAlertRuleMsSecurityIncidents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AlertRuleMsSecurityIncident, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(alertrulemssecurityincidentsResource, c.ns, name), &v1alpha1.AlertRuleMsSecurityIncident{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertRuleMsSecurityIncident), err
}

// List takes label and field selectors, and returns the list of AlertRuleMsSecurityIncidents that match those selectors.
func (c *FakeAlertRuleMsSecurityIncidents) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AlertRuleMsSecurityIncidentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(alertrulemssecurityincidentsResource, alertrulemssecurityincidentsKind, c.ns, opts), &v1alpha1.AlertRuleMsSecurityIncidentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AlertRuleMsSecurityIncidentList{ListMeta: obj.(*v1alpha1.AlertRuleMsSecurityIncidentList).ListMeta}
	for _, item := range obj.(*v1alpha1.AlertRuleMsSecurityIncidentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alertRuleMsSecurityIncidents.
func (c *FakeAlertRuleMsSecurityIncidents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(alertrulemssecurityincidentsResource, c.ns, opts))

}

// Create takes the representation of a alertRuleMsSecurityIncident and creates it.  Returns the server's representation of the alertRuleMsSecurityIncident, and an error, if there is any.
func (c *FakeAlertRuleMsSecurityIncidents) Create(ctx context.Context, alertRuleMsSecurityIncident *v1alpha1.AlertRuleMsSecurityIncident, opts v1.CreateOptions) (result *v1alpha1.AlertRuleMsSecurityIncident, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(alertrulemssecurityincidentsResource, c.ns, alertRuleMsSecurityIncident), &v1alpha1.AlertRuleMsSecurityIncident{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertRuleMsSecurityIncident), err
}

// Update takes the representation of a alertRuleMsSecurityIncident and updates it. Returns the server's representation of the alertRuleMsSecurityIncident, and an error, if there is any.
func (c *FakeAlertRuleMsSecurityIncidents) Update(ctx context.Context, alertRuleMsSecurityIncident *v1alpha1.AlertRuleMsSecurityIncident, opts v1.UpdateOptions) (result *v1alpha1.AlertRuleMsSecurityIncident, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(alertrulemssecurityincidentsResource, c.ns, alertRuleMsSecurityIncident), &v1alpha1.AlertRuleMsSecurityIncident{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertRuleMsSecurityIncident), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAlertRuleMsSecurityIncidents) UpdateStatus(ctx context.Context, alertRuleMsSecurityIncident *v1alpha1.AlertRuleMsSecurityIncident, opts v1.UpdateOptions) (*v1alpha1.AlertRuleMsSecurityIncident, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(alertrulemssecurityincidentsResource, "status", c.ns, alertRuleMsSecurityIncident), &v1alpha1.AlertRuleMsSecurityIncident{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertRuleMsSecurityIncident), err
}

// Delete takes name of the alertRuleMsSecurityIncident and deletes it. Returns an error if one occurs.
func (c *FakeAlertRuleMsSecurityIncidents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(alertrulemssecurityincidentsResource, c.ns, name), &v1alpha1.AlertRuleMsSecurityIncident{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlertRuleMsSecurityIncidents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(alertrulemssecurityincidentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AlertRuleMsSecurityIncidentList{})
	return err
}

// Patch applies the patch and returns the patched alertRuleMsSecurityIncident.
func (c *FakeAlertRuleMsSecurityIncidents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AlertRuleMsSecurityIncident, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(alertrulemssecurityincidentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AlertRuleMsSecurityIncident{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertRuleMsSecurityIncident), err
}
