package ldcontext

const DidContext string = "https://www.w3.org/ns/did/v1"
const Ed25519Context2020 string = "https://w3id.org/security/suites/ed25519-2020/v1"
const X25519KeyAgreement2020Context string = "https://ns.did.ai/suites/x25519-2020/v1"
const Secp256k1Recovery2020Context string = "https://ns.did.ai/suites/secp256k1-2020/v1"
const BbsSignature2020Context string = "https://ns.did.ai/suites/bls12381-2020/v1"
const Secp256k12019Context string = "https://ns.did.ai/suites/secp256k1-2019/v1"
const X25519KeyAgreementKeyEIP5630Context string = "https://raw.githubusercontent.com/hypersign-protocol/hypersign-contexts/main/X25519KeyAgreementKeyEIP5630.jsonld"
const CredentialStatusContext string = "https://raw.githubusercontent.com/hypersign-protocol/hypersign-contexts/main/CredentialStatus.jsonld"
const CredentialSchemaContext string = "https://raw.githubusercontent.com/hypersign-protocol/hypersign-contexts/main/CredentialSchema.jsonld"

// As hid-node is not supposed to perform any GET request, the complete Context body of their
// respective Context urls has been maintained below.
var ContextUrlMap map[string]contextObject = map[string]contextObject{
	DidContext: {
		"@protected": true,
		"id":         "@id",
		"type":       "@type",
		"alsoKnownAs": map[string]interface{}{
			"@id":   "https://www.w3.org/ns/activitystreams#alsoKnownAs",
			"@type": "@id",
		},
		"assertionMethod": map[string]interface{}{
			"@id":        "https://w3id.org/security#assertionMethod",
			"@type":      "@id",
			"@container": "@set",
		},
		"authentication": map[string]interface{}{
			"@id":        "https://w3id.org/security#authenticationMethod",
			"@type":      "@id",
			"@container": "@set",
		},
		"capabilityDelegation": map[string]interface{}{
			"@id":        "https://w3id.org/security#capabilityDelegationMethod",
			"@type":      "@id",
			"@container": "@set",
		},
		"capabilityInvocation": map[string]interface{}{
			"@id":        "https://w3id.org/security#capabilityInvocationMethod",
			"@type":      "@id",
			"@container": "@set",
		},
		"controller": map[string]interface{}{
			"@id":   "https://w3id.org/security#controller",
			"@type": "@id",
		},
		"keyAgreement": map[string]interface{}{
			"@id":        "https://w3id.org/security#keyAgreementMethod",
			"@type":      "@id",
			"@container": "@set",
		},
		"service": map[string]interface{}{
			"@id":   "https://www.w3.org/ns/did#service",
			"@type": "@id",
			"@context": map[string]interface{}{
				"@protected": true,
				"id":         "@id",
				"type":       "@type",
				"serviceEndpoint": map[string]interface{}{
					"@id":   "https://www.w3.org/ns/did#serviceEndpoint",
					"@type": "@id",
				},
			},
		},
		"verificationMethod": map[string]interface{}{
			"@id":   "https://w3id.org/security#verificationMethod",
			"@type": "@id",
		},
	},
	Ed25519Context2020: {
		"id":         "@id",
		"type":       "@type",
		"@protected": true,
		"proof": map[string]interface{}{
			"@id":        "https://w3id.org/security#proof",
			"@type":      "@id",
			"@container": "@graph",
		},
		"Ed25519VerificationKey2020": map[string]interface{}{
			"@id": "https://w3id.org/security#Ed25519VerificationKey2020",
			"@context": map[string]interface{}{
				"@protected": true,
				"id":         "@id",
				"type":       "@type",
				"controller": map[string]interface{}{
					"@id":   "https://w3id.org/security#controller",
					"@type": "@id",
				},
				"revoked": map[string]interface{}{
					"@id":   "https://w3id.org/security#revoked",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"publicKeyMultibase": map[string]interface{}{
					"@id":   "https://w3id.org/security#publicKeyMultibase",
					"@type": "https://w3id.org/security#multibase",
				},
			},
		},
		"Ed25519Signature2020": map[string]interface{}{
			"@id": "https://w3id.org/security#Ed25519Signature2020",
			"@context": map[string]interface{}{
				"@protected": true,
				"id":         "@id",
				"type":       "@type",
				"challenge":  "https://w3id.org/security#challenge",
				"created": map[string]interface{}{
					"@id":   "http://purl.org/dc/terms/created",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"domain": "https://w3id.org/security#domain",
				"expires": map[string]interface{}{
					"@id":   "https://w3id.org/security#expiration",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"nonce": "https://w3id.org/security#nonce",
				"proofPurpose": map[string]interface{}{
					"@id":   "https://w3id.org/security#proofPurpose",
					"@type": "@vocab",
					"@context": map[string]interface{}{
						"@protected": true,
						"id":         "@id",
						"type":       "@type",
						"assertionMethod": map[string]interface{}{
							"@id":        "https://w3id.org/security#assertionMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"authentication": map[string]interface{}{
							"@id":        "https://w3id.org/security#authenticationMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"capabilityInvocation": map[string]interface{}{
							"@id":        "https://w3id.org/security#capabilityInvocationMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"capabilityDelegation": map[string]interface{}{
							"@id":        "https://w3id.org/security#capabilityDelegationMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"keyAgreement": map[string]interface{}{
							"@id":        "https://w3id.org/security#keyAgreementMethod",
							"@type":      "@id",
							"@container": "@set",
						},
					},
				},
				"proofValue": map[string]interface{}{
					"@id":   "https://w3id.org/security#proofValue",
					"@type": "https://w3id.org/security#multibase",
				},
				"verificationMethod": map[string]interface{}{
					"@id":   "https://w3id.org/security#verificationMethod",
					"@type": "@id",
				},
			},
		},
	},
	X25519KeyAgreement2020Context: {
		"id":         "@id",
		"type":       "@type",
		"@protected": true,
		"X25519KeyAgreementKey2020": map[string]interface{}{
			"@id": "https://w3id.org/security#X25519KeyAgreementKey2020",
			"@context": map[string]interface{}{
				"@protected": true,
				"id":         "@id",
				"type":       "@type",
				"controller": map[string]interface{}{
					"@id":   "https://w3id.org/security#controller",
					"@type": "@id",
				},
				"revoked": map[string]interface{}{
					"@id":   "https://w3id.org/security#revoked",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"publicKeyMultibase": map[string]interface{}{
					"@id":   "https://w3id.org/security#publicKeyMultibase",
					"@type": "https://w3id.org/security#multibase",
				},
			},
		},
	},
	Secp256k1Recovery2020Context: {
		"id":         "@id",
		"type":       "@type",
		"@protected": true,
		"EcdsaSecp256k1VerificationKey2020": map[string]interface{}{
			"@id": "https://w3id.org/security#EcdsaSecp256k1VerificationKey2020",
			"@context": map[string]interface{}{
				"@protected": true,
				"id":         "@id",
				"type":       "@type",
				"controller": map[string]interface{}{
					"@id":   "https://w3id.org/security#controller",
					"@type": "@id",
				},
				"revoked": map[string]interface{}{
					"@id":   "https://w3id.org/security#revoked",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"publicKeyMultibase": map[string]interface{}{
					"@id":   "https://w3id.org/security#publicKeyMultibase",
					"@type": "https://w3id.org/security#multibase",
				},
			},
		},
		"EcdsaSecp256k1Signature2020": map[string]interface{}{
			"@id": "https://w3id.org/security#EcdsaSecp256k1Signature2020",
			"@context": map[string]interface{}{
				"@protected": true,
				"id":         "@id",
				"type":       "@type",
				"challenge":  "https://w3id.org/security#challenge",
				"created": map[string]interface{}{
					"@id":   "http://purl.org/dc/terms/created",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"domain": "https://w3id.org/security#domain",
				"expires": map[string]interface{}{
					"@id":   "https://w3id.org/security#expiration",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"nonce": "https://w3id.org/security#nonce",
				"proofPurpose": map[string]interface{}{
					"@id":   "https://w3id.org/security#proofPurpose",
					"@type": "@vocab",
					"@context": map[string]interface{}{
						"@protected": true,
						"id":         "@id",
						"type":       "@type",
						"assertionMethod": map[string]interface{}{
							"@id":        "https://w3id.org/security#assertionMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"authentication": map[string]interface{}{
							"@id":        "https://w3id.org/security#authenticationMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"capabilityInvocation": map[string]interface{}{
							"@id":        "https://w3id.org/security#capabilityInvocationMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"capabilityDelegation": map[string]interface{}{
							"@id":        "https://w3id.org/security#capabilityDelegationMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"keyAgreement": map[string]interface{}{
							"@id":        "https://w3id.org/security#keyAgreementMethod",
							"@type":      "@id",
							"@container": "@set",
						},
					},
				},
				"jws": map[string]interface{}{
					"@id": "https://w3id.org/security#jws",
				},
				"verificationMethod": map[string]interface{}{
					"@id":   "https://w3id.org/security#verificationMethod",
					"@type": "@id",
				},
			},
		},
		"EcdsaSecp256k1RecoveryMethod2020": map[string]interface{}{
			"@id": "https://w3id.org/security#EcdsaSecp256k1RecoveryMethod2020",
			"@context": map[string]interface{}{
				"@protected": true,
				"id":         "@id",
				"type":       "@type",
				"controller": map[string]interface{}{
					"@id":   "https://w3id.org/security#controller",
					"@type": "@id",
				},
				"revoked": map[string]interface{}{
					"@id":   "https://w3id.org/security#revoked",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"blockchainAccountId": map[string]interface{}{
					"@id": "https://w3id.org/security#blockchainAccountId",
				},
				"ethereumAddress": map[string]interface{}{
					"@id": "https://w3id.org/security#ethereumAddress",
				},
				"publicKeyJwk": map[string]interface{}{
					"@id":   "https://w3id.org/security#publicKeyJwk",
					"@type": "@json",
				},
				"publicKeyBase58": map[string]interface{}{
					"@id": "https://w3id.org/security#publicKeyBase58",
				},
				"publicKeyHex": map[string]interface{}{
					"@id": "https://w3id.org/security#publicKeyHex",
				},
				"publicKeyMultibase": map[string]interface{}{
					"@id":   "https://w3id.org/security#publicKeyMultibase",
					"@type": "https://w3id.org/security#multibase",
				},
			},
		},
		"EcdsaSecp256k1RecoverySignature2020": map[string]interface{}{
			"@id": "https://w3id.org/security#EcdsaSecp256k1RecoverySignature2020",
			"@context": map[string]interface{}{
				"@protected": true,
				"id":         "@id",
				"type":       "@type",
				"challenge":  "https://w3id.org/security#challenge",
				"created": map[string]interface{}{
					"@id":   "http://purl.org/dc/terms/created",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"domain": "https://w3id.org/security#domain",
				"expires": map[string]interface{}{
					"@id":   "https://w3id.org/security#expiration",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"nonce": "https://w3id.org/security#nonce",
				"proofPurpose": map[string]interface{}{
					"@id":   "https://w3id.org/security#proofPurpose",
					"@type": "@vocab",
					"@context": map[string]interface{}{
						"@protected": true,
						"id":         "@id",
						"type":       "@type",
						"assertionMethod": map[string]interface{}{
							"@id":        "https://w3id.org/security#assertionMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"authentication": map[string]interface{}{
							"@id":        "https://w3id.org/security#authenticationMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"capabilityInvocation": map[string]interface{}{
							"@id":        "https://w3id.org/security#capabilityInvocationMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"capabilityDelegation": map[string]interface{}{
							"@id":        "https://w3id.org/security#capabilityDelegationMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"keyAgreement": map[string]interface{}{
							"@id":        "https://w3id.org/security#keyAgreementMethod",
							"@type":      "@id",
							"@container": "@set",
						},
					},
				},
				"jws": map[string]interface{}{
					"@id": "https://w3id.org/security#jws",
				},
				"verificationMethod": map[string]interface{}{
					"@id":   "https://w3id.org/security#verificationMethod",
					"@type": "@id",
				},
			},
		},
	},
	BbsSignature2020Context: {
		"@version": 1.1,
		"id":       "@id",
		"type":     "@type",
		"proof": map[string]interface{}{
			"@id":        "https://w3id.org/security#proof",
			"@type":      "@id",
			"@container": "@graph",
		},
		"BbsBlsSignature2020": map[string]interface{}{
			"@id": "https://w3id.org/security#BbsBlsSignature2020",
			"@context": map[string]interface{}{
				"@version":   1.1,
				"@protected": true,
				"id":         "@id",
				"type":       "@type",
				"challenge":  "https://w3id.org/security#challenge",
				"created": map[string]interface{}{
					"@id":   "http://purl.org/dc/terms/created",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"domain":     "https://w3id.org/security#domain",
				"proofValue": "https://w3id.org/security#proofValue",
				"nonce":      "https://w3id.org/security#nonce",
				"proofPurpose": map[string]interface{}{
					"@id":   "https://w3id.org/security#proofPurpose",
					"@type": "@vocab",
					"@context": map[string]interface{}{
						"@version":   1.1,
						"@protected": true,
						"id":         "@id",
						"type":       "@type",
						"assertionMethod": map[string]interface{}{
							"@id":        "https://w3id.org/security#assertionMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"authentication": map[string]interface{}{
							"@id":        "https://w3id.org/security#authenticationMethod",
							"@type":      "@id",
							"@container": "@set",
						},
					},
				},
				"verificationMethod": map[string]interface{}{
					"@id":   "https://w3id.org/security#verificationMethod",
					"@type": "@id",
				},
			},
		},
		"BbsBlsSignatureProof2020": map[string]interface{}{
			"@id": "https://w3id.org/security#BbsBlsSignatureProof2020",
			"@context": map[string]interface{}{
				"@version":   1.1,
				"@protected": true,
				"id":         "@id",
				"type":       "@type",
				"challenge":  "https://w3id.org/security#challenge",
				"created": map[string]interface{}{
					"@id":   "http://purl.org/dc/terms/created",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"domain": "https://w3id.org/security#domain",
				"nonce":  "https://w3id.org/security#nonce",
				"proofPurpose": map[string]interface{}{
					"@id":   "https://w3id.org/security#proofPurpose",
					"@type": "@vocab",
					"@context": map[string]interface{}{
						"@version":   1.1,
						"@protected": true,
						"id":         "@id",
						"type":       "@type",
						"sec":        "https://w3id.org/security#",
						"assertionMethod": map[string]interface{}{
							"@id":        "https://w3id.org/security#assertionMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"authentication": map[string]interface{}{
							"@id":        "https://w3id.org/security#authenticationMethod",
							"@type":      "@id",
							"@container": "@set",
						},
					},
				},
				"proofValue": "https://w3id.org/security#proofValue",
				"verificationMethod": map[string]interface{}{
					"@id":   "https://w3id.org/security#verificationMethod",
					"@type": "@id",
				},
			},
		},
		"Bls12381G1Key2020": map[string]interface{}{
			"@id": "https://w3id.org/security#Bls12381G1Key2020",
			"@context": map[string]interface{}{
				"@protected": true,
				"id":         "@id",
				"type":       "@type",
				"controller": map[string]interface{}{
					"@id":   "https://w3id.org/security#controller",
					"@type": "@id",
				},
				"revoked": map[string]interface{}{
					"@id":   "https://w3id.org/security#revoked",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"publicKeyJwk": map[string]interface{}{
					"@id":   "https://w3id.org/security#publicKeyJwk",
					"@type": "@json",
				},
				"publicKeyBase58": map[string]interface{}{
					"@id": "https://w3id.org/security#publicKeyBase58",
				},
				"publicKeyMultibase": map[string]interface{}{
					"@id":   "https://w3id.org/security#publicKeyMultibase",
					"@type": "https://w3id.org/security#multibase",
				},
			},
		},
		"Bls12381G2Key2020": map[string]interface{}{
			"@id": "https://w3id.org/security#Bls12381G2Key2020",
			"@context": map[string]interface{}{
				"@protected": true,
				"id":         "@id",
				"type":       "@type",
				"controller": map[string]interface{}{
					"@id":   "https://w3id.org/security#controller",
					"@type": "@id",
				},
				"revoked": map[string]interface{}{
					"@id":   "https://w3id.org/security#revoked",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"publicKeyJwk": map[string]interface{}{
					"@id":   "https://w3id.org/security#publicKeyJwk",
					"@type": "@json",
				},
				"publicKeyBase58": map[string]interface{}{
					"@id": "https://w3id.org/security#publicKeyBase58",
				},
				"publicKeyMultibase": map[string]interface{}{
					"@id":   "https://w3id.org/security#publicKeyMultibase",
					"@type": "https://w3id.org/security#multibase",
				},
			},
		},
	},
	Secp256k12019Context: {
		"id":         "@id",
		"type":       "@type",
		"@protected": true,
		"proof": map[string]interface{}{
			"@id":        "https://w3id.org/security#proof",
			"@type":      "@id",
			"@container": "@graph",
		},
		"EcdsaSecp256k1VerificationKey2019": map[string]interface{}{
			"@id": "https://w3id.org/security#EcdsaSecp256k1VerificationKey2019",
			"@context": map[string]interface{}{
				"@protected": true,
				"id":         "@id",
				"type":       "@type",
				"controller": map[string]interface{}{
					"@id":   "https://w3id.org/security#controller",
					"@type": "@id",
				},
				"revoked": map[string]interface{}{
					"@id":   "https://w3id.org/security#revoked",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"blockchainAccountId": map[string]interface{}{
					"@id": "https://w3id.org/security#blockchainAccountId",
				},
				"publicKeyJwk": map[string]interface{}{
					"@id":   "https://w3id.org/security#publicKeyJwk",
					"@type": "@json",
				},
				"publicKeyBase58": map[string]interface{}{
					"@id": "https://w3id.org/security#publicKeyBase58",
				},
				"publicKeyMultibase": map[string]interface{}{
					"@id":   "https://w3id.org/security#publicKeyMultibase",
					"@type": "https://w3id.org/security#multibase",
				},
			},
		},
		"EcdsaSecp256k1Signature2019": map[string]interface{}{
			"@id": "https://w3id.org/security#EcdsaSecp256k1Signature2019",
			"@context": map[string]interface{}{
				"@protected": true,
				"id":         "@id",
				"type":       "@type",
				"challenge":  "https://w3id.org/security#challenge",
				"created": map[string]interface{}{
					"@id":   "http://purl.org/dc/terms/created",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"domain": "https://w3id.org/security#domain",
				"expires": map[string]interface{}{
					"@id":   "https://w3id.org/security#expiration",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"nonce": "https://w3id.org/security#nonce",
				"proofPurpose": map[string]interface{}{
					"@id":   "https://w3id.org/security#proofPurpose",
					"@type": "@vocab",
					"@context": map[string]interface{}{
						"@protected": true,
						"id":         "@id",
						"type":       "@type",
						"assertionMethod": map[string]interface{}{
							"@id":        "https://w3id.org/security#assertionMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"authentication": map[string]interface{}{
							"@id":        "https://w3id.org/security#authenticationMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"capabilityInvocation": map[string]interface{}{
							"@id":        "https://w3id.org/security#capabilityInvocationMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"capabilityDelegation": map[string]interface{}{
							"@id":        "https://w3id.org/security#capabilityDelegationMethod",
							"@type":      "@id",
							"@container": "@set",
						},
						"keyAgreement": map[string]interface{}{
							"@id":        "https://w3id.org/security#keyAgreementMethod",
							"@type":      "@id",
							"@container": "@set",
						},
					},
				},
				"jws": map[string]interface{}{
					"@id": "https://w3id.org/security#jws",
				},
				"verificationMethod": map[string]interface{}{
					"@id":   "https://w3id.org/security#verificationMethod",
					"@type": "@id",
				},
			},
		},
	},
	X25519KeyAgreementKeyEIP5630Context: {
		"id":         "@id",
		"type":       "@type",
		"@protected": true,
		"proof": map[string]interface{}{
			"@id":        "https://w3id.org/security#proof",
			"@type":      "@id",
			"@container": "@graph",
		},
		"X25519KeyAgreementKeyEIP5630": map[string]interface{}{
			"@id": "https://w3id.org/security#X25519KeyAgreementKeyEIP5630",
			"@context": map[string]interface{}{
				"@protected": true,
				"id":         "@id",
				"type":       "@type",
				"controller": map[string]interface{}{
					"@id":   "https://w3id.org/security#controller",
					"@type": "@id",
				},
				"revoked": map[string]interface{}{
					"@id":   "https://w3id.org/security#revoked",
					"@type": "http://www.w3.org/2001/XMLSchema#dateTime",
				},
				"publicKeyMultibase": map[string]interface{}{
					"@id":   "https://w3id.org/security#publicKeyMultibase",
					"@type": "https://w3id.org/security#multibase",
				},
				"blockchainAccountId": map[string]interface{}{
					"@id":   "https://w3c.github.io/vc-data-integrity/vocab/security/vocabulary.jsonld#blockchainAccountId",
					"@type": "https://w3id.org/security#blockchainAccountId",
				},
			},
		},
	},
	CredentialStatusContext: {
		"@protected":      true,
		"@version":        1.1,
		"hypersign-vocab": "urn:uuid:13fe9318-bb82-4d95-8bf5-8e7fdf8b2026#",
		"xsd":             "http://www.w3.org/2001/XMLSchema#",
		"id":              "@id",
		"revoked": map[string]interface{}{
			"@id":   "hypersign-vocab:revoked",
			"@type": "xsd:boolean",
		},
		"suspended": map[string]interface{}{
			"@id":   "hypersign-vocab:suspended",
			"@type": "xsd:boolean",
		},
		"remarks": map[string]interface{}{
			"@id":   "hypersign-vocab:remarks",
			"@type": "xsd:string",
		},
		"issuer": map[string]interface{}{
			"@id":   "hypersign-vocab:issuer",
			"@type": "xsd:string",
		},
		"issuanceDate": map[string]interface{}{
			"@id":   "hypersign-vocab:issuanceDate",
			"@type": "xsd:dateTime",
		},
		"credentialMerkleRootHash": map[string]interface{}{
			"@id":   "hypersign-vocab:credentialMerkleRootHash",
			"@type": "xsd:string",
		},
	},
	CredentialSchemaContext: {
		"@version":        1.1,
		"hypersign-vocab": "urn:uuid:13fe9318-bb82-4d95-8bf5-8e7fdf8b2026#",
		"xsd":             "http://www.w3.org/2001/XMLSchema#",
		"id":              "@id",
		"type": map[string]interface{}{
			"@id": "hypersign-vocab:type",
		},
		"modelVersion": map[string]interface{}{
			"@id":   "hypersign-vocab:modelVersion",
			"@type": "xsd:string",
		},
		"name": map[string]interface{}{
			"@id":   "hypersign-vocab:name",
			"@type": "xsd:string",
		},
		"author": map[string]interface{}{
			"@id":   "hypersign-vocab:author",
			"@type": "xsd:string",
		},
		"authored": map[string]interface{}{
			"@id":   "hypersign-vocab:authored",
			"@type": "xsd:dateTime",
		},
		"schema": map[string]interface{}{
			"@id":   "hypersign-vocab:schema",
			"@type": "xsd:string",
		},
		"additionalProperties": map[string]interface{}{
			"@id":   "hypersign-vocab:additionalProperties",
			"@type": "xsd:boolean",
		},
		"description": map[string]interface{}{
			"@id":   "hypersign-vocab:description",
			"@type": "xsd:string",
		},
		"properties": map[string]interface{}{
			"@id":   "hypersign-vocab:properties",
			"@type": "xsd:string",
		},
		"required": map[string]interface{}{
			"@id":        "hypersign-vocab:required",
			"@container": "@set",
		},
	},
}
