package ssi

import (
	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
)

type SsiDocSigningElements struct {
	KeyPair *testcrypto.KeyPair
	VmId    string
}
