// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: shipper_user.proto

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

type ShipperUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Username   string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password   string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Phone      string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	IsBlocked  bool   `protobuf:"varint,6,opt,name=is_blocked,json=isBlocked,proto3" json:"is_blocked,omitempty"`
	CreatedAt  string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ShipperId  string `protobuf:"bytes,9,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	UserRoleId string `protobuf:"bytes,10,opt,name=user_role_id,json=userRoleId,proto3" json:"user_role_id,omitempty"`
	FcmToken   string `protobuf:"bytes,11,opt,name=fcm_token,json=fcmToken,proto3" json:"fcm_token,omitempty"`
}

func (x *ShipperUser) Reset() {
	*x = ShipperUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipper_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipperUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipperUser) ProtoMessage() {}

func (x *ShipperUser) ProtoReflect() protoreflect.Message {
	mi := &file_shipper_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipperUser.ProtoReflect.Descriptor instead.
func (*ShipperUser) Descriptor() ([]byte, []int) {
	return file_shipper_user_proto_rawDescGZIP(), []int{0}
}

func (x *ShipperUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShipperUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShipperUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ShipperUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ShipperUser) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *ShipperUser) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

func (x *ShipperUser) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ShipperUser) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ShipperUser) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *ShipperUser) GetUserRoleId() string {
	if x != nil {
		return x.UserRoleId
	}
	return ""
}

func (x *ShipperUser) GetFcmToken() string {
	if x != nil {
		return x.FcmToken
	}
	return ""
}

var File_shipper_user_proto protoreflect.FileDescriptor

var file_shipper_user_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba,
	0x02, 0x0a, 0x0b, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x63, 0x6d, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x63, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x17, 0x5a, 0x15, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shipper_user_proto_rawDescOnce sync.Once
	file_shipper_user_proto_rawDescData = file_shipper_user_proto_rawDesc
)

func file_shipper_user_proto_rawDescGZIP() []byte {
	file_shipper_user_proto_rawDescOnce.Do(func() {
		file_shipper_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_shipper_user_proto_rawDescData)
	})
	return file_shipper_user_proto_rawDescData
}

var file_shipper_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_shipper_user_proto_goTypes = []interface{}{
	(*ShipperUser)(nil), // 0: genproto.ShipperUser
}
var file_shipper_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shipper_user_proto_init() }
func file_shipper_user_proto_init() {
	if File_shipper_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shipper_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipperUser); i {
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
			RawDescriptor: file_shipper_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shipper_user_proto_goTypes,
		DependencyIndexes: file_shipper_user_proto_depIdxs,
		MessageInfos:      file_shipper_user_proto_msgTypes,
	}.Build()
	File_shipper_user_proto = out.File
	file_shipper_user_proto_rawDesc = nil
	file_shipper_user_proto_goTypes = nil
	file_shipper_user_proto_depIdxs = nil
}
