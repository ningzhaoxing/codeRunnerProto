// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.9.0
// source: tokenIssuer.proto

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

type GenerateTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateTokenRequest) Reset() {
	*x = GenerateTokenRequest{}
	mi := &file_tokenIssuer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenRequest) ProtoMessage() {}

func (x *GenerateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tokenIssuer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenRequest.ProtoReflect.Descriptor instead.
func (*GenerateTokenRequest) Descriptor() ([]byte, []int) {
	return file_tokenIssuer_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateTokenRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GenerateTokenRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GenerateTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateTokenResponse) Reset() {
	*x = GenerateTokenResponse{}
	mi := &file_tokenIssuer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenResponse) ProtoMessage() {}

func (x *GenerateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tokenIssuer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenResponse.ProtoReflect.Descriptor instead.
func (*GenerateTokenResponse) Descriptor() ([]byte, []int) {
	return file_tokenIssuer_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_tokenIssuer_proto protoreflect.FileDescriptor

var file_tokenIssuer_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x22, 0x46, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2d, 0x0a, 0x15, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x6b, 0x0a, 0x0b, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x5c, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_tokenIssuer_proto_rawDescOnce sync.Once
	file_tokenIssuer_proto_rawDescData []byte
)

func file_tokenIssuer_proto_rawDescGZIP() []byte {
	file_tokenIssuer_proto_rawDescOnce.Do(func() {
		file_tokenIssuer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tokenIssuer_proto_rawDesc), len(file_tokenIssuer_proto_rawDesc)))
	})
	return file_tokenIssuer_proto_rawDescData
}

var file_tokenIssuer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tokenIssuer_proto_goTypes = []any{
	(*GenerateTokenRequest)(nil),  // 0: codeRunner.v1.GenerateTokenRequest
	(*GenerateTokenResponse)(nil), // 1: codeRunner.v1.GenerateTokenResponse
}
var file_tokenIssuer_proto_depIdxs = []int32{
	0, // 0: codeRunner.v1.tokenIssuer.GenerateToken:input_type -> codeRunner.v1.GenerateTokenRequest
	1, // 1: codeRunner.v1.tokenIssuer.GenerateToken:output_type -> codeRunner.v1.GenerateTokenResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tokenIssuer_proto_init() }
func file_tokenIssuer_proto_init() {
	if File_tokenIssuer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tokenIssuer_proto_rawDesc), len(file_tokenIssuer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tokenIssuer_proto_goTypes,
		DependencyIndexes: file_tokenIssuer_proto_depIdxs,
		MessageInfos:      file_tokenIssuer_proto_msgTypes,
	}.Build()
	File_tokenIssuer_proto = out.File
	file_tokenIssuer_proto_goTypes = nil
	file_tokenIssuer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TokenIssuerClient is the client API for TokenIssuer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokenIssuerClient interface {
	GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error)
}

type tokenIssuerClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenIssuerClient(cc grpc.ClientConnInterface) TokenIssuerClient {
	return &tokenIssuerClient{cc}
}

func (c *tokenIssuerClient) GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error) {
	out := new(GenerateTokenResponse)
	err := c.cc.Invoke(ctx, "/codeRunner.v1.tokenIssuer/GenerateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenIssuerServer is the server API for TokenIssuer service.
type TokenIssuerServer interface {
	GenerateToken(context.Context, *GenerateTokenRequest) (*GenerateTokenResponse, error)
}

// UnimplementedTokenIssuerServer can be embedded to have forward compatible implementations.
type UnimplementedTokenIssuerServer struct {
}

func (*UnimplementedTokenIssuerServer) GenerateToken(context.Context, *GenerateTokenRequest) (*GenerateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}

func RegisterTokenIssuerServer(s *grpc.Server, srv TokenIssuerServer) {
	s.RegisterService(&_TokenIssuer_serviceDesc, srv)
}

func _TokenIssuer_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenIssuerServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codeRunner.v1.tokenIssuer/GenerateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenIssuerServer).GenerateToken(ctx, req.(*GenerateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokenIssuer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "codeRunner.v1.tokenIssuer",
	HandlerType: (*TokenIssuerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateToken",
			Handler:    _TokenIssuer_GenerateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tokenIssuer.proto",
}
