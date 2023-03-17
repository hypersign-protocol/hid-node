package verification

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// Read more about Cosmos's ADR Spec from the following:
// https://docs.cosmos.network/v0.45/architecture/adr-036-arbitrary-signature.html
func getCosmosADR036SignDocBytes(ssiMsg types.SsiMsg, clientSpec *types.ClientSpec) ([]byte, error) {
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
	ssiDocBytes := ssiMsg.GetSignBytes()
	baseCosmosADR036SignDoc.Msgs[0].Value.Data = base64.StdEncoding.EncodeToString(
		ssiDocBytes)
	baseCosmosADR036SignDoc.Msgs[0].Value.Signer = clientSpec.Adr036SignerAddress

	updatedSignDocBytes, err := json.Marshal(baseCosmosADR036SignDoc)
	if err != nil {
		return nil, err
	}

	return updatedSignDocBytes, nil
}

// More info on the `personal_sign` here: https://docs.metamask.io/guide/signing-data.html#personal-sign
func getPersonalSignSpecDocBytes(ssiMsg types.SsiMsg) ([]byte, error) {
	return json.Marshal(ssiMsg)
}

// Get the updated marshaled SSI document for the respective ClientSpec
func getDocBytesByClientSpec(ssiMsg types.SsiMsg, extendedVm *types.ExtendedVerificationMethod) ([]byte, error) {
	if extendedVm.ClientSpec != nil {
		switch extendedVm.ClientSpec.Type {
		case types.ADR036ClientSpec:
			return getCosmosADR036SignDocBytes(ssiMsg, extendedVm.ClientSpec)
		case types.PersonalSignClientSpec:
			return getPersonalSignSpecDocBytes(ssiMsg)
		default:
			return nil, fmt.Errorf(
				"supported clientSpecs: %v",
				types.SupportedClientSpecs,
			)
		}
	} else {
		return ssiMsg.GetSignBytes(), nil
	}
}
