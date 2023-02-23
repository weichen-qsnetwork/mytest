// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stratos/pot/v1/genesis.proto

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

// GenesisState defines the register module's genesis state.
type GenesisState struct {
	Params               *Params          `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty" yaml:"params"`
	TotalMinedToken      *types.Coin      `protobuf:"bytes,2,opt,name=total_mined_token,json=totalMinedToken,proto3" json:"total_mined_token,omitempty" yaml:"total_mined_token"`
	LastReportedEpoch    int64            `protobuf:"varint,3,opt,name=last_reported_epoch,json=lastReportedEpoch,proto3" json:"last_reported_epoch,omitempty" yaml:"last_reported_epoch"`
	ImmatureTotalInfo    []*ImmatureTotal `protobuf:"bytes,4,rep,name=immature_total_info,json=immatureTotalInfo,proto3" json:"immature_total_info,omitempty" yaml:"immature_total_info"`
	MatureTotalInfo      []*MatureTotal   `protobuf:"bytes,5,rep,name=mature_total_info,json=matureTotalInfo,proto3" json:"mature_total_info,omitempty" yaml:"mature_total_info"`
	IndividualRewardInfo []*Reward        `protobuf:"bytes,6,rep,name=individual_reward_info,json=individualRewardInfo,proto3" json:"individual_reward_info,omitempty" yaml:"individual_reward_info"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_abdde08c2564316a, []int{0}
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

func (m *GenesisState) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *GenesisState) GetTotalMinedToken() *types.Coin {
	if m != nil {
		return m.TotalMinedToken
	}
	return nil
}

func (m *GenesisState) GetLastReportedEpoch() int64 {
	if m != nil {
		return m.LastReportedEpoch
	}
	return 0
}

func (m *GenesisState) GetImmatureTotalInfo() []*ImmatureTotal {
	if m != nil {
		return m.ImmatureTotalInfo
	}
	return nil
}

func (m *GenesisState) GetMatureTotalInfo() []*MatureTotal {
	if m != nil {
		return m.MatureTotalInfo
	}
	return nil
}

func (m *GenesisState) GetIndividualRewardInfo() []*Reward {
	if m != nil {
		return m.IndividualRewardInfo
	}
	return nil
}

type ImmatureTotal struct {
	WalletAddress string                                   `protobuf:"bytes,1,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address" yaml:"wallet_address"`
	Value         github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=value,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"value"`
}

func (m *ImmatureTotal) Reset()         { *m = ImmatureTotal{} }
func (m *ImmatureTotal) String() string { return proto.CompactTextString(m) }
func (*ImmatureTotal) ProtoMessage()    {}
func (*ImmatureTotal) Descriptor() ([]byte, []int) {
	return fileDescriptor_abdde08c2564316a, []int{1}
}
func (m *ImmatureTotal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ImmatureTotal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ImmatureTotal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ImmatureTotal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImmatureTotal.Merge(m, src)
}
func (m *ImmatureTotal) XXX_Size() int {
	return m.Size()
}
func (m *ImmatureTotal) XXX_DiscardUnknown() {
	xxx_messageInfo_ImmatureTotal.DiscardUnknown(m)
}

var xxx_messageInfo_ImmatureTotal proto.InternalMessageInfo

func (m *ImmatureTotal) GetWalletAddress() string {
	if m != nil {
		return m.WalletAddress
	}
	return ""
}

func (m *ImmatureTotal) GetValue() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Value
	}
	return nil
}

