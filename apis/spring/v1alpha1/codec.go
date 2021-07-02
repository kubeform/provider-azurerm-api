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
		jsoniter.MustGetKind(reflect2.TypeOf(CloudAppSpecIdentity{}).Type1()):                                          CloudAppSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CloudAppSpecPersistentDisk{}).Type1()):                                    CloudAppSpecPersistentDiskCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSetting{}).Type1()):                        CloudServiceSpecConfigServerGitSettingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingHttpBasicAuth{}).Type1()):           CloudServiceSpecConfigServerGitSettingHttpBasicAuthCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth{}).Type1()): CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuthCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingRepositorySshAuth{}).Type1()):       CloudServiceSpecConfigServerGitSettingRepositorySshAuthCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingSshAuth{}).Type1()):                 CloudServiceSpecConfigServerGitSettingSshAuthCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecNetwork{}).Type1()):                                       CloudServiceSpecNetworkCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecTrace{}).Type1()):                                         CloudServiceSpecTraceCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(CloudAppSpecIdentity{}).Type1()):                                          CloudAppSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CloudAppSpecPersistentDisk{}).Type1()):                                    CloudAppSpecPersistentDiskCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSetting{}).Type1()):                        CloudServiceSpecConfigServerGitSettingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingHttpBasicAuth{}).Type1()):           CloudServiceSpecConfigServerGitSettingHttpBasicAuthCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth{}).Type1()): CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuthCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingRepositorySshAuth{}).Type1()):       CloudServiceSpecConfigServerGitSettingRepositorySshAuthCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingSshAuth{}).Type1()):                 CloudServiceSpecConfigServerGitSettingSshAuthCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecNetwork{}).Type1()):                                       CloudServiceSpecNetworkCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecTrace{}).Type1()):                                         CloudServiceSpecTraceCodec{},
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
type CloudAppSpecIdentityCodec struct {
}

func (CloudAppSpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CloudAppSpecIdentity)(ptr) == nil
}

func (CloudAppSpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CloudAppSpecIdentity)(ptr)
	var objs []CloudAppSpecIdentity
	if obj != nil {
		objs = []CloudAppSpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudAppSpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CloudAppSpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CloudAppSpecIdentity)(ptr) = CloudAppSpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CloudAppSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudAppSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CloudAppSpecIdentity)(ptr) = objs[0]
			} else {
				*(*CloudAppSpecIdentity)(ptr) = CloudAppSpecIdentity{}
			}
		} else {
			*(*CloudAppSpecIdentity)(ptr) = CloudAppSpecIdentity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CloudAppSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudAppSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CloudAppSpecIdentity)(ptr) = obj
		} else {
			*(*CloudAppSpecIdentity)(ptr) = CloudAppSpecIdentity{}
		}
	default:
		iter.ReportError("decode CloudAppSpecIdentity", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CloudAppSpecPersistentDiskCodec struct {
}

func (CloudAppSpecPersistentDiskCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CloudAppSpecPersistentDisk)(ptr) == nil
}

func (CloudAppSpecPersistentDiskCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CloudAppSpecPersistentDisk)(ptr)
	var objs []CloudAppSpecPersistentDisk
	if obj != nil {
		objs = []CloudAppSpecPersistentDisk{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudAppSpecPersistentDisk{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CloudAppSpecPersistentDiskCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CloudAppSpecPersistentDisk)(ptr) = CloudAppSpecPersistentDisk{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CloudAppSpecPersistentDisk

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudAppSpecPersistentDisk{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CloudAppSpecPersistentDisk)(ptr) = objs[0]
			} else {
				*(*CloudAppSpecPersistentDisk)(ptr) = CloudAppSpecPersistentDisk{}
			}
		} else {
			*(*CloudAppSpecPersistentDisk)(ptr) = CloudAppSpecPersistentDisk{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CloudAppSpecPersistentDisk

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudAppSpecPersistentDisk{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CloudAppSpecPersistentDisk)(ptr) = obj
		} else {
			*(*CloudAppSpecPersistentDisk)(ptr) = CloudAppSpecPersistentDisk{}
		}
	default:
		iter.ReportError("decode CloudAppSpecPersistentDisk", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CloudServiceSpecConfigServerGitSettingCodec struct {
}

func (CloudServiceSpecConfigServerGitSettingCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CloudServiceSpecConfigServerGitSetting)(ptr) == nil
}

func (CloudServiceSpecConfigServerGitSettingCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CloudServiceSpecConfigServerGitSetting)(ptr)
	var objs []CloudServiceSpecConfigServerGitSetting
	if obj != nil {
		objs = []CloudServiceSpecConfigServerGitSetting{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSetting{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CloudServiceSpecConfigServerGitSettingCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CloudServiceSpecConfigServerGitSetting)(ptr) = CloudServiceSpecConfigServerGitSetting{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CloudServiceSpecConfigServerGitSetting

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSetting{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CloudServiceSpecConfigServerGitSetting)(ptr) = objs[0]
			} else {
				*(*CloudServiceSpecConfigServerGitSetting)(ptr) = CloudServiceSpecConfigServerGitSetting{}
			}
		} else {
			*(*CloudServiceSpecConfigServerGitSetting)(ptr) = CloudServiceSpecConfigServerGitSetting{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CloudServiceSpecConfigServerGitSetting

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSetting{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CloudServiceSpecConfigServerGitSetting)(ptr) = obj
		} else {
			*(*CloudServiceSpecConfigServerGitSetting)(ptr) = CloudServiceSpecConfigServerGitSetting{}
		}
	default:
		iter.ReportError("decode CloudServiceSpecConfigServerGitSetting", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CloudServiceSpecConfigServerGitSettingHttpBasicAuthCodec struct {
}

func (CloudServiceSpecConfigServerGitSettingHttpBasicAuthCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CloudServiceSpecConfigServerGitSettingHttpBasicAuth)(ptr) == nil
}

func (CloudServiceSpecConfigServerGitSettingHttpBasicAuthCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CloudServiceSpecConfigServerGitSettingHttpBasicAuth)(ptr)
	var objs []CloudServiceSpecConfigServerGitSettingHttpBasicAuth
	if obj != nil {
		objs = []CloudServiceSpecConfigServerGitSettingHttpBasicAuth{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingHttpBasicAuth{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CloudServiceSpecConfigServerGitSettingHttpBasicAuthCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CloudServiceSpecConfigServerGitSettingHttpBasicAuth)(ptr) = CloudServiceSpecConfigServerGitSettingHttpBasicAuth{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CloudServiceSpecConfigServerGitSettingHttpBasicAuth

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingHttpBasicAuth{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CloudServiceSpecConfigServerGitSettingHttpBasicAuth)(ptr) = objs[0]
			} else {
				*(*CloudServiceSpecConfigServerGitSettingHttpBasicAuth)(ptr) = CloudServiceSpecConfigServerGitSettingHttpBasicAuth{}
			}
		} else {
			*(*CloudServiceSpecConfigServerGitSettingHttpBasicAuth)(ptr) = CloudServiceSpecConfigServerGitSettingHttpBasicAuth{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CloudServiceSpecConfigServerGitSettingHttpBasicAuth

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingHttpBasicAuth{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CloudServiceSpecConfigServerGitSettingHttpBasicAuth)(ptr) = obj
		} else {
			*(*CloudServiceSpecConfigServerGitSettingHttpBasicAuth)(ptr) = CloudServiceSpecConfigServerGitSettingHttpBasicAuth{}
		}
	default:
		iter.ReportError("decode CloudServiceSpecConfigServerGitSettingHttpBasicAuth", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuthCodec struct {
}

func (CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuthCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth)(ptr) == nil
}

func (CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuthCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth)(ptr)
	var objs []CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth
	if obj != nil {
		objs = []CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuthCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth)(ptr) = CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth)(ptr) = objs[0]
			} else {
				*(*CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth)(ptr) = CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth{}
			}
		} else {
			*(*CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth)(ptr) = CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth)(ptr) = obj
		} else {
			*(*CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth)(ptr) = CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth{}
		}
	default:
		iter.ReportError("decode CloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CloudServiceSpecConfigServerGitSettingRepositorySshAuthCodec struct {
}

func (CloudServiceSpecConfigServerGitSettingRepositorySshAuthCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CloudServiceSpecConfigServerGitSettingRepositorySshAuth)(ptr) == nil
}

func (CloudServiceSpecConfigServerGitSettingRepositorySshAuthCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CloudServiceSpecConfigServerGitSettingRepositorySshAuth)(ptr)
	var objs []CloudServiceSpecConfigServerGitSettingRepositorySshAuth
	if obj != nil {
		objs = []CloudServiceSpecConfigServerGitSettingRepositorySshAuth{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingRepositorySshAuth{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CloudServiceSpecConfigServerGitSettingRepositorySshAuthCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CloudServiceSpecConfigServerGitSettingRepositorySshAuth)(ptr) = CloudServiceSpecConfigServerGitSettingRepositorySshAuth{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CloudServiceSpecConfigServerGitSettingRepositorySshAuth

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingRepositorySshAuth{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CloudServiceSpecConfigServerGitSettingRepositorySshAuth)(ptr) = objs[0]
			} else {
				*(*CloudServiceSpecConfigServerGitSettingRepositorySshAuth)(ptr) = CloudServiceSpecConfigServerGitSettingRepositorySshAuth{}
			}
		} else {
			*(*CloudServiceSpecConfigServerGitSettingRepositorySshAuth)(ptr) = CloudServiceSpecConfigServerGitSettingRepositorySshAuth{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CloudServiceSpecConfigServerGitSettingRepositorySshAuth

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingRepositorySshAuth{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CloudServiceSpecConfigServerGitSettingRepositorySshAuth)(ptr) = obj
		} else {
			*(*CloudServiceSpecConfigServerGitSettingRepositorySshAuth)(ptr) = CloudServiceSpecConfigServerGitSettingRepositorySshAuth{}
		}
	default:
		iter.ReportError("decode CloudServiceSpecConfigServerGitSettingRepositorySshAuth", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CloudServiceSpecConfigServerGitSettingSshAuthCodec struct {
}

func (CloudServiceSpecConfigServerGitSettingSshAuthCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CloudServiceSpecConfigServerGitSettingSshAuth)(ptr) == nil
}

func (CloudServiceSpecConfigServerGitSettingSshAuthCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CloudServiceSpecConfigServerGitSettingSshAuth)(ptr)
	var objs []CloudServiceSpecConfigServerGitSettingSshAuth
	if obj != nil {
		objs = []CloudServiceSpecConfigServerGitSettingSshAuth{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingSshAuth{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CloudServiceSpecConfigServerGitSettingSshAuthCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CloudServiceSpecConfigServerGitSettingSshAuth)(ptr) = CloudServiceSpecConfigServerGitSettingSshAuth{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CloudServiceSpecConfigServerGitSettingSshAuth

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingSshAuth{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CloudServiceSpecConfigServerGitSettingSshAuth)(ptr) = objs[0]
			} else {
				*(*CloudServiceSpecConfigServerGitSettingSshAuth)(ptr) = CloudServiceSpecConfigServerGitSettingSshAuth{}
			}
		} else {
			*(*CloudServiceSpecConfigServerGitSettingSshAuth)(ptr) = CloudServiceSpecConfigServerGitSettingSshAuth{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CloudServiceSpecConfigServerGitSettingSshAuth

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecConfigServerGitSettingSshAuth{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CloudServiceSpecConfigServerGitSettingSshAuth)(ptr) = obj
		} else {
			*(*CloudServiceSpecConfigServerGitSettingSshAuth)(ptr) = CloudServiceSpecConfigServerGitSettingSshAuth{}
		}
	default:
		iter.ReportError("decode CloudServiceSpecConfigServerGitSettingSshAuth", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CloudServiceSpecNetworkCodec struct {
}

func (CloudServiceSpecNetworkCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CloudServiceSpecNetwork)(ptr) == nil
}

func (CloudServiceSpecNetworkCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CloudServiceSpecNetwork)(ptr)
	var objs []CloudServiceSpecNetwork
	if obj != nil {
		objs = []CloudServiceSpecNetwork{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecNetwork{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CloudServiceSpecNetworkCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CloudServiceSpecNetwork)(ptr) = CloudServiceSpecNetwork{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CloudServiceSpecNetwork

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecNetwork{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CloudServiceSpecNetwork)(ptr) = objs[0]
			} else {
				*(*CloudServiceSpecNetwork)(ptr) = CloudServiceSpecNetwork{}
			}
		} else {
			*(*CloudServiceSpecNetwork)(ptr) = CloudServiceSpecNetwork{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CloudServiceSpecNetwork

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecNetwork{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CloudServiceSpecNetwork)(ptr) = obj
		} else {
			*(*CloudServiceSpecNetwork)(ptr) = CloudServiceSpecNetwork{}
		}
	default:
		iter.ReportError("decode CloudServiceSpecNetwork", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CloudServiceSpecTraceCodec struct {
}

func (CloudServiceSpecTraceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CloudServiceSpecTrace)(ptr) == nil
}

func (CloudServiceSpecTraceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CloudServiceSpecTrace)(ptr)
	var objs []CloudServiceSpecTrace
	if obj != nil {
		objs = []CloudServiceSpecTrace{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecTrace{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CloudServiceSpecTraceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CloudServiceSpecTrace)(ptr) = CloudServiceSpecTrace{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CloudServiceSpecTrace

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecTrace{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CloudServiceSpecTrace)(ptr) = objs[0]
			} else {
				*(*CloudServiceSpecTrace)(ptr) = CloudServiceSpecTrace{}
			}
		} else {
			*(*CloudServiceSpecTrace)(ptr) = CloudServiceSpecTrace{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CloudServiceSpecTrace

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CloudServiceSpecTrace{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CloudServiceSpecTrace)(ptr) = obj
		} else {
			*(*CloudServiceSpecTrace)(ptr) = CloudServiceSpecTrace{}
		}
	default:
		iter.ReportError("decode CloudServiceSpecTrace", "unexpected JSON type")
	}
}