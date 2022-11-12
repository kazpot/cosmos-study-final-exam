package keeper

import (
	"context"
	"strconv"

	"github.com/b9lab/toll-road/x/tollroad/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	RoadOperatorCreatedEventType      = "new-road-operator-created"
	RoadOperatorCreatedEventCreator   = "creator"
	RoadOperatorCreatedEventGameIndex = "road-operator-index"
	RoadOperatorCreatedEventName      = "name"
	RoadOperatorCreatedEventToken     = "token"
	RoadOperatorCreatedEventActive    = "active"
)

func (k msgServer) CreateRoadOperator(goCtx context.Context, msg *types.MsgCreateRoadOperator) (*types.MsgCreateRoadOperatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}

	nextIndex := strconv.FormatUint(systemInfo.GetNextOperatorId(), 10)

	// Check if the value already exists
	_, isFound := k.GetRoadOperator(
		ctx,
		nextIndex,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var roadOperator = types.RoadOperator{
		Creator: msg.Creator,
		Index:   nextIndex,
		Name:    msg.Name,
		Token:   msg.Token,
		Active:  msg.Active,
	}

	k.SetRoadOperator(
		ctx,
		roadOperator,
	)

	systemInfo.NextOperatorId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			RoadOperatorCreatedEventType,
			sdk.NewAttribute(RoadOperatorCreatedEventCreator, msg.Creator),
			sdk.NewAttribute(RoadOperatorCreatedEventGameIndex, nextIndex),
			sdk.NewAttribute(RoadOperatorCreatedEventName, msg.Name),
			sdk.NewAttribute(RoadOperatorCreatedEventToken, msg.Token),
			sdk.NewAttribute(RoadOperatorCreatedEventActive, strconv.FormatBool(msg.Active)),
		))

	return &types.MsgCreateRoadOperatorResponse{Index: nextIndex}, nil
}

func (k msgServer) UpdateRoadOperator(goCtx context.Context, msg *types.MsgUpdateRoadOperator) (*types.MsgUpdateRoadOperatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetRoadOperator(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var roadOperator = types.RoadOperator{
		Creator: msg.Creator,
		Index:   msg.Index,
		Name:    msg.Name,
		Token:   msg.Token,
		Active:  msg.Active,
	}

	k.SetRoadOperator(ctx, roadOperator)

	return &types.MsgUpdateRoadOperatorResponse{}, nil
}

func (k msgServer) DeleteRoadOperator(goCtx context.Context, msg *types.MsgDeleteRoadOperator) (*types.MsgDeleteRoadOperatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetRoadOperator(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveRoadOperator(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteRoadOperatorResponse{}, nil
}
