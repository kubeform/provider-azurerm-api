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

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/synapse/v1alpha1"
	"kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type SynapseV1alpha1Interface interface {
	RESTClient() rest.Interface
	FirewallRulesGetter
	IntegrationRuntimeAzuresGetter
	IntegrationRuntimeSelfHostedsGetter
	LinkedServicesGetter
	ManagedPrivateEndpointsGetter
	PrivateLinkHubsGetter
	RoleAssignmentsGetter
	SparkPoolsGetter
	SqlPoolsGetter
	SqlPoolExtendedAuditingPoliciesGetter
	SqlPoolSecurityAlertPoliciesGetter
	SqlPoolVulnerabilityAssessmentsGetter
	SqlPoolVulnerabilityAssessmentBaselinesGetter
	SqlPoolWorkloadClassifiersGetter
	SqlPoolWorkloadGroupsGetter
	WorkspacesGetter
	WorkspaceAadAdminsGetter
	WorkspaceExtendedAuditingPoliciesGetter
	WorkspaceKeysGetter
	WorkspaceSQLAadAdminsGetter
	WorkspaceSecurityAlertPoliciesGetter
	WorkspaceVulnerabilityAssessmentsGetter
}

// SynapseV1alpha1Client is used to interact with features provided by the synapse.azurerm.kubeform.com group.
type SynapseV1alpha1Client struct {
	restClient rest.Interface
}

func (c *SynapseV1alpha1Client) FirewallRules(namespace string) FirewallRuleInterface {
	return newFirewallRules(c, namespace)
}

func (c *SynapseV1alpha1Client) IntegrationRuntimeAzures(namespace string) IntegrationRuntimeAzureInterface {
	return newIntegrationRuntimeAzures(c, namespace)
}

func (c *SynapseV1alpha1Client) IntegrationRuntimeSelfHosteds(namespace string) IntegrationRuntimeSelfHostedInterface {
	return newIntegrationRuntimeSelfHosteds(c, namespace)
}

func (c *SynapseV1alpha1Client) LinkedServices(namespace string) LinkedServiceInterface {
	return newLinkedServices(c, namespace)
}

func (c *SynapseV1alpha1Client) ManagedPrivateEndpoints(namespace string) ManagedPrivateEndpointInterface {
	return newManagedPrivateEndpoints(c, namespace)
}

func (c *SynapseV1alpha1Client) PrivateLinkHubs(namespace string) PrivateLinkHubInterface {
	return newPrivateLinkHubs(c, namespace)
}

func (c *SynapseV1alpha1Client) RoleAssignments(namespace string) RoleAssignmentInterface {
	return newRoleAssignments(c, namespace)
}

func (c *SynapseV1alpha1Client) SparkPools(namespace string) SparkPoolInterface {
	return newSparkPools(c, namespace)
}

func (c *SynapseV1alpha1Client) SqlPools(namespace string) SqlPoolInterface {
	return newSqlPools(c, namespace)
}

func (c *SynapseV1alpha1Client) SqlPoolExtendedAuditingPolicies(namespace string) SqlPoolExtendedAuditingPolicyInterface {
	return newSqlPoolExtendedAuditingPolicies(c, namespace)
}

func (c *SynapseV1alpha1Client) SqlPoolSecurityAlertPolicies(namespace string) SqlPoolSecurityAlertPolicyInterface {
	return newSqlPoolSecurityAlertPolicies(c, namespace)
}

func (c *SynapseV1alpha1Client) SqlPoolVulnerabilityAssessments(namespace string) SqlPoolVulnerabilityAssessmentInterface {
	return newSqlPoolVulnerabilityAssessments(c, namespace)
}

func (c *SynapseV1alpha1Client) SqlPoolVulnerabilityAssessmentBaselines(namespace string) SqlPoolVulnerabilityAssessmentBaselineInterface {
	return newSqlPoolVulnerabilityAssessmentBaselines(c, namespace)
}

func (c *SynapseV1alpha1Client) SqlPoolWorkloadClassifiers(namespace string) SqlPoolWorkloadClassifierInterface {
	return newSqlPoolWorkloadClassifiers(c, namespace)
}

func (c *SynapseV1alpha1Client) SqlPoolWorkloadGroups(namespace string) SqlPoolWorkloadGroupInterface {
	return newSqlPoolWorkloadGroups(c, namespace)
}

func (c *SynapseV1alpha1Client) Workspaces(namespace string) WorkspaceInterface {
	return newWorkspaces(c, namespace)
}

func (c *SynapseV1alpha1Client) WorkspaceAadAdmins(namespace string) WorkspaceAadAdminInterface {
	return newWorkspaceAadAdmins(c, namespace)
}

func (c *SynapseV1alpha1Client) WorkspaceExtendedAuditingPolicies(namespace string) WorkspaceExtendedAuditingPolicyInterface {
	return newWorkspaceExtendedAuditingPolicies(c, namespace)
}

func (c *SynapseV1alpha1Client) WorkspaceKeys(namespace string) WorkspaceKeyInterface {
	return newWorkspaceKeys(c, namespace)
}

func (c *SynapseV1alpha1Client) WorkspaceSQLAadAdmins(namespace string) WorkspaceSQLAadAdminInterface {
	return newWorkspaceSQLAadAdmins(c, namespace)
}

func (c *SynapseV1alpha1Client) WorkspaceSecurityAlertPolicies(namespace string) WorkspaceSecurityAlertPolicyInterface {
	return newWorkspaceSecurityAlertPolicies(c, namespace)
}

func (c *SynapseV1alpha1Client) WorkspaceVulnerabilityAssessments(namespace string) WorkspaceVulnerabilityAssessmentInterface {
	return newWorkspaceVulnerabilityAssessments(c, namespace)
}

// NewForConfig creates a new SynapseV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*SynapseV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &SynapseV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new SynapseV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *SynapseV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SynapseV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *SynapseV1alpha1Client {
	return &SynapseV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *SynapseV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
