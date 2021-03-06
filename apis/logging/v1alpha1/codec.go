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
		jsoniter.MustGetKind(reflect2.TypeOf(BillingAccountSinkSpecBigqueryOptions{}).Type1()):     BillingAccountSinkSpecBigqueryOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FolderSinkSpecBigqueryOptions{}).Type1()):             FolderSinkSpecBigqueryOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptions{}).Type1()):                   MetricSpecBucketOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptionsExplicitBuckets{}).Type1()):    MetricSpecBucketOptionsExplicitBucketsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptionsExponentialBuckets{}).Type1()): MetricSpecBucketOptionsExponentialBucketsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptionsLinearBuckets{}).Type1()):      MetricSpecBucketOptionsLinearBucketsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecMetricDescriptor{}).Type1()):                MetricSpecMetricDescriptorCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(OrganizationSinkSpecBigqueryOptions{}).Type1()):       OrganizationSinkSpecBigqueryOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSinkSpecBigqueryOptions{}).Type1()):            ProjectSinkSpecBigqueryOptionsCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(BillingAccountSinkSpecBigqueryOptions{}).Type1()):     BillingAccountSinkSpecBigqueryOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FolderSinkSpecBigqueryOptions{}).Type1()):             FolderSinkSpecBigqueryOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptions{}).Type1()):                   MetricSpecBucketOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptionsExplicitBuckets{}).Type1()):    MetricSpecBucketOptionsExplicitBucketsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptionsExponentialBuckets{}).Type1()): MetricSpecBucketOptionsExponentialBucketsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptionsLinearBuckets{}).Type1()):      MetricSpecBucketOptionsLinearBucketsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecMetricDescriptor{}).Type1()):                MetricSpecMetricDescriptorCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(OrganizationSinkSpecBigqueryOptions{}).Type1()):       OrganizationSinkSpecBigqueryOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProjectSinkSpecBigqueryOptions{}).Type1()):            ProjectSinkSpecBigqueryOptionsCodec{},
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
type BillingAccountSinkSpecBigqueryOptionsCodec struct {
}

func (BillingAccountSinkSpecBigqueryOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BillingAccountSinkSpecBigqueryOptions)(ptr) == nil
}

func (BillingAccountSinkSpecBigqueryOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BillingAccountSinkSpecBigqueryOptions)(ptr)
	var objs []BillingAccountSinkSpecBigqueryOptions
	if obj != nil {
		objs = []BillingAccountSinkSpecBigqueryOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BillingAccountSinkSpecBigqueryOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BillingAccountSinkSpecBigqueryOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BillingAccountSinkSpecBigqueryOptions)(ptr) = BillingAccountSinkSpecBigqueryOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BillingAccountSinkSpecBigqueryOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BillingAccountSinkSpecBigqueryOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BillingAccountSinkSpecBigqueryOptions)(ptr) = objs[0]
			} else {
				*(*BillingAccountSinkSpecBigqueryOptions)(ptr) = BillingAccountSinkSpecBigqueryOptions{}
			}
		} else {
			*(*BillingAccountSinkSpecBigqueryOptions)(ptr) = BillingAccountSinkSpecBigqueryOptions{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BillingAccountSinkSpecBigqueryOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BillingAccountSinkSpecBigqueryOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BillingAccountSinkSpecBigqueryOptions)(ptr) = obj
		} else {
			*(*BillingAccountSinkSpecBigqueryOptions)(ptr) = BillingAccountSinkSpecBigqueryOptions{}
		}
	default:
		iter.ReportError("decode BillingAccountSinkSpecBigqueryOptions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type FolderSinkSpecBigqueryOptionsCodec struct {
}

func (FolderSinkSpecBigqueryOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*FolderSinkSpecBigqueryOptions)(ptr) == nil
}

func (FolderSinkSpecBigqueryOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*FolderSinkSpecBigqueryOptions)(ptr)
	var objs []FolderSinkSpecBigqueryOptions
	if obj != nil {
		objs = []FolderSinkSpecBigqueryOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FolderSinkSpecBigqueryOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (FolderSinkSpecBigqueryOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*FolderSinkSpecBigqueryOptions)(ptr) = FolderSinkSpecBigqueryOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []FolderSinkSpecBigqueryOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FolderSinkSpecBigqueryOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*FolderSinkSpecBigqueryOptions)(ptr) = objs[0]
			} else {
				*(*FolderSinkSpecBigqueryOptions)(ptr) = FolderSinkSpecBigqueryOptions{}
			}
		} else {
			*(*FolderSinkSpecBigqueryOptions)(ptr) = FolderSinkSpecBigqueryOptions{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj FolderSinkSpecBigqueryOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FolderSinkSpecBigqueryOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*FolderSinkSpecBigqueryOptions)(ptr) = obj
		} else {
			*(*FolderSinkSpecBigqueryOptions)(ptr) = FolderSinkSpecBigqueryOptions{}
		}
	default:
		iter.ReportError("decode FolderSinkSpecBigqueryOptions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MetricSpecBucketOptionsCodec struct {
}

