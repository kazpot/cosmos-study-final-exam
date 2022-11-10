package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/b9lab/toll-road/x/tollroad/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteUserVault(goCtx context.Context, msg *types.MsgDeleteUserVault) (*types.MsgDeleteUserVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userVault, found := k.Keeper.GetUserVault(ctx, msg.Creator, msg.RoadOperatorIndex, msg.Token)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	k.Keeper.RefundWager(ctx, &userVault)
	k.Keeper.RemoveUserVault(
		ctx,
		userVault.GetOwner(),
		userVault.GetRoadOperatorIndex(),
		userVault.GetToken(),
	)

	return &types.MsgDeleteUserVaultResponse{}, nil
}
