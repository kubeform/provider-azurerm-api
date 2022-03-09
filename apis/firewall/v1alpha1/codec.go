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
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(FirewallSpecManagementIPConfiguration{}).Type1()): FirewallSpecManagementIPConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FirewallSpecVirtualHub{}).Type1()):                FirewallSpecVirtualHubCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecDns{}).Type1()):                         PolicySpecDnsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecIdentity{}).Type1()):                    PolicySpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecInsights{}).Type1()):                    PolicySpecInsightsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecIntrusionDetection{}).Type1()):          PolicySpecIntrusionDetectionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecThreatIntelligenceAllowlist{}).Type1()): PolicySpecThreatIntelligenceAllowlistCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecTlsCertificate{}).Type1()):              PolicySpecTlsCertificateCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(FirewallSpecManagementIPConfiguration{}).Type1()): FirewallSpecManagementIPConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FirewallSpecVirtualHub{}).Type1()):                FirewallSpecVirtualHubCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecDns{}).Type1()):                         PolicySpecDnsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecIdentity{}).Type1()):                    PolicySpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecInsights{}).Type1()):                    PolicySpecInsightsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecIntrusionDetection{}).Type1()):          PolicySpecIntrusionDetectionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecThreatIntelligenceAllowlist{}).Type1()): PolicySpecThreatIntelligenceAllowlistCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecTlsCertificate{}).Type1()):              PolicySpecTlsCertificateCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type FirewallSpecManagementIPConfigurationCodec struct {
}

func (FirewallSpecManagementIPConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*FirewallSpecManagementIPConfiguration)(ptr) == nil
}

func (FirewallSpecManagementIPConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*FirewallSpecManagementIPConfiguration)(ptr)
	var objs []FirewallSpecManagementIPConfiguration
	if obj != nil {
		objs = []FirewallSpecManagementIPConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FirewallSpecManagementIPConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (FirewallSpecManagementIPConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*FirewallSpecManagementIPConfiguration)(ptr) = FirewallSpecManagementIPConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []FirewallSpecManagementIPConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FirewallSpecManagementIPConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*FirewallSpecManagementIPConfiguration)(ptr) = objs[0]
			} else {
				*(*FirewallSpecManagementIPConfiguration)(ptr) = FirewallSpecManagementIPConfiguration{}
			}
		} else {
			*(*FirewallSpecManagementIPConfiguration)(ptr) = FirewallSpecManagementIPConfiguration{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj FirewallSpecManagementIPConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FirewallSpecManagementIPConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*FirewallSpecManagementIPConfiguration)(ptr) = obj
		} else {
			*(*FirewallSpecManagementIPConfiguration)(ptr) = FirewallSpecManagementIPConfiguration{}
		}
	default:
		iter.ReportError("decode FirewallSpecManagementIPConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type FirewallSpecVirtualHubCodec struct {
}

func (FirewallSpecVirtualHubCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*FirewallSpecVirtualHub)(ptr) == nil
}

func (FirewallSpecVirtualHubCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*FirewallSpecVirtualHub)(ptr)
	var objs []FirewallSpecVirtualHub
	if obj != nil {
		objs = []FirewallSpecVirtualHub{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FirewallSpecVirtualHub{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (FirewallSpecVirtualHubCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*FirewallSpecVirtualHub)(ptr) = FirewallSpecVirtualHub{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []FirewallSpecVirtualHub

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FirewallSpecVirtualHub{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*FirewallSpecVirtualHub)(ptr) = objs[0]
			} else {
				*(*FirewallSpecVirtualHub)(ptr) = FirewallSpecVirtualHub{}
			}
		} else {
			*(*FirewallSpecVirtualHub)(ptr) = FirewallSpecVirtualHub{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj FirewallSpecVirtualHub

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FirewallSpecVirtualHub{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*FirewallSpecVirtualHub)(ptr) = obj
		} else {
			*(*FirewallSpecVirtualHub)(ptr) = FirewallSpecVirtualHub{}
		}
	default:
		iter.ReportError("decode FirewallSpecVirtualHub", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PolicySpecDnsCodec struct {
}

func (PolicySpecDnsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PolicySpecDns)(ptr) == nil
}

func (PolicySpecDnsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PolicySpecDns)(ptr)
	var objs []PolicySpecDns
	if obj != nil {
		objs = []PolicySpecDns{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecDns{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PolicySpecDnsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PolicySpecDns)(ptr) = PolicySpecDns{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PolicySpecDns

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecDns{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PolicySpecDns)(ptr) = objs[0]
			} else {
				*(*PolicySpecDns)(ptr) = PolicySpecDns{}
			}
		} else {
			*(*PolicySpecDns)(ptr) = PolicySpecDns{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj PolicySpecDns

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecDns{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*PolicySpecDns)(ptr) = obj
		} else {
			*(*PolicySpecDns)(ptr) = PolicySpecDns{}
		}
	default:
		iter.ReportError("decode PolicySpecDns", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PolicySpecIdentityCodec struct {
}

func (PolicySpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PolicySpecIdentity)(ptr) == nil
}

func (PolicySpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PolicySpecIdentity)(ptr)
	var objs []PolicySpecIdentity
	if obj != nil {
		objs = []PolicySpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PolicySpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PolicySpecIdentity)(ptr) = PolicySpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PolicySpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PolicySpecIdentity)(ptr) = objs[0]
			} else {
				*(*PolicySpecIdentity)(ptr) = PolicySpecIdentity{}
			}
		} else {
			*(*PolicySpecIdentity)(ptr) = PolicySpecIdentity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj PolicySpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*PolicySpecIdentity)(ptr) = obj
		} else {
			*(*PolicySpecIdentity)(ptr) = PolicySpecIdentity{}
		}
	default:
		iter.ReportError("decode PolicySpecIdentity", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PolicySpecInsightsCodec struct {
}

func (PolicySpecInsightsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PolicySpecInsights)(ptr) == nil
}

func (PolicySpecInsightsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PolicySpecInsights)(ptr)
	var objs []PolicySpecInsights
	if obj != nil {
		objs = []PolicySpecInsights{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecInsights{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PolicySpecInsightsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PolicySpecInsights)(ptr) = PolicySpecInsights{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PolicySpecInsights

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecInsights{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PolicySpecInsights)(ptr) = objs[0]
			} else {
				*(*PolicySpecInsights)(ptr) = PolicySpecInsights{}
			}
		} else {
			*(*PolicySpecInsights)(ptr) = PolicySpecInsights{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj PolicySpecInsights

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecInsights{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*PolicySpecInsights)(ptr) = obj
		} else {
			*(*PolicySpecInsights)(ptr) = PolicySpecInsights{}
		}
	default:
		iter.ReportError("decode PolicySpecInsights", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PolicySpecIntrusionDetectionCodec struct {
}

func (PolicySpecIntrusionDetectionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PolicySpecIntrusionDetection)(ptr) == nil
}

func (PolicySpecIntrusionDetectionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PolicySpecIntrusionDetection)(ptr)
	var objs []PolicySpecIntrusionDetection
	if obj != nil {
		objs = []PolicySpecIntrusionDetection{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecIntrusionDetection{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PolicySpecIntrusionDetectionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PolicySpecIntrusionDetection)(ptr) = PolicySpecIntrusionDetection{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PolicySpecIntrusionDetection

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecIntrusionDetection{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PolicySpecIntrusionDetection)(ptr) = objs[0]
			} else {
				*(*PolicySpecIntrusionDetection)(ptr) = PolicySpecIntrusionDetection{}
			}
		} else {
			*(*PolicySpecIntrusionDetection)(ptr) = PolicySpecIntrusionDetection{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj PolicySpecIntrusionDetection

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecIntrusionDetection{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*PolicySpecIntrusionDetection)(ptr) = obj
		} else {
			*(*PolicySpecIntrusionDetection)(ptr) = PolicySpecIntrusionDetection{}
		}
	default:
		iter.ReportError("decode PolicySpecIntrusionDetection", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PolicySpecThreatIntelligenceAllowlistCodec struct {
}

func (PolicySpecThreatIntelligenceAllowlistCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PolicySpecThreatIntelligenceAllowlist)(ptr) == nil
}

func (PolicySpecThreatIntelligenceAllowlistCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PolicySpecThreatIntelligenceAllowlist)(ptr)
	var objs []PolicySpecThreatIntelligenceAllowlist
	if obj != nil {
		objs = []PolicySpecThreatIntelligenceAllowlist{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecThreatIntelligenceAllowlist{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PolicySpecThreatIntelligenceAllowlistCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PolicySpecThreatIntelligenceAllowlist)(ptr) = PolicySpecThreatIntelligenceAllowlist{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PolicySpecThreatIntelligenceAllowlist

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecThreatIntelligenceAllowlist{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PolicySpecThreatIntelligenceAllowlist)(ptr) = objs[0]
			} else {
				*(*PolicySpecThreatIntelligenceAllowlist)(ptr) = PolicySpecThreatIntelligenceAllowlist{}
			}
		} else {
			*(*PolicySpecThreatIntelligenceAllowlist)(ptr) = PolicySpecThreatIntelligenceAllowlist{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj PolicySpecThreatIntelligenceAllowlist

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecThreatIntelligenceAllowlist{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*PolicySpecThreatIntelligenceAllowlist)(ptr) = obj
		} else {
			*(*PolicySpecThreatIntelligenceAllowlist)(ptr) = PolicySpecThreatIntelligenceAllowlist{}
		}
	default:
		iter.ReportError("decode PolicySpecThreatIntelligenceAllowlist", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PolicySpecTlsCertificateCodec struct {
}

func (PolicySpecTlsCertificateCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PolicySpecTlsCertificate)(ptr) == nil
}

func (PolicySpecTlsCertificateCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PolicySpecTlsCertificate)(ptr)
	var objs []PolicySpecTlsCertificate
	if obj != nil {
		objs = []PolicySpecTlsCertificate{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecTlsCertificate{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PolicySpecTlsCertificateCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PolicySpecTlsCertificate)(ptr) = PolicySpecTlsCertificate{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PolicySpecTlsCertificate

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecTlsCertificate{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PolicySpecTlsCertificate)(ptr) = objs[0]
			} else {
				*(*PolicySpecTlsCertificate)(ptr) = PolicySpecTlsCertificate{}
			}
		} else {
			*(*PolicySpecTlsCertificate)(ptr) = PolicySpecTlsCertificate{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj PolicySpecTlsCertificate

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecTlsCertificate{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*PolicySpecTlsCertificate)(ptr) = obj
		} else {
			*(*PolicySpecTlsCertificate)(ptr) = PolicySpecTlsCertificate{}
		}
	default:
		iter.ReportError("decode PolicySpecTlsCertificate", "unexpected JSON type")
	}
}
