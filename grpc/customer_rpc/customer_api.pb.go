// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: grpc/customer_rpc/customer_api.proto

package customer_rpc

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

type GetCustomerByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCustomerByIdRequest) Reset() {
	*x = GetCustomerByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_customer_rpc_customer_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerByIdRequest) ProtoMessage() {}

func (x *GetCustomerByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_customer_rpc_customer_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerByIdRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerByIdRequest) Descriptor() ([]byte, []int) {
	return file_grpc_customer_rpc_customer_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetCustomerByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCustomerByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCustomerByIdResponse) Reset() {
	*x = GetCustomerByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_customer_rpc_customer_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerByIdResponse) ProtoMessage() {}

func (x *GetCustomerByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_customer_rpc_customer_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerByIdResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerByIdResponse) Descriptor() ([]byte, []int) {
	return file_grpc_customer_rpc_customer_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetCustomerByIdResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GetCustomerByIdResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCustomerByIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Ids    []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetCustomerByIdsResponse) Reset() {
	*x = GetCustomerByIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_customer_rpc_customer_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerByIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerByIdsResponse) ProtoMessage() {}

func (x *GetCustomerByIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_customer_rpc_customer_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerByIdsResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerByIdsResponse) Descriptor() ([]byte, []int) {
	return file_grpc_customer_rpc_customer_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetCustomerByIdsResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GetCustomerByIdsResponse) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetCustomerInterestsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCustomerInterestsRequest) Reset() {
	*x = GetCustomerInterestsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_customer_rpc_customer_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerInterestsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerInterestsRequest) ProtoMessage() {}

func (x *GetCustomerInterestsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_customer_rpc_customer_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerInterestsRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerInterestsRequest) Descriptor() ([]byte, []int) {
	return file_grpc_customer_rpc_customer_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetCustomerInterestsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCustomerInterestsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interest string `protobuf:"bytes,1,opt,name=interest,proto3" json:"interest,omitempty"`
}

func (x *GetCustomerInterestsResponse) Reset() {
	*x = GetCustomerInterestsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_customer_rpc_customer_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerInterestsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerInterestsResponse) ProtoMessage() {}

func (x *GetCustomerInterestsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_customer_rpc_customer_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerInterestsResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerInterestsResponse) Descriptor() ([]byte, []int) {
	return file_grpc_customer_rpc_customer_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetCustomerInterestsResponse) GetInterest() string {
	if x != nil {
		return x.Interest
	}
	return ""
}

var File_grpc_customer_rpc_customer_api_proto protoreflect.FileDescriptor

var file_grpc_customer_rpc_customer_api_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x72, 0x70, 0x63, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x44, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x2d, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x1c, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x32, 0xdd, 0x03, 0x0a, 0x0b, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x41, 0x50, 0x49, 0x12, 0x68, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x29, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x6c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x29, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12,
	0x79, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73, 0x12, 0x2e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x7b, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73,
	0x74, 0x73, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x29, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_customer_rpc_customer_api_proto_rawDescOnce sync.Once
	file_grpc_customer_rpc_customer_api_proto_rawDescData = file_grpc_customer_rpc_customer_api_proto_rawDesc
)

func file_grpc_customer_rpc_customer_api_proto_rawDescGZIP() []byte {
	file_grpc_customer_rpc_customer_api_proto_rawDescOnce.Do(func() {
		file_grpc_customer_rpc_customer_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_customer_rpc_customer_api_proto_rawDescData)
	})
	return file_grpc_customer_rpc_customer_api_proto_rawDescData
}

var file_grpc_customer_rpc_customer_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_customer_rpc_customer_api_proto_goTypes = []any{
	(*GetCustomerByIdRequest)(nil),       // 0: grpc.customer_rpc.GetCustomerByIdRequest
	(*GetCustomerByIdResponse)(nil),      // 1: grpc.customer_rpc.GetCustomerByIdResponse
	(*GetCustomerByIdsResponse)(nil),     // 2: grpc.customer_rpc.GetCustomerByIdsResponse
	(*GetCustomerInterestsRequest)(nil),  // 3: grpc.customer_rpc.GetCustomerInterestsRequest
	(*GetCustomerInterestsResponse)(nil), // 4: grpc.customer_rpc.GetCustomerInterestsResponse
}
var file_grpc_customer_rpc_customer_api_proto_depIdxs = []int32{
	0, // 0: grpc.customer_rpc.CustomerAPI.GetCustomerById:input_type -> grpc.customer_rpc.GetCustomerByIdRequest
	0, // 1: grpc.customer_rpc.CustomerAPI.GetCustomerByIds:input_type -> grpc.customer_rpc.GetCustomerByIdRequest
	3, // 2: grpc.customer_rpc.CustomerAPI.GetCustomerInterests:input_type -> grpc.customer_rpc.GetCustomerInterestsRequest
	0, // 3: grpc.customer_rpc.CustomerAPI.GetCustomerInterestsByIds:input_type -> grpc.customer_rpc.GetCustomerByIdRequest
	1, // 4: grpc.customer_rpc.CustomerAPI.GetCustomerById:output_type -> grpc.customer_rpc.GetCustomerByIdResponse
	2, // 5: grpc.customer_rpc.CustomerAPI.GetCustomerByIds:output_type -> grpc.customer_rpc.GetCustomerByIdsResponse
	4, // 6: grpc.customer_rpc.CustomerAPI.GetCustomerInterests:output_type -> grpc.customer_rpc.GetCustomerInterestsResponse
	4, // 7: grpc.customer_rpc.CustomerAPI.GetCustomerInterestsByIds:output_type -> grpc.customer_rpc.GetCustomerInterestsResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_customer_rpc_customer_api_proto_init() }
func file_grpc_customer_rpc_customer_api_proto_init() {
	if File_grpc_customer_rpc_customer_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_customer_rpc_customer_api_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetCustomerByIdRequest); i {
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
		file_grpc_customer_rpc_customer_api_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetCustomerByIdResponse); i {
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
		file_grpc_customer_rpc_customer_api_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetCustomerByIdsResponse); i {
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
		file_grpc_customer_rpc_customer_api_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetCustomerInterestsRequest); i {
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
		file_grpc_customer_rpc_customer_api_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetCustomerInterestsResponse); i {
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
			RawDescriptor: file_grpc_customer_rpc_customer_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_customer_rpc_customer_api_proto_goTypes,
		DependencyIndexes: file_grpc_customer_rpc_customer_api_proto_depIdxs,
		MessageInfos:      file_grpc_customer_rpc_customer_api_proto_msgTypes,
	}.Build()
	File_grpc_customer_rpc_customer_api_proto = out.File
	file_grpc_customer_rpc_customer_api_proto_rawDesc = nil
	file_grpc_customer_rpc_customer_api_proto_goTypes = nil
	file_grpc_customer_rpc_customer_api_proto_depIdxs = nil
}
