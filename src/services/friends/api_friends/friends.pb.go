// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: friends.proto

package api_friends

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

// 获取好友列表
type GetFriendsListReq struct {
	ClientInfo *im_home_proto.ClientOnlineInfo  `protobuf:"bytes,1,opt,name=ClientInfo,proto3" json:"ClientInfo,omitempty"`
	Data       *im_home_proto.GetFriendsListReq `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *GetFriendsListReq) Reset()         { *m = GetFriendsListReq{} }
func (m *GetFriendsListReq) String() string { return proto.CompactTextString(m) }
func (*GetFriendsListReq) ProtoMessage()    {}
func (*GetFriendsListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef2cbb5c12e56bfb, []int{0}
}
func (m *GetFriendsListReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetFriendsListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetFriendsListReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetFriendsListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFriendsListReq.Merge(m, src)
}
func (m *GetFriendsListReq) XXX_Size() int {
	return m.Size()
}
func (m *GetFriendsListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFriendsListReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetFriendsListReq proto.InternalMessageInfo

func (m *GetFriendsListReq) GetClientInfo() *im_home_proto.ClientOnlineInfo {
	if m != nil {
		return m.ClientInfo
	}
	return nil
}

func (m *GetFriendsListReq) GetData() *im_home_proto.GetFriendsListReq {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetFriendsListRes struct {
	ErrCode im_error_proto.ErrCode           `protobuf:"varint,1,opt,name=ErrCode,proto3,enum=im_error_proto.ErrCode" json:"ErrCode,omitempty"`
	Data    *im_home_proto.GetFriendsListRes `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *GetFriendsListRes) Reset()         { *m = GetFriendsListRes{} }
func (m *GetFriendsListRes) String() string { return proto.CompactTextString(m) }
func (*GetFriendsListRes) ProtoMessage()    {}
func (*GetFriendsListRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef2cbb5c12e56bfb, []int{1}
}
func (m *GetFriendsListRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetFriendsListRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetFriendsListRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetFriendsListRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFriendsListRes.Merge(m, src)
}
func (m *GetFriendsListRes) XXX_Size() int {
	return m.Size()
}
func (m *GetFriendsListRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFriendsListRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetFriendsListRes proto.InternalMessageInfo

func (m *GetFriendsListRes) GetErrCode() im_error_proto.ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return im_error_proto.ErrCode_success
}

func (m *GetFriendsListRes) GetData() *im_home_proto.GetFriendsListRes {
	if m != nil {
		return m.Data
	}
	return nil
}

// 添加好友
type ApplyFriendsReq struct {
	ClientInfo *im_home_proto.ClientOnlineInfo `protobuf:"bytes,1,opt,name=ClientInfo,proto3" json:"ClientInfo,omitempty"`
	Data       *im_home_proto.ApplyFriendsReq  `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *ApplyFriendsReq) Reset()         { *m = ApplyFriendsReq{} }
func (m *ApplyFriendsReq) String() string { return proto.CompactTextString(m) }
func (*ApplyFriendsReq) ProtoMessage()    {}
func (*ApplyFriendsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef2cbb5c12e56bfb, []int{2}
}
func (m *ApplyFriendsReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApplyFriendsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApplyFriendsReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApplyFriendsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyFriendsReq.Merge(m, src)
}
func (m *ApplyFriendsReq) XXX_Size() int {
	return m.Size()
}
func (m *ApplyFriendsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyFriendsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyFriendsReq proto.InternalMessageInfo

func (m *ApplyFriendsReq) GetClientInfo() *im_home_proto.ClientOnlineInfo {
	if m != nil {
		return m.ClientInfo
	}
	return nil
}

func (m *ApplyFriendsReq) GetData() *im_home_proto.ApplyFriendsReq {
	if m != nil {
		return m.Data
	}
	return nil
}

type ApplyFriendsRes struct {
	ErrCode im_error_proto.ErrCode         `protobuf:"varint,1,opt,name=ErrCode,proto3,enum=im_error_proto.ErrCode" json:"ErrCode,omitempty"`
	Data    *im_home_proto.ApplyFriendsRes `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *ApplyFriendsRes) Reset()         { *m = ApplyFriendsRes{} }
