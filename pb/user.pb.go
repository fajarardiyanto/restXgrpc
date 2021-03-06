// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: pb/user.proto

package pb

import (
	context "context"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fullname  string `protobuf:"bytes,2,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	CreatedAt int64  `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type UserCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *UserCreateRequest) Reset() {
	*x = UserCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateRequest) ProtoMessage() {}

func (x *UserCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateRequest.ProtoReflect.Descriptor instead.
func (*UserCreateRequest) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserCreateRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *UserCreateResponse) Reset() {
	*x = UserCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateResponse) ProtoMessage() {}

func (x *UserCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateResponse.ProtoReflect.Descriptor instead.
func (*UserCreateResponse) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserCreateResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserReadRequest) Reset() {
	*x = UserReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReadRequest) ProtoMessage() {}

func (x *UserReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReadRequest.ProtoReflect.Descriptor instead.
func (*UserReadRequest) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserReadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UserReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *UserReadResponse) Reset() {
	*x = UserReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReadResponse) ProtoMessage() {}

func (x *UserReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReadResponse.ProtoReflect.Descriptor instead.
func (*UserReadResponse) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserReadResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *UserUpdateRequest) Reset() {
	*x = UserUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdateRequest) ProtoMessage() {}

func (x *UserUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdateRequest.ProtoReflect.Descriptor instead.
func (*UserUpdateRequest) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserUpdateRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *UserUpdateResponse) Reset() {
	*x = UserUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdateResponse) ProtoMessage() {}

func (x *UserUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdateResponse.ProtoReflect.Descriptor instead.
func (*UserUpdateResponse) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserUpdateResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserDeleteRequest) Reset() {
	*x = UserDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeleteRequest) ProtoMessage() {}

func (x *UserDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeleteRequest.ProtoReflect.Descriptor instead.
func (*UserDeleteRequest) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UserDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserDeleteResponse) Reset() {
	*x = UserDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeleteResponse) ProtoMessage() {}

func (x *UserDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeleteResponse.ProtoReflect.Descriptor instead.
func (*UserDeleteResponse) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{8}
}

func (x *UserDeleteResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UserReadAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserReadAllRequest) Reset() {
	*x = UserReadAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReadAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReadAllRequest) ProtoMessage() {}

func (x *UserReadAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReadAllRequest.ProtoReflect.Descriptor instead.
func (*UserReadAllRequest) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{9}
}

type UserReadAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User []*User `protobuf:"bytes,1,rep,name=User,proto3" json:"User,omitempty"`
}

func (x *UserReadAllResponse) Reset() {
	*x = UserReadAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReadAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReadAllResponse) ProtoMessage() {}

func (x *UserReadAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReadAllResponse.ProtoReflect.Descriptor instead.
func (*UserReadAllResponse) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{10}
}

func (x *UserReadAllResponse) GetUser() []*User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_pb_user_proto protoreflect.FileDescriptor

var file_pb_user_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x22, 0x9f, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x31, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x22, 0x21, 0x0a, 0x0f,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x30, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x22, 0x31, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a,
	0x12, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x13, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x32, 0xa7,
	0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12,
	0x13, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x07,
	0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_user_proto_rawDescOnce sync.Once
	file_pb_user_proto_rawDescData = file_pb_user_proto_rawDesc
)

func file_pb_user_proto_rawDescGZIP() []byte {
	file_pb_user_proto_rawDescOnce.Do(func() {
		file_pb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_user_proto_rawDescData)
	})
	return file_pb_user_proto_rawDescData
}

