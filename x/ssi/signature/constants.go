package signature

// Supported Verification Method Types
const Ed25519VerificationKey2020 = "Ed25519VerificationKey2020"

// Mapping between Verification Key and its corresponding Signature
var verificationKeySignatureMap = map[string]string{
	"Ed25519VerificationKey2020": "Ed25519Signature2020",
}
