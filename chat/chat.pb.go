// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: chat.proto

package chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MensajeBienvenida struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *MensajeBienvenida) Reset() {
	*x = MensajeBienvenida{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MensajeBienvenida) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MensajeBienvenida) ProtoMessage() {}

func (x *MensajeBienvenida) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MensajeBienvenida.ProtoReflect.Descriptor instead.
func (*MensajeBienvenida) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *MensajeBienvenida) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type MensajeEntreEtapas struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *MensajeEntreEtapas) Reset() {
	*x = MensajeEntreEtapas{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MensajeEntreEtapas) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MensajeEntreEtapas) ProtoMessage() {}

func (x *MensajeEntreEtapas) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MensajeEntreEtapas.ProtoReflect.Descriptor instead.
func (*MensajeEntreEtapas) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *MensajeEntreEtapas) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type MensajeEtapa1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body int32 `protobuf:"varint,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *MensajeEtapa1) Reset() {
	*x = MensajeEtapa1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MensajeEtapa1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MensajeEtapa1) ProtoMessage() {}

func (x *MensajeEtapa1) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MensajeEtapa1.ProtoReflect.Descriptor instead.
func (*MensajeEtapa1) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *MensajeEtapa1) GetBody() int32 {
	if x != nil {
		return x.Body
	}
	return 0
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x22, 0x27, 0x0a, 0x11, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x42, 0x69, 0x65,
	0x6e, 0x76, 0x65, 0x6e, 0x69, 0x64, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x28, 0x0a, 0x12, 0x6d,
	0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x65, 0x45, 0x74, 0x61, 0x70, 0x61,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65,
	0x45, 0x74, 0x61, 0x70, 0x61, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0xca, 0x01, 0x0a, 0x0b, 0x43,
	0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x42, 0x69,
	0x65, 0x6e, 0x76, 0x65, 0x6e, 0x69, 0x64, 0x61, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x42, 0x69, 0x65, 0x6e, 0x76, 0x65, 0x6e, 0x69, 0x64,
	0x61, 0x1a, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65,
	0x42, 0x69, 0x65, 0x6e, 0x76, 0x65, 0x6e, 0x69, 0x64, 0x61, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b,
	0x45, 0x6e, 0x74, 0x72, 0x65, 0x45, 0x74, 0x61, 0x70, 0x61, 0x73, 0x12, 0x18, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x65, 0x45,
	0x74, 0x61, 0x70, 0x61, 0x73, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x6d, 0x65, 0x6e,
	0x73, 0x61, 0x6a, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x65, 0x45, 0x74, 0x61, 0x70, 0x61, 0x73, 0x22,
	0x00, 0x12, 0x34, 0x0a, 0x06, 0x45, 0x74, 0x61, 0x70, 0x61, 0x31, 0x12, 0x13, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x45, 0x74, 0x61, 0x70, 0x61, 0x31,
	0x1a, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x45,
	0x74, 0x61, 0x70, 0x61, 0x31, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x63, 0x68,
	0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chat_proto_goTypes = []interface{}{
	(*MensajeBienvenida)(nil),  // 0: chat.mensajeBienvenida
	(*MensajeEntreEtapas)(nil), // 1: chat.mensajeEntreEtapas
	(*MensajeEtapa1)(nil),      // 2: chat.mensajeEtapa1
}
var file_chat_proto_depIdxs = []int32{
	0, // 0: chat.ChatService.Bienvenida:input_type -> chat.mensajeBienvenida
	1, // 1: chat.ChatService.EntreEtapas:input_type -> chat.mensajeEntreEtapas
	2, // 2: chat.ChatService.Etapa1:input_type -> chat.mensajeEtapa1
	0, // 3: chat.ChatService.Bienvenida:output_type -> chat.mensajeBienvenida
	1, // 4: chat.ChatService.EntreEtapas:output_type -> chat.mensajeEntreEtapas
	2, // 5: chat.ChatService.Etapa1:output_type -> chat.mensajeEtapa1
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MensajeBienvenida); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MensajeEntreEtapas); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MensajeEtapa1); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	Bienvenida(ctx context.Context, in *MensajeBienvenida, opts ...grpc.CallOption) (*MensajeBienvenida, error)
	EntreEtapas(ctx context.Context, in *MensajeEntreEtapas, opts ...grpc.CallOption) (*MensajeEntreEtapas, error)
	Etapa1(ctx context.Context, in *MensajeEtapa1, opts ...grpc.CallOption) (*MensajeEtapa1, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Bienvenida(ctx context.Context, in *MensajeBienvenida, opts ...grpc.CallOption) (*MensajeBienvenida, error) {
	out := new(MensajeBienvenida)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Bienvenida", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) EntreEtapas(ctx context.Context, in *MensajeEntreEtapas, opts ...grpc.CallOption) (*MensajeEntreEtapas, error) {
	out := new(MensajeEntreEtapas)
	err := c.cc.Invoke(ctx, "/chat.ChatService/EntreEtapas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Etapa1(ctx context.Context, in *MensajeEtapa1, opts ...grpc.CallOption) (*MensajeEtapa1, error) {
	out := new(MensajeEtapa1)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Etapa1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	Bienvenida(context.Context, *MensajeBienvenida) (*MensajeBienvenida, error)
	EntreEtapas(context.Context, *MensajeEntreEtapas) (*MensajeEntreEtapas, error)
	Etapa1(context.Context, *MensajeEtapa1) (*MensajeEtapa1, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) Bienvenida(context.Context, *MensajeBienvenida) (*MensajeBienvenida, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bienvenida not implemented")
}
func (*UnimplementedChatServiceServer) EntreEtapas(context.Context, *MensajeEntreEtapas) (*MensajeEntreEtapas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntreEtapas not implemented")
}
func (*UnimplementedChatServiceServer) Etapa1(context.Context, *MensajeEtapa1) (*MensajeEtapa1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Etapa1 not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_Bienvenida_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MensajeBienvenida)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Bienvenida(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Bienvenida",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Bienvenida(ctx, req.(*MensajeBienvenida))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_EntreEtapas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MensajeEntreEtapas)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).EntreEtapas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/EntreEtapas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).EntreEtapas(ctx, req.(*MensajeEntreEtapas))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Etapa1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MensajeEtapa1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Etapa1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Etapa1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Etapa1(ctx, req.(*MensajeEtapa1))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bienvenida",
			Handler:    _ChatService_Bienvenida_Handler,
		},
		{
			MethodName: "EntreEtapas",
			Handler:    _ChatService_EntreEtapas_Handler,
		},
		{
			MethodName: "Etapa1",
			Handler:    _ChatService_Etapa1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
