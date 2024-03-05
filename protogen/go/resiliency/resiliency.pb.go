// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: proto/resiliency/resiliency.proto

package resiliency

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ResiliencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinDelaySecond int32    `protobuf:"varint,1,opt,name=min_delay_second,proto3" json:"min_delay_second,omitempty"`
	MaxDelaySecond int32    `protobuf:"varint,2,opt,name=max_delay_second,proto3" json:"max_delay_second,omitempty"`
	StatusCodes    []uint32 `protobuf:"varint,3,rep,packed,name=status_codes,proto3" json:"status_codes,omitempty"`
}

func (x *ResiliencyRequest) Reset() {
	*x = ResiliencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_resiliency_resiliency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResiliencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResiliencyRequest) ProtoMessage() {}

func (x *ResiliencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_resiliency_resiliency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResiliencyRequest.ProtoReflect.Descriptor instead.
func (*ResiliencyRequest) Descriptor() ([]byte, []int) {
	return file_proto_resiliency_resiliency_proto_rawDescGZIP(), []int{0}
}

func (x *ResiliencyRequest) GetMinDelaySecond() int32 {
	if x != nil {
		return x.MinDelaySecond
	}
	return 0
}

func (x *ResiliencyRequest) GetMaxDelaySecond() int32 {
	if x != nil {
		return x.MaxDelaySecond
	}
	return 0
}

func (x *ResiliencyRequest) GetStatusCodes() []uint32 {
	if x != nil {
		return x.StatusCodes
	}
	return nil
}

type ResiliencyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestString string `protobuf:"bytes,1,opt,name=test_string,proto3" json:"test_string,omitempty"`
}

func (x *ResiliencyResponse) Reset() {
	*x = ResiliencyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_resiliency_resiliency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResiliencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResiliencyResponse) ProtoMessage() {}

