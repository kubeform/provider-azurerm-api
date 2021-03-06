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
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/clientset/versioned/typed/stream/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeStreamV1alpha1 struct {
	*testing.Fake
}

func (c *FakeStreamV1alpha1) AnalyticsClusters(namespace string) v1alpha1.AnalyticsClusterInterface {
	return &FakeAnalyticsClusters{c, namespace}
}

func (c *FakeStreamV1alpha1) AnalyticsFunctionJavascriptUdves(namespace string) v1alpha1.AnalyticsFunctionJavascriptUdfInterface {
	return &FakeAnalyticsFunctionJavascriptUdves{c, namespace}
}

func (c *FakeStreamV1alpha1) AnalyticsJobs(namespace string) v1alpha1.AnalyticsJobInterface {
	return &FakeAnalyticsJobs{c, namespace}
}

func (c *FakeStreamV1alpha1) AnalyticsManagedPrivateEndpoints(namespace string) v1alpha1.AnalyticsManagedPrivateEndpointInterface {
	return &FakeAnalyticsManagedPrivateEndpoints{c, namespace}
}

func (c *FakeStreamV1alpha1) AnalyticsOutputBlobs(namespace string) v1alpha1.AnalyticsOutputBlobInterface {
	return &FakeAnalyticsOutputBlobs{c, namespace}
}

func (c *FakeStreamV1alpha1) AnalyticsOutputEventhubs(namespace string) v1alpha1.AnalyticsOutputEventhubInterface {
	return &FakeAnalyticsOutputEventhubs{c, namespace}
}

func (c *FakeStreamV1alpha1) AnalyticsOutputFunctions(namespace string) v1alpha1.AnalyticsOutputFunctionInterface {
	return &FakeAnalyticsOutputFunctions{c, namespace}
}

func (c *FakeStreamV1alpha1) AnalyticsOutputMssqls(namespace string) v1alpha1.AnalyticsOutputMssqlInterface {
	return &FakeAnalyticsOutputMssqls{c, namespace}
}

func (c *FakeStreamV1alpha1) AnalyticsOutputServicebusQueues(namespace string) v1alpha1.AnalyticsOutputServicebusQueueInterface {
	return &FakeAnalyticsOutputServicebusQueues{c, namespace}
}

func (c *FakeStreamV1alpha1) AnalyticsOutputServicebusTopics(namespace string) v1alpha1.AnalyticsOutputServicebusTopicInterface {
	return &FakeAnalyticsOutputServicebusTopics{c, namespace}
}

func (c *FakeStreamV1alpha1) AnalyticsOutputSynapses(namespace string) v1alpha1.AnalyticsOutputSynapseInterface {
	return &FakeAnalyticsOutputSynapses{c, namespace}
}

func (c *FakeStreamV1alpha1) AnalyticsOutputTables(namespace string) v1alpha1.AnalyticsOutputTableInterface {
	return &FakeAnalyticsOutputTables{c, namespace}
}

func (c *FakeStreamV1alpha1) AnalyticsReferenceInputBlobs(namespace string) v1alpha1.AnalyticsReferenceInputBlobInterface {
	return &FakeAnalyticsReferenceInputBlobs{c, namespace}
}

func (c *FakeStreamV1alpha1) AnalyticsReferenceInputMssqls(namespace string) v1alpha1.AnalyticsReferenceInputMssqlInterface {
	return &FakeAnalyticsReferenceInputMssqls{c, namespace}
}

func (c *FakeStreamV1alpha1) AnalyticsStreamInputBlobs(namespace string) v1alpha1.AnalyticsStreamInputBlobInterface {
	return &FakeAnalyticsStreamInputBlobs{c, namespace}
}

func (c *FakeStreamV1alpha1) AnalyticsStreamInputEventhubs(namespace string) v1alpha1.AnalyticsStreamInputEventhubInterface {
	return &FakeAnalyticsStreamInputEventhubs{c, namespace}
}

func (c *FakeStreamV1alpha1) AnalyticsStreamInputIothubs(namespace string) v1alpha1.AnalyticsStreamInputIothubInterface {
	return &FakeAnalyticsStreamInputIothubs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeStreamV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
