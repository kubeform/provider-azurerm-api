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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Endpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointSpec   `json:"spec,omitempty"`
	Status            EndpointStatus `json:"status,omitempty"`
}

type EndpointSpecDeliveryRuleCacheExpirationAction struct {
	Behavior *string `json:"behavior" tf:"behavior"`
	// +optional
	Duration *string `json:"duration,omitempty" tf:"duration"`
}

type EndpointSpecDeliveryRuleCacheKeyQueryStringAction struct {
	Behavior *string `json:"behavior" tf:"behavior"`
	// +optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters"`
}

type EndpointSpecDeliveryRuleCookiesCondition struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	MatchValues []string `json:"matchValues,omitempty" tf:"match_values"`
	// +optional
	NegateCondition *bool   `json:"negateCondition,omitempty" tf:"negate_condition"`
	Operator        *string `json:"operator" tf:"operator"`
	Selector        *string `json:"selector" tf:"selector"`
	// +optional
	Transforms []string `json:"transforms,omitempty" tf:"transforms"`
}

type EndpointSpecDeliveryRuleDeviceCondition struct {
	// +kubebuilder:validation:MinItems=1
	MatchValues []string `json:"matchValues" tf:"match_values"`
	// +optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition"`
	// +optional
	Operator *string `json:"operator,omitempty" tf:"operator"`
}

type EndpointSpecDeliveryRuleHttpVersionCondition struct {
	// +kubebuilder:validation:MinItems=1
	MatchValues []string `json:"matchValues" tf:"match_values"`
	// +optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition"`
	// +optional
	Operator *string `json:"operator,omitempty" tf:"operator"`
}

type EndpointSpecDeliveryRuleModifyRequestHeaderAction struct {
	Action *string `json:"action" tf:"action"`
	Name   *string `json:"name" tf:"name"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type EndpointSpecDeliveryRuleModifyResponseHeaderAction struct {
	Action *string `json:"action" tf:"action"`
	Name   *string `json:"name" tf:"name"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type EndpointSpecDeliveryRulePostArgCondition struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	MatchValues []string `json:"matchValues,omitempty" tf:"match_values"`
	// +optional
	NegateCondition *bool   `json:"negateCondition,omitempty" tf:"negate_condition"`
	Operator        *string `json:"operator" tf:"operator"`
	Selector        *string `json:"selector" tf:"selector"`
	// +optional
	Transforms []string `json:"transforms,omitempty" tf:"transforms"`
}

type EndpointSpecDeliveryRuleQueryStringCondition struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	MatchValues []string `json:"matchValues,omitempty" tf:"match_values"`
	// +optional
	NegateCondition *bool   `json:"negateCondition,omitempty" tf:"negate_condition"`
	Operator        *string `json:"operator" tf:"operator"`
	// +optional
	Transforms []string `json:"transforms,omitempty" tf:"transforms"`
}

type EndpointSpecDeliveryRuleRemoteAddressCondition struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	MatchValues []string `json:"matchValues,omitempty" tf:"match_values"`
	// +optional
	NegateCondition *bool   `json:"negateCondition,omitempty" tf:"negate_condition"`
	Operator        *string `json:"operator" tf:"operator"`
}

type EndpointSpecDeliveryRuleRequestBodyCondition struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	MatchValues []string `json:"matchValues,omitempty" tf:"match_values"`
	// +optional
	NegateCondition *bool   `json:"negateCondition,omitempty" tf:"negate_condition"`
	Operator        *string `json:"operator" tf:"operator"`
	// +optional
	Transforms []string `json:"transforms,omitempty" tf:"transforms"`
}

type EndpointSpecDeliveryRuleRequestHeaderCondition struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	MatchValues []string `json:"matchValues,omitempty" tf:"match_values"`
	// +optional
	NegateCondition *bool   `json:"negateCondition,omitempty" tf:"negate_condition"`
	Operator        *string `json:"operator" tf:"operator"`
	Selector        *string `json:"selector" tf:"selector"`
	// +optional
	Transforms []string `json:"transforms,omitempty" tf:"transforms"`
}

