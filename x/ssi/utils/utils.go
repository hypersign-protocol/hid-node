package utils

import (
	"crypto/ed25519"
	"strings"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func FindPublicKey(signer types.Signer, id string) (ed25519.PublicKey, error) {
	for _, authentication := range signer.Authentication {
		if authentication == id {
			vm := FindVerificationMethod(signer.VerificationMethod, id)
			if vm == nil {
				return nil, types.ErrVerificationMethodNotFound.Wrap(id)
			}
			return vm.GetPublicKey()
		}
	}

	return nil, types.ErrVerificationMethodNotFound.Wrap(id)
}

func FindVerificationMethod(vms []*types.VerificationMethod, id string) *types.VerificationMethod {
	for _, vm := range vms {
		if vm.Id == id {
			return vm
		}
	}

	return nil
}

func MergeUrlWithResource(url string, resource string) string {
	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}

	if resource[0] == '/' {
		resource = resource[1:]
	}

	return url + "/" + resource
}

func SplitDidUrlIntoDid(didUrl string) (string, string) {
	segments := strings.Split(didUrl, "#")
	return segments[0], segments[1]
}
