// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: iiko.proto

package user_service

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

type IikoCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipperId    string `protobuf:"bytes,1,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	ApiLogin     string `protobuf:"bytes,2,opt,name=api_login,json=apiLogin,proto3" json:"api_login,omitempty"`
	CreatedAt    string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DispatcherId string `protobuf:"bytes,5,opt,name=dispatcher_id,json=dispatcherId,proto3" json:"dispatcher_id,omitempty"`
}

func (x *IikoCredentials) Reset() {
	*x = IikoCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iiko_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IikoCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IikoCredentials) ProtoMessage() {}

func (x *IikoCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_iiko_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IikoCredentials.ProtoReflect.Descriptor instead.
func (*IikoCredentials) Descriptor() ([]byte, []int) {
	return file_iiko_proto_rawDescGZIP(), []int{0}
}

func (x *IikoCredentials) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *IikoCredentials) GetApiLogin() string {
	if x != nil {
		return x.ApiLogin
	}
	return ""
}

func (x *IikoCredentials) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *IikoCredentials) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *IikoCredentials) GetDispatcherId() string {
	if x != nil {
		return x.DispatcherId
	}
	return ""
}

var File_iiko_proto protoreflect.FileDescriptor

var file_iiko_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x69, 0x69, 0x6b, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x0f, 0x49, 0x69, 0x6b, 0x6f, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68,
	0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x70, 0x69,
	0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70,
	0x69, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iiko_proto_rawDescOnce sync.Once
	file_iiko_proto_rawDescData = file_iiko_proto_rawDesc
)

func file_iiko_proto_rawDescGZIP() []byte {
	file_iiko_proto_rawDescOnce.Do(func() {
		file_iiko_proto_rawDescData = protoimpl.X.CompressGZIP(file_iiko_proto_rawDescData)
	})
	return file_iiko_proto_rawDescData
}

var file_iiko_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_iiko_proto_goTypes = []interface{}{
	(*IikoCredentials)(nil), // 0: genproto.IikoCredentials
}
var file_iiko_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_iiko_proto_init() }
func file_iiko_proto_init() {
	if File_iiko_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iiko_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IikoCredentials); i {
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
			RawDescriptor: file_iiko_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_iiko_proto_goTypes,
		DependencyIndexes: file_iiko_proto_depIdxs,
		MessageInfos:      file_iiko_proto_msgTypes,
	}.Build()
	File_iiko_proto = out.File
	file_iiko_proto_rawDesc = nil
	file_iiko_proto_goTypes = nil
	file_iiko_proto_depIdxs = nil
}
