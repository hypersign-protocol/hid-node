package ssi

import (
	ldcontext "github.com/hypersign-protocol/hid-node/x/ssi/ld-context"
	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
)

func GetContextFromKeyPair(kp testcrypto.IKeyPair) string {
	switch kp.(type) {
	case *testcrypto.Ed25519KeyPair:
		return ldcontext.Ed25519Context2020
	case *testcrypto.Secp256k1Pair:
		return ldcontext.Secp256k12019Context
	case *testcrypto.Secp256k1RecoveryPair:
		return ldcontext.Secp256k1Recovery2020Context
	case *testcrypto.BabyJubJubKeyPair:
		return ldcontext.BabyJubJubKey2021Context
	case *testcrypto.BbsBlsKeyPair:
		return ldcontext.BbsSignature2020Context
	default:
		panic("Unsupported IKeyPair type")
	}
}
