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
		jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecDestination{}).Type1()):                        ConnectionMonitorSpecDestinationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecEndpointFilter{}).Type1()):                     ConnectionMonitorSpecEndpointFilterCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecSource{}).Type1()):                             ConnectionMonitorSpecSourceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationHttpConfiguration{}).Type1()): ConnectionMonitorSpecTestConfigurationHttpConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationIcmpConfiguration{}).Type1()): ConnectionMonitorSpecTestConfigurationIcmpConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationSuccessThreshold{}).Type1()):  ConnectionMonitorSpecTestConfigurationSuccessThresholdCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationTcpConfiguration{}).Type1()):  ConnectionMonitorSpecTestConfigurationTcpConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PacketCaptureSpecStorageLocation{}).Type1()):                        PacketCaptureSpecStorageLocationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProfileSpecContainerNetworkInterface{}).Type1()):                    ProfileSpecContainerNetworkInterfaceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WatcherFlowLogSpecRetentionPolicy{}).Type1()):                       WatcherFlowLogSpecRetentionPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WatcherFlowLogSpecTrafficAnalytics{}).Type1()):                      WatcherFlowLogSpecTrafficAnalyticsCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecDestination{}).Type1()):                        ConnectionMonitorSpecDestinationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecEndpointFilter{}).Type1()):                     ConnectionMonitorSpecEndpointFilterCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecSource{}).Type1()):                             ConnectionMonitorSpecSourceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationHttpConfiguration{}).Type1()): ConnectionMonitorSpecTestConfigurationHttpConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationIcmpConfiguration{}).Type1()): ConnectionMonitorSpecTestConfigurationIcmpConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationSuccessThreshold{}).Type1()):  ConnectionMonitorSpecTestConfigurationSuccessThresholdCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationTcpConfiguration{}).Type1()):  ConnectionMonitorSpecTestConfigurationTcpConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PacketCaptureSpecStorageLocation{}).Type1()):                        PacketCaptureSpecStorageLocationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProfileSpecContainerNetworkInterface{}).Type1()):                    ProfileSpecContainerNetworkInterfaceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WatcherFlowLogSpecRetentionPolicy{}).Type1()):                       WatcherFlowLogSpecRetentionPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WatcherFlowLogSpecTrafficAnalytics{}).Type1()):                      WatcherFlowLogSpecTrafficAnalyticsCodec{},
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
type ConnectionMonitorSpecDestinationCodec struct {
}

func (ConnectionMonitorSpecDestinationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ConnectionMonitorSpecDestination)(ptr) == nil
}

func (ConnectionMonitorSpecDestinationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ConnectionMonitorSpecDestination)(ptr)
	var objs []ConnectionMonitorSpecDestination
	if obj != nil {
		objs = []ConnectionMonitorSpecDestination{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecDestination{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ConnectionMonitorSpecDestinationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ConnectionMonitorSpecDestination)(ptr) = ConnectionMonitorSpecDestination{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ConnectionMonitorSpecDestination

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecDestination{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ConnectionMonitorSpecDestination)(ptr) = objs[0]
			} else {
				*(*ConnectionMonitorSpecDestination)(ptr) = ConnectionMonitorSpecDestination{}
			}
		} else {
			*(*ConnectionMonitorSpecDestination)(ptr) = ConnectionMonitorSpecDestination{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ConnectionMonitorSpecDestination

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecDestination{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ConnectionMonitorSpecDestination)(ptr) = obj
		} else {
			*(*ConnectionMonitorSpecDestination)(ptr) = ConnectionMonitorSpecDestination{}
		}
	default:
		iter.ReportError("decode ConnectionMonitorSpecDestination", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ConnectionMonitorSpecEndpointFilterCodec struct {
}

func (ConnectionMonitorSpecEndpointFilterCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ConnectionMonitorSpecEndpointFilter)(ptr) == nil
}

func (ConnectionMonitorSpecEndpointFilterCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ConnectionMonitorSpecEndpointFilter)(ptr)
	var objs []ConnectionMonitorSpecEndpointFilter
	if obj != nil {
		objs = []ConnectionMonitorSpecEndpointFilter{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecEndpointFilter{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ConnectionMonitorSpecEndpointFilterCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ConnectionMonitorSpecEndpointFilter)(ptr) = ConnectionMonitorSpecEndpointFilter{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ConnectionMonitorSpecEndpointFilter

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecEndpointFilter{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ConnectionMonitorSpecEndpointFilter)(ptr) = objs[0]
			} else {
				*(*ConnectionMonitorSpecEndpointFilter)(ptr) = ConnectionMonitorSpecEndpointFilter{}
			}
		} else {
			*(*ConnectionMonitorSpecEndpointFilter)(ptr) = ConnectionMonitorSpecEndpointFilter{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ConnectionMonitorSpecEndpointFilter

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecEndpointFilter{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ConnectionMonitorSpecEndpointFilter)(ptr) = obj
		} else {
			*(*ConnectionMonitorSpecEndpointFilter)(ptr) = ConnectionMonitorSpecEndpointFilter{}
		}
	default:
		iter.ReportError("decode ConnectionMonitorSpecEndpointFilter", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ConnectionMonitorSpecSourceCodec struct {
}

func (ConnectionMonitorSpecSourceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ConnectionMonitorSpecSource)(ptr) == nil
}

func (ConnectionMonitorSpecSourceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ConnectionMonitorSpecSource)(ptr)
	var objs []ConnectionMonitorSpecSource
	if obj != nil {
		objs = []ConnectionMonitorSpecSource{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecSource{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ConnectionMonitorSpecSourceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ConnectionMonitorSpecSource)(ptr) = ConnectionMonitorSpecSource{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ConnectionMonitorSpecSource

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecSource{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ConnectionMonitorSpecSource)(ptr) = objs[0]
			} else {
				*(*ConnectionMonitorSpecSource)(ptr) = ConnectionMonitorSpecSource{}
			}
		} else {
			*(*ConnectionMonitorSpecSource)(ptr) = ConnectionMonitorSpecSource{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ConnectionMonitorSpecSource

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecSource{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ConnectionMonitorSpecSource)(ptr) = obj
		} else {
			*(*ConnectionMonitorSpecSource)(ptr) = ConnectionMonitorSpecSource{}
		}
	default:
		iter.ReportError("decode ConnectionMonitorSpecSource", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ConnectionMonitorSpecTestConfigurationHttpConfigurationCodec struct {
}

func (ConnectionMonitorSpecTestConfigurationHttpConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ConnectionMonitorSpecTestConfigurationHttpConfiguration)(ptr) == nil
}

func (ConnectionMonitorSpecTestConfigurationHttpConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ConnectionMonitorSpecTestConfigurationHttpConfiguration)(ptr)
	var objs []ConnectionMonitorSpecTestConfigurationHttpConfiguration
	if obj != nil {
		objs = []ConnectionMonitorSpecTestConfigurationHttpConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationHttpConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ConnectionMonitorSpecTestConfigurationHttpConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ConnectionMonitorSpecTestConfigurationHttpConfiguration)(ptr) = ConnectionMonitorSpecTestConfigurationHttpConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ConnectionMonitorSpecTestConfigurationHttpConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationHttpConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ConnectionMonitorSpecTestConfigurationHttpConfiguration)(ptr) = objs[0]
			} else {
				*(*ConnectionMonitorSpecTestConfigurationHttpConfiguration)(ptr) = ConnectionMonitorSpecTestConfigurationHttpConfiguration{}
			}
		} else {
			*(*ConnectionMonitorSpecTestConfigurationHttpConfiguration)(ptr) = ConnectionMonitorSpecTestConfigurationHttpConfiguration{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ConnectionMonitorSpecTestConfigurationHttpConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationHttpConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ConnectionMonitorSpecTestConfigurationHttpConfiguration)(ptr) = obj
		} else {
			*(*ConnectionMonitorSpecTestConfigurationHttpConfiguration)(ptr) = ConnectionMonitorSpecTestConfigurationHttpConfiguration{}
		}
	default:
		iter.ReportError("decode ConnectionMonitorSpecTestConfigurationHttpConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ConnectionMonitorSpecTestConfigurationIcmpConfigurationCodec struct {
}

func (ConnectionMonitorSpecTestConfigurationIcmpConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ConnectionMonitorSpecTestConfigurationIcmpConfiguration)(ptr) == nil
}

func (ConnectionMonitorSpecTestConfigurationIcmpConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ConnectionMonitorSpecTestConfigurationIcmpConfiguration)(ptr)
	var objs []ConnectionMonitorSpecTestConfigurationIcmpConfiguration
	if obj != nil {
		objs = []ConnectionMonitorSpecTestConfigurationIcmpConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationIcmpConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ConnectionMonitorSpecTestConfigurationIcmpConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ConnectionMonitorSpecTestConfigurationIcmpConfiguration)(ptr) = ConnectionMonitorSpecTestConfigurationIcmpConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ConnectionMonitorSpecTestConfigurationIcmpConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationIcmpConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ConnectionMonitorSpecTestConfigurationIcmpConfiguration)(ptr) = objs[0]
			} else {
				*(*ConnectionMonitorSpecTestConfigurationIcmpConfiguration)(ptr) = ConnectionMonitorSpecTestConfigurationIcmpConfiguration{}
			}
		} else {
			*(*ConnectionMonitorSpecTestConfigurationIcmpConfiguration)(ptr) = ConnectionMonitorSpecTestConfigurationIcmpConfiguration{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ConnectionMonitorSpecTestConfigurationIcmpConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationIcmpConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ConnectionMonitorSpecTestConfigurationIcmpConfiguration)(ptr) = obj
		} else {
			*(*ConnectionMonitorSpecTestConfigurationIcmpConfiguration)(ptr) = ConnectionMonitorSpecTestConfigurationIcmpConfiguration{}
		}
	default:
		iter.ReportError("decode ConnectionMonitorSpecTestConfigurationIcmpConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ConnectionMonitorSpecTestConfigurationSuccessThresholdCodec struct {
}

