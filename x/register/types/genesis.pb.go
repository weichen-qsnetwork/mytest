// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stratos/register/v1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/x/staking/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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
	Params            *Params                                `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty" yaml:"params"`
	ResourceNodes     ResourceNodes                          `protobuf:"bytes,2,rep,name=resource_nodes,json=resourceNodes,proto3,castrepeated=ResourceNodes" json:"resource_nodes" yaml:"resource_nodes"`
	MetaNodes         MetaNodes                              `protobuf:"bytes,3,rep,name=meta_nodes,json=metaNodes,proto3,castrepeated=MetaNodes" json:"meta_nodes" yaml:"meta_nodes"`
	RemainingNozLimit github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=remaining_noz_limit,json=remainingNozLimit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"remaining_noz_limit" yaml:"remaining_noz_limit"`
	Slashing          []*Slashing                            `protobuf:"bytes,5,rep,name=slashing,proto3" json:"slashing,omitempty" yaml:"slashing_info"`
	StakeNozRate      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=stake_noz_rate,json=stakeNozRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"stake_noz_rate" yaml:"stake_noz_rate"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bdab54ebea9e48e, []int{0}
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

func (m *GenesisState) GetResourceNodes() ResourceNodes {
	if m != nil {
		return m.ResourceNodes
	}
	return nil
}

func (m *GenesisState) GetMetaNodes() MetaNodes {
	if m != nil {
		return m.MetaNodes
	}
	return nil
}

func (m *GenesisState) GetSlashing() []*Slashing {
	if m != nil {
		return m.Slashing
	}
	return nil
}

type GenesisMetaNode struct {
	NetworkAddress string            `protobuf:"bytes,1,opt,name=network_address,json=networkAddress,proto3" json:"network_address,omitempty" yaml:"network_address"`
	Pubkey         *types.Any        `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty" yaml:"pubkey"`
	Suspend        bool              `protobuf:"varint,3,opt,name=suspend,proto3" json:"suspend,omitempty" yaml:"suspend"`
	Status         types1.BondStatus `protobuf:"varint,4,opt,name=status,proto3,enum=cosmos.staking.v1beta1.BondStatus" json:"status,omitempty" yaml:"status"`
	Tokens         string            `protobuf:"bytes,5,opt,name=tokens,proto3" json:"tokens,omitempty" yaml:"token"`
	OwnerAddress   string            `protobuf:"bytes,6,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty" yaml:"owner_address"`
	Description    *Description      `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty" yaml:"description",omitempty`
}

func (m *GenesisMetaNode) Reset()         { *m = GenesisMetaNode{} }
func (m *GenesisMetaNode) String() string { return proto.CompactTextString(m) }
func (*GenesisMetaNode) ProtoMessage()    {}
func (*GenesisMetaNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bdab54ebea9e48e, []int{1}
}
func (m *GenesisMetaNode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisMetaNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisMetaNode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisMetaNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisMetaNode.Merge(m, src)
}
func (m *GenesisMetaNode) XXX_Size() int {
	return m.Size()
}
func (m *GenesisMetaNode) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisMetaNode.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisMetaNode proto.InternalMessageInfo

func (m *GenesisMetaNode) GetNetworkAddress() string {
	if m != nil {
		return m.NetworkAddress
	}
	return ""
}

func (m *GenesisMetaNode) GetPubkey() *types.Any {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *GenesisMetaNode) GetSuspend() bool {
	if m != nil {
		return m.Suspend
	}
	return false
}

func (m *GenesisMetaNode) GetStatus() types1.BondStatus {
	if m != nil {
		return m.Status
	}
	return types1.Unspecified
}

func (m *GenesisMetaNode) GetTokens() string {
	if m != nil {
		return m.Tokens
	}
	return ""
}

func (m *GenesisMetaNode) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *GenesisMetaNode) GetDescription() *Description {
	if m != nil {
		return m.Description
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "stratos.register.v1.GenesisState")
	proto.RegisterType((*GenesisMetaNode)(nil), "stratos.register.v1.GenesisMetaNode")
}

func init() { proto.RegisterFile("stratos/register/v1/genesis.proto", fileDescriptor_5bdab54ebea9e48e) }

