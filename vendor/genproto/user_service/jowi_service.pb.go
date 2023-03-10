// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: jowi_service.proto

package user_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type JowiCredentialsId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipperId string `protobuf:"bytes,1,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
}

func (x *JowiCredentialsId) Reset() {
	*x = JowiCredentialsId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jowi_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JowiCredentialsId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JowiCredentialsId) ProtoMessage() {}

func (x *JowiCredentialsId) ProtoReflect() protoreflect.Message {
	mi := &file_jowi_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JowiCredentialsId.ProtoReflect.Descriptor instead.
func (*JowiCredentialsId) Descriptor() ([]byte, []int) {
	return file_jowi_service_proto_rawDescGZIP(), []int{0}
}

func (x *JowiCredentialsId) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

var File_jowi_service_proto protoreflect.FileDescriptor

var file_jowi_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6a, 0x6f, 0x77, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a,
	0x6a, 0x6f, 0x77, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x11, 0x4a, 0x6f, 0x77, 0x69, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x32, 0x9d, 0x02, 0x0a, 0x16,
	0x4a, 0x6f, 0x77, 0x69, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x19, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x77, 0x69,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0x1b, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x77, 0x69, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x64, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4a, 0x6f, 0x77, 0x69, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x1b, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x77, 0x69,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x64, 0x1a, 0x19, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x77, 0x69, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4a, 0x6f, 0x77, 0x69, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49,
	0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jowi_service_proto_rawDescOnce sync.Once
	file_jowi_service_proto_rawDescData = file_jowi_service_proto_rawDesc
)

func file_jowi_service_proto_rawDescGZIP() []byte {
	file_jowi_service_proto_rawDescOnce.Do(func() {
		file_jowi_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_jowi_service_proto_rawDescData)
	})
	return file_jowi_service_proto_rawDescData
}

var file_jowi_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_jowi_service_proto_goTypes = []interface{}{
	(*JowiCredentialsId)(nil), // 0: genproto.JowiCredentialsId
	(*JowiCredentials)(nil),   // 1: genproto.JowiCredentials
	(*empty.Empty)(nil),       // 2: google.protobuf.Empty
}
var file_jowi_service_proto_depIdxs = []int32{
	1, // 0: genproto.JowiCredentialsService.Create:input_type -> genproto.JowiCredentials
	1, // 1: genproto.JowiCredentialsService.Update:input_type -> genproto.JowiCredentials
	0, // 2: genproto.JowiCredentialsService.Get:input_type -> genproto.JowiCredentialsId
	0, // 3: genproto.JowiCredentialsService.Delete:input_type -> genproto.JowiCredentialsId
	0, // 4: genproto.JowiCredentialsService.Create:output_type -> genproto.JowiCredentialsId
	2, // 5: genproto.JowiCredentialsService.Update:output_type -> google.protobuf.Empty
	1, // 6: genproto.JowiCredentialsService.Get:output_type -> genproto.JowiCredentials
	2, // 7: genproto.JowiCredentialsService.Delete:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_jowi_service_proto_init() }
func file_jowi_service_proto_init() {
	if File_jowi_service_proto != nil {
		return
	}
	file_jowi_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_jowi_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JowiCredentialsId); i {
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
			RawDescriptor: file_jowi_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_jowi_service_proto_goTypes,
		DependencyIndexes: file_jowi_service_proto_depIdxs,
		MessageInfos:      file_jowi_service_proto_msgTypes,
	}.Build()
	File_jowi_service_proto = out.File
	file_jowi_service_proto_rawDesc = nil
	file_jowi_service_proto_goTypes = nil
	file_jowi_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JowiCredentialsServiceClient is the client API for JowiCredentialsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JowiCredentialsServiceClient interface {
	Create(ctx context.Context, in *JowiCredentials, opts ...grpc.CallOption) (*JowiCredentialsId, error)
	Update(ctx context.Context, in *JowiCredentials, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *JowiCredentialsId, opts ...grpc.CallOption) (*JowiCredentials, error)
	Delete(ctx context.Context, in *JowiCredentialsId, opts ...grpc.CallOption) (*empty.Empty, error)
}

type jowiCredentialsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJowiCredentialsServiceClient(cc grpc.ClientConnInterface) JowiCredentialsServiceClient {
	return &jowiCredentialsServiceClient{cc}
}

func (c *jowiCredentialsServiceClient) Create(ctx context.Context, in *JowiCredentials, opts ...grpc.CallOption) (*JowiCredentialsId, error) {
	out := new(JowiCredentialsId)
	err := c.cc.Invoke(ctx, "/genproto.JowiCredentialsService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jowiCredentialsServiceClient) Update(ctx context.Context, in *JowiCredentials, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.JowiCredentialsService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jowiCredentialsServiceClient) Get(ctx context.Context, in *JowiCredentialsId, opts ...grpc.CallOption) (*JowiCredentials, error) {
	out := new(JowiCredentials)
	err := c.cc.Invoke(ctx, "/genproto.JowiCredentialsService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jowiCredentialsServiceClient) Delete(ctx context.Context, in *JowiCredentialsId, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.JowiCredentialsService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JowiCredentialsServiceServer is the server API for JowiCredentialsService service.
type JowiCredentialsServiceServer interface {
	Create(context.Context, *JowiCredentials) (*JowiCredentialsId, error)
	Update(context.Context, *JowiCredentials) (*empty.Empty, error)
	Get(context.Context, *JowiCredentialsId) (*JowiCredentials, error)
	Delete(context.Context, *JowiCredentialsId) (*empty.Empty, error)
}

// UnimplementedJowiCredentialsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJowiCredentialsServiceServer struct {
}

func (*UnimplementedJowiCredentialsServiceServer) Create(context.Context, *JowiCredentials) (*JowiCredentialsId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedJowiCredentialsServiceServer) Update(context.Context, *JowiCredentials) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedJowiCredentialsServiceServer) Get(context.Context, *JowiCredentialsId) (*JowiCredentials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedJowiCredentialsServiceServer) Delete(context.Context, *JowiCredentialsId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterJowiCredentialsServiceServer(s *grpc.Server, srv JowiCredentialsServiceServer) {
	s.RegisterService(&_JowiCredentialsService_serviceDesc, srv)
}

func _JowiCredentialsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JowiCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JowiCredentialsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.JowiCredentialsService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JowiCredentialsServiceServer).Create(ctx, req.(*JowiCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _JowiCredentialsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JowiCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JowiCredentialsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.JowiCredentialsService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JowiCredentialsServiceServer).Update(ctx, req.(*JowiCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _JowiCredentialsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JowiCredentialsId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JowiCredentialsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.JowiCredentialsService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JowiCredentialsServiceServer).Get(ctx, req.(*JowiCredentialsId))
	}
	return interceptor(ctx, in, info, handler)
}

func _JowiCredentialsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JowiCredentialsId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JowiCredentialsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.JowiCredentialsService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JowiCredentialsServiceServer).Delete(ctx, req.(*JowiCredentialsId))
	}
	return interceptor(ctx, in, info, handler)
}

var _JowiCredentialsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.JowiCredentialsService",
	HandlerType: (*JowiCredentialsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _JowiCredentialsService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _JowiCredentialsService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _JowiCredentialsService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _JowiCredentialsService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jowi_service.proto",
}
