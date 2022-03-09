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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/iothub/v1alpha1"
	"kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type IothubV1alpha1Interface interface {
	RESTClient() rest.Interface
	CertificatesGetter
	ConsumerGroupsGetter
	DpsesGetter
	DpsCertificatesGetter
	DpsSharedAccessPoliciesGetter
	EndpointEventhubsGetter
	EndpointServicebusQueuesGetter
	EndpointServicebusTopicsGetter
	EndpointStorageContainersGetter
	EnrichmentsGetter
	FallbackRoutesGetter
	IothubsGetter
	RoutesGetter
	SharedAccessPoliciesGetter
}

// IothubV1alpha1Client is used to interact with features provided by the iothub.azurerm.kubeform.com group.
type IothubV1alpha1Client struct {
	restClient rest.Interface
}

func (c *IothubV1alpha1Client) Certificates(namespace string) CertificateInterface {
	return newCertificates(c, namespace)
}

func (c *IothubV1alpha1Client) ConsumerGroups(namespace string) ConsumerGroupInterface {
	return newConsumerGroups(c, namespace)
}

func (c *IothubV1alpha1Client) Dpses(namespace string) DpsInterface {
	return newDpses(c, namespace)
}

func (c *IothubV1alpha1Client) DpsCertificates(namespace string) DpsCertificateInterface {
	return newDpsCertificates(c, namespace)
}

func (c *IothubV1alpha1Client) DpsSharedAccessPolicies(namespace string) DpsSharedAccessPolicyInterface {
	return newDpsSharedAccessPolicies(c, namespace)
}

func (c *IothubV1alpha1Client) EndpointEventhubs(namespace string) EndpointEventhubInterface {
	return newEndpointEventhubs(c, namespace)
}

func (c *IothubV1alpha1Client) EndpointServicebusQueues(namespace string) EndpointServicebusQueueInterface {
	return newEndpointServicebusQueues(c, namespace)
}

func (c *IothubV1alpha1Client) EndpointServicebusTopics(namespace string) EndpointServicebusTopicInterface {
	return newEndpointServicebusTopics(c, namespace)
}

func (c *IothubV1alpha1Client) EndpointStorageContainers(namespace string) EndpointStorageContainerInterface {
	return newEndpointStorageContainers(c, namespace)
}

func (c *IothubV1alpha1Client) Enrichments(namespace string) EnrichmentInterface {
	return newEnrichments(c, namespace)
}

func (c *IothubV1alpha1Client) FallbackRoutes(namespace string) FallbackRouteInterface {
	return newFallbackRoutes(c, namespace)
}

func (c *IothubV1alpha1Client) Iothubs(namespace string) IothubInterface {
	return newIothubs(c, namespace)
}

func (c *IothubV1alpha1Client) Routes(namespace string) RouteInterface {
	return newRoutes(c, namespace)
}

func (c *IothubV1alpha1Client) SharedAccessPolicies(namespace string) SharedAccessPolicyInterface {
	return newSharedAccessPolicies(c, namespace)
}

// NewForConfig creates a new IothubV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*IothubV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &IothubV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new IothubV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *IothubV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new IothubV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *IothubV1alpha1Client {
	return &IothubV1alpha1Client{c}
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
func (c *IothubV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
