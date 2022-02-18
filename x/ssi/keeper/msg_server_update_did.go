package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func (k msgServer) UpdateDID(goCtx context.Context, msg *types.MsgUpdateDID) (*types.MsgUpdateDIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	didMsg := msg.GetDidDocString()
	did := msg.GetDidDocString().GetId()

	oldDIDDoc, err := k.GetDid(&ctx, didMsg.Id)
	if err != nil {
		return nil, err
	}

	// TODO: Implement this when we have generic type for Create and Update DID
	// didDocCheck := utils.IsValidDidDoc(didMsg)
	// if didDocCheck != "" {
	// 	return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, didDocCheck)
	// }

	// Checks if the DID is not present in the store
	if !k.HasDid(ctx, did) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("DID doesnt exists %s", did))
	}

	if k.ValidateDidControllers(&ctx, did, didMsg.GetController(), didMsg.GetVerificationMethod()) != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, "DID controller is not valid")
	}
	
	if err := k.VerifySignatureOnDidUpdate(&ctx, oldDIDDoc, didMsg, msg.Signatures); err != nil {
		return nil, err
	}

	// TODO: Implement this when the version ID is used
	// if oldStateValue.Metadata.VersionId != didMsg.VersionId {
	// 	errMsg := fmt.Sprintf("Ecpected %s with version %s. Got version %s", didMsg.Id, oldStateValue.Metadata.VersionId, didMsg.VersionId)
	// 	return nil, sdkerrors.Wrap(types.ErrUnexpectedDidVersion, errMsg)
	// }

	var didSpec = types.Did{
		Context:              didMsg.GetContext(),
		Id:                   didMsg.GetId(),
		Controller:           didMsg.GetController(),
		VerificationMethod:   didMsg.GetVerificationMethod(),
		Authentication:       didMsg.GetAuthentication(),
		AssertionMethod:      didMsg.GetAssertionMethod(),
		KeyAgreement:         didMsg.GetKeyAgreement(),
		CapabilityInvocation: didMsg.GetCapabilityInvocation(),
		Created:              didMsg.GetCreated(),
		Updated:              didMsg.GetUpdated(),
	}

	didSpec.Updated = ctx.BlockTime().String()

	if err := k.SetDid(ctx, didSpec); err != nil {
		return nil, err
	}

	return &types.MsgUpdateDIDResponse{UpdateId: didSpec.Id}, nil
}
