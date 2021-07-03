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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/stream/v1alpha1"
	"kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type StreamV1alpha1Interface interface {
	RESTClient() rest.Interface
	AnalyticsFunctionJavascriptUdvesGetter
	AnalyticsJobsGetter
	AnalyticsOutputBlobsGetter
	AnalyticsOutputEventhubsGetter
	AnalyticsOutputMssqlsGetter
	AnalyticsOutputServicebusQueuesGetter
	AnalyticsOutputServicebusTopicsGetter
	AnalyticsReferenceInputBlobsGetter
	AnalyticsStreamInputBlobsGetter
	AnalyticsStreamInputEventhubsGetter
	AnalyticsStreamInputIothubsGetter
}

// StreamV1alpha1Client is used to interact with features provided by the stream.azurerm.kubeform.com group.
type StreamV1alpha1Client struct {
	restClient rest.Interface
}

func (c *StreamV1alpha1Client) AnalyticsFunctionJavascriptUdves(namespace string) AnalyticsFunctionJavascriptUdfInterface {
	return newAnalyticsFunctionJavascriptUdves(c, namespace)
}

func (c *StreamV1alpha1Client) AnalyticsJobs(namespace string) AnalyticsJobInterface {
	return newAnalyticsJobs(c, namespace)
}

func (c *StreamV1alpha1Client) AnalyticsOutputBlobs(namespace string) AnalyticsOutputBlobInterface {
	return newAnalyticsOutputBlobs(c, namespace)
}

func (c *StreamV1alpha1Client) AnalyticsOutputEventhubs(namespace string) AnalyticsOutputEventhubInterface {
	return newAnalyticsOutputEventhubs(c, namespace)
}

func (c *StreamV1alpha1Client) AnalyticsOutputMssqls(namespace string) AnalyticsOutputMssqlInterface {
	return newAnalyticsOutputMssqls(c, namespace)
}

func (c *StreamV1alpha1Client) AnalyticsOutputServicebusQueues(namespace string) AnalyticsOutputServicebusQueueInterface {
	return newAnalyticsOutputServicebusQueues(c, namespace)
}

func (c *StreamV1alpha1Client) AnalyticsOutputServicebusTopics(namespace string) AnalyticsOutputServicebusTopicInterface {
	return newAnalyticsOutputServicebusTopics(c, namespace)
}

func (c *StreamV1alpha1Client) AnalyticsReferenceInputBlobs(namespace string) AnalyticsReferenceInputBlobInterface {
	return newAnalyticsReferenceInputBlobs(c, namespace)
}

func (c *StreamV1alpha1Client) AnalyticsStreamInputBlobs(namespace string) AnalyticsStreamInputBlobInterface {
	return newAnalyticsStreamInputBlobs(c, namespace)
}

func (c *StreamV1alpha1Client) AnalyticsStreamInputEventhubs(namespace string) AnalyticsStreamInputEventhubInterface {
	return newAnalyticsStreamInputEventhubs(c, namespace)
}

func (c *StreamV1alpha1Client) AnalyticsStreamInputIothubs(namespace string) AnalyticsStreamInputIothubInterface {
	return newAnalyticsStreamInputIothubs(c, namespace)
}

// NewForConfig creates a new StreamV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*StreamV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &StreamV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new StreamV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *StreamV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new StreamV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *StreamV1alpha1Client {
	return &StreamV1alpha1Client{c}
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
func (c *StreamV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
