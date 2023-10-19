package types

// Supported method-specific-id Formats
const MSIBlockchainAccountId = "MSIBlockchainAccountId"
const MSINonBlockchainAccountId = "MSINonBlockchainAccountId"

// Valid Proof Purpose values


// Supported Verification Method Types
const Ed25519VerificationKey2020 = "Ed25519VerificationKey2020"
const EcdsaSecp256k1VerificationKey2019 = "EcdsaSecp256k1VerificationKey2019"
const EcdsaSecp256k1RecoveryMethod2020 = "EcdsaSecp256k1RecoveryMethod2020"
const X25519KeyAgreementKey2020 = "X25519KeyAgreementKey2020"
const X25519KeyAgreementKeyEIP5630 = "X25519KeyAgreementKeyEIP5630" // TODO: Temporary spec name for KeyAgreement type from Metamask
const Bls12381G2Key2020 = "Bls12381G2Key2020"
const BabyJubJubVerificationKey2023 = "BabyJubJubVerificationKey2023"

// Supported Proof Types
const Ed25519Signature2020 = "Ed25519Signature2020"
const EcdsaSecp256k1Signature2019 = "EcdsaSecp256k1Signature2019"
const EcdsaSecp256k1RecoverySignature2020 = "EcdsaSecp256k1RecoverySignature2020"
const BabyJubJubSignature2023 = "BabyJubJubSignature2023"
const BbsBlsSignature2020 = "BbsBlsSignature2020"

// Mapping between Verification Key and its corresponding Signature
var VerificationKeySignatureMap = map[string]string{
	Ed25519VerificationKey2020:        Ed25519Signature2020,
	EcdsaSecp256k1VerificationKey2019: EcdsaSecp256k1Signature2019,
	EcdsaSecp256k1RecoveryMethod2020:  EcdsaSecp256k1RecoverySignature2020,
	X25519KeyAgreementKey2020:         "", // Authentication and Assertion are not allowed
	X25519KeyAgreementKeyEIP5630:      "", // Authentication and Assertion are not allowed
	BabyJubJubVerificationKey2023:     BabyJubJubSignature2023,
	Bls12381G2Key2020:                 BbsBlsSignature2020,
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

// Supported CAIP-10 Prefixes
var SupportedCAIP10Prefixes = []string{
	EthereumCAIP10Prefix,
	CosmosCAIP10Prefix,
}

var SupportedCAIP10EthereumChainIds = []string{
	// Ethereum-Based Mainnet Chains
	"1",   // Ethereum Mainnet
	"137", // Polygon Mainnet
	"56",  // Binance Smart Chain

	// Ethereum-Based Testnet Chains
	"3",     // Ropsten (Ethereum Testnet)
	"4",     // Rinkeby (Ethereum Testnet)
	"5",     // Goerli (Ethereum Testnet)
	"80001", // Polygon Mumbai Testnet
	"97",    // Binance Smart Chain Testnet
}

var SupportedCAIP10CosmosChainIds = []string{
	"cosmoshub-4",                // Cosmos Hub
	"osmosis-1",                  // Osmosis
	"akashnet-2",                 // Akash
	"stargaze-1",                 // Stargaze
	"core-1",                     // Persistence
	"crypto-org-chain-mainnet-1", // Crypto.Org Chain

	"theta-testnet-001", // Cosmos Hub Theta Testnet
	"osmo-test-4",       // Osmosis Testnet
	"elgafar-1",         // Stargaze Testnet
	"test-core-1",       // Persistence Testnet
	"jagrat",            // Hypersign Identity Network - Jagrat Testnet
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

// Map between support CAIP-10 prefix and list of chain-ids
var SupportedCAIP10PrefixChainIdsMap = map[string][]string{
	EthereumCAIP10Prefix: SupportedCAIP10EthereumChainIds,
	CosmosCAIP10Prefix:   SupportedCAIP10CosmosChainIds,
}
