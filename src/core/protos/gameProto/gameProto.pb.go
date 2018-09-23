// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gameProto.proto

/*
Package gameProto is a generated protocol buffer package.

It is generated from these files:
	gameProto.proto

It has these top-level messages:
	ErrorNoticeS2C
	ClientPingC2S
	UserLoginC2S
	UserLoginS2C
	UserOtherLoginNoticeS2C
	UserInfo
	UserGetInfoC2S
	UserGetInfoS2C
*/
package gameProto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 错误消息(500)
type ErrorNoticeS2C struct {
	ErrorCode        *int32 `protobuf:"varint,1,req,name=errorCode" json:"errorCode,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ErrorNoticeS2C) Reset()                    { *m = ErrorNoticeS2C{} }
func (m *ErrorNoticeS2C) String() string            { return proto.CompactTextString(m) }
func (*ErrorNoticeS2C) ProtoMessage()               {}
func (*ErrorNoticeS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ErrorNoticeS2C) GetErrorCode() int32 {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

// 客户端ping(1001)
type ClientPingC2S struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ClientPingC2S) Reset()                    { *m = ClientPingC2S{} }
func (m *ClientPingC2S) String() string            { return proto.CompactTextString(m) }
func (*ClientPingC2S) ProtoMessage()               {}
func (*ClientPingC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 用户登录C2S(2001)
type UserLoginC2S struct {
	Account          *string `protobuf:"bytes,1,req,name=account" json:"account,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserLoginC2S) Reset()                    { *m = UserLoginC2S{} }
func (m *UserLoginC2S) String() string            { return proto.CompactTextString(m) }
func (*UserLoginC2S) ProtoMessage()               {}
func (*UserLoginC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserLoginC2S) GetAccount() string {
	if m != nil && m.Account != nil {
		return *m.Account
	}
	return ""
}

// 用户登录S2C(2002)
type UserLoginS2C struct {
	Token            *string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserLoginS2C) Reset()                    { *m = UserLoginS2C{} }
func (m *UserLoginS2C) String() string            { return proto.CompactTextString(m) }
func (*UserLoginS2C) ProtoMessage()               {}
func (*UserLoginS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UserLoginS2C) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

// 其他客户端登录S2C(2004)
type UserOtherLoginNoticeS2C struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *UserOtherLoginNoticeS2C) Reset()                    { *m = UserOtherLoginNoticeS2C{} }
func (m *UserOtherLoginNoticeS2C) String() string            { return proto.CompactTextString(m) }
func (*UserOtherLoginNoticeS2C) ProtoMessage()               {}
func (*UserOtherLoginNoticeS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// 用户数据
type UserInfo struct {
	Id               *uint64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Money            *int32  `protobuf:"varint,3,req,name=money" json:"money,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserInfo) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *UserInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *UserInfo) GetMoney() int32 {
	if m != nil && m.Money != nil {
		return *m.Money
	}
	return 0
}

// 获取用户信息C2S(3001)
type UserGetInfoC2S struct {
	Token            *string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserGetInfoC2S) Reset()                    { *m = UserGetInfoC2S{} }
func (m *UserGetInfoC2S) String() string            { return proto.CompactTextString(m) }
func (*UserGetInfoC2S) ProtoMessage()               {}
func (*UserGetInfoC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UserGetInfoC2S) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

// 获取用户信息S2C(3002)
type UserGetInfoS2C struct {
	Data             *UserInfo `protobuf:"bytes,1,req,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *UserGetInfoS2C) Reset()                    { *m = UserGetInfoS2C{} }
func (m *UserGetInfoS2C) String() string            { return proto.CompactTextString(m) }
func (*UserGetInfoS2C) ProtoMessage()               {}
func (*UserGetInfoS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UserGetInfoS2C) GetData() *UserInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ErrorNoticeS2C)(nil), "error_notice_s2c")
	proto.RegisterType((*ClientPingC2S)(nil), "client_ping_c2s")
	proto.RegisterType((*UserLoginC2S)(nil), "user_login_c2s")
	proto.RegisterType((*UserLoginS2C)(nil), "user_login_s2c")
	proto.RegisterType((*UserOtherLoginNoticeS2C)(nil), "user_otherLogin_notice_s2c")
	proto.RegisterType((*UserInfo)(nil), "userInfo")
	proto.RegisterType((*UserGetInfoC2S)(nil), "user_getInfo_c2s")
	proto.RegisterType((*UserGetInfoS2C)(nil), "user_getInfo_s2c")
}

func init() { proto.RegisterFile("gameProto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x31, 0xa6, 0x68, 0x46, 0x68, 0xeb, 0xe2, 0x21, 0x48, 0x05, 0xd9, 0x83, 0x14, 0x0f,
	0x45, 0xf3, 0x17, 0xf4, 0x22, 0x78, 0x90, 0xfc, 0x81, 0x65, 0xd9, 0x8c, 0x71, 0xb1, 0x99, 0x29,
	0x9b, 0xe9, 0xc1, 0x7f, 0x2f, 0x3b, 0xb5, 0x04, 0xc5, 0x4b, 0xc8, 0x7b, 0xfb, 0xed, 0x7b, 0x6f,
	0x61, 0xd1, 0xfb, 0x01, 0xdf, 0x12, 0x0b, 0x6f, 0x76, 0xf9, 0x6b, 0x1f, 0x60, 0x89, 0x29, 0x71,
	0x72, 0xc4, 0x12, 0x03, 0xba, 0xb1, 0x09, 0x66, 0x05, 0x95, 0x7a, 0x4f, 0xdc, 0x61, 0x7d, 0x72,
	0x5b, 0xac, 0x67, 0xed, 0x64, 0xd8, 0x4b, 0x58, 0x84, 0x6d, 0x44, 0x12, 0xb7, 0x8b, 0xd4, 0xbb,
	0xd0, 0x8c, 0xf6, 0x1e, 0xe6, 0xfb, 0x11, 0x93, 0xdb, 0x72, 0x1f, 0x29, 0x3b, 0xa6, 0x86, 0x33,
	0x1f, 0x02, 0xef, 0x49, 0x34, 0xa0, 0x6a, 0x8f, 0xd2, 0xde, 0xfd, 0x62, 0x73, 0xdd, 0x15, 0xcc,
	0x84, 0x3f, 0x91, 0x7e, 0xc8, 0x83, 0xb0, 0x2b, 0xb8, 0x56, 0x8e, 0xe5, 0x03, 0xd3, 0xab, 0xc2,
	0xd3, 0x44, 0xfb, 0x0c, 0xe7, 0xf9, 0xf4, 0x85, 0xde, 0xd9, 0xcc, 0xa1, 0x88, 0x9d, 0x5e, 0x2e,
	0xdb, 0x22, 0x76, 0xc6, 0x40, 0x49, 0x7e, 0xc0, 0xba, 0xd0, 0x38, 0xfd, 0xcf, 0x1d, 0x03, 0x13,
	0x7e, 0xd5, 0xa7, 0xfa, 0x9c, 0x83, 0xb0, 0x6b, 0x58, 0x6a, 0x47, 0x8f, 0x92, 0x93, 0x74, 0xf9,
	0xff, 0x6b, 0x1e, 0xff, 0x90, 0x79, 0xf7, 0x0d, 0x94, 0x9d, 0x17, 0xaf, 0xe0, 0x45, 0x53, 0x6d,
	0x8e, 0x83, 0x5a, 0xb5, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x87, 0x2d, 0x52, 0xda, 0x6b, 0x01,
	0x00, 0x00,
}