func (ConnectionMonitorSpecTestConfigurationSuccessThresholdCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ConnectionMonitorSpecTestConfigurationSuccessThreshold)(ptr) == nil
}

func (ConnectionMonitorSpecTestConfigurationSuccessThresholdCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ConnectionMonitorSpecTestConfigurationSuccessThreshold)(ptr)
	var objs []ConnectionMonitorSpecTestConfigurationSuccessThreshold
	if obj != nil {
		objs = []ConnectionMonitorSpecTestConfigurationSuccessThreshold{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationSuccessThreshold{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ConnectionMonitorSpecTestConfigurationSuccessThresholdCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ConnectionMonitorSpecTestConfigurationSuccessThreshold)(ptr) = ConnectionMonitorSpecTestConfigurationSuccessThreshold{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ConnectionMonitorSpecTestConfigurationSuccessThreshold

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationSuccessThreshold{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ConnectionMonitorSpecTestConfigurationSuccessThreshold)(ptr) = objs[0]
			} else {
				*(*ConnectionMonitorSpecTestConfigurationSuccessThreshold)(ptr) = ConnectionMonitorSpecTestConfigurationSuccessThreshold{}
			}
		} else {
			*(*ConnectionMonitorSpecTestConfigurationSuccessThreshold)(ptr) = ConnectionMonitorSpecTestConfigurationSuccessThreshold{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ConnectionMonitorSpecTestConfigurationSuccessThreshold

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationSuccessThreshold{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ConnectionMonitorSpecTestConfigurationSuccessThreshold)(ptr) = obj
		} else {
			*(*ConnectionMonitorSpecTestConfigurationSuccessThreshold)(ptr) = ConnectionMonitorSpecTestConfigurationSuccessThreshold{}
		}
	default:
		iter.ReportError("decode ConnectionMonitorSpecTestConfigurationSuccessThreshold", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ConnectionMonitorSpecTestConfigurationTcpConfigurationCodec struct {
}

func (ConnectionMonitorSpecTestConfigurationTcpConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ConnectionMonitorSpecTestConfigurationTcpConfiguration)(ptr) == nil
}

func (ConnectionMonitorSpecTestConfigurationTcpConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ConnectionMonitorSpecTestConfigurationTcpConfiguration)(ptr)
	var objs []ConnectionMonitorSpecTestConfigurationTcpConfiguration
	if obj != nil {
		objs = []ConnectionMonitorSpecTestConfigurationTcpConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationTcpConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ConnectionMonitorSpecTestConfigurationTcpConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ConnectionMonitorSpecTestConfigurationTcpConfiguration)(ptr) = ConnectionMonitorSpecTestConfigurationTcpConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ConnectionMonitorSpecTestConfigurationTcpConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationTcpConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ConnectionMonitorSpecTestConfigurationTcpConfiguration)(ptr) = objs[0]
			} else {
				*(*ConnectionMonitorSpecTestConfigurationTcpConfiguration)(ptr) = ConnectionMonitorSpecTestConfigurationTcpConfiguration{}
			}
		} else {
			*(*ConnectionMonitorSpecTestConfigurationTcpConfiguration)(ptr) = ConnectionMonitorSpecTestConfigurationTcpConfiguration{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ConnectionMonitorSpecTestConfigurationTcpConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ConnectionMonitorSpecTestConfigurationTcpConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ConnectionMonitorSpecTestConfigurationTcpConfiguration)(ptr) = obj
		} else {
			*(*ConnectionMonitorSpecTestConfigurationTcpConfiguration)(ptr) = ConnectionMonitorSpecTestConfigurationTcpConfiguration{}
		}
	default:
		iter.ReportError("decode ConnectionMonitorSpecTestConfigurationTcpConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PacketCaptureSpecStorageLocationCodec struct {
}

func (PacketCaptureSpecStorageLocationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PacketCaptureSpecStorageLocation)(ptr) == nil
}

func (PacketCaptureSpecStorageLocationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PacketCaptureSpecStorageLocation)(ptr)
	var objs []PacketCaptureSpecStorageLocation
	if obj != nil {
		objs = []PacketCaptureSpecStorageLocation{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PacketCaptureSpecStorageLocation{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PacketCaptureSpecStorageLocationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PacketCaptureSpecStorageLocation)(ptr) = PacketCaptureSpecStorageLocation{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PacketCaptureSpecStorageLocation

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PacketCaptureSpecStorageLocation{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PacketCaptureSpecStorageLocation)(ptr) = objs[0]
			} else {
				*(*PacketCaptureSpecStorageLocation)(ptr) = PacketCaptureSpecStorageLocation{}
			}
		} else {
			*(*PacketCaptureSpecStorageLocation)(ptr) = PacketCaptureSpecStorageLocation{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj PacketCaptureSpecStorageLocation

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PacketCaptureSpecStorageLocation{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*PacketCaptureSpecStorageLocation)(ptr) = obj
		} else {
			*(*PacketCaptureSpecStorageLocation)(ptr) = PacketCaptureSpecStorageLocation{}
		}
	default:
		iter.ReportError("decode PacketCaptureSpecStorageLocation", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ProfileSpecContainerNetworkInterfaceCodec struct {
}

func (ProfileSpecContainerNetworkInterfaceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProfileSpecContainerNetworkInterface)(ptr) == nil
}

func (ProfileSpecContainerNetworkInterfaceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProfileSpecContainerNetworkInterface)(ptr)
	var objs []ProfileSpecContainerNetworkInterface
	if obj != nil {
		objs = []ProfileSpecContainerNetworkInterface{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProfileSpecContainerNetworkInterface{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProfileSpecContainerNetworkInterfaceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProfileSpecContainerNetworkInterface)(ptr) = ProfileSpecContainerNetworkInterface{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProfileSpecContainerNetworkInterface

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProfileSpecContainerNetworkInterface{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProfileSpecContainerNetworkInterface)(ptr) = objs[0]
			} else {
				*(*ProfileSpecContainerNetworkInterface)(ptr) = ProfileSpecContainerNetworkInterface{}
			}
		} else {
			*(*ProfileSpecContainerNetworkInterface)(ptr) = ProfileSpecContainerNetworkInterface{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ProfileSpecContainerNetworkInterface

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProfileSpecContainerNetworkInterface{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ProfileSpecContainerNetworkInterface)(ptr) = obj
		} else {
			*(*ProfileSpecContainerNetworkInterface)(ptr) = ProfileSpecContainerNetworkInterface{}
		}
	default:
		iter.ReportError("decode ProfileSpecContainerNetworkInterface", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type WatcherFlowLogSpecRetentionPolicyCodec struct {
}

func (WatcherFlowLogSpecRetentionPolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*WatcherFlowLogSpecRetentionPolicy)(ptr) == nil
}

func (WatcherFlowLogSpecRetentionPolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*WatcherFlowLogSpecRetentionPolicy)(ptr)
	var objs []WatcherFlowLogSpecRetentionPolicy
	if obj != nil {
		objs = []WatcherFlowLogSpecRetentionPolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WatcherFlowLogSpecRetentionPolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (WatcherFlowLogSpecRetentionPolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*WatcherFlowLogSpecRetentionPolicy)(ptr) = WatcherFlowLogSpecRetentionPolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []WatcherFlowLogSpecRetentionPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WatcherFlowLogSpecRetentionPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*WatcherFlowLogSpecRetentionPolicy)(ptr) = objs[0]
			} else {
				*(*WatcherFlowLogSpecRetentionPolicy)(ptr) = WatcherFlowLogSpecRetentionPolicy{}
			}
		} else {
			*(*WatcherFlowLogSpecRetentionPolicy)(ptr) = WatcherFlowLogSpecRetentionPolicy{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj WatcherFlowLogSpecRetentionPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WatcherFlowLogSpecRetentionPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*WatcherFlowLogSpecRetentionPolicy)(ptr) = obj
		} else {
			*(*WatcherFlowLogSpecRetentionPolicy)(ptr) = WatcherFlowLogSpecRetentionPolicy{}
		}
	default:
		iter.ReportError("decode WatcherFlowLogSpecRetentionPolicy", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type WatcherFlowLogSpecTrafficAnalyticsCodec struct {
}

func (WatcherFlowLogSpecTrafficAnalyticsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*WatcherFlowLogSpecTrafficAnalytics)(ptr) == nil
}

func (WatcherFlowLogSpecTrafficAnalyticsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*WatcherFlowLogSpecTrafficAnalytics)(ptr)
	var objs []WatcherFlowLogSpecTrafficAnalytics
	if obj != nil {
		objs = []WatcherFlowLogSpecTrafficAnalytics{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WatcherFlowLogSpecTrafficAnalytics{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (WatcherFlowLogSpecTrafficAnalyticsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*WatcherFlowLogSpecTrafficAnalytics)(ptr) = WatcherFlowLogSpecTrafficAnalytics{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []WatcherFlowLogSpecTrafficAnalytics

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WatcherFlowLogSpecTrafficAnalytics{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*WatcherFlowLogSpecTrafficAnalytics)(ptr) = objs[0]
			} else {
				*(*WatcherFlowLogSpecTrafficAnalytics)(ptr) = WatcherFlowLogSpecTrafficAnalytics{}
			}
		} else {
			*(*WatcherFlowLogSpecTrafficAnalytics)(ptr) = WatcherFlowLogSpecTrafficAnalytics{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj WatcherFlowLogSpecTrafficAnalytics

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WatcherFlowLogSpecTrafficAnalytics{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*WatcherFlowLogSpecTrafficAnalytics)(ptr) = obj
		} else {
			*(*WatcherFlowLogSpecTrafficAnalytics)(ptr) = WatcherFlowLogSpecTrafficAnalytics{}
		}
	default:
		iter.ReportError("decode WatcherFlowLogSpecTrafficAnalytics", "unexpected JSON type")
	}
}
