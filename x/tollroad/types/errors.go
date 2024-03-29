package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/tollroad module sentinel errors
var (
	ErrSample            = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrZeroTokens        = sdkerrors.Register(ModuleName, 1101, "zero tokens")
	ErrIndexSet          = sdkerrors.Register(ModuleName, 1102, "index already set")
	ErrSenderCannotPay   = sdkerrors.Register(ModuleName, 1103, "sender cannot pay")
	ErrInvalidSender     = sdkerrors.Register(ModuleName, 1104, "invalid sender")
	ErrCannotRefundWager = sdkerrors.Register(ModuleName, 1105, "cannot refund wager")
	ErrVaultNotFound     = sdkerrors.Register(ModuleName, 1106, "user vault not found")
)
