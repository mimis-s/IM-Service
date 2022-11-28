// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: home_relay.proto

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

// 转发给对方的聊天消息
type ChatSingleToReceiver struct {
	Data *ChatMessage `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *ChatSingleToReceiver) Reset()         { *m = ChatSingleToReceiver{} }
func (m *ChatSingleToReceiver) String() string { return proto.CompactTextString(m) }
func (*ChatSingleToReceiver) ProtoMessage()    {}
func (*ChatSingleToReceiver) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfa52dd254ee19fc, []int{0}
}
func (m *ChatSingleToReceiver) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChatSingleToReceiver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChatSingleToReceiver.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChatSingleToReceiver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatSingleToReceiver.Merge(m, src)
}
func (m *ChatSingleToReceiver) XXX_Size() int {
	return m.Size()
}
func (m *ChatSingleToReceiver) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatSingleToReceiver.DiscardUnknown(m)
}

var xxx_messageInfo_ChatSingleToReceiver proto.InternalMessageInfo

func (m *ChatSingleToReceiver) GetData() *ChatMessage {
	if m != nil {
		return m.Data
	}
	return nil
}

// 好友申请转发
type ApplyFriendsToReceiver struct {
	SenderID    int64     `protobuf:"varint,1,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	ReceiverID  int64     `protobuf:"varint,2,opt,name=ReceiverID,proto3" json:"ReceiverID,omitempty"`
	ApplyerInfo *UserInfo `protobuf:"bytes,3,opt,name=ApplyerInfo,proto3" json:"ApplyerInfo,omitempty"`
}

func (m *ApplyFriendsToReceiver) Reset()         { *m = ApplyFriendsToReceiver{} }
func (m *ApplyFriendsToReceiver) String() string { return proto.CompactTextString(m) }
func (*ApplyFriendsToReceiver) ProtoMessage()    {}
func (*ApplyFriendsToReceiver) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfa52dd254ee19fc, []int{1}
}
func (m *ApplyFriendsToReceiver) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApplyFriendsToReceiver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApplyFriendsToReceiver.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApplyFriendsToReceiver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyFriendsToReceiver.Merge(m, src)
}
func (m *ApplyFriendsToReceiver) XXX_Size() int {
	return m.Size()
}
func (m *ApplyFriendsToReceiver) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyFriendsToReceiver.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyFriendsToReceiver proto.InternalMessageInfo

func (m *ApplyFriendsToReceiver) GetSenderID() int64 {
	if m != nil {
		return m.SenderID
	}
	return 0
}

func (m *ApplyFriendsToReceiver) GetReceiverID() int64 {
	if m != nil {
		return m.ReceiverID
	}
	return 0
}

func (m *ApplyFriendsToReceiver) GetApplyerInfo() *UserInfo {
	if m != nil {
		return m.ApplyerInfo
	}
	return nil
}

