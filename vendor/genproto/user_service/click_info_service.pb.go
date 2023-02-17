// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: click_info_service.proto

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

type GetAllClickInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipperId string `protobuf:"bytes,1,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	Search    string `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
	Page      uint64 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit     uint64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetAllClickInfoRequest) Reset() {
	*x = GetAllClickInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_click_info_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllClickInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllClickInfoRequest) ProtoMessage() {}

func (x *GetAllClickInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_click_info_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllClickInfoRequest.ProtoReflect.Descriptor instead.
func (*GetAllClickInfoRequest) Descriptor() ([]byte, []int) {
	return file_click_info_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllClickInfoRequest) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *GetAllClickInfoRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetAllClickInfoRequest) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllClickInfoRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAllClickInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClickInfos []*Click `protobuf:"bytes,1,rep,name=click_infos,json=clickInfos,proto3" json:"click_infos,omitempty"`
	Count      uint64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllClickInfoResponse) Reset() {
	*x = GetAllClickInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_click_info_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllClickInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllClickInfoResponse) ProtoMessage() {}

func (x *GetAllClickInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_click_info_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllClickInfoResponse.ProtoReflect.Descriptor instead.
func (*GetAllClickInfoResponse) Descriptor() ([]byte, []int) {
	return file_click_info_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllClickInfoResponse) GetClickInfos() []*Click {
	if x != nil {
		return x.ClickInfos
	}
	return nil
}

func (x *GetAllClickInfoResponse) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ClickInfoId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BranchId string `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
}

func (x *ClickInfoId) Reset() {
	*x = ClickInfoId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_click_info_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClickInfoId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickInfoId) ProtoMessage() {}

func (x *ClickInfoId) ProtoReflect() protoreflect.Message {
	mi := &file_click_info_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickInfoId.ProtoReflect.Descriptor instead.
func (*ClickInfoId) Descriptor() ([]byte, []int) {
	return file_click_info_service_proto_rawDescGZIP(), []int{2}
}

func (x *ClickInfoId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClickInfoId) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

type ClickShipperIdAndKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ClickShipperIdAndKey) Reset() {
	*x = ClickShipperIdAndKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_click_info_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClickShipperIdAndKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickShipperIdAndKey) ProtoMessage() {}

func (x *ClickShipperIdAndKey) ProtoReflect() protoreflect.Message {
	mi := &file_click_info_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickShipperIdAndKey.ProtoReflect.Descriptor instead.
func (*ClickShipperIdAndKey) Descriptor() ([]byte, []int) {
	return file_click_info_service_proto_rawDescGZIP(), []int{3}
}

func (x *ClickShipperIdAndKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClickShipperIdAndKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_click_info_service_proto protoreflect.FileDescriptor

var file_click_info_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6c, 0x69,
	0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x61,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x63, 0x6c, 0x69,
	0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x52,
	0x0a, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x3a, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x22, 0x38, 0x0a,
	0x14, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x41,
	0x6e, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x32, 0x97, 0x03, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x63,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x1a, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x64, 0x22, 0x00,
	0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x64, 0x1a, 0x0f,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x22,
	0x00, 0x12, 0x4f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x20, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6c, 0x69,
	0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43,
	0x6c, 0x69, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69,
	0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x5d, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x41, 0x6e, 0x64, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x12, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6c, 0x69, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x1a,
	0x1e, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b,
	0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x4b, 0x65, 0x79, 0x22,
	0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_click_info_service_proto_rawDescOnce sync.Once
	file_click_info_service_proto_rawDescData = file_click_info_service_proto_rawDesc
)

func file_click_info_service_proto_rawDescGZIP() []byte {
	file_click_info_service_proto_rawDescOnce.Do(func() {
		file_click_info_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_click_info_service_proto_rawDescData)
	})
	return file_click_info_service_proto_rawDescData
}

var file_click_info_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_click_info_service_proto_goTypes = []interface{}{
	(*GetAllClickInfoRequest)(nil),  // 0: genproto.GetAllClickInfoRequest
	(*GetAllClickInfoResponse)(nil), // 1: genproto.GetAllClickInfoResponse
	(*ClickInfoId)(nil),             // 2: genproto.ClickInfoId
	(*ClickShipperIdAndKey)(nil),    // 3: genproto.ClickShipperIdAndKey
	(*Click)(nil),                   // 4: genproto.Click
	(*ClickCredentials)(nil),        // 5: genproto.ClickCredentials
	(*empty.Empty)(nil),             // 6: google.protobuf.Empty
}
var file_click_info_service_proto_depIdxs = []int32{
	4, // 0: genproto.GetAllClickInfoResponse.click_infos:type_name -> genproto.Click
	4, // 1: genproto.ClickInfoService.Create:input_type -> genproto.Click
	2, // 2: genproto.ClickInfoService.Get:input_type -> genproto.ClickInfoId
	0, // 3: genproto.ClickInfoService.GetAll:input_type -> genproto.GetAllClickInfoRequest
	4, // 4: genproto.ClickInfoService.Update:input_type -> genproto.Click
	2, // 5: genproto.ClickInfoService.Delete:input_type -> genproto.ClickInfoId
	5, // 6: genproto.ClickInfoService.GetShipperAndKeyByCredentials:input_type -> genproto.ClickCredentials
	2, // 7: genproto.ClickInfoService.Create:output_type -> genproto.ClickInfoId
	4, // 8: genproto.ClickInfoService.Get:output_type -> genproto.Click
	1, // 9: genproto.ClickInfoService.GetAll:output_type -> genproto.GetAllClickInfoResponse
	6, // 10: genproto.ClickInfoService.Update:output_type -> google.protobuf.Empty
	6, // 11: genproto.ClickInfoService.Delete:output_type -> google.protobuf.Empty
	3, // 12: genproto.ClickInfoService.GetShipperAndKeyByCredentials:output_type -> genproto.ClickShipperIdAndKey
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_click_info_service_proto_init() }
func file_click_info_service_proto_init() {
	if File_click_info_service_proto != nil {
		return
	}
	file_click_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_click_info_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllClickInfoRequest); i {
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
		file_click_info_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllClickInfoResponse); i {
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
		file_click_info_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClickInfoId); i {
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
		file_click_info_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClickShipperIdAndKey); i {
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
			RawDescriptor: file_click_info_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_click_info_service_proto_goTypes,
		DependencyIndexes: file_click_info_service_proto_depIdxs,
		MessageInfos:      file_click_info_service_proto_msgTypes,
	}.Build()
	File_click_info_service_proto = out.File
	file_click_info_service_proto_rawDesc = nil
	file_click_info_service_proto_goTypes = nil
	file_click_info_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClickInfoServiceClient is the client API for ClickInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClickInfoServiceClient interface {
	Create(ctx context.Context, in *Click, opts ...grpc.CallOption) (*ClickInfoId, error)
	Get(ctx context.Context, in *ClickInfoId, opts ...grpc.CallOption) (*Click, error)
	GetAll(ctx context.Context, in *GetAllClickInfoRequest, opts ...grpc.CallOption) (*GetAllClickInfoResponse, error)
	Update(ctx context.Context, in *Click, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *ClickInfoId, opts ...grpc.CallOption) (*empty.Empty, error)
	GetShipperAndKeyByCredentials(ctx context.Context, in *ClickCredentials, opts ...grpc.CallOption) (*ClickShipperIdAndKey, error)
}

type clickInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClickInfoServiceClient(cc grpc.ClientConnInterface) ClickInfoServiceClient {
	return &clickInfoServiceClient{cc}
}

func (c *clickInfoServiceClient) Create(ctx context.Context, in *Click, opts ...grpc.CallOption) (*ClickInfoId, error) {
	out := new(ClickInfoId)
	err := c.cc.Invoke(ctx, "/genproto.ClickInfoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clickInfoServiceClient) Get(ctx context.Context, in *ClickInfoId, opts ...grpc.CallOption) (*Click, error) {
	out := new(Click)
	err := c.cc.Invoke(ctx, "/genproto.ClickInfoService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clickInfoServiceClient) GetAll(ctx context.Context, in *GetAllClickInfoRequest, opts ...grpc.CallOption) (*GetAllClickInfoResponse, error) {
	out := new(GetAllClickInfoResponse)
	err := c.cc.Invoke(ctx, "/genproto.ClickInfoService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clickInfoServiceClient) Update(ctx context.Context, in *Click, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.ClickInfoService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clickInfoServiceClient) Delete(ctx context.Context, in *ClickInfoId, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.ClickInfoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clickInfoServiceClient) GetShipperAndKeyByCredentials(ctx context.Context, in *ClickCredentials, opts ...grpc.CallOption) (*ClickShipperIdAndKey, error) {
	out := new(ClickShipperIdAndKey)
	err := c.cc.Invoke(ctx, "/genproto.ClickInfoService/GetShipperAndKeyByCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClickInfoServiceServer is the server API for ClickInfoService service.
type ClickInfoServiceServer interface {
	Create(context.Context, *Click) (*ClickInfoId, error)
	Get(context.Context, *ClickInfoId) (*Click, error)
	GetAll(context.Context, *GetAllClickInfoRequest) (*GetAllClickInfoResponse, error)
	Update(context.Context, *Click) (*empty.Empty, error)
	Delete(context.Context, *ClickInfoId) (*empty.Empty, error)
	GetShipperAndKeyByCredentials(context.Context, *ClickCredentials) (*ClickShipperIdAndKey, error)
}

// UnimplementedClickInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClickInfoServiceServer struct {
}

func (*UnimplementedClickInfoServiceServer) Create(context.Context, *Click) (*ClickInfoId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedClickInfoServiceServer) Get(context.Context, *ClickInfoId) (*Click, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedClickInfoServiceServer) GetAll(context.Context, *GetAllClickInfoRequest) (*GetAllClickInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedClickInfoServiceServer) Update(context.Context, *Click) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedClickInfoServiceServer) Delete(context.Context, *ClickInfoId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedClickInfoServiceServer) GetShipperAndKeyByCredentials(context.Context, *ClickCredentials) (*ClickShipperIdAndKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShipperAndKeyByCredentials not implemented")
}

func RegisterClickInfoServiceServer(s *grpc.Server, srv ClickInfoServiceServer) {
	s.RegisterService(&_ClickInfoService_serviceDesc, srv)
}

func _ClickInfoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Click)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClickInfoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ClickInfoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClickInfoServiceServer).Create(ctx, req.(*Click))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClickInfoService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClickInfoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClickInfoServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ClickInfoService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClickInfoServiceServer).Get(ctx, req.(*ClickInfoId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClickInfoService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllClickInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClickInfoServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ClickInfoService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClickInfoServiceServer).GetAll(ctx, req.(*GetAllClickInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClickInfoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Click)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClickInfoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ClickInfoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClickInfoServiceServer).Update(ctx, req.(*Click))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClickInfoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClickInfoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClickInfoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ClickInfoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClickInfoServiceServer).Delete(ctx, req.(*ClickInfoId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClickInfoService_GetShipperAndKeyByCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClickCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClickInfoServiceServer).GetShipperAndKeyByCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ClickInfoService/GetShipperAndKeyByCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClickInfoServiceServer).GetShipperAndKeyByCredentials(ctx, req.(*ClickCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClickInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.ClickInfoService",
	HandlerType: (*ClickInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ClickInfoService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ClickInfoService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ClickInfoService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ClickInfoService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ClickInfoService_Delete_Handler,
		},
		{
			MethodName: "GetShipperAndKeyByCredentials",
			Handler:    _ClickInfoService_GetShipperAndKeyByCredentials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "click_info_service.proto",
}