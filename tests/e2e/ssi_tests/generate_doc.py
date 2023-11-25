import os
import sys
sys.path.insert(1, os.getcwd())

import json
from utils import run_command, generate_document_id, get_document_signature, \
    secp256k1_pubkey_to_address

ED25519_CONTEXT = "https://w3id.org/security/suites/ed25519-2020/v1"
DID_CONTEXT = "https://www.w3.org/ns/did/v1"
SECP256K1_RECOVERY_CONTEXT = "https://ns.did.ai/suites/secp256k1-2020/v1"
SECP256K1_VER_KEY_2019_CONTEXT = "https://ns.did.ai/suites/secp256k1-2019/v1"
BBS_CONTEXT = "https://ns.did.ai/suites/bls12381-2020/v1"
CREDENTIAL_STATUS_CONTEXT = "https://raw.githubusercontent.com/hypersign-protocol/hypersign-contexts/main/CredentialStatus.jsonld"

def generate_did_document(key_pair, algo="Ed25519Signature2020", bech32prefix="hid", is_uuid=False):
    base_document = {
        "context" : [
            DID_CONTEXT,
        ],
        "id": "",
        "controller": [],
        "verificationMethod": [],
        "authentication": [],
    }
    if algo == "Ed25519Signature2020":
        base_document["context"].append(ED25519_CONTEXT)
    if algo == "EcdsaSecp256k1RecoverySignature2020":
        base_document["context"].append(SECP256K1_RECOVERY_CONTEXT)
    if algo == "EcdsaSecp256k1Signature2019":
        base_document["context"].append(SECP256K1_VER_KEY_2019_CONTEXT)
    if algo == "BbsBlsSignature2020":
        base_document["context"].append(BBS_CONTEXT)
    did_id = generate_document_id("did", key_pair, algo, is_uuid)
    
    # Form the DID Document
    vm_type = ""
    if algo == "Ed25519Signature2020":
        vm_type = "Ed25519VerificationKey2020"
    elif algo == "EcdsaSecp256k1Signature2019":
        vm_type = "EcdsaSecp256k1VerificationKey2019"
    elif algo == "EcdsaSecp256k1RecoverySignature2020":
        vm_type = "EcdsaSecp256k1RecoveryMethod2020"
    elif algo == "BbsBlsSignature2020":
        vm_type = "Bls12381G2Key2020"
    elif algo == "BabyJubJubSignature2023":
        vm_type = "BabyJubJubVerificationKey2023"
    else:
        raise Exception("unknown signing algorithm: " + algo)

    verification_method = {}
    if algo == "EcdsaSecp256k1RecoverySignature2020":
        verification_method = {
            "id": "",
            "type": "",
            "controller": "",
            "blockchainAccountId": ""
        }
    else:
        verification_method = {
            "id": "",
            "type": "",
            "controller": "",
            "publicKeyMultibase": ""
        }

    if algo == "EcdsaSecp256k1RecoverySignature2020":
        verification_method["blockchainAccountId"] = "eip155:1:" + key_pair["ethereum_address"]
    elif algo == "EcdsaSecp256k1Signature2019":

        if bech32prefix == "hid":
            verification_method["blockchainAccountId"] = "cosmos:jagrat:" + \
                secp256k1_pubkey_to_address(key_pair["pub_key_base_64"], bech32prefix)
            did_id = "did:hid:devnet:" + verification_method["blockchainAccountId"]
        elif bech32prefix == "osmo":
            verification_method["blockchainAccountId"] = "cosmos:osmosis-1:" + \
                secp256k1_pubkey_to_address(key_pair["pub_key_base_64"], bech32prefix)
            did_id = "did:hid:devnet:" + verification_method["blockchainAccountId"]
        else:
            verification_method["blockchainAccountId"] = ""
    
        verification_method["publicKeyMultibase"] = key_pair["pub_key_multibase"]
    else:
        verification_method["publicKeyMultibase"] = key_pair["pub_key_multibase"]
    
    verification_method["controller"] = did_id
    verification_method["type"] = vm_type
    verification_method["id"] = did_id + "#k1"

    base_document["id"] = did_id
    base_document["controller"] = [did_id]
    base_document["verificationMethod"] = [verification_method]
    base_document["authentication"] = []
    base_document["assertionMethod"] = []
    return base_document
    
