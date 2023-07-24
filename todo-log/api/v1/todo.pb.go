// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: todo-log/api/v1/todo.proto

package api_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusType int32

const (
	StatusType_PHONE_TYPE_TODO  StatusType = 0
	StatusType_PHONE_TYPE_DOING StatusType = 1
	StatusType_PHONE_TYPE_DONE  StatusType = 2
)

// Enum value maps for StatusType.
var (
	StatusType_name = map[int32]string{
		0: "PHONE_TYPE_TODO",
		1: "PHONE_TYPE_DOING",
		2: "PHONE_TYPE_DONE",
	}
	StatusType_value = map[string]int32{
		"PHONE_TYPE_TODO":  0,
		"PHONE_TYPE_DOING": 1,
		"PHONE_TYPE_DONE":  2,
	}
)

func (x StatusType) Enum() *StatusType {
	p := new(StatusType)
	*p = x
	return p
}

func (x StatusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusType) Descriptor() protoreflect.EnumDescriptor {
	return file_todo_log_api_v1_todo_proto_enumTypes[0].Descriptor()
}

func (StatusType) Type() protoreflect.EnumType {
	return &file_todo_log_api_v1_todo_proto_enumTypes[0]
}

func (x StatusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusType.Descriptor instead.
func (StatusType) EnumDescriptor() ([]byte, []int) {
	return file_todo_log_api_v1_todo_proto_rawDescGZIP(), []int{0}
}

type RetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RetrieveRequest) Reset() {
	*x = RetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_log_api_v1_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveRequest) ProtoMessage() {}

func (x *RetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_log_api_v1_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveRequest.ProtoReflect.Descriptor instead.
func (*RetrieveRequest) Descriptor() ([]byte, []int) {
	return file_todo_log_api_v1_todo_proto_rawDescGZIP(), []int{0}
}

func (x *RetrieveRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string     `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string     `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Status      StatusType `protobuf:"varint,5,opt,name=status,proto3,enum=api.v1.StatusType" json:"status,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_log_api_v1_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_todo_log_api_v1_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_todo_log_api_v1_todo_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Response) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Response) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Response) GetStatus() StatusType {
	if x != nil {
		return x.Status
	}
	return StatusType_PHONE_TYPE_TODO
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_log_api_v1_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_log_api_v1_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_todo_log_api_v1_todo_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Title       string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Status      StatusType             `protobuf:"varint,5,opt,name=status,proto3,enum=api.v1.StatusType" json:"status,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_log_api_v1_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_todo_log_api_v1_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_todo_log_api_v1_todo_proto_rawDescGZIP(), []int{3}
}

func (x *Todo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Todo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Todo) GetStatus() StatusType {
	if x != nil {
		return x.Status
	}
	return StatusType_PHONE_TYPE_TODO
}

type Activities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*Todo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *Activities) Reset() {
	*x = Activities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_log_api_v1_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activities) ProtoMessage() {}

func (x *Activities) ProtoReflect() protoreflect.Message {
	mi := &file_todo_log_api_v1_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activities.ProtoReflect.Descriptor instead.
func (*Activities) Descriptor() ([]byte, []int) {
	return file_todo_log_api_v1_todo_proto_rawDescGZIP(), []int{4}
}

func (x *Activities) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

type TodoQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *TodoQuery) Reset() {
	*x = TodoQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_log_api_v1_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoQuery) ProtoMessage() {}

