package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateUserVault = "update_user_vault"

var _ sdk.Msg = &MsgUpdateUserVault{}

func NewMsgUpdateUserVault(creator string, roadOperatorIndex string, token string, balance uint64) *MsgUpdateUserVault {
	return &MsgUpdateUserVault{
		Creator:           creator,
		RoadOperatorIndex: roadOperatorIndex,
		Token:             token,
		Balance:           balance,
	}
}

func (msg *MsgUpdateUserVault) Route() string {
	return RouterKey
}

func (msg *MsgUpdateUserVault) Type() string {
	return TypeMsgUpdateUserVault
}

func (msg *MsgUpdateUserVault) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateUserVault) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateUserVault) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
