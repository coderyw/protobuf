// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: a.proto

package import_public

import (
	fmt "fmt"
	proto "github.com/coderyw/protobuf/proto"
	sub "github.com/coderyw/protobuf/protoc-gen-gogo/testdata/import_public/sub"
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

const Default_M_DefaultField = sub.Default_M_DefaultField

// M from public import import_public/sub/a.proto
type M = sub.M
type M_OneofInt32 = sub.M_OneofInt32
type M_OneofInt64 = sub.M_OneofInt64

// M_Grouping from public import import_public/sub/a.proto
type M_Grouping = sub.M_Grouping

// M_Submessage from public import import_public/sub/a.proto
type M_Submessage = sub.M_Submessage
type M_Submessage_SubmessageOneofInt32 = sub.M_Submessage_SubmessageOneofInt32
type M_Submessage_SubmessageOneofInt64 = sub.M_Submessage_SubmessageOneofInt64

// E from public import import_public/sub/a.proto
type E = sub.E

var E_name = sub.E_name
var E_value = sub.E_value

const E_ZERO = E(sub.E_ZERO)

// M_Subenum from public import import_public/sub/a.proto
type M_Subenum = sub.M_Subenum

var M_Subenum_name = sub.M_Subenum_name
var M_Subenum_value = sub.M_Subenum_value

const M_M_ZERO = M_Subenum(sub.M_M_ZERO)

// M_Submessage_Submessage_Subenum from public import import_public/sub/a.proto
type M_Submessage_Submessage_Subenum = sub.M_Submessage_Submessage_Subenum

var M_Submessage_Submessage_Subenum_name = sub.M_Submessage_Submessage_Subenum_name
var M_Submessage_Submessage_Subenum_value = sub.M_Submessage_Submessage_Subenum_value

const M_Submessage_M_SUBMESSAGE_ZERO = M_Submessage_Submessage_Subenum(sub.M_Submessage_M_SUBMESSAGE_ZERO)

var E_ExtensionField = sub.E_ExtensionField

type Public struct {
	M                    *sub.M   `protobuf:"bytes,1,opt,name=m" json:"m,omitempty"`
	E                    *sub.E   `protobuf:"varint,2,opt,name=e,enum=goproto.test.import_public.sub.E" json:"e,omitempty"`
	Local                *Local   `protobuf:"bytes,3,opt,name=local" json:"local,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Public) Reset()         { *m = Public{} }
func (m *Public) String() string { return proto.CompactTextString(m) }
func (*Public) ProtoMessage()    {}
func (*Public) Descriptor() ([]byte, []int) {
	return fileDescriptor_44591ff926c61195, []int{0}
}
func (m *Public) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Public.Unmarshal(m, b)
}
func (m *Public) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Public.Marshal(b, m, deterministic)
}
func (m *Public) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Public.Merge(m, src)
}
func (m *Public) XXX_Size() int {
	return xxx_messageInfo_Public.Size(m)
}
func (m *Public) XXX_DiscardUnknown() {
	xxx_messageInfo_Public.DiscardUnknown(m)
}

var xxx_messageInfo_Public proto.InternalMessageInfo

func (m *Public) GetM() *sub.M {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *Public) GetE() sub.E {
	if m != nil && m.E != nil {
		return *m.E
	}
	return sub.E_ZERO
}

func (m *Public) GetLocal() *Local {
	if m != nil {
		return m.Local
	}
	return nil
}

func (m *Public) SetM_(val *sub.M) {
	if m != nil {
		m.M = val
	}
	return
}

func (m *Public) SetE_(val sub.E) {
	if m != nil {
		*m.E = val
	}

}

func (m *Public) SetLocal_(val *Local) {
	if m != nil {
		m.Local = val
	}
	return
}

func (m *Public) IsNil() bool {
	return m == nil

}

func init() {
	proto.RegisterType((*Public)(nil), "goproto.test.import_public.Public")
}

func init() { proto.RegisterFile("a.proto", fileDescriptor_44591ff926c61195) }

var fileDescriptor_44591ff926c61195 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x62, 0x4f, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4a, 0xcf, 0x07, 0x33, 0xf4, 0x4a, 0x52, 0x8b, 0x4b, 0xf4, 0x32,
	0x73, 0x0b, 0xf2, 0x8b, 0x4a, 0xe2, 0x0b, 0x4a, 0x93, 0x72, 0x32, 0x93, 0xa5, 0x24, 0x51, 0xb8,
	0xfa, 0xc5, 0xa5, 0x49, 0xfa, 0x50, 0x6d, 0x52, 0xa2, 0xa8, 0x52, 0x49, 0x10, 0x61, 0xa5, 0xc5,
	0x8c, 0x5c, 0x6c, 0x01, 0x60, 0x21, 0x21, 0x7d, 0x2e, 0xc6, 0x5c, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0x6e, 0x23, 0x45, 0x3d, 0xdc, 0x96, 0xe8, 0x15, 0x97, 0x26, 0xe9, 0xf9, 0x06, 0x31, 0xe6, 0x82,
	0x34, 0xa4, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x11, 0xd6, 0xe0, 0x1a, 0xc4, 0x98, 0x2a, 0x64,
	0xce, 0xc5, 0x9a, 0x93, 0x9f, 0x9c, 0x98, 0x23, 0xc1, 0x4c, 0xd8, 0x16, 0x1f, 0x90, 0xc2, 0x20,
	0x88, 0x7a, 0x27, 0x97, 0x28, 0xa7, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c,
	0xfd, 0xe4, 0xfc, 0x94, 0xd4, 0xa2, 0xca, 0x72, 0x7d, 0xb0, 0xde, 0xa4, 0xd2, 0x34, 0x08, 0x23,
	0x59, 0x37, 0x3d, 0x35, 0x4f, 0x37, 0x3d, 0x3f, 0x3d, 0x5f, 0x1f, 0x64, 0x5c, 0x4a, 0x62, 0x49,
	0xa2, 0x3e, 0x8a, 0x91, 0x01, 0x0c, 0x01, 0x8c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x08,
	0x5a, 0x22, 0x47, 0x01, 0x00, 0x00,
}