func (x *TodoQuery) ProtoReflect() protoreflect.Message {
	mi := &file_todo_log_api_v1_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoQuery.ProtoReflect.Descriptor instead.
func (*TodoQuery) Descriptor() ([]byte, []int) {
	return file_todo_log_api_v1_todo_proto_rawDescGZIP(), []int{5}
}

func (x *TodoQuery) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_todo_log_api_v1_todo_proto protoreflect.FileDescriptor

var file_todo_log_api_v1_todo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x25, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22,
	0xb5, 0x01, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x30, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x22, 0x23, 0x0a, 0x09, 0x54, 0x6f, 0x64,
	0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x2a, 0x4c,
	0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f,
	0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x4f, 0x44, 0x4f, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x44, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x48, 0x4f, 0x4e, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x32, 0xcc, 0x01, 0x0a,
	0x08, 0x54, 0x6f, 0x64, 0x6f, 0x5f, 0x4c, 0x6f, 0x67, 0x12, 0x2a, 0x0a, 0x06, 0x49, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x1a, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2c, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x61, 0x74, 0x68, 0x69, 0x65, 0x66, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x6c, 0x6f, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_log_api_v1_todo_proto_rawDescOnce sync.Once
	file_todo_log_api_v1_todo_proto_rawDescData = file_todo_log_api_v1_todo_proto_rawDesc
)

func file_todo_log_api_v1_todo_proto_rawDescGZIP() []byte {
	file_todo_log_api_v1_todo_proto_rawDescOnce.Do(func() {
		file_todo_log_api_v1_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_log_api_v1_todo_proto_rawDescData)
	})
	return file_todo_log_api_v1_todo_proto_rawDescData
}

var file_todo_log_api_v1_todo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_todo_log_api_v1_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_todo_log_api_v1_todo_proto_goTypes = []interface{}{
	(StatusType)(0),               // 0: api.v1.StatusType
	(*RetrieveRequest)(nil),       // 1: api.v1.RetrieveRequest
	(*Response)(nil),              // 2: api.v1.Response
	(*ListRequest)(nil),           // 3: api.v1.ListRequest
	(*Todo)(nil),                  // 4: api.v1.Todo
	(*Activities)(nil),            // 5: api.v1.Activities
	(*TodoQuery)(nil),             // 6: api.v1.TodoQuery
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_todo_log_api_v1_todo_proto_depIdxs = []int32{
	0, // 0: api.v1.Response.status:type_name -> api.v1.StatusType
	7, // 1: api.v1.Todo.created_at:type_name -> google.protobuf.Timestamp
	0, // 2: api.v1.Todo.status:type_name -> api.v1.StatusType
	4, // 3: api.v1.Activities.todos:type_name -> api.v1.Todo
	4, // 4: api.v1.Todo_Log.Insert:input_type -> api.v1.Todo
	1, // 5: api.v1.Todo_Log.Retrieve:input_type -> api.v1.RetrieveRequest
	3, // 6: api.v1.Todo_Log.List:input_type -> api.v1.ListRequest
	4, // 7: api.v1.Todo_Log.Update:input_type -> api.v1.Todo
	2, // 8: api.v1.Todo_Log.Insert:output_type -> api.v1.Response
	4, // 9: api.v1.Todo_Log.Retrieve:output_type -> api.v1.Todo
	5, // 10: api.v1.Todo_Log.List:output_type -> api.v1.Activities
	5, // 11: api.v1.Todo_Log.Update:output_type -> api.v1.Activities
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_todo_log_api_v1_todo_proto_init() }
func file_todo_log_api_v1_todo_proto_init() {
	if File_todo_log_api_v1_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_log_api_v1_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveRequest); i {
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
		file_todo_log_api_v1_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_todo_log_api_v1_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_todo_log_api_v1_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_todo_log_api_v1_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activities); i {
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
		file_todo_log_api_v1_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoQuery); i {
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
			RawDescriptor: file_todo_log_api_v1_todo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_log_api_v1_todo_proto_goTypes,
		DependencyIndexes: file_todo_log_api_v1_todo_proto_depIdxs,
		EnumInfos:         file_todo_log_api_v1_todo_proto_enumTypes,
		MessageInfos:      file_todo_log_api_v1_todo_proto_msgTypes,
	}.Build()
	File_todo_log_api_v1_todo_proto = out.File
	file_todo_log_api_v1_todo_proto_rawDesc = nil
	file_todo_log_api_v1_todo_proto_goTypes = nil
	file_todo_log_api_v1_todo_proto_depIdxs = nil
}