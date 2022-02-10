package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utils "github.com/hypersign-protocol/hid-node/x/ssi/utils"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func (k msgServer) CreateDID(goCtx context.Context, msg *types.MsgCreateDID) (*types.MsgCreateDIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	didMsg := msg.GetDidDocString()

	// Checks if the DID has a valid format
	if !utils.IsValidDid(msg.Did) {
		return nil, types.ErrBadRequestIsNotDid.Wrap(msg.Did)
	}

	// Checks if the DidDoc is a valid format
	didDocCheck := utils.IsValidDidDoc(msg.DidDocString)
	if didDocCheck != "" {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, didDocCheck)
	}  
	
	// Signature check
	if err := k.VerifySignature(&ctx, didMsg, didMsg.GetSigners(), msg.GetSignatures()); err != nil {
		return nil, err
	}

	// Checks if the DID is already present in the store
	if k.HasDid(ctx, msg.Did) {
		return nil, sdkerrors.Wrap(types.ErrDidDocExists, fmt.Sprintf("DID already exists %s", msg.Did))
	}

	var didSpec = types.Did{
		Did:          msg.Did,
		DidDocString: msg.DidDocString,
	}
	// Add a DID to the store and get back the ID
	id := k.AppendDID(ctx, didSpec)
	// Return the Id of the DID
	return &types.MsgCreateDIDResponse{Id: id}, nil
}
