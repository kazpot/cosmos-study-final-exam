package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteUserVault = "delete_user_vault"

var _ sdk.Msg = &MsgDeleteUserVault{}

func NewMsgDeleteUserVault(creator string, roadOperatorIndex string, token string) *MsgDeleteUserVault {
	return &MsgDeleteUserVault{
		Creator:           creator,
		RoadOperatorIndex: roadOperatorIndex,
		Token:             token,
	}
}

func (msg *MsgDeleteUserVault) Route() string {
	return RouterKey
}

func (msg *MsgDeleteUserVault) Type() string {
	return TypeMsgDeleteUserVault
}

func (msg *MsgDeleteUserVault) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteUserVault) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteUserVault) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