func (x *ResiliencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_resiliency_resiliency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResiliencyResponse.ProtoReflect.Descriptor instead.
func (*ResiliencyResponse) Descriptor() ([]byte, []int) {
	return file_proto_resiliency_resiliency_proto_rawDescGZIP(), []int{1}
}

func (x *ResiliencyResponse) GetTestString() string {
	if x != nil {
		return x.TestString
	}
	return ""
}

var File_proto_resiliency_resiliency_proto protoreflect.FileDescriptor

var file_proto_resiliency_resiliency_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e,
	0x63, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x1a,
	0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x02, 0x0a, 0x11, 0x52, 0x65,
	0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x73, 0x0a, 0x10, 0x6d, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x47, 0x92, 0x41, 0x44, 0x2a, 0x10,
	0x6d, 0x69, 0x6e, 0x2d, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x2d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x32, 0x23, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x20, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x20,
	0x69, 0x6e, 0x20, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x20, 0x28, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x73, 0x69, 0x76, 0x65, 0x29, 0x4a, 0x01, 0x32, 0xa2, 0x02, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x65, 0x72, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x12, 0x73, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x65, 0x6c, 0x61,
	0x79, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x47,
	0x92, 0x41, 0x44, 0x2a, 0x10, 0x6d, 0x61, 0x78, 0x2d, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x2d, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x32, 0x23, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x20, 0x64,
	0x65, 0x6c, 0x61, 0x79, 0x20, 0x69, 0x6e, 0x20, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x20, 0x28,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x29, 0x4a, 0x01, 0x37, 0xa2, 0x02, 0x07,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x65, 0x6c,
	0x61, 0x79, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x6b, 0x0a, 0x0c, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x42,
	0x47, 0x92, 0x41, 0x44, 0x2a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2d, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0x32, 0x34, 0x41, 0x72, 0x72, 0x61, 0x79, 0x20, 0x6f, 0x66, 0x20, 0x67, 0x52, 0x50,
	0x43, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x20, 0x74,
	0x6f, 0x20, 0x62, 0x65, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x20, 0x28,
	0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x29, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x62, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69,
	0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x2a, 0x92, 0x41, 0x27, 0x2a, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x32, 0x18, 0x54, 0x65, 0x73, 0x74, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0xd8, 0x04, 0x0a, 0x11, 0x52,
	0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x72, 0x0a, 0x0f, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x6e, 0x61, 0x72, 0x79, 0x12, 0xaf, 0x01, 0x0a, 0x19, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x52,
	0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x51, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4b, 0x12, 0x49, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x7b,
	0x6d, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x7d, 0x2f, 0x7b, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x7d, 0x30, 0x01, 0x12, 0x89, 0x01, 0x0a, 0x19, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x23, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x28, 0x01, 0x12, 0x90, 0x01, 0x0a, 0x17, 0x42, 0x69, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1d,
	0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x69,
	0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x6c,
	0x69, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x12, 0x2a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x69,
	0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x28, 0x01, 0x30, 0x01, 0x32, 0xb7, 0x09, 0x0a, 0x1d, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69,
	0x65, 0x6e, 0x63, 0x79, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x1b, 0x55, 0x6e, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x57, 0x69, 0x74, 0x68, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69,
	0x65, 0x6e, 0x63, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65,
	0x6e, 0x63, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01,
	0x2a, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e,
	0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x12, 0xd8, 0x02, 0x0a, 0x25,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0xed, 0x01, 0x92, 0x41, 0xbb, 0x01, 0x12, 0x32, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x3a, 0x20, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x57,
	0x69, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x20, 0x72, 0x70, 0x63, 0x1a,
	0x35, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4a, 0x26, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x1f, 0x0a,
	0x1d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2c, 0x20, 0x32, 0x30, 0x30, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x2e, 0x2e, 0x2e, 0x4a, 0x26,
	0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x1f, 0x0a, 0x1d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2c, 0x20, 0x34, 0x30, 0x30, 0x20, 0x77, 0x68,
	0x65, 0x6e, 0x20, 0x2e, 0x2e, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22,
	0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x30, 0x01, 0x12, 0xd8, 0x02, 0x0a, 0x25, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69,
	0x65, 0x6e, 0x63, 0x79, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x52, 0x65,
	0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x52, 0x65, 0x73,
	0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xed, 0x01, 0x92, 0x41, 0xbb, 0x01, 0x12, 0x32, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x3a,
	0x20, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x20, 0x72, 0x70, 0x63, 0x1a, 0x35, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x69, 0x6c,
	0x69, 0x65, 0x6e, 0x63, 0x79, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x4a, 0x26, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x1f, 0x0a, 0x1d, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2c, 0x20, 0x32, 0x30, 0x30,
	0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x2e, 0x2e, 0x2e, 0x4a, 0x26, 0x0a, 0x03, 0x34, 0x30, 0x30,
	0x12, 0x1f, 0x0a, 0x1d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2c, 0x20, 0x34, 0x30, 0x30, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x2e, 0x2e,
	0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x28,
	0x01, 0x12, 0xdb, 0x02, 0x0a, 0x23, 0x42, 0x69, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x57, 0x69, 0x74,
	0x68, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x69,
	0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x6c,
	0x69, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xf0, 0x01, 0x92, 0x41, 0xb7, 0x01, 0x12,
	0x30, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x3a, 0x20, 0x42, 0x69, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63,
	0x79, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x20, 0x72, 0x70,
	0x63, 0x1a, 0x33, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x42, 0x69, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4a, 0x26, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x1f, 0x0a,
	0x1d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2c, 0x20, 0x32, 0x30, 0x30, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x2e, 0x2e, 0x2e, 0x4a, 0x26,
	0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x1f, 0x0a, 0x1d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2c, 0x20, 0x34, 0x30, 0x30, 0x20, 0x77, 0x68,
	0x65, 0x6e, 0x20, 0x2e, 0x2e, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x01, 0x2a, 0x22,
	0x2a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65, 0x6e, 0x63, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x81, 0x01, 0x92, 0x41, 0x48, 0x12, 0x21, 0x0a, 0x18, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x20,
	0x2d, 0x20, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x20, 0x26, 0x20, 0x67, 0x52, 0x50,
	0x43, 0x32, 0x05, 0x31, 0x2e, 0x30, 0x2e, 0x30, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68,
	0x6f, 0x73, 0x74, 0x3a, 0x38, 0x30, 0x38, 0x31, 0x2a, 0x01, 0x01, 0x3a, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x74, 0x72, 0x61, 0x6e, 0x32,
	0x39, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x6c, 0x69, 0x65,
	0x6e, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_resiliency_resiliency_proto_rawDescOnce sync.Once
	file_proto_resiliency_resiliency_proto_rawDescData = file_proto_resiliency_resiliency_proto_rawDesc
)

