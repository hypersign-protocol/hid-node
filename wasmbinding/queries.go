package wasmbinding

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/wasmbinding/bindings"
)

func (qp *QueryPlugin) GetDidDocumentExists(ctx sdk.Context, didId string) (*bindings.QueryDidDocumentExistsResponse, error) {
	_, err := qp.ssiKeeper.GetDidDocumentState(&ctx, didId)
	
	if err != nil {
		return &bindings.QueryDidDocumentExistsResponse{
			Result: false,
		}, nil
	} else {
		return &bindings.QueryDidDocumentExistsResponse{
			Result: true,
		}, nil
	}
}