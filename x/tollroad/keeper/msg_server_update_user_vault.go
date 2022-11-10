package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/b9lab/toll-road/x/tollroad/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateUserVault(goCtx context.Context, msg *types.MsgUpdateUserVault) (*types.MsgUpdateUserVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userVault, found := k.Keeper.GetUserVault(ctx, msg.Creator, msg.RoadOperatorIndex, msg.Token)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	if msg.Balance == 0 {
		return nil, sdkerrors.Wrapf(types.ErrZeroTokens, "%s", msg.Creator)
	}

	if msg.Balance > userVault.GetBalance() {
		uv := types.UserVault{
			Owner:             msg.Creator,
			RoadOperatorIndex: msg.RoadOperatorIndex,
			Token:             msg.Token,
			Balance:           msg.Balance - userVault.GetBalance(),
		}

		err := k.Keeper.CollectWager(ctx, &uv)
		if err != nil {
			return nil, err
		}
		userVault.Balance = msg.Balance
		k.Keeper.SetUserVault(ctx, userVault)
	} else if msg.Balance < userVault.GetBalance() {
		uv := types.UserVault{
			Owner:             msg.Creator,
			RoadOperatorIndex: msg.RoadOperatorIndex,
			Token:             msg.Token,
			Balance:           userVault.GetBalance() - msg.Balance,
		}
		k.Keeper.RefundWager(ctx, &uv)
		userVault.Balance = msg.Balance
		k.Keeper.SetUserVault(ctx, userVault)
	}

	return &types.MsgUpdateUserVaultResponse{}, nil
}
