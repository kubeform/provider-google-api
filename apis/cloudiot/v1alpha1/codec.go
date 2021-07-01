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
		jsoniter.MustGetKind(reflect2.TypeOf(DeviceSpecCredentialsPublicKey{}).Type1()): DeviceSpecCredentialsPublicKeyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(DeviceSpecGatewayConfig{}).Type1()):        DeviceSpecGatewayConfigCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(DeviceSpecCredentialsPublicKey{}).Type1()): DeviceSpecCredentialsPublicKeyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(DeviceSpecGatewayConfig{}).Type1()):        DeviceSpecGatewayConfigCodec{},
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
type DeviceSpecCredentialsPublicKeyCodec struct {
}

func (DeviceSpecCredentialsPublicKeyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*DeviceSpecCredentialsPublicKey)(ptr) == nil
}

func (DeviceSpecCredentialsPublicKeyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*DeviceSpecCredentialsPublicKey)(ptr)
	var objs []DeviceSpecCredentialsPublicKey
	if obj != nil {
		objs = []DeviceSpecCredentialsPublicKey{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DeviceSpecCredentialsPublicKey{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (DeviceSpecCredentialsPublicKeyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*DeviceSpecCredentialsPublicKey)(ptr) = DeviceSpecCredentialsPublicKey{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []DeviceSpecCredentialsPublicKey

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DeviceSpecCredentialsPublicKey{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*DeviceSpecCredentialsPublicKey)(ptr) = objs[0]
			} else {
				*(*DeviceSpecCredentialsPublicKey)(ptr) = DeviceSpecCredentialsPublicKey{}
			}
		} else {
			*(*DeviceSpecCredentialsPublicKey)(ptr) = DeviceSpecCredentialsPublicKey{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj DeviceSpecCredentialsPublicKey

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DeviceSpecCredentialsPublicKey{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*DeviceSpecCredentialsPublicKey)(ptr) = obj
		} else {
			*(*DeviceSpecCredentialsPublicKey)(ptr) = DeviceSpecCredentialsPublicKey{}
		}
	default:
		iter.ReportError("decode DeviceSpecCredentialsPublicKey", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type DeviceSpecGatewayConfigCodec struct {
}

func (DeviceSpecGatewayConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*DeviceSpecGatewayConfig)(ptr) == nil
}

func (DeviceSpecGatewayConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*DeviceSpecGatewayConfig)(ptr)
	var objs []DeviceSpecGatewayConfig
	if obj != nil {
		objs = []DeviceSpecGatewayConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DeviceSpecGatewayConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (DeviceSpecGatewayConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*DeviceSpecGatewayConfig)(ptr) = DeviceSpecGatewayConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []DeviceSpecGatewayConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DeviceSpecGatewayConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*DeviceSpecGatewayConfig)(ptr) = objs[0]
			} else {
				*(*DeviceSpecGatewayConfig)(ptr) = DeviceSpecGatewayConfig{}
			}
		} else {
			*(*DeviceSpecGatewayConfig)(ptr) = DeviceSpecGatewayConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj DeviceSpecGatewayConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DeviceSpecGatewayConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*DeviceSpecGatewayConfig)(ptr) = obj
		} else {
			*(*DeviceSpecGatewayConfig)(ptr) = DeviceSpecGatewayConfig{}
		}
	default:
		iter.ReportError("decode DeviceSpecGatewayConfig", "unexpected JSON type")
	}
}
