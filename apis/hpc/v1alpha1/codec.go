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
		jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDefaultAccessPolicy{}).Type1()):      CacheSpecDefaultAccessPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryActiveDirectory{}).Type1()): CacheSpecDirectoryActiveDirectoryCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryFlatFile{}).Type1()):        CacheSpecDirectoryFlatFileCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryLdap{}).Type1()):            CacheSpecDirectoryLdapCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryLdapBind{}).Type1()):        CacheSpecDirectoryLdapBindCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDns{}).Type1()):                      CacheSpecDnsCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDefaultAccessPolicy{}).Type1()):      CacheSpecDefaultAccessPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryActiveDirectory{}).Type1()): CacheSpecDirectoryActiveDirectoryCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryFlatFile{}).Type1()):        CacheSpecDirectoryFlatFileCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryLdap{}).Type1()):            CacheSpecDirectoryLdapCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryLdapBind{}).Type1()):        CacheSpecDirectoryLdapBindCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDns{}).Type1()):                      CacheSpecDnsCodec{},
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
type CacheSpecDefaultAccessPolicyCodec struct {
}

func (CacheSpecDefaultAccessPolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CacheSpecDefaultAccessPolicy)(ptr) == nil
}

func (CacheSpecDefaultAccessPolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CacheSpecDefaultAccessPolicy)(ptr)
	var objs []CacheSpecDefaultAccessPolicy
	if obj != nil {
		objs = []CacheSpecDefaultAccessPolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDefaultAccessPolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CacheSpecDefaultAccessPolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CacheSpecDefaultAccessPolicy)(ptr) = CacheSpecDefaultAccessPolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CacheSpecDefaultAccessPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDefaultAccessPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CacheSpecDefaultAccessPolicy)(ptr) = objs[0]
			} else {
				*(*CacheSpecDefaultAccessPolicy)(ptr) = CacheSpecDefaultAccessPolicy{}
			}
		} else {
			*(*CacheSpecDefaultAccessPolicy)(ptr) = CacheSpecDefaultAccessPolicy{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CacheSpecDefaultAccessPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDefaultAccessPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CacheSpecDefaultAccessPolicy)(ptr) = obj
		} else {
			*(*CacheSpecDefaultAccessPolicy)(ptr) = CacheSpecDefaultAccessPolicy{}
		}
	default:
		iter.ReportError("decode CacheSpecDefaultAccessPolicy", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CacheSpecDirectoryActiveDirectoryCodec struct {
}

func (CacheSpecDirectoryActiveDirectoryCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CacheSpecDirectoryActiveDirectory)(ptr) == nil
}

