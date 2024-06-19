// define syntax used in proto file

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: device/v3/device.proto

package v3

import (
	_ "github.com/osguydch/diosproc/proto/google/api"
	_ "github.com/osguydch/diosproc/proto/protoc-gen-openapiv2/options"
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

// simple message
type OpenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceUri   string `protobuf:"bytes,1,opt,name=deviceUri,proto3" json:"deviceUri,omitempty"`
	DeviceGroup string `protobuf:"bytes,2,opt,name=deviceGroup,proto3" json:"deviceGroup,omitempty"`
}

func (x *OpenRequest) Reset() {
	*x = OpenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_v3_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenRequest) ProtoMessage() {}

func (x *OpenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_device_v3_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenRequest.ProtoReflect.Descriptor instead.
func (*OpenRequest) Descriptor() ([]byte, []int) {
	return file_device_v3_device_proto_rawDescGZIP(), []int{0}
}

func (x *OpenRequest) GetDeviceUri() string {
	if x != nil {
		return x.DeviceUri
	}
	return ""
}

func (x *OpenRequest) GetDeviceGroup() string {
	if x != nil {
		return x.DeviceGroup
	}
	return ""
}

type OpenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	RetCode    string `protobuf:"bytes,2,opt,name=retCode,proto3" json:"retCode,omitempty"`
	RetContext string `protobuf:"bytes,3,opt,name=retContext,proto3" json:"retContext,omitempty"`
}

func (x *OpenReply) Reset() {
	*x = OpenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_v3_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenReply) ProtoMessage() {}

func (x *OpenReply) ProtoReflect() protoreflect.Message {
	mi := &file_device_v3_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenReply.ProtoReflect.Descriptor instead.
func (*OpenReply) Descriptor() ([]byte, []int) {
	return file_device_v3_device_proto_rawDescGZIP(), []int{1}
}

func (x *OpenReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OpenReply) GetRetCode() string {
	if x != nil {
		return x.RetCode
	}
	return ""
}

func (x *OpenReply) GetRetContext() string {
	if x != nil {
		return x.RetContext
	}
	return ""
}

type CloseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	RetCode string `protobuf:"bytes,2,opt,name=retCode,proto3" json:"retCode,omitempty"`
}

func (x *CloseReply) Reset() {
	*x = CloseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_v3_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseReply) ProtoMessage() {}

func (x *CloseReply) ProtoReflect() protoreflect.Message {
	mi := &file_device_v3_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseReply.ProtoReflect.Descriptor instead.
func (*CloseReply) Descriptor() ([]byte, []int) {
	return file_device_v3_device_proto_rawDescGZIP(), []int{2}
}

func (x *CloseReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CloseReply) GetRetCode() string {
	if x != nil {
		return x.RetCode
	}
	return ""
}

type DoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceUri      string `protobuf:"bytes,1,opt,name=deviceUri,proto3" json:"deviceUri,omitempty"`
	TypeName       string `protobuf:"bytes,2,opt,name=typeName,proto3" json:"typeName,omitempty"`
	ReqName        string `protobuf:"bytes,3,opt,name=reqName,proto3" json:"reqName,omitempty"`
	ReqParam       string `protobuf:"bytes,4,opt,name=reqParam,proto3" json:"reqParam,omitempty"`
	ReqTransaction string `protobuf:"bytes,5,opt,name=reqTransaction,proto3" json:"reqTransaction,omitempty"`
	ReqClientID    string `protobuf:"bytes,6,opt,name=reqClientID,proto3" json:"reqClientID,omitempty"`
	ReqTimeStamp   string `protobuf:"bytes,7,opt,name=reqTimeStamp,proto3" json:"reqTimeStamp,omitempty"`
}

func (x *DoRequest) Reset() {
	*x = DoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_v3_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoRequest) ProtoMessage() {}

func (x *DoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_device_v3_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoRequest.ProtoReflect.Descriptor instead.
func (*DoRequest) Descriptor() ([]byte, []int) {
	return file_device_v3_device_proto_rawDescGZIP(), []int{3}
}

func (x *DoRequest) GetDeviceUri() string {
	if x != nil {
		return x.DeviceUri
	}
	return ""
}

func (x *DoRequest) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *DoRequest) GetReqName() string {
	if x != nil {
		return x.ReqName
	}
	return ""
}

