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
		jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecInitialReplicaSet{}).Type1()): DirectoryDomainServiceSpecInitialReplicaSetCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecNotifications{}).Type1()):     DirectoryDomainServiceSpecNotificationsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecSecureLdap{}).Type1()):        DirectoryDomainServiceSpecSecureLdapCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecSecurity{}).Type1()):          DirectoryDomainServiceSpecSecurityCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecInitialReplicaSet{}).Type1()): DirectoryDomainServiceSpecInitialReplicaSetCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecNotifications{}).Type1()):     DirectoryDomainServiceSpecNotificationsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecSecureLdap{}).Type1()):        DirectoryDomainServiceSpecSecureLdapCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecSecurity{}).Type1()):          DirectoryDomainServiceSpecSecurityCodec{},
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
type DirectoryDomainServiceSpecInitialReplicaSetCodec struct {
}

func (DirectoryDomainServiceSpecInitialReplicaSetCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*DirectoryDomainServiceSpecInitialReplicaSet)(ptr) == nil
}

func (DirectoryDomainServiceSpecInitialReplicaSetCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*DirectoryDomainServiceSpecInitialReplicaSet)(ptr)
	var objs []DirectoryDomainServiceSpecInitialReplicaSet
	if obj != nil {
		objs = []DirectoryDomainServiceSpecInitialReplicaSet{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecInitialReplicaSet{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (DirectoryDomainServiceSpecInitialReplicaSetCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*DirectoryDomainServiceSpecInitialReplicaSet)(ptr) = DirectoryDomainServiceSpecInitialReplicaSet{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []DirectoryDomainServiceSpecInitialReplicaSet

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecInitialReplicaSet{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*DirectoryDomainServiceSpecInitialReplicaSet)(ptr) = objs[0]
			} else {
				*(*DirectoryDomainServiceSpecInitialReplicaSet)(ptr) = DirectoryDomainServiceSpecInitialReplicaSet{}
			}
		} else {
			*(*DirectoryDomainServiceSpecInitialReplicaSet)(ptr) = DirectoryDomainServiceSpecInitialReplicaSet{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj DirectoryDomainServiceSpecInitialReplicaSet

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecInitialReplicaSet{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*DirectoryDomainServiceSpecInitialReplicaSet)(ptr) = obj
		} else {
			*(*DirectoryDomainServiceSpecInitialReplicaSet)(ptr) = DirectoryDomainServiceSpecInitialReplicaSet{}
		}
	default:
		iter.ReportError("decode DirectoryDomainServiceSpecInitialReplicaSet", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type DirectoryDomainServiceSpecNotificationsCodec struct {
}

func (DirectoryDomainServiceSpecNotificationsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*DirectoryDomainServiceSpecNotifications)(ptr) == nil
}

func (DirectoryDomainServiceSpecNotificationsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*DirectoryDomainServiceSpecNotifications)(ptr)
	var objs []DirectoryDomainServiceSpecNotifications
	if obj != nil {
		objs = []DirectoryDomainServiceSpecNotifications{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecNotifications{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (DirectoryDomainServiceSpecNotificationsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*DirectoryDomainServiceSpecNotifications)(ptr) = DirectoryDomainServiceSpecNotifications{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []DirectoryDomainServiceSpecNotifications

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecNotifications{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*DirectoryDomainServiceSpecNotifications)(ptr) = objs[0]
			} else {
				*(*DirectoryDomainServiceSpecNotifications)(ptr) = DirectoryDomainServiceSpecNotifications{}
			}
		} else {
			*(*DirectoryDomainServiceSpecNotifications)(ptr) = DirectoryDomainServiceSpecNotifications{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj DirectoryDomainServiceSpecNotifications

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecNotifications{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*DirectoryDomainServiceSpecNotifications)(ptr) = obj
		} else {
			*(*DirectoryDomainServiceSpecNotifications)(ptr) = DirectoryDomainServiceSpecNotifications{}
		}
	default:
		iter.ReportError("decode DirectoryDomainServiceSpecNotifications", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type DirectoryDomainServiceSpecSecureLdapCodec struct {
}

func (DirectoryDomainServiceSpecSecureLdapCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*DirectoryDomainServiceSpecSecureLdap)(ptr) == nil
}

func (DirectoryDomainServiceSpecSecureLdapCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*DirectoryDomainServiceSpecSecureLdap)(ptr)
	var objs []DirectoryDomainServiceSpecSecureLdap
	if obj != nil {
		objs = []DirectoryDomainServiceSpecSecureLdap{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecSecureLdap{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (DirectoryDomainServiceSpecSecureLdapCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*DirectoryDomainServiceSpecSecureLdap)(ptr) = DirectoryDomainServiceSpecSecureLdap{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []DirectoryDomainServiceSpecSecureLdap

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecSecureLdap{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*DirectoryDomainServiceSpecSecureLdap)(ptr) = objs[0]
			} else {
				*(*DirectoryDomainServiceSpecSecureLdap)(ptr) = DirectoryDomainServiceSpecSecureLdap{}
			}
		} else {
			*(*DirectoryDomainServiceSpecSecureLdap)(ptr) = DirectoryDomainServiceSpecSecureLdap{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj DirectoryDomainServiceSpecSecureLdap

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecSecureLdap{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*DirectoryDomainServiceSpecSecureLdap)(ptr) = obj
		} else {
			*(*DirectoryDomainServiceSpecSecureLdap)(ptr) = DirectoryDomainServiceSpecSecureLdap{}
		}
	default:
		iter.ReportError("decode DirectoryDomainServiceSpecSecureLdap", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type DirectoryDomainServiceSpecSecurityCodec struct {
}

func (DirectoryDomainServiceSpecSecurityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*DirectoryDomainServiceSpecSecurity)(ptr) == nil
}

func (DirectoryDomainServiceSpecSecurityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*DirectoryDomainServiceSpecSecurity)(ptr)
	var objs []DirectoryDomainServiceSpecSecurity
	if obj != nil {
		objs = []DirectoryDomainServiceSpecSecurity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecSecurity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (DirectoryDomainServiceSpecSecurityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*DirectoryDomainServiceSpecSecurity)(ptr) = DirectoryDomainServiceSpecSecurity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []DirectoryDomainServiceSpecSecurity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecSecurity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*DirectoryDomainServiceSpecSecurity)(ptr) = objs[0]
			} else {
				*(*DirectoryDomainServiceSpecSecurity)(ptr) = DirectoryDomainServiceSpecSecurity{}
			}
		} else {
			*(*DirectoryDomainServiceSpecSecurity)(ptr) = DirectoryDomainServiceSpecSecurity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj DirectoryDomainServiceSpecSecurity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DirectoryDomainServiceSpecSecurity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*DirectoryDomainServiceSpecSecurity)(ptr) = obj
		} else {
			*(*DirectoryDomainServiceSpecSecurity)(ptr) = DirectoryDomainServiceSpecSecurity{}
		}
	default:
		iter.ReportError("decode DirectoryDomainServiceSpecSecurity", "unexpected JSON type")
	}
}
