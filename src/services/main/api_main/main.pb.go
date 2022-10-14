// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: main.proto

package api_main

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	im_main_proto "github.com/mimis-s/IM-Service/src/common/commonproto/im_main_proto"
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

type ClientRequestHandleReq struct {
	MsgID   uint32                          `protobuf:"varint,1,opt,name=msgID,proto3" json:"msgID,omitempty"`
	Payload []byte                          `protobuf:"bytes,2,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Client  *im_main_proto.ClientOnlineInfo `protobuf:"bytes,3,opt,name=Client,proto3" json:"Client,omitempty"`
}

func (m *ClientRequestHandleReq) Reset()         { *m = ClientRequestHandleReq{} }
func (m *ClientRequestHandleReq) String() string { return proto.CompactTextString(m) }
func (*ClientRequestHandleReq) ProtoMessage()    {}
func (*ClientRequestHandleReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{0}
}
func (m *ClientRequestHandleReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientRequestHandleReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientRequestHandleReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientRequestHandleReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientRequestHandleReq.Merge(m, src)
}
func (m *ClientRequestHandleReq) XXX_Size() int {
	return m.Size()
}
func (m *ClientRequestHandleReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientRequestHandleReq.DiscardUnknown(m)
}

var xxx_messageInfo_ClientRequestHandleReq proto.InternalMessageInfo

func (m *ClientRequestHandleReq) GetMsgID() uint32 {
	if m != nil {
		return m.MsgID
	}
	return 0
}

func (m *ClientRequestHandleReq) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ClientRequestHandleReq) GetClient() *im_main_proto.ClientOnlineInfo {
	if m != nil {
		return m.Client
	}
	return nil
}

type ClientRequestHandleRes struct {
	ErrCode int32                           `protobuf:"varint,1,opt,name=ErrCode,proto3" json:"ErrCode,omitempty"`
	MsgID   uint32                          `protobuf:"varint,2,opt,name=msgID,proto3" json:"msgID,omitempty"`
	Payload []byte                          `protobuf:"bytes,3,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Client  *im_main_proto.ClientOnlineInfo `protobuf:"bytes,4,opt,name=Client,proto3" json:"Client,omitempty"`
}

func (m *ClientRequestHandleRes) Reset()         { *m = ClientRequestHandleRes{} }
func (m *ClientRequestHandleRes) String() string { return proto.CompactTextString(m) }
func (*ClientRequestHandleRes) ProtoMessage()    {}
func (*ClientRequestHandleRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{1}
}
func (m *ClientRequestHandleRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientRequestHandleRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientRequestHandleRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientRequestHandleRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientRequestHandleRes.Merge(m, src)
}
func (m *ClientRequestHandleRes) XXX_Size() int {
	return m.Size()
}
func (m *ClientRequestHandleRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientRequestHandleRes.DiscardUnknown(m)
}

var xxx_messageInfo_ClientRequestHandleRes proto.InternalMessageInfo

func (m *ClientRequestHandleRes) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *ClientRequestHandleRes) GetMsgID() uint32 {
	if m != nil {
		return m.MsgID
	}
	return 0
}

func (m *ClientRequestHandleRes) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ClientRequestHandleRes) GetClient() *im_main_proto.ClientOnlineInfo {
	if m != nil {
		return m.Client
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientRequestHandleReq)(nil), "api_main.ClientRequestHandleReq")
	proto.RegisterType((*ClientRequestHandleRes)(nil), "api_main.ClientRequestHandleRes")
}

func init() { proto.RegisterFile("main.proto", fileDescriptor_7ed94b0a22d11796) }