func (x *DoRequest) GetReqParam() string {
	if x != nil {
		return x.ReqParam
	}
	return ""
}

func (x *DoRequest) GetReqTransaction() string {
	if x != nil {
		return x.ReqTransaction
	}
	return ""
}

func (x *DoRequest) GetReqClientID() string {
	if x != nil {
		return x.ReqClientID
	}
	return ""
}

func (x *DoRequest) GetReqTimeStamp() string {
	if x != nil {
		return x.ReqTimeStamp
	}
	return ""
}

type DoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceUri      string `protobuf:"bytes,1,opt,name=deviceUri,proto3" json:"deviceUri,omitempty"`
	TypeName       string `protobuf:"bytes,2,opt,name=typeName,proto3" json:"typeName,omitempty"`
	ReqName        string `protobuf:"bytes,3,opt,name=reqName,proto3" json:"reqName,omitempty"`
	RespParam      string `protobuf:"bytes,4,opt,name=RespParam,proto3" json:"RespParam,omitempty"`
	ReqTransaction string `protobuf:"bytes,5,opt,name=reqTransaction,proto3" json:"reqTransaction,omitempty"`
	ReqClientID    string `protobuf:"bytes,6,opt,name=reqClientID,proto3" json:"reqClientID,omitempty"`
	RespTimeStamp  string `protobuf:"bytes,7,opt,name=RespTimeStamp,proto3" json:"RespTimeStamp,omitempty"`
	RetContext     string `protobuf:"bytes,8,opt,name=retContext,proto3" json:"retContext,omitempty"`
}

func (x *DoResponse) Reset() {
	*x = DoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_v3_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoResponse) ProtoMessage() {}

func (x *DoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_device_v3_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoResponse.ProtoReflect.Descriptor instead.
func (*DoResponse) Descriptor() ([]byte, []int) {
	return file_device_v3_device_proto_rawDescGZIP(), []int{4}
}

func (x *DoResponse) GetDeviceUri() string {
	if x != nil {
		return x.DeviceUri
	}
	return ""
}

func (x *DoResponse) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *DoResponse) GetReqName() string {
	if x != nil {
		return x.ReqName
	}
	return ""
}

func (x *DoResponse) GetRespParam() string {
	if x != nil {
		return x.RespParam
	}
	return ""
}

func (x *DoResponse) GetReqTransaction() string {
	if x != nil {
		return x.ReqTransaction
	}
	return ""
}

func (x *DoResponse) GetReqClientID() string {
	if x != nil {
		return x.ReqClientID
	}
	return ""
}

func (x *DoResponse) GetRespTimeStamp() string {
	if x != nil {
		return x.RespTimeStamp
	}
	return ""
}

func (x *DoResponse) GetRetContext() string {
	if x != nil {
		return x.RetContext
	}
	return ""
}

var File_device_v3_device_proto protoreflect.FileDescriptor

