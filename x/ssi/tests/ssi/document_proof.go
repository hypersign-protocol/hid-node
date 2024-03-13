package ssi

import (
	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func getDocumentProof(ssiDoc types.SsiMsg, keyPairs []testcrypto.IKeyPair) []*types.DocumentProof {
	var docProofs []*types.DocumentProof
	
	for i := 0; i < len(keyPairs); i++ {
		var genericDocumentProof *types.DocumentProof = &types.DocumentProof{
			Type: testcrypto.GetSignatureTypeFromVmType(keyPairs[i].GetType()),
			Created:      "2023-08-16T09:37:12Z",
			ProofPurpose: "assertionMethod",
		}
		genericDocumentProof.VerificationMethod = keyPairs[i].GetVerificationMethodId()

		signature := testcrypto.SignGeneric(keyPairs[i], ssiDoc, genericDocumentProof)
		genericDocumentProof.ProofValue = signature

		docProofs = append(docProofs, genericDocumentProof)
	}
	return docProofs
}
