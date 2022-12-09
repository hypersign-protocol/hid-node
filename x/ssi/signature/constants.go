package signature

// Supported Verification Method Types
const Ed25519VerificationKey2020 = "Ed25519VerificationKey2020"
const EcdsaSecp256k1VerificationKey2019 = "EcdsaSecp256k1VerificationKey2019"

// Mapping between Verification Key and its corresponding Signature
var VerificationKeySignatureMap = map[string]string{
	"Ed25519VerificationKey2020":        "Ed25519Signature2020",
	"EcdsaSecp256k1VerificationKey2019": "EcdsaSecp256k1Signature2019",
}
