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
		case contractQuery.DidDocumentExists != nil:
			didId := contractQuery.DidDocumentExists.DidId

			response, err := qp.GetDidDocumentExists(ctx, didId)
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
