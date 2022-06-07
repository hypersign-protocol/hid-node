package utils

import (
	"crypto/ed25519"
	"errors"
	"fmt"
	"regexp"
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

func IsValidSchemaID(schemaId string, authorDid string) error {
	IdComponents := strings.Split(schemaId, ";")
	if len(IdComponents) < 2 {
		return errors.New("Expected 3 components in schema ID after being seperated by `;`, got " + fmt.Sprint(len(IdComponents)) + " components. The Schema ID is `" + schemaId + "` ")
	}

	//Checking the prefix
	if !strings.HasPrefix(IdComponents[0], "did:hs:") {
		return errors.New("Expected did:hs as prefix in schema ID, The Schema ID is " + schemaId)
	}

	// Check if the first component matches with author Did
	if authorDid != IdComponents[0] {
		return errors.New("author`s did doesn`t match with the first component of schema id")
	} 

	//Checking the type of version
	versionNumber := strings.Split(IdComponents[2], "=")[1]
	// TODO: The regex pattern should be configurable to match the version format.
	// Currently it's set for floating point validation
	regexPatternForVersion := regexp.MustCompile(`^(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)$`)
	if !regexPatternForVersion.MatchString(versionNumber) {
		return fmt.Errorf("input version Id -> `%s`, is an invalid format", versionNumber)
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
