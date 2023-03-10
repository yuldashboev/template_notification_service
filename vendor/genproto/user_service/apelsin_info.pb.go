// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: apelsin_info.proto

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

type Apelsin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipperId  string `protobuf:"bytes,1,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	CashId     string `protobuf:"bytes,2,opt,name=cash_id,json=cashId,proto3" json:"cash_id,omitempty"`
	CreatedAt  string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	BranchId   string `protobuf:"bytes,5,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	BranchName string `protobuf:"bytes,6,opt,name=branch_name,json=branchName,proto3" json:"branch_name,omitempty"`
}

func (x *Apelsin) Reset() {
	*x = Apelsin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apelsin_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Apelsin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Apelsin) ProtoMessage() {}

func (x *Apelsin) ProtoReflect() protoreflect.Message {
	mi := &file_apelsin_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Apelsin.ProtoReflect.Descriptor instead.
func (*Apelsin) Descriptor() ([]byte, []int) {
	return file_apelsin_info_proto_rawDescGZIP(), []int{0}
}

func (x *Apelsin) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *Apelsin) GetCashId() string {
	if x != nil {
		return x.CashId
	}
	return ""
}

func (x *Apelsin) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Apelsin) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Apelsin) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *Apelsin) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

type ApelsinInfoId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BranchId string `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
}

func (x *ApelsinInfoId) Reset() {
	*x = ApelsinInfoId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apelsin_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApelsinInfoId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApelsinInfoId) ProtoMessage() {}

func (x *ApelsinInfoId) ProtoReflect() protoreflect.Message {
	mi := &file_apelsin_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApelsinInfoId.ProtoReflect.Descriptor instead.
func (*ApelsinInfoId) Descriptor() ([]byte, []int) {
	return file_apelsin_info_proto_rawDescGZIP(), []int{1}
}

func (x *ApelsinInfoId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ApelsinInfoId) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

type GetAllApelsinInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipperId string `protobuf:"bytes,1,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	Search    string `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
	Page      uint64 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit     uint64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetAllApelsinInfoRequest) Reset() {
	*x = GetAllApelsinInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apelsin_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllApelsinInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllApelsinInfoRequest) ProtoMessage() {}

