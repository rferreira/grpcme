// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: grpcme.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Limit *durationpb.Duration `protobuf:"bytes,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Args  *string              `protobuf:"bytes,3,opt,name=args,proto3,oneof" json:"args,omitempty"`
}

func (x *ExecRequest) Reset() {
	*x = ExecRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcme_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecRequest) ProtoMessage() {}

func (x *ExecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcme_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecRequest.ProtoReflect.Descriptor instead.
func (*ExecRequest) Descriptor() ([]byte, []int) {
	return file_grpcme_proto_rawDescGZIP(), []int{0}
}

func (x *ExecRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExecRequest) GetLimit() *durationpb.Duration {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ExecRequest) GetArgs() string {
	if x != nil && x.Args != nil {
		return *x.Args
	}
	return ""
}

type ExecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StdOut     string `protobuf:"bytes,1,opt,name=stdOut,proto3" json:"stdOut,omitempty"`
	StdError   string `protobuf:"bytes,2,opt,name=stdError,proto3" json:"stdError,omitempty"`
	ResultCode int32  `protobuf:"varint,3,opt,name=resultCode,proto3" json:"resultCode,omitempty"`
}

func (x *ExecResponse) Reset() {
	*x = ExecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcme_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecResponse) ProtoMessage() {}

func (x *ExecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcme_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecResponse.ProtoReflect.Descriptor instead.
func (*ExecResponse) Descriptor() ([]byte, []int) {
	return file_grpcme_proto_rawDescGZIP(), []int{1}
}

func (x *ExecResponse) GetStdOut() string {
	if x != nil {
		return x.StdOut
	}
	return ""
}

func (x *ExecResponse) GetStdError() string {
	if x != nil {
		return x.StdError
	}
	return ""
}

func (x *ExecResponse) GetResultCode() int32 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

var File_grpcme_proto protoreflect.FileDescriptor

var file_grpcme_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x67, 0x72, 0x70, 0x63, 0x6d, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x0b, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x61, 0x72, 0x67,
	0x73, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x22, 0x62, 0x0a, 0x0c, 0x45, 0x78, 0x65, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x4f, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x4f, 0x75, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x74, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x74, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x3b, 0x0a, 0x04, 0x45,
	0x78, 0x65, 0x63, 0x12, 0x33, 0x0a, 0x04, 0x45, 0x78, 0x65, 0x63, 0x12, 0x13, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x6d, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcme_proto_rawDescOnce sync.Once
	file_grpcme_proto_rawDescData = file_grpcme_proto_rawDesc
)

func file_grpcme_proto_rawDescGZIP() []byte {
	file_grpcme_proto_rawDescOnce.Do(func() {
		file_grpcme_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcme_proto_rawDescData)
	})
	return file_grpcme_proto_rawDescData
}

var file_grpcme_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpcme_proto_goTypes = []interface{}{
	(*ExecRequest)(nil),         // 0: grpcme.ExecRequest
	(*ExecResponse)(nil),        // 1: grpcme.ExecResponse
	(*durationpb.Duration)(nil), // 2: google.protobuf.Duration
}
var file_grpcme_proto_depIdxs = []int32{
	2, // 0: grpcme.ExecRequest.limit:type_name -> google.protobuf.Duration
	0, // 1: grpcme.Exec.Exec:input_type -> grpcme.ExecRequest
	1, // 2: grpcme.Exec.Exec:output_type -> grpcme.ExecResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpcme_proto_init() }
func file_grpcme_proto_init() {
	if File_grpcme_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcme_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecRequest); i {
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
		file_grpcme_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecResponse); i {
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
	file_grpcme_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpcme_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcme_proto_goTypes,
		DependencyIndexes: file_grpcme_proto_depIdxs,
		MessageInfos:      file_grpcme_proto_msgTypes,
	}.Build()
	File_grpcme_proto = out.File
	file_grpcme_proto_rawDesc = nil
	file_grpcme_proto_goTypes = nil
	file_grpcme_proto_depIdxs = nil
}
