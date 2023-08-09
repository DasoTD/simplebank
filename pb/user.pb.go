// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Username             string               `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FullName             string               `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PasswordChangedAt    *timestamp.Timestamp `protobuf:"bytes,4,opt,name=password_changed_at,json=passwordChangedAt,proto3" json:"password_changed_at,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPasswordChangedAt() *timestamp.Timestamp {
	if m != nil {
		return m.PasswordChangedAt
	}
	return nil
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8e, 0x31, 0x4f, 0xc3, 0x30,
	0x14, 0x84, 0x95, 0xd2, 0x22, 0xf2, 0x98, 0x30, 0x0c, 0x51, 0x18, 0x5a, 0x98, 0x3a, 0xd9, 0x12,
	0x4c, 0x8c, 0x85, 0x8d, 0x81, 0xa1, 0x82, 0x85, 0x25, 0x7a, 0x4e, 0x5e, 0xd3, 0x08, 0x3b, 0xb6,
	0xec, 0x17, 0xf1, 0x67, 0xf9, 0x31, 0x28, 0x31, 0x61, 0xed, 0x78, 0xf7, 0xdd, 0x9d, 0x0e, 0x60,
	0x88, 0x14, 0xa4, 0x0f, 0x8e, 0x9d, 0x58, 0x78, 0x5d, 0xae, 0x5b, 0xe7, 0x5a, 0x43, 0x6a, 0x72,
	0xf4, 0x70, 0x50, 0xdc, 0x59, 0x8a, 0x8c, 0xd6, 0xa7, 0xd0, 0xfd, 0x4f, 0x06, 0xcb, 0x8f, 0x48,
	0x41, 0x94, 0x70, 0x31, 0x76, 0x7b, 0xb4, 0x54, 0x64, 0x9b, 0x6c, 0x9b, 0xef, 0xff, 0xb5, 0xb8,
	0x85, 0xfc, 0x30, 0x18, 0x53, 0x4d, 0x70, 0x91, 0xe0, 0x68, 0xbc, 0x8d, 0xf0, 0x06, 0x56, 0x64,
	0xb1, 0x33, 0xc5, 0xd9, 0x04, 0x92, 0x10, 0xaf, 0x70, 0xed, 0x31, 0xc6, 0x6f, 0x17, 0x9a, 0xaa,
	0x3e, 0x62, 0xdf, 0x52, 0x53, 0x21, 0x17, 0xcb, 0x4d, 0xb6, 0xbd, 0x7c, 0x28, 0x65, 0xba, 0x25,
	0xe7, 0x5b, 0xf2, 0x7d, 0xbe, 0xb5, 0xbf, 0x9a, 0x6b, 0x2f, 0xa9, 0xb5, 0x63, 0xf1, 0x04, 0x50,
	0x07, 0x42, 0x4e, 0x13, 0xab, 0x93, 0x13, 0xf9, 0x5f, 0x7a, 0xc7, 0xcf, 0x77, 0x9f, 0xeb, 0xb6,
	0xe3, 0xe3, 0xa0, 0x65, 0xed, 0xac, 0x6a, 0x30, 0x3a, 0x6e, 0x54, 0xec, 0xac, 0x37, 0xa4, 0xb1,
	0xff, 0x52, 0x5e, 0xeb, 0xf3, 0x69, 0xe1, 0xf1, 0x37, 0x00, 0x00, 0xff, 0xff, 0xee, 0x0f, 0x58,
	0x18, 0x3b, 0x01, 0x00, 0x00,
}
