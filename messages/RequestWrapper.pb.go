// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages/RequestWrapper.proto

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

type RequestWrapper struct {
	RequestType          string   `protobuf:"bytes,1,opt,name=requestType,proto3" json:"requestType,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestWrapper) Reset()         { *m = RequestWrapper{} }
func (m *RequestWrapper) String() string { return proto.CompactTextString(m) }
func (*RequestWrapper) ProtoMessage()    {}
func (*RequestWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_RequestWrapper_a3728b92a618eb0a, []int{0}
}
func (m *RequestWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestWrapper.Unmarshal(m, b)
}
func (m *RequestWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestWrapper.Marshal(b, m, deterministic)
}
func (dst *RequestWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestWrapper.Merge(dst, src)
}
func (m *RequestWrapper) XXX_Size() int {
	return xxx_messageInfo_RequestWrapper.Size(m)
}
func (m *RequestWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_RequestWrapper proto.InternalMessageInfo

func (m *RequestWrapper) GetRequestType() string {
	if m != nil {
		return m.RequestType
	}
	return ""
}

func (m *RequestWrapper) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestWrapper)(nil), "messages.RequestWrapper")
}

func init() {
	proto.RegisterFile("messages/RequestWrapper.proto", fileDescriptor_RequestWrapper_a3728b92a618eb0a)
}

var fileDescriptor_RequestWrapper_a3728b92a618eb0a = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x09, 0x2f, 0x4a, 0x2c,
	0x28, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0x49, 0x2b, 0xf9, 0x70,
	0xf1, 0xa1, 0xaa, 0x10, 0x52, 0xe0, 0xe2, 0x2e, 0x82, 0x88, 0x84, 0x54, 0x16, 0xa4, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x0b, 0x09, 0x49, 0x70, 0xb1, 0x17, 0x24, 0x56, 0xe6, 0xe4,
	0x27, 0xa6, 0x48, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x04, 0xc1, 0xb8, 0x49, 0x6c, 0x60, 0xe3, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x26, 0x91, 0x9e, 0x7f, 0x00, 0x00, 0x00,
}