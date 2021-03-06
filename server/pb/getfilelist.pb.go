// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getfilelist.proto

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

type DownFileInfo struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	FileSize             int32    `protobuf:"varint,2,opt,name=FileSize,proto3" json:"FileSize,omitempty"`
	FileName             string   `protobuf:"bytes,3,opt,name=FileName,proto3" json:"FileName,omitempty"`
	DownUrl              string   `protobuf:"bytes,4,opt,name=DownUrl,proto3" json:"DownUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownFileInfo) Reset()         { *m = DownFileInfo{} }
func (m *DownFileInfo) String() string { return proto.CompactTextString(m) }
func (*DownFileInfo) ProtoMessage()    {}
func (*DownFileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e4e237099156d3, []int{0}
}

func (m *DownFileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownFileInfo.Unmarshal(m, b)
}
func (m *DownFileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownFileInfo.Marshal(b, m, deterministic)
}
func (m *DownFileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownFileInfo.Merge(m, src)
}
func (m *DownFileInfo) XXX_Size() int {
	return xxx_messageInfo_DownFileInfo.Size(m)
}
func (m *DownFileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DownFileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DownFileInfo proto.InternalMessageInfo

func (m *DownFileInfo) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *DownFileInfo) GetFileSize() int32 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *DownFileInfo) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *DownFileInfo) GetDownUrl() string {
	if m != nil {
		return m.DownUrl
	}
	return ""
}

type GetFileListRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	UserType             int32    `protobuf:"varint,2,opt,name=UserType,proto3" json:"UserType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFileListRequest) Reset()         { *m = GetFileListRequest{} }
func (m *GetFileListRequest) String() string { return proto.CompactTextString(m) }
func (*GetFileListRequest) ProtoMessage()    {}
func (*GetFileListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e4e237099156d3, []int{1}
}

func (m *GetFileListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileListRequest.Unmarshal(m, b)
}
func (m *GetFileListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileListRequest.Marshal(b, m, deterministic)
}
func (m *GetFileListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileListRequest.Merge(m, src)
}
func (m *GetFileListRequest) XXX_Size() int {
	return xxx_messageInfo_GetFileListRequest.Size(m)
}
func (m *GetFileListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileListRequest proto.InternalMessageInfo

func (m *GetFileListRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GetFileListRequest) GetUserType() int32 {
	if m != nil {
		return m.UserType
	}
	return 0
}

type GetFileListResponse struct {
	Result               *Result         `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	FileList             []*DownFileInfo `protobuf:"bytes,2,rep,name=FileList,proto3" json:"FileList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetFileListResponse) Reset()         { *m = GetFileListResponse{} }
func (m *GetFileListResponse) String() string { return proto.CompactTextString(m) }
func (*GetFileListResponse) ProtoMessage()    {}
func (*GetFileListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e4e237099156d3, []int{2}
}

func (m *GetFileListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileListResponse.Unmarshal(m, b)
}
func (m *GetFileListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileListResponse.Marshal(b, m, deterministic)
}
func (m *GetFileListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileListResponse.Merge(m, src)
}
func (m *GetFileListResponse) XXX_Size() int {
	return xxx_messageInfo_GetFileListResponse.Size(m)
}
func (m *GetFileListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileListResponse proto.InternalMessageInfo

func (m *GetFileListResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetFileListResponse) GetFileList() []*DownFileInfo {
	if m != nil {
		return m.FileList
	}
	return nil
}

func init() {
	proto.RegisterType((*DownFileInfo)(nil), "pb.DownFileInfo")
	proto.RegisterType((*GetFileListRequest)(nil), "pb.GetFileListRequest")
	proto.RegisterType((*GetFileListResponse)(nil), "pb.GetFileListResponse")
}

func init() { proto.RegisterFile("getfilelist.proto", fileDescriptor_87e4e237099156d3) }

var fileDescriptor_87e4e237099156d3 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4f, 0x2d, 0x49,
	0xcb, 0xcc, 0x49, 0xcd, 0xc9, 0x2c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a,
	0x48, 0x92, 0xe2, 0x29, 0x4a, 0x2d, 0x2e, 0xcd, 0x81, 0x8a, 0x28, 0x55, 0x70, 0xf1, 0xb8, 0xe4,
	0x97, 0xe7, 0xb9, 0x65, 0xe6, 0xa4, 0x7a, 0xe6, 0xa5, 0xe5, 0x0b, 0x89, 0x71, 0xb1, 0x85, 0x16,
	0xa7, 0x16, 0x79, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x42, 0x52, 0x5c,
	0x1c, 0x20, 0x35, 0xc1, 0x99, 0x55, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x70, 0x3e,
	0x4c, 0xce, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x19, 0xac, 0x0b, 0xce, 0x17, 0x92, 0xe0, 0x62, 0x07,
	0x99, 0x1f, 0x5a, 0x94, 0x23, 0xc1, 0x02, 0x96, 0x82, 0x71, 0x95, 0x3c, 0xb8, 0x84, 0xdc, 0x53,
	0x4b, 0x40, 0x0a, 0x7d, 0x32, 0x8b, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xf0, 0xd9,
	0x0f, 0x62, 0x85, 0x54, 0x16, 0xc0, 0xed, 0x87, 0xf1, 0x95, 0xd2, 0xb9, 0x84, 0x51, 0x4c, 0x2a,
	0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x52, 0xe2, 0x62, 0x0b, 0x02, 0x7b, 0x15, 0x6c, 0x14, 0xb7,
	0x11, 0x97, 0x5e, 0x41, 0x92, 0x1e, 0x44, 0x24, 0x08, 0x2a, 0x23, 0xa4, 0x03, 0x71, 0x3a, 0x48,
	0x9f, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x00, 0x48, 0x15, 0x72, 0x90, 0x04, 0xc1, 0x55,
	0x38, 0xb1, 0xae, 0x62, 0x62, 0x2a, 0x48, 0x4a, 0x62, 0x03, 0x07, 0x9d, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x14, 0xb2, 0x44, 0x81, 0x61, 0x01, 0x00, 0x00,
}
