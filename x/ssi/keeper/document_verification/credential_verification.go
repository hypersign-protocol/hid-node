package verification

import (
	"regexp"
)

func VerifyCredentialHash(credHash string) bool {
	var matchFound bool = false
	var supportedCredentialHash map[string]string = map[string]string{
		"sha256": "[a-f0-9]{64}",
	}

	for _, regexPattern := range supportedCredentialHash {
		matchFound, _ = regexp.MatchString(regexPattern, credHash)
		if matchFound {
			return true
		}
	}

	return matchFound
}