def generate_schema_document(key_pair, schema_author, vm, signature=None, algo="Ed25519Signature2020", updated_schema=None):
    base_schema_doc = {
        "type": "https://schema.org/Person",
        "modelVersion": "v1.0",
        "id": "",
        "name": "Person",
        "author": "",
        "authored": "2022-08-16T10:22:12Z",
        "schema": {
            "schema":"https://json-schema.org/draft-07/schema#",
            "description":"Person Schema",
            "type":"object",
            "properties":"{givenName:{type:string},gender:{type:string},email:{type:text},address:{type:text}}",
            "required":["givenName","address"],
        }
    }
    
    proof_type = ""
    if algo == "Ed25519Signature2020":
        proof_type = "Ed25519Signature2020"
    elif algo == "EcdsaSecp256k1Signature2019":
        proof_type = "EcdsaSecp256k1Signature2019"
    elif algo == "EcdsaSecp256k1RecoverySignature2020":
        proof_type = "EcdsaSecp256k1RecoverySignature2020"
    elif algo == "BbsBlsSignature2020":
        proof_type = "BbsBlsSignature2020"
    elif algo == "BabyJubJubSignature2023":
        proof_type = "BabyJubJubSignature2023"
    else:
        raise Exception("Invalid signing algo: " + algo)

    base_schema_proof = {
        "type": proof_type,
        "created": "2022-08-16T10:22:12Z",
        "verificationMethod": "",
        "proofValue": "",
        "proofPurpose": "assertionMethod"
    }
    
    schema_id = generate_document_id("schema", algo=algo)
    base_schema_doc["id"] = schema_id
    base_schema_doc["author"] = schema_author
    base_schema_proof["verificationMethod"] = vm

    # Form Signature
    if not updated_schema:
        if not signature:
            signature = get_document_signature(base_schema_doc, "schema", key_pair, algo)
        base_schema_proof["proofValue"] = signature
        return base_schema_doc, base_schema_proof
    else:
        if not signature:
            signature = get_document_signature(updated_schema, "schema", key_pair, algo)
        base_schema_proof["proofValue"] = signature
        return updated_schema, base_schema_proof

def generate_cred_status_document(key_pair, cred_author, vm, signature=None, algo="Ed25519Signature2020", updated_credstatus_doc=None):
    base_cred_status_doc = {
        "@context": [CREDENTIAL_STATUS_CONTEXT],
        "id": "",
        "issuer": "did:hid:devnet:z3861habXtUFLNuu6J7m5p8VPsoBMduYbYeUxfx9CnWZR",
        "issuanceDate": "2022-08-16T09:37:12Z",
        "credentialMerkleRootHash": "f35c3a4e3f1b8ba54ee3cf59d3de91b8b357f707fdb72a46473b65b46f92f80b"
    }
    
    proof_type = ""
    if algo == "Ed25519Signature2020":
        proof_type = "Ed25519Signature2020"
        base_cred_status_doc["@context"].append(ED25519_CONTEXT)
    elif algo == "EcdsaSecp256k1Signature2019":
        proof_type = "EcdsaSecp256k1Signature2019"
        base_cred_status_doc["@context"].append(SECP256K1_VER_KEY_2019_CONTEXT)
    elif algo == "EcdsaSecp256k1RecoverySignature2020":
        proof_type = "EcdsaSecp256k1RecoverySignature2020"
        base_cred_status_doc["@context"].append(SECP256K1_RECOVERY_CONTEXT)
    elif algo == "BbsBlsSignature2020":
        proof_type = "BbsBlsSignature2020"
        base_cred_status_doc["@context"].append(BBS_CONTEXT)
    elif algo == "BabyJubJubSignature2023":
        proof_type = "BabyJubJubSignature2023"
    else:
        raise Exception("Invalid signing algo: " + algo)

    base_cred_status_proof = {
        "type": proof_type,
        "created": "2022-08-16T09:37:12Z",
        "verificationMethod": "",
        "proofValue": "",
        "proofPurpose": "assertionMethod"
    }

    cred_id = generate_document_id("cred-status", algo=algo)
    base_cred_status_doc["id"] = cred_id
    base_cred_status_doc["issuer"] = cred_author
    base_cred_status_proof["verificationMethod"] = vm

    # Form Signature
    if not updated_credstatus_doc:
        if not signature:
            signature = get_document_signature(base_cred_status_doc, "cred-status", key_pair, algo, proofObj=base_cred_status_proof)
        base_cred_status_proof["proofValue"] = signature
        return base_cred_status_doc, base_cred_status_proof
    else:
        if not signature:
            signature = get_document_signature(updated_credstatus_doc, "cred-status", key_pair, algo, proofObj=base_cred_status_proof)
        base_cred_status_proof["proofValue"] = signature
        return updated_credstatus_doc, base_cred_status_proof