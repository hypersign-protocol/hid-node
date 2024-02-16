package ssi

import (
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	testcrypto  "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
)

func getDocumentProof(ssiDoc types.SsiMsg, signingElements []*SsiDocSigningElements) []*types.DocumentProof {
	var docProofs []*types.DocumentProof

	genericDocumentProof := &types.DocumentProof{
		Created:      "2023-08-16T09:37:12Z",
		ProofPurpose: "assertionMethod",
	}

	for i := 0; i < len(signingElements); i++ {
		genericDocumentProof.VerificationMethod = signingElements[i].VmId

		signature := testcrypto.SignGeneric(signingElements[i].KeyPair, ssiDoc, genericDocumentProof)
		genericDocumentProof.ProofValue = signature

		docProofs = append(docProofs, genericDocumentProof)
	}
	return docProofs
}
