// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: m2.proto

package beta

import (
	fmt "fmt"
	proto "github.com/coderyw/protobuf/proto"
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

type M2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2) Reset()         { *m = M2{} }
func (m *M2) String() string { return proto.CompactTextString(m) }
func (*M2) ProtoMessage()    {}
func (*M2) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eacb2783a20cf9d, []int{0}
}
func (m *M2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2.Unmarshal(m, b)
}
func (m *M2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2.Marshal(b, m, deterministic)
}
func (m *M2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2.Merge(m, src)
}
func (m *M2) XXX_Size() int {
	return xxx_messageInfo_M2.Size(m)
}
func (m *M2) XXX_DiscardUnknown() {
	xxx_messageInfo_M2.DiscardUnknown(m)
}

var xxx_messageInfo_M2 proto.InternalMessageInfo

func (m *M2) IsNil() bool {
	return m == nil

}

func init() {
	proto.RegisterType((*M2)(nil), "test.b.part2.M2")
}

func init() { proto.RegisterFile("m2.proto", fileDescriptor_8eacb2783a20cf9d) }

var fileDescriptor_8eacb2783a20cf9d = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x35, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x29, 0x49, 0x2d, 0x2e, 0xd1, 0x4b, 0xd2, 0x2b, 0x48, 0x2c,
	0x2a, 0x31, 0x52, 0x62, 0xe1, 0x62, 0xf2, 0x35, 0x72, 0xf2, 0x89, 0xf2, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x4f, 0x49, 0x2d, 0xaa, 0x2c, 0xd7, 0x07,
	0xab, 0x4f, 0x2a, 0x4d, 0x83, 0x30, 0x92, 0x75, 0xd3, 0x53, 0xf3, 0x74, 0xd3, 0xf3, 0xd3, 0xf3,
	0xf5, 0x41, 0x26, 0xa4, 0x24, 0x96, 0x24, 0xea, 0x67, 0xe6, 0x16, 0xe4, 0x17, 0x95, 0x14, 0x83,
	0x05, 0xe2, 0x93, 0xe2, 0x0d, 0xad, 0x93, 0x52, 0x4b, 0x12, 0x93, 0xd8, 0xc0, 0xea, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x94, 0xf0, 0x33, 0x74, 0x00, 0x00, 0x00,
}
