package ssi

import (
	ldcontext "github.com/hypersign-protocol/hid-node/x/ssi/ld-context"
	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
)

func GetContextFromKeyPair(kp testcrypto.IKeyPair) []string {
	switch kp.(type) {
	case *testcrypto.Ed25519KeyPair:
		return []string{ldcontext.Ed25519Context2020}
	case *testcrypto.Secp256k1Pair:
		return []string{ldcontext.Secp256k12019Context}
	case *testcrypto.Secp256k1RecoveryPair:
		return []string{ldcontext.Secp256k1Recovery2020Context}
	case *testcrypto.BabyJubJubKeyPair:
		return []string{ldcontext.BabyJubJubKey2021Context, ldcontext.BJJSignature2021Context}
	case *testcrypto.BbsBlsKeyPair:
		return []string{ldcontext.BbsSignature2020Context}
	default:
		panic("Unsupported IKeyPair type")
	}
}
