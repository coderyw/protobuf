// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: extension_extra.proto

package extension_extra

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

type ExtraMessage struct {
	Width                *int32   `protobuf:"varint,1,opt,name=width" json:"width,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtraMessage) Reset()         { *m = ExtraMessage{} }
func (m *ExtraMessage) String() string { return proto.CompactTextString(m) }
func (*ExtraMessage) ProtoMessage()    {}
func (*ExtraMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_66f87514ae0ade66, []int{0}
}
func (m *ExtraMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraMessage.Unmarshal(m, b)
}
func (m *ExtraMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraMessage.Marshal(b, m, deterministic)
}
func (m *ExtraMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraMessage.Merge(m, src)
}
func (m *ExtraMessage) XXX_Size() int {
	return xxx_messageInfo_ExtraMessage.Size(m)
}
func (m *ExtraMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraMessage proto.InternalMessageInfo

func (m *ExtraMessage) GetWidth() int32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *ExtraMessage) SetWidth_(val int32) {
	if m != nil {
		*m.Width = val
	}

}

func (m *ExtraMessage) IsNil() bool {
	return m == nil

}

func init() {
	proto.RegisterType((*ExtraMessage)(nil), "extension_extra.ExtraMessage")
}

func init() { proto.RegisterFile("extension_extra.proto", fileDescriptor_66f87514ae0ade66) }

var fileDescriptor_66f87514ae0ade66 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xad, 0x28, 0x49,
	0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x8b, 0x4f, 0xad, 0x28, 0x29, 0x4a, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x47, 0x13, 0x56, 0x52, 0xe1, 0xe2, 0x71, 0x05, 0x31, 0x7c, 0x53, 0x8b, 0x8b,
	0x13, 0xd3, 0x53, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x33, 0x53, 0x4a, 0x32, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x58, 0x83, 0x20, 0x1c, 0x27, 0xb7, 0x28, 0x97, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd,
	0xe4, 0xfc, 0x5c, 0xfd, 0xe4, 0xfc, 0x94, 0xd4, 0xa2, 0xca, 0x72, 0x7d, 0xb0, 0x91, 0x49, 0xa5,
	0x69, 0x10, 0x46, 0xb2, 0x6e, 0x7a, 0x6a, 0x9e, 0x6e, 0x7a, 0x7e, 0x7a, 0xbe, 0x7e, 0x49, 0x6a,
	0x71, 0x49, 0x4a, 0x62, 0x49, 0xa2, 0x3e, 0x9a, 0x6d, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x44,
	0xc3, 0x93, 0xc9, 0x96, 0x00, 0x00, 0x00,
}
