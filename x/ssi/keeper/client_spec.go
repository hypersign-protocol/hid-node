package keeper

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// Read more about Cosmos's ADR Spec from the following:
// https://docs.cosmos.network/v0.45/architecture/adr-036-arbitrary-signature.html
func getCosmosADR036SignDocBytes(clientSpecOpts types.ClientSpecOpts) ([]byte, error) {
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
	ssiDocBytes := clientSpecOpts.SSIDoc.GetSignBytes()
	baseCosmosADR036SignDoc.Msgs[0].Value.Data = base64.StdEncoding.EncodeToString(
		ssiDocBytes)
	baseCosmosADR036SignDoc.Msgs[0].Value.Signer = clientSpecOpts.SignerAddress

	updatedSignDocBytes, err := json.Marshal(baseCosmosADR036SignDoc)
	if err != nil {
		return nil, err
	}

	return updatedSignDocBytes, nil
}

// More info on the `personal_sign` here: https://docs.metamask.io/guide/signing-data.html#personal-sign
func getPersonalSignSpecDocBytes(clientSpecOpts types.ClientSpecOpts) ([]byte, error) {
	return json.Marshal(clientSpecOpts.SSIDoc)
}

// Get the updated marshaled SSI document for the respective ClientSpec
func getClientSpecDocBytes(clientSpecOpts types.ClientSpecOpts) ([]byte, error) {
	switch clientSpecOpts.ClientSpecType {
	case types.ADR036Spec:
		return getCosmosADR036SignDocBytes(clientSpecOpts)
	case types.PersonalSignSpec:
		return getPersonalSignSpecDocBytes(clientSpecOpts)
	// Non-ClientSpec RPC Request
	// Return marshaled SSI document as-is
	case "":
		return clientSpecOpts.SSIDoc.GetSignBytes(), nil
	default:
		return nil, fmt.Errorf(
			"supported client specs are : [%s]",
			strings.Join(types.SupportedClientSpecs, ", "),
		)
	}
}
