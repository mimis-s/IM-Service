// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: home_message.proto

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

type GetSingleChatHistoryReq struct {
	MaxNotGainMessageID int64 `protobuf:"varint,1,opt,name=MaxNotGainMessageID,proto3" json:"MaxNotGainMessageID,omitempty"`
	FriendID            int64 `protobuf:"varint,2,opt,name=FriendID,proto3" json:"FriendID,omitempty"`
}

func (m *GetSingleChatHistoryReq) Reset()         { *m = GetSingleChatHistoryReq{} }
func (m *GetSingleChatHistoryReq) String() string { return proto.CompactTextString(m) }
func (*GetSingleChatHistoryReq) ProtoMessage()    {}
func (*GetSingleChatHistoryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e858025c0727086, []int{0}
}
func (m *GetSingleChatHistoryReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSingleChatHistoryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetSingleChatHistoryReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetSingleChatHistoryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSingleChatHistoryReq.Merge(m, src)
}
func (m *GetSingleChatHistoryReq) XXX_Size() int {
	return m.Size()
}
func (m *GetSingleChatHistoryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSingleChatHistoryReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetSingleChatHistoryReq proto.InternalMessageInfo

func (m *GetSingleChatHistoryReq) GetMaxNotGainMessageID() int64 {
	if m != nil {
		return m.MaxNotGainMessageID
	}
	return 0
}

func (m *GetSingleChatHistoryReq) GetFriendID() int64 {
	if m != nil {
		return m.FriendID
	}
	return 0
}

type GetSingleChatHistoryRes struct {
	FriendID int64          `protobuf:"varint,1,opt,name=FriendID,proto3" json:"FriendID,omitempty"`
	Data     []*ChatMessage `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (m *GetSingleChatHistoryRes) Reset()         { *m = GetSingleChatHistoryRes{} }
func (m *GetSingleChatHistoryRes) String() string { return proto.CompactTextString(m) }
func (*GetSingleChatHistoryRes) ProtoMessage()    {}
func (*GetSingleChatHistoryRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e858025c0727086, []int{1}
}
func (m *GetSingleChatHistoryRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSingleChatHistoryRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetSingleChatHistoryRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetSingleChatHistoryRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSingleChatHistoryRes.Merge(m, src)
}
func (m *GetSingleChatHistoryRes) XXX_Size() int {
	return m.Size()
}
func (m *GetSingleChatHistoryRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSingleChatHistoryRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetSingleChatHistoryRes proto.InternalMessageInfo

func (m *GetSingleChatHistoryRes) GetFriendID() int64 {
	if m != nil {
		return m.FriendID
	}
	return 0
}

func (m *GetSingleChatHistoryRes) GetData() []*ChatMessage {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetSingleChatHistoryReq)(nil), "im_home_proto.GetSingleChatHistoryReq")
	proto.RegisterType((*GetSingleChatHistoryRes)(nil), "im_home_proto.GetSingleChatHistoryRes")
}

func init() { proto.RegisterFile("home_message.proto", fileDescriptor_6e858025c0727086) }

var fileDescriptor_6e858025c0727086 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xc8, 0xcf, 0x4d,
	0x8d, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0xcd, 0xcc, 0x8d, 0x07, 0x0b, 0x83, 0xb9, 0x52, 0xfc, 0x60, 0x76, 0x72, 0x46, 0x62, 0x09, 0x44,
	0x5e, 0x29, 0x9d, 0x4b, 0xdc, 0x3d, 0xb5, 0x24, 0x38, 0x33, 0x2f, 0x3d, 0x27, 0xd5, 0x39, 0x23,
	0xb1, 0xc4, 0x23, 0xb3, 0xb8, 0x24, 0xbf, 0xa8, 0x32, 0x28, 0xb5, 0x50, 0xc8, 0x80, 0x4b, 0xd8,
	0x37, 0xb1, 0xc2, 0x2f, 0xbf, 0xc4, 0x3d, 0x31, 0x33, 0xcf, 0x17, 0x62, 0xaa, 0xa7, 0x8b, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x36, 0x29, 0x21, 0x29, 0x2e, 0x0e, 0xb7, 0xa2, 0xcc, 0xd4,
	0xbc, 0x14, 0x4f, 0x17, 0x09, 0x26, 0xb0, 0x32, 0x38, 0x5f, 0x29, 0x15, 0x97, 0x45, 0xc5, 0x28,
	0xda, 0x18, 0x51, 0xb5, 0x09, 0xe9, 0x71, 0xb1, 0xb8, 0x24, 0x96, 0x24, 0x4a, 0x30, 0x29, 0x30,
	0x6b, 0x70, 0x1b, 0x49, 0xe9, 0xa1, 0x78, 0x47, 0x0f, 0x64, 0x10, 0xd4, 0xfa, 0x20, 0xb0, 0x3a,
	0xa7, 0x98, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x72, 0x4a, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0xcc, 0xcd, 0x2c, 0xd6, 0x2d, 0xd6,
	0xf7, 0xf4, 0xd5, 0x0d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2f, 0x2e, 0x4a, 0xd6, 0x4f,
	0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x52, 0x60, 0x1b, 0xf4, 0x51, 0xec, 0x4b, 0x62, 0x03, 0x53,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xe5, 0xee, 0x61, 0x6a, 0x01, 0x00, 0x00,
}

func (m *GetSingleChatHistoryReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSingleChatHistoryReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSingleChatHistoryReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FriendID != 0 {
		i = encodeVarintHomeMessage(dAtA, i, uint64(m.FriendID))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxNotGainMessageID != 0 {
		i = encodeVarintHomeMessage(dAtA, i, uint64(m.MaxNotGainMessageID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetSingleChatHistoryRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSingleChatHistoryRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSingleChatHistoryRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHomeMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.FriendID != 0 {
		i = encodeVarintHomeMessage(dAtA, i, uint64(m.FriendID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintHomeMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovHomeMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetSingleChatHistoryReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxNotGainMessageID != 0 {
		n += 1 + sovHomeMessage(uint64(m.MaxNotGainMessageID))
	}
	if m.FriendID != 0 {
		n += 1 + sovHomeMessage(uint64(m.FriendID))
	}
	return n
}

func (m *GetSingleChatHistoryRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FriendID != 0 {
		n += 1 + sovHomeMessage(uint64(m.FriendID))
	}
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovHomeMessage(uint64(l))
		}
	}
	return n
}

func sovHomeMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHomeMessage(x uint64) (n int) {
	return sovHomeMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetSingleChatHistoryReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHomeMessage
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
			return fmt.Errorf("proto: GetSingleChatHistoryReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSingleChatHistoryReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNotGainMessageID", wireType)
			}
			m.MaxNotGainMessageID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNotGainMessageID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FriendID", wireType)
			}
			m.FriendID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FriendID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHomeMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHomeMessage
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
func (m *GetSingleChatHistoryRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHomeMessage
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
			return fmt.Errorf("proto: GetSingleChatHistoryRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSingleChatHistoryRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FriendID", wireType)
			}
			m.FriendID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FriendID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHomeMessage
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
				return ErrInvalidLengthHomeMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHomeMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &ChatMessage{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHomeMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHomeMessage
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
func skipHomeMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHomeMessage
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
					return 0, ErrIntOverflowHomeMessage
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
					return 0, ErrIntOverflowHomeMessage
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
				return 0, ErrInvalidLengthHomeMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHomeMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHomeMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHomeMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHomeMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHomeMessage = fmt.Errorf("proto: unexpected end of group")
)