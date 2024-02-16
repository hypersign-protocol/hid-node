package ssi

import "strings"

func stripDidFromVerificationMethod(vmId string) string {
	segments := strings.Split(vmId, "#")
	return segments[0]
}