func file_proto_resiliency_resiliency_proto_rawDescGZIP() []byte {
	file_proto_resiliency_resiliency_proto_rawDescOnce.Do(func() {
		file_proto_resiliency_resiliency_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_resiliency_resiliency_proto_rawDescData)
	})
	return file_proto_resiliency_resiliency_proto_rawDescData
}

var file_proto_resiliency_resiliency_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_resiliency_resiliency_proto_goTypes = []interface{}{
	(*ResiliencyRequest)(nil),  // 0: resiliency.ResiliencyRequest
	(*ResiliencyResponse)(nil), // 1: resiliency.ResiliencyResponse
}
var file_proto_resiliency_resiliency_proto_depIdxs = []int32{
	0, // 0: resiliency.ResiliencyService.UnaryResiliency:input_type -> resiliency.ResiliencyRequest
	0, // 1: resiliency.ResiliencyService.ServerStreamingResiliency:input_type -> resiliency.ResiliencyRequest
	0, // 2: resiliency.ResiliencyService.ClientStreamingResiliency:input_type -> resiliency.ResiliencyRequest
	0, // 3: resiliency.ResiliencyService.BiDirectionalResiliency:input_type -> resiliency.ResiliencyRequest
	0, // 4: resiliency.ResiliencyWithMetadataService.UnaryResiliencyWithMetadata:input_type -> resiliency.ResiliencyRequest
	0, // 5: resiliency.ResiliencyWithMetadataService.ServerStreamingResiliencyWithMetadata:input_type -> resiliency.ResiliencyRequest
	0, // 6: resiliency.ResiliencyWithMetadataService.ClientStreamingResiliencyWithMetadata:input_type -> resiliency.ResiliencyRequest
	0, // 7: resiliency.ResiliencyWithMetadataService.BiDirectionalResiliencyWithMetadata:input_type -> resiliency.ResiliencyRequest
	1, // 8: resiliency.ResiliencyService.UnaryResiliency:output_type -> resiliency.ResiliencyResponse
	1, // 9: resiliency.ResiliencyService.ServerStreamingResiliency:output_type -> resiliency.ResiliencyResponse
	1, // 10: resiliency.ResiliencyService.ClientStreamingResiliency:output_type -> resiliency.ResiliencyResponse
	1, // 11: resiliency.ResiliencyService.BiDirectionalResiliency:output_type -> resiliency.ResiliencyResponse
	1, // 12: resiliency.ResiliencyWithMetadataService.UnaryResiliencyWithMetadata:output_type -> resiliency.ResiliencyResponse
	1, // 13: resiliency.ResiliencyWithMetadataService.ServerStreamingResiliencyWithMetadata:output_type -> resiliency.ResiliencyResponse
	1, // 14: resiliency.ResiliencyWithMetadataService.ClientStreamingResiliencyWithMetadata:output_type -> resiliency.ResiliencyResponse
	1, // 15: resiliency.ResiliencyWithMetadataService.BiDirectionalResiliencyWithMetadata:output_type -> resiliency.ResiliencyResponse
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_resiliency_resiliency_proto_init() }
func file_proto_resiliency_resiliency_proto_init() {
	if File_proto_resiliency_resiliency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_resiliency_resiliency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResiliencyRequest); i {
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
		file_proto_resiliency_resiliency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResiliencyResponse); i {
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
			RawDescriptor: file_proto_resiliency_resiliency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_resiliency_resiliency_proto_goTypes,
		DependencyIndexes: file_proto_resiliency_resiliency_proto_depIdxs,
		MessageInfos:      file_proto_resiliency_resiliency_proto_msgTypes,
	}.Build()
	File_proto_resiliency_resiliency_proto = out.File
	file_proto_resiliency_resiliency_proto_rawDesc = nil
	file_proto_resiliency_resiliency_proto_goTypes = nil
	file_proto_resiliency_resiliency_proto_depIdxs = nil
}
