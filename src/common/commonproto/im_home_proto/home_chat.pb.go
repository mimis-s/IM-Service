// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: home_chat.proto

package im_home_proto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// 消息类型
type MessageType_Enum int32

const (
	MessageType_Enum_MessageType_Enum_Default MessageType_Enum = 0
	MessageType_Enum_EnumTextType             MessageType_Enum = 1
	MessageType_Enum_EnumImgType              MessageType_Enum = 2
	MessageType_Enum_EnumFileType             MessageType_Enum = 3
)

var MessageType_Enum_name = map[int32]string{
	0: "MessageType_Enum_Default",
	1: "EnumTextType",
	2: "EnumImgType",
	3: "EnumFileType",
}

var MessageType_Enum_value = map[string]int32{
	"MessageType_Enum_Default": 0,
	"EnumTextType":             1,
	"EnumImgType":              2,
	"EnumFileType":             3,
}

func (x MessageType_Enum) String() string {
	return proto.EnumName(MessageType_Enum_name, int32(x))
}

func (MessageType_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_75856de7ef776b75, []int{0}
}

// 消息状态
type MessageStatus_Enum int32

const (
	MessageStatus_Enum_MessageStatus_Enum_Default MessageStatus_Enum = 0
	MessageStatus_Enum_EnumSend                   MessageStatus_Enum = 1
	MessageStatus_Enum_EnumArrive                 MessageStatus_Enum = 2
	MessageStatus_Enum_EnumRead                   MessageStatus_Enum = 3
)

var MessageStatus_Enum_name = map[int32]string{
	0: "MessageStatus_Enum_Default",
	1: "EnumSend",
	2: "EnumArrive",
	3: "EnumRead",
}

var MessageStatus_Enum_value = map[string]int32{
	"MessageStatus_Enum_Default": 0,
	"EnumSend":                   1,
	"EnumArrive":                 2,
	"EnumRead":                   3,
}

func (x MessageStatus_Enum) String() string {
	return proto.EnumName(MessageStatus_Enum_name, int32(x))
}

func (MessageStatus_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_75856de7ef776b75, []int{1}
}

// 消息概要(当消息为文件或者图片的时候才需要)
type MessageRecap struct {
	FileName      string `protobuf:"bytes,1,opt,name=FileName,proto3" json:"FileName,omitempty"`
	FileExtension string `protobuf:"bytes,2,opt,name=FileExtension,proto3" json:"FileExtension,omitempty"`
	FileSize      int64  `protobuf:"varint,3,opt,name=FileSize,proto3" json:"FileSize,omitempty"`
}

func (m *MessageRecap) Reset()         { *m = MessageRecap{} }
func (m *MessageRecap) String() string { return proto.CompactTextString(m) }
func (*MessageRecap) ProtoMessage()    {}
func (*MessageRecap) Descriptor() ([]byte, []int) {
	return fileDescriptor_75856de7ef776b75, []int{0}
}
func (m *MessageRecap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MessageRecap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MessageRecap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MessageRecap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRecap.Merge(m, src)
}
func (m *MessageRecap) XXX_Size() int {
	return m.Size()
}
func (m *MessageRecap) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRecap.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRecap proto.InternalMessageInfo

func (m *MessageRecap) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *MessageRecap) GetFileExtension() string {
	if m != nil {
		return m.FileExtension
	}
	return ""
}

func (m *MessageRecap) GetFileSize() int64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

