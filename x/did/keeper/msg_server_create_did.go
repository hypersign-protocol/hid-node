package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/did/types"
)

func (k msgServer) CreateDID(goCtx context.Context, msg *types.MsgCreateDID) (*types.MsgCreateDIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var didSpec = types.DidSpec{
		Creator:      msg.Creator,
		Did:          msg.Did,
		DidDocString: msg.DidDocString,
		CreatedAt:    msg.CreatedAt,
	}
	// Add a DID to the store and get back the ID
	id := k.AppendDID(ctx, didSpec)
	// Return the Id of the DID
	return &types.MsgCreateDIDResponse{Id: id}, nil
}
