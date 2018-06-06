// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/proto/verify/verify.proto

/*
Package verify is a generated protocol buffer package.

It is generated from these files:
	auth/proto/verify/verify.proto

It has these top-level messages:
	VerifyReq
	VerifyResp
*/
package verify

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

type VerifyReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Pwd  string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *VerifyReq) Reset()                    { *m = VerifyReq{} }
func (m *VerifyReq) String() string            { return proto.CompactTextString(m) }
func (*VerifyReq) ProtoMessage()               {}
func (*VerifyReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VerifyReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VerifyReq) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type VerifyResp struct {
	Token    string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Verified bool   `protobuf:"varint,2,opt,name=verified" json:"verified,omitempty"`
}

func (m *VerifyResp) Reset()                    { *m = VerifyResp{} }
func (m *VerifyResp) String() string            { return proto.CompactTextString(m) }
func (*VerifyResp) ProtoMessage()               {}
func (*VerifyResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *VerifyResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *VerifyResp) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

func init() {
	proto.RegisterType((*VerifyReq)(nil), "verify.VerifyReq")
	proto.RegisterType((*VerifyResp)(nil), "verify.VerifyResp")
}

func init() { proto.RegisterFile("auth/proto/verify/verify.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x4b, 0x2d, 0xca, 0x4c, 0xab, 0x84, 0x52, 0x7a,
	0x60, 0x31, 0x21, 0x36, 0x08, 0x4f, 0xc9, 0x90, 0x8b, 0x33, 0x0c, 0xcc, 0x0a, 0x4a, 0x2d, 0x14,
	0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3,
	0x85, 0x04, 0xb8, 0x98, 0x0b, 0xca, 0x53, 0x24, 0x98, 0xc0, 0x42, 0x20, 0xa6, 0x92, 0x1d, 0x17,
	0x17, 0x4c, 0x4b, 0x71, 0x81, 0x90, 0x08, 0x17, 0x6b, 0x49, 0x7e, 0x76, 0x6a, 0x1e, 0x54, 0x13,
	0x84, 0x23, 0x24, 0xc5, 0xc5, 0x01, 0xb6, 0x20, 0x33, 0x15, 0xa2, 0x95, 0x23, 0x08, 0xce, 0x4f,
	0x62, 0x03, 0xbb, 0xc0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x87, 0x09, 0x75, 0x81, 0xa3, 0x00,
	0x00, 0x00,
}