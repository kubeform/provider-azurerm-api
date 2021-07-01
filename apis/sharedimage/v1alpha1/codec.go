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
		jsoniter.MustGetKind(reflect2.TypeOf(SharedImageSpecIdentifier{}).Type1()):   SharedImageSpecIdentifierCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SharedImageSpecPurchasePlan{}).Type1()): SharedImageSpecPurchasePlanCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(SharedImageSpecIdentifier{}).Type1()):   SharedImageSpecIdentifierCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SharedImageSpecPurchasePlan{}).Type1()): SharedImageSpecPurchasePlanCodec{},
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
type SharedImageSpecIdentifierCodec struct {
}

func (SharedImageSpecIdentifierCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SharedImageSpecIdentifier)(ptr) == nil
}

func (SharedImageSpecIdentifierCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SharedImageSpecIdentifier)(ptr)
	var objs []SharedImageSpecIdentifier
	if obj != nil {
		objs = []SharedImageSpecIdentifier{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SharedImageSpecIdentifier{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SharedImageSpecIdentifierCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SharedImageSpecIdentifier)(ptr) = SharedImageSpecIdentifier{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SharedImageSpecIdentifier

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SharedImageSpecIdentifier{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SharedImageSpecIdentifier)(ptr) = objs[0]
			} else {
				*(*SharedImageSpecIdentifier)(ptr) = SharedImageSpecIdentifier{}
			}
		} else {
			*(*SharedImageSpecIdentifier)(ptr) = SharedImageSpecIdentifier{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj SharedImageSpecIdentifier

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SharedImageSpecIdentifier{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*SharedImageSpecIdentifier)(ptr) = obj
		} else {
			*(*SharedImageSpecIdentifier)(ptr) = SharedImageSpecIdentifier{}
		}
	default:
		iter.ReportError("decode SharedImageSpecIdentifier", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type SharedImageSpecPurchasePlanCodec struct {
}

func (SharedImageSpecPurchasePlanCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SharedImageSpecPurchasePlan)(ptr) == nil
}

func (SharedImageSpecPurchasePlanCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SharedImageSpecPurchasePlan)(ptr)
	var objs []SharedImageSpecPurchasePlan
	if obj != nil {
		objs = []SharedImageSpecPurchasePlan{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SharedImageSpecPurchasePlan{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SharedImageSpecPurchasePlanCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SharedImageSpecPurchasePlan)(ptr) = SharedImageSpecPurchasePlan{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SharedImageSpecPurchasePlan

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SharedImageSpecPurchasePlan{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SharedImageSpecPurchasePlan)(ptr) = objs[0]
			} else {
				*(*SharedImageSpecPurchasePlan)(ptr) = SharedImageSpecPurchasePlan{}
			}
		} else {
			*(*SharedImageSpecPurchasePlan)(ptr) = SharedImageSpecPurchasePlan{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj SharedImageSpecPurchasePlan

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SharedImageSpecPurchasePlan{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*SharedImageSpecPurchasePlan)(ptr) = obj
		} else {
			*(*SharedImageSpecPurchasePlan)(ptr) = SharedImageSpecPurchasePlan{}
		}
	default:
		iter.ReportError("decode SharedImageSpecPurchasePlan", "unexpected JSON type")
	}
}