var file_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pb_user_proto_goTypes = []interface{}{
	(*User)(nil),                // 0: v1.User
	(*UserCreateRequest)(nil),   // 1: v1.UserCreateRequest
	(*UserCreateResponse)(nil),  // 2: v1.UserCreateResponse
	(*UserReadRequest)(nil),     // 3: v1.UserReadRequest
	(*UserReadResponse)(nil),    // 4: v1.UserReadResponse
	(*UserUpdateRequest)(nil),   // 5: v1.UserUpdateRequest
	(*UserUpdateResponse)(nil),  // 6: v1.UserUpdateResponse
	(*UserDeleteRequest)(nil),   // 7: v1.UserDeleteRequest
	(*UserDeleteResponse)(nil),  // 8: v1.UserDeleteResponse
	(*UserReadAllRequest)(nil),  // 9: v1.UserReadAllRequest
	(*UserReadAllResponse)(nil), // 10: v1.UserReadAllResponse
}
var file_pb_user_proto_depIdxs = []int32{
	0,  // 0: v1.UserCreateRequest.User:type_name -> v1.User
	0,  // 1: v1.UserCreateResponse.User:type_name -> v1.User
	0,  // 2: v1.UserReadResponse.User:type_name -> v1.User
	0,  // 3: v1.UserUpdateRequest.User:type_name -> v1.User
	0,  // 4: v1.UserUpdateResponse.User:type_name -> v1.User
	0,  // 5: v1.UserReadAllResponse.User:type_name -> v1.User
	1,  // 6: v1.UserService.Create:input_type -> v1.UserCreateRequest
	3,  // 7: v1.UserService.Read:input_type -> v1.UserReadRequest
	5,  // 8: v1.UserService.Update:input_type -> v1.UserUpdateRequest
	7,  // 9: v1.UserService.Delete:input_type -> v1.UserDeleteRequest
	9,  // 10: v1.UserService.ReadAll:input_type -> v1.UserReadAllRequest
	2,  // 11: v1.UserService.Create:output_type -> v1.UserCreateResponse
	4,  // 12: v1.UserService.Read:output_type -> v1.UserReadResponse
	6,  // 13: v1.UserService.Update:output_type -> v1.UserUpdateResponse
	8,  // 14: v1.UserService.Delete:output_type -> v1.UserDeleteResponse
	10, // 15: v1.UserService.ReadAll:output_type -> v1.UserReadAllResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pb_user_proto_init() }
func file_pb_user_proto_init() {
	if File_pb_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_pb_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreateRequest); i {
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
		file_pb_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreateResponse); i {
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
		file_pb_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReadRequest); i {
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
		file_pb_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReadResponse); i {
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
		file_pb_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdateRequest); i {
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
		file_pb_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdateResponse); i {
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
		file_pb_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDeleteRequest); i {
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
		file_pb_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDeleteResponse); i {
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
		file_pb_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReadAllRequest); i {
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
		file_pb_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReadAllResponse); i {
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
			RawDescriptor: file_pb_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_user_proto_goTypes,
		DependencyIndexes: file_pb_user_proto_depIdxs,
		MessageInfos:      file_pb_user_proto_msgTypes,
	}.Build()
	File_pb_user_proto = out.File
	file_pb_user_proto_rawDesc = nil
	file_pb_user_proto_goTypes = nil
	file_pb_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
	Read(ctx context.Context, in *UserReadRequest, opts ...grpc.CallOption) (*UserReadResponse, error)
	Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserUpdateResponse, error)
	Delete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error)
	ReadAll(ctx context.Context, in *UserReadAllRequest, opts ...grpc.CallOption) (*UserReadAllResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	out := new(UserCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Read(ctx context.Context, in *UserReadRequest, opts ...grpc.CallOption) (*UserReadResponse, error) {
	out := new(UserReadResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserUpdateResponse, error) {
	out := new(UserUpdateResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error) {
	out := new(UserDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ReadAll(ctx context.Context, in *UserReadAllRequest, opts ...grpc.CallOption) (*UserReadAllResponse, error) {
	out := new(UserReadAllResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Create(context.Context, *UserCreateRequest) (*UserCreateResponse, error)
	Read(context.Context, *UserReadRequest) (*UserReadResponse, error)
	Update(context.Context, *UserUpdateRequest) (*UserUpdateResponse, error)
	Delete(context.Context, *UserDeleteRequest) (*UserDeleteResponse, error)
	ReadAll(context.Context, *UserReadAllRequest) (*UserReadAllResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Create(context.Context, *UserCreateRequest) (*UserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedUserServiceServer) Read(context.Context, *UserReadRequest) (*UserReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedUserServiceServer) Update(context.Context, *UserUpdateRequest) (*UserUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedUserServiceServer) Delete(context.Context, *UserDeleteRequest) (*UserDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedUserServiceServer) ReadAll(context.Context, *UserReadAllRequest) (*UserReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Read(ctx, req.(*UserReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ReadAll(ctx, req.(*UserReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _UserService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _UserService_ReadAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/user.proto",
}
