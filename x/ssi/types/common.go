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

var supportedVerificationMethodTypes []string = func() []string {
	result := []string{}

	for vmType := range VerificationKeySignatureMap {
		result = append(result, vmType)
	}

	return result
}()

// Supported Service Types
var SupportedServiceTypes = []string{
	"LinkedDomains",
}

// Did Document ID
const DocumentIdentifierDid = "did"
const DidMethod = "hid"

// CAIP-10 prefixes
const EthereumCAIP10Prefix string = "eip155" // Ethereum Based Chains
const CosmosCAIP10Prefix string = "cosmos"   // Cosmos Based Chains

// Supported CAIP-10 prefixes
var CAIP10PrefixForEcdsaSecp256k1RecoveryMethod2020 []string = []string{
	EthereumCAIP10Prefix,
}

var CAIP10PrefixForEcdsaSecp256k1VerificationKey2019 []string = []string{
	CosmosCAIP10Prefix,
}

const ADR036ClientSpec string = "cosmos-ADR036"
const PersonalSignClientSpec string = "eth-personalSign"

// Supported Client Specs
var SupportedClientSpecs []string = []string{
	ADR036ClientSpec,
	PersonalSignClientSpec,
}

// Map between supported cosmos chain-id and their respective blockhchain address prefix
var CosmosCAIP10ChainIdBech32PrefixMap = map[string]string{
	// Mainnet Chains
	"cosmoshub-4":                "cosmos",
	"osmosis-1":                  "osmo",
	"akashnet-2":                 "akash",
	"stargaze-1":                 "stars",
	"core-1":                     "persistence",
	"crypto-org-chain-mainnet-1": "cro",

	// Testnet Chains
	"theta-testnet-001": "cosmos",
	"osmo-test-4":       "osmo",
	"elgafar-1":         "stars",
	"test-core-1":       "persistence",
	"jagrat":            "hid",
}