func (MetricSpecBucketOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MetricSpecBucketOptions)(ptr) == nil
}

func (MetricSpecBucketOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MetricSpecBucketOptions)(ptr)
	var objs []MetricSpecBucketOptions
	if obj != nil {
		objs = []MetricSpecBucketOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MetricSpecBucketOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MetricSpecBucketOptions)(ptr) = MetricSpecBucketOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MetricSpecBucketOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MetricSpecBucketOptions)(ptr) = objs[0]
			} else {
				*(*MetricSpecBucketOptions)(ptr) = MetricSpecBucketOptions{}
			}
		} else {
			*(*MetricSpecBucketOptions)(ptr) = MetricSpecBucketOptions{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MetricSpecBucketOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MetricSpecBucketOptions)(ptr) = obj
		} else {
			*(*MetricSpecBucketOptions)(ptr) = MetricSpecBucketOptions{}
		}
	default:
		iter.ReportError("decode MetricSpecBucketOptions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MetricSpecBucketOptionsExplicitBucketsCodec struct {
}

func (MetricSpecBucketOptionsExplicitBucketsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MetricSpecBucketOptionsExplicitBuckets)(ptr) == nil
}

func (MetricSpecBucketOptionsExplicitBucketsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MetricSpecBucketOptionsExplicitBuckets)(ptr)
	var objs []MetricSpecBucketOptionsExplicitBuckets
	if obj != nil {
		objs = []MetricSpecBucketOptionsExplicitBuckets{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptionsExplicitBuckets{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MetricSpecBucketOptionsExplicitBucketsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MetricSpecBucketOptionsExplicitBuckets)(ptr) = MetricSpecBucketOptionsExplicitBuckets{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MetricSpecBucketOptionsExplicitBuckets

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptionsExplicitBuckets{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MetricSpecBucketOptionsExplicitBuckets)(ptr) = objs[0]
			} else {
				*(*MetricSpecBucketOptionsExplicitBuckets)(ptr) = MetricSpecBucketOptionsExplicitBuckets{}
			}
		} else {
			*(*MetricSpecBucketOptionsExplicitBuckets)(ptr) = MetricSpecBucketOptionsExplicitBuckets{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MetricSpecBucketOptionsExplicitBuckets

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptionsExplicitBuckets{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MetricSpecBucketOptionsExplicitBuckets)(ptr) = obj
		} else {
			*(*MetricSpecBucketOptionsExplicitBuckets)(ptr) = MetricSpecBucketOptionsExplicitBuckets{}
		}
	default:
		iter.ReportError("decode MetricSpecBucketOptionsExplicitBuckets", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MetricSpecBucketOptionsExponentialBucketsCodec struct {
}

func (MetricSpecBucketOptionsExponentialBucketsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MetricSpecBucketOptionsExponentialBuckets)(ptr) == nil
}

func (MetricSpecBucketOptionsExponentialBucketsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MetricSpecBucketOptionsExponentialBuckets)(ptr)
	var objs []MetricSpecBucketOptionsExponentialBuckets
	if obj != nil {
		objs = []MetricSpecBucketOptionsExponentialBuckets{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptionsExponentialBuckets{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MetricSpecBucketOptionsExponentialBucketsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MetricSpecBucketOptionsExponentialBuckets)(ptr) = MetricSpecBucketOptionsExponentialBuckets{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MetricSpecBucketOptionsExponentialBuckets

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptionsExponentialBuckets{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MetricSpecBucketOptionsExponentialBuckets)(ptr) = objs[0]
			} else {
				*(*MetricSpecBucketOptionsExponentialBuckets)(ptr) = MetricSpecBucketOptionsExponentialBuckets{}
			}
		} else {
			*(*MetricSpecBucketOptionsExponentialBuckets)(ptr) = MetricSpecBucketOptionsExponentialBuckets{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MetricSpecBucketOptionsExponentialBuckets

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptionsExponentialBuckets{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MetricSpecBucketOptionsExponentialBuckets)(ptr) = obj
		} else {
			*(*MetricSpecBucketOptionsExponentialBuckets)(ptr) = MetricSpecBucketOptionsExponentialBuckets{}
		}
	default:
		iter.ReportError("decode MetricSpecBucketOptionsExponentialBuckets", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MetricSpecBucketOptionsLinearBucketsCodec struct {
}

func (MetricSpecBucketOptionsLinearBucketsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MetricSpecBucketOptionsLinearBuckets)(ptr) == nil
}

func (MetricSpecBucketOptionsLinearBucketsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MetricSpecBucketOptionsLinearBuckets)(ptr)
	var objs []MetricSpecBucketOptionsLinearBuckets
	if obj != nil {
		objs = []MetricSpecBucketOptionsLinearBuckets{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptionsLinearBuckets{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MetricSpecBucketOptionsLinearBucketsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MetricSpecBucketOptionsLinearBuckets)(ptr) = MetricSpecBucketOptionsLinearBuckets{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MetricSpecBucketOptionsLinearBuckets

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptionsLinearBuckets{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MetricSpecBucketOptionsLinearBuckets)(ptr) = objs[0]
			} else {
				*(*MetricSpecBucketOptionsLinearBuckets)(ptr) = MetricSpecBucketOptionsLinearBuckets{}
			}
		} else {
			*(*MetricSpecBucketOptionsLinearBuckets)(ptr) = MetricSpecBucketOptionsLinearBuckets{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MetricSpecBucketOptionsLinearBuckets

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecBucketOptionsLinearBuckets{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MetricSpecBucketOptionsLinearBuckets)(ptr) = obj
		} else {
			*(*MetricSpecBucketOptionsLinearBuckets)(ptr) = MetricSpecBucketOptionsLinearBuckets{}
		}
	default:
		iter.ReportError("decode MetricSpecBucketOptionsLinearBuckets", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MetricSpecMetricDescriptorCodec struct {
}

func (MetricSpecMetricDescriptorCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MetricSpecMetricDescriptor)(ptr) == nil
}

func (MetricSpecMetricDescriptorCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MetricSpecMetricDescriptor)(ptr)
	var objs []MetricSpecMetricDescriptor
	if obj != nil {
		objs = []MetricSpecMetricDescriptor{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecMetricDescriptor{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MetricSpecMetricDescriptorCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MetricSpecMetricDescriptor)(ptr) = MetricSpecMetricDescriptor{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MetricSpecMetricDescriptor

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecMetricDescriptor{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MetricSpecMetricDescriptor)(ptr) = objs[0]
			} else {
				*(*MetricSpecMetricDescriptor)(ptr) = MetricSpecMetricDescriptor{}
			}
		} else {
			*(*MetricSpecMetricDescriptor)(ptr) = MetricSpecMetricDescriptor{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MetricSpecMetricDescriptor

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MetricSpecMetricDescriptor{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MetricSpecMetricDescriptor)(ptr) = obj
		} else {
			*(*MetricSpecMetricDescriptor)(ptr) = MetricSpecMetricDescriptor{}
		}
	default:
		iter.ReportError("decode MetricSpecMetricDescriptor", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type OrganizationSinkSpecBigqueryOptionsCodec struct {
}

func (OrganizationSinkSpecBigqueryOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*OrganizationSinkSpecBigqueryOptions)(ptr) == nil
}

func (OrganizationSinkSpecBigqueryOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*OrganizationSinkSpecBigqueryOptions)(ptr)
	var objs []OrganizationSinkSpecBigqueryOptions
	if obj != nil {
		objs = []OrganizationSinkSpecBigqueryOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(OrganizationSinkSpecBigqueryOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (OrganizationSinkSpecBigqueryOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*OrganizationSinkSpecBigqueryOptions)(ptr) = OrganizationSinkSpecBigqueryOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []OrganizationSinkSpecBigqueryOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(OrganizationSinkSpecBigqueryOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*OrganizationSinkSpecBigqueryOptions)(ptr) = objs[0]
			} else {
				*(*OrganizationSinkSpecBigqueryOptions)(ptr) = OrganizationSinkSpecBigqueryOptions{}
			}
		} else {
			*(*OrganizationSinkSpecBigqueryOptions)(ptr) = OrganizationSinkSpecBigqueryOptions{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj OrganizationSinkSpecBigqueryOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(OrganizationSinkSpecBigqueryOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*OrganizationSinkSpecBigqueryOptions)(ptr) = obj
		} else {
			*(*OrganizationSinkSpecBigqueryOptions)(ptr) = OrganizationSinkSpecBigqueryOptions{}
		}
	default:
		iter.ReportError("decode OrganizationSinkSpecBigqueryOptions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ProjectSinkSpecBigqueryOptionsCodec struct {
}

func (ProjectSinkSpecBigqueryOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProjectSinkSpecBigqueryOptions)(ptr) == nil
}

func (ProjectSinkSpecBigqueryOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProjectSinkSpecBigqueryOptions)(ptr)
	var objs []ProjectSinkSpecBigqueryOptions
	if obj != nil {
		objs = []ProjectSinkSpecBigqueryOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSinkSpecBigqueryOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProjectSinkSpecBigqueryOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProjectSinkSpecBigqueryOptions)(ptr) = ProjectSinkSpecBigqueryOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProjectSinkSpecBigqueryOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSinkSpecBigqueryOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProjectSinkSpecBigqueryOptions)(ptr) = objs[0]
			} else {
				*(*ProjectSinkSpecBigqueryOptions)(ptr) = ProjectSinkSpecBigqueryOptions{}
			}
		} else {
			*(*ProjectSinkSpecBigqueryOptions)(ptr) = ProjectSinkSpecBigqueryOptions{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ProjectSinkSpecBigqueryOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProjectSinkSpecBigqueryOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ProjectSinkSpecBigqueryOptions)(ptr) = obj
		} else {
			*(*ProjectSinkSpecBigqueryOptions)(ptr) = ProjectSinkSpecBigqueryOptions{}
		}
	default:
		iter.ReportError("decode ProjectSinkSpecBigqueryOptions", "unexpected JSON type")
	}
}
