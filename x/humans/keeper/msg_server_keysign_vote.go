package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/humansdotai/humans/x/humans/types"
)

func (k msgServer) KeysignVote(goCtx context.Context, msg *types.MsgKeysignVote) (*types.MsgKeysignVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgKeysignVoteResponse{}, nil
}
