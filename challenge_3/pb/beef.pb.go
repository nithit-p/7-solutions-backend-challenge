// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0--rc1
// source: beef.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BeefSummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *BeefSummaryRequest) Reset() {
	*x = BeefSummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beef_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeefSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeefSummaryRequest) ProtoMessage() {}

func (x *BeefSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beef_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeefSummaryRequest.ProtoReflect.Descriptor instead.
func (*BeefSummaryRequest) Descriptor() ([]byte, []int) {
	return file_beef_proto_rawDescGZIP(), []int{0}
}

func (x *BeefSummaryRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type BeefSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beef map[string]int32 `protobuf:"bytes,1,rep,name=beef,proto3" json:"beef,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *BeefSummary) Reset() {
	*x = BeefSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beef_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeefSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeefSummary) ProtoMessage() {}

func (x *BeefSummary) ProtoReflect() protoreflect.Message {
	mi := &file_beef_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeefSummary.ProtoReflect.Descriptor instead.
func (*BeefSummary) Descriptor() ([]byte, []int) {
	return file_beef_proto_rawDescGZIP(), []int{1}
}

func (x *BeefSummary) GetBeef() map[string]int32 {
	if x != nil {
		return x.Beef
	}
	return nil
}

var File_beef_proto protoreflect.FileDescriptor

var file_beef_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x65, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x65,
	0x65, 0x66, 0x22, 0x28, 0x0a, 0x12, 0x42, 0x65, 0x65, 0x66, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x77, 0x0a, 0x0b,
	0x42, 0x65, 0x65, 0x66, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x2f, 0x0a, 0x04, 0x62,
	0x65, 0x65, 0x66, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x65, 0x65, 0x66,
	0x2e, 0x42, 0x65, 0x65, 0x66, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x65, 0x65,
	0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x62, 0x65, 0x65, 0x66, 0x1a, 0x37, 0x0a, 0x09,
	0x42, 0x65, 0x65, 0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x4e, 0x0a, 0x0b, 0x42, 0x65, 0x65, 0x66, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x65, 0x65, 0x66, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x18, 0x2e, 0x62, 0x65, 0x65, 0x66, 0x2e, 0x42, 0x65,
	0x65, 0x66, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x62, 0x65, 0x65, 0x66, 0x2e, 0x42, 0x65, 0x65, 0x66, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_beef_proto_rawDescOnce sync.Once
	file_beef_proto_rawDescData = file_beef_proto_rawDesc
)

func file_beef_proto_rawDescGZIP() []byte {
	file_beef_proto_rawDescOnce.Do(func() {
		file_beef_proto_rawDescData = protoimpl.X.CompressGZIP(file_beef_proto_rawDescData)
	})
	return file_beef_proto_rawDescData
}

var file_beef_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_beef_proto_goTypes = []interface{}{
	(*BeefSummaryRequest)(nil), // 0: beef.BeefSummaryRequest
	(*BeefSummary)(nil),        // 1: beef.BeefSummary
	nil,                        // 2: beef.BeefSummary.BeefEntry
}
var file_beef_proto_depIdxs = []int32{
	2, // 0: beef.BeefSummary.beef:type_name -> beef.BeefSummary.BeefEntry
	0, // 1: beef.BeefService.GetBeefSummary:input_type -> beef.BeefSummaryRequest
	1, // 2: beef.BeefService.GetBeefSummary:output_type -> beef.BeefSummary
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_beef_proto_init() }
func file_beef_proto_init() {
	if File_beef_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_beef_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeefSummaryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_beef_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeefSummary); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_beef_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_beef_proto_goTypes,
		DependencyIndexes: file_beef_proto_depIdxs,
		MessageInfos:      file_beef_proto_msgTypes,
	}.Build()
	File_beef_proto = out.File
	file_beef_proto_rawDesc = nil
	file_beef_proto_goTypes = nil
	file_beef_proto_depIdxs = nil
}
