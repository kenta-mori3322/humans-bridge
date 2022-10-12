package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/humansdotai/humans/x/humans/types"
)

func (k msgServer) RequestTransaction(goCtx context.Context, msg *types.MsgRequestTransaction) (*types.MsgRequestTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRequestTransactionResponse{}, nil
}
