package wasmbinding

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/wasmbinding/bindings"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func CustomQuerier(qp *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var contractQuery bindings.SsiContractQuery
		if err := json.Unmarshal(request, &contractQuery); err != nil {
			return nil, err
		}

		switch {
		case contractQuery.DidDocument != nil:
			didId := contractQuery.DidDocument.DidId

			response, err := qp.GetDidDocument(ctx, didId)
			if err != nil {
				return nil, err
			}

			resultBytes, err := json.Marshal(response)
			if err != nil {
				return nil, err
			}

			return resultBytes, nil
		case contractQuery.DidDocumentFromAddress != nil:
			addr := contractQuery.DidDocumentFromAddress.Address
			
			if _, err :=sdk.AccAddressFromBech32(addr); err != nil {
				return nil, err
			}
			
			response, err := qp.GetDidIdFromAddress(ctx, addr)
			if err != nil {
				return nil, err
			}

			resultBytes, err := json.Marshal(response)
			if err != nil {
				return nil, err
			}

			return resultBytes, nil
		default:
			return nil, sdkerrors.Wrapf(types.ErrUnsupportedWasmRequest, string(request))
		}
	}
}
