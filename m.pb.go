// Code generated by protoc-gen-go. DO NOT EDIT.
// source: m.proto

package m

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MyMessage struct {
	Msg                  *any.Any `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyMessage) Reset()         { *m = MyMessage{} }
func (m *MyMessage) String() string { return proto.CompactTextString(m) }
func (*MyMessage) ProtoMessage()    {}
func (*MyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e37c38dd62f49b, []int{0}
}

func (m *MyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyMessage.Unmarshal(m, b)
}
func (m *MyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyMessage.Marshal(b, m, deterministic)
}
func (m *MyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyMessage.Merge(m, src)
}
func (m *MyMessage) XXX_Size() int {
	return xxx_messageInfo_MyMessage.Size(m)
}
func (m *MyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MyMessage proto.InternalMessageInfo

func (m *MyMessage) GetMsg() *any.Any {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterType((*MyMessage)(nil), "MyMessage")
}

func init() { proto.RegisterFile("m.proto", fileDescriptor_28e37c38dd62f49b) }

var fileDescriptor_28e37c38dd62f49b = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x62, 0xcf, 0xd5, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4c, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a,
	0xd3, 0xf4, 0x13, 0xf3, 0x2a, 0x21, 0x52, 0x4a, 0xc6, 0x5c, 0x9c, 0xbe, 0x95, 0xbe, 0xa9, 0xc5,
	0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x6a, 0x5c, 0xcc, 0xb9, 0xc5, 0xe9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0xdc, 0x46, 0x22, 0x7a, 0x10, 0x5d, 0x7a, 0x30, 0x5d, 0x7a, 0x8e, 0x79, 0x95, 0x41, 0x20, 0x05,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x80, 0x07, 0xc7, 0x59, 0x00, 0x00, 0x00,
}