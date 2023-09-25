package verification

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"

	ldcontext "github.com/hypersign-protocol/hid-node/x/ssi/ld-context"
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
			if didDoc, ok := ssiMsg.(*types.Did); ok {
				canonizedDidDocHash, err := ldcontext.EcdsaSecp256k1RecoverySignature2020Canonize(didDoc)
				if err != nil {
					return nil, err
				}

				// TODO: This is temporary fix eth.personal.sign() client function, since it only signs JSON 
				// stringified document and hence the following struct was used to sign from the Client end.
				return json.Marshal(struct{
					DidId string `json:"didId"`
					DidDocDigest string `json:"didDocDigest"`
				} {
					DidId: didDoc.Id,
					DidDocDigest: hex.EncodeToString(canonizedDidDocHash),
				})
			}
			return getPersonalSignSpecDocBytes(ssiMsg)
		default:
			return nil, fmt.Errorf(
				"supported clientSpecs: %v",
				types.SupportedClientSpecs,
			)
		}
	} else {
		// If DID Document, perform RDF normalisation and return its SHA-256 Hash
		didDoc, ok := ssiMsg.(*types.Did)
		if ok {
			return ldcontext.NormalizeByVerificationMethodType(didDoc, extendedVm.Type)
		}
		return ssiMsg.GetSignBytes(), nil
	}
}
