// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: pkg/pb/company.proto

package pb

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

type CreateCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Amount      int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Registered  string `protobuf:"bytes,4,opt,name=registered,proto3" json:"registered,omitempty"`
	Type        string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *CreateCompanyRequest) Reset() {
	*x = CreateCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_company_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCompanyRequest) ProtoMessage() {}

func (x *CreateCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_company_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCompanyRequest.ProtoReflect.Descriptor instead.
func (*CreateCompanyRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_company_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCompanyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCompanyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateCompanyRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateCompanyRequest) GetRegistered() string {
	if x != nil {
		return x.Registered
	}
	return ""
}

func (x *CreateCompanyRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type CreateCompanyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id     int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateCompanyResponse) Reset() {
	*x = CreateCompanyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_company_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCompanyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCompanyResponse) ProtoMessage() {}

func (x *CreateCompanyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_company_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCompanyResponse.ProtoReflect.Descriptor instead.
func (*CreateCompanyResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_company_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCompanyResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateCompanyResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreateCompanyResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PatchCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Amount      int64  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Registered  string `protobuf:"bytes,5,opt,name=registered,proto3" json:"registered,omitempty"`
	Type        string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *PatchCompanyRequest) Reset() {
	*x = PatchCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_company_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchCompanyRequest) ProtoMessage() {}

func (x *PatchCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_company_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchCompanyRequest.ProtoReflect.Descriptor instead.
func (*PatchCompanyRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_company_proto_rawDescGZIP(), []int{2}
}

func (x *PatchCompanyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PatchCompanyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PatchCompanyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PatchCompanyRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PatchCompanyRequest) GetRegistered() string {
	if x != nil {
		return x.Registered
	}
	return ""
}

func (x *PatchCompanyRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type PatchCompanyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id     int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PatchCompanyResponse) Reset() {
	*x = PatchCompanyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_company_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchCompanyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchCompanyResponse) ProtoMessage() {}

func (x *PatchCompanyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_company_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchCompanyResponse.ProtoReflect.Descriptor instead.
func (*PatchCompanyResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_company_proto_rawDescGZIP(), []int{3}
}

func (x *PatchCompanyResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PatchCompanyResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *PatchCompanyResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetOneCompanyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Amount      int64  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Registered  string `protobuf:"bytes,5,opt,name=registered,proto3" json:"registered,omitempty"`
	Type        string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GetOneCompanyData) Reset() {
	*x = GetOneCompanyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_company_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneCompanyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneCompanyData) ProtoMessage() {}

func (x *GetOneCompanyData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_company_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneCompanyData.ProtoReflect.Descriptor instead.
func (*GetOneCompanyData) Descriptor() ([]byte, []int) {
	return file_pkg_pb_company_proto_rawDescGZIP(), []int{4}
}

func (x *GetOneCompanyData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetOneCompanyData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetOneCompanyData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetOneCompanyData) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *GetOneCompanyData) GetRegistered() string {
	if x != nil {
		return x.Registered
	}
	return ""
}

func (x *GetOneCompanyData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetOneCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOneCompanyRequest) Reset() {
	*x = GetOneCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_company_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneCompanyRequest) ProtoMessage() {}

func (x *GetOneCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_company_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneCompanyRequest.ProtoReflect.Descriptor instead.
func (*GetOneCompanyRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_company_proto_rawDescGZIP(), []int{5}
}

func (x *GetOneCompanyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetOneCompanyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64              `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string             `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   *GetOneCompanyData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetOneCompanyResponse) Reset() {
	*x = GetOneCompanyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_company_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneCompanyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneCompanyResponse) ProtoMessage() {}

func (x *GetOneCompanyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_company_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneCompanyResponse.ProtoReflect.Descriptor instead.
func (*GetOneCompanyResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_company_proto_rawDescGZIP(), []int{6}
}

func (x *GetOneCompanyResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetOneCompanyResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetOneCompanyResponse) GetData() *GetOneCompanyData {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCompanyRequest) Reset() {
	*x = DeleteCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_company_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCompanyRequest) ProtoMessage() {}

func (x *DeleteCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_company_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCompanyRequest.ProtoReflect.Descriptor instead.
func (*DeleteCompanyRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_company_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCompanyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCompanyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteCompanyResponse) Reset() {
	*x = DeleteCompanyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_company_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCompanyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCompanyResponse) ProtoMessage() {}

func (x *DeleteCompanyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_company_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCompanyResponse.ProtoReflect.Descriptor instead.
func (*DeleteCompanyResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_company_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCompanyResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteCompanyResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_pb_company_proto protoreflect.FileDescriptor

var file_pkg_pb_company_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22,
	0x98, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x55, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xa7, 0x01, 0x0a, 0x13, 0x50, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x54, 0x0a, 0x14, 0x50,
	0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xa5, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x4f, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x75, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x45, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xd5, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1d, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1c, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x4f, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1d, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1d,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_company_proto_rawDescOnce sync.Once
	file_pkg_pb_company_proto_rawDescData = file_pkg_pb_company_proto_rawDesc
)

func file_pkg_pb_company_proto_rawDescGZIP() []byte {
	file_pkg_pb_company_proto_rawDescOnce.Do(func() {
		file_pkg_pb_company_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_company_proto_rawDescData)
	})
	return file_pkg_pb_company_proto_rawDescData
}

var file_pkg_pb_company_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_pb_company_proto_goTypes = []interface{}{
	(*CreateCompanyRequest)(nil),  // 0: company.CreateCompanyRequest
	(*CreateCompanyResponse)(nil), // 1: company.CreateCompanyResponse
	(*PatchCompanyRequest)(nil),   // 2: company.PatchCompanyRequest
	(*PatchCompanyResponse)(nil),  // 3: company.PatchCompanyResponse
	(*GetOneCompanyData)(nil),     // 4: company.GetOneCompanyData
	(*GetOneCompanyRequest)(nil),  // 5: company.GetOneCompanyRequest
	(*GetOneCompanyResponse)(nil), // 6: company.GetOneCompanyResponse
	(*DeleteCompanyRequest)(nil),  // 7: company.DeleteCompanyRequest
	(*DeleteCompanyResponse)(nil), // 8: company.DeleteCompanyResponse
}
var file_pkg_pb_company_proto_depIdxs = []int32{
	4, // 0: company.GetOneCompanyResponse.data:type_name -> company.GetOneCompanyData
	0, // 1: company.CompanyService.CreateCompany:input_type -> company.CreateCompanyRequest
	2, // 2: company.CompanyService.PatchCompany:input_type -> company.PatchCompanyRequest
	5, // 3: company.CompanyService.GetOneCompany:input_type -> company.GetOneCompanyRequest
	7, // 4: company.CompanyService.DeleteCompany:input_type -> company.DeleteCompanyRequest
	1, // 5: company.CompanyService.CreateCompany:output_type -> company.CreateCompanyResponse
	3, // 6: company.CompanyService.PatchCompany:output_type -> company.PatchCompanyResponse
	6, // 7: company.CompanyService.GetOneCompany:output_type -> company.GetOneCompanyResponse
	8, // 8: company.CompanyService.DeleteCompany:output_type -> company.DeleteCompanyResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_pb_company_proto_init() }
func file_pkg_pb_company_proto_init() {
	if File_pkg_pb_company_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_company_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCompanyRequest); i {
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
		file_pkg_pb_company_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCompanyResponse); i {
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
		file_pkg_pb_company_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchCompanyRequest); i {
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
		file_pkg_pb_company_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchCompanyResponse); i {
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
		file_pkg_pb_company_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneCompanyData); i {
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
		file_pkg_pb_company_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneCompanyRequest); i {
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
		file_pkg_pb_company_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneCompanyResponse); i {
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
		file_pkg_pb_company_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCompanyRequest); i {
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
		file_pkg_pb_company_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCompanyResponse); i {
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
			RawDescriptor: file_pkg_pb_company_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_company_proto_goTypes,
		DependencyIndexes: file_pkg_pb_company_proto_depIdxs,
		MessageInfos:      file_pkg_pb_company_proto_msgTypes,
	}.Build()
	File_pkg_pb_company_proto = out.File
	file_pkg_pb_company_proto_rawDesc = nil
	file_pkg_pb_company_proto_goTypes = nil
	file_pkg_pb_company_proto_depIdxs = nil
}