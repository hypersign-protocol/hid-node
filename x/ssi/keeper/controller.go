package keeper

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/utils"
)

func (k msgServer) ValidateDidControllers(ctx *sdk.Context, id string, controllers []string, verMethods []*types.VerificationMethod) error {

	for _, verificationMethod := range verMethods {
		if err := k.validateController(ctx, id, verificationMethod.Controller); err != nil {
			return err
		}
	}

	for _, didController := range controllers {
		if err := k.validateController(ctx, id, didController); err != nil {
			return err
		}
	}
	return nil
}

func (k *Keeper) validateController(ctx *sdk.Context, id string, controller string) error {
	if id == controller {
		return nil
	}
	didDoc, err := k.GetDid(ctx, controller)
	if err != nil {
		return types.ErrDidDocNotFound.Wrap(controller)
	}
	if len(didDoc.DidDocument.Authentication) == 0 {
		return types.ErrBadRequestInvalidVerMethod.Wrap(
			fmt.Sprintf("Verificatition method controller %s doesn't have an authentication keys", controller))
	}
	return nil
}

// Get Verification Method and Verification Relationship fields for Signers, if they don't have any
func (k *Keeper) GetVMForSigners(ctx *sdk.Context, signers []types.Signer) ([]types.Signer, error) {
	for i := 0; i < len(signers); i++ {
		if signers[i].VerificationMethod == nil {
			fetchedDidDoc, err := k.GetDid(ctx, signers[i].Signer)
			if err != nil {
				return nil, types.ErrDidDocNotFound.Wrap(signers[i].Signer)
			}

			signers[i].Authentication = fetchedDidDoc.DidDocument.Authentication
			signers[i].AssertionMethod = fetchedDidDoc.DidDocument.AssertionMethod
			signers[i].KeyAgreement = fetchedDidDoc.DidDocument.KeyAgreement
			signers[i].CapabilityInvocation = fetchedDidDoc.DidDocument.CapabilityInvocation
			signers[i].CapabilityDelegation = fetchedDidDoc.DidDocument.CapabilityDelegation
			signers[i].VerificationMethod = fetchedDidDoc.DidDocument.VerificationMethod
		}
	}

	return signers, nil
}

// Get the updated signers from the new DID Document
func GetUpdatedSigners(ctx *sdk.Context, oldDIDDoc *types.Did, newDIDDoc *types.Did, signatures []*types.SignInfo) []types.Signer {
	var signers []types.Signer

	oldController := oldDIDDoc.Controller
	if len(oldController) == 0 {
		oldController = []string{oldDIDDoc.Id}
	}

	for _, controller := range oldController {
		signers = append(signers, types.Signer{Signer: controller})
	}

	for _, oldVM := range oldDIDDoc.VerificationMethod {
		newVM := utils.FindVerificationMethod(newDIDDoc.VerificationMethod, oldVM.Id)

		// Verification Method has been deleted
		if newVM == nil {
			signers = appendSignerIfNeed(signers, oldVM.Controller, newDIDDoc)
		}

		// Verification Method has been changed
		if !reflect.DeepEqual(oldVM, newVM) {
			signers = appendSignerIfNeed(signers, newVM.Controller, newDIDDoc)
		}

		// Verification Method Controller has been changed, need to add old controller
		if newVM.Controller != oldVM.Controller {
			signers = appendSignerIfNeed(signers, oldVM.Controller, newDIDDoc)
		}
	}

	return signers
}

func appendSignerIfNeed(signers []types.Signer, controller string, msg *types.Did) []types.Signer {
	for _, signer := range signers {
		if signer.Signer == controller {
			return signers
		}
	}

	signer := types.Signer{
		Signer: controller,
	}

	if controller == msg.Id {
		signer.VerificationMethod = msg.VerificationMethod
		signer.Authentication = msg.Authentication
	}

	return append(signers, signer)
}
