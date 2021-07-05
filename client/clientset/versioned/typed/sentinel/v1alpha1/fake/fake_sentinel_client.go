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
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/clientset/versioned/typed/sentinel/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSentinelV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSentinelV1alpha1) AlertRuleFusions(namespace string) v1alpha1.AlertRuleFusionInterface {
	return &FakeAlertRuleFusions{c, namespace}
}

func (c *FakeSentinelV1alpha1) AlertRuleMachineLearningBehaviorAnalyticses(namespace string) v1alpha1.AlertRuleMachineLearningBehaviorAnalyticsInterface {
	return &FakeAlertRuleMachineLearningBehaviorAnalyticses{c, namespace}
}

func (c *FakeSentinelV1alpha1) AlertRuleMsSecurityIncidents(namespace string) v1alpha1.AlertRuleMsSecurityIncidentInterface {
	return &FakeAlertRuleMsSecurityIncidents{c, namespace}
}

func (c *FakeSentinelV1alpha1) AlertRuleScheduleds(namespace string) v1alpha1.AlertRuleScheduledInterface {
	return &FakeAlertRuleScheduleds{c, namespace}
}

func (c *FakeSentinelV1alpha1) DataConnectorAwsCloudTrails(namespace string) v1alpha1.DataConnectorAwsCloudTrailInterface {
	return &FakeDataConnectorAwsCloudTrails{c, namespace}
}

func (c *FakeSentinelV1alpha1) DataConnectorAzureActiveDirectories(namespace string) v1alpha1.DataConnectorAzureActiveDirectoryInterface {
	return &FakeDataConnectorAzureActiveDirectories{c, namespace}
}

func (c *FakeSentinelV1alpha1) DataConnectorAzureAdvancedThreatProtections(namespace string) v1alpha1.DataConnectorAzureAdvancedThreatProtectionInterface {
	return &FakeDataConnectorAzureAdvancedThreatProtections{c, namespace}
}

func (c *FakeSentinelV1alpha1) DataConnectorAzureSecurityCenters(namespace string) v1alpha1.DataConnectorAzureSecurityCenterInterface {
	return &FakeDataConnectorAzureSecurityCenters{c, namespace}
}

func (c *FakeSentinelV1alpha1) DataConnectorMicrosoftCloudAppSecurities(namespace string) v1alpha1.DataConnectorMicrosoftCloudAppSecurityInterface {
	return &FakeDataConnectorMicrosoftCloudAppSecurities{c, namespace}
}

func (c *FakeSentinelV1alpha1) DataConnectorMicrosoftDefenderAdvancedThreatProtections(namespace string) v1alpha1.DataConnectorMicrosoftDefenderAdvancedThreatProtectionInterface {
	return &FakeDataConnectorMicrosoftDefenderAdvancedThreatProtections{c, namespace}
}

func (c *FakeSentinelV1alpha1) DataConnectorOffice365s(namespace string) v1alpha1.DataConnectorOffice365Interface {
	return &FakeDataConnectorOffice365s{c, namespace}
}

func (c *FakeSentinelV1alpha1) DataConnectorThreatIntelligences(namespace string) v1alpha1.DataConnectorThreatIntelligenceInterface {
	return &FakeDataConnectorThreatIntelligences{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSentinelV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}