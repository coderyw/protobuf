// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue42.proto

package issue42

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/coderyw/protobuf/gogoproto"
	proto "github.com/coderyw/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type UnorderedFields struct {
	A                    *int64   `protobuf:"varint,10,opt,name=A" json:"A,omitempty"`
	B                    *uint64  `protobuf:"fixed64,1,opt,name=B" json:"B,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnorderedFields) Reset()         { *m = UnorderedFields{} }
func (m *UnorderedFields) String() string { return proto.CompactTextString(m) }
func (*UnorderedFields) ProtoMessage()    {}
func (*UnorderedFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4aafed97be2033, []int{0}
}
func (m *UnorderedFields) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnorderedFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnorderedFields.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnorderedFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnorderedFields.Merge(m, src)
}
func (m *UnorderedFields) XXX_Size() int {
	return m.Size()
}
func (m *UnorderedFields) XXX_DiscardUnknown() {
	xxx_messageInfo_UnorderedFields.DiscardUnknown(m)
}

var xxx_messageInfo_UnorderedFields proto.InternalMessageInfo

func (m *UnorderedFields) GetA() int64 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

func (m *UnorderedFields) GetB() uint64 {
	if m != nil && m.B != nil {
		return *m.B
	}
	return 0
}

func (m *UnorderedFields) SetA_(val int64) {
	if m != nil {
		*m.A = val
	}

}

func (m *UnorderedFields) SetB_(val uint64) {
	if m != nil {
		*m.B = val
	}

}

func (m *UnorderedFields) IsNil() bool {
	return m == nil

}

type OrderedFields struct {
	B                    *uint64  `protobuf:"fixed64,1,opt,name=B" json:"B,omitempty"`
	A                    *int64   `protobuf:"varint,10,opt,name=A" json:"A,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderedFields) Reset()         { *m = OrderedFields{} }
func (m *OrderedFields) String() string { return proto.CompactTextString(m) }
func (*OrderedFields) ProtoMessage()    {}
func (*OrderedFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4aafed97be2033, []int{1}
}
func (m *OrderedFields) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderedFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderedFields.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderedFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderedFields.Merge(m, src)
}
func (m *OrderedFields) XXX_Size() int {
	return m.Size()
}
func (m *OrderedFields) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderedFields.DiscardUnknown(m)
}

var xxx_messageInfo_OrderedFields proto.InternalMessageInfo

func (m *OrderedFields) GetB() uint64 {
	if m != nil && m.B != nil {
		return *m.B
	}
	return 0
}

func (m *OrderedFields) GetA() int64 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

func (m *OrderedFields) SetB_(val uint64) {
	if m != nil {
		*m.B = val
	}

}

func (m *OrderedFields) SetA_(val int64) {
	if m != nil {
		*m.A = val
	}

}

func (m *OrderedFields) IsNil() bool {
	return m == nil

}

func init() {
	proto.RegisterType((*UnorderedFields)(nil), "issue42.UnorderedFields")
	proto.RegisterType((*OrderedFields)(nil), "issue42.OrderedFields")
}

func init() { proto.RegisterFile("issue42.proto", fileDescriptor_fb4aafed97be2033) }

var fileDescriptor_fb4aafed97be2033 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x31, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0x0c, 0xd2,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0xf3, 0x53, 0x52, 0x8b, 0x2a,
	0xcb, 0xf5, 0xc1, 0x4a, 0x92, 0x4a, 0xd3, 0xf4, 0xd3, 0xf3, 0xd3, 0xf3, 0xc1, 0x1c, 0x30, 0x0b,
	0xa2, 0x55, 0x49, 0x97, 0x8b, 0x3f, 0x34, 0x2f, 0xbf, 0x28, 0x25, 0xb5, 0x28, 0x35, 0xc5, 0x2d,
	0x33, 0x35, 0x27, 0xa5, 0x58, 0x88, 0x87, 0x8b, 0xd1, 0x51, 0x82, 0x4b, 0x81, 0x51, 0x83, 0x39,
	0x88, 0xd1, 0x11, 0xc4, 0x73, 0x92, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0b, 0x62, 0x74, 0x52, 0xd2,
	0xe6, 0xe2, 0xf5, 0x47, 0x57, 0x8c, 0x90, 0x46, 0xd5, 0xea, 0x24, 0xf0, 0xe3, 0xa1, 0x1c, 0xe3,
	0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x08, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xb4, 0xe7, 0x8a, 0x00, 0xb8, 0x00, 0x00, 0x00,
}

func (m *UnorderedFields) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnorderedFields) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnorderedFields) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.A != nil {
		i = encodeVarintIssue42(dAtA, i, uint64(*m.A))
		i--
		dAtA[i] = 0x50
	}
	if m.B != nil {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(*m.B))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func (m *OrderedFields) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderedFields) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderedFields) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.A != nil {
		i = encodeVarintIssue42(dAtA, i, uint64(*m.A))
		i--
		dAtA[i] = 0x50
	}
	if m.B != nil {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(*m.B))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func encodeVarintIssue42(dAtA []byte, offset int, v uint64) int {
	offset -= sovIssue42(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedUnorderedFields(r randyIssue42, easy bool) *UnorderedFields {
	this := &UnorderedFields{}
	if r.Intn(5) != 0 {
		v1 := uint64(uint64(r.Uint32()))
		this.B = &v1
	}
	if r.Intn(5) != 0 {
		v2 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		this.A = &v2
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedIssue42(r, 11)
	}
	return this
}

func NewPopulatedOrderedFields(r randyIssue42, easy bool) *OrderedFields {
	this := &OrderedFields{}
	if r.Intn(5) != 0 {
		v3 := uint64(uint64(r.Uint32()))
		this.B = &v3
	}
	if r.Intn(5) != 0 {
		v4 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		this.A = &v4
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedIssue42(r, 11)
	}
	return this
}

type randyIssue42 interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneIssue42(r randyIssue42) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringIssue42(r randyIssue42) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneIssue42(r)
	}
	return string(tmps)
}
func randUnrecognizedIssue42(r randyIssue42, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldIssue42(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldIssue42(dAtA []byte, r randyIssue42, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateIssue42(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateIssue42(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateIssue42(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateIssue42(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateIssue42(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateIssue42(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateIssue42(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *UnorderedFields) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.B != nil {
		n += 9
	}
	if m.A != nil {
		n += 1 + sovIssue42(uint64(*m.A))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *OrderedFields) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.B != nil {
		n += 9
	}
	if m.A != nil {
		n += 1 + sovIssue42(uint64(*m.A))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIssue42(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIssue42(x uint64) (n int) {
	return sovIssue42(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UnorderedFields) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssue42
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UnorderedFields: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnorderedFields: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.B = &v
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue42
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.A = &v
		default:
			iNdEx = preIndex
			skippy, err := skipIssue42(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIssue42
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OrderedFields) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssue42
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OrderedFields: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderedFields: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.B = &v
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue42
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.A = &v
		default:
			iNdEx = preIndex
			skippy, err := skipIssue42(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIssue42
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipIssue42(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIssue42
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIssue42
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIssue42
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthIssue42
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIssue42
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIssue42
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIssue42        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIssue42          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIssue42 = fmt.Errorf("proto: unexpected end of group")
)
