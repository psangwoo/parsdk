package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc auth module wide codec
var ModuleCdc = codec.New()

// RegisterCodec registers concrete types on the codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgIssue{}, "issue/MsgIssue", nil)
	cdc.RegisterConcrete(MsgTransfer{}, "issue/MsgTransfer", nil)
	cdc.RegisterConcrete(MsgApprove{}, "issue/MsgApprove", nil)
	cdc.RegisterConcrete(MsgIncreaseAllowance{}, "issue/MsgIncreaseAllowance", nil)
	cdc.RegisterConcrete(MsgDecreaseAllowance{}, "issue/MsgDecreaseAllowance", nil)

	cdc.RegisterInterface((*IIssue)(nil), nil)
	cdc.RegisterConcrete(&CoinIssue{}, "issue/CoinIssue", nil)
}

func init() {
	RegisterCodec(ModuleCdc)
}
