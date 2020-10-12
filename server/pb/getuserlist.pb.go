// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getuserlist.proto

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

type GetUserListRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserListRequest) Reset()         { *m = GetUserListRequest{} }
func (m *GetUserListRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserListRequest) ProtoMessage()    {}
func (*GetUserListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f06eb88822d1f2a, []int{0}
}

func (m *GetUserListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserListRequest.Unmarshal(m, b)
}
func (m *GetUserListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserListRequest.Marshal(b, m, deterministic)
}
func (m *GetUserListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserListRequest.Merge(m, src)
}
func (m *GetUserListRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserListRequest.Size(m)
}
func (m *GetUserListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserListRequest proto.InternalMessageInfo

func (m *GetUserListRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type GetUserListResponse struct {
	Result               *Result     `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	UserList             []*UserInfo `protobuf:"bytes,2,rep,name=UserList,proto3" json:"UserList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetUserListResponse) Reset()         { *m = GetUserListResponse{} }
func (m *GetUserListResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserListResponse) ProtoMessage()    {}
func (*GetUserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f06eb88822d1f2a, []int{1}
}

func (m *GetUserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserListResponse.Unmarshal(m, b)
}
func (m *GetUserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserListResponse.Marshal(b, m, deterministic)
}
func (m *GetUserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserListResponse.Merge(m, src)
}
func (m *GetUserListResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserListResponse.Size(m)
}
func (m *GetUserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserListResponse proto.InternalMessageInfo

func (m *GetUserListResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetUserListResponse) GetUserList() []*UserInfo {
	if m != nil {
		return m.UserList
	}
	return nil
}

func init() {
	proto.RegisterType((*GetUserListRequest)(nil), "pb.GetUserListRequest")
	proto.RegisterType((*GetUserListResponse)(nil), "pb.GetUserListResponse")
}

func init() { proto.RegisterFile("getuserlist.proto", fileDescriptor_8f06eb88822d1f2a) }

var fileDescriptor_8f06eb88822d1f2a = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4f, 0x2d, 0x29,
	0x2d, 0x4e, 0x2d, 0xca, 0xc9, 0x2c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a,
	0x48, 0x92, 0xe2, 0x29, 0x4a, 0x2d, 0x2e, 0xcd, 0x81, 0x8a, 0x48, 0xf1, 0x81, 0x54, 0x64, 0xe6,
	0xa5, 0xe5, 0x43, 0xf8, 0x4a, 0x3a, 0x5c, 0x42, 0xee, 0xa9, 0x25, 0xa1, 0xc5, 0xa9, 0x45, 0x3e,
	0x99, 0xc5, 0x25, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x62, 0x5c, 0x6c, 0x20, 0x21,
	0x4f, 0x17, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x28, 0x4f, 0x29, 0x99, 0x4b, 0x18, 0x45,
	0x75, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x12, 0x17, 0x5b, 0x10, 0xd8, 0x12, 0xb0, 0x72,
	0x6e, 0x23, 0x2e, 0xbd, 0x82, 0x24, 0x3d, 0x88, 0x48, 0x10, 0x54, 0x46, 0x48, 0x83, 0x8b, 0x03,
	0xa6, 0x4f, 0x82, 0x49, 0x81, 0x59, 0x83, 0xdb, 0x88, 0x07, 0xa4, 0x0a, 0x6c, 0x70, 0x5e, 0x5a,
	0x7e, 0x10, 0x5c, 0xd6, 0x89, 0x75, 0x15, 0x13, 0x53, 0x41, 0x52, 0x12, 0x1b, 0xd8, 0x81, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x06, 0x79, 0x5e, 0x31, 0xd7, 0x00, 0x00, 0x00,
}