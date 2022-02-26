package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
	cdc.RegisterConcrete(MsgVolumeReport{}, "pot/VolumeReportTx", nil)
	cdc.RegisterConcrete(MsgWithdraw{}, "pot/WithdrawTx", nil)
	cdc.RegisterConcrete(MsgFoundationDeposit{}, "pot/FoundationDepositTx", nil)
	cdc.RegisterConcrete(MsgSlashingResourceNode{}, "pot/SlashingResourceNodeTx", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