var fileDescriptor_7ed94b0a22d11796 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4d, 0xcc, 0xcc,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x2c, 0xc8, 0x8c, 0x07, 0xf1, 0xa5, 0x84,
	0x40, 0x64, 0x7c, 0x62, 0x72, 0x72, 0x7e, 0x69, 0x5e, 0x09, 0x44, 0x56, 0xa9, 0x91, 0x91, 0x4b,
	0xcc, 0x39, 0x27, 0x33, 0x35, 0xaf, 0x24, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0xc4, 0x23, 0x31,
	0x2f, 0x25, 0x27, 0x35, 0x28, 0xb5, 0x50, 0x48, 0x84, 0x8b, 0x35, 0xb7, 0x38, 0xdd, 0xd3, 0x45,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x37, 0x08, 0xc2, 0x11, 0x92, 0xe0, 0x62, 0x0f, 0x48, 0xac, 0xcc,
	0xc9, 0x4f, 0x4c, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0x71, 0x85, 0xcc, 0xb9, 0xd8,
	0x20, 0x26, 0x49, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x1b, 0xc9, 0xeb, 0x65, 0xe6, 0x82, 0x2d, 0x8e,
	0x07, 0x5b, 0xa5, 0x07, 0x91, 0xf4, 0xcf, 0xcb, 0xc9, 0xcc, 0x4b, 0xf5, 0xcc, 0x4b, 0xcb, 0x0f,
	0x82, 0x2a, 0x57, 0x9a, 0x8d, 0xcb, 0x0d, 0xc5, 0x20, 0xdb, 0x5c, 0x8b, 0x8a, 0x9c, 0xf3, 0x53,
	0x52, 0xc1, 0xae, 0x60, 0x0d, 0x82, 0x71, 0x11, 0xae, 0x63, 0xc2, 0xe1, 0x3a, 0x66, 0x5c, 0xae,
	0x63, 0x21, 0xc9, 0x75, 0x46, 0x27, 0x19, 0xb9, 0x58, 0x7c, 0x13, 0x33, 0xf3, 0x84, 0x12, 0xb8,
	0x24, 0xb0, 0xb8, 0x32, 0x00, 0x1c, 0xc8, 0x0a, 0x7a, 0xb0, 0x50, 0xd6, 0xc3, 0x1e, 0x9a, 0x52,
	0x84, 0x54, 0x14, 0x2b, 0x31, 0x08, 0xc5, 0x73, 0x89, 0x63, 0x91, 0xf3, 0x2a, 0xce, 0xcf, 0xa3,
	0x8e, 0x05, 0x4e, 0x12, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c,
	0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06,
	0x0e, 0x05, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x8f, 0x7a, 0x69, 0x3a, 0x02, 0x00,
	0x00,
}

func (m *ClientRequestHandleReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientRequestHandleReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientRequestHandleReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Client != nil {
		{
			size, err := m.Client.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMain(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintMain(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x12
	}
	if m.MsgID != 0 {
		i = encodeVarintMain(dAtA, i, uint64(m.MsgID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ClientRequestHandleRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientRequestHandleRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientRequestHandleRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Client != nil {
		{
			size, err := m.Client.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMain(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintMain(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x1a
	}
	if m.MsgID != 0 {
		i = encodeVarintMain(dAtA, i, uint64(m.MsgID))
		i--
		dAtA[i] = 0x10
	}
	if m.ErrCode != 0 {
		i = encodeVarintMain(dAtA, i, uint64(m.ErrCode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMain(dAtA []byte, offset int, v uint64) int {
	offset -= sovMain(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientRequestHandleReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MsgID != 0 {
		n += 1 + sovMain(uint64(m.MsgID))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovMain(uint64(l))
	}
	if m.Client != nil {
		l = m.Client.Size()
		n += 1 + l + sovMain(uint64(l))
	}
	return n
}

func (m *ClientRequestHandleRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ErrCode != 0 {
		n += 1 + sovMain(uint64(m.ErrCode))
	}
	if m.MsgID != 0 {
		n += 1 + sovMain(uint64(m.MsgID))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovMain(uint64(l))
	}
	if m.Client != nil {
		l = m.Client.Size()
		n += 1 + l + sovMain(uint64(l))
	}
	return n
}

func sovMain(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMain(x uint64) (n int) {
	return sovMain(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientRequestHandleReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMain
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
			return fmt.Errorf("proto: ClientRequestHandleReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientRequestHandleReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgID", wireType)
			}
			m.MsgID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMain
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Client", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMain
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
				return ErrInvalidLengthMain
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Client == nil {
				m.Client = &im_main_proto.ClientOnlineInfo{}
			}
			if err := m.Client.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMain
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
func (m *ClientRequestHandleRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMain
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
			return fmt.Errorf("proto: ClientRequestHandleRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientRequestHandleRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrCode", wireType)
			}
			m.ErrCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgID", wireType)
			}
			m.MsgID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMain
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Client", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMain
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
				return ErrInvalidLengthMain
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Client == nil {
				m.Client = &im_main_proto.ClientOnlineInfo{}
			}
			if err := m.Client.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMain
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
func skipMain(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMain
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
					return 0, ErrIntOverflowMain
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
					return 0, ErrIntOverflowMain
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
				return 0, ErrInvalidLengthMain
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMain
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMain
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMain        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMain          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMain = fmt.Errorf("proto: unexpected end of group")
)
