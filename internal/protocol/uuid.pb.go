// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protocol/uuid.proto

package protocol

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type UUID struct {
	Low  uint64 `protobuf:"fixed64,1,opt,name=low,proto3" json:"low,omitempty"`
	High uint64 `protobuf:"fixed64,2,opt,name=high,proto3" json:"high,omitempty"`
}

func (m *UUID) Reset()         { *m = UUID{} }
func (m *UUID) String() string { return proto.CompactTextString(m) }
func (*UUID) ProtoMessage()    {}
func (*UUID) Descriptor() ([]byte, []int) {
	return fileDescriptor_70cf74d951327a30, []int{0}
}
func (m *UUID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UUID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UUID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UUID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UUID.Merge(m, src)
}
func (m *UUID) XXX_Size() int {
	return m.Size()
}
func (m *UUID) XXX_DiscardUnknown() {
	xxx_messageInfo_UUID.DiscardUnknown(m)
}

var xxx_messageInfo_UUID proto.InternalMessageInfo

func (m *UUID) GetLow() uint64 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *UUID) GetHigh() uint64 {
	if m != nil {
		return m.High
	}
	return 0
}

func init() {
	proto.RegisterType((*UUID)(nil), "pbrpc.protocol.UUID")
}

func init() { proto.RegisterFile("protocol/uuid.proto", fileDescriptor_70cf74d951327a30) }

var fileDescriptor_70cf74d951327a30 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x2d, 0xcd, 0x4c, 0xd1, 0x03, 0xf3, 0x84, 0xf8, 0x0a, 0x92,
	0x8a, 0x0a, 0x92, 0xf5, 0x60, 0x52, 0x4a, 0x3a, 0x5c, 0x2c, 0xa1, 0xa1, 0x9e, 0x2e, 0x42, 0x02,
	0x5c, 0xcc, 0x39, 0xf9, 0xe5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x6c, 0x41, 0x20, 0xa6, 0x90, 0x10,
	0x17, 0x4b, 0x46, 0x66, 0x7a, 0x86, 0x04, 0x13, 0x58, 0x08, 0xcc, 0x76, 0x72, 0x3d, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0,
	0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xed, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd,
	0xe4, 0xfc, 0x5c, 0xfd, 0x9c, 0xd4, 0x12, 0xdd, 0x2a, 0xdd, 0xf4, 0x7c, 0x7d, 0xb0, 0x5d, 0xfa,
	0x99, 0x79, 0x25, 0xa9, 0x45, 0x79, 0x89, 0x39, 0xfa, 0x30, 0x4b, 0x93, 0xd8, 0xc0, 0x2c, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0x84, 0xc7, 0x76, 0xa2, 0x00, 0x00, 0x00,
}

func (m *UUID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UUID) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Low != 0 {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Low))
		i += 8
	}
	if m.High != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.High))
		i += 8
	}
	return i, nil
}

func encodeVarintUuid(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UUID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Low != 0 {
		n += 9
	}
	if m.High != 0 {
		n += 9
	}
	return n
}

func sovUuid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUuid(x uint64) (n int) {
	return sovUuid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UUID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUuid
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
			return fmt.Errorf("proto: UUID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UUID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Low", wireType)
			}
			m.Low = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Low = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field High", wireType)
			}
			m.High = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.High = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		default:
			iNdEx = preIndex
			skippy, err := skipUuid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUuid
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUuid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUuid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUuid
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
					return 0, ErrIntOverflowUuid
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUuid
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
				return 0, ErrInvalidLengthUuid
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthUuid
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUuid
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipUuid(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthUuid
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthUuid = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUuid   = fmt.Errorf("proto: integer overflow")
)