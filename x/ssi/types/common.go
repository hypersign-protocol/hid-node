package types

// Supported Verification Method Types
const Ed25519VerificationKey2020 = "Ed25519VerificationKey2020"
const EcdsaSecp256k1VerificationKey2019 = "EcdsaSecp256k1VerificationKey2019"
const EcdsaSecp256k1RecoveryMethod2020 = "EcdsaSecp256k1RecoveryMethod2020"

// Mapping between Verification Key and its corresponding Signature
var VerificationKeySignatureMap = map[string]string{
	Ed25519VerificationKey2020:        "Ed25519Signature2020",
	EcdsaSecp256k1VerificationKey2019: "EcdsaSecp256k1Signature2019",
	EcdsaSecp256k1RecoveryMethod2020:  "EcdsaSecp256k1RecoverySignature2020",
}

// Supported Service Types
var SupportedServiceTypes = []string{
	"LinkedDomains",
}

// Did Document ID
const DocumentIdentifierDid = "did"
const DidMethod = "hid"

// Supported CAIP-10 prefixes
const EIP155 string = "eip155"

// Support Client Specs
const ADR036Spec string = "cosmos-ADR036"
const PersonalSignSpec string = "eth-personalSign"

var SupportedClientSpecs []string = []string{
	ADR036Spec,
	PersonalSignSpec,
}
