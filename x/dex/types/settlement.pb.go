// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/settlement.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type SettlementEntry struct {
	Account                string                                 `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	PriceDenom             string                                 `protobuf:"bytes,2,opt,name=priceDenom,proto3" json:"price_denom"`
	AssetDenom             string                                 `protobuf:"bytes,3,opt,name=assetDenom,proto3" json:"asset_denom"`
	Quantity               github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=quantity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"quantity" yaml:"quantity"`
	ExecutionCostOrProceed github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=executionCostOrProceed,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"execution_cost_or_proceed" yaml:"execution_cost_or_proceed"`
	ExpectedCostOrProceed  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=expectedCostOrProceed,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"expected_cost_or_proceed" yaml:"expected_cost_or_proceed"`
	PositionDirection      string                                 `protobuf:"bytes,7,opt,name=positionDirection,proto3" json:"position_direction"`
	PositionEffect         string                                 `protobuf:"bytes,8,opt,name=positionEffect,proto3" json:"position_effect"`
	Leverage               github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=leverage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"leverage" yaml:"leverage"`
}

func (m *SettlementEntry) Reset()         { *m = SettlementEntry{} }
func (m *SettlementEntry) String() string { return proto.CompactTextString(m) }
func (*SettlementEntry) ProtoMessage()    {}
func (*SettlementEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_c24d83c09612bb1c, []int{0}
}
func (m *SettlementEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SettlementEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SettlementEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SettlementEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettlementEntry.Merge(m, src)
}
func (m *SettlementEntry) XXX_Size() int {
	return m.Size()
}
func (m *SettlementEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SettlementEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SettlementEntry proto.InternalMessageInfo

func (m *SettlementEntry) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *SettlementEntry) GetPriceDenom() string {
	if m != nil {
		return m.PriceDenom
	}
	return ""
}

func (m *SettlementEntry) GetAssetDenom() string {
	if m != nil {
		return m.AssetDenom
	}
	return ""
}

func (m *SettlementEntry) GetPositionDirection() string {
	if m != nil {
		return m.PositionDirection
	}
	return ""
}

func (m *SettlementEntry) GetPositionEffect() string {
	if m != nil {
		return m.PositionEffect
	}
	return ""
}

type Settlements struct {
	Epoch   int64              `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch"`
	Entries []*SettlementEntry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries"`
}

func (m *Settlements) Reset()         { *m = Settlements{} }
func (m *Settlements) String() string { return proto.CompactTextString(m) }
func (*Settlements) ProtoMessage()    {}
func (*Settlements) Descriptor() ([]byte, []int) {
	return fileDescriptor_c24d83c09612bb1c, []int{1}
}
func (m *Settlements) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Settlements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Settlements.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Settlements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settlements.Merge(m, src)
}
func (m *Settlements) XXX_Size() int {
	return m.Size()
}
func (m *Settlements) XXX_DiscardUnknown() {
	xxx_messageInfo_Settlements.DiscardUnknown(m)
}

var xxx_messageInfo_Settlements proto.InternalMessageInfo

func (m *Settlements) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *Settlements) GetEntries() []*SettlementEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*SettlementEntry)(nil), "seiprotocol.seichain.dex.SettlementEntry")
	proto.RegisterType((*Settlements)(nil), "seiprotocol.seichain.dex.Settlements")
}

func init() { proto.RegisterFile("dex/settlement.proto", fileDescriptor_c24d83c09612bb1c) }

