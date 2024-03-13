package crypto

import (
	"fmt"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

type IKeyPair interface {
	GetType() string
	GetPublicKey() string
	GetPrivateKey() string
	GetVerificationMethodId() string
	GetOptionalID() string
}

type Ed25519KeyPair struct {
	Type       string
	PublicKey  string
	PrivateKey string
	VerificationMethodId string
	OptionalID string // If this field is not empty, it will override publicKey as the method specific id
}

func (kp *Ed25519KeyPair) GetType() string {
	return kp.Type
}

func (kp *Ed25519KeyPair) GetPublicKey() string {
	return kp.PublicKey
}

func (kp *Ed25519KeyPair) GetPrivateKey() string {
	return kp.PrivateKey
}

func (kp *Ed25519KeyPair) GetVerificationMethodId() string {
	return kp.VerificationMethodId
}

func (kp *Ed25519KeyPair) GetOptionalID() string {
	return kp.OptionalID
}

type Secp256k1Pair struct {
	Type       string
	PublicKey  string
	PrivateKey string
	VerificationMethodId string
	OptionalID string // If this field is not empty, it will override publicKey as the method specific id
}

func (kp *Secp256k1Pair) GetType() string {
	return kp.Type
}

func (kp *Secp256k1Pair) GetPublicKey() string {
	return kp.PublicKey
}

func (kp *Secp256k1Pair) GetPrivateKey() string {
	return kp.PrivateKey
}

func (kp *Secp256k1Pair) GetVerificationMethodId() string {
	return kp.VerificationMethodId
}

func (kp *Secp256k1Pair) GetOptionalID() string {
	return kp.OptionalID
}


type Secp256k1RecoveryPair struct {
	Type       string
	PublicKey  string
	PrivateKey string
	VerificationMethodId string
	OptionalID string // If this field is not empty, it will override publicKey as the method specific id
}

func (kp *Secp256k1RecoveryPair) GetType() string {
	return kp.Type
}

func (kp *Secp256k1RecoveryPair) GetPublicKey() string {
	return kp.PublicKey
}

func (kp *Secp256k1RecoveryPair) GetPrivateKey() string {
	return kp.PrivateKey
}

func (kp *Secp256k1RecoveryPair) GetVerificationMethodId() string {
	return kp.VerificationMethodId
}

func (kp *Secp256k1RecoveryPair) GetOptionalID() string {
	return kp.OptionalID
}

type BabyJubJubKeyPair struct {
	Type       string
	PublicKey  string
	PrivateKey string
	VerificationMethodId string
	OptionalID string // If this field is not empty, it will override publicKey as the method specific id
}

func (kp *BabyJubJubKeyPair) GetType() string {
	return kp.Type
}

func (kp *BabyJubJubKeyPair) GetPublicKey() string {
	return kp.PublicKey
}

func (kp *BabyJubJubKeyPair) GetPrivateKey() string {
	return kp.PrivateKey
}

func (kp *BabyJubJubKeyPair) GetVerificationMethodId() string {
	return kp.VerificationMethodId
}

func (kp *BabyJubJubKeyPair) GetOptionalID() string {
	return kp.OptionalID
}

type BbsBlsKeyPair struct {
	Type       string
	PublicKey  string
	PrivateKey string
	VerificationMethodId string
	OptionalID string // If this field is not empty, it will override publicKey as the method specific id
}

func (kp *BbsBlsKeyPair) GetType() string {
	return kp.Type
}

func (kp *BbsBlsKeyPair) GetPublicKey() string {
	return kp.PublicKey
}

func (kp *BbsBlsKeyPair) GetPrivateKey() string {
	return kp.PrivateKey
}

func (kp *BbsBlsKeyPair) GetVerificationMethodId() string {
	return kp.VerificationMethodId
}

func (kp *BbsBlsKeyPair) GetOptionalID() string {
	return kp.OptionalID
}


func CollectKeysPairs(kps ...IKeyPair) []IKeyPair {
	return kps
}

func GetSignatureTypeFromVmType(vmType string) string {
	switch vmType {
	case types.Ed25519VerificationKey2020:
		return types.Ed25519Signature2020
	case types.EcdsaSecp256k1VerificationKey2019:
		return types.EcdsaSecp256k1Signature2019
	case types.EcdsaSecp256k1RecoveryMethod2020:
		return types.EcdsaSecp256k1RecoverySignature2020
	case types.Bls12381G2Key2020:
		return types.BbsBlsSignature2020
	case types.BabyJubJubKey2021:
		return types.BJJSignature2021
	default:
		panic(fmt.Sprintf("Unsupported vm Type: %v", vmType))
	}
}