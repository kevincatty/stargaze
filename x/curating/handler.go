package curating

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// NewHandler creates an sdk.Handler for all the x/curating type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgPost:
			return handleMsgPost(ctx, k, msg)
		case types.MsgUpvote:
			return handleMsgUpvote(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgUpvote calls the keeper to perform the upvote operation
func handleMsgUpvote(ctx sdk.Context, k Keeper, msg types.MsgUpvote) (*sdk.Result, error) {
	err := k.CreateUpvote(
		ctx, msg.VendorID, msg.PostID, msg.Curator, msg.RewardAccount, msg.VoteNum, msg.Deposit)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Curator.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgPost(ctx sdk.Context, k Keeper, msg types.MsgPost) (*sdk.Result, error) {
	err := k.CreatePost(
		ctx, msg.VendorID, msg.PostID, msg.Body, msg.Deposit, msg.Creator, msg.RewardAccount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