var file_device_v3_device_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x56, 0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4d, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x72, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x72, 0x69, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x22, 0x5f, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x22, 0x40, 0x0a, 0x0a, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0xe9, 0x01, 0x0a, 0x09, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x72, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x72, 0x69, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65,
	0x71, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x71, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x71, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x72,
	0x65, 0x71, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x8e, 0x02, 0x0a, 0x0a, 0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x72, 0x69, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x73, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x71, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65,
	0x73, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x32, 0xaa, 0x0b, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x92, 0x01, 0x0a, 0x04,
	0x4f, 0x70, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x33,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x33, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x5c, 0x92, 0x41, 0x33, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x0d, 0x4f, 0x70, 0x65, 0x6e, 0x20, 0x61, 0x20, 0x44, 0x65, 0x69, 0x76, 0x63, 0x65, 0x1a, 0x1a,
	0x4f, 0x70, 0x65, 0x6e, 0x20, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x20, 0x66, 0x72, 0x6f, 0x6d,
	0x20, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20,
	0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x4f, 0x70, 0x65, 0x6e, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x72, 0x69, 0x7d,
	0x12, 0x9c, 0x01, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x16, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x56, 0x33, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x33, 0x2e, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x64, 0x92, 0x41, 0x3a, 0x0a, 0x06,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x13, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x20, 0x4f, 0x70,
	0x65, 0x6e, 0x65, 0x64, 0x20, 0x44, 0x65, 0x69, 0x76, 0x63, 0x65, 0x1a, 0x1b, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x20, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01,
	0x2a, 0x22, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x72, 0x69, 0x7d, 0x12,
	0xab, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x56, 0x33, 0x2e, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x33, 0x2e, 0x44, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x77, 0x92, 0x41, 0x56, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x20, 0x47, 0x65, 0x74, 0x20, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x44, 0x65, 0x69,
	0x76, 0x63, 0x65, 0x1a, 0x2a, 0x47, 0x65, 0x74, 0x20, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x44,
	0x65, 0x69, 0x76, 0x63, 0x65, 0x20, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x72, 0x69, 0x7d, 0x12, 0xc3, 0x01,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x56, 0x33, 0x2e, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x33, 0x2e, 0x44, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8b, 0x01, 0x92, 0x41, 0x67, 0x0a, 0x06, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x27, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x28, 0x53, 0x65, 0x74, 0x29,
	0x20, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x20, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x44, 0x65, 0x69, 0x76, 0x63, 0x65, 0x1a, 0x34, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x28, 0x53, 0x65, 0x74, 0x29, 0x20, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x44,
	0x65, 0x69, 0x76, 0x63, 0x65, 0x20, 0x61, 0x74, 0x20, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x1a, 0x16, 0x2f, 0x76, 0x31,
	0x2f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55,
	0x72, 0x69, 0x7d, 0x12, 0xc8, 0x01, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x14, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x33, 0x2e, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x33, 0x2e, 0x44, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x93, 0x01, 0x92, 0x41, 0x6f, 0x0a, 0x06,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x41, 0x64, 0x64, 0x20, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x44, 0x65, 0x69, 0x76, 0x63, 0x65, 0x1a, 0x44, 0x41, 0x64, 0x64, 0x20, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x44, 0x65, 0x69, 0x76, 0x63, 0x65, 0x20, 0x61, 0x74, 0x20, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x20, 0x69, 0x66, 0x20, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x20, 0x63, 0x61, 0x6e, 0x62, 0x65, 0x20, 0x61, 0x64, 0x64, 0x65, 0x64, 0x2e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x72, 0x69, 0x7d, 0x12, 0xcd,
	0x01, 0x0a, 0x03, 0x44, 0x65, 0x6c, 0x12, 0x14, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x56, 0x33, 0x2e, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x33, 0x2e, 0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x98, 0x01, 0x92, 0x41, 0x77, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x22, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x44,
	0x65, 0x69, 0x76, 0x63, 0x65, 0x1a, 0x49, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x44, 0x65, 0x69, 0x76, 0x63, 0x65, 0x20, 0x61, 0x74, 0x20, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x69, 0x66, 0x20, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x20, 0x63, 0x61, 0x6e, 0x62, 0x65, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x72, 0x69, 0x7d, 0x12, 0xb0,
	0x01, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x56, 0x33, 0x2e, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x33, 0x2e, 0x44, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x79, 0x92, 0x41, 0x4e, 0x0a, 0x06, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x1a, 0x45, 0x78, 0x65, 0x63, 0x75, 0x61, 0x74, 0x65, 0x20, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x44, 0x65, 0x69, 0x76, 0x63, 0x65, 0x1a,
	0x28, 0x45, 0x78, 0x65, 0x63, 0x75, 0x61, 0x74, 0x65, 0x20, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x44, 0x65, 0x69, 0x76, 0x63, 0x65, 0x20, 0x61, 0x74, 0x20, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a,
	0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x72, 0x69,
	0x7d, 0x12, 0xa9, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x33, 0x2e, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x33, 0x2e,
	0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x71, 0x92, 0x41, 0x45, 0x0a,
	0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x20, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x44, 0x65, 0x69, 0x76, 0x63, 0x65, 0x1a,
	0x23, 0x53, 0x65, 0x6e, 0x64, 0x20, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x74, 0x6f,
	0x20, 0x44, 0x65, 0x69, 0x76, 0x63, 0x65, 0x20, 0x61, 0x74, 0x20, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f,
	0x76, 0x31, 0x2f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x72, 0x69, 0x7d, 0x42, 0xe9, 0x01,
	0x92, 0x41, 0x57, 0x12, 0x05, 0x32, 0x03, 0x33, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x72, 0x4b, 0x0a,
	0x23, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x20, 0x62, 0x6f,
	0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x24, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x73, 0x67, 0x75, 0x79, 0x64, 0x63,
	0x68, 0x2f, 0x64, 0x69, 0x6f, 0x73, 0x70, 0x72, 0x6f, 0x63, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x33, 0x42, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x73, 0x67, 0x75, 0x79, 0x64, 0x63, 0x68, 0x2f, 0x64, 0x69,
	0x6f, 0x73, 0x70, 0x72, 0x6f, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x33, 0xa2, 0x02, 0x03, 0x44, 0x56, 0x58, 0xaa, 0x02, 0x09, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x33, 0xca, 0x02, 0x09, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5c, 0x56, 0x33, 0xe2, 0x02, 0x15, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x33,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_device_v3_device_proto_rawDescOnce sync.Once
	file_device_v3_device_proto_rawDescData = file_device_v3_device_proto_rawDesc
)

