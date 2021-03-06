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
		jsoniter.MustGetKind(reflect2.TypeOf(ProfileSpecDnsConfig{}).Type1()):     ProfileSpecDnsConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProfileSpecMonitorConfig{}).Type1()): ProfileSpecMonitorConfigCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ProfileSpecDnsConfig{}).Type1()):     ProfileSpecDnsConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProfileSpecMonitorConfig{}).Type1()): ProfileSpecMonitorConfigCodec{},
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
type ProfileSpecDnsConfigCodec struct {
}

func (ProfileSpecDnsConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProfileSpecDnsConfig)(ptr) == nil
}

func (ProfileSpecDnsConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProfileSpecDnsConfig)(ptr)
	var objs []ProfileSpecDnsConfig
	if obj != nil {
		objs = []ProfileSpecDnsConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProfileSpecDnsConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProfileSpecDnsConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProfileSpecDnsConfig)(ptr) = ProfileSpecDnsConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProfileSpecDnsConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProfileSpecDnsConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProfileSpecDnsConfig)(ptr) = objs[0]
			} else {
				*(*ProfileSpecDnsConfig)(ptr) = ProfileSpecDnsConfig{}
			}
		} else {
			*(*ProfileSpecDnsConfig)(ptr) = ProfileSpecDnsConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ProfileSpecDnsConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProfileSpecDnsConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ProfileSpecDnsConfig)(ptr) = obj
		} else {
			*(*ProfileSpecDnsConfig)(ptr) = ProfileSpecDnsConfig{}
		}
	default:
		iter.ReportError("decode ProfileSpecDnsConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ProfileSpecMonitorConfigCodec struct {
}

func (ProfileSpecMonitorConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProfileSpecMonitorConfig)(ptr) == nil
}

func (ProfileSpecMonitorConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProfileSpecMonitorConfig)(ptr)
	var objs []ProfileSpecMonitorConfig
	if obj != nil {
		objs = []ProfileSpecMonitorConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProfileSpecMonitorConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProfileSpecMonitorConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProfileSpecMonitorConfig)(ptr) = ProfileSpecMonitorConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProfileSpecMonitorConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProfileSpecMonitorConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProfileSpecMonitorConfig)(ptr) = objs[0]
			} else {
				*(*ProfileSpecMonitorConfig)(ptr) = ProfileSpecMonitorConfig{}
			}
		} else {
			*(*ProfileSpecMonitorConfig)(ptr) = ProfileSpecMonitorConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ProfileSpecMonitorConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProfileSpecMonitorConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ProfileSpecMonitorConfig)(ptr) = obj
		} else {
			*(*ProfileSpecMonitorConfig)(ptr) = ProfileSpecMonitorConfig{}
		}
	default:
		iter.ReportError("decode ProfileSpecMonitorConfig", "unexpected JSON type")
	}
}
