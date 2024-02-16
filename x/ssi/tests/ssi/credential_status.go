package ssi

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	testconstants "github.com/hypersign-protocol/hid-node/x/ssi/tests/constants"
	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func GenerateCredStatusRPCElements(keyPair *testcrypto.KeyPair, issuerId string, verficationMethod *types.VerificationMethod) *types.MsgRegisterCredentialStatus {
	var credentialId = "vc:" + testconstants.DidMethod + ":" + testconstants.ChainNamespace + ":" + strings.Split(issuerId, ":")[3]
	var credHash = sha256.Sum256([]byte("Hash1234"))
	var credentialStatus *types.CredentialStatusDocument = &types.CredentialStatusDocument{
		Id:                       credentialId,
		Remarks:                  "Live",
		Revoked:                  false,
		Suspended:                false,
		Issuer:                   issuerId,
		IssuanceDate:             "2022-04-10T04:07:12Z",
		CredentialMerkleRootHash: hex.EncodeToString(credHash[:]),
	}

	var credentialProof *types.DocumentProof = &types.DocumentProof{
		Created:            "2022-04-10T04:07:12Z",
		VerificationMethod: verficationMethod.Id,
		ProofPurpose:       "assertionMethod",
	}

	var credentialStatusSignature string = testcrypto.SignGeneric(keyPair, credentialStatus, credentialProof)
	credentialProof.ProofValue = credentialStatusSignature

	return &types.MsgRegisterCredentialStatus{
		CredentialStatusDocument: credentialStatus,
		CredentialStatusProof:    credentialProof,
		TxAuthor:                 testconstants.Creator,
	}
}

func UpdateCredStatusToSuspended(
	msgUpdateCred *types.MsgUpdateCredentialStatus,
	keyPair *testcrypto.KeyPair,
) *types.MsgUpdateCredentialStatus {
	msgUpdateCred.CredentialStatusDocument.Suspended = true
	msgUpdateCred.CredentialStatusDocument.Remarks = "Status changed for Testing"

	updatedSignature := ""

	msgUpdateCred.CredentialStatusProof.ProofValue = updatedSignature

	return msgUpdateCred
}