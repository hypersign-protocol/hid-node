package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	verify "github.com/hypersign-protocol/hid-node/x/ssi/keeper/document_verification"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func (k msgServer) CreateDID(goCtx context.Context, msg *types.MsgCreateDID) (*types.MsgCreateDIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	didMsg := msg.GetDidDocString()
	did := didMsg.GetId()

	chainNamespace := k.GetChainNamespace(&ctx)

	// Checks if the DidDoc is a valid format
	err := verify.IsValidDidDoc(msg.DidDocString, chainNamespace)
	if err != nil {
		return nil, err
	}

	// Checks if the DID is already present in the store
	if k.HasDid(ctx, did) {
		return nil, sdkerrors.Wrap(types.ErrDidDocExists, fmt.Sprintf("DID already exists %s", did))
	}

	// Signature check
	if err := k.VerifySignature(&ctx, didMsg, didMsg.GetSigners(), msg.GetSignatures()); err != nil {
		return nil, err
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
		CapabilityDelegation: didMsg.GetCapabilityDelegation(),
		Service:              didMsg.GetService(),
	}

	// Create the Metadata
	metadata := types.CreateNewMetadata(ctx)

	// Form the DID Document
	didDoc := types.DidDocument{
		Did:      &didSpec,
		Metadata: &metadata,
	}
	// Add a DID to the store
	id := k.AppendDID(ctx, &didDoc)
	// Return the Id of the DID
	return &types.MsgCreateDIDResponse{Id: id}, nil
}

func (k msgServer) UpdateDID(goCtx context.Context, msg *types.MsgUpdateDID) (*types.MsgUpdateDIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	didMsg := msg.GetDidDocString()
	did := msg.GetDidDocString().GetId()
	versionId := msg.GetVersionId()

	chainNamespace := k.GetChainNamespace(&ctx)

	oldDIDDoc, err := k.GetDid(&ctx, didMsg.Id)
	if err != nil {
		return nil, err
	}
	oldDid := oldDIDDoc.GetDid()
	oldMetaData := oldDIDDoc.GetMetadata()

	// Check if the didDoc is valid
	err = verify.IsValidDidDoc(didMsg, chainNamespace)
	if err != nil {
		return nil, err
	}

	// Checks if the DID is not present in the store
	if !k.HasDid(ctx, did) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("DID doesnt exists %s", did))
	}

	// Check if the DID is already deactivated
	if err := VerifyDidDeactivate(oldMetaData, didMsg.Id); err != nil {
		return nil, err
	}

	if k.ValidateDidControllers(&ctx, did, didMsg.GetController(), didMsg.GetVerificationMethod()) != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, "DID controller is not valid")
	}

	if err := k.VerifySignatureOnDidUpdate(&ctx, oldDid, didMsg, msg.Signatures); err != nil {
		return nil, err
	}

	// Check if the versionId passed is the same as the one in the Latest DID Document in store
	if oldMetaData.VersionId != versionId {
		errMsg := fmt.Sprintf("Expected %s with version %s. Got version %s", didMsg.Id, oldMetaData.VersionId, versionId)
		return nil, sdkerrors.Wrap(types.ErrUnexpectedDidVersion, errMsg)
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
		Service:              didMsg.GetService(),
	}

	// Create the Metadata
	metadata := types.CreateNewMetadata(ctx)
	// Assign `created` and `deactivated` to previous DIDDoc's metadata values
	metadata.Created = oldDIDDoc.GetMetadata().Created
	metadata.Deactivated = oldDIDDoc.GetMetadata().Deactivated

	// Form the DID Document
	didDoc := types.DidDocument{
		Did:      &didSpec,
		Metadata: &metadata,
	}
	if err := k.SetDid(ctx, didDoc); err != nil {
		return nil, err
	}

	return &types.MsgUpdateDIDResponse{UpdateId: didSpec.Id}, nil
}

func (k msgServer) DeactivateDID(goCtx context.Context, msg *types.MsgDeactivateDID) (*types.MsgDeactivateDIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := msg.GetDidId()
	versionId := msg.GetVersionId()

	didDocument, err := k.GetDid(&ctx, id)
	if err != nil {
		return nil, err
	}
	did := didDocument.GetDid()
	metadata := didDocument.GetMetadata()

	chainNamespace := k.GetChainNamespace(&ctx)

	// Check if the didDoc is valid
	err = verify.IsValidDidDoc(did, chainNamespace)
	if err != nil {
		return nil, err
	}

	if k.ValidateDidControllers(&ctx, id, did.GetController(), did.GetVerificationMethod()) != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, "DID controller is not valid")
	}

	if err := k.VerifySignature(&ctx, did, did.GetSigners(), msg.Signatures); err != nil {
		return nil, err
	}

	// Check if the versionId passed is the same as the one in the Latest DID Document in store
	if metadata.GetVersionId() != versionId {
		errMsg := fmt.Sprintf("Expected %s with version %s. Got version %s", did.GetId(), metadata.GetVersionId(), versionId)
		return nil, sdkerrors.Wrap(types.ErrUnexpectedDidVersion, errMsg)
	}

	// Check if the DID is already deactivated
	if err := VerifyDidDeactivate(metadata, did.GetId()); err != nil {
		return nil, err
	}

	// Create updated metadata
	updatedMetadata := types.CreateNewMetadata(ctx)
	updatedMetadata.Created = didDocument.GetMetadata().Created
	updatedMetadata.Deactivated = true

	// Form the updated DID Document
	updatedDidDocument := types.DidDocument{
		Did:      did,
		Metadata: &updatedMetadata,
	}

	if err := k.SetDidDeactivate(ctx, updatedDidDocument, id); err != nil {
		return nil, err
	}

	return &types.MsgDeactivateDIDResponse{Id: 1}, nil
}