var fileDescriptor_c24d83c09612bb1c = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x3d, 0x6f, 0xd3, 0x40,
	0x18, 0x8e, 0x5b, 0xfa, 0x91, 0x8b, 0x44, 0xc4, 0xd1, 0x56, 0x47, 0x07, 0x5f, 0xe4, 0x01, 0x95,
	0x21, 0xb6, 0x04, 0x1b, 0x6c, 0x21, 0x15, 0x13, 0xa2, 0x32, 0x1b, 0x8b, 0xe5, 0x9e, 0xdf, 0x26,
	0x27, 0x62, 0x9f, 0xf1, 0x5d, 0x50, 0xb2, 0xf1, 0x13, 0xf8, 0x0f, 0x0c, 0xfc, 0x03, 0x7e, 0x43,
	0xc7, 0x8e, 0x88, 0xc1, 0x42, 0xc9, 0xe6, 0xb1, 0xbf, 0x00, 0xdd, 0xb9, 0xe7, 0x94, 0x92, 0x0c,
	0x65, 0xf2, 0x7b, 0xcf, 0xfb, 0x3c, 0xcf, 0xfb, 0x9c, 0x74, 0xaf, 0xd1, 0x41, 0x02, 0xb3, 0x40,
	0x82, 0x52, 0x13, 0x48, 0x21, 0x53, 0x7e, 0x5e, 0x08, 0x25, 0x30, 0x91, 0xc0, 0x4d, 0xc5, 0xc4,
	0xc4, 0x97, 0xc0, 0xd9, 0x38, 0xe6, 0x99, 0x9f, 0xc0, 0xec, 0xf8, 0x50, 0xf3, 0x45, 0x91, 0x40,
	0x11, 0x41, 0xa6, 0x8a, 0x79, 0x2d, 0x38, 0x3e, 0x18, 0x89, 0x91, 0x30, 0x65, 0xa0, 0xab, 0x1b,
	0xb4, 0xab, 0xc9, 0x90, 0x4d, 0x53, 0x59, 0x03, 0xde, 0x8f, 0x5d, 0xd4, 0x7d, 0xdf, 0x0c, 0x3b,
	0xd5, 0x06, 0x98, 0xa0, 0xbd, 0x98, 0x31, 0x31, 0xcd, 0x14, 0x71, 0x7a, 0xce, 0x49, 0x3b, 0xb4,
	0x47, 0x1c, 0x20, 0x94, 0x17, 0x9c, 0xc1, 0x10, 0x32, 0x91, 0x92, 0x2d, 0xdd, 0x1c, 0x74, 0xab,
	0x92, 0x76, 0x0c, 0x1a, 0x25, 0x1a, 0x0e, 0x6f, 0x51, 0xb4, 0x20, 0x96, 0x12, 0x54, 0x2d, 0xd8,
	0x5e, 0x09, 0x0c, 0x6a, 0x05, 0x2b, 0x0a, 0xe6, 0x68, 0xff, 0xd3, 0x34, 0xce, 0x14, 0x57, 0x73,
	0xf2, 0xc0, 0xd0, 0xdf, 0x5e, 0x96, 0xb4, 0xf5, 0xab, 0xa4, 0x4f, 0x47, 0x5c, 0x8d, 0xa7, 0xe7,
	0x3e, 0x13, 0x69, 0xc0, 0x84, 0x4c, 0x85, 0xbc, 0xf9, 0xf4, 0x65, 0xf2, 0x31, 0x50, 0xf3, 0x1c,
	0xa4, 0x3f, 0x04, 0x56, 0x95, 0xb4, 0x71, 0xb8, 0x2e, 0x69, 0x77, 0x1e, 0xa7, 0x93, 0x97, 0x9e,
	0x45, 0xbc, 0xb0, 0x69, 0xe2, 0xef, 0x0e, 0x3a, 0x82, 0x19, 0xb0, 0xa9, 0xe2, 0x22, 0x7b, 0x2d,
	0xa4, 0x7a, 0x57, 0x9c, 0x15, 0x82, 0x01, 0x24, 0x64, 0xc7, 0x4c, 0x16, 0xf7, 0x9e, 0xfc, 0xa4,
	0xf1, 0x8b, 0x98, 0x90, 0x2a, 0x12, 0x45, 0x94, 0xd7, 0x96, 0xd7, 0x25, 0xed, 0xd5, 0x51, 0x36,
	0x52, 0xbc, 0x70, 0x43, 0x1c, 0xfc, 0xcd, 0x41, 0x87, 0x30, 0xcb, 0x81, 0x29, 0x48, 0xfe, 0x0e,
	0xba, 0x6b, 0x82, 0xa6, 0xf7, 0x0e, 0x4a, 0xac, 0xdd, 0x9a, 0x9c, 0xd4, 0xe6, 0x5c, 0xcf, 0xf0,
	0xc2, 0xf5, 0x59, 0xf0, 0x10, 0x3d, 0xca, 0x85, 0xe4, 0x3a, 0xfe, 0x90, 0x17, 0xc0, 0x74, 0x41,
	0xf6, 0x4c, 0xc0, 0xa3, 0xaa, 0xa4, 0xd8, 0x36, 0xa3, 0xc4, 0x76, 0xc3, 0x7f, 0x05, 0xf8, 0x15,
	0x7a, 0x68, 0xc1, 0xd3, 0x8b, 0x0b, 0x60, 0x8a, 0xec, 0x1b, 0x8b, 0xc7, 0x55, 0x49, 0xbb, 0x8d,
	0x05, 0x98, 0x56, 0x78, 0x87, 0xaa, 0x5f, 0xcf, 0x04, 0x3e, 0x43, 0x11, 0x8f, 0x80, 0xb4, 0xff,
	0xf7, 0xf5, 0x58, 0x87, 0xd5, 0xeb, 0xb1, 0x88, 0x17, 0x36, 0x4d, 0xef, 0x8b, 0x83, 0x3a, 0xab,
	0xc5, 0x91, 0x98, 0xa2, 0x1d, 0xc8, 0x05, 0x1b, 0x9b, 0x95, 0xd9, 0x1e, 0xb4, 0xab, 0x92, 0xd6,
	0x40, 0x58, 0x7f, 0xf0, 0x19, 0xda, 0xd3, 0xfb, 0xc9, 0x41, 0x92, 0xad, 0xde, 0xf6, 0x49, 0xe7,
	0xf9, 0x33, 0x7f, 0xd3, 0x4e, 0xfb, 0x77, 0x36, 0x72, 0xd0, 0xa9, 0x4a, 0x6a, 0xd5, 0xa1, 0x2d,
	0x06, 0x6f, 0x2e, 0x17, 0xae, 0x73, 0xb5, 0x70, 0x9d, 0xdf, 0x0b, 0xd7, 0xf9, 0xba, 0x74, 0x5b,
	0x57, 0x4b, 0xb7, 0xf5, 0x73, 0xe9, 0xb6, 0x3e, 0xf4, 0x6f, 0xdd, 0x56, 0x02, 0xef, 0xdb, 0x29,
	0xe6, 0x60, 0xc6, 0x04, 0xb3, 0x40, 0xff, 0x0a, 0xcc, 0xc5, 0xcf, 0x77, 0x4d, 0xff, 0xc5, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x4b, 0x69, 0x87, 0x7b, 0x04, 0x00, 0x00,
}

