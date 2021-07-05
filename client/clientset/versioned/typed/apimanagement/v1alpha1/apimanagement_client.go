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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/apimanagement/v1alpha1"
	"kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type ApimanagementV1alpha1Interface interface {
	RESTClient() rest.Interface
	ApisGetter
	ApiDiagnosticsGetter
	ApiManagementsGetter
	ApiOperationsGetter
	ApiOperationPoliciesGetter
	ApiOperationTagsGetter
	ApiPoliciesGetter
	ApiSchemasGetter
	ApiVersionSetsGetter
	AuthorizationServersGetter
	BackendsGetter
	CertificatesGetter
	CustomDomainsGetter
	DiagnosticsGetter
	EmailTemplatesGetter
	GroupsGetter
	GroupUsersGetter
	IdentityProviderAadsGetter
	IdentityProviderAadb2csGetter
	IdentityProviderFacebooksGetter
	IdentityProviderGooglesGetter
	IdentityProviderMicrosoftsGetter
	IdentityProviderTwittersGetter
	LoggersGetter
	NamedValuesGetter
	OpenidConnectProvidersGetter
	PoliciesGetter
	ProductsGetter
	ProductAPIsGetter
	ProductGroupsGetter
	ProductPoliciesGetter
	PropertiesGetter
	RedisCachesGetter
	SubscriptionsGetter
	UsersGetter
}

// ApimanagementV1alpha1Client is used to interact with features provided by the apimanagement.azurerm.kubeform.com group.
type ApimanagementV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ApimanagementV1alpha1Client) Apis(namespace string) ApiInterface {
	return newApis(c, namespace)
}

func (c *ApimanagementV1alpha1Client) ApiDiagnostics(namespace string) ApiDiagnosticInterface {
	return newApiDiagnostics(c, namespace)
}

func (c *ApimanagementV1alpha1Client) ApiManagements(namespace string) ApiManagementInterface {
	return newApiManagements(c, namespace)
}

func (c *ApimanagementV1alpha1Client) ApiOperations(namespace string) ApiOperationInterface {
	return newApiOperations(c, namespace)
}

func (c *ApimanagementV1alpha1Client) ApiOperationPolicies(namespace string) ApiOperationPolicyInterface {
	return newApiOperationPolicies(c, namespace)
}

func (c *ApimanagementV1alpha1Client) ApiOperationTags(namespace string) ApiOperationTagInterface {
	return newApiOperationTags(c, namespace)
}

func (c *ApimanagementV1alpha1Client) ApiPolicies(namespace string) ApiPolicyInterface {
	return newApiPolicies(c, namespace)
}

func (c *ApimanagementV1alpha1Client) ApiSchemas(namespace string) ApiSchemaInterface {
	return newApiSchemas(c, namespace)
}

func (c *ApimanagementV1alpha1Client) ApiVersionSets(namespace string) ApiVersionSetInterface {
	return newApiVersionSets(c, namespace)
}

func (c *ApimanagementV1alpha1Client) AuthorizationServers(namespace string) AuthorizationServerInterface {
	return newAuthorizationServers(c, namespace)
}

func (c *ApimanagementV1alpha1Client) Backends(namespace string) BackendInterface {
	return newBackends(c, namespace)
}

func (c *ApimanagementV1alpha1Client) Certificates(namespace string) CertificateInterface {
	return newCertificates(c, namespace)
}

func (c *ApimanagementV1alpha1Client) CustomDomains(namespace string) CustomDomainInterface {
	return newCustomDomains(c, namespace)
}

func (c *ApimanagementV1alpha1Client) Diagnostics(namespace string) DiagnosticInterface {
	return newDiagnostics(c, namespace)
}

func (c *ApimanagementV1alpha1Client) EmailTemplates(namespace string) EmailTemplateInterface {
	return newEmailTemplates(c, namespace)
}

func (c *ApimanagementV1alpha1Client) Groups(namespace string) GroupInterface {
	return newGroups(c, namespace)
}

func (c *ApimanagementV1alpha1Client) GroupUsers(namespace string) GroupUserInterface {
	return newGroupUsers(c, namespace)
}

func (c *ApimanagementV1alpha1Client) IdentityProviderAads(namespace string) IdentityProviderAadInterface {
	return newIdentityProviderAads(c, namespace)
}

func (c *ApimanagementV1alpha1Client) IdentityProviderAadb2cs(namespace string) IdentityProviderAadb2cInterface {
	return newIdentityProviderAadb2cs(c, namespace)
}

func (c *ApimanagementV1alpha1Client) IdentityProviderFacebooks(namespace string) IdentityProviderFacebookInterface {
	return newIdentityProviderFacebooks(c, namespace)
}

func (c *ApimanagementV1alpha1Client) IdentityProviderGoogles(namespace string) IdentityProviderGoogleInterface {
	return newIdentityProviderGoogles(c, namespace)
}

func (c *ApimanagementV1alpha1Client) IdentityProviderMicrosofts(namespace string) IdentityProviderMicrosoftInterface {
	return newIdentityProviderMicrosofts(c, namespace)
}

func (c *ApimanagementV1alpha1Client) IdentityProviderTwitters(namespace string) IdentityProviderTwitterInterface {
	return newIdentityProviderTwitters(c, namespace)
}

func (c *ApimanagementV1alpha1Client) Loggers(namespace string) LoggerInterface {
	return newLoggers(c, namespace)
}

func (c *ApimanagementV1alpha1Client) NamedValues(namespace string) NamedValueInterface {
	return newNamedValues(c, namespace)
}

func (c *ApimanagementV1alpha1Client) OpenidConnectProviders(namespace string) OpenidConnectProviderInterface {
	return newOpenidConnectProviders(c, namespace)
}

func (c *ApimanagementV1alpha1Client) Policies(namespace string) PolicyInterface {
	return newPolicies(c, namespace)
}

func (c *ApimanagementV1alpha1Client) Products(namespace string) ProductInterface {
	return newProducts(c, namespace)
}

func (c *ApimanagementV1alpha1Client) ProductAPIs(namespace string) ProductAPIInterface {
	return newProductAPIs(c, namespace)
}

func (c *ApimanagementV1alpha1Client) ProductGroups(namespace string) ProductGroupInterface {
	return newProductGroups(c, namespace)
}

func (c *ApimanagementV1alpha1Client) ProductPolicies(namespace string) ProductPolicyInterface {
	return newProductPolicies(c, namespace)
}

func (c *ApimanagementV1alpha1Client) Properties(namespace string) PropertyInterface {
	return newProperties(c, namespace)
}

func (c *ApimanagementV1alpha1Client) RedisCaches(namespace string) RedisCacheInterface {
	return newRedisCaches(c, namespace)
}

func (c *ApimanagementV1alpha1Client) Subscriptions(namespace string) SubscriptionInterface {
	return newSubscriptions(c, namespace)
}

func (c *ApimanagementV1alpha1Client) Users(namespace string) UserInterface {
	return newUsers(c, namespace)
}

// NewForConfig creates a new ApimanagementV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*ApimanagementV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ApimanagementV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ApimanagementV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ApimanagementV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ApimanagementV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ApimanagementV1alpha1Client {
	return &ApimanagementV1alpha1Client{c}
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
func (c *ApimanagementV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}