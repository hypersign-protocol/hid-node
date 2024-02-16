package wasmbinding

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/wasmbinding/bindings"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func (qp *QueryPlugin) GetDidDocument(ctx sdk.Context, didId string) (*bindings.QueryDidDocumentResponse, error) {
	response, err := qp.ssiKeeper.DidDocumentByID(&ctx, &types.QueryDidDocumentRequest{
		DidId: didId,
	})

	if err != nil {
		return nil, types.ErrDidDocNotFound
	}

	didDoc := response.DidDocument

	return &bindings.QueryDidDocumentResponse{
		Context:              didDoc.Context,
		Id:                   didDoc.Id,
		Controller:           didDoc.Controller,
		AlsoKnownAs:          didDoc.AlsoKnownAs,
		VerificationMethod:   didDoc.VerificationMethod,
		Authentication:       didDoc.Authentication,
		AssertionMethod:      didDoc.AssertionMethod,
		KeyAgreement:         didDoc.KeyAgreement,
		CapabilityInvocation: didDoc.CapabilityInvocation,
		CapabilityDelegation: didDoc.CapabilityDelegation,
		Service:              didDoc.Service,
	}, nil
}

func (qp *QueryPlugin) GetDidIdFromAddress(ctx sdk.Context, address string) (*bindings.QueryDidDocumentFromAddressResponse, error) {
	blockchainAccountId := "cosmos:prajna:" + address 
	didIdBytes := qp.ssiKeeper.GetBlockchainAddressFromStore(&ctx, blockchainAccountId)
	if len(didIdBytes) == 0 {
		return nil, fmt.Errorf("did document doesn't exists for address: %v", address)
	}

	didId := string(didIdBytes)

	return &bindings.QueryDidDocumentFromAddressResponse{
		DidId: didId,
	}, nil
}