type EndpointSpecDeliveryRuleRequestMethodCondition struct {
	// +kubebuilder:validation:MinItems=1
	MatchValues []string `json:"matchValues" tf:"match_values"`
	// +optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition"`
	// +optional
	Operator *string `json:"operator,omitempty" tf:"operator"`
}

type EndpointSpecDeliveryRuleRequestSchemeCondition struct {
	// +kubebuilder:validation:MinItems=1
	MatchValues []string `json:"matchValues" tf:"match_values"`
	// +optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition"`
	// +optional
	Operator *string `json:"operator,omitempty" tf:"operator"`
}

type EndpointSpecDeliveryRuleRequestURICondition struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	MatchValues []string `json:"matchValues,omitempty" tf:"match_values"`
	// +optional
	NegateCondition *bool   `json:"negateCondition,omitempty" tf:"negate_condition"`
	Operator        *string `json:"operator" tf:"operator"`
	// +optional
	Transforms []string `json:"transforms,omitempty" tf:"transforms"`
}

type EndpointSpecDeliveryRuleUrlFileExtensionCondition struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	MatchValues []string `json:"matchValues,omitempty" tf:"match_values"`
	// +optional
	NegateCondition *bool   `json:"negateCondition,omitempty" tf:"negate_condition"`
	Operator        *string `json:"operator" tf:"operator"`
	// +optional
	Transforms []string `json:"transforms,omitempty" tf:"transforms"`
}

type EndpointSpecDeliveryRuleUrlFileNameCondition struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	MatchValues []string `json:"matchValues,omitempty" tf:"match_values"`
	// +optional
	NegateCondition *bool   `json:"negateCondition,omitempty" tf:"negate_condition"`
	Operator        *string `json:"operator" tf:"operator"`
	// +optional
	Transforms []string `json:"transforms,omitempty" tf:"transforms"`
}

type EndpointSpecDeliveryRuleUrlPathCondition struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	MatchValues []string `json:"matchValues,omitempty" tf:"match_values"`
	// +optional
	NegateCondition *bool   `json:"negateCondition,omitempty" tf:"negate_condition"`
	Operator        *string `json:"operator" tf:"operator"`
	// +optional
	Transforms []string `json:"transforms,omitempty" tf:"transforms"`
}

type EndpointSpecDeliveryRuleUrlRedirectAction struct {
	// +optional
	Fragment *string `json:"fragment,omitempty" tf:"fragment"`
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
	// +optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
	// +optional
	QueryString  *string `json:"queryString,omitempty" tf:"query_string"`
	RedirectType *string `json:"redirectType" tf:"redirect_type"`
}

type EndpointSpecDeliveryRuleUrlRewriteAction struct {
	Destination *string `json:"destination" tf:"destination"`
	// +optional
	PreserveUnmatchedPath *bool   `json:"preserveUnmatchedPath,omitempty" tf:"preserve_unmatched_path"`
	SourcePattern         *string `json:"sourcePattern" tf:"source_pattern"`
}

