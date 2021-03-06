// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userinfo.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type UserInfo struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	NickName             string   `protobuf:"bytes,2,opt,name=NickName,proto3" json:"NickName,omitempty"`
	Signature            string   `protobuf:"bytes,3,opt,name=Signature,proto3" json:"Signature,omitempty"`
	HeadImageUrl         string   `protobuf:"bytes,4,opt,name=HeadImageUrl,proto3" json:"HeadImageUrl,omitempty"`
	CreateTime           int64    `protobuf:"varint,5,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	Status               int32    `protobuf:"varint,6,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_785a78c34699a93d, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UserInfo) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *UserInfo) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *UserInfo) GetHeadImageUrl() string {
	if m != nil {
		return m.HeadImageUrl
	}
	return ""
}

func (m *UserInfo) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *UserInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "pb.UserInfo")
}

func init() { proto.RegisterFile("userinfo.proto", fileDescriptor_785a78c34699a93d) }

var fileDescriptor_785a78c34699a93d = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2d, 0x4e, 0x2d,
	0xca, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xda,
	0xc1, 0xc8, 0xc5, 0x11, 0x5a, 0x9c, 0x5a, 0xe4, 0x99, 0x97, 0x96, 0x2f, 0x24, 0xc6, 0xc5, 0x06,
	0x66, 0xbb, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x42, 0x52, 0x5c, 0x1c, 0x7e,
	0x99, 0xc9, 0xd9, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x60, 0x19, 0x38, 0x5f, 0x48, 0x86, 0x8b,
	0x33, 0x38, 0x33, 0x3d, 0x2f, 0xb1, 0xa4, 0xb4, 0x28, 0x55, 0x82, 0x19, 0x2c, 0x89, 0x10, 0x10,
	0x52, 0xe2, 0xe2, 0xf1, 0x48, 0x4d, 0x4c, 0xf1, 0xcc, 0x4d, 0x4c, 0x4f, 0x0d, 0x2d, 0xca, 0x91,
	0x60, 0x01, 0x2b, 0x40, 0x11, 0x13, 0x92, 0xe3, 0xe2, 0x72, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x0d,
	0xc9, 0xcc, 0x4d, 0x95, 0x60, 0x55, 0x60, 0xd4, 0x60, 0x0e, 0x42, 0x12, 0x01, 0xb9, 0x2a, 0xb8,
	0x24, 0xb1, 0xa4, 0xb4, 0x58, 0x82, 0x4d, 0x81, 0x51, 0x83, 0x35, 0x08, 0xca, 0x73, 0x62, 0x5d,
	0xc5, 0xc4, 0x54, 0x90, 0x94, 0xc4, 0x06, 0xf6, 0x8c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x6a,
	0x66, 0xd9, 0x99, 0xde, 0x00, 0x00, 0x00,
}
