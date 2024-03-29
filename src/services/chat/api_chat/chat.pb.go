// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.proto

package api_chat

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	im_error_proto "github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	im_home_proto "github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
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

type ChatSingleReq struct {
	ClientInfo *im_home_proto.ClientOnlineInfo `protobuf:"bytes,1,opt,name=ClientInfo,proto3" json:"ClientInfo,omitempty"`
	Data       *im_home_proto.ChatMessage      `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *ChatSingleReq) Reset()         { *m = ChatSingleReq{} }
func (m *ChatSingleReq) String() string { return proto.CompactTextString(m) }
func (*ChatSingleReq) ProtoMessage()    {}
func (*ChatSingleReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
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

func (m *ChatSingleReq) GetClientInfo() *im_home_proto.ClientOnlineInfo {
	if m != nil {
		return m.ClientInfo
	}
	return nil
}

func (m *ChatSingleReq) GetData() *im_home_proto.ChatMessage {
	if m != nil {
		return m.Data
	}
	return nil
}

type ChatSingleRes struct {
	ErrCode im_error_proto.ErrCode     `protobuf:"varint,1,opt,name=ErrCode,proto3,enum=im_error_proto.ErrCode" json:"ErrCode,omitempty"`
	Data    *im_home_proto.ChatMessage `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *ChatSingleRes) Reset()         { *m = ChatSingleRes{} }
func (m *ChatSingleRes) String() string { return proto.CompactTextString(m) }
func (*ChatSingleRes) ProtoMessage()    {}
func (*ChatSingleRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
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

func (m *ChatSingleRes) GetErrCode() im_error_proto.ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return im_error_proto.ErrCode_success
}

func (m *ChatSingleRes) GetData() *im_home_proto.ChatMessage {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ChatSingleReq)(nil), "api_chat.ChatSingleReq")
	proto.RegisterType((*ChatSingleRes)(nil), "api_chat.ChatSingleRes")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x2c, 0xc8, 0x8c, 0x07, 0xf1, 0xa5, 0xf8,
	0x33, 0xf2, 0x73, 0x53, 0xe3, 0x11, 0x52, 0x52, 0x42, 0x60, 0x81, 0xc4, 0xe4, 0xe4, 0xfc, 0xd2,
	0x3c, 0x98, 0x18, 0x4f, 0x6a, 0x51, 0x51, 0x7e, 0x51, 0x31, 0x84, 0xa7, 0xd4, 0xc0, 0xc8, 0xc5,
	0xeb, 0x9c, 0x91, 0x58, 0x12, 0x9c, 0x99, 0x97, 0x9e, 0x93, 0x1a, 0x94, 0x5a, 0x28, 0x64, 0xcf,
	0xc5, 0xe5, 0x9c, 0x93, 0x99, 0x9a, 0x57, 0xe2, 0x99, 0x97, 0x96, 0x2f, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x6d, 0x24, 0xaf, 0x97, 0x99, 0x1b, 0x0f, 0x36, 0x0b, 0xac, 0x4b, 0x0f, 0xa2, 0xc0, 0x3f,
	0x2f, 0x27, 0x33, 0x2f, 0x15, 0xa4, 0x2c, 0x08, 0x49, 0x8b, 0x90, 0x1e, 0x17, 0x8b, 0x4b, 0x62,
	0x49, 0xa2, 0x04, 0x13, 0x58, 0xab, 0x14, 0xba, 0xd6, 0x8c, 0xc4, 0x12, 0xdf, 0xd4, 0xe2, 0xe2,
	0xc4, 0xf4, 0xd4, 0x20, 0xb0, 0x3a, 0xa5, 0x22, 0x54, 0x17, 0x14, 0x0b, 0x19, 0x72, 0xb1, 0xbb,
	0x16, 0x15, 0x39, 0xe7, 0xa7, 0xa4, 0x82, 0xad, 0xe7, 0x33, 0x12, 0x07, 0x99, 0x01, 0x76, 0x36,
	0xd4, 0x10, 0xa8, 0x74, 0x10, 0x4c, 0x1d, 0xa9, 0x76, 0x1a, 0x79, 0x70, 0xb1, 0x80, 0x04, 0x85,
	0x1c, 0xb8, 0xb8, 0x10, 0x76, 0x0b, 0x89, 0xeb, 0xc1, 0x82, 0x52, 0x0f, 0x25, 0x4c, 0xa4, 0x70,
	0x48, 0x14, 0x2b, 0x31, 0x38, 0x49, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83,
	0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43,
	0x12, 0x1b, 0xd8, 0x72, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0xad, 0x4b, 0xc6, 0xac,
	0x01, 0x00, 0x00,
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
			i = encodeVarintChat(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ClientInfo != nil {
		{
			size, err := m.ClientInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintChat(dAtA, i, uint64(size))
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
			i = encodeVarintChat(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ErrCode != 0 {
		i = encodeVarintChat(dAtA, i, uint64(m.ErrCode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintChat(dAtA []byte, offset int, v uint64) int {
	offset -= sovChat(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChatSingleReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClientInfo != nil {
		l = m.ClientInfo.Size()
		n += 1 + l + sovChat(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovChat(uint64(l))
	}
	return n
}

func (m *ChatSingleRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ErrCode != 0 {
		n += 1 + sovChat(uint64(m.ErrCode))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovChat(uint64(l))
	}
	return n
}

func sovChat(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChat(x uint64) (n int) {
	return sovChat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChatSingleReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChat
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
				return fmt.Errorf("proto: wrong wireType = %d for field ClientInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
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
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClientInfo == nil {
				m.ClientInfo = &im_home_proto.ClientOnlineInfo{}
			}
			if err := m.ClientInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
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
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &im_home_proto.ChatMessage{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChat
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
				return ErrIntOverflowChat
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrCode", wireType)
			}
			m.ErrCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrCode |= im_error_proto.ErrCode(b&0x7F) << shift
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
					return ErrIntOverflowChat
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
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &im_home_proto.ChatMessage{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChat
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
func skipChat(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChat
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
					return 0, ErrIntOverflowChat
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
					return 0, ErrIntOverflowChat
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
				return 0, ErrInvalidLengthChat
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChat
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChat
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChat        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChat          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChat = fmt.Errorf("proto: unexpected end of group")
)