type EndpointSpecDeliveryRule struct {
	// +optional
	CacheExpirationAction *EndpointSpecDeliveryRuleCacheExpirationAction `json:"cacheExpirationAction,omitempty" tf:"cache_expiration_action"`
	// +optional
	CacheKeyQueryStringAction *EndpointSpecDeliveryRuleCacheKeyQueryStringAction `json:"cacheKeyQueryStringAction,omitempty" tf:"cache_key_query_string_action"`
	// +optional
	CookiesCondition []EndpointSpecDeliveryRuleCookiesCondition `json:"cookiesCondition,omitempty" tf:"cookies_condition"`
	// +optional
	DeviceCondition *EndpointSpecDeliveryRuleDeviceCondition `json:"deviceCondition,omitempty" tf:"device_condition"`
	// +optional
	HttpVersionCondition []EndpointSpecDeliveryRuleHttpVersionCondition `json:"httpVersionCondition,omitempty" tf:"http_version_condition"`
	// +optional
	ModifyRequestHeaderAction []EndpointSpecDeliveryRuleModifyRequestHeaderAction `json:"modifyRequestHeaderAction,omitempty" tf:"modify_request_header_action"`
	// +optional
	ModifyResponseHeaderAction []EndpointSpecDeliveryRuleModifyResponseHeaderAction `json:"modifyResponseHeaderAction,omitempty" tf:"modify_response_header_action"`
	Name                       *string                                              `json:"name" tf:"name"`
	Order                      *int64                                               `json:"order" tf:"order"`
	// +optional
	PostArgCondition []EndpointSpecDeliveryRulePostArgCondition `json:"postArgCondition,omitempty" tf:"post_arg_condition"`
	// +optional
	QueryStringCondition []EndpointSpecDeliveryRuleQueryStringCondition `json:"queryStringCondition,omitempty" tf:"query_string_condition"`
	// +optional
	RemoteAddressCondition []EndpointSpecDeliveryRuleRemoteAddressCondition `json:"remoteAddressCondition,omitempty" tf:"remote_address_condition"`
	// +optional
	RequestBodyCondition []EndpointSpecDeliveryRuleRequestBodyCondition `json:"requestBodyCondition,omitempty" tf:"request_body_condition"`
	// +optional
	RequestHeaderCondition []EndpointSpecDeliveryRuleRequestHeaderCondition `json:"requestHeaderCondition,omitempty" tf:"request_header_condition"`
	// +optional
	RequestMethodCondition *EndpointSpecDeliveryRuleRequestMethodCondition `json:"requestMethodCondition,omitempty" tf:"request_method_condition"`
	// +optional
	RequestSchemeCondition *EndpointSpecDeliveryRuleRequestSchemeCondition `json:"requestSchemeCondition,omitempty" tf:"request_scheme_condition"`
	// +optional
	RequestURICondition []EndpointSpecDeliveryRuleRequestURICondition `json:"requestURICondition,omitempty" tf:"request_uri_condition"`
	// +optional
	UrlFileExtensionCondition []EndpointSpecDeliveryRuleUrlFileExtensionCondition `json:"urlFileExtensionCondition,omitempty" tf:"url_file_extension_condition"`
	// +optional
	UrlFileNameCondition []EndpointSpecDeliveryRuleUrlFileNameCondition `json:"urlFileNameCondition,omitempty" tf:"url_file_name_condition"`
	// +optional
	UrlPathCondition []EndpointSpecDeliveryRuleUrlPathCondition `json:"urlPathCondition,omitempty" tf:"url_path_condition"`
	// +optional
	UrlRedirectAction *EndpointSpecDeliveryRuleUrlRedirectAction `json:"urlRedirectAction,omitempty" tf:"url_redirect_action"`
	// +optional
	UrlRewriteAction *EndpointSpecDeliveryRuleUrlRewriteAction `json:"urlRewriteAction,omitempty" tf:"url_rewrite_action"`
}

type EndpointSpecGeoFilter struct {
	Action       *string  `json:"action" tf:"action"`
	CountryCodes []string `json:"countryCodes" tf:"country_codes"`
	RelativePath *string  `json:"relativePath" tf:"relative_path"`
}

type EndpointSpecGlobalDeliveryRuleCacheExpirationAction struct {
	Behavior *string `json:"behavior" tf:"behavior"`
	// +optional
	Duration *string `json:"duration,omitempty" tf:"duration"`
}

type EndpointSpecGlobalDeliveryRuleCacheKeyQueryStringAction struct {
	Behavior *string `json:"behavior" tf:"behavior"`
	// +optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters"`
}