var fileDescriptor_5bdab54ebea9e48e = []byte{
	// 748 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6f, 0xfb, 0x34,
	0x14, 0x6e, 0x56, 0xd6, 0xad, 0x5e, 0xdb, 0xb1, 0xac, 0xa0, 0xac, 0x40, 0xd3, 0x59, 0x68, 0x2a,
	0xd2, 0x96, 0xa8, 0x83, 0x13, 0x12, 0x87, 0x85, 0x89, 0x09, 0xc1, 0xaa, 0x29, 0x3d, 0x20, 0x71,
	0xa9, 0xdc, 0xc4, 0xcb, 0xa2, 0x36, 0x76, 0x14, 0xbb, 0xdb, 0xb2, 0xbf, 0x62, 0x12, 0xff, 0x05,
	0x67, 0xfe, 0x88, 0x89, 0xd3, 0x8e, 0x88, 0x43, 0x40, 0xdb, 0x8d, 0x63, 0xae, 0x5c, 0x50, 0x6d,
	0xa7, 0xcd, 0xa0, 0x42, 0xfa, 0x9d, 0xea, 0xf7, 0xde, 0xf7, 0xbe, 0x2f, 0x7e, 0xfe, 0x5e, 0xc1,
	0x21, 0xe3, 0x09, 0xe2, 0x94, 0xd9, 0x09, 0x0e, 0x42, 0xc6, 0x71, 0x62, 0xdf, 0x0e, 0xec, 0x00,
	0x13, 0xcc, 0x42, 0x66, 0xc5, 0x09, 0xe5, 0x54, 0xdf, 0x57, 0x10, 0xab, 0x80, 0x58, 0xb7, 0x83,
	0xce, 0x41, 0x40, 0x69, 0x30, 0xc3, 0xb6, 0x80, 0x4c, 0xe6, 0xd7, 0x36, 0x22, 0xa9, 0xc4, 0x77,
	0xda, 0x01, 0x0d, 0xa8, 0x38, 0xda, 0x8b, 0x93, 0xca, 0x1e, 0x78, 0x94, 0x45, 0x94, 0x8d, 0x65,
	0x41, 0x06, 0xaa, 0xf4, 0xa9, 0x8c, 0x6c, 0xc6, 0xd1, 0x34, 0x24, 0x81, 0x7d, 0x3b, 0x98, 0x60,
	0x8e, 0x06, 0x45, 0xac, 0x50, 0x70, 0xdd, 0x97, 0x2e, 0x3f, 0x49, 0x60, 0xe0, 0xe3, 0x26, 0x68,
	0x5c, 0xc8, 0x8f, 0x1f, 0x71, 0xc4, 0xb1, 0xfe, 0x0d, 0xa8, 0xc5, 0x28, 0x41, 0x11, 0x33, 0xb4,
	0x9e, 0xd6, 0xdf, 0x39, 0xfd, 0xc8, 0x5a, 0x73, 0x19, 0xeb, 0x4a, 0x40, 0x9c, 0xbd, 0x3c, 0x33,
	0x9b, 0x29, 0x8a, 0x66, 0x5f, 0x42, 0xd9, 0x04, 0x5d, 0xd5, 0xad, 0xdf, 0x83, 0x56, 0x82, 0x19,
	0x9d, 0x27, 0x1e, 0x1e, 0x13, 0xea, 0x63, 0x66, 0x6c, 0xf4, 0xaa, 0xfd, 0x9d, 0xd3, 0xc3, 0xb5,
	0x7c, 0xae, 0x82, 0x0e, 0xa9, 0x8f, 0x1d, 0xeb, 0x29, 0x33, 0x2b, 0x79, 0x66, 0x7e, 0x20, 0x99,
	0xdf, 0xd2, 0xc0, 0x9f, 0xff, 0x30, 0x9b, 0x65, 0x38, 0x73, 0x9b, 0x49, 0x39, 0xd4, 0x7d, 0x00,
	0x22, 0xcc, 0x91, 0x52, 0xad, 0x0a, 0xd5, 0x4f, 0xd6, 0xaa, 0x5e, 0x62, 0x8e, 0x84, 0xe2, 0x91,
	0x52, 0xdc, 0x93, 0x8a, 0xab, 0xf6, 0x85, 0x5a, 0xbd, 0x80, 0x31, 0xb7, 0x1e, 0x15, 0x47, 0xfd,
	0x27, 0x0d, 0xec, 0x27, 0x38, 0x42, 0x21, 0x09, 0x49, 0x30, 0x26, 0xf4, 0x61, 0x3c, 0x0b, 0xa3,
	0x90, 0x1b, 0xef, 0xf5, 0xb4, 0x7e, 0xdd, 0xf1, 0x16, 0x84, 0xbf, 0x67, 0xe6, 0x51, 0x10, 0xf2,
	0x9b, 0xf9, 0xc4, 0xf2, 0x68, 0xa4, 0x5e, 0x50, 0xfd, 0x9c, 0x30, 0x7f, 0x6a, 0xf3, 0x34, 0xc6,
	0xcc, 0xfa, 0x96, 0xf0, 0xbf, 0x32, 0x73, 0x1d, 0x59, 0x9e, 0x99, 0x9d, 0x62, 0x06, 0xff, 0x29,
	0x42, 0x77, 0x6f, 0x99, 0x1d, 0xd2, 0x87, 0xef, 0x17, 0x39, 0x7d, 0x04, 0xb6, 0xd9, 0x0c, 0xb1,
	0x9b, 0x90, 0x04, 0xc6, 0xe6, 0xff, 0xdc, 0x7c, 0xa4, 0x40, 0x8e, 0x91, 0x67, 0x66, 0x5b, 0x6a,
	0x14, 0x8d, 0xe3, 0x90, 0x5c, 0x53, 0xe8, 0x2e, 0x89, 0xf4, 0x08, 0xb4, 0x16, 0xc6, 0xc2, 0x42,
	0x3b, 0x41, 0x1c, 0x1b, 0x35, 0x71, 0xc9, 0x8b, 0x77, 0xb8, 0xe4, 0x39, 0xf6, 0x56, 0x2f, 0xfa,
	0x96, 0x0d, 0xba, 0x0d, 0x91, 0x18, 0xd2, 0x07, 0x77, 0x11, 0xfe, 0x5d, 0x05, 0xbb, 0xca, 0x92,
	0xc5, 0xe4, 0xf5, 0xaf, 0xc1, 0x2e, 0xc1, 0xfc, 0x8e, 0x26, 0xd3, 0x31, 0xf2, 0xfd, 0x04, 0x33,
	0x69, 0xcf, 0xba, 0xd3, 0xc9, 0x33, 0xf3, 0x43, 0xc9, 0xfa, 0x2f, 0x00, 0x74, 0x5b, 0x2a, 0x73,
	0x26, 0x13, 0xfa, 0x0f, 0xa0, 0x16, 0xcf, 0x27, 0x53, 0x9c, 0x1a, 0x1b, 0xc2, 0xda, 0x6d, 0x4b,
	0xae, 0xa4, 0x55, 0xac, 0xa4, 0x75, 0x46, 0x52, 0xe7, 0xb3, 0x92, 0xa7, 0x05, 0x1a, 0xfe, 0xfa,
	0xcb, 0x49, 0x5b, 0xad, 0x9f, 0x97, 0xa4, 0x31, 0xa7, 0xd6, 0xd5, 0x7c, 0xf2, 0x1d, 0x4e, 0x5d,
	0x45, 0xa7, 0x1f, 0x83, 0x2d, 0x36, 0x67, 0x31, 0x26, 0xbe, 0x51, 0xed, 0x69, 0xfd, 0x6d, 0x47,
	0xcf, 0x33, 0xb3, 0xa5, 0xee, 0x2a, 0x0b, 0xd0, 0x2d, 0x20, 0xfa, 0x25, 0xa8, 0x31, 0x8e, 0xf8,
	0x9c, 0x09, 0xaf, 0xb4, 0x4e, 0xa1, 0xa5, 0xc8, 0x8b, 0xed, 0x55, 0xdb, 0x6c, 0x39, 0x94, 0xf8,
	0x23, 0x81, 0x2c, 0x2f, 0x9a, 0xec, 0x85, 0xae, 0x22, 0xd1, 0xfb, 0xa0, 0xc6, 0xe9, 0x14, 0x13,
	0x66, 0x6c, 0x8a, 0x89, 0xbc, 0x9f, 0x67, 0x66, 0x43, 0x42, 0x45, 0x1e, 0xba, 0xaa, 0xae, 0x7f,
	0x05, 0x9a, 0xf4, 0x8e, 0xe0, 0x64, 0x39, 0x42, 0xf9, 0x8c, 0x25, 0x0b, 0xbc, 0x29, 0x43, 0xb7,
	0x21, 0xe2, 0x62, 0x7c, 0x3e, 0xd8, 0xf1, 0x31, 0xf3, 0x92, 0x30, 0xe6, 0x21, 0x25, 0xc6, 0x96,
	0x98, 0x61, 0x6f, 0xad, 0xbd, 0xce, 0x57, 0x38, 0xa7, 0x97, 0x67, 0xe6, 0xc7, 0x92, 0xbe, 0xd4,
	0x0e, 0x8f, 0x69, 0x14, 0x72, 0x1c, 0xc5, 0x3c, 0x75, 0xcb, 0xb4, 0xce, 0xf0, 0xe9, 0xa5, 0xab,
	0x3d, 0xbf, 0x74, 0xb5, 0x3f, 0x5f, 0xba, 0xda, 0xe3, 0x6b, 0xb7, 0xf2, 0xfc, 0xda, 0xad, 0xfc,
	0xf6, 0xda, 0xad, 0xfc, 0xf8, 0x45, 0xc9, 0x66, 0x4a, 0x94, 0x60, 0x5e, 0x1c, 0x4f, 0xbc, 0x1b,
	0x14, 0x12, 0xfb, 0x7e, 0xf5, 0x67, 0x27, 0x8c, 0x37, 0xa9, 0x89, 0xc7, 0xfd, 0xfc, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xe6, 0xa3, 0xcc, 0x8a, 0xb7, 0x05, 0x00, 0x00,
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
	{
		size := m.StakeNozRate.Size()
		i -= size
		if _, err := m.StakeNozRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Slashing) > 0 {
		for iNdEx := len(m.Slashing) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Slashing[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size := m.RemainingNozLimit.Size()
		i -= size
		if _, err := m.RemainingNozLimit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.MetaNodes) > 0 {
		for iNdEx := len(m.MetaNodes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MetaNodes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ResourceNodes) > 0 {
		for iNdEx := len(m.ResourceNodes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ResourceNodes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *GenesisMetaNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisMetaNode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisMetaNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Description != nil {
		{
			size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Tokens) > 0 {
		i -= len(m.Tokens)
		copy(dAtA[i:], m.Tokens)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Tokens)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Status != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.Suspend {
		i--
		if m.Suspend {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Pubkey != nil {
		{
			size, err := m.Pubkey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.NetworkAddress) > 0 {
		i -= len(m.NetworkAddress)
		copy(dAtA[i:], m.NetworkAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.NetworkAddress)))
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
	if len(m.ResourceNodes) > 0 {
		for _, e := range m.ResourceNodes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MetaNodes) > 0 {
		for _, e := range m.MetaNodes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.RemainingNozLimit.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Slashing) > 0 {
		for _, e := range m.Slashing {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.StakeNozRate.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *GenesisMetaNode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NetworkAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Pubkey != nil {
		l = m.Pubkey.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Suspend {
		n += 2
	}
	if m.Status != 0 {
		n += 1 + sovGenesis(uint64(m.Status))
	}
	l = len(m.Tokens)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Description != nil {
		l = m.Description.Size()
		n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceNodes", wireType)
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
			m.ResourceNodes = append(m.ResourceNodes, ResourceNode{})
			if err := m.ResourceNodes[len(m.ResourceNodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaNodes", wireType)
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
			m.MetaNodes = append(m.MetaNodes, MetaNode{})
			if err := m.MetaNodes[len(m.MetaNodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingNozLimit", wireType)
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
			if err := m.RemainingNozLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slashing", wireType)
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
			m.Slashing = append(m.Slashing, &Slashing{})
			if err := m.Slashing[len(m.Slashing)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeNozRate", wireType)
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
			if err := m.StakeNozRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GenesisMetaNode) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisMetaNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisMetaNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkAddress", wireType)
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
			m.NetworkAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
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
			if m.Pubkey == nil {
				m.Pubkey = &types.Any{}
			}
			if err := m.Pubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suspend", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Suspend = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types1.BondStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
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
			m.Tokens = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
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
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			if m.Description == nil {
				m.Description = &Description{}
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
