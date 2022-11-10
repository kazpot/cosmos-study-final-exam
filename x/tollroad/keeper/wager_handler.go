package keeper

import (
	"github.com/b9lab/toll-road/x/tollroad/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func coinsOf(amount uint64, denom string) sdk.Coins {
	return sdk.Coins{
		sdk.Coin{
			Denom:  denom,
			Amount: sdk.NewInt(int64(amount)),
		},
	}
}

func (k *Keeper) CollectWager(ctx sdk.Context, userVault *types.UserVault) error {
	sender, err := sdk.AccAddressFromBech32(userVault.GetOwner())
	if err != nil {
		return sdkerrors.Wrapf(err, types.ErrInvalidSender.Error())
	}

	err = k.bank.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, coinsOf(userVault.GetBalance(), userVault.GetToken()))
	if err != nil {
		return sdkerrors.Wrapf(err, types.ErrSenderCannotPay.Error())
	}
	return nil
}

func (k *Keeper) RefundWager(ctx sdk.Context, userVault *types.UserVault) {
	sender, err := sdk.AccAddressFromBech32(userVault.GetOwner())
	if err != nil {
		panic(err.Error())
	}

	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, coinsOf(userVault.GetBalance(), userVault.GetToken()))
	if err != nil {
		panic("bank error")
	}
}
