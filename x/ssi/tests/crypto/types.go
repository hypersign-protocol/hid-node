package crypto

type KeyPair struct {
	Type       string
	PublicKey  string
	PrivateKey string
	OptionalID string // If this field is not empty, it will override publicKey as the method specific id
}

const Ed25519KeyPair string = "Ed25519KeyPair"
const Secp256k1Pair string = "Secp256k1Pair"
