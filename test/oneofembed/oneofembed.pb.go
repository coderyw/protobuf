// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oneofembed.proto

package proto

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/coderyw/protobuf/gogoproto"
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

type Foo struct {
	*Bar                 `protobuf:"bytes,1,opt,name=bar,proto3,embedded=bar" json:"bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_70b95aca3b5d76ed, []int{0}
}
func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (m *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(m, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) IsNil() bool {
	return m == nil

}

type Bar struct {
	// Types that are valid to be assigned to Pick:
	//
	//	*Bar_A
	//	*Bar_B
	Pick                 isBar_Pick `protobuf_oneof:"pick"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Bar) Reset()         { *m = Bar{} }
func (m *Bar) String() string { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()    {}
func (*Bar) Descriptor() ([]byte, []int) {
	return fileDescriptor_70b95aca3b5d76ed, []int{1}
}
func (m *Bar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bar.Unmarshal(m, b)
}
func (m *Bar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bar.Marshal(b, m, deterministic)
}
func (m *Bar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bar.Merge(m, src)
}
func (m *Bar) XXX_Size() int {
	return xxx_messageInfo_Bar.Size(m)
}
func (m *Bar) XXX_DiscardUnknown() {
	xxx_messageInfo_Bar.DiscardUnknown(m)
}

var xxx_messageInfo_Bar proto.InternalMessageInfo

type isBar_Pick interface {
	isBar_Pick()
	Equal(interface{}) bool
	Compare(interface{}) int
}

type Bar_A struct {
	A bool `protobuf:"varint,11,opt,name=a,proto3,oneof" json:"a,omitempty"`
}
type Bar_B struct {
	B bool `protobuf:"varint,12,opt,name=b,proto3,oneof" json:"b,omitempty"`
}

func (*Bar_A) isBar_Pick() {}
func (*Bar_B) isBar_Pick() {}

func (m *Bar) GetPick() isBar_Pick {
	if m != nil {
		return m.Pick
	}
	return nil
}

func (m *Bar) GetA() bool {
	if x, ok := m.GetPick().(*Bar_A); ok {
		return x.A
	}
	return false
}

func (m *Bar) GetB() bool {
	if x, ok := m.GetPick().(*Bar_B); ok {
		return x.B
	}
	return false
}

func (m *Bar) IsNil() bool {
	return m == nil

}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Bar) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Bar_A)(nil),
		(*Bar_B)(nil),
	}
}

func init() {
	proto.RegisterType((*Foo)(nil), "proto.Foo")
	proto.RegisterType((*Bar)(nil), "proto.Bar")
}

func init() { proto.RegisterFile("oneofembed.proto", fileDescriptor_70b95aca3b5d76ed) }

var fileDescriptor_70b95aca3b5d76ed = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xcf, 0x4b, 0xcd,
	0x4f, 0x4b, 0xcd, 0x4d, 0x4a, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53,
	0x52, 0x06, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc9, 0xf9, 0x29,
	0xa9, 0x45, 0x95, 0xe5, 0xfa, 0x60, 0x99, 0xa4, 0xd2, 0x34, 0xfd, 0xf4, 0xfc, 0xf4, 0x7c, 0x30,
	0x07, 0xcc, 0x82, 0x68, 0x54, 0xd2, 0xe4, 0x62, 0x76, 0xcb, 0xcf, 0x17, 0x52, 0xe2, 0x62, 0x4e,
	0x4a, 0x2c, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0xe2, 0x82, 0xc8, 0xe9, 0x39, 0x25, 0x16,
	0x39, 0xb1, 0x5c, 0xb8, 0x27, 0xcf, 0x18, 0x04, 0x92, 0x54, 0xd2, 0xe5, 0x62, 0x76, 0x4a, 0x2c,
	0x12, 0xe2, 0xe3, 0x62, 0x4c, 0x94, 0xe0, 0x56, 0x60, 0xd4, 0xe0, 0xf0, 0x60, 0x08, 0x62, 0x4c,
	0x04, 0xf1, 0x93, 0x24, 0x78, 0x60, 0xfc, 0x24, 0x27, 0x36, 0x2e, 0x96, 0x82, 0xcc, 0xe4, 0x6c,
	0x27, 0x81, 0x1f, 0x0f, 0xe5, 0x18, 0x57, 0x3c, 0x92, 0x63, 0xdc, 0xf1, 0x48, 0x8e, 0x71, 0xc5,
	0x63, 0x39, 0xc6, 0x24, 0x36, 0xb0, 0xb1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xd1,
	0x32, 0x3e, 0xbf, 0x00, 0x00, 0x00,
}

