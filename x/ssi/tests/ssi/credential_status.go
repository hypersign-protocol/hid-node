package ssi

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	ldcontext "github.com/hypersign-protocol/hid-node/x/ssi/ld-context"
	testconstants "github.com/hypersign-protocol/hid-node/x/ssi/tests/constants"
	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func GenerateCredentialStatus(keyPair testcrypto.IKeyPair, issuerId string) *types.CredentialStatusDocument {
	var credentialId = "vc:" + testconstants.DidMethod + ":" + testconstants.ChainNamespace + ":" + strings.Split(issuerId, ":")[3]
	var credHash = sha256.Sum256([]byte("Hash1234"))
	var vmContextUrl = GetContextFromKeyPair(keyPair)

	var credentialStatus *types.CredentialStatusDocument = &types.CredentialStatusDocument{
		Context: []string{
			ldcontext.CredentialStatusContext,
			vmContextUrl,
		},
		Id:                       credentialId,
		Remarks:                  "Live",
		Revoked:                  false,
		Suspended:                false,
		Issuer:                   issuerId,
		IssuanceDate:             "2022-04-10T04:07:12Z",
		CredentialMerkleRootHash: hex.EncodeToString(credHash[:]),
	}
	return credentialStatus
}

func GenerateRegisterCredStatusRPCElements(keyPair testcrypto.IKeyPair, credentialStatus *types.CredentialStatusDocument, verficationMethod *types.VerificationMethod) *types.MsgRegisterCredentialStatus {
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

func GenerateUpdateCredStatusRPCElements(keyPair testcrypto.IKeyPair, credentialStatus *types.CredentialStatusDocument, verficationMethod *types.VerificationMethod) *types.MsgUpdateCredentialStatus {
	var credentialProof *types.DocumentProof = &types.DocumentProof{
		Created:            "2022-04-10T04:07:12Z",
		VerificationMethod: verficationMethod.Id,
		ProofPurpose:       "assertionMethod",
	}

	var credentialStatusSignature string = testcrypto.SignGeneric(keyPair, credentialStatus, credentialProof)
	credentialProof.ProofValue = credentialStatusSignature

	return &types.MsgUpdateCredentialStatus{
		CredentialStatusDocument: credentialStatus,
		CredentialStatusProof:    credentialProof,
		TxAuthor:                 testconstants.Creator,
	}
}
