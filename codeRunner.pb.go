// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.9.0
// source: codeRunner.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExecuteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid           uint32                 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Language      string                 `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`
	CodeBlock     string                 `protobuf:"bytes,4,opt,name=codeBlock,proto3" json:"codeBlock,omitempty"`
	CallBackUrl   string                 `protobuf:"bytes,5,opt,name=callBackUrl,proto3" json:"callBackUrl,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecuteRequest) Reset() {
	*x = ExecuteRequest{}
	mi := &file_codeRunner_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteRequest) ProtoMessage() {}

func (x *ExecuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_codeRunner_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteRequest.ProtoReflect.Descriptor instead.
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return file_codeRunner_proto_rawDescGZIP(), []int{0}
}

func (x *ExecuteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExecuteRequest) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ExecuteRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *ExecuteRequest) GetCodeBlock() string {
	if x != nil {
		return x.CodeBlock
	}
	return ""
}

func (x *ExecuteRequest) GetCallBackUrl() string {
	if x != nil {
		return x.CallBackUrl
	}
	return ""
}

type ExecuteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid           uint32                 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	GrpcCode      string                 `protobuf:"bytes,3,opt,name=grpcCode,proto3" json:"grpcCode,omitempty"`
	Result        string                 `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	CallBackUrl   string                 `protobuf:"bytes,5,opt,name=callBackUrl,proto3" json:"callBackUrl,omitempty"`
	Err           string                 `protobuf:"bytes,6,opt,name=err,proto3" json:"err,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecuteResponse) Reset() {
	*x = ExecuteResponse{}
	mi := &file_codeRunner_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteResponse) ProtoMessage() {}

func (x *ExecuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_codeRunner_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteResponse.ProtoReflect.Descriptor instead.
func (*ExecuteResponse) Descriptor() ([]byte, []int) {
	return file_codeRunner_proto_rawDescGZIP(), []int{1}
}

func (x *ExecuteResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExecuteResponse) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ExecuteResponse) GetGrpcCode() string {
	if x != nil {
		return x.GrpcCode
	}
	return ""
}

func (x *ExecuteResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *ExecuteResponse) GetCallBackUrl() string {
	if x != nil {
		return x.CallBackUrl
	}
	return ""
}

func (x *ExecuteResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_codeRunner_proto protoreflect.FileDescriptor

var file_codeRunner_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x22, 0x8e, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x55,
	0x72, 0x6c, 0x22, 0x9b, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x70, 0x63,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x70, 0x63,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72,
	0x32, 0x58, 0x0a, 0x0a, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x4a,
	0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x52,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_codeRunner_proto_rawDescOnce sync.Once
	file_codeRunner_proto_rawDescData []byte
)

func file_codeRunner_proto_rawDescGZIP() []byte {
	file_codeRunner_proto_rawDescOnce.Do(func() {
		file_codeRunner_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_codeRunner_proto_rawDesc), len(file_codeRunner_proto_rawDesc)))
	})
	return file_codeRunner_proto_rawDescData
}

var file_codeRunner_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_codeRunner_proto_goTypes = []any{
	(*ExecuteRequest)(nil),  // 0: codeRunner.v1.ExecuteRequest
	(*ExecuteResponse)(nil), // 1: codeRunner.v1.ExecuteResponse
}
var file_codeRunner_proto_depIdxs = []int32{
	0, // 0: codeRunner.v1.CodeRunner.Execute:input_type -> codeRunner.v1.ExecuteRequest
	1, // 1: codeRunner.v1.CodeRunner.Execute:output_type -> codeRunner.v1.ExecuteResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_codeRunner_proto_init() }
func file_codeRunner_proto_init() {
	if File_codeRunner_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_codeRunner_proto_rawDesc), len(file_codeRunner_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_codeRunner_proto_goTypes,
		DependencyIndexes: file_codeRunner_proto_depIdxs,
		MessageInfos:      file_codeRunner_proto_msgTypes,
	}.Build()
	File_codeRunner_proto = out.File
	file_codeRunner_proto_goTypes = nil
	file_codeRunner_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CodeRunnerClient is the client API for CodeRunner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CodeRunnerClient interface {
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
}

type codeRunnerClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeRunnerClient(cc grpc.ClientConnInterface) CodeRunnerClient {
	return &codeRunnerClient{cc}
}

func (c *codeRunnerClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.cc.Invoke(ctx, "/codeRunner.v1.CodeRunner/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeRunnerServer is the server API for CodeRunner service.
type CodeRunnerServer interface {
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
}

// UnimplementedCodeRunnerServer can be embedded to have forward compatible implementations.
type UnimplementedCodeRunnerServer struct {
}

func (*UnimplementedCodeRunnerServer) Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

func RegisterCodeRunnerServer(s *grpc.Server, srv CodeRunnerServer) {
	s.RegisterService(&_CodeRunner_serviceDesc, srv)
}

func _CodeRunner_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeRunnerServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codeRunner.v1.CodeRunner/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeRunnerServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CodeRunner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "codeRunner.v1.CodeRunner",
	HandlerType: (*CodeRunnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _CodeRunner_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "codeRunner.proto",
}
