// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: menus_service.proto

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
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_menus_service_proto protoreflect.FileDescriptor

var file_menus_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0b, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb9, 0x02, 0x0a, 0x0b, 0x4d, 0x65,
	0x6e, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x1a, 0x19, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x1c, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d,
	0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65,
	0x71, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_menus_service_proto_goTypes = []interface{}{
	(*MenuRequest)(nil),         // 0: genproto.MenuRequest
	(*GetMenu)(nil),             // 1: genproto.GetMenu
	(*GetAllMenusRequest)(nil),  // 2: genproto.GetAllMenusRequest
	(*UpdateMenuReqest)(nil),    // 3: genproto.UpdateMenuReqest
	(*MenuId)(nil),              // 4: genproto.MenuId
	(*GetMenuResponse)(nil),     // 5: genproto.GetMenuResponse
	(*GetAllMenusResponse)(nil), // 6: genproto.GetAllMenusResponse
	(*empty.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_menus_service_proto_depIdxs = []int32{
	0, // 0: genproto.MenuService.Create:input_type -> genproto.MenuRequest
	1, // 1: genproto.MenuService.Get:input_type -> genproto.GetMenu
	2, // 2: genproto.MenuService.GetAll:input_type -> genproto.GetAllMenusRequest
	3, // 3: genproto.MenuService.Update:input_type -> genproto.UpdateMenuReqest
	1, // 4: genproto.MenuService.Delete:input_type -> genproto.GetMenu
	4, // 5: genproto.MenuService.Create:output_type -> genproto.MenuId
	5, // 6: genproto.MenuService.Get:output_type -> genproto.GetMenuResponse
	6, // 7: genproto.MenuService.GetAll:output_type -> genproto.GetAllMenusResponse
	7, // 8: genproto.MenuService.Update:output_type -> google.protobuf.Empty
	7, // 9: genproto.MenuService.Delete:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_menus_service_proto_init() }
func file_menus_service_proto_init() {
	if File_menus_service_proto != nil {
		return
	}
	file_menus_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_menus_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_menus_service_proto_goTypes,
		DependencyIndexes: file_menus_service_proto_depIdxs,
	}.Build()
	File_menus_service_proto = out.File
	file_menus_service_proto_rawDesc = nil
	file_menus_service_proto_goTypes = nil
	file_menus_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MenuServiceClient is the client API for MenuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MenuServiceClient interface {
	Create(ctx context.Context, in *MenuRequest, opts ...grpc.CallOption) (*MenuId, error)
	Get(ctx context.Context, in *GetMenu, opts ...grpc.CallOption) (*GetMenuResponse, error)
	GetAll(ctx context.Context, in *GetAllMenusRequest, opts ...grpc.CallOption) (*GetAllMenusResponse, error)
	Update(ctx context.Context, in *UpdateMenuReqest, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *GetMenu, opts ...grpc.CallOption) (*empty.Empty, error)
}

type menuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMenuServiceClient(cc grpc.ClientConnInterface) MenuServiceClient {
	return &menuServiceClient{cc}
}

func (c *menuServiceClient) Create(ctx context.Context, in *MenuRequest, opts ...grpc.CallOption) (*MenuId, error) {
	out := new(MenuId)
	err := c.cc.Invoke(ctx, "/genproto.MenuService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) Get(ctx context.Context, in *GetMenu, opts ...grpc.CallOption) (*GetMenuResponse, error) {
	out := new(GetMenuResponse)
	err := c.cc.Invoke(ctx, "/genproto.MenuService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) GetAll(ctx context.Context, in *GetAllMenusRequest, opts ...grpc.CallOption) (*GetAllMenusResponse, error) {
	out := new(GetAllMenusResponse)
	err := c.cc.Invoke(ctx, "/genproto.MenuService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) Update(ctx context.Context, in *UpdateMenuReqest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.MenuService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) Delete(ctx context.Context, in *GetMenu, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.MenuService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MenuServiceServer is the server API for MenuService service.
type MenuServiceServer interface {
	Create(context.Context, *MenuRequest) (*MenuId, error)
	Get(context.Context, *GetMenu) (*GetMenuResponse, error)
	GetAll(context.Context, *GetAllMenusRequest) (*GetAllMenusResponse, error)
	Update(context.Context, *UpdateMenuReqest) (*empty.Empty, error)
	Delete(context.Context, *GetMenu) (*empty.Empty, error)
}

// UnimplementedMenuServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMenuServiceServer struct {
}

func (*UnimplementedMenuServiceServer) Create(context.Context, *MenuRequest) (*MenuId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedMenuServiceServer) Get(context.Context, *GetMenu) (*GetMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedMenuServiceServer) GetAll(context.Context, *GetAllMenusRequest) (*GetAllMenusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedMenuServiceServer) Update(context.Context, *UpdateMenuReqest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedMenuServiceServer) Delete(context.Context, *GetMenu) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterMenuServiceServer(s *grpc.Server, srv MenuServiceServer) {
	s.RegisterService(&_MenuService_serviceDesc, srv)
}

func _MenuService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.MenuService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).Create(ctx, req.(*MenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMenu)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.MenuService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).Get(ctx, req.(*GetMenu))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllMenusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.MenuService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).GetAll(ctx, req.(*GetAllMenusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMenuReqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.MenuService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).Update(ctx, req.(*UpdateMenuReqest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMenu)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.MenuService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).Delete(ctx, req.(*GetMenu))
	}
	return interceptor(ctx, in, info, handler)
}

var _MenuService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.MenuService",
	HandlerType: (*MenuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _MenuService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _MenuService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _MenuService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MenuService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _MenuService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "menus_service.proto",
}