// 同意好友申请转发
type AgreeApplyFriendsToReceiver struct {
	SenderID   int64 `protobuf:"varint,1,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	ReceiverID int64 `protobuf:"varint,2,opt,name=ReceiverID,proto3" json:"ReceiverID,omitempty"`
}

func (m *AgreeApplyFriendsToReceiver) Reset()         { *m = AgreeApplyFriendsToReceiver{} }
func (m *AgreeApplyFriendsToReceiver) String() string { return proto.CompactTextString(m) }
func (*AgreeApplyFriendsToReceiver) ProtoMessage()    {}
func (*AgreeApplyFriendsToReceiver) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfa52dd254ee19fc, []int{2}
}
func (m *AgreeApplyFriendsToReceiver) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AgreeApplyFriendsToReceiver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AgreeApplyFriendsToReceiver.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AgreeApplyFriendsToReceiver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgreeApplyFriendsToReceiver.Merge(m, src)
}
func (m *AgreeApplyFriendsToReceiver) XXX_Size() int {
	return m.Size()
}
func (m *AgreeApplyFriendsToReceiver) XXX_DiscardUnknown() {
	xxx_messageInfo_AgreeApplyFriendsToReceiver.DiscardUnknown(m)
}

var xxx_messageInfo_AgreeApplyFriendsToReceiver proto.InternalMessageInfo

func (m *AgreeApplyFriendsToReceiver) GetSenderID() int64 {
	if m != nil {
		return m.SenderID
	}
	return 0
}

func (m *AgreeApplyFriendsToReceiver) GetReceiverID() int64 {
	if m != nil {
		return m.ReceiverID
	}
	return 0
}

func init() {
	proto.RegisterType((*ChatSingleToReceiver)(nil), "im_home_proto.ChatSingleToReceiver")
	proto.RegisterType((*ApplyFriendsToReceiver)(nil), "im_home_proto.ApplyFriendsToReceiver")
	proto.RegisterType((*AgreeApplyFriendsToReceiver)(nil), "im_home_proto.AgreeApplyFriendsToReceiver")
}

func init() { proto.RegisterFile("home_relay.proto", fileDescriptor_bfa52dd254ee19fc) }

var fileDescriptor_bfa52dd254ee19fc = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x8f, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x1b, 0x27, 0x22, 0x19, 0xa2, 0x14, 0xd1, 0x52, 0x21, 0x48, 0x4f, 0x5e, 0xd6, 0x82,
	0x9e, 0x3c, 0x6e, 0x96, 0x41, 0x0f, 0xbb, 0xb4, 0x7a, 0x50, 0x84, 0x91, 0x65, 0xcf, 0x36, 0xd0,
	0x24, 0x25, 0xe9, 0x06, 0xfb, 0x2b, 0xf4, 0xcf, 0xf2, 0xb8, 0xa3, 0x47, 0x69, 0xff, 0x11, 0x31,
	0x9b, 0xd2, 0x79, 0xf6, 0xf4, 0xf2, 0xbe, 0xef, 0xcb, 0x8f, 0xef, 0xe1, 0x93, 0x42, 0x09, 0x98,
	0x6a, 0x28, 0xe9, 0x2a, 0xac, 0xb4, 0xaa, 0x95, 0x7b, 0xc4, 0xc5, 0xd4, 0x8a, 0x76, 0xf5, 0x8f,
	0xed, 0x9b, 0x15, 0xb4, 0xde, 0xf8, 0xbe, 0x6b, 0x05, 0xca, 0x98, 0x5a, 0xc8, 0xad, 0x16, 0x8c,
	0xf1, 0xe9, 0x5d, 0x41, 0xeb, 0x8c, 0xcb, 0xbc, 0x84, 0x7b, 0x95, 0x02, 0x03, 0xbe, 0x04, 0xed,
	0x86, 0x78, 0x3f, 0xa6, 0x35, 0xf5, 0xd0, 0x25, 0xba, 0xea, 0x5f, 0xfb, 0xe1, 0x0e, 0x3a, 0xfc,
	0xfe, 0x32, 0x01, 0x63, 0x68, 0x0e, 0xa9, 0xcd, 0x05, 0xaf, 0x08, 0x9f, 0x0d, 0xab, 0xaa, 0x5c,
	0x8d, 0x35, 0x07, 0x39, 0x37, 0x1d, 0x94, 0x8f, 0x0f, 0x33, 0x90, 0x73, 0xd0, 0x49, 0x6c, 0x71,
	0xbd, 0xf4, 0x77, 0x77, 0x09, 0xc6, 0x3f, 0xb9, 0x24, 0xf6, 0xf6, 0xac, 0xdb, 0x51, 0xdc, 0x5b,
	0xdc, 0xb7, 0x54, 0xd0, 0x89, 0x7c, 0x51, 0x5e, 0xcf, 0xb6, 0x39, 0xff, 0xd3, 0xe6, 0xc1, 0x6c,
	0xec, 0xb4, 0x9b, 0x0d, 0x1e, 0xf1, 0xc5, 0x30, 0xd7, 0x00, 0xff, 0xdf, 0x6a, 0xf4, 0xfc, 0xde,
	0x10, 0xb4, 0x6e, 0x08, 0xfa, 0x6c, 0x08, 0x7a, 0x6b, 0x89, 0xb3, 0x6e, 0x89, 0xf3, 0xd1, 0x12,
	0xe7, 0x69, 0x94, 0xf3, 0xba, 0x58, 0xcc, 0x42, 0xa6, 0x44, 0x24, 0xb8, 0xe0, 0x66, 0x60, 0xa2,
	0x64, 0x32, 0xc8, 0x40, 0x2f, 0x39, 0x83, 0xc8, 0x68, 0x16, 0x31, 0x25, 0x84, 0x92, 0xdb, 0x61,
	0x0f, 0x88, 0x76, 0xce, 0x99, 0x1d, 0xd8, 0x71, 0xf3, 0x15, 0x00, 0x00, 0xff, 0xff, 0x51, 0x1a,
	0x05, 0x1c, 0xe1, 0x01, 0x00, 0x00,
}

func (m *ChatSingleToReceiver) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChatSingleToReceiver) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChatSingleToReceiver) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintHomeRelay(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ApplyFriendsToReceiver) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApplyFriendsToReceiver) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApplyFriendsToReceiver) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ApplyerInfo != nil {
		{
			size, err := m.ApplyerInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHomeRelay(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ReceiverID != 0 {
		i = encodeVarintHomeRelay(dAtA, i, uint64(m.ReceiverID))
		i--
		dAtA[i] = 0x10
	}
	if m.SenderID != 0 {
		i = encodeVarintHomeRelay(dAtA, i, uint64(m.SenderID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AgreeApplyFriendsToReceiver) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AgreeApplyFriendsToReceiver) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AgreeApplyFriendsToReceiver) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReceiverID != 0 {
		i = encodeVarintHomeRelay(dAtA, i, uint64(m.ReceiverID))
		i--
		dAtA[i] = 0x10
	}
	if m.SenderID != 0 {
		i = encodeVarintHomeRelay(dAtA, i, uint64(m.SenderID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintHomeRelay(dAtA []byte, offset int, v uint64) int {
	offset -= sovHomeRelay(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChatSingleToReceiver) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovHomeRelay(uint64(l))
	}
	return n
}

func (m *ApplyFriendsToReceiver) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SenderID != 0 {
		n += 1 + sovHomeRelay(uint64(m.SenderID))
	}
	if m.ReceiverID != 0 {
		n += 1 + sovHomeRelay(uint64(m.ReceiverID))
	}
	if m.ApplyerInfo != nil {
		l = m.ApplyerInfo.Size()
		n += 1 + l + sovHomeRelay(uint64(l))
	}
	return n
}

func (m *AgreeApplyFriendsToReceiver) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SenderID != 0 {
		n += 1 + sovHomeRelay(uint64(m.SenderID))
	}
	if m.ReceiverID != 0 {
		n += 1 + sovHomeRelay(uint64(m.ReceiverID))
	}
	return n
}

func sovHomeRelay(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHomeRelay(x uint64) (n int) {
	return sovHomeRelay(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChatSingleToReceiver) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHomeRelay
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
			return fmt.Errorf("proto: ChatSingleToReceiver: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChatSingleToReceiver: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeRelay
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
				return ErrInvalidLengthHomeRelay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHomeRelay
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
			skippy, err := skipHomeRelay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHomeRelay
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
func (m *ApplyFriendsToReceiver) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHomeRelay
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
			return fmt.Errorf("proto: ApplyFriendsToReceiver: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApplyFriendsToReceiver: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderID", wireType)
			}
			m.SenderID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeRelay
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
					return ErrIntOverflowHomeRelay
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplyerInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeRelay
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
				return ErrInvalidLengthHomeRelay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHomeRelay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ApplyerInfo == nil {
				m.ApplyerInfo = &UserInfo{}
			}
			if err := m.ApplyerInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHomeRelay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHomeRelay
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
func (m *AgreeApplyFriendsToReceiver) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHomeRelay
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
			return fmt.Errorf("proto: AgreeApplyFriendsToReceiver: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AgreeApplyFriendsToReceiver: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderID", wireType)
			}
			m.SenderID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeRelay
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
					return ErrIntOverflowHomeRelay
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
		default:
			iNdEx = preIndex
			skippy, err := skipHomeRelay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHomeRelay
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
func skipHomeRelay(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHomeRelay
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
					return 0, ErrIntOverflowHomeRelay
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
					return 0, ErrIntOverflowHomeRelay
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
				return 0, ErrInvalidLengthHomeRelay
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHomeRelay
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHomeRelay
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHomeRelay        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHomeRelay          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHomeRelay = fmt.Errorf("proto: unexpected end of group")
)
