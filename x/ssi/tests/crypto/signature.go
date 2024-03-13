package crypto

import (
	ldcontext "github.com/hypersign-protocol/hid-node/x/ssi/ld-context"
	cli "github.com/hypersign-protocol/hid-node/x/ssi/client/cli"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func SignGeneric(keyPair IKeyPair, doc types.SsiMsg, docProof *types.DocumentProof) string {
	docProof.Type = GetSignatureTypeFromVmType(keyPair.GetType())

	signature, err := GetDocumentSignature(doc, docProof, keyPair.GetPrivateKey())
	if err != nil {
		panic(err)
	}
	return signature
}

func GetPublicKeyAndOptionalID(keyPair IKeyPair) (string, string) {
	return keyPair.GetPublicKey(), keyPair.GetOptionalID()
}

func GetDocumentSignature(doc types.SsiMsg, docProof *types.DocumentProof, privateKey string) (string, error) {
	var signature string

	switch docProof.Type {
	case types.Ed25519Signature2020:
		var docBytes []byte
		docBytes, err := ldcontext.Ed25519Signature2020Normalize(doc, docProof)
		if err != nil {
			return "", err
		}

		signature, err = cli.GetEd25519Signature2020(privateKey, docBytes)
		if err != nil {
			return "", err
		}
	case types.EcdsaSecp256k1Signature2019:
		var docBytes []byte
		docBytes, err := ldcontext.EcdsaSecp256k1Signature2019Normalize(doc, docProof)
		if err != nil {
			return "", err
		}

		signature, err = cli.GetEcdsaSecp256k1Signature2019(privateKey, docBytes)
		if err != nil {
			return "", err
		}
	case types.EcdsaSecp256k1RecoverySignature2020:
		var docBytes []byte
		docBytes, err := ldcontext.EcdsaSecp256k1RecoverySignature2020Normalize(doc, docProof)
		if err != nil {
			return "", err
		}

		signature, err = cli.GetEcdsaSecp256k1RecoverySignature2020(privateKey, docBytes)
		if err != nil {
			return "", err
		}
	case types.BbsBlsSignature2020:
		var docBytes []byte
		docBytes, err := ldcontext.BbsBlsSignature2020Normalize(doc, docProof)
		if err != nil {
			return "", err
		}

		signature, err = cli.GetBbsBlsSignature2020(privateKey, docBytes)
		if err != nil {
			return "", err
		}
	case types.BJJSignature2021:
		var docBytes []byte
		docBytes, err := ldcontext.BJJSignature2021Normalize(doc)
		if err != nil {
			return "", err
		}

		signature, err = cli.GetBJJSignature2021(privateKey, docBytes)
		if err != nil {
			return "", err
		}
	default:
		panic("recieved unsupported signing-algo. Supported algorithms are: [Ed25519Signature2020, EcdsaSecp256k1Signature2019, EcdsaSecp256k1RecoverySignature2020, BbsBlsSignature2020, BJJSignature2021]")
	}

	return signature, nil
}
