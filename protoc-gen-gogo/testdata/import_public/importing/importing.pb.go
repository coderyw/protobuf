// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: import_public/importing/importing.proto

package importing

import (
	fmt "fmt"
	proto "github.com/coderyw/protobuf/proto"
	_ "github.com/coderyw/protobuf/protoc-gen-gogo/testdata/import_public"
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

type M struct {
	// Message type defined in a file publicly imported by a file we import.
	M                    *sub.M   `protobuf:"bytes,1,opt,name=m" json:"m,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M) Reset()         { *m = M{} }
func (m *M) String() string { return proto.CompactTextString(m) }
func (*M) ProtoMessage()    {}
func (*M) Descriptor() ([]byte, []int) {
	return fileDescriptor_36b835b3b8f6171a, []int{0}
}
func (m *M) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M.Unmarshal(m, b)
}
func (m *M) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M.Marshal(b, m, deterministic)
}
func (m *M) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M.Merge(m, src)
}
func (m *M) XXX_Size() int {
	return xxx_messageInfo_M.Size(m)
}
func (m *M) XXX_DiscardUnknown() {
	xxx_messageInfo_M.DiscardUnknown(m)
}

var xxx_messageInfo_M proto.InternalMessageInfo

func (m *M) GetM() *sub.M {
	if m != nil {
		return m.M
	}
	return nil
}

func init() {
	proto.RegisterType((*M)(nil), "goproto.test.import_public.importing.M")
}

func init() {
	proto.RegisterFile("import_public/importing/importing.proto", fileDescriptor_36b835b3b8f6171a)
}

var fileDescriptor_36b835b3b8f6171a = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcf, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x89, 0x2f, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0xd6, 0x87, 0xf0, 0x32, 0xf3, 0xd2, 0x11,
	0x2c, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x95, 0xf4, 0x7c, 0x30, 0x43, 0xaf, 0x24, 0xb5,
	0xb8, 0x44, 0x0f, 0x45, 0x97, 0x1e, 0x5c, 0xad, 0x94, 0x28, 0xaa, 0x71, 0x89, 0x10, 0xcd, 0x4a,
	0x26, 0x5c, 0x8c, 0xbe, 0x42, 0xfa, 0x5c, 0x8c, 0xb9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46,
	0x8a, 0x7a, 0x78, 0x4c, 0x2b, 0x2e, 0x4d, 0xd2, 0xf3, 0x0d, 0x62, 0xcc, 0x75, 0xf2, 0x8e, 0xf2,
	0x4c, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x49, 0xcc,
	0x4b, 0xd7, 0x07, 0x6b, 0x4b, 0x2a, 0x4d, 0x83, 0x30, 0x92, 0x75, 0xd3, 0x53, 0xf3, 0x74, 0xd3,
	0xf3, 0xf5, 0x41, 0xe6, 0xa4, 0x24, 0x96, 0x24, 0xea, 0xe3, 0xf0, 0x0f, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xd8, 0x7e, 0x58, 0x1c, 0xe9, 0x00, 0x00, 0x00,
}