type MatureTotal struct {
	WalletAddress string                                   `protobuf:"bytes,1,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address" yaml:"wallet_address"`
	Value         github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=value,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"value"`
}

func (m *MatureTotal) Reset()         { *m = MatureTotal{} }
func (m *MatureTotal) String() string { return proto.CompactTextString(m) }
func (*MatureTotal) ProtoMessage()    {}
func (*MatureTotal) Descriptor() ([]byte, []int) {
	return fileDescriptor_abdde08c2564316a, []int{2}
}
func (m *MatureTotal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MatureTotal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MatureTotal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MatureTotal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatureTotal.Merge(m, src)
}
func (m *MatureTotal) XXX_Size() int {
	return m.Size()
}
func (m *MatureTotal) XXX_DiscardUnknown() {
	xxx_messageInfo_MatureTotal.DiscardUnknown(m)
}

var xxx_messageInfo_MatureTotal proto.InternalMessageInfo

func (m *MatureTotal) GetWalletAddress() string {
	if m != nil {
		return m.WalletAddress
	}
	return ""
}

func (m *MatureTotal) GetValue() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "stratos.pot.v1.GenesisState")
	proto.RegisterType((*ImmatureTotal)(nil), "stratos.pot.v1.ImmatureTotal")
	proto.RegisterType((*MatureTotal)(nil), "stratos.pot.v1.MatureTotal")
}

func init() { proto.RegisterFile("stratos/pot/v1/genesis.proto", fileDescriptor_abdde08c2564316a) }

var fileDescriptor_abdde08c2564316a = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x9b, 0x75, 0xab, 0x84, 0x4b, 0x37, 0x35, 0x1b, 0x53, 0x28, 0x5b, 0x52, 0x72, 0xaa,
	0x84, 0x16, 0xd3, 0x71, 0xe3, 0xb6, 0x20, 0x84, 0x7a, 0x28, 0x42, 0x66, 0x27, 0x2e, 0x91, 0x9b,
	0x78, 0xad, 0xb5, 0xc4, 0x8e, 0x62, 0xb7, 0x63, 0x6f, 0xc1, 0x73, 0xf0, 0x1c, 0x08, 0xed, 0x82,
	0xb4, 0x23, 0xa7, 0x80, 0xda, 0x1b, 0xc7, 0x3e, 0x01, 0x8a, 0x1d, 0x44, 0x9b, 0x56, 0xdc, 0x39,
	0xd5, 0xfd, 0x7f, 0xff, 0xef, 0xff, 0xfb, 0xec, 0xba, 0x06, 0x27, 0x42, 0x66, 0x58, 0x72, 0x01,
	0x53, 0x2e, 0xe1, 0xac, 0x0f, 0xc7, 0x84, 0x11, 0x41, 0x85, 0x97, 0x66, 0x5c, 0x72, 0x73, 0xbf,
	0xac, 0x7a, 0x29, 0x97, 0xde, 0xac, 0xdf, 0x39, 0x1a, 0xf3, 0x31, 0x57, 0x25, 0x58, 0xac, 0xb4,
	0xab, 0x63, 0x87, 0x5c, 0x24, 0x5c, 0xc0, 0x11, 0x16, 0x04, 0xce, 0xfa, 0x23, 0x22, 0x71, 0x1f,
	0x86, 0x9c, 0xb2, 0xb2, 0x6e, 0x55, 0x18, 0x45, 0x98, 0xaa, 0xb8, 0xdf, 0x76, 0xc1, 0xc3, 0x37,
	0x9a, 0xf8, 0x5e, 0x62, 0x49, 0xcc, 0x0b, 0xd0, 0x48, 0x71, 0x86, 0x13, 0x61, 0x19, 0x5d, 0xa3,
	0xd7, 0x3c, 0x3f, 0xf6, 0xd6, 0x27, 0xf0, 0xde, 0xa9, 0xaa, 0xdf, 0x5e, 0xe6, 0x4e, 0xeb, 0x16,
	0x27, 0xf1, 0x4b, 0x57, 0xfb, 0x5d, 0x54, 0x36, 0x9a, 0x21, 0x68, 0x4b, 0x2e, 0x71, 0x1c, 0x24,
	0x94, 0x91, 0x28, 0x90, 0xfc, 0x9a, 0x30, 0x6b, 0x47, 0xa5, 0x3d, 0xf6, 0xf4, 0xa4, 0x5e, 0x31,
	0xa9, 0x57, 0x4e, 0xea, 0xbd, 0xe2, 0x94, 0xf9, 0x27, 0xcb, 0xdc, 0xb1, 0x74, 0xe0, 0x46, 0xb7,
	0x8b, 0x0e, 0x94, 0x36, 0x2c, 0xa4, 0xcb, 0x42, 0x31, 0xdf, 0x82, 0xc3, 0x18, 0x0b, 0x19, 0x64,
	0x24, 0xe5, 0x99, 0x24, 0x51, 0x40, 0x52, 0x1e, 0x4e, 0xac, 0x7a, 0xd7, 0xe8, 0xd5, 0x7d, 0x7b,
	0x99, 0x3b, 0x1d, 0x9d, 0xb5, 0xc5, 0xe4, 0xa2, 0x76, 0xa1, 0xa2, 0x52, 0x7c, 0x5d, 0x68, 0x66,
	0x02, 0x0e, 0x69, 0x92, 0x60, 0x39, 0xcd, 0x48, 0xa0, 0xf9, 0x94, 0x5d, 0x71, 0x6b, 0xb7, 0x5b,
	0xef, 0x35, 0xcf, 0x4f, 0xab, 0x87, 0x30, 0x28, 0xad, 0x97, 0x85, 0x73, 0x15, 0xb7, 0x25, 0xc3,
	0x45, 0x6d, 0xba, 0x6a, 0x1f, 0xb0, 0x2b, 0x6e, 0x12, 0xd0, 0xde, 0x84, 0xed, 0x29, 0xd8, 0x93,
	0x2a, 0x6c, 0xb8, 0x82, 0x5a, 0x39, 0xa5, 0x2d, 0xa0, 0x83, 0x2a, 0x86, 0x83, 0x63, 0xca, 0x22,
	0x3a, 0xa3, 0xd1, 0x14, 0xc7, 0x41, 0x46, 0x6e, 0x70, 0x16, 0x69, 0x56, 0x43, 0xb1, 0x36, 0x7e,
	0x5d, 0xa4, 0x2c, 0xfe, 0xd3, 0x65, 0xee, 0x9c, 0x96, 0x3b, 0xda, 0xda, 0xef, 0xa2, 0xa3, 0xbf,
	0x05, 0xdd, 0x54, 0x00, 0xdd, 0xaf, 0x06, 0x68, 0xad, 0x1d, 0x8e, 0x89, 0xc0, 0xfe, 0x0d, 0x8e,
	0x63, 0x22, 0x03, 0x1c, 0x45, 0x19, 0x11, 0xfa, 0x62, 0x3d, 0xf0, 0x9f, 0xfd, 0xca, 0x9d, 0x4a,
	0x65, 0x99, 0x3b, 0x8f, 0x34, 0x74, 0x5d, 0x77, 0x51, 0x4b, 0x0b, 0x17, 0xfa, 0xbb, 0x89, 0xc1,
	0xde, 0x0c, 0xc7, 0x53, 0x62, 0xed, 0xa8, 0x5d, 0xfc, 0xe3, 0x56, 0x3d, 0xbf, 0xcb, 0x9d, 0xda,
	0xe7, 0x1f, 0x4e, 0x6f, 0x4c, 0xe5, 0x64, 0x3a, 0xf2, 0x42, 0x9e, 0xc0, 0xf2, 0xcf, 0xa2, 0x3f,
	0xce, 0x44, 0x74, 0x0d, 0xe5, 0x6d, 0x4a, 0x84, 0x6a, 0x10, 0x48, 0x27, 0xbb, 0x5f, 0x0c, 0xd0,
	0x1c, 0xfe, 0xf7, 0xdb, 0xf0, 0x07, 0x77, 0x73, 0xdb, 0xb8, 0x9f, 0xdb, 0xc6, 0xcf, 0xb9, 0x6d,
	0x7c, 0x5a, 0xd8, 0xb5, 0xfb, 0x85, 0x5d, 0xfb, 0xbe, 0xb0, 0x6b, 0x1f, 0xe0, 0x4a, 0x54, 0x79,
	0x09, 0x18, 0x91, 0x7f, 0x96, 0x67, 0xe1, 0x04, 0x53, 0x06, 0x3f, 0xaa, 0x17, 0x43, 0xe5, 0x8e,
	0x1a, 0xea, 0xc5, 0x78, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x54, 0xa4, 0x8b, 0xb1, 0x04,
	0x00, 0x00,
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
	if len(m.IndividualRewardInfo) > 0 {
		for iNdEx := len(m.IndividualRewardInfo) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IndividualRewardInfo[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.MatureTotalInfo) > 0 {
		for iNdEx := len(m.MatureTotalInfo) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MatureTotalInfo[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ImmatureTotalInfo) > 0 {
		for iNdEx := len(m.ImmatureTotalInfo) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ImmatureTotalInfo[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.LastReportedEpoch != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastReportedEpoch))
		i--
		dAtA[i] = 0x18
	}
	if m.TotalMinedToken != nil {
		{
			size, err := m.TotalMinedToken.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ImmatureTotal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ImmatureTotal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ImmatureTotal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		for iNdEx := len(m.Value) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Value[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.WalletAddress) > 0 {
		i -= len(m.WalletAddress)
		copy(dAtA[i:], m.WalletAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.WalletAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MatureTotal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MatureTotal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MatureTotal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		for iNdEx := len(m.Value) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Value[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.WalletAddress) > 0 {
		i -= len(m.WalletAddress)
		copy(dAtA[i:], m.WalletAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.WalletAddress)))
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
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.TotalMinedToken != nil {
		l = m.TotalMinedToken.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.LastReportedEpoch != 0 {
		n += 1 + sovGenesis(uint64(m.LastReportedEpoch))
	}
	if len(m.ImmatureTotalInfo) > 0 {
		for _, e := range m.ImmatureTotalInfo {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MatureTotalInfo) > 0 {
		for _, e := range m.MatureTotalInfo {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.IndividualRewardInfo) > 0 {
		for _, e := range m.IndividualRewardInfo {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *ImmatureTotal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.WalletAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Value) > 0 {
		for _, e := range m.Value {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *MatureTotal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.WalletAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Value) > 0 {
		for _, e := range m.Value {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalMinedToken", wireType)
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
			if m.TotalMinedToken == nil {
				m.TotalMinedToken = &types.Coin{}
			}
			if err := m.TotalMinedToken.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastReportedEpoch", wireType)
			}
			m.LastReportedEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastReportedEpoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmatureTotalInfo", wireType)
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
			m.ImmatureTotalInfo = append(m.ImmatureTotalInfo, &ImmatureTotal{})
			if err := m.ImmatureTotalInfo[len(m.ImmatureTotalInfo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MatureTotalInfo", wireType)
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
			m.MatureTotalInfo = append(m.MatureTotalInfo, &MatureTotal{})
			if err := m.MatureTotalInfo[len(m.MatureTotalInfo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndividualRewardInfo", wireType)
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
			m.IndividualRewardInfo = append(m.IndividualRewardInfo, &Reward{})
			if err := m.IndividualRewardInfo[len(m.IndividualRewardInfo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *ImmatureTotal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ImmatureTotal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImmatureTotal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WalletAddress", wireType)
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
			m.WalletAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
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
			m.Value = append(m.Value, types.Coin{})
			if err := m.Value[len(m.Value)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MatureTotal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MatureTotal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MatureTotal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WalletAddress", wireType)
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
			m.WalletAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
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
			m.Value = append(m.Value, types.Coin{})
			if err := m.Value[len(m.Value)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