// 聊天消息
type ChatMessage struct {
	SenderID         int64              `protobuf:"varint,1,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	ReceiverID       int64              `protobuf:"varint,2,opt,name=ReceiverID,proto3" json:"ReceiverID,omitempty"`
	MessageID        int64              `protobuf:"varint,3,opt,name=MessageID,proto3" json:"MessageID,omitempty"`
	MessageType      MessageType_Enum   `protobuf:"varint,4,opt,name=MessageType,proto3,enum=im_home_proto.MessageType_Enum" json:"MessageType,omitempty"`
	SendTimeStamp    int64              `protobuf:"varint,5,opt,name=SendTimeStamp,proto3" json:"SendTimeStamp,omitempty"`
	MessageStatus    MessageStatus_Enum `protobuf:"varint,6,opt,name=MessageStatus,proto3,enum=im_home_proto.MessageStatus_Enum" json:"MessageStatus,omitempty"`
	MessageRecapInfo *MessageRecap      `protobuf:"bytes,7,opt,name=MessageRecapInfo,proto3" json:"MessageRecapInfo,omitempty"`
	Data             string             `protobuf:"bytes,8,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *ChatMessage) Reset()         { *m = ChatMessage{} }
func (m *ChatMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()    {}
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_75856de7ef776b75, []int{1}
}
func (m *ChatMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChatMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessage.Merge(m, src)
}
func (m *ChatMessage) XXX_Size() int {
	return m.Size()
}
func (m *ChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessage proto.InternalMessageInfo

func (m *ChatMessage) GetSenderID() int64 {
	if m != nil {
		return m.SenderID
	}
	return 0
}

func (m *ChatMessage) GetReceiverID() int64 {
	if m != nil {
		return m.ReceiverID
	}
	return 0
}

func (m *ChatMessage) GetMessageID() int64 {
	if m != nil {
		return m.MessageID
	}
	return 0
}

func (m *ChatMessage) GetMessageType() MessageType_Enum {
	if m != nil {
		return m.MessageType
	}
	return MessageType_Enum_MessageType_Enum_Default
}

func (m *ChatMessage) GetSendTimeStamp() int64 {
	if m != nil {
		return m.SendTimeStamp
	}
	return 0
}

func (m *ChatMessage) GetMessageStatus() MessageStatus_Enum {
	if m != nil {
		return m.MessageStatus
	}
	return MessageStatus_Enum_MessageStatus_Enum_Default
}

func (m *ChatMessage) GetMessageRecapInfo() *MessageRecap {
	if m != nil {
		return m.MessageRecapInfo
	}
	return nil
}

func (m *ChatMessage) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ChatSingleReq struct {
	Data *ChatMessage `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *ChatSingleReq) Reset()         { *m = ChatSingleReq{} }
func (m *ChatSingleReq) String() string { return proto.CompactTextString(m) }
func (*ChatSingleReq) ProtoMessage()    {}
func (*ChatSingleReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_75856de7ef776b75, []int{2}
}
func (m *ChatSingleReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChatSingleReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChatSingleReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChatSingleReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatSingleReq.Merge(m, src)
}
func (m *ChatSingleReq) XXX_Size() int {
	return m.Size()
}
func (m *ChatSingleReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatSingleReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChatSingleReq proto.InternalMessageInfo

func (m *ChatSingleReq) GetData() *ChatMessage {
	if m != nil {
		return m.Data
	}
	return nil
}

type ChatSingleRes struct {
	Data *ChatMessage `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *ChatSingleRes) Reset()         { *m = ChatSingleRes{} }
func (m *ChatSingleRes) String() string { return proto.CompactTextString(m) }
func (*ChatSingleRes) ProtoMessage()    {}
func (*ChatSingleRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_75856de7ef776b75, []int{3}
}
func (m *ChatSingleRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChatSingleRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChatSingleRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChatSingleRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatSingleRes.Merge(m, src)
}
func (m *ChatSingleRes) XXX_Size() int {
	return m.Size()
}
func (m *ChatSingleRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatSingleRes.DiscardUnknown(m)
}

var xxx_messageInfo_ChatSingleRes proto.InternalMessageInfo

func (m *ChatSingleRes) GetData() *ChatMessage {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("im_home_proto.MessageType_Enum", MessageType_Enum_name, MessageType_Enum_value)
	proto.RegisterEnum("im_home_proto.MessageStatus_Enum", MessageStatus_Enum_name, MessageStatus_Enum_value)
	proto.RegisterType((*MessageRecap)(nil), "im_home_proto.MessageRecap")
	proto.RegisterType((*ChatMessage)(nil), "im_home_proto.ChatMessage")
	proto.RegisterType((*ChatSingleReq)(nil), "im_home_proto.ChatSingleReq")
	proto.RegisterType((*ChatSingleRes)(nil), "im_home_proto.ChatSingleRes")
}

func init() { proto.RegisterFile("home_chat.proto", fileDescriptor_75856de7ef776b75) }

var fileDescriptor_75856de7ef776b75 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x6e, 0xdb, 0x3c,
	0x10, 0xc7, 0x4d, 0x2b, 0x5f, 0x3e, 0xe7, 0x6c, 0x27, 0x02, 0x27, 0xc2, 0x0d, 0x54, 0xd7, 0xe8,
	0x60, 0x04, 0x88, 0x04, 0xa4, 0x0f, 0x50, 0x24, 0x75, 0x1a, 0x68, 0x48, 0x07, 0xca, 0x53, 0x51,
	0xc0, 0x65, 0x94, 0x8b, 0x4d, 0xc0, 0x94, 0x5c, 0x89, 0x36, 0xd2, 0x3e, 0x45, 0x1f, 0xab, 0x63,
	0xc6, 0x8e, 0x85, 0xfd, 0x0e, 0x9d, 0x0b, 0x32, 0x8a, 0x2a, 0xc5, 0x5e, 0x3a, 0x91, 0xf7, 0xbb,
	0xff, 0xdd, 0x5f, 0xbc, 0x13, 0x1c, 0xcd, 0x52, 0x85, 0x93, 0x78, 0x26, 0xb4, 0xbf, 0xc8, 0x52,
	0x9d, 0xd2, 0xae, 0x54, 0x13, 0xcb, 0x6c, 0x38, 0x98, 0x43, 0xe7, 0x1a, 0xf3, 0x5c, 0x4c, 0x91,
	0x63, 0x2c, 0x16, 0xb4, 0x07, 0xad, 0xf7, 0x72, 0x8e, 0x1f, 0x84, 0x42, 0x46, 0xfa, 0x64, 0x78,
	0xc0, 0xcb, 0x98, 0xbe, 0x86, 0xae, 0xb9, 0x5f, 0xde, 0x6b, 0x4c, 0x72, 0x99, 0x26, 0xac, 0x69,
	0x05, 0x75, 0xf8, 0xd4, 0x21, 0x92, 0xdf, 0x90, 0x39, 0x7d, 0x32, 0x74, 0x78, 0x19, 0x0f, 0x7e,
	0x37, 0xa1, 0xfd, 0x6e, 0x26, 0x74, 0x61, 0x69, 0xb4, 0x11, 0x26, 0xb7, 0x98, 0x85, 0x23, 0xeb,
	0xe6, 0xf0, 0x32, 0xa6, 0x1e, 0x00, 0xc7, 0x18, 0xe5, 0xca, 0x66, 0x9b, 0x36, 0x5b, 0x21, 0xf4,
	0x18, 0x0e, 0x8a, 0x36, 0xe1, 0xa8, 0x30, 0xfa, 0x0b, 0xe8, 0x39, 0xb4, 0x8b, 0x60, 0xfc, 0x75,
	0x81, 0x6c, 0xaf, 0x4f, 0x86, 0x87, 0x67, 0x2f, 0xfd, 0xda, 0xe3, 0xfd, 0x8a, 0x62, 0x72, 0x99,
	0x2c, 0x15, 0xaf, 0xd6, 0x98, 0xe7, 0x9a, 0x8f, 0x19, 0x4b, 0x85, 0x91, 0x16, 0x6a, 0xc1, 0xfe,
	0xb3, 0x26, 0x75, 0x48, 0xaf, 0xa0, 0x5b, 0x14, 0x45, 0x5a, 0xe8, 0x65, 0xce, 0xf6, 0xad, 0xd5,
	0xab, 0xdd, 0x56, 0x8f, 0x9a, 0x47, 0xb3, 0x7a, 0x1d, 0xbd, 0x02, 0xb7, 0xba, 0x89, 0x30, 0xb9,
	0x4b, 0xd9, 0xff, 0x7d, 0x32, 0x6c, 0x9f, 0xbd, 0xd8, 0xdd, 0xcb, 0xca, 0xf8, 0x56, 0x11, 0xa5,
	0xb0, 0x37, 0x12, 0x5a, 0xb0, 0x96, 0xdd, 0x8e, 0xbd, 0x0f, 0xde, 0x42, 0xd7, 0xcc, 0x3d, 0x92,
	0xc9, 0x74, 0x8e, 0x1c, 0xbf, 0x50, 0xbf, 0x10, 0x11, 0xeb, 0xd0, 0x7b, 0xe6, 0x50, 0xd9, 0xd1,
	0xee, 0x06, 0xf9, 0xbf, 0x36, 0x38, 0xc1, 0xf2, 0x79, 0xe5, 0xb8, 0xe9, 0x31, 0xb0, 0xe7, 0x6c,
	0x32, 0xc2, 0x3b, 0xb1, 0x9c, 0x6b, 0xb7, 0x41, 0x5d, 0xe8, 0x18, 0x32, 0xc6, 0x7b, 0x6d, 0xd2,
	0x2e, 0xa1, 0x47, 0xd0, 0x36, 0x24, 0x54, 0x53, 0x0b, 0x9a, 0x4f, 0x12, 0xf3, 0x7f, 0x59, 0xe2,
	0x9c, 0x7c, 0x06, 0xba, 0x3d, 0x6a, 0xea, 0x41, 0x6f, 0x9b, 0x56, 0xac, 0x3a, 0xd0, 0x32, 0xc4,
	0x6c, 0xd6, 0x25, 0xf4, 0x10, 0xc0, 0x44, 0xe7, 0x59, 0x26, 0x57, 0xc6, 0xa5, 0xc8, 0x72, 0x14,
	0xb7, 0xae, 0x73, 0xf1, 0xe9, 0xc7, 0xda, 0x23, 0x0f, 0x6b, 0x8f, 0xfc, 0x5a, 0x7b, 0xe4, 0xfb,
	0xc6, 0x6b, 0x3c, 0x6c, 0xbc, 0xc6, 0xcf, 0x8d, 0xd7, 0xf8, 0x78, 0x31, 0x95, 0x7a, 0xb6, 0xbc,
	0xf1, 0xe3, 0x54, 0x05, 0x4a, 0x2a, 0x99, 0x9f, 0xe6, 0x41, 0x78, 0x7d, 0x1a, 0x61, 0xb6, 0x92,
	0x31, 0x06, 0x79, 0x16, 0x07, 0x71, 0xaa, 0x54, 0x9a, 0x14, 0x87, 0x1d, 0x55, 0x50, 0x1b, 0xdc,
	0xcd, 0xbe, 0x3d, 0xde, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x46, 0x71, 0xac, 0x71, 0xb8, 0x03,
	0x00, 0x00,
}

func (m *MessageRecap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MessageRecap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MessageRecap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FileSize != 0 {
		i = encodeVarintHomeChat(dAtA, i, uint64(m.FileSize))
		i--
		dAtA[i] = 0x18
	}
	if len(m.FileExtension) > 0 {
		i -= len(m.FileExtension)
		copy(dAtA[i:], m.FileExtension)
		i = encodeVarintHomeChat(dAtA, i, uint64(len(m.FileExtension)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FileName) > 0 {
		i -= len(m.FileName)
		copy(dAtA[i:], m.FileName)
		i = encodeVarintHomeChat(dAtA, i, uint64(len(m.FileName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChatMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChatMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChatMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintHomeChat(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x42
	}
	if m.MessageRecapInfo != nil {
		{
			size, err := m.MessageRecapInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHomeChat(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.MessageStatus != 0 {
		i = encodeVarintHomeChat(dAtA, i, uint64(m.MessageStatus))
		i--
		dAtA[i] = 0x30
	}
	if m.SendTimeStamp != 0 {
		i = encodeVarintHomeChat(dAtA, i, uint64(m.SendTimeStamp))
		i--
		dAtA[i] = 0x28
	}
	if m.MessageType != 0 {
		i = encodeVarintHomeChat(dAtA, i, uint64(m.MessageType))
		i--
		dAtA[i] = 0x20
	}
	if m.MessageID != 0 {
		i = encodeVarintHomeChat(dAtA, i, uint64(m.MessageID))
		i--
		dAtA[i] = 0x18
	}
	if m.ReceiverID != 0 {
		i = encodeVarintHomeChat(dAtA, i, uint64(m.ReceiverID))
		i--
		dAtA[i] = 0x10
	}
	if m.SenderID != 0 {
		i = encodeVarintHomeChat(dAtA, i, uint64(m.SenderID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ChatSingleReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChatSingleReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChatSingleReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHomeChat(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChatSingleRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChatSingleRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChatSingleRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHomeChat(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHomeChat(dAtA []byte, offset int, v uint64) int {
	offset -= sovHomeChat(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MessageRecap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FileName)
	if l > 0 {
		n += 1 + l + sovHomeChat(uint64(l))
	}
	l = len(m.FileExtension)
	if l > 0 {
		n += 1 + l + sovHomeChat(uint64(l))
	}
	if m.FileSize != 0 {
		n += 1 + sovHomeChat(uint64(m.FileSize))
	}
	return n
}

func (m *ChatMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SenderID != 0 {
		n += 1 + sovHomeChat(uint64(m.SenderID))
	}
	if m.ReceiverID != 0 {
		n += 1 + sovHomeChat(uint64(m.ReceiverID))
	}
	if m.MessageID != 0 {
		n += 1 + sovHomeChat(uint64(m.MessageID))
	}
	if m.MessageType != 0 {
		n += 1 + sovHomeChat(uint64(m.MessageType))
	}
	if m.SendTimeStamp != 0 {
		n += 1 + sovHomeChat(uint64(m.SendTimeStamp))
	}
	if m.MessageStatus != 0 {
		n += 1 + sovHomeChat(uint64(m.MessageStatus))
	}
	if m.MessageRecapInfo != nil {
		l = m.MessageRecapInfo.Size()
		n += 1 + l + sovHomeChat(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovHomeChat(uint64(l))
	}
	return n
}

func (m *ChatSingleReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovHomeChat(uint64(l))
	}
	return n
}

func (m *ChatSingleRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovHomeChat(uint64(l))
	}
	return n
}

func sovHomeChat(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHomeChat(x uint64) (n int) {
	return sovHomeChat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MessageRecap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHomeChat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MessageRecap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessageRecap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHomeChat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHomeChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileExtension", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHomeChat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHomeChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileExtension = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileSize", wireType)
			}
			m.FileSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FileSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHomeChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHomeChat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ChatMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHomeChat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChatMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChatMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderID", wireType)
			}
			m.SenderID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SenderID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiverID", wireType)
			}
			m.ReceiverID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReceiverID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageID", wireType)
			}
			m.MessageID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageType", wireType)
			}
			m.MessageType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageType |= MessageType_Enum(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendTimeStamp", wireType)
			}
			m.SendTimeStamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SendTimeStamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageStatus", wireType)
			}
			m.MessageStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageStatus |= MessageStatus_Enum(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageRecapInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHomeChat
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHomeChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MessageRecapInfo == nil {
				m.MessageRecapInfo = &MessageRecap{}
			}
			if err := m.MessageRecapInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHomeChat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHomeChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHomeChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHomeChat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ChatSingleReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHomeChat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChatSingleReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChatSingleReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHomeChat
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHomeChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &ChatMessage{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHomeChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHomeChat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ChatSingleRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHomeChat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChatSingleRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChatSingleRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHomeChat
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHomeChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &ChatMessage{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHomeChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHomeChat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHomeChat(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHomeChat
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHomeChat
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHomeChat
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthHomeChat
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHomeChat
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHomeChat
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHomeChat        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHomeChat          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHomeChat = fmt.Errorf("proto: unexpected end of group")
)
