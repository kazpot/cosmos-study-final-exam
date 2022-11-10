package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/b9lab/toll-road/x/tollroad/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateUserVault(goCtx context.Context, msg *types.MsgCreateUserVault) (*types.MsgCreateUserVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.Keeper.GetUserVault(ctx, msg.Creator, msg.RoadOperatorIndex, msg.Token)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrIndexSet, "%s", msg.Creator)
	}

	if msg.Balance == 0 {
		return nil, sdkerrors.Wrapf(types.ErrZeroTokens, "%s", msg.Creator)
	}

	uv := types.UserVault{
		Owner:             msg.Creator,
		RoadOperatorIndex: msg.RoadOperatorIndex,
		Token:             msg.Token,
		Balance:           msg.Balance,
	}

	err := k.Keeper.CollectWager(ctx, &uv)
	if err != nil {
		return nil, err
	}

	k.Keeper.SetUserVault(ctx, uv)

	return &types.MsgCreateUserVaultResponse{}, nil
}
