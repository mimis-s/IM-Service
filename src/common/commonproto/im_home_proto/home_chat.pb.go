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

// 聊天消息
type ChatMessage struct {
	SenderID      int64              `protobuf:"varint,1,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	ReceiverID    int64              `protobuf:"varint,2,opt,name=ReceiverID,proto3" json:"ReceiverID,omitempty"`
	MessageID     int32              `protobuf:"varint,3,opt,name=MessageID,proto3" json:"MessageID,omitempty"`
	MessageType   MessageType_Enum   `protobuf:"varint,4,opt,name=MessageType,proto3,enum=im_home_proto.MessageType_Enum" json:"MessageType,omitempty"`
	SendTimeStamp int64              `protobuf:"varint,5,opt,name=SendTimeStamp,proto3" json:"SendTimeStamp,omitempty"`
	MessageStatus MessageStatus_Enum `protobuf:"varint,6,opt,name=MessageStatus,proto3,enum=im_home_proto.MessageStatus_Enum" json:"MessageStatus,omitempty"`
	Data          string             `protobuf:"bytes,7,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *ChatMessage) Reset()         { *m = ChatMessage{} }
func (m *ChatMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()    {}
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_75856de7ef776b75, []int{0}
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

func (m *ChatMessage) GetMessageID() int32 {
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
	return fileDescriptor_75856de7ef776b75, []int{1}
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
	return fileDescriptor_75856de7ef776b75, []int{2}
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
	proto.RegisterType((*ChatMessage)(nil), "im_home_proto.ChatMessage")
	proto.RegisterType((*ChatSingleReq)(nil), "im_home_proto.ChatSingleReq")
	proto.RegisterType((*ChatSingleRes)(nil), "im_home_proto.ChatSingleRes")
}

func init() { proto.RegisterFile("home_chat.proto", fileDescriptor_75856de7ef776b75) }

var fileDescriptor_75856de7ef776b75 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb3, 0x49, 0x5b, 0xda, 0x49, 0xd3, 0x5a, 0x73, 0xb2, 0xa2, 0x6a, 0x09, 0x15, 0x87,
	0xa8, 0x52, 0x6d, 0xa9, 0x3c, 0x00, 0x6a, 0x09, 0x20, 0x1f, 0x7a, 0xb1, 0x73, 0x42, 0x48, 0x66,
	0xeb, 0x0e, 0xf6, 0x4a, 0x59, 0x3b, 0xd8, 0xeb, 0x08, 0xde, 0x82, 0x17, 0xe1, 0x3d, 0x38, 0xf6,
	0xc8, 0x11, 0x25, 0x2f, 0x82, 0x76, 0x6b, 0xc0, 0x4e, 0x7b, 0xe9, 0x69, 0x3d, 0xdf, 0xcc, 0xfc,
	0xff, 0xe8, 0x37, 0x1c, 0x67, 0x85, 0xa2, 0x38, 0xc9, 0x84, 0xf6, 0x96, 0x65, 0xa1, 0x0b, 0x1c,
	0x49, 0x15, 0x5b, 0x66, 0xcb, 0xd3, 0x1f, 0x7d, 0x18, 0xbe, 0xc9, 0x84, 0xbe, 0xa6, 0xaa, 0x12,
	0x29, 0xe1, 0x18, 0xf6, 0x23, 0xca, 0x6f, 0xa9, 0x0c, 0x66, 0x2e, 0x9b, 0xb0, 0xe9, 0x20, 0xfc,
	0x57, 0x23, 0x07, 0x08, 0x29, 0x21, 0xb9, 0xb2, 0xdd, 0xbe, 0xed, 0xb6, 0x08, 0x9e, 0xc0, 0x41,
	0x23, 0x13, 0xcc, 0xdc, 0xc1, 0x84, 0x4d, 0x77, 0xc3, 0xff, 0x00, 0x2f, 0x61, 0xd8, 0x14, 0xf3,
	0x6f, 0x4b, 0x72, 0x77, 0x26, 0x6c, 0x7a, 0x74, 0xf1, 0xdc, 0xeb, 0x9c, 0xe3, 0xb5, 0x26, 0xe2,
	0xb7, 0x79, 0xad, 0xc2, 0xf6, 0x0e, 0xbe, 0x84, 0x91, 0x39, 0x66, 0x2e, 0x15, 0x45, 0x5a, 0xa8,
	0xa5, 0xbb, 0x6b, 0x6f, 0xe8, 0x42, 0x7c, 0x0f, 0xa3, 0x66, 0x29, 0xd2, 0x42, 0xd7, 0x95, 0xbb,
	0x67, 0xad, 0x5e, 0x3c, 0x6e, 0x75, 0x3f, 0x73, 0x6f, 0xd6, 0xdd, 0x43, 0x84, 0x9d, 0x99, 0xd0,
	0xc2, 0x7d, 0x36, 0x61, 0xd3, 0x83, 0xd0, 0x7e, 0x9f, 0xbe, 0x86, 0x91, 0x89, 0x2b, 0x92, 0x79,
	0xba, 0xa0, 0x90, 0xbe, 0xa0, 0xd7, 0x0c, 0x99, 0xb0, 0x86, 0x17, 0xe3, 0x2d, 0x93, 0x56, 0xb4,
	0x8f, 0x0b, 0x54, 0x4f, 0x15, 0x38, 0x23, 0x70, 0xb6, 0x53, 0xc2, 0x13, 0x70, 0xb7, 0x59, 0x3c,
	0xa3, 0xcf, 0xa2, 0x5e, 0x68, 0xa7, 0x87, 0x0e, 0x1c, 0x1a, 0x32, 0xa7, 0xaf, 0xda, 0xb4, 0x1d,
	0x86, 0xc7, 0x30, 0x34, 0x24, 0x50, 0xa9, 0x05, 0xfd, 0xbf, 0x23, 0xef, 0xe4, 0xc2, 0x2a, 0x38,
	0x83, 0xb3, 0x4f, 0x80, 0x0f, 0x13, 0x42, 0x0e, 0xe3, 0x87, 0xb4, 0x65, 0x75, 0x08, 0xfb, 0x86,
	0x98, 0x1f, 0xe2, 0x30, 0x3c, 0x02, 0x30, 0xd5, 0x65, 0x59, 0xca, 0x95, 0x71, 0x69, 0xba, 0x21,
	0x89, 0x5b, 0x67, 0x70, 0xf5, 0xf1, 0xe7, 0x9a, 0xb3, 0xbb, 0x35, 0x67, 0xbf, 0xd7, 0x9c, 0x7d,
	0xdf, 0xf0, 0xde, 0xdd, 0x86, 0xf7, 0x7e, 0x6d, 0x78, 0xef, 0xc3, 0x55, 0x2a, 0x75, 0x56, 0xdf,
	0x78, 0x49, 0xa1, 0x7c, 0x25, 0x95, 0xac, 0xce, 0x2b, 0x3f, 0xb8, 0x3e, 0x8f, 0xa8, 0x5c, 0xc9,
	0x84, 0xfc, 0xaa, 0x4c, 0xfc, 0xa4, 0x50, 0xaa, 0xc8, 0x9b, 0xc7, 0x46, 0xe5, 0x77, 0x82, 0xbb,
	0xd9, 0xb3, 0xcf, 0xab, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x59, 0x6a, 0x4a, 0x01, 0x03,
	0x00, 0x00,
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
				m.MessageID |= int32(b&0x7F) << shift
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