func (this *Foo) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*Foo)
	if !ok {
		that2, ok := that.(Foo)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if c := this.Bar.Compare(that1.Bar); c != 0 {
		return c
	}
	if c := bytes.Compare(this.XXX_unrecognized, that1.XXX_unrecognized); c != 0 {
		return c
	}
	return 0
}
func (this *Bar) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*Bar)
	if !ok {
		that2, ok := that.(Bar)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if that1.Pick == nil {
		if this.Pick != nil {
			return 1
		}
	} else if this.Pick == nil {
		return -1
	} else {
		thisType := -1
		switch this.Pick.(type) {
		case *Bar_A:
			thisType = 0
		case *Bar_B:
			thisType = 1
		default:
			panic(fmt.Sprintf("compare: unexpected type %T in oneof", this.Pick))
		}
		that1Type := -1
		switch that1.Pick.(type) {
		case *Bar_A:
			that1Type = 0
		case *Bar_B:
			that1Type = 1
		default:
			panic(fmt.Sprintf("compare: unexpected type %T in oneof", that1.Pick))
		}
		if thisType == that1Type {
			if c := this.Pick.Compare(that1.Pick); c != 0 {
				return c
			}
		} else if thisType < that1Type {
			return -1
		} else if thisType > that1Type {
			return 1
		}
	}
	if c := bytes.Compare(this.XXX_unrecognized, that1.XXX_unrecognized); c != 0 {
		return c
	}
	return 0
}
func (this *Bar_A) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*Bar_A)
	if !ok {
		that2, ok := that.(Bar_A)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.A != that1.A {
		if !this.A {
			return -1
		}
		return 1
	}
	return 0
}
func (this *Bar_B) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*Bar_B)
	if !ok {
		that2, ok := that.(Bar_B)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.B != that1.B {
		if !this.B {
			return -1
		}
		return 1
	}
	return 0
}
func (this *Foo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Foo)
	if !ok {
		that2, ok := that.(Foo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Bar.Equal(that1.Bar) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Bar) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Bar)
	if !ok {
		that2, ok := that.(Bar)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Pick == nil {
		if this.Pick != nil {
			return false
		}
	} else if this.Pick == nil {
		return false
	} else if !this.Pick.Equal(that1.Pick) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Bar_A) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Bar_A)
	if !ok {
		that2, ok := that.(Bar_A)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.A != that1.A {
		return false
	}
	return true
}
func (this *Bar_B) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Bar_B)
	if !ok {
		that2, ok := that.(Bar_B)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.B != that1.B {
		return false
	}
	return true
}
func NewPopulatedFoo(r randyOneofembed, easy bool) *Foo {
	this := &Foo{}
	if r.Intn(5) != 0 {
		this.Bar = NewPopulatedBar(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedOneofembed(r, 2)
	}
	return this
}

func NewPopulatedBar(r randyOneofembed, easy bool) *Bar {
	this := &Bar{}
	oneofNumber_Pick := []int32{11, 12}[r.Intn(2)]
	switch oneofNumber_Pick {
	case 11:
		this.Pick = NewPopulatedBar_A(r, easy)
	case 12:
		this.Pick = NewPopulatedBar_B(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedOneofembed(r, 13)
	}
	return this
}

func NewPopulatedBar_A(r randyOneofembed, easy bool) *Bar_A {
	this := &Bar_A{}
	this.A = bool(bool(r.Intn(2) == 0))
	return this
}
func NewPopulatedBar_B(r randyOneofembed, easy bool) *Bar_B {
	this := &Bar_B{}
	this.B = bool(bool(r.Intn(2) == 0))
	return this
}

type randyOneofembed interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneOneofembed(r randyOneofembed) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringOneofembed(r randyOneofembed) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneOneofembed(r)
	}
	return string(tmps)
}
func randUnrecognizedOneofembed(r randyOneofembed, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldOneofembed(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldOneofembed(dAtA []byte, r randyOneofembed, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateOneofembed(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateOneofembed(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateOneofembed(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateOneofembed(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateOneofembed(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateOneofembed(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateOneofembed(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
