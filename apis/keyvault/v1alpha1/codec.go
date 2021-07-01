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
		jsoniter.MustGetKind(reflect2.TypeOf(KeyVaultSpecNetworkAcls{}).Type1()):                                                          KeyVaultSpecNetworkAclsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificate{}).Type1()):                                                       CertificateSpecCertificateCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicy{}).Type1()):                                                 CertificateSpecCertificatePolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyIssuerParameters{}).Type1()):                                 CertificateSpecCertificatePolicyIssuerParametersCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyKeyProperties{}).Type1()):                                    CertificateSpecCertificatePolicyKeyPropertiesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyLifetimeActionAction{}).Type1()):                             CertificateSpecCertificatePolicyLifetimeActionActionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyLifetimeActionTrigger{}).Type1()):                            CertificateSpecCertificatePolicyLifetimeActionTriggerCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicySecretProperties{}).Type1()):                                 CertificateSpecCertificatePolicySecretPropertiesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyX509CertificateProperties{}).Type1()):                        CertificateSpecCertificatePolicyX509CertificatePropertiesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames{}).Type1()): CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(KeyVaultSpecNetworkAcls{}).Type1()):                                                          KeyVaultSpecNetworkAclsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificate{}).Type1()):                                                       CertificateSpecCertificateCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicy{}).Type1()):                                                 CertificateSpecCertificatePolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyIssuerParameters{}).Type1()):                                 CertificateSpecCertificatePolicyIssuerParametersCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyKeyProperties{}).Type1()):                                    CertificateSpecCertificatePolicyKeyPropertiesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyLifetimeActionAction{}).Type1()):                             CertificateSpecCertificatePolicyLifetimeActionActionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyLifetimeActionTrigger{}).Type1()):                            CertificateSpecCertificatePolicyLifetimeActionTriggerCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicySecretProperties{}).Type1()):                                 CertificateSpecCertificatePolicySecretPropertiesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyX509CertificateProperties{}).Type1()):                        CertificateSpecCertificatePolicyX509CertificatePropertiesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames{}).Type1()): CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesCodec{},
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
type KeyVaultSpecNetworkAclsCodec struct {
}

func (KeyVaultSpecNetworkAclsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*KeyVaultSpecNetworkAcls)(ptr) == nil
}

func (KeyVaultSpecNetworkAclsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*KeyVaultSpecNetworkAcls)(ptr)
	var objs []KeyVaultSpecNetworkAcls
	if obj != nil {
		objs = []KeyVaultSpecNetworkAcls{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(KeyVaultSpecNetworkAcls{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (KeyVaultSpecNetworkAclsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*KeyVaultSpecNetworkAcls)(ptr) = KeyVaultSpecNetworkAcls{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []KeyVaultSpecNetworkAcls

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(KeyVaultSpecNetworkAcls{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*KeyVaultSpecNetworkAcls)(ptr) = objs[0]
			} else {
				*(*KeyVaultSpecNetworkAcls)(ptr) = KeyVaultSpecNetworkAcls{}
			}
		} else {
			*(*KeyVaultSpecNetworkAcls)(ptr) = KeyVaultSpecNetworkAcls{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj KeyVaultSpecNetworkAcls

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(KeyVaultSpecNetworkAcls{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*KeyVaultSpecNetworkAcls)(ptr) = obj
		} else {
			*(*KeyVaultSpecNetworkAcls)(ptr) = KeyVaultSpecNetworkAcls{}
		}
	default:
		iter.ReportError("decode KeyVaultSpecNetworkAcls", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CertificateSpecCertificateCodec struct {
}

func (CertificateSpecCertificateCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CertificateSpecCertificate)(ptr) == nil
}

func (CertificateSpecCertificateCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CertificateSpecCertificate)(ptr)
	var objs []CertificateSpecCertificate
	if obj != nil {
		objs = []CertificateSpecCertificate{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificate{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CertificateSpecCertificateCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CertificateSpecCertificate)(ptr) = CertificateSpecCertificate{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CertificateSpecCertificate

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificate{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CertificateSpecCertificate)(ptr) = objs[0]
			} else {
				*(*CertificateSpecCertificate)(ptr) = CertificateSpecCertificate{}
			}
		} else {
			*(*CertificateSpecCertificate)(ptr) = CertificateSpecCertificate{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CertificateSpecCertificate

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificate{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CertificateSpecCertificate)(ptr) = obj
		} else {
			*(*CertificateSpecCertificate)(ptr) = CertificateSpecCertificate{}
		}
	default:
		iter.ReportError("decode CertificateSpecCertificate", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CertificateSpecCertificatePolicyCodec struct {
}

func (CertificateSpecCertificatePolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CertificateSpecCertificatePolicy)(ptr) == nil
}

func (CertificateSpecCertificatePolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CertificateSpecCertificatePolicy)(ptr)
	var objs []CertificateSpecCertificatePolicy
	if obj != nil {
		objs = []CertificateSpecCertificatePolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CertificateSpecCertificatePolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CertificateSpecCertificatePolicy)(ptr) = CertificateSpecCertificatePolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CertificateSpecCertificatePolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CertificateSpecCertificatePolicy)(ptr) = objs[0]
			} else {
				*(*CertificateSpecCertificatePolicy)(ptr) = CertificateSpecCertificatePolicy{}
			}
		} else {
			*(*CertificateSpecCertificatePolicy)(ptr) = CertificateSpecCertificatePolicy{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CertificateSpecCertificatePolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CertificateSpecCertificatePolicy)(ptr) = obj
		} else {
			*(*CertificateSpecCertificatePolicy)(ptr) = CertificateSpecCertificatePolicy{}
		}
	default:
		iter.ReportError("decode CertificateSpecCertificatePolicy", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CertificateSpecCertificatePolicyIssuerParametersCodec struct {
}

func (CertificateSpecCertificatePolicyIssuerParametersCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CertificateSpecCertificatePolicyIssuerParameters)(ptr) == nil
}

func (CertificateSpecCertificatePolicyIssuerParametersCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CertificateSpecCertificatePolicyIssuerParameters)(ptr)
	var objs []CertificateSpecCertificatePolicyIssuerParameters
	if obj != nil {
		objs = []CertificateSpecCertificatePolicyIssuerParameters{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyIssuerParameters{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CertificateSpecCertificatePolicyIssuerParametersCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CertificateSpecCertificatePolicyIssuerParameters)(ptr) = CertificateSpecCertificatePolicyIssuerParameters{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CertificateSpecCertificatePolicyIssuerParameters

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyIssuerParameters{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CertificateSpecCertificatePolicyIssuerParameters)(ptr) = objs[0]
			} else {
				*(*CertificateSpecCertificatePolicyIssuerParameters)(ptr) = CertificateSpecCertificatePolicyIssuerParameters{}
			}
		} else {
			*(*CertificateSpecCertificatePolicyIssuerParameters)(ptr) = CertificateSpecCertificatePolicyIssuerParameters{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CertificateSpecCertificatePolicyIssuerParameters

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyIssuerParameters{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CertificateSpecCertificatePolicyIssuerParameters)(ptr) = obj
		} else {
			*(*CertificateSpecCertificatePolicyIssuerParameters)(ptr) = CertificateSpecCertificatePolicyIssuerParameters{}
		}
	default:
		iter.ReportError("decode CertificateSpecCertificatePolicyIssuerParameters", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CertificateSpecCertificatePolicyKeyPropertiesCodec struct {
}

func (CertificateSpecCertificatePolicyKeyPropertiesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CertificateSpecCertificatePolicyKeyProperties)(ptr) == nil
}

func (CertificateSpecCertificatePolicyKeyPropertiesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CertificateSpecCertificatePolicyKeyProperties)(ptr)
	var objs []CertificateSpecCertificatePolicyKeyProperties
	if obj != nil {
		objs = []CertificateSpecCertificatePolicyKeyProperties{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyKeyProperties{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CertificateSpecCertificatePolicyKeyPropertiesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CertificateSpecCertificatePolicyKeyProperties)(ptr) = CertificateSpecCertificatePolicyKeyProperties{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CertificateSpecCertificatePolicyKeyProperties

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyKeyProperties{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CertificateSpecCertificatePolicyKeyProperties)(ptr) = objs[0]
			} else {
				*(*CertificateSpecCertificatePolicyKeyProperties)(ptr) = CertificateSpecCertificatePolicyKeyProperties{}
			}
		} else {
			*(*CertificateSpecCertificatePolicyKeyProperties)(ptr) = CertificateSpecCertificatePolicyKeyProperties{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CertificateSpecCertificatePolicyKeyProperties

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyKeyProperties{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CertificateSpecCertificatePolicyKeyProperties)(ptr) = obj
		} else {
			*(*CertificateSpecCertificatePolicyKeyProperties)(ptr) = CertificateSpecCertificatePolicyKeyProperties{}
		}
	default:
		iter.ReportError("decode CertificateSpecCertificatePolicyKeyProperties", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CertificateSpecCertificatePolicyLifetimeActionActionCodec struct {
}

func (CertificateSpecCertificatePolicyLifetimeActionActionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CertificateSpecCertificatePolicyLifetimeActionAction)(ptr) == nil
}

func (CertificateSpecCertificatePolicyLifetimeActionActionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CertificateSpecCertificatePolicyLifetimeActionAction)(ptr)
	var objs []CertificateSpecCertificatePolicyLifetimeActionAction
	if obj != nil {
		objs = []CertificateSpecCertificatePolicyLifetimeActionAction{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyLifetimeActionAction{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CertificateSpecCertificatePolicyLifetimeActionActionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CertificateSpecCertificatePolicyLifetimeActionAction)(ptr) = CertificateSpecCertificatePolicyLifetimeActionAction{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CertificateSpecCertificatePolicyLifetimeActionAction

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyLifetimeActionAction{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CertificateSpecCertificatePolicyLifetimeActionAction)(ptr) = objs[0]
			} else {
				*(*CertificateSpecCertificatePolicyLifetimeActionAction)(ptr) = CertificateSpecCertificatePolicyLifetimeActionAction{}
			}
		} else {
			*(*CertificateSpecCertificatePolicyLifetimeActionAction)(ptr) = CertificateSpecCertificatePolicyLifetimeActionAction{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CertificateSpecCertificatePolicyLifetimeActionAction

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyLifetimeActionAction{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CertificateSpecCertificatePolicyLifetimeActionAction)(ptr) = obj
		} else {
			*(*CertificateSpecCertificatePolicyLifetimeActionAction)(ptr) = CertificateSpecCertificatePolicyLifetimeActionAction{}
		}
	default:
		iter.ReportError("decode CertificateSpecCertificatePolicyLifetimeActionAction", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CertificateSpecCertificatePolicyLifetimeActionTriggerCodec struct {
}

func (CertificateSpecCertificatePolicyLifetimeActionTriggerCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CertificateSpecCertificatePolicyLifetimeActionTrigger)(ptr) == nil
}

func (CertificateSpecCertificatePolicyLifetimeActionTriggerCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CertificateSpecCertificatePolicyLifetimeActionTrigger)(ptr)
	var objs []CertificateSpecCertificatePolicyLifetimeActionTrigger
	if obj != nil {
		objs = []CertificateSpecCertificatePolicyLifetimeActionTrigger{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyLifetimeActionTrigger{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CertificateSpecCertificatePolicyLifetimeActionTriggerCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CertificateSpecCertificatePolicyLifetimeActionTrigger)(ptr) = CertificateSpecCertificatePolicyLifetimeActionTrigger{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CertificateSpecCertificatePolicyLifetimeActionTrigger

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyLifetimeActionTrigger{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CertificateSpecCertificatePolicyLifetimeActionTrigger)(ptr) = objs[0]
			} else {
				*(*CertificateSpecCertificatePolicyLifetimeActionTrigger)(ptr) = CertificateSpecCertificatePolicyLifetimeActionTrigger{}
			}
		} else {
			*(*CertificateSpecCertificatePolicyLifetimeActionTrigger)(ptr) = CertificateSpecCertificatePolicyLifetimeActionTrigger{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CertificateSpecCertificatePolicyLifetimeActionTrigger

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyLifetimeActionTrigger{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CertificateSpecCertificatePolicyLifetimeActionTrigger)(ptr) = obj
		} else {
			*(*CertificateSpecCertificatePolicyLifetimeActionTrigger)(ptr) = CertificateSpecCertificatePolicyLifetimeActionTrigger{}
		}
	default:
		iter.ReportError("decode CertificateSpecCertificatePolicyLifetimeActionTrigger", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CertificateSpecCertificatePolicySecretPropertiesCodec struct {
}

func (CertificateSpecCertificatePolicySecretPropertiesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CertificateSpecCertificatePolicySecretProperties)(ptr) == nil
}

func (CertificateSpecCertificatePolicySecretPropertiesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CertificateSpecCertificatePolicySecretProperties)(ptr)
	var objs []CertificateSpecCertificatePolicySecretProperties
	if obj != nil {
		objs = []CertificateSpecCertificatePolicySecretProperties{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicySecretProperties{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CertificateSpecCertificatePolicySecretPropertiesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CertificateSpecCertificatePolicySecretProperties)(ptr) = CertificateSpecCertificatePolicySecretProperties{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CertificateSpecCertificatePolicySecretProperties

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicySecretProperties{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CertificateSpecCertificatePolicySecretProperties)(ptr) = objs[0]
			} else {
				*(*CertificateSpecCertificatePolicySecretProperties)(ptr) = CertificateSpecCertificatePolicySecretProperties{}
			}
		} else {
			*(*CertificateSpecCertificatePolicySecretProperties)(ptr) = CertificateSpecCertificatePolicySecretProperties{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CertificateSpecCertificatePolicySecretProperties

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicySecretProperties{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CertificateSpecCertificatePolicySecretProperties)(ptr) = obj
		} else {
			*(*CertificateSpecCertificatePolicySecretProperties)(ptr) = CertificateSpecCertificatePolicySecretProperties{}
		}
	default:
		iter.ReportError("decode CertificateSpecCertificatePolicySecretProperties", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CertificateSpecCertificatePolicyX509CertificatePropertiesCodec struct {
}

func (CertificateSpecCertificatePolicyX509CertificatePropertiesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CertificateSpecCertificatePolicyX509CertificateProperties)(ptr) == nil
}

func (CertificateSpecCertificatePolicyX509CertificatePropertiesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CertificateSpecCertificatePolicyX509CertificateProperties)(ptr)
	var objs []CertificateSpecCertificatePolicyX509CertificateProperties
	if obj != nil {
		objs = []CertificateSpecCertificatePolicyX509CertificateProperties{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyX509CertificateProperties{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CertificateSpecCertificatePolicyX509CertificatePropertiesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CertificateSpecCertificatePolicyX509CertificateProperties)(ptr) = CertificateSpecCertificatePolicyX509CertificateProperties{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CertificateSpecCertificatePolicyX509CertificateProperties

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyX509CertificateProperties{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CertificateSpecCertificatePolicyX509CertificateProperties)(ptr) = objs[0]
			} else {
				*(*CertificateSpecCertificatePolicyX509CertificateProperties)(ptr) = CertificateSpecCertificatePolicyX509CertificateProperties{}
			}
		} else {
			*(*CertificateSpecCertificatePolicyX509CertificateProperties)(ptr) = CertificateSpecCertificatePolicyX509CertificateProperties{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CertificateSpecCertificatePolicyX509CertificateProperties

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyX509CertificateProperties{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CertificateSpecCertificatePolicyX509CertificateProperties)(ptr) = obj
		} else {
			*(*CertificateSpecCertificatePolicyX509CertificateProperties)(ptr) = CertificateSpecCertificatePolicyX509CertificateProperties{}
		}
	default:
		iter.ReportError("decode CertificateSpecCertificatePolicyX509CertificateProperties", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesCodec struct {
}

func (CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames)(ptr) == nil
}

func (CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames)(ptr)
	var objs []CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames
	if obj != nil {
		objs = []CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames)(ptr) = CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames)(ptr) = objs[0]
			} else {
				*(*CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames)(ptr) = CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames{}
			}
		} else {
			*(*CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames)(ptr) = CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames)(ptr) = obj
		} else {
			*(*CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames)(ptr) = CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames{}
		}
	default:
		iter.ReportError("decode CertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames", "unexpected JSON type")
	}
}
