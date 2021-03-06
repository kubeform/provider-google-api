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
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMaintenancePolicy{}).Type1()):                                 InstanceSpecMaintenancePolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime{}).Type1()): InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTimeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMaintenanceSchedule{}).Type1()):                               InstanceSpecMaintenanceScheduleCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMaintenancePolicy{}).Type1()):                                 InstanceSpecMaintenancePolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime{}).Type1()): InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTimeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMaintenanceSchedule{}).Type1()):                               InstanceSpecMaintenanceScheduleCodec{},
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
type InstanceSpecMaintenancePolicyCodec struct {
}

func (InstanceSpecMaintenancePolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InstanceSpecMaintenancePolicy)(ptr) == nil
}

func (InstanceSpecMaintenancePolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InstanceSpecMaintenancePolicy)(ptr)
	var objs []InstanceSpecMaintenancePolicy
	if obj != nil {
		objs = []InstanceSpecMaintenancePolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMaintenancePolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InstanceSpecMaintenancePolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InstanceSpecMaintenancePolicy)(ptr) = InstanceSpecMaintenancePolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InstanceSpecMaintenancePolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMaintenancePolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InstanceSpecMaintenancePolicy)(ptr) = objs[0]
			} else {
				*(*InstanceSpecMaintenancePolicy)(ptr) = InstanceSpecMaintenancePolicy{}
			}
		} else {
			*(*InstanceSpecMaintenancePolicy)(ptr) = InstanceSpecMaintenancePolicy{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj InstanceSpecMaintenancePolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMaintenancePolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*InstanceSpecMaintenancePolicy)(ptr) = obj
		} else {
			*(*InstanceSpecMaintenancePolicy)(ptr) = InstanceSpecMaintenancePolicy{}
		}
	default:
		iter.ReportError("decode InstanceSpecMaintenancePolicy", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTimeCodec struct {
}

func (InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTimeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime)(ptr) == nil
}

func (InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTimeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime)(ptr)
	var objs []InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime
	if obj != nil {
		objs = []InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTimeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime)(ptr) = InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime)(ptr) = objs[0]
			} else {
				*(*InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime)(ptr) = InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime{}
			}
		} else {
			*(*InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime)(ptr) = InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime)(ptr) = obj
		} else {
			*(*InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime)(ptr) = InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime{}
		}
	default:
		iter.ReportError("decode InstanceSpecMaintenancePolicyWeeklyMaintenanceWindowStartTime", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InstanceSpecMaintenanceScheduleCodec struct {
}

func (InstanceSpecMaintenanceScheduleCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InstanceSpecMaintenanceSchedule)(ptr) == nil
}

func (InstanceSpecMaintenanceScheduleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InstanceSpecMaintenanceSchedule)(ptr)
	var objs []InstanceSpecMaintenanceSchedule
	if obj != nil {
		objs = []InstanceSpecMaintenanceSchedule{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMaintenanceSchedule{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InstanceSpecMaintenanceScheduleCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InstanceSpecMaintenanceSchedule)(ptr) = InstanceSpecMaintenanceSchedule{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InstanceSpecMaintenanceSchedule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMaintenanceSchedule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InstanceSpecMaintenanceSchedule)(ptr) = objs[0]
			} else {
				*(*InstanceSpecMaintenanceSchedule)(ptr) = InstanceSpecMaintenanceSchedule{}
			}
		} else {
			*(*InstanceSpecMaintenanceSchedule)(ptr) = InstanceSpecMaintenanceSchedule{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj InstanceSpecMaintenanceSchedule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMaintenanceSchedule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*InstanceSpecMaintenanceSchedule)(ptr) = obj
		} else {
			*(*InstanceSpecMaintenanceSchedule)(ptr) = InstanceSpecMaintenanceSchedule{}
		}
	default:
		iter.ReportError("decode InstanceSpecMaintenanceSchedule", "unexpected JSON type")
	}
}
