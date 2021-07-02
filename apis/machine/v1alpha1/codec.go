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
		jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecIdentity{}).Type1()):      LearningComputeClusterSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecScaleSettings{}).Type1()): LearningComputeClusterSpecScaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningInferenceClusterSpecSsl{}).Type1()):         LearningInferenceClusterSpecSslCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningWorkspaceSpecIdentity{}).Type1()):           LearningWorkspaceSpecIdentityCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecIdentity{}).Type1()):      LearningComputeClusterSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecScaleSettings{}).Type1()): LearningComputeClusterSpecScaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningInferenceClusterSpecSsl{}).Type1()):         LearningInferenceClusterSpecSslCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningWorkspaceSpecIdentity{}).Type1()):           LearningWorkspaceSpecIdentityCodec{},
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
type LearningComputeClusterSpecIdentityCodec struct {
}

func (LearningComputeClusterSpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LearningComputeClusterSpecIdentity)(ptr) == nil
}

func (LearningComputeClusterSpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LearningComputeClusterSpecIdentity)(ptr)
	var objs []LearningComputeClusterSpecIdentity
	if obj != nil {
		objs = []LearningComputeClusterSpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LearningComputeClusterSpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LearningComputeClusterSpecIdentity)(ptr) = LearningComputeClusterSpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LearningComputeClusterSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LearningComputeClusterSpecIdentity)(ptr) = objs[0]
			} else {
				*(*LearningComputeClusterSpecIdentity)(ptr) = LearningComputeClusterSpecIdentity{}
			}
		} else {
			*(*LearningComputeClusterSpecIdentity)(ptr) = LearningComputeClusterSpecIdentity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LearningComputeClusterSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LearningComputeClusterSpecIdentity)(ptr) = obj
		} else {
			*(*LearningComputeClusterSpecIdentity)(ptr) = LearningComputeClusterSpecIdentity{}
		}
	default:
		iter.ReportError("decode LearningComputeClusterSpecIdentity", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LearningComputeClusterSpecScaleSettingsCodec struct {
}

func (LearningComputeClusterSpecScaleSettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LearningComputeClusterSpecScaleSettings)(ptr) == nil
}

func (LearningComputeClusterSpecScaleSettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LearningComputeClusterSpecScaleSettings)(ptr)
	var objs []LearningComputeClusterSpecScaleSettings
	if obj != nil {
		objs = []LearningComputeClusterSpecScaleSettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecScaleSettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LearningComputeClusterSpecScaleSettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LearningComputeClusterSpecScaleSettings)(ptr) = LearningComputeClusterSpecScaleSettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LearningComputeClusterSpecScaleSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecScaleSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LearningComputeClusterSpecScaleSettings)(ptr) = objs[0]
			} else {
				*(*LearningComputeClusterSpecScaleSettings)(ptr) = LearningComputeClusterSpecScaleSettings{}
			}
		} else {
			*(*LearningComputeClusterSpecScaleSettings)(ptr) = LearningComputeClusterSpecScaleSettings{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LearningComputeClusterSpecScaleSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecScaleSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LearningComputeClusterSpecScaleSettings)(ptr) = obj
		} else {
			*(*LearningComputeClusterSpecScaleSettings)(ptr) = LearningComputeClusterSpecScaleSettings{}
		}
	default:
		iter.ReportError("decode LearningComputeClusterSpecScaleSettings", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LearningInferenceClusterSpecSslCodec struct {
}

func (LearningInferenceClusterSpecSslCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LearningInferenceClusterSpecSsl)(ptr) == nil
}

func (LearningInferenceClusterSpecSslCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LearningInferenceClusterSpecSsl)(ptr)
	var objs []LearningInferenceClusterSpecSsl
	if obj != nil {
		objs = []LearningInferenceClusterSpecSsl{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningInferenceClusterSpecSsl{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LearningInferenceClusterSpecSslCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LearningInferenceClusterSpecSsl)(ptr) = LearningInferenceClusterSpecSsl{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LearningInferenceClusterSpecSsl

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningInferenceClusterSpecSsl{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LearningInferenceClusterSpecSsl)(ptr) = objs[0]
			} else {
				*(*LearningInferenceClusterSpecSsl)(ptr) = LearningInferenceClusterSpecSsl{}
			}
		} else {
			*(*LearningInferenceClusterSpecSsl)(ptr) = LearningInferenceClusterSpecSsl{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LearningInferenceClusterSpecSsl

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningInferenceClusterSpecSsl{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LearningInferenceClusterSpecSsl)(ptr) = obj
		} else {
			*(*LearningInferenceClusterSpecSsl)(ptr) = LearningInferenceClusterSpecSsl{}
		}
	default:
		iter.ReportError("decode LearningInferenceClusterSpecSsl", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LearningWorkspaceSpecIdentityCodec struct {
}

func (LearningWorkspaceSpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LearningWorkspaceSpecIdentity)(ptr) == nil
}

func (LearningWorkspaceSpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LearningWorkspaceSpecIdentity)(ptr)
	var objs []LearningWorkspaceSpecIdentity
	if obj != nil {
		objs = []LearningWorkspaceSpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningWorkspaceSpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LearningWorkspaceSpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LearningWorkspaceSpecIdentity)(ptr) = LearningWorkspaceSpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LearningWorkspaceSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningWorkspaceSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LearningWorkspaceSpecIdentity)(ptr) = objs[0]
			} else {
				*(*LearningWorkspaceSpecIdentity)(ptr) = LearningWorkspaceSpecIdentity{}
			}
		} else {
			*(*LearningWorkspaceSpecIdentity)(ptr) = LearningWorkspaceSpecIdentity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LearningWorkspaceSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningWorkspaceSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LearningWorkspaceSpecIdentity)(ptr) = obj
		} else {
			*(*LearningWorkspaceSpecIdentity)(ptr) = LearningWorkspaceSpecIdentity{}
		}
	default:
		iter.ReportError("decode LearningWorkspaceSpecIdentity", "unexpected JSON type")
	}
}