type EndpointSpecGlobalDeliveryRuleModifyRequestHeaderAction struct {
	Action *string `json:"action" tf:"action"`
	Name   *string `json:"name" tf:"name"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type EndpointSpecGlobalDeliveryRuleModifyResponseHeaderAction struct {
	Action *string `json:"action" tf:"action"`
	Name   *string `json:"name" tf:"name"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type EndpointSpecGlobalDeliveryRuleUrlRedirectAction struct {
	// +optional
	Fragment *string `json:"fragment,omitempty" tf:"fragment"`
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
	// +optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
	// +optional
	QueryString  *string `json:"queryString,omitempty" tf:"query_string"`
	RedirectType *string `json:"redirectType" tf:"redirect_type"`
}

type EndpointSpecGlobalDeliveryRuleUrlRewriteAction struct {
	Destination *string `json:"destination" tf:"destination"`
	// +optional
	PreserveUnmatchedPath *bool   `json:"preserveUnmatchedPath,omitempty" tf:"preserve_unmatched_path"`
	SourcePattern         *string `json:"sourcePattern" tf:"source_pattern"`
}

type EndpointSpecGlobalDeliveryRule struct {
	// +optional
	CacheExpirationAction *EndpointSpecGlobalDeliveryRuleCacheExpirationAction `json:"cacheExpirationAction,omitempty" tf:"cache_expiration_action"`
	// +optional
	CacheKeyQueryStringAction *EndpointSpecGlobalDeliveryRuleCacheKeyQueryStringAction `json:"cacheKeyQueryStringAction,omitempty" tf:"cache_key_query_string_action"`
	// +optional
	ModifyRequestHeaderAction []EndpointSpecGlobalDeliveryRuleModifyRequestHeaderAction `json:"modifyRequestHeaderAction,omitempty" tf:"modify_request_header_action"`
	// +optional
	ModifyResponseHeaderAction []EndpointSpecGlobalDeliveryRuleModifyResponseHeaderAction `json:"modifyResponseHeaderAction,omitempty" tf:"modify_response_header_action"`
	// +optional
	UrlRedirectAction *EndpointSpecGlobalDeliveryRuleUrlRedirectAction `json:"urlRedirectAction,omitempty" tf:"url_redirect_action"`
	// +optional
	UrlRewriteAction *EndpointSpecGlobalDeliveryRuleUrlRewriteAction `json:"urlRewriteAction,omitempty" tf:"url_rewrite_action"`
}

type EndpointSpecOrigin struct {
	HostName *string `json:"hostName" tf:"host_name"`
	// +optional
	HttpPort *int64 `json:"httpPort,omitempty" tf:"http_port"`
	// +optional
	HttpsPort *int64  `json:"httpsPort,omitempty" tf:"https_port"`
	Name      *string `json:"name" tf:"name"`
}

type EndpointSpec struct {
	State *EndpointSpecResource `json:"state,omitempty" tf:"-"`

	Resource EndpointSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type EndpointSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ContentTypesToCompress []string `json:"contentTypesToCompress,omitempty" tf:"content_types_to_compress"`
	// +optional
	DeliveryRule []EndpointSpecDeliveryRule `json:"deliveryRule,omitempty" tf:"delivery_rule"`
	// +optional
	GeoFilter []EndpointSpecGeoFilter `json:"geoFilter,omitempty" tf:"geo_filter"`
	// +optional
	GlobalDeliveryRule *EndpointSpecGlobalDeliveryRule `json:"globalDeliveryRule,omitempty" tf:"global_delivery_rule"`
	// +optional
	HostName *string `json:"hostName,omitempty" tf:"host_name"`
	// +optional
	IsCompressionEnabled *bool `json:"isCompressionEnabled,omitempty" tf:"is_compression_enabled"`
	// +optional
	IsHTTPAllowed *bool `json:"isHTTPAllowed,omitempty" tf:"is_http_allowed"`
	// +optional
	IsHTTPSAllowed *bool   `json:"isHTTPSAllowed,omitempty" tf:"is_https_allowed"`
	Location       *string `json:"location" tf:"location"`
	Name           *string `json:"name" tf:"name"`
	// +optional
	OptimizationType *string              `json:"optimizationType,omitempty" tf:"optimization_type"`
	Origin           []EndpointSpecOrigin `json:"origin" tf:"origin"`
	// +optional
	OriginHostHeader *string `json:"originHostHeader,omitempty" tf:"origin_host_header"`
	// +optional
	OriginPath *string `json:"originPath,omitempty" tf:"origin_path"`
	// +optional
	ProbePath   *string `json:"probePath,omitempty" tf:"probe_path"`
	ProfileName *string `json:"profileName" tf:"profile_name"`
	// +optional
	QuerystringCachingBehaviour *string `json:"querystringCachingBehaviour,omitempty" tf:"querystring_caching_behaviour"`
	ResourceGroupName           *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type EndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EndpointList is a list of Endpoints
type EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Endpoint CRD objects
	Items []Endpoint `json:"items,omitempty"`
}
