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
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/clientset/versioned/typed/keyvault/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKeyvaultV1alpha1 struct {
	*testing.Fake
}

func (c *FakeKeyvaultV1alpha1) AccessPolicies(namespace string) v1alpha1.AccessPolicyInterface {
	return &FakeAccessPolicies{c, namespace}
}

func (c *FakeKeyvaultV1alpha1) Certificates(namespace string) v1alpha1.CertificateInterface {
	return &FakeCertificates{c, namespace}
}

func (c *FakeKeyvaultV1alpha1) CertificateIssuers(namespace string) v1alpha1.CertificateIssuerInterface {
	return &FakeCertificateIssuers{c, namespace}
}

func (c *FakeKeyvaultV1alpha1) Keys(namespace string) v1alpha1.KeyInterface {
	return &FakeKeys{c, namespace}
}

func (c *FakeKeyvaultV1alpha1) KeyVaults(namespace string) v1alpha1.KeyVaultInterface {
	return &FakeKeyVaults{c, namespace}
}

func (c *FakeKeyvaultV1alpha1) ManagedHardwareSecurityModules(namespace string) v1alpha1.ManagedHardwareSecurityModuleInterface {
	return &FakeManagedHardwareSecurityModules{c, namespace}
}

func (c *FakeKeyvaultV1alpha1) ManagedStorageAccounts(namespace string) v1alpha1.ManagedStorageAccountInterface {
	return &FakeManagedStorageAccounts{c, namespace}
}

func (c *FakeKeyvaultV1alpha1) ManagedStorageAccountSasTokenDefinitions(namespace string) v1alpha1.ManagedStorageAccountSasTokenDefinitionInterface {
	return &FakeManagedStorageAccountSasTokenDefinitions{c, namespace}
}

func (c *FakeKeyvaultV1alpha1) Secrets(namespace string) v1alpha1.SecretInterface {
	return &FakeSecrets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKeyvaultV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
