package tests

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"

	"github.com/btcsuite/btcutil/base58"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/utils"
)

type ed25519KeyPair struct {
	publicKeyBase64 ed25519.PublicKey
	privateKeyBase64 ed25519.PrivateKey
	publicKeyBase58 string
}

var Creator = "hid1kxqk5ejca8nfpw8pg47484rppv359xh7qcasy4"

var didMethod = "did:hs"
// TODO: Need to look for an alternate approach for
// generating uuid as the current method works for Linux
// based systems
var didId = didMethod + ":" + utils.UUID()

var keyPair ed25519KeyPair = generatePublicPrivateKeyPair()
var verificationMethodId string = didId + "#" + keyPair.publicKeyBase58
var vm = &types.VerificationMethod{
	Id: verificationMethodId,
	Type: "Ed25519VerificationKey2020",
	Controller: didId,
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

var ValidSignInfo []*types.SignInfo = getSigningInfo(ValidDidDocumet, keyPair, vm)

func getSigningInfo(didDoc *types.Did, keyPair ed25519KeyPair, vm *types.VerificationMethod) []*types.SignInfo {
	signature := ed25519.Sign(keyPair.privateKeyBase64, didDoc.GetSignBytes())
	signInfo := &types.SignInfo{
		VerificationMethodId: vm.GetId(),
		Signature: base64.StdEncoding.EncodeToString(signature),
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
		publicKeyBase64: publicKey,
		privateKeyBase64: privateKey,
		publicKeyBase58: publicKeyBase58Encoded,
	}
}