func (x *GetAllApelsinInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apelsin_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllApelsinInfoRequest.ProtoReflect.Descriptor instead.
func (*GetAllApelsinInfoRequest) Descriptor() ([]byte, []int) {
	return file_apelsin_info_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllApelsinInfoRequest) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *GetAllApelsinInfoRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetAllApelsinInfoRequest) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllApelsinInfoRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAllApelsinInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApelsinInfos []*Apelsin `protobuf:"bytes,1,rep,name=apelsin_infos,json=apelsinInfos,proto3" json:"apelsin_infos,omitempty"`
	Count        uint64     `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllApelsinInfoResponse) Reset() {
	*x = GetAllApelsinInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apelsin_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllApelsinInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllApelsinInfoResponse) ProtoMessage() {}

func (x *GetAllApelsinInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apelsin_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllApelsinInfoResponse.ProtoReflect.Descriptor instead.
func (*GetAllApelsinInfoResponse) Descriptor() ([]byte, []int) {
	return file_apelsin_info_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllApelsinInfoResponse) GetApelsinInfos() []*Apelsin {
	if x != nil {
		return x.ApelsinInfos
	}
	return nil
}

func (x *GetAllApelsinInfoResponse) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_apelsin_info_proto protoreflect.FileDescriptor

var file_apelsin_info_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x65, 0x6c, 0x73, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x07,
	0x41, 0x70, 0x65, 0x6c, 0x73, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x73, 0x68, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x73, 0x68, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x0d, 0x41,
	0x70, 0x65, 0x6c, 0x73, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x22, 0x7b, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x41, 0x70, 0x65, 0x6c, 0x73, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x69, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x41, 0x70, 0x65, 0x6c, 0x73, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x61, 0x70, 0x65, 0x6c, 0x73, 0x69, 0x6e, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x70, 0x65, 0x6c, 0x73, 0x69, 0x6e, 0x52, 0x0c, 0x61,
	0x70, 0x65, 0x6c, 0x73, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x32, 0xca, 0x02, 0x0a, 0x12, 0x41, 0x70, 0x65, 0x6c, 0x73, 0x69, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x70,
	0x65, 0x6c, 0x73, 0x69, 0x6e, 0x1a, 0x17, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x70, 0x65, 0x6c, 0x73, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x64, 0x22, 0x00,
	0x12, 0x33, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x70, 0x65, 0x6c, 0x73, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x64,
	0x1a, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x70, 0x65, 0x6c,
	0x73, 0x69, 0x6e, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x22, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x41, 0x70, 0x65, 0x6c, 0x73, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x70, 0x65, 0x6c, 0x73, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x70, 0x65, 0x6c, 0x73, 0x69, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x70, 0x65, 0x6c, 0x73, 0x69, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x17,
	0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apelsin_info_proto_rawDescOnce sync.Once
	file_apelsin_info_proto_rawDescData = file_apelsin_info_proto_rawDesc
)

func file_apelsin_info_proto_rawDescGZIP() []byte {
	file_apelsin_info_proto_rawDescOnce.Do(func() {
		file_apelsin_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_apelsin_info_proto_rawDescData)
	})
	return file_apelsin_info_proto_rawDescData
}

var file_apelsin_info_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apelsin_info_proto_goTypes = []interface{}{
	(*Apelsin)(nil),                   // 0: genproto.Apelsin
	(*ApelsinInfoId)(nil),             // 1: genproto.ApelsinInfoId
	(*GetAllApelsinInfoRequest)(nil),  // 2: genproto.GetAllApelsinInfoRequest
	(*GetAllApelsinInfoResponse)(nil), // 3: genproto.GetAllApelsinInfoResponse
	(*empty.Empty)(nil),               // 4: google.protobuf.Empty
}
var file_apelsin_info_proto_depIdxs = []int32{
	0, // 0: genproto.GetAllApelsinInfoResponse.apelsin_infos:type_name -> genproto.Apelsin
	0, // 1: genproto.ApelsinInfoService.Create:input_type -> genproto.Apelsin
	1, // 2: genproto.ApelsinInfoService.Get:input_type -> genproto.ApelsinInfoId
	2, // 3: genproto.ApelsinInfoService.GetAll:input_type -> genproto.GetAllApelsinInfoRequest
	0, // 4: genproto.ApelsinInfoService.Update:input_type -> genproto.Apelsin
	1, // 5: genproto.ApelsinInfoService.Delete:input_type -> genproto.ApelsinInfoId
	1, // 6: genproto.ApelsinInfoService.Create:output_type -> genproto.ApelsinInfoId
	0, // 7: genproto.ApelsinInfoService.Get:output_type -> genproto.Apelsin
	3, // 8: genproto.ApelsinInfoService.GetAll:output_type -> genproto.GetAllApelsinInfoResponse
	4, // 9: genproto.ApelsinInfoService.Update:output_type -> google.protobuf.Empty
	4, // 10: genproto.ApelsinInfoService.Delete:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apelsin_info_proto_init() }
func file_apelsin_info_proto_init() {
	if File_apelsin_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apelsin_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Apelsin); i {
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
		file_apelsin_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApelsinInfoId); i {
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
		file_apelsin_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllApelsinInfoRequest); i {
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
		file_apelsin_info_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllApelsinInfoResponse); i {
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
			RawDescriptor: file_apelsin_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apelsin_info_proto_goTypes,
		DependencyIndexes: file_apelsin_info_proto_depIdxs,
		MessageInfos:      file_apelsin_info_proto_msgTypes,
	}.Build()
	File_apelsin_info_proto = out.File
	file_apelsin_info_proto_rawDesc = nil
	file_apelsin_info_proto_goTypes = nil
	file_apelsin_info_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ApelsinInfoServiceClient is the client API for ApelsinInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApelsinInfoServiceClient interface {
	Create(ctx context.Context, in *Apelsin, opts ...grpc.CallOption) (*ApelsinInfoId, error)
	Get(ctx context.Context, in *ApelsinInfoId, opts ...grpc.CallOption) (*Apelsin, error)
	GetAll(ctx context.Context, in *GetAllApelsinInfoRequest, opts ...grpc.CallOption) (*GetAllApelsinInfoResponse, error)
	Update(ctx context.Context, in *Apelsin, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *ApelsinInfoId, opts ...grpc.CallOption) (*empty.Empty, error)
}

type apelsinInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApelsinInfoServiceClient(cc grpc.ClientConnInterface) ApelsinInfoServiceClient {
	return &apelsinInfoServiceClient{cc}
}

func (c *apelsinInfoServiceClient) Create(ctx context.Context, in *Apelsin, opts ...grpc.CallOption) (*ApelsinInfoId, error) {
	out := new(ApelsinInfoId)
	err := c.cc.Invoke(ctx, "/genproto.ApelsinInfoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apelsinInfoServiceClient) Get(ctx context.Context, in *ApelsinInfoId, opts ...grpc.CallOption) (*Apelsin, error) {
	out := new(Apelsin)
	err := c.cc.Invoke(ctx, "/genproto.ApelsinInfoService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apelsinInfoServiceClient) GetAll(ctx context.Context, in *GetAllApelsinInfoRequest, opts ...grpc.CallOption) (*GetAllApelsinInfoResponse, error) {
	out := new(GetAllApelsinInfoResponse)
	err := c.cc.Invoke(ctx, "/genproto.ApelsinInfoService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apelsinInfoServiceClient) Update(ctx context.Context, in *Apelsin, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.ApelsinInfoService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apelsinInfoServiceClient) Delete(ctx context.Context, in *ApelsinInfoId, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.ApelsinInfoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApelsinInfoServiceServer is the server API for ApelsinInfoService service.
type ApelsinInfoServiceServer interface {
	Create(context.Context, *Apelsin) (*ApelsinInfoId, error)
	Get(context.Context, *ApelsinInfoId) (*Apelsin, error)
	GetAll(context.Context, *GetAllApelsinInfoRequest) (*GetAllApelsinInfoResponse, error)
	Update(context.Context, *Apelsin) (*empty.Empty, error)
	Delete(context.Context, *ApelsinInfoId) (*empty.Empty, error)
}

// UnimplementedApelsinInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedApelsinInfoServiceServer struct {
}

func (*UnimplementedApelsinInfoServiceServer) Create(context.Context, *Apelsin) (*ApelsinInfoId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedApelsinInfoServiceServer) Get(context.Context, *ApelsinInfoId) (*Apelsin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedApelsinInfoServiceServer) GetAll(context.Context, *GetAllApelsinInfoRequest) (*GetAllApelsinInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedApelsinInfoServiceServer) Update(context.Context, *Apelsin) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedApelsinInfoServiceServer) Delete(context.Context, *ApelsinInfoId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterApelsinInfoServiceServer(s *grpc.Server, srv ApelsinInfoServiceServer) {
	s.RegisterService(&_ApelsinInfoService_serviceDesc, srv)
}

func _ApelsinInfoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Apelsin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApelsinInfoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ApelsinInfoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApelsinInfoServiceServer).Create(ctx, req.(*Apelsin))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApelsinInfoService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApelsinInfoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApelsinInfoServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ApelsinInfoService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApelsinInfoServiceServer).Get(ctx, req.(*ApelsinInfoId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApelsinInfoService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllApelsinInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApelsinInfoServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ApelsinInfoService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApelsinInfoServiceServer).GetAll(ctx, req.(*GetAllApelsinInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApelsinInfoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Apelsin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApelsinInfoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ApelsinInfoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApelsinInfoServiceServer).Update(ctx, req.(*Apelsin))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApelsinInfoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApelsinInfoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApelsinInfoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ApelsinInfoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApelsinInfoServiceServer).Delete(ctx, req.(*ApelsinInfoId))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApelsinInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.ApelsinInfoService",
	HandlerType: (*ApelsinInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ApelsinInfoService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ApelsinInfoService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ApelsinInfoService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ApelsinInfoService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ApelsinInfoService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apelsin_info.proto",
}
