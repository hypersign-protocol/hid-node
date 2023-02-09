package verification

import (
	"strings"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/multiformats/go-multibase"
)

type verificationMethodElements struct {
	PublicKey              string
	VerificationMethodType string
	BlockchainAccountId    string
}

func FindVerificationMethod(vms []*types.VerificationMethod, id string) *types.VerificationMethod {
	for _, vm := range vms {
		if vm.Id == id {
			return vm
		}
	}
	return nil
}

func findVerificationMethodElements(signer types.Signer, vmId string) (verificationMethodElements, error) {
	if signer.Authentication != nil {
		for _, authentication := range signer.Authentication {
			if authentication == vmId {
				vm := FindVerificationMethod(signer.VerificationMethod, vmId)
				if vm == nil {
					return verificationMethodElements{}, types.ErrVerificationMethodNotFound.Wrap(vmId)
				}
				return getVerificationMethodElements(vm)
			}
		}
	}

	if signer.AssertionMethod != nil {
		for _, assertionMethod := range signer.AssertionMethod {
			if assertionMethod == vmId {
				vm := FindVerificationMethod(signer.VerificationMethod, vmId)
				if vm == nil {
					return verificationMethodElements{}, types.ErrVerificationMethodNotFound.Wrap(vmId)
				}
				return getVerificationMethodElements(vm)
			}
		}
	}
	return verificationMethodElements{}, types.ErrVerificationMethodNotFound.Wrap(vmId)
}

func splitDidUrlIntoDid(didUrl string) (string, string) {
	segments := strings.Split(didUrl, "#")
	return segments[0], segments[1]
}

func getVerificationMethodElements(v *types.VerificationMethod) (verificationMethodElements, error) {
	var publicKey string = v.PublicKeyMultibase
	var verificationMethodType string = v.Type
	var blockchainAccountId string = v.BlockchainAccountId

	var vmElements verificationMethodElements
	vmElements.VerificationMethodType = verificationMethodType

	// Check if the public key is decoded successfully
	if publicKey != "" {
		if publicKey[0] != 'z' {
			return verificationMethodElements{}, types.ErrInvalidPublicKey.Wrapf(
				"public key is expected to be in multibase base58btc encoding, recieved public key: %s",
				publicKey,
			)
		}

		_, _, err := multibase.Decode(publicKey)
		if err != nil {
			return verificationMethodElements{}, types.ErrInvalidPublicKey.Wrapf(
				"multibase decoding failed, recieved public key: %s",
				publicKey,
			)
		}

		vmElements.PublicKey = publicKey
	}

	// Check if the blockchainAccountId is of supported type
	if blockchainAccountId != "" {
		vmElements.BlockchainAccountId = blockchainAccountId
	}

	return vmElements, nil
}
