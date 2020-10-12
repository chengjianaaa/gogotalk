// Code generated by protoc-gen-go. DO NOT EDIT.
// source: uploadfile.proto

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

type UploadFileRequest struct {
	FromUserID           string   `protobuf:"bytes,1,opt,name=FromUserID,proto3" json:"FromUserID,omitempty"`
	ToUserID             string   `protobuf:"bytes,2,opt,name=ToUserID,proto3" json:"ToUserID,omitempty"`
	UserType             int32    `protobuf:"varint,3,opt,name=UserType,proto3" json:"UserType,omitempty"`
	FileName             string   `protobuf:"bytes,4,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Data                 []byte   `protobuf:"bytes,5,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadFileRequest) Reset()         { *m = UploadFileRequest{} }
func (m *UploadFileRequest) String() string { return proto.CompactTextString(m) }
func (*UploadFileRequest) ProtoMessage()    {}
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b4910f4aaf3992e, []int{0}
}

func (m *UploadFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileRequest.Unmarshal(m, b)
}
func (m *UploadFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileRequest.Marshal(b, m, deterministic)
}
func (m *UploadFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileRequest.Merge(m, src)
}
func (m *UploadFileRequest) XXX_Size() int {
	return xxx_messageInfo_UploadFileRequest.Size(m)
}
func (m *UploadFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileRequest proto.InternalMessageInfo

func (m *UploadFileRequest) GetFromUserID() string {
	if m != nil {
		return m.FromUserID
	}
	return ""
}

func (m *UploadFileRequest) GetToUserID() string {
	if m != nil {
		return m.ToUserID
	}
	return ""
}

func (m *UploadFileRequest) GetUserType() int32 {
	if m != nil {
		return m.UserType
	}
	return 0
}

func (m *UploadFileRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *UploadFileRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type UploadFileResponse struct {
	Result               *Result  `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	FileName             string   `protobuf:"bytes,2,opt,name=FileName,proto3" json:"FileName,omitempty"`
	DownUrl              string   `protobuf:"bytes,3,opt,name=DownUrl,proto3" json:"DownUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadFileResponse) Reset()         { *m = UploadFileResponse{} }
func (m *UploadFileResponse) String() string { return proto.CompactTextString(m) }
func (*UploadFileResponse) ProtoMessage()    {}
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b4910f4aaf3992e, []int{1}
}

func (m *UploadFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileResponse.Unmarshal(m, b)
}
func (m *UploadFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileResponse.Marshal(b, m, deterministic)
}
func (m *UploadFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileResponse.Merge(m, src)
}
func (m *UploadFileResponse) XXX_Size() int {
	return xxx_messageInfo_UploadFileResponse.Size(m)
}
func (m *UploadFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileResponse proto.InternalMessageInfo

func (m *UploadFileResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *UploadFileResponse) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *UploadFileResponse) GetDownUrl() string {
	if m != nil {
		return m.DownUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*UploadFileRequest)(nil), "pb.UploadFileRequest")
	proto.RegisterType((*UploadFileResponse)(nil), "pb.UploadFileResponse")
}

func init() { proto.RegisterFile("uploadfile.proto", fileDescriptor_9b4910f4aaf3992e) }

var fileDescriptor_9b4910f4aaf3992e = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4e, 0xc4, 0x20,
	0x14, 0x86, 0x03, 0x4e, 0x47, 0xe7, 0x39, 0x0b, 0x7d, 0x2b, 0x32, 0x0b, 0xd3, 0x74, 0xd5, 0x55,
	0x17, 0x7a, 0x03, 0xd3, 0x34, 0x71, 0xe3, 0x82, 0x4c, 0x0f, 0x00, 0xf1, 0x99, 0x4c, 0xc2, 0x14,
	0x04, 0x1a, 0xe3, 0x59, 0xbc, 0x81, 0xa7, 0x34, 0xd0, 0x32, 0x99, 0xdd, 0xfb, 0xff, 0x0f, 0xf2,
	0x3e, 0x80, 0x87, 0xd9, 0x19, 0xab, 0x3e, 0x3e, 0x4f, 0x86, 0x3a, 0xe7, 0x6d, 0xb4, 0xc8, 0x9d,
	0x3e, 0xec, 0x3d, 0x85, 0xd9, 0xc4, 0xa5, 0x69, 0x7e, 0x19, 0x3c, 0x8e, 0xf9, 0xd8, 0x70, 0x32,
	0x24, 0xe9, 0x6b, 0xa6, 0x10, 0xf1, 0x09, 0x60, 0xf0, 0xf6, 0x3c, 0x06, 0xf2, 0x6f, 0xbd, 0x60,
	0x35, 0x6b, 0x77, 0xf2, 0xaa, 0xc1, 0x03, 0xdc, 0x1d, 0xed, 0x4a, 0x79, 0xa6, 0x97, 0x9c, 0x58,
	0x9a, 0x8e, 0x3f, 0x8e, 0xc4, 0x4d, 0xcd, 0xda, 0x4a, 0x5e, 0x72, 0x62, 0xc9, 0xe6, 0x5d, 0x9d,
	0x49, 0x6c, 0x96, 0x7b, 0x25, 0x23, 0xc2, 0xa6, 0x57, 0x51, 0x89, 0xaa, 0x66, 0xed, 0x5e, 0xe6,
	0xb9, 0x99, 0x00, 0xaf, 0xe5, 0x82, 0xb3, 0x53, 0x20, 0x6c, 0x60, 0x2b, 0xf3, 0x1b, 0xb2, 0xd9,
	0xfd, 0x33, 0x74, 0x4e, 0x77, 0x4b, 0x23, 0x57, 0x92, 0x36, 0x0d, 0x65, 0xd3, 0x6a, 0x58, 0x32,
	0x0a, 0xb8, 0xed, 0xed, 0xf7, 0x34, 0x7a, 0x93, 0x05, 0x77, 0xb2, 0xc4, 0xd7, 0xea, 0x8f, 0x73,
	0xa7, 0xf5, 0x36, 0xff, 0xcd, 0xcb, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xe9, 0xa9, 0x92,
	0x41, 0x01, 0x00, 0x00,
}