func (m *ApplyFriendsRes) String() string { return proto.CompactTextString(m) }
func (*ApplyFriendsRes) ProtoMessage()    {}
func (*ApplyFriendsRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef2cbb5c12e56bfb, []int{3}
}
func (m *ApplyFriendsRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApplyFriendsRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApplyFriendsRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApplyFriendsRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyFriendsRes.Merge(m, src)
}
func (m *ApplyFriendsRes) XXX_Size() int {
	return m.Size()
}
func (m *ApplyFriendsRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyFriendsRes.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyFriendsRes proto.InternalMessageInfo

func (m *ApplyFriendsRes) GetErrCode() im_error_proto.ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return im_error_proto.ErrCode_success
}

func (m *ApplyFriendsRes) GetData() *im_home_proto.ApplyFriendsRes {
	if m != nil {
		return m.Data
	}
	return nil
}

// 删除好友
type DelFriendsReq struct {
	ClientInfo *im_home_proto.ClientOnlineInfo `protobuf:"bytes,1,opt,name=ClientInfo,proto3" json:"ClientInfo,omitempty"`
	Data       *im_home_proto.DelFriendsReq    `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *DelFriendsReq) Reset()         { *m = DelFriendsReq{} }
func (m *DelFriendsReq) String() string { return proto.CompactTextString(m) }
func (*DelFriendsReq) ProtoMessage()    {}
func (*DelFriendsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef2cbb5c12e56bfb, []int{4}
}
func (m *DelFriendsReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelFriendsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelFriendsReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelFriendsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelFriendsReq.Merge(m, src)
}
func (m *DelFriendsReq) XXX_Size() int {
	return m.Size()
}
func (m *DelFriendsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DelFriendsReq.DiscardUnknown(m)
}

var xxx_messageInfo_DelFriendsReq proto.InternalMessageInfo

func (m *DelFriendsReq) GetClientInfo() *im_home_proto.ClientOnlineInfo {
	if m != nil {
		return m.ClientInfo
	}
	return nil
}

func (m *DelFriendsReq) GetData() *im_home_proto.DelFriendsReq {
	if m != nil {
		return m.Data
	}
	return nil
}

type DelFriendsRes struct {
	ErrCode im_error_proto.ErrCode       `protobuf:"varint,1,opt,name=ErrCode,proto3,enum=im_error_proto.ErrCode" json:"ErrCode,omitempty"`
	Data    *im_home_proto.DelFriendsRes `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *DelFriendsRes) Reset()         { *m = DelFriendsRes{} }
func (m *DelFriendsRes) String() string { return proto.CompactTextString(m) }
func (*DelFriendsRes) ProtoMessage()    {}
func (*DelFriendsRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef2cbb5c12e56bfb, []int{5}
}
func (m *DelFriendsRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelFriendsRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelFriendsRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelFriendsRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelFriendsRes.Merge(m, src)
}
func (m *DelFriendsRes) XXX_Size() int {
	return m.Size()
}
func (m *DelFriendsRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DelFriendsRes.DiscardUnknown(m)
}

var xxx_messageInfo_DelFriendsRes proto.InternalMessageInfo

func (m *DelFriendsRes) GetErrCode() im_error_proto.ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return im_error_proto.ErrCode_success
}

func (m *DelFriendsRes) GetData() *im_home_proto.DelFriendsRes {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetFriendsListReq)(nil), "api_friends.GetFriendsListReq")
	proto.RegisterType((*GetFriendsListRes)(nil), "api_friends.GetFriendsListRes")
	proto.RegisterType((*ApplyFriendsReq)(nil), "api_friends.ApplyFriendsReq")
	proto.RegisterType((*ApplyFriendsRes)(nil), "api_friends.ApplyFriendsRes")
	proto.RegisterType((*DelFriendsReq)(nil), "api_friends.DelFriendsReq")
	proto.RegisterType((*DelFriendsRes)(nil), "api_friends.DelFriendsRes")
}

func init() { proto.RegisterFile("friends.proto", fileDescriptor_ef2cbb5c12e56bfb) }

