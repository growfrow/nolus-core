// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: treasury/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the treasury module's genesis state.
type GenesisState struct {
	FeeRate     github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,1,opt,name=fee_rate,json=feeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fee_rate"`
	FeeCaps     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=fee_caps,json=feeCapstest,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee_capstest"`
	FeeProceeds github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=fee_proceeds,json=feeProceeds,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee_proceeds"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_35e88080a694bc53, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetFeeCaps() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.FeeCaps
	}
	return nil
}

func (m *GenesisState) GetFeeProceeds() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.FeeProceeds
	}
	return nil
}

// TaxCap is the max tax amount can be charged for the given denom
type FeeCap struct {
	Denom  string                                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	TaxCap github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=tax_cap,json=taxCap,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"tax_cap"`
}

func (m *FeeCap) Reset()         { *m = FeeCap{} }
func (m *FeeCap) String() string { return proto.CompactTextString(m) }
func (*FeeCap) ProtoMessage()    {}
func (*FeeCap) Descriptor() ([]byte, []int) {
	return fileDescriptor_35e88080a694bc53, []int{1}
}
func (m *FeeCap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeCap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeCap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeCap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeCap.Merge(m, src)
}
func (m *FeeCap) XXX_Size() int {
	return m.Size()
}
func (m *FeeCap) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeCap.DiscardUnknown(m)
}

var xxx_messageInfo_FeeCap proto.InternalMessageInfo

func (m *FeeCap) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nomo.cosmzone.suspend.GenesisState")
	proto.RegisterType((*FeeCap)(nil), "nomo.cosmzone.suspend.FeeCap")
}

func init() { proto.RegisterFile("treasury/genesis.proto", fileDescriptor_35e88080a694bc53) }

var fileDescriptor_35e88080a694bc53 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xb1, 0x6e, 0xea, 0x30,
	0x14, 0x86, 0x13, 0xd0, 0x85, 0x7b, 0x03, 0x53, 0x84, 0x50, 0x2e, 0x43, 0x40, 0x0c, 0x57, 0x2c,
	0xd8, 0x97, 0x76, 0xea, 0x0a, 0x55, 0x11, 0x5b, 0x15, 0xb6, 0x2e, 0x95, 0x93, 0x9c, 0xa4, 0x56,
	0x1b, 0x9f, 0x28, 0x36, 0x15, 0xf4, 0x29, 0xfa, 0x1c, 0xdd, 0xfa, 0x16, 0x8c, 0x8c, 0x55, 0x07,
	0x5a, 0xc1, 0x8b, 0x54, 0x8e, 0xa1, 0x73, 0x2b, 0x75, 0x4a, 0x6c, 0xff, 0xe7, 0xff, 0x8e, 0x8f,
	0x7f, 0xa7, 0xad, 0x0a, 0x60, 0x72, 0x51, 0xac, 0x68, 0x0a, 0x02, 0x24, 0x97, 0x24, 0x2f, 0x50,
	0xa1, 0xdb, 0x16, 0x98, 0x21, 0x89, 0x50, 0x66, 0x0f, 0x28, 0x80, 0x1c, 0x55, 0x9d, 0x56, 0x8a,
	0x29, 0x96, 0x12, 0xaa, 0xff, 0x8c, 0xba, 0xe3, 0x6b, 0x21, 0x4a, 0x1a, 0x32, 0x09, 0xf4, 0x7e,
	0x14, 0x82, 0x62, 0x23, 0x1a, 0x21, 0x17, 0xe6, 0xbc, 0xff, 0x5c, 0x71, 0x9a, 0x53, 0xe3, 0x3f,
	0x57, 0x4c, 0x81, 0x3b, 0x73, 0x7e, 0x27, 0x00, 0xd7, 0x05, 0x53, 0xe0, 0xd9, 0x3d, 0x7b, 0xf0,
	0x67, 0x4c, 0xd6, 0xdb, 0xae, 0xf5, 0xba, 0xed, 0xfe, 0x4b, 0xb9, 0xba, 0x59, 0x84, 0x24, 0xc2,
	0x8c, 0x1e, 0x5c, 0xcd, 0x67, 0x28, 0xe3, 0x5b, 0xaa, 0x56, 0x39, 0x48, 0x72, 0x0e, 0x51, 0x50,
	0x4f, 0x00, 0x02, 0x6d, 0x95, 0x18, 0xab, 0x88, 0xe5, 0xd2, 0xab, 0xf4, 0xaa, 0x83, 0xc6, 0xc9,
	0x5f, 0x62, 0x2a, 0x88, 0x6e, 0x87, 0x1c, 0xda, 0x21, 0x13, 0xe4, 0x62, 0xfc, 0x5f, 0x53, 0x9e,
	0xde, 0xba, 0x83, 0x2f, 0x50, 0x74, 0x81, 0x2c, 0x39, 0x13, 0x96, 0x4b, 0x57, 0x38, 0x4d, 0xcd,
	0xc9, 0x0b, 0x8c, 0x00, 0x62, 0xe9, 0x55, 0x7f, 0x9e, 0xd5, 0x48, 0x00, 0x2e, 0x0f, 0xfe, 0xfd,
	0xd4, 0xa9, 0x5d, 0x94, 0x68, 0xb7, 0xe5, 0xfc, 0x8a, 0x41, 0x60, 0x66, 0x26, 0x15, 0x98, 0x85,
	0x3b, 0x75, 0xea, 0x8a, 0x2d, 0xf5, 0xbd, 0xbd, 0xca, 0xb7, 0x27, 0x38, 0x13, 0x2a, 0xa8, 0x29,
	0xb6, 0x9c, 0xb0, 0x7c, 0x3c, 0x5f, 0xef, 0x7c, 0x7b, 0xb3, 0xf3, 0xed, 0xf7, 0x9d, 0x6f, 0x3f,
	0xee, 0x7d, 0x6b, 0xb3, 0xf7, 0xad, 0x97, 0xbd, 0x6f, 0x5d, 0x9d, 0xa5, 0x5c, 0xdd, 0xb1, 0x70,
	0x68, 0xb2, 0x50, 0x40, 0xcc, 0xa5, 0xe4, 0x19, 0x12, 0x01, 0x8a, 0xea, 0x3d, 0x7a, 0xcc, 0x07,
	0x5d, 0xd2, 0xcf, 0x1c, 0x95, 0x80, 0xb0, 0x56, 0x3e, 0xfc, 0xe9, 0x47, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf4, 0x1b, 0x98, 0x72, 0x60, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeeProceeds) > 0 {
		for iNdEx := len(m.FeeProceeds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeProceeds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.FeeCaps) > 0 {
		for iNdEx := len(m.FeeCaps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeCaps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.FeeRate.Size()
		i -= size
		if _, err := m.FeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *FeeCap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeCap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeCap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TaxCap.Size()
		i -= size
		if _, err := m.TaxCap.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.FeeRate.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.FeeCaps) > 0 {
		for _, e := range m.FeeCaps {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.FeeProceeds) > 0 {
		for _, e := range m.FeeProceeds {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *FeeCap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.TaxCap.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeCaps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeCaps = append(m.FeeCaps, types.Coin{})
			if err := m.FeeCaps[len(m.FeeCaps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeProceeds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeProceeds = append(m.FeeProceeds, types.Coin{})
			if err := m.FeeProceeds[len(m.FeeProceeds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *FeeCap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: FeeCap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeCap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxCap", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TaxCap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
