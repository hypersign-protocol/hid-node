package keeper

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

var supportedClientSpecs []string = []string{
	"cosmos-ADR036",
}

// Read more about Cosmos's ADR Spec from the following:
// https://docs.cosmos.network/v0.45/architecture/adr-036-arbitrary-signature.html
func getCosmosADR036SignDocBytes(clientSpecOpts types.ClientSpecOpts) []byte {
	var msgSignData types.Msg = types.Msg{
		Type: "sign/MsgSignData",
		Value: types.Val{
			Data:   "",
			Signer: "",
		},
	}

	var baseCosmosADR036SignDoc types.SignDoc = types.SignDoc{
		AccountNumber: "0",
		ChainId:       "",
		Fee: types.Fees{
			Amount: []string{},
			Gas:    "0",
		},
		Memo: "",
		Msgs: []types.Msg{
			msgSignData,
		},
		Sequence: "0",
	}

	baseCosmosADR036SignDoc.Msgs[0].Value.Data = base64.StdEncoding.EncodeToString(
		clientSpecOpts.SSIDocBytes)
	baseCosmosADR036SignDoc.Msgs[0].Value.Signer = clientSpecOpts.SignerAddress
	
	
	updatedSignDocBytes, err := json.Marshal(baseCosmosADR036SignDoc)
	if err != nil {
		panic(err)
	}

	return updatedSignDocBytes
}

// Get the updated marshaled SSI document for the respective ClientSpec
func getClientSpecDocBytes(clientSpecType string, clientSpecOpts types.ClientSpecOpts) ([]byte, error) {
	switch clientSpecType {
	case "cosmos-ADR036":
		return getCosmosADR036SignDocBytes(clientSpecOpts), nil
	// Non-ClientSpec RPC Request
	// Return marshaled SSI document as-is 
	case "":
		return clientSpecOpts.SSIDocBytes, nil
	default:
		return nil, sdkerrors.Wrap(
			types.ErrInvalidClientSpecType,
			fmt.Sprintf(
				"supported client specs are : [%s]", 
				strings.Join(supportedClientSpecs, ", "),
			),
		)
	}
}
