package tests

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/utils"
)

type ed25519KeyPair struct {
	publicKeyBase64  ed25519.PublicKey
	privateKeyBase64 ed25519.PrivateKey
	publicKeyBase58  string
}

var Creator = "hid1kxqk5ejca8nfpw8pg47484rppv359xh7qcasy4"

/**************** DID Test *****************/

//var didMethod = "did:hs"
// TODO: Need to look for an alternate approach for
// generating uuid as the current method works for Linux
// based systems
var didId = "did:hs:5c3b799a-30f8-4bd8-9799-fd57d03b11c"

var keyPair ed25519KeyPair = generatePublicPrivateKeyPair()
var verificationMethodId string = didId + "#" + keyPair.publicKeyBase58
var vm = &types.VerificationMethod{
	Id:                 verificationMethodId,
	Type:               "Ed25519VerificationKey2020",
	Controller:         didId,
	PublicKeyMultibase: keyPair.publicKeyBase58,
}

var ValidDidDocumet *types.Did = &types.Did{
	Context: []string{
		"www.context.something",
	},
	Id:          didId,
	Controller:  []string{didId},
	AlsoKnownAs: []string{"some name"},
	VerificationMethod: []*types.VerificationMethod{
		vm,
	},
	Authentication: []string{verificationMethodId},
}

var DidDocumentValidSignInfo []*types.SignInfo = getDidSigningInfo(ValidDidDocumet, keyPair, vm)

/**************** Schema Test *****************/

var ValidSchemaDocument *types.Schema = &types.Schema{
	Type:         "https://w3c-ccg.github.io/vc-json-schemas/schema/1.0/schema.json",
	ModelVersion: "v1.0",
	Name:         "HS Credential",
	Author:       didId,
	Id:           fmt.Sprintf("%s;id=%s;version=1.0", didId, utils.UUID()),
	Authored:     "Tue Apr 06 2021 00:09:56 GMT+0530 (India Standard Time)",
	Schema: &types.SchemaProperty{
		Schema:               "https://json-schema.org/draft-07/schema#",
		Description:          "test",
		Type:                 "Object",
		Properties:           "{myString:{type:string}}",
		Required:             []string{"myString"},
		AdditionalProperties: false,
	},
}

var SchemaValidSignInfo []*types.SignInfo = getSchemaSigningInfo(ValidSchemaDocument, keyPair, vm)

/**********************************************/

func getDidSigningInfo(didDoc *types.Did, keyPair ed25519KeyPair, vm *types.VerificationMethod) []*types.SignInfo {
	signature := ed25519.Sign(keyPair.privateKeyBase64, didDoc.GetSignBytes())
	signInfo := &types.SignInfo{
		VerificationMethodId: vm.GetId(),
		Signature:            base64.StdEncoding.EncodeToString(signature),
	}

	return []*types.SignInfo{
		signInfo,
	}
}

func getSchemaSigningInfo(schemaDoc *types.Schema, keyPair ed25519KeyPair, vm *types.VerificationMethod) []*types.SignInfo {
	signature := ed25519.Sign(keyPair.privateKeyBase64, schemaDoc.GetSignBytes())
	signInfo := &types.SignInfo{
		VerificationMethodId: vm.GetId(),
		Signature:            base64.StdEncoding.EncodeToString(signature),
	}

	return []*types.SignInfo{
		signInfo,
	}
}

func generatePublicPrivateKeyPair() ed25519KeyPair {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	publicKeyBase58Encoded := "z" + base58.Encode(publicKey)

	return ed25519KeyPair{
		publicKeyBase64:  publicKey,
		privateKeyBase64: privateKey,
		publicKeyBase58:  publicKeyBase58Encoded,
	}
}