func (CacheSpecDirectoryActiveDirectoryCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CacheSpecDirectoryActiveDirectory)(ptr)
	var objs []CacheSpecDirectoryActiveDirectory
	if obj != nil {
		objs = []CacheSpecDirectoryActiveDirectory{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryActiveDirectory{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CacheSpecDirectoryActiveDirectoryCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CacheSpecDirectoryActiveDirectory)(ptr) = CacheSpecDirectoryActiveDirectory{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CacheSpecDirectoryActiveDirectory

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryActiveDirectory{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CacheSpecDirectoryActiveDirectory)(ptr) = objs[0]
			} else {
				*(*CacheSpecDirectoryActiveDirectory)(ptr) = CacheSpecDirectoryActiveDirectory{}
			}
		} else {
			*(*CacheSpecDirectoryActiveDirectory)(ptr) = CacheSpecDirectoryActiveDirectory{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CacheSpecDirectoryActiveDirectory

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryActiveDirectory{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CacheSpecDirectoryActiveDirectory)(ptr) = obj
		} else {
			*(*CacheSpecDirectoryActiveDirectory)(ptr) = CacheSpecDirectoryActiveDirectory{}
		}
	default:
		iter.ReportError("decode CacheSpecDirectoryActiveDirectory", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CacheSpecDirectoryFlatFileCodec struct {
}

func (CacheSpecDirectoryFlatFileCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CacheSpecDirectoryFlatFile)(ptr) == nil
}

func (CacheSpecDirectoryFlatFileCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CacheSpecDirectoryFlatFile)(ptr)
	var objs []CacheSpecDirectoryFlatFile
	if obj != nil {
		objs = []CacheSpecDirectoryFlatFile{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryFlatFile{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CacheSpecDirectoryFlatFileCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CacheSpecDirectoryFlatFile)(ptr) = CacheSpecDirectoryFlatFile{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CacheSpecDirectoryFlatFile

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryFlatFile{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CacheSpecDirectoryFlatFile)(ptr) = objs[0]
			} else {
				*(*CacheSpecDirectoryFlatFile)(ptr) = CacheSpecDirectoryFlatFile{}
			}
		} else {
			*(*CacheSpecDirectoryFlatFile)(ptr) = CacheSpecDirectoryFlatFile{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CacheSpecDirectoryFlatFile

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryFlatFile{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CacheSpecDirectoryFlatFile)(ptr) = obj
		} else {
			*(*CacheSpecDirectoryFlatFile)(ptr) = CacheSpecDirectoryFlatFile{}
		}
	default:
		iter.ReportError("decode CacheSpecDirectoryFlatFile", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CacheSpecDirectoryLdapCodec struct {
}

func (CacheSpecDirectoryLdapCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CacheSpecDirectoryLdap)(ptr) == nil
}

func (CacheSpecDirectoryLdapCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CacheSpecDirectoryLdap)(ptr)
	var objs []CacheSpecDirectoryLdap
	if obj != nil {
		objs = []CacheSpecDirectoryLdap{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryLdap{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CacheSpecDirectoryLdapCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CacheSpecDirectoryLdap)(ptr) = CacheSpecDirectoryLdap{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CacheSpecDirectoryLdap

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryLdap{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CacheSpecDirectoryLdap)(ptr) = objs[0]
			} else {
				*(*CacheSpecDirectoryLdap)(ptr) = CacheSpecDirectoryLdap{}
			}
		} else {
			*(*CacheSpecDirectoryLdap)(ptr) = CacheSpecDirectoryLdap{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CacheSpecDirectoryLdap

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryLdap{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CacheSpecDirectoryLdap)(ptr) = obj
		} else {
			*(*CacheSpecDirectoryLdap)(ptr) = CacheSpecDirectoryLdap{}
		}
	default:
		iter.ReportError("decode CacheSpecDirectoryLdap", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CacheSpecDirectoryLdapBindCodec struct {
}

func (CacheSpecDirectoryLdapBindCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CacheSpecDirectoryLdapBind)(ptr) == nil
}

func (CacheSpecDirectoryLdapBindCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CacheSpecDirectoryLdapBind)(ptr)
	var objs []CacheSpecDirectoryLdapBind
	if obj != nil {
		objs = []CacheSpecDirectoryLdapBind{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryLdapBind{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CacheSpecDirectoryLdapBindCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CacheSpecDirectoryLdapBind)(ptr) = CacheSpecDirectoryLdapBind{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CacheSpecDirectoryLdapBind

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryLdapBind{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CacheSpecDirectoryLdapBind)(ptr) = objs[0]
			} else {
				*(*CacheSpecDirectoryLdapBind)(ptr) = CacheSpecDirectoryLdapBind{}
			}
		} else {
			*(*CacheSpecDirectoryLdapBind)(ptr) = CacheSpecDirectoryLdapBind{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CacheSpecDirectoryLdapBind

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDirectoryLdapBind{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CacheSpecDirectoryLdapBind)(ptr) = obj
		} else {
			*(*CacheSpecDirectoryLdapBind)(ptr) = CacheSpecDirectoryLdapBind{}
		}
	default:
		iter.ReportError("decode CacheSpecDirectoryLdapBind", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CacheSpecDnsCodec struct {
}

func (CacheSpecDnsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CacheSpecDns)(ptr) == nil
}

func (CacheSpecDnsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CacheSpecDns)(ptr)
	var objs []CacheSpecDns
	if obj != nil {
		objs = []CacheSpecDns{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDns{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CacheSpecDnsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CacheSpecDns)(ptr) = CacheSpecDns{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CacheSpecDns

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDns{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CacheSpecDns)(ptr) = objs[0]
			} else {
				*(*CacheSpecDns)(ptr) = CacheSpecDns{}
			}
		} else {
			*(*CacheSpecDns)(ptr) = CacheSpecDns{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CacheSpecDns

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CacheSpecDns{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CacheSpecDns)(ptr) = obj
		} else {
			*(*CacheSpecDns)(ptr) = CacheSpecDns{}
		}
	default:
		iter.ReportError("decode CacheSpecDns", "unexpected JSON type")
	}
}
