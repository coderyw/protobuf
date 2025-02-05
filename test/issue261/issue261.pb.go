// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue261.proto

package issue261

import (
	fmt "fmt"
	_ "github.com/coderyw/protobuf/gogoproto"
	proto "github.com/coderyw/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MapStdTypes struct {
	NullableDuration     map[int32]*time.Duration `protobuf:"bytes,3,rep,name=nullableDuration,proto3,stdduration" json:"nullableDuration,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MapStdTypes) Reset()         { *m = MapStdTypes{} }
func (m *MapStdTypes) String() string { return proto.CompactTextString(m) }
func (*MapStdTypes) ProtoMessage()    {}
func (*MapStdTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_35fa31fa2023935a, []int{0}
}
func (m *MapStdTypes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapStdTypes.Unmarshal(m, b)
}
func (m *MapStdTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapStdTypes.Marshal(b, m, deterministic)
}
func (m *MapStdTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapStdTypes.Merge(m, src)
}
func (m *MapStdTypes) XXX_Size() int {
	return xxx_messageInfo_MapStdTypes.Size(m)
}
func (m *MapStdTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_MapStdTypes.DiscardUnknown(m)
}

var xxx_messageInfo_MapStdTypes proto.InternalMessageInfo

func (m *MapStdTypes) GetNullableDuration() map[int32]*time.Duration {
	if m != nil {
		return m.NullableDuration
	}
	return nil
}

func (m *MapStdTypes) SetNullableDuration_(val map[int32]*time.Duration) {
	if m != nil {
		m.NullableDuration = val
	}
	return
}

func (m *MapStdTypes) IsNil() bool {
	return m == nil

}

func init() {
	proto.RegisterType((*MapStdTypes)(nil), "issue261.MapStdTypes")
	proto.RegisterMapType((map[int32]*time.Duration)(nil), "issue261.MapStdTypes.NullableDurationEntry")
}

func init() { proto.RegisterFile("issue261.proto", fileDescriptor_35fa31fa2023935a) }

var fileDescriptor_35fa31fa2023935a = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x32, 0x33, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x0c,
	0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0xf3, 0x53, 0x52, 0x8b,
	0x2a, 0xcb, 0xf5, 0xc1, 0x6a, 0x92, 0x4a, 0xd3, 0xf4, 0xd3, 0xf3, 0xd3, 0xf3, 0xc1, 0x1c, 0x30,
	0x0b, 0xa2, 0x57, 0x4a, 0x2e, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x15, 0xa1, 0x2a, 0xa5, 0xb4, 0x28,
	0xb1, 0x24, 0x33, 0x3f, 0x0f, 0x22, 0xaf, 0x74, 0x86, 0x91, 0x8b, 0xdb, 0x37, 0xb1, 0x20, 0xb8,
	0x24, 0x25, 0xa4, 0xb2, 0x20, 0xb5, 0x58, 0x28, 0x96, 0x4b, 0x20, 0xaf, 0x34, 0x27, 0x27, 0x31,
	0x29, 0x27, 0xd5, 0x05, 0xaa, 0x52, 0x82, 0x59, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x5b, 0x0f, 0xee,
	0x2c, 0x24, 0x0d, 0x7a, 0x7e, 0x68, 0xaa, 0x5d, 0xf3, 0x4a, 0x8a, 0x2a, 0x9d, 0x58, 0x66, 0xdc,
	0x97, 0x67, 0x0c, 0xc2, 0x30, 0x4a, 0x2a, 0x8e, 0x4b, 0x14, 0xab, 0x06, 0x21, 0x01, 0x2e, 0xe6,
	0xec, 0xd4, 0x4a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x10, 0x53, 0x48, 0x9f, 0x8b, 0xb5,
	0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x52, 0x0f, 0xe2, 0x13,
	0x3d, 0x98, 0x4f, 0xf4, 0x60, 0x06, 0x04, 0x41, 0xd4, 0x59, 0x31, 0x59, 0x30, 0x26, 0xb1, 0x81,
	0x65, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xee, 0x58, 0xb1, 0xad, 0x43, 0x01, 0x00, 0x00,
}