func file_device_v3_device_proto_rawDescGZIP() []byte {
	file_device_v3_device_proto_rawDescOnce.Do(func() {
		file_device_v3_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_device_v3_device_proto_rawDescData)
	})
	return file_device_v3_device_proto_rawDescData
}

var file_device_v3_device_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_device_v3_device_proto_goTypes = []interface{}{
	(*OpenRequest)(nil), // 0: Device.V3.OpenRequest
	(*OpenReply)(nil),   // 1: Device.V3.OpenReply
	(*CloseReply)(nil),  // 2: Device.V3.CloseReply
	(*DoRequest)(nil),   // 3: Device.V3.DoRequest
	(*DoResponse)(nil),  // 4: Device.V3.DoResponse
}
var file_device_v3_device_proto_depIdxs = []int32{
	0, // 0: Device.V3.Device.Open:input_type -> Device.V3.OpenRequest
	0, // 1: Device.V3.Device.Close:input_type -> Device.V3.OpenRequest
	3, // 2: Device.V3.Device.Get:input_type -> Device.V3.DoRequest
	3, // 3: Device.V3.Device.Update:input_type -> Device.V3.DoRequest
	3, // 4: Device.V3.Device.Add:input_type -> Device.V3.DoRequest
	3, // 5: Device.V3.Device.Del:input_type -> Device.V3.DoRequest
	3, // 6: Device.V3.Device.Action:input_type -> Device.V3.DoRequest
	3, // 7: Device.V3.Device.Message:input_type -> Device.V3.DoRequest
	1, // 8: Device.V3.Device.Open:output_type -> Device.V3.OpenReply
	2, // 9: Device.V3.Device.Close:output_type -> Device.V3.CloseReply
	4, // 10: Device.V3.Device.Get:output_type -> Device.V3.DoResponse
	4, // 11: Device.V3.Device.Update:output_type -> Device.V3.DoResponse
	4, // 12: Device.V3.Device.Add:output_type -> Device.V3.DoResponse
	4, // 13: Device.V3.Device.Del:output_type -> Device.V3.DoResponse
	4, // 14: Device.V3.Device.Action:output_type -> Device.V3.DoResponse
	4, // 15: Device.V3.Device.Message:output_type -> Device.V3.DoResponse
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_device_v3_device_proto_init() }
func file_device_v3_device_proto_init() {
	if File_device_v3_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_device_v3_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenRequest); i {
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
		file_device_v3_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenReply); i {
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
		file_device_v3_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseReply); i {
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
		file_device_v3_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoRequest); i {
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
		file_device_v3_device_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoResponse); i {
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
			RawDescriptor: file_device_v3_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_device_v3_device_proto_goTypes,
		DependencyIndexes: file_device_v3_device_proto_depIdxs,
		MessageInfos:      file_device_v3_device_proto_msgTypes,
	}.Build()
	File_device_v3_device_proto = out.File
	file_device_v3_device_proto_rawDesc = nil
	file_device_v3_device_proto_goTypes = nil
	file_device_v3_device_proto_depIdxs = nil
}