var fileDescriptor_ef2cbb5c12e56bfb = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x3d, 0x4b, 0x03, 0x41,
	0x14, 0xbc, 0x15, 0x31, 0xf0, 0xf2, 0x21, 0x6e, 0x63, 0x38, 0xc2, 0x1a, 0xae, 0xb2, 0x3a, 0xf4,
	0xb4, 0x17, 0x4d, 0x8c, 0x08, 0x01, 0xe1, 0xfe, 0x40, 0x38, 0x93, 0x0d, 0x1e, 0x5c, 0x6e, 0xcf,
	0xbd, 0x15, 0x14, 0xac, 0x2c, 0x2c, 0xac, 0xfc, 0x59, 0x96, 0x29, 0x2d, 0x25, 0xf9, 0x09, 0xfe,
	0x01, 0xc9, 0xe6, 0x05, 0xb3, 0x1b, 0x0c, 0x26, 0xa4, 0x7d, 0x33, 0xf3, 0x66, 0xde, 0xb0, 0x0b,
	0xe5, 0xbe, 0x8c, 0x79, 0xda, 0xcb, 0xfd, 0x4c, 0x0a, 0x25, 0x68, 0x31, 0xca, 0xe2, 0x0e, 0x8e,
	0x5c, 0x7a, 0x27, 0x06, 0xbc, 0x63, 0x10, 0x70, 0x16, 0x75, 0xbb, 0xe2, 0x21, 0x55, 0x38, 0x2b,
	0x71, 0x29, 0x85, 0x44, 0x86, 0xf7, 0x46, 0x60, 0xef, 0x8a, 0xab, 0xd6, 0x54, 0xd6, 0x8e, 0x73,
	0x15, 0xf2, 0x7b, 0x7a, 0x06, 0xd0, 0x48, 0x62, 0x9e, 0xaa, 0xeb, 0xb4, 0x2f, 0xaa, 0xa4, 0x4e,
	0x0e, 0x8b, 0xc1, 0x81, 0x1f, 0x0f, 0x3a, 0x7a, 0x9f, 0x56, 0xfa, 0x53, 0xc2, 0x4d, 0x9a, 0xc4,
	0x29, 0x9f, 0xd0, 0xc2, 0x39, 0x09, 0x3d, 0x85, 0xed, 0x66, 0xa4, 0xa2, 0xea, 0x96, 0x96, 0xd6,
	0x2d, 0xe9, 0x82, 0x61, 0xa8, 0xd9, 0xde, 0xf3, 0x62, 0x96, 0x9c, 0x1e, 0x43, 0xe1, 0x52, 0xca,
	0x86, 0xe8, 0x71, 0x1d, 0xa4, 0x12, 0xec, 0x4f, 0xb6, 0xe9, 0x23, 0x70, 0x1d, 0xc2, 0xe1, 0x8c,
	0xb7, 0xa2, 0x7b, 0x8e, 0xee, 0xaf, 0x04, 0x76, 0xcf, 0xb3, 0x2c, 0x79, 0x42, 0x74, 0x23, 0x45,
	0x04, 0x46, 0x14, 0x66, 0x49, 0x2d, 0x3b, 0x0c, 0xf2, 0x68, 0xe7, 0x58, 0xab, 0x84, 0x55, 0x9c,
	0x67, 0x15, 0xbc, 0x10, 0x28, 0x37, 0x79, 0xb2, 0xc9, 0x02, 0x8e, 0x8c, 0x18, 0x35, 0x4b, 0x6a,
	0x98, 0x61, 0x08, 0x65, 0x66, 0x58, 0xeb, 0xf8, 0xff, 0xbb, 0xe2, 0xe9, 0xc1, 0x37, 0x81, 0x02,
	0x0e, 0x69, 0x08, 0x15, 0xf3, 0x91, 0x50, 0xe6, 0xcf, 0x7d, 0xb5, 0xc5, 0xf7, 0xeb, 0x2e, 0xc7,
	0x73, 0xcf, 0xa1, 0x6d, 0x28, 0xcd, 0x77, 0x4e, 0x6b, 0x86, 0xc2, 0x7a, 0x08, 0xee, 0x32, 0x74,
	0xb2, 0xad, 0x05, 0xf0, 0x7b, 0x04, 0x75, 0x0d, 0xb6, 0xd1, 0xa9, 0xfb, 0x37, 0x96, 0x7b, 0xce,
	0x45, 0xf5, 0x63, 0xc4, 0xc8, 0x70, 0xc4, 0xc8, 0xd7, 0x88, 0x91, 0xf7, 0x31, 0x73, 0x86, 0x63,
	0xe6, 0x7c, 0x8e, 0x99, 0x73, 0xbb, 0xa3, 0xab, 0x3a, 0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0x91,
	0x7d, 0x8f, 0x23, 0x73, 0x04, 0x00, 0x00,
}

