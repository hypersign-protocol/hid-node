package keeper

import (
	"encoding/base64"
	"encoding/json"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)


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

// Get the updated marshaled SSI Document for the corresponding ClientSpec
func getClientSpecDocBytes(clientSpecType string, clientSpecOpts types.ClientSpecOpts) ([]byte, error) {
	switch clientSpecType {
	case "cosmos-ADR036":
		return getCosmosADR036SignDocBytes(clientSpecOpts), nil
	// Non-ClientSpec RPC Request
	case "":
		return clientSpecOpts.SSIDocBytes, nil
	default:
		return nil, sdkerrors.Wrap(
			types.ErrInvalidClientSpecType,
			"supported client specs are : ['cosmos-ADR036']",
		)
	}
}
