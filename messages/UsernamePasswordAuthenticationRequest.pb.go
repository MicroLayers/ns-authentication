// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages/UsernamePasswordAuthenticationRequest.proto

package messages

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

type UsernamePasswordAuthenticationRequest struct {
	RequestType          string   `protobuf:"bytes,1,opt,name=requestType,proto3" json:"requestType,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsernamePasswordAuthenticationRequest) Reset()         { *m = UsernamePasswordAuthenticationRequest{} }
func (m *UsernamePasswordAuthenticationRequest) String() string { return proto.CompactTextString(m) }
func (*UsernamePasswordAuthenticationRequest) ProtoMessage()    {}
func (*UsernamePasswordAuthenticationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_UsernamePasswordAuthenticationRequest_596c244c52c4a2b2, []int{0}
}
func (m *UsernamePasswordAuthenticationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsernamePasswordAuthenticationRequest.Unmarshal(m, b)
}
func (m *UsernamePasswordAuthenticationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsernamePasswordAuthenticationRequest.Marshal(b, m, deterministic)
}
func (dst *UsernamePasswordAuthenticationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsernamePasswordAuthenticationRequest.Merge(dst, src)
}
func (m *UsernamePasswordAuthenticationRequest) XXX_Size() int {
	return xxx_messageInfo_UsernamePasswordAuthenticationRequest.Size(m)
}
func (m *UsernamePasswordAuthenticationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UsernamePasswordAuthenticationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UsernamePasswordAuthenticationRequest proto.InternalMessageInfo

func (m *UsernamePasswordAuthenticationRequest) GetRequestType() string {
	if m != nil {
		return m.RequestType
	}
	return ""
}

func (m *UsernamePasswordAuthenticationRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UsernamePasswordAuthenticationRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*UsernamePasswordAuthenticationRequest)(nil), "messages.UsernamePasswordAuthenticationRequest")
}

func init() {
	proto.RegisterFile("messages/UsernamePasswordAuthenticationRequest.proto", fileDescriptor_UsernamePasswordAuthenticationRequest_596c244c52c4a2b2)
}

var fileDescriptor_UsernamePasswordAuthenticationRequest_596c244c52c4a2b2 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xc9, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x0f, 0x2d, 0x4e, 0x2d, 0xca, 0x4b, 0xcc, 0x4d, 0x0d, 0x48, 0x2c,
	0x2e, 0x2e, 0xcf, 0x2f, 0x4a, 0x71, 0x2c, 0x2d, 0xc9, 0x48, 0xcd, 0x2b, 0xc9, 0x4c, 0x4e, 0x2c,
	0xc9, 0xcc, 0xcf, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x80, 0xe9, 0x52, 0x6a, 0x64, 0xe4, 0x52, 0x25, 0x4a, 0xa7, 0x90, 0x02, 0x17, 0x77,
	0x11, 0x84, 0x19, 0x52, 0x59, 0x90, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x2c, 0x24,
	0x24, 0xc5, 0xc5, 0x51, 0x0a, 0x35, 0x4a, 0x82, 0x09, 0x2c, 0x0d, 0xe7, 0x83, 0xe4, 0x0a, 0xa0,
	0xc6, 0x4b, 0x30, 0x43, 0xe4, 0x60, 0xfc, 0x24, 0x36, 0xb0, 0xa3, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xfd, 0xb3, 0xf5, 0x06, 0xcc, 0x00, 0x00, 0x00,
}
