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
		jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecDestination{}).Type1()):                TriggerSpecDestinationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecDestinationCloudRunService{}).Type1()): TriggerSpecDestinationCloudRunServiceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecTransport{}).Type1()):                  TriggerSpecTransportCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecTransportPubsub{}).Type1()):            TriggerSpecTransportPubsubCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecDestination{}).Type1()):                TriggerSpecDestinationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecDestinationCloudRunService{}).Type1()): TriggerSpecDestinationCloudRunServiceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecTransport{}).Type1()):                  TriggerSpecTransportCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecTransportPubsub{}).Type1()):            TriggerSpecTransportPubsubCodec{},
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
type TriggerSpecDestinationCodec struct {
}

func (TriggerSpecDestinationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*TriggerSpecDestination)(ptr) == nil
}

func (TriggerSpecDestinationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*TriggerSpecDestination)(ptr)
	var objs []TriggerSpecDestination
	if obj != nil {
		objs = []TriggerSpecDestination{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecDestination{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (TriggerSpecDestinationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*TriggerSpecDestination)(ptr) = TriggerSpecDestination{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []TriggerSpecDestination

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecDestination{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*TriggerSpecDestination)(ptr) = objs[0]
			} else {
				*(*TriggerSpecDestination)(ptr) = TriggerSpecDestination{}
			}
		} else {
			*(*TriggerSpecDestination)(ptr) = TriggerSpecDestination{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj TriggerSpecDestination

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecDestination{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*TriggerSpecDestination)(ptr) = obj
		} else {
			*(*TriggerSpecDestination)(ptr) = TriggerSpecDestination{}
		}
	default:
		iter.ReportError("decode TriggerSpecDestination", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type TriggerSpecDestinationCloudRunServiceCodec struct {
}

func (TriggerSpecDestinationCloudRunServiceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*TriggerSpecDestinationCloudRunService)(ptr) == nil
}

func (TriggerSpecDestinationCloudRunServiceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*TriggerSpecDestinationCloudRunService)(ptr)
	var objs []TriggerSpecDestinationCloudRunService
	if obj != nil {
		objs = []TriggerSpecDestinationCloudRunService{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecDestinationCloudRunService{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (TriggerSpecDestinationCloudRunServiceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*TriggerSpecDestinationCloudRunService)(ptr) = TriggerSpecDestinationCloudRunService{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []TriggerSpecDestinationCloudRunService

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecDestinationCloudRunService{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*TriggerSpecDestinationCloudRunService)(ptr) = objs[0]
			} else {
				*(*TriggerSpecDestinationCloudRunService)(ptr) = TriggerSpecDestinationCloudRunService{}
			}
		} else {
			*(*TriggerSpecDestinationCloudRunService)(ptr) = TriggerSpecDestinationCloudRunService{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj TriggerSpecDestinationCloudRunService

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecDestinationCloudRunService{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*TriggerSpecDestinationCloudRunService)(ptr) = obj
		} else {
			*(*TriggerSpecDestinationCloudRunService)(ptr) = TriggerSpecDestinationCloudRunService{}
		}
	default:
		iter.ReportError("decode TriggerSpecDestinationCloudRunService", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type TriggerSpecTransportCodec struct {
}

func (TriggerSpecTransportCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*TriggerSpecTransport)(ptr) == nil
}

func (TriggerSpecTransportCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*TriggerSpecTransport)(ptr)
	var objs []TriggerSpecTransport
	if obj != nil {
		objs = []TriggerSpecTransport{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecTransport{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (TriggerSpecTransportCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*TriggerSpecTransport)(ptr) = TriggerSpecTransport{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []TriggerSpecTransport

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecTransport{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*TriggerSpecTransport)(ptr) = objs[0]
			} else {
				*(*TriggerSpecTransport)(ptr) = TriggerSpecTransport{}
			}
		} else {
			*(*TriggerSpecTransport)(ptr) = TriggerSpecTransport{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj TriggerSpecTransport

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecTransport{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*TriggerSpecTransport)(ptr) = obj
		} else {
			*(*TriggerSpecTransport)(ptr) = TriggerSpecTransport{}
		}
	default:
		iter.ReportError("decode TriggerSpecTransport", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type TriggerSpecTransportPubsubCodec struct {
}

func (TriggerSpecTransportPubsubCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*TriggerSpecTransportPubsub)(ptr) == nil
}

func (TriggerSpecTransportPubsubCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*TriggerSpecTransportPubsub)(ptr)
	var objs []TriggerSpecTransportPubsub
	if obj != nil {
		objs = []TriggerSpecTransportPubsub{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecTransportPubsub{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (TriggerSpecTransportPubsubCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*TriggerSpecTransportPubsub)(ptr) = TriggerSpecTransportPubsub{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []TriggerSpecTransportPubsub

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecTransportPubsub{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*TriggerSpecTransportPubsub)(ptr) = objs[0]
			} else {
				*(*TriggerSpecTransportPubsub)(ptr) = TriggerSpecTransportPubsub{}
			}
		} else {
			*(*TriggerSpecTransportPubsub)(ptr) = TriggerSpecTransportPubsub{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj TriggerSpecTransportPubsub

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TriggerSpecTransportPubsub{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*TriggerSpecTransportPubsub)(ptr) = obj
		} else {
			*(*TriggerSpecTransportPubsub)(ptr) = TriggerSpecTransportPubsub{}
		}
	default:
		iter.ReportError("decode TriggerSpecTransportPubsub", "unexpected JSON type")
	}
}
