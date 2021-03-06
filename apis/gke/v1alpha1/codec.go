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
		jsoniter.MustGetKind(reflect2.TypeOf(HubMembershipSpecAuthority{}).Type1()):          HubMembershipSpecAuthorityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(HubMembershipSpecEndpoint{}).Type1()):           HubMembershipSpecEndpointCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(HubMembershipSpecEndpointGkeCluster{}).Type1()): HubMembershipSpecEndpointGkeClusterCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(HubMembershipSpecAuthority{}).Type1()):          HubMembershipSpecAuthorityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(HubMembershipSpecEndpoint{}).Type1()):           HubMembershipSpecEndpointCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(HubMembershipSpecEndpointGkeCluster{}).Type1()): HubMembershipSpecEndpointGkeClusterCodec{},
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
type HubMembershipSpecAuthorityCodec struct {
}

func (HubMembershipSpecAuthorityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*HubMembershipSpecAuthority)(ptr) == nil
}

func (HubMembershipSpecAuthorityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*HubMembershipSpecAuthority)(ptr)
	var objs []HubMembershipSpecAuthority
	if obj != nil {
		objs = []HubMembershipSpecAuthority{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(HubMembershipSpecAuthority{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (HubMembershipSpecAuthorityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*HubMembershipSpecAuthority)(ptr) = HubMembershipSpecAuthority{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []HubMembershipSpecAuthority

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(HubMembershipSpecAuthority{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*HubMembershipSpecAuthority)(ptr) = objs[0]
			} else {
				*(*HubMembershipSpecAuthority)(ptr) = HubMembershipSpecAuthority{}
			}
		} else {
			*(*HubMembershipSpecAuthority)(ptr) = HubMembershipSpecAuthority{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj HubMembershipSpecAuthority

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(HubMembershipSpecAuthority{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*HubMembershipSpecAuthority)(ptr) = obj
		} else {
			*(*HubMembershipSpecAuthority)(ptr) = HubMembershipSpecAuthority{}
		}
	default:
		iter.ReportError("decode HubMembershipSpecAuthority", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type HubMembershipSpecEndpointCodec struct {
}

func (HubMembershipSpecEndpointCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*HubMembershipSpecEndpoint)(ptr) == nil
}

func (HubMembershipSpecEndpointCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*HubMembershipSpecEndpoint)(ptr)
	var objs []HubMembershipSpecEndpoint
	if obj != nil {
		objs = []HubMembershipSpecEndpoint{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(HubMembershipSpecEndpoint{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (HubMembershipSpecEndpointCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*HubMembershipSpecEndpoint)(ptr) = HubMembershipSpecEndpoint{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []HubMembershipSpecEndpoint

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(HubMembershipSpecEndpoint{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*HubMembershipSpecEndpoint)(ptr) = objs[0]
			} else {
				*(*HubMembershipSpecEndpoint)(ptr) = HubMembershipSpecEndpoint{}
			}
		} else {
			*(*HubMembershipSpecEndpoint)(ptr) = HubMembershipSpecEndpoint{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj HubMembershipSpecEndpoint

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(HubMembershipSpecEndpoint{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*HubMembershipSpecEndpoint)(ptr) = obj
		} else {
			*(*HubMembershipSpecEndpoint)(ptr) = HubMembershipSpecEndpoint{}
		}
	default:
		iter.ReportError("decode HubMembershipSpecEndpoint", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type HubMembershipSpecEndpointGkeClusterCodec struct {
}

func (HubMembershipSpecEndpointGkeClusterCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*HubMembershipSpecEndpointGkeCluster)(ptr) == nil
}

func (HubMembershipSpecEndpointGkeClusterCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*HubMembershipSpecEndpointGkeCluster)(ptr)
	var objs []HubMembershipSpecEndpointGkeCluster
	if obj != nil {
		objs = []HubMembershipSpecEndpointGkeCluster{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(HubMembershipSpecEndpointGkeCluster{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (HubMembershipSpecEndpointGkeClusterCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*HubMembershipSpecEndpointGkeCluster)(ptr) = HubMembershipSpecEndpointGkeCluster{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []HubMembershipSpecEndpointGkeCluster

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(HubMembershipSpecEndpointGkeCluster{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*HubMembershipSpecEndpointGkeCluster)(ptr) = objs[0]
			} else {
				*(*HubMembershipSpecEndpointGkeCluster)(ptr) = HubMembershipSpecEndpointGkeCluster{}
			}
		} else {
			*(*HubMembershipSpecEndpointGkeCluster)(ptr) = HubMembershipSpecEndpointGkeCluster{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj HubMembershipSpecEndpointGkeCluster

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(HubMembershipSpecEndpointGkeCluster{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*HubMembershipSpecEndpointGkeCluster)(ptr) = obj
		} else {
			*(*HubMembershipSpecEndpointGkeCluster)(ptr) = HubMembershipSpecEndpointGkeCluster{}
		}
	default:
		iter.ReportError("decode HubMembershipSpecEndpointGkeCluster", "unexpected JSON type")
	}
}
