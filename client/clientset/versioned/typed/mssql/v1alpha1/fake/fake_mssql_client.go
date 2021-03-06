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
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/clientset/versioned/typed/mssql/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMssqlV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMssqlV1alpha1) Databases(namespace string) v1alpha1.DatabaseInterface {
	return &FakeDatabases{c, namespace}
}

func (c *FakeMssqlV1alpha1) DatabaseExtendedAuditingPolicies(namespace string) v1alpha1.DatabaseExtendedAuditingPolicyInterface {
	return &FakeDatabaseExtendedAuditingPolicies{c, namespace}
}

func (c *FakeMssqlV1alpha1) DatabaseVulnerabilityAssessmentRuleBaselines(namespace string) v1alpha1.DatabaseVulnerabilityAssessmentRuleBaselineInterface {
	return &FakeDatabaseVulnerabilityAssessmentRuleBaselines{c, namespace}
}

func (c *FakeMssqlV1alpha1) Elasticpools(namespace string) v1alpha1.ElasticpoolInterface {
	return &FakeElasticpools{c, namespace}
}

func (c *FakeMssqlV1alpha1) FailoverGroups(namespace string) v1alpha1.FailoverGroupInterface {
	return &FakeFailoverGroups{c, namespace}
}

func (c *FakeMssqlV1alpha1) FirewallRules(namespace string) v1alpha1.FirewallRuleInterface {
	return &FakeFirewallRules{c, namespace}
}

func (c *FakeMssqlV1alpha1) JobAgents(namespace string) v1alpha1.JobAgentInterface {
	return &FakeJobAgents{c, namespace}
}

func (c *FakeMssqlV1alpha1) JobCredentials(namespace string) v1alpha1.JobCredentialInterface {
	return &FakeJobCredentials{c, namespace}
}

func (c *FakeMssqlV1alpha1) ManagedDatabases(namespace string) v1alpha1.ManagedDatabaseInterface {
	return &FakeManagedDatabases{c, namespace}
}

func (c *FakeMssqlV1alpha1) ManagedInstances(namespace string) v1alpha1.ManagedInstanceInterface {
	return &FakeManagedInstances{c, namespace}
}

func (c *FakeMssqlV1alpha1) ManagedInstanceActiveDirectoryAdministrators(namespace string) v1alpha1.ManagedInstanceActiveDirectoryAdministratorInterface {
	return &FakeManagedInstanceActiveDirectoryAdministrators{c, namespace}
}

func (c *FakeMssqlV1alpha1) ManagedInstanceFailoverGroups(namespace string) v1alpha1.ManagedInstanceFailoverGroupInterface {
	return &FakeManagedInstanceFailoverGroups{c, namespace}
}

func (c *FakeMssqlV1alpha1) OutboundFirewallRules(namespace string) v1alpha1.OutboundFirewallRuleInterface {
	return &FakeOutboundFirewallRules{c, namespace}
}

func (c *FakeMssqlV1alpha1) Servers(namespace string) v1alpha1.ServerInterface {
	return &FakeServers{c, namespace}
}

func (c *FakeMssqlV1alpha1) ServerExtendedAuditingPolicies(namespace string) v1alpha1.ServerExtendedAuditingPolicyInterface {
	return &FakeServerExtendedAuditingPolicies{c, namespace}
}

func (c *FakeMssqlV1alpha1) ServerSecurityAlertPolicies(namespace string) v1alpha1.ServerSecurityAlertPolicyInterface {
	return &FakeServerSecurityAlertPolicies{c, namespace}
}

func (c *FakeMssqlV1alpha1) ServerTransparentDataEncryptions(namespace string) v1alpha1.ServerTransparentDataEncryptionInterface {
	return &FakeServerTransparentDataEncryptions{c, namespace}
}

func (c *FakeMssqlV1alpha1) ServerVulnerabilityAssessments(namespace string) v1alpha1.ServerVulnerabilityAssessmentInterface {
	return &FakeServerVulnerabilityAssessments{c, namespace}
}

func (c *FakeMssqlV1alpha1) VirtualMachines(namespace string) v1alpha1.VirtualMachineInterface {
	return &FakeVirtualMachines{c, namespace}
}

func (c *FakeMssqlV1alpha1) VirtualNetworkRules(namespace string) v1alpha1.VirtualNetworkRuleInterface {
	return &FakeVirtualNetworkRules{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMssqlV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
