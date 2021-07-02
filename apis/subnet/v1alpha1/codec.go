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
		jsoniter.MustGetKind(reflect2.TypeOf(SubnetSpecDelegationServiceDelegation{}).Type1()):      SubnetSpecDelegationServiceDelegationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceEndpointStoragePolicySpecDefinition{}).Type1()): ServiceEndpointStoragePolicySpecDefinitionCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(SubnetSpecDelegationServiceDelegation{}).Type1()):      SubnetSpecDelegationServiceDelegationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceEndpointStoragePolicySpecDefinition{}).Type1()): ServiceEndpointStoragePolicySpecDefinitionCodec{},
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
type SubnetSpecDelegationServiceDelegationCodec struct {
}

func (SubnetSpecDelegationServiceDelegationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SubnetSpecDelegationServiceDelegation)(ptr) == nil
}

func (SubnetSpecDelegationServiceDelegationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SubnetSpecDelegationServiceDelegation)(ptr)
	var objs []SubnetSpecDelegationServiceDelegation
	if obj != nil {
		objs = []SubnetSpecDelegationServiceDelegation{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SubnetSpecDelegationServiceDelegation{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SubnetSpecDelegationServiceDelegationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SubnetSpecDelegationServiceDelegation)(ptr) = SubnetSpecDelegationServiceDelegation{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SubnetSpecDelegationServiceDelegation

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SubnetSpecDelegationServiceDelegation{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SubnetSpecDelegationServiceDelegation)(ptr) = objs[0]
			} else {
				*(*SubnetSpecDelegationServiceDelegation)(ptr) = SubnetSpecDelegationServiceDelegation{}
			}
		} else {
			*(*SubnetSpecDelegationServiceDelegation)(ptr) = SubnetSpecDelegationServiceDelegation{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj SubnetSpecDelegationServiceDelegation

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SubnetSpecDelegationServiceDelegation{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*SubnetSpecDelegationServiceDelegation)(ptr) = obj
		} else {
			*(*SubnetSpecDelegationServiceDelegation)(ptr) = SubnetSpecDelegationServiceDelegation{}
		}
	default:
		iter.ReportError("decode SubnetSpecDelegationServiceDelegation", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ServiceEndpointStoragePolicySpecDefinitionCodec struct {
}

func (ServiceEndpointStoragePolicySpecDefinitionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServiceEndpointStoragePolicySpecDefinition)(ptr) == nil
}

func (ServiceEndpointStoragePolicySpecDefinitionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServiceEndpointStoragePolicySpecDefinition)(ptr)
	var objs []ServiceEndpointStoragePolicySpecDefinition
	if obj != nil {
		objs = []ServiceEndpointStoragePolicySpecDefinition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceEndpointStoragePolicySpecDefinition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServiceEndpointStoragePolicySpecDefinitionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServiceEndpointStoragePolicySpecDefinition)(ptr) = ServiceEndpointStoragePolicySpecDefinition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServiceEndpointStoragePolicySpecDefinition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceEndpointStoragePolicySpecDefinition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServiceEndpointStoragePolicySpecDefinition)(ptr) = objs[0]
			} else {
				*(*ServiceEndpointStoragePolicySpecDefinition)(ptr) = ServiceEndpointStoragePolicySpecDefinition{}
			}
		} else {
			*(*ServiceEndpointStoragePolicySpecDefinition)(ptr) = ServiceEndpointStoragePolicySpecDefinition{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ServiceEndpointStoragePolicySpecDefinition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceEndpointStoragePolicySpecDefinition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ServiceEndpointStoragePolicySpecDefinition)(ptr) = obj
		} else {
			*(*ServiceEndpointStoragePolicySpecDefinition)(ptr) = ServiceEndpointStoragePolicySpecDefinition{}
		}
	default:
		iter.ReportError("decode ServiceEndpointStoragePolicySpecDefinition", "unexpected JSON type")
	}
}