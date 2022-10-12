package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/humansdotai/humans/x/humans/types"
)

func (k msgServer) AddWhitelisted(goCtx context.Context, msg *types.MsgAddWhitelisted) (*types.MsgAddWhitelistedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddWhitelistedResponse{}, nil
}