func (m *GetFriendsListReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetFriendsListReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetFriendsListReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintFriends(dAtA, i, uint64(size))
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
			i = encodeVarintFriends(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetFriendsListRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetFriendsListRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetFriendsListRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintFriends(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ErrCode != 0 {
		i = encodeVarintFriends(dAtA, i, uint64(m.ErrCode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ApplyFriendsReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApplyFriendsReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApplyFriendsReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintFriends(dAtA, i, uint64(size))
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
			i = encodeVarintFriends(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ApplyFriendsRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApplyFriendsRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApplyFriendsRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintFriends(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ErrCode != 0 {
		i = encodeVarintFriends(dAtA, i, uint64(m.ErrCode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DelFriendsReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelFriendsReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelFriendsReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintFriends(dAtA, i, uint64(size))
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
			i = encodeVarintFriends(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DelFriendsRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelFriendsRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelFriendsRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintFriends(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ErrCode != 0 {
		i = encodeVarintFriends(dAtA, i, uint64(m.ErrCode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintFriends(dAtA []byte, offset int, v uint64) int {
	offset -= sovFriends(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetFriendsListReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClientInfo != nil {
		l = m.ClientInfo.Size()
		n += 1 + l + sovFriends(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovFriends(uint64(l))
	}
	return n
}

func (m *GetFriendsListRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ErrCode != 0 {
		n += 1 + sovFriends(uint64(m.ErrCode))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovFriends(uint64(l))
	}
	return n
}

func (m *ApplyFriendsReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClientInfo != nil {
		l = m.ClientInfo.Size()
		n += 1 + l + sovFriends(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovFriends(uint64(l))
	}
	return n
}

func (m *ApplyFriendsRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ErrCode != 0 {
		n += 1 + sovFriends(uint64(m.ErrCode))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovFriends(uint64(l))
	}
	return n
}

func (m *DelFriendsReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClientInfo != nil {
		l = m.ClientInfo.Size()
		n += 1 + l + sovFriends(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovFriends(uint64(l))
	}
	return n
}

func (m *DelFriendsRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ErrCode != 0 {
		n += 1 + sovFriends(uint64(m.ErrCode))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovFriends(uint64(l))
	}
	return n
}

func sovFriends(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFriends(x uint64) (n int) {
	return sovFriends(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetFriendsListReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFriends
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
			return fmt.Errorf("proto: GetFriendsListReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetFriendsListReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFriends
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
				return ErrInvalidLengthFriends
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFriends
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
					return ErrIntOverflowFriends
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
				return ErrInvalidLengthFriends
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFriends
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &im_home_proto.GetFriendsListReq{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFriends(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFriends
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
func (m *GetFriendsListRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFriends
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
			return fmt.Errorf("proto: GetFriendsListRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetFriendsListRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrCode", wireType)
			}
			m.ErrCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFriends
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
					return ErrIntOverflowFriends
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
				return ErrInvalidLengthFriends
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFriends
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &im_home_proto.GetFriendsListRes{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFriends(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFriends
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
func (m *ApplyFriendsReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFriends
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
			return fmt.Errorf("proto: ApplyFriendsReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApplyFriendsReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFriends
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
				return ErrInvalidLengthFriends
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFriends
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
					return ErrIntOverflowFriends
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
				return ErrInvalidLengthFriends
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFriends
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &im_home_proto.ApplyFriendsReq{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFriends(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFriends
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
func (m *ApplyFriendsRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFriends
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
			return fmt.Errorf("proto: ApplyFriendsRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApplyFriendsRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrCode", wireType)
			}
			m.ErrCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFriends
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
					return ErrIntOverflowFriends
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
				return ErrInvalidLengthFriends
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFriends
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &im_home_proto.ApplyFriendsRes{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFriends(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFriends
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
func (m *DelFriendsReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFriends
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
			return fmt.Errorf("proto: DelFriendsReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelFriendsReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFriends
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
				return ErrInvalidLengthFriends
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFriends
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
					return ErrIntOverflowFriends
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
				return ErrInvalidLengthFriends
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFriends
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &im_home_proto.DelFriendsReq{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFriends(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFriends
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
func (m *DelFriendsRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFriends
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
			return fmt.Errorf("proto: DelFriendsRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelFriendsRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrCode", wireType)
			}
			m.ErrCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFriends
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
					return ErrIntOverflowFriends
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
				return ErrInvalidLengthFriends
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFriends
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &im_home_proto.DelFriendsRes{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFriends(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFriends
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
func skipFriends(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFriends
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
					return 0, ErrIntOverflowFriends
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
					return 0, ErrIntOverflowFriends
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
				return 0, ErrInvalidLengthFriends
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFriends
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFriends
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFriends        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFriends          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFriends = fmt.Errorf("proto: unexpected end of group")
)
