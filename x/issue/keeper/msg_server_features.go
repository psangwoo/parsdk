package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

func (k msgServer) Features(goCtx context.Context, msg *types.MsgFeatures) (*types.MsgFeaturesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgFeaturesResponse{}, nil
}