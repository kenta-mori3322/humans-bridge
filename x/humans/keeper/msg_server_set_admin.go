package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/humansdotai/humans/x/humans/types"
)

func (k msgServer) SetAdmin(goCtx context.Context, msg *types.MsgSetAdmin) (*types.MsgSetAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetAdminResponse{}, nil
}
