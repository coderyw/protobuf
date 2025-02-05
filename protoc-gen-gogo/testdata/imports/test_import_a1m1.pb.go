// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test_import_a1m1.proto

package imports

import (
	fmt "fmt"
	proto "github.com/coderyw/protobuf/proto"
	test_a_1 "github.com/coderyw/protobuf/protoc-gen-gogo/testdata/imports/test_a_1"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type A1M1 struct {
	F                    *test_a_1.M1 `protobuf:"bytes,1,opt,name=f,proto3" json:"f,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *A1M1) Reset()         { *m = A1M1{} }
func (m *A1M1) String() string { return proto.CompactTextString(m) }
func (*A1M1) ProtoMessage()    {}
func (*A1M1) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbb6455be6296fbe, []int{0}
}
func (m *A1M1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_A1M1.Unmarshal(m, b)
}
func (m *A1M1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_A1M1.Marshal(b, m, deterministic)
}
func (m *A1M1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_A1M1.Merge(m, src)
}
func (m *A1M1) XXX_Size() int {
	return xxx_messageInfo_A1M1.Size(m)
}
func (m *A1M1) XXX_DiscardUnknown() {
	xxx_messageInfo_A1M1.DiscardUnknown(m)
}

var xxx_messageInfo_A1M1 proto.InternalMessageInfo

func (m *A1M1) GetF() *test_a_1.M1 {
	if m != nil {
		return m.F
	}
	return nil
}

func (m *A1M1) SetF_(val *test_a_1.M1) {
	if m != nil {
		m.F = val
	}
	return
}

func (m *A1M1) IsNil() bool {
	return m == nil

}

func init() {
	proto.RegisterType((*A1M1)(nil), "test.A1M1")
}

func init() { proto.RegisterFile("test_import_a1m1.proto", fileDescriptor_fbb6455be6296fbe) }

var fileDescriptor_fbb6455be6296fbe = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x49, 0x2d, 0x2e,
	0x89, 0xcf, 0xcc, 0x2d, 0xc8, 0x2f, 0x2a, 0x89, 0x4f, 0x34, 0xcc, 0x35, 0xd4, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x01, 0x89, 0x4b, 0x49, 0x42, 0x24, 0x8a, 0xf5, 0xc1, 0xaa, 0x12, 0xe3,
	0x0d, 0xf5, 0x61, 0x0a, 0x94, 0x14, 0xb8, 0x58, 0x1c, 0x0d, 0x7d, 0x0d, 0x85, 0x24, 0xb8, 0x18,
	0xd3, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xb8, 0xf4, 0x40, 0xca, 0xf4, 0x12, 0xf5, 0x7c,
	0x0d, 0x83, 0x18, 0xd3, 0x9c, 0xec, 0xa2, 0x6c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92,
	0xf3, 0x73, 0xf5, 0x93, 0xf3, 0x53, 0x52, 0x8b, 0x2a, 0xcb, 0xf5, 0xc1, 0xba, 0x93, 0x4a, 0xd3,
	0x20, 0x8c, 0x64, 0xdd, 0xf4, 0xd4, 0x3c, 0xdd, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0x15, 0x29, 0x89,
	0x25, 0x89, 0xfa, 0x50, 0x3b, 0x93, 0xd8, 0xc0, 0x2a, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x85, 0x36, 0x5e, 0xe9, 0xa3, 0x00, 0x00, 0x00,
}
