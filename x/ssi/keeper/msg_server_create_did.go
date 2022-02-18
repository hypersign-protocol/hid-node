package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	utils "github.com/hypersign-protocol/hid-node/x/ssi/utils"
)

func (k msgServer) CreateDID(goCtx context.Context, msg *types.MsgCreateDID) (*types.MsgCreateDIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	didMsg := msg.GetDidDocString()
	did := msg.GetDidDocString().GetId()
	// Checks if the DID has a valid format
	if !utils.IsValidDid(did) {
		return nil, types.ErrBadRequestIsNotDid.Wrap(did)
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
	if k.HasDid(ctx, did) {
		return nil, sdkerrors.Wrap(types.ErrDidDocExists, fmt.Sprintf("DID already exists %s", did))
	}

	if k.ValidateDidControllers(&ctx, did, didMsg.GetController(), didMsg.GetVerificationMethod()) != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, "DID controller is not valid")
	}

	var didSpec = types.Did{
		Context:              didMsg.GetContext(),
		Id:                   didMsg.GetId(),
		Controller:           didMsg.GetController(),
		AlsoKnownAs:          didMsg.GetAlsoKnownAs(),
		VerificationMethod:   didMsg.GetVerificationMethod(),
		Authentication:       didMsg.GetAuthentication(),
		AssertionMethod:      didMsg.GetAssertionMethod(),
		KeyAgreement:         didMsg.GetKeyAgreement(),
		CapabilityInvocation: didMsg.GetCapabilityInvocation(),
		Created:              didMsg.GetCreated(),
		Updated:              didMsg.GetUpdated(),
	}
	// Add a DID to the store and get back the ID
	id := k.AppendDID(ctx, didSpec)
	// Return the Id of the DID
	return &types.MsgCreateDIDResponse{Id: id}, nil
}