func (m *SettlementEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SettlementEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SettlementEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Leverage.Size()
		i -= size
		if _, err := m.Leverage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSettlement(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.PositionEffect) > 0 {
		i -= len(m.PositionEffect)
		copy(dAtA[i:], m.PositionEffect)
		i = encodeVarintSettlement(dAtA, i, uint64(len(m.PositionEffect)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.PositionDirection) > 0 {
		i -= len(m.PositionDirection)
		copy(dAtA[i:], m.PositionDirection)
		i = encodeVarintSettlement(dAtA, i, uint64(len(m.PositionDirection)))
		i--
		dAtA[i] = 0x3a
	}
	{
		size := m.ExpectedCostOrProceed.Size()
		i -= size
		if _, err := m.ExpectedCostOrProceed.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSettlement(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.ExecutionCostOrProceed.Size()
		i -= size
		if _, err := m.ExecutionCostOrProceed.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSettlement(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Quantity.Size()
		i -= size
		if _, err := m.Quantity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSettlement(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.AssetDenom) > 0 {
		i -= len(m.AssetDenom)
		copy(dAtA[i:], m.AssetDenom)
		i = encodeVarintSettlement(dAtA, i, uint64(len(m.AssetDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PriceDenom) > 0 {
		i -= len(m.PriceDenom)
		copy(dAtA[i:], m.PriceDenom)
		i = encodeVarintSettlement(dAtA, i, uint64(len(m.PriceDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintSettlement(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Settlements) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Settlements) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Settlements) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSettlement(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Epoch != 0 {
		i = encodeVarintSettlement(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSettlement(dAtA []byte, offset int, v uint64) int {
	offset -= sovSettlement(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SettlementEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovSettlement(uint64(l))
	}
	l = len(m.PriceDenom)
	if l > 0 {
		n += 1 + l + sovSettlement(uint64(l))
	}
	l = len(m.AssetDenom)
	if l > 0 {
		n += 1 + l + sovSettlement(uint64(l))
	}
	l = m.Quantity.Size()
	n += 1 + l + sovSettlement(uint64(l))
	l = m.ExecutionCostOrProceed.Size()
	n += 1 + l + sovSettlement(uint64(l))
	l = m.ExpectedCostOrProceed.Size()
	n += 1 + l + sovSettlement(uint64(l))
	l = len(m.PositionDirection)
	if l > 0 {
		n += 1 + l + sovSettlement(uint64(l))
	}
	l = len(m.PositionEffect)
	if l > 0 {
		n += 1 + l + sovSettlement(uint64(l))
	}
	l = m.Leverage.Size()
	n += 1 + l + sovSettlement(uint64(l))
	return n
}

func (m *Settlements) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Epoch != 0 {
		n += 1 + sovSettlement(uint64(m.Epoch))
	}
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovSettlement(uint64(l))
		}
	}
	return n
}

func sovSettlement(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSettlement(x uint64) (n int) {
	return sovSettlement(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SettlementEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSettlement
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
			return fmt.Errorf("proto: SettlementEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SettlementEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quantity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionCostOrProceed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExecutionCostOrProceed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpectedCostOrProceed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExpectedCostOrProceed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PositionDirection", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PositionDirection = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PositionEffect", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PositionEffect = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leverage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Leverage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSettlement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSettlement
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
func (m *Settlements) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSettlement
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
			return fmt.Errorf("proto: Settlements: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Settlements: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, &SettlementEntry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSettlement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSettlement
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
func skipSettlement(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSettlement
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
					return 0, ErrIntOverflowSettlement
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
					return 0, ErrIntOverflowSettlement
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
				return 0, ErrInvalidLengthSettlement
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSettlement
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSettlement
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSettlement        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSettlement          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSettlement = fmt.Errorf("proto: unexpected end of group")
)
