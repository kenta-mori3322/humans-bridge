package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/humansdotai/humans/x/humans/types"
)

func (k msgServer) TransferPoolcoin(goCtx context.Context, msg *types.MsgTransferPoolcoin) (*types.MsgTransferPoolcoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTransferPoolcoinResponse{}, nil
}
