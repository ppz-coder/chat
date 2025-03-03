// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request.proto

package protocol

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

type HeartBeatReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartBeatReq) Reset()         { *m = HeartBeatReq{} }
func (m *HeartBeatReq) String() string { return proto.CompactTextString(m) }
func (*HeartBeatReq) ProtoMessage()    {}
func (*HeartBeatReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{0}
}

func (m *HeartBeatReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeatReq.Unmarshal(m, b)
}
func (m *HeartBeatReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeatReq.Marshal(b, m, deterministic)
}
func (m *HeartBeatReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeatReq.Merge(m, src)
}
func (m *HeartBeatReq) XXX_Size() int {
	return xxx_messageInfo_HeartBeatReq.Size(m)
}
func (m *HeartBeatReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeatReq.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeatReq proto.InternalMessageInfo

//创建房间
type CreateRoomReq struct {
	RoomName             *string  `protobuf:"bytes,1,req,name=RoomName" json:"RoomName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoomReq) Reset()         { *m = CreateRoomReq{} }
func (m *CreateRoomReq) String() string { return proto.CompactTextString(m) }
func (*CreateRoomReq) ProtoMessage()    {}
func (*CreateRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{1}
}

func (m *CreateRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoomReq.Unmarshal(m, b)
}
func (m *CreateRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoomReq.Marshal(b, m, deterministic)
}
func (m *CreateRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomReq.Merge(m, src)
}
func (m *CreateRoomReq) XXX_Size() int {
	return xxx_messageInfo_CreateRoomReq.Size(m)
}
func (m *CreateRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomReq proto.InternalMessageInfo

func (m *CreateRoomReq) GetRoomName() string {
	if m != nil && m.RoomName != nil {
		return *m.RoomName
	}
	return ""
}

//进入房间
type EnterRoomReq struct {
	RoomName             *string  `protobuf:"bytes,1,req,name=RoomName" json:"RoomName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnterRoomReq) Reset()         { *m = EnterRoomReq{} }
func (m *EnterRoomReq) String() string { return proto.CompactTextString(m) }
func (*EnterRoomReq) ProtoMessage()    {}
func (*EnterRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{2}
}

func (m *EnterRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterRoomReq.Unmarshal(m, b)
}
func (m *EnterRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterRoomReq.Marshal(b, m, deterministic)
}
func (m *EnterRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterRoomReq.Merge(m, src)
}
func (m *EnterRoomReq) XXX_Size() int {
	return xxx_messageInfo_EnterRoomReq.Size(m)
}
func (m *EnterRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_EnterRoomReq proto.InternalMessageInfo

func (m *EnterRoomReq) GetRoomName() string {
	if m != nil && m.RoomName != nil {
		return *m.RoomName
	}
	return ""
}

//离开房间
type LeaveRoomReq struct {
	RoomName             *string  `protobuf:"bytes,1,req,name=RoomName" json:"RoomName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveRoomReq) Reset()         { *m = LeaveRoomReq{} }
func (m *LeaveRoomReq) String() string { return proto.CompactTextString(m) }
func (*LeaveRoomReq) ProtoMessage()    {}
func (*LeaveRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{3}
}

func (m *LeaveRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRoomReq.Unmarshal(m, b)
}
func (m *LeaveRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRoomReq.Marshal(b, m, deterministic)
}
func (m *LeaveRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRoomReq.Merge(m, src)
}
func (m *LeaveRoomReq) XXX_Size() int {
	return xxx_messageInfo_LeaveRoomReq.Size(m)
}
func (m *LeaveRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRoomReq proto.InternalMessageInfo

func (m *LeaveRoomReq) GetRoomName() string {
	if m != nil && m.RoomName != nil {
		return *m.RoomName
	}
	return ""
}

//聊天请求
type ChatReq struct {
	Data                 *string  `protobuf:"bytes,1,req,name=Data" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatReq) Reset()         { *m = ChatReq{} }
func (m *ChatReq) String() string { return proto.CompactTextString(m) }
func (*ChatReq) ProtoMessage()    {}
func (*ChatReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{4}
}

func (m *ChatReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatReq.Unmarshal(m, b)
}
func (m *ChatReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatReq.Marshal(b, m, deterministic)
}
func (m *ChatReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatReq.Merge(m, src)
}
func (m *ChatReq) XXX_Size() int {
	return xxx_messageInfo_ChatReq.Size(m)
}
func (m *ChatReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChatReq proto.InternalMessageInfo

func (m *ChatReq) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*HeartBeatReq)(nil), "protocol.HeartBeatReq")
	proto.RegisterType((*CreateRoomReq)(nil), "protocol.CreateRoomReq")
	proto.RegisterType((*EnterRoomReq)(nil), "protocol.EnterRoomReq")
	proto.RegisterType((*LeaveRoomReq)(nil), "protocol.LeaveRoomReq")
	proto.RegisterType((*ChatReq)(nil), "protocol.ChatReq")
}

func init() { proto.RegisterFile("request.proto", fileDescriptor_7f73548e33e655fe) }

var fileDescriptor_7f73548e33e655fe = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39,
	0x4a, 0x7c, 0x5c, 0x3c, 0x1e, 0xa9, 0x89, 0x45, 0x25, 0x4e, 0xa9, 0x89, 0x25, 0x41, 0xa9, 0x85,
	0x4a, 0xda, 0x5c, 0xbc, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0x41, 0xf9, 0xf9, 0xb9, 0x41, 0xa9,
	0x85, 0x42, 0x52, 0x5c, 0x1c, 0x20, 0xa6, 0x5f, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0x93, 0x06,
	0x67, 0x10, 0x9c, 0xaf, 0xa4, 0xc5, 0xc5, 0xe3, 0x9a, 0x57, 0x92, 0x5a, 0x44, 0xa4, 0x5a, 0x9f,
	0xd4, 0xc4, 0x32, 0xa2, 0xcc, 0x95, 0xe5, 0x62, 0x77, 0xce, 0x00, 0xbb, 0x47, 0x48, 0x88, 0x8b,
	0xc5, 0x25, 0xb1, 0x24, 0x11, 0xaa, 0x04, 0xcc, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xee,
	0xf7, 0xc4, 0xcd, 0x00, 0x00, 0x00,
}
