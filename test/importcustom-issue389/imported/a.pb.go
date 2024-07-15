// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: a.proto

package imported

import (
	bytes "bytes"
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

type A struct {
	F1                   string   `protobuf:"bytes,1,opt,name=f1,proto3" json:"f1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *A) Reset()         { *m = A{} }
func (m *A) String() string { return proto.CompactTextString(m) }
func (*A) ProtoMessage()    {}
func (*A) Descriptor() ([]byte, []int) {
	return fileDescriptor_44591ff926c61195, []int{0}
}
func (m *A) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_A.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_A.Merge(m, src)
}
func (m *A) XXX_Size() int {
	return m.Size()
}
func (m *A) XXX_DiscardUnknown() {
	xxx_messageInfo_A.DiscardUnknown(m)
}

var xxx_messageInfo_A proto.InternalMessageInfo

func (m *A) GetF1() string {
	if m != nil {
		return m.F1
	}
	return ""
}

func (m *A) SetF1_(val string) {
	if m != nil {
		m.F1 = val
	}

}

func (m *A) IsNil() bool {
	return m == nil

}

func init() {
	proto.RegisterType((*A)(nil), "imported.A")
}

func init() { proto.RegisterFile("a.proto", fileDescriptor_44591ff926c61195) }

var fileDescriptor_44591ff926c61195 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x62, 0x4f, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0xcc, 0x2d, 0xc8, 0x2f, 0x2a, 0x49, 0x4d, 0x91, 0x32, 0x48,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x4f, 0x49, 0x2d, 0xaa,
	0x2c, 0xd7, 0x07, 0xab, 0x49, 0x2a, 0x4d, 0xd3, 0x4f, 0xcf, 0x4f, 0xcf, 0x07, 0x73, 0xc0, 0x2c,
	0x88, 0x5e, 0x25, 0x61, 0x2e, 0x46, 0x47, 0x21, 0x3e, 0x2e, 0xa6, 0x34, 0x43, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0xa6, 0x34, 0x43, 0x27, 0x89, 0x1f, 0x0f, 0xe5, 0x18, 0x57, 0x3c, 0x92,
	0x63, 0xdc, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x4c, 0x62, 0x03, 0xeb, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x0e, 0x2c,
	0xd9, 0x7c, 0x00, 0x00, 0x00,
}

func (this *A) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*A)
	if !ok {
		that2, ok := that.(A)
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
	if this.F1 != that1.F1 {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *A) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *A) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *A) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.F1) > 0 {
		i -= len(m.F1)
		copy(dAtA[i:], m.F1)
		i = encodeVarintA(dAtA, i, uint64(len(m.F1)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintA(dAtA []byte, offset int, v uint64) int {
	offset -= sovA(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedA(r randyA, easy bool) *A {
	this := &A{}
	this.F1 = string(randStringA(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedA(r, 2)
	}
	return this
}

type randyA interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneA(r randyA) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringA(r randyA) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneA(r)
	}
	return string(tmps)
}
func randUnrecognizedA(r randyA, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldA(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldA(dAtA []byte, r randyA, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateA(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateA(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateA(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateA(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateA(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateA(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateA(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *A) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.F1)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovA(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozA(x uint64) (n int) {
	return sovA(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *A) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowA
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
			return fmt.Errorf("proto: A: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: A: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field F1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthA
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.F1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipA(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthA
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
func skipA(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowA
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
					return 0, ErrIntOverflowA
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
					return 0, ErrIntOverflowA
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
				return 0, ErrInvalidLengthA
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupA
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthA
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthA        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowA          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupA = fmt.Errorf("proto: unexpected end of group")
)
