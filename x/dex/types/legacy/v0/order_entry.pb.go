// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: legacy/dex/v0/order_entry.proto

package v0

import (
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type OrderEntry struct {
	Price             uint64   `protobuf:"varint,1,opt,name=price,proto3" json:"price,omitempty"`
	Quantity          uint64   `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	AllocationCreator []string `protobuf:"bytes,3,rep,name=allocationCreator,proto3" json:"allocationCreator,omitempty"`
	Allocation        []uint64 `protobuf:"varint,4,rep,packed,name=allocation,proto3" json:"allocation,omitempty"`
	PriceDenom        string   `protobuf:"bytes,5,opt,name=priceDenom,proto3" json:"priceDenom,omitempty"`
	AssetDenom        string   `protobuf:"bytes,6,opt,name=assetDenom,proto3" json:"assetDenom,omitempty"`
}

func (m *OrderEntry) Reset()         { *m = OrderEntry{} }
func (m *OrderEntry) String() string { return proto.CompactTextString(m) }
func (*OrderEntry) ProtoMessage()    {}
func (*OrderEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_742efccd6b0db773, []int{0}
}
func (m *OrderEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderEntry.Merge(m, src)
}
func (m *OrderEntry) XXX_Size() int {
	return m.Size()
}
func (m *OrderEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderEntry.DiscardUnknown(m)
}

var xxx_messageInfo_OrderEntry proto.InternalMessageInfo

func (m *OrderEntry) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderEntry) GetQuantity() uint64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *OrderEntry) GetAllocationCreator() []string {
	if m != nil {
		return m.AllocationCreator
	}
	return nil
}

func (m *OrderEntry) GetAllocation() []uint64 {
	if m != nil {
		return m.Allocation
	}
	return nil
}

func (m *OrderEntry) GetPriceDenom() string {
	if m != nil {
		return m.PriceDenom
	}
	return ""
}

func (m *OrderEntry) GetAssetDenom() string {
	if m != nil {
		return m.AssetDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*OrderEntry)(nil), "seiprotocol.seichain.legacy.dex.v0.OrderEntry")
}

func init() { proto.RegisterFile("legacy/dex/v0/order_entry.proto", fileDescriptor_742efccd6b0db773) }

var fileDescriptor_742efccd6b0db773 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xbd, 0x4e, 0xf3, 0x30,
	0x18, 0x85, 0xe3, 0x2f, 0x69, 0xf5, 0xd5, 0x1b, 0x11, 0x43, 0xc4, 0x60, 0xa2, 0x4e, 0x19, 0xc0,
	0x8e, 0xc4, 0xc0, 0xce, 0xcf, 0x8c, 0xc8, 0xc8, 0x82, 0x5c, 0xe7, 0x55, 0x6b, 0x29, 0xb5, 0x83,
	0xed, 0x46, 0xc9, 0x5d, 0x70, 0x59, 0x0c, 0x0c, 0x1d, 0x19, 0x51, 0x72, 0x23, 0x28, 0x0e, 0xd0,
	0x4a, 0x6c, 0x3e, 0xe7, 0xf1, 0xd1, 0x2b, 0x3d, 0xf8, 0xbc, 0x82, 0x35, 0x17, 0x1d, 0x2b, 0xa1,
	0x65, 0x4d, 0xce, 0xb4, 0x29, 0xc1, 0x3c, 0x83, 0x72, 0xa6, 0xa3, 0xb5, 0xd1, 0x4e, 0xc7, 0x4b,
	0x0b, 0xd2, 0xbf, 0x84, 0xae, 0xa8, 0x05, 0x29, 0x36, 0x5c, 0x2a, 0x3a, 0xad, 0x68, 0x09, 0x2d,
	0x6d, 0xf2, 0xe5, 0x3b, 0xc2, 0xf8, 0x61, 0x5c, 0xde, 0x8f, 0xc3, 0xf8, 0x14, 0xcf, 0x6a, 0x23,
	0x05, 0x24, 0x28, 0x45, 0x59, 0x54, 0x4c, 0x21, 0x3e, 0xc3, 0xff, 0x5f, 0x76, 0x5c, 0x39, 0xe9,
	0xba, 0xe4, 0x9f, 0x07, 0xbf, 0x39, 0xbe, 0xc0, 0x27, 0xbc, 0xaa, 0xb4, 0xe0, 0x4e, 0x6a, 0x75,
	0x6b, 0x80, 0x3b, 0x6d, 0x92, 0x30, 0x0d, 0xb3, 0x45, 0xf1, 0x17, 0xc4, 0x04, 0xe3, 0x43, 0x99,
	0x44, 0x69, 0x98, 0x45, 0xc5, 0x51, 0x33, 0x72, 0x7f, 0xf2, 0x0e, 0x94, 0xde, 0x26, 0xb3, 0x14,
	0x65, 0x8b, 0xe2, 0xa8, 0xf1, 0x7b, 0x6b, 0xc1, 0x4d, 0x7c, 0x3e, 0xf1, 0x43, 0x73, 0xf3, 0xf8,
	0xd6, 0x13, 0xb4, 0xef, 0x09, 0xfa, 0xec, 0x09, 0x7a, 0x1d, 0x48, 0xb0, 0x1f, 0x48, 0xf0, 0x31,
	0x90, 0xe0, 0xe9, 0x7a, 0x2d, 0xdd, 0x66, 0xb7, 0xa2, 0x42, 0x6f, 0x99, 0x05, 0x79, 0xf9, 0x23,
	0xc6, 0x07, 0x6f, 0x86, 0xb5, 0x5e, 0xa5, 0xeb, 0x6a, 0xb0, 0xec, 0xdb, 0x6d, 0x93, 0xaf, 0xe6,
	0xfe, 0xe7, 0xd5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0x72, 0x94, 0x8c, 0x6f, 0x01, 0x00,
	0x00,
}

func (m *OrderEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AssetDenom) > 0 {
		i -= len(m.AssetDenom)
		copy(dAtA[i:], m.AssetDenom)
		i = encodeVarintOrderEntry(dAtA, i, uint64(len(m.AssetDenom)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.PriceDenom) > 0 {
		i -= len(m.PriceDenom)
		copy(dAtA[i:], m.PriceDenom)
		i = encodeVarintOrderEntry(dAtA, i, uint64(len(m.PriceDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Allocation) > 0 {
		dAtA2 := make([]byte, len(m.Allocation)*10)
		var j1 int
		for _, num := range m.Allocation {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintOrderEntry(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AllocationCreator) > 0 {
		for iNdEx := len(m.AllocationCreator) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllocationCreator[iNdEx])
			copy(dAtA[i:], m.AllocationCreator[iNdEx])
			i = encodeVarintOrderEntry(dAtA, i, uint64(len(m.AllocationCreator[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Quantity != 0 {
		i = encodeVarintOrderEntry(dAtA, i, uint64(m.Quantity))
		i--
		dAtA[i] = 0x10
	}
	if m.Price != 0 {
		i = encodeVarintOrderEntry(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrderEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrderEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrderEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Price != 0 {
		n += 1 + sovOrderEntry(uint64(m.Price))
	}
	if m.Quantity != 0 {
		n += 1 + sovOrderEntry(uint64(m.Quantity))
	}
	if len(m.AllocationCreator) > 0 {
		for _, s := range m.AllocationCreator {
			l = len(s)
			n += 1 + l + sovOrderEntry(uint64(l))
		}
	}
	if len(m.Allocation) > 0 {
		l = 0
		for _, e := range m.Allocation {
			l += sovOrderEntry(uint64(e))
		}
		n += 1 + sovOrderEntry(uint64(l)) + l
	}
	l = len(m.PriceDenom)
	if l > 0 {
		n += 1 + l + sovOrderEntry(uint64(l))
	}
	l = len(m.AssetDenom)
	if l > 0 {
		n += 1 + l + sovOrderEntry(uint64(l))
	}
	return n
}

func sovOrderEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrderEntry(x uint64) (n int) {
	return sovOrderEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrderEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrderEntry
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
			return fmt.Errorf("proto: OrderEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			m.Quantity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Quantity |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocationCreator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderEntry
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
				return ErrInvalidLengthOrderEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllocationCreator = append(m.AllocationCreator, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOrderEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Allocation = append(m.Allocation, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOrderEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthOrderEntry
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthOrderEntry
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Allocation) == 0 {
					m.Allocation = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowOrderEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Allocation = append(m.Allocation, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocation", wireType)
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderEntry
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
				return ErrInvalidLengthOrderEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderEntry
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
				return ErrInvalidLengthOrderEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrderEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrderEntry
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
func skipOrderEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrderEntry
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
					return 0, ErrIntOverflowOrderEntry
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
					return 0, ErrIntOverflowOrderEntry
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
				return 0, ErrInvalidLengthOrderEntry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrderEntry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrderEntry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrderEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrderEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrderEntry = fmt.Errorf("proto: unexpected end of group")
)
