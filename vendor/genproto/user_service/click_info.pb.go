// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: click_info.proto

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

type Click struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipperId      string `protobuf:"bytes,1,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	MerchantId     int64  `protobuf:"varint,2,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	ServiceId      int64  `protobuf:"varint,3,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	MerchantUserId int64  `protobuf:"varint,4,opt,name=merchant_user_id,json=merchantUserId,proto3" json:"merchant_user_id,omitempty"` // optional
	Key            string `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	CreatedAt      string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	BranchId       string `protobuf:"bytes,8,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	BranchName     string `protobuf:"bytes,9,opt,name=branch_name,json=branchName,proto3" json:"branch_name,omitempty"`
}

func (x *Click) Reset() {
	*x = Click{}
	if protoimpl.UnsafeEnabled {
		mi := &file_click_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Click) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Click) ProtoMessage() {}

func (x *Click) ProtoReflect() protoreflect.Message {
	mi := &file_click_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Click.ProtoReflect.Descriptor instead.
func (*Click) Descriptor() ([]byte, []int) {
	return file_click_info_proto_rawDescGZIP(), []int{0}
}

func (x *Click) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *Click) GetMerchantId() int64 {
	if x != nil {
		return x.MerchantId
	}
	return 0
}

func (x *Click) GetServiceId() int64 {
	if x != nil {
		return x.ServiceId
	}
	return 0
}

func (x *Click) GetMerchantUserId() int64 {
	if x != nil {
		return x.MerchantUserId
	}
	return 0
}

func (x *Click) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Click) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Click) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Click) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *Click) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

type ClickCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId int64 `protobuf:"varint,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
}

func (x *ClickCredentials) Reset() {
	*x = ClickCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_click_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClickCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickCredentials) ProtoMessage() {}

func (x *ClickCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_click_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickCredentials.ProtoReflect.Descriptor instead.
func (*ClickCredentials) Descriptor() ([]byte, []int) {
	return file_click_info_proto_rawDescGZIP(), []int{1}
}

func (x *ClickCredentials) GetServiceId() int64 {
	if x != nil {
		return x.ServiceId
	}
	return 0
}

var File_click_info_proto protoreflect.FileDescriptor

var file_click_info_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x02, 0x0a,
	0x05, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70,
	0x70, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a,
	0x10, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_click_info_proto_rawDescOnce sync.Once
	file_click_info_proto_rawDescData = file_click_info_proto_rawDesc
)

func file_click_info_proto_rawDescGZIP() []byte {
	file_click_info_proto_rawDescOnce.Do(func() {
		file_click_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_click_info_proto_rawDescData)
	})
	return file_click_info_proto_rawDescData
}

var file_click_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_click_info_proto_goTypes = []interface{}{
	(*Click)(nil),            // 0: genproto.Click
	(*ClickCredentials)(nil), // 1: genproto.ClickCredentials
}
var file_click_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_click_info_proto_init() }
func file_click_info_proto_init() {
	if File_click_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_click_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Click); i {
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
		file_click_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClickCredentials); i {
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
			RawDescriptor: file_click_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_click_info_proto_goTypes,
		DependencyIndexes: file_click_info_proto_depIdxs,
		MessageInfos:      file_click_info_proto_msgTypes,
	}.Build()
	File_click_info_proto = out.File
	file_click_info_proto_rawDesc = nil
	file_click_info_proto_goTypes = nil
	file_click_info_proto_depIdxs = nil
}
