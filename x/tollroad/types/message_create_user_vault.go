package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateUserVault = "create_user_vault"

var _ sdk.Msg = &MsgCreateUserVault{}

func NewMsgCreateUserVault(creator string, roadOperatorIndex string, token string, balance uint64) *MsgCreateUserVault {
	return &MsgCreateUserVault{
		Creator:           creator,
		RoadOperatorIndex: roadOperatorIndex,
		Token:             token,
		Balance:           balance,
	}
}

func (msg *MsgCreateUserVault) Route() string {
	return RouterKey
}

func (msg *MsgCreateUserVault) Type() string {
	return TypeMsgCreateUserVault
}

func (msg *MsgCreateUserVault) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateUserVault) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateUserVault) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
