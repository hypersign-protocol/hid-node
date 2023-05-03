import os
import sys
sys.path.insert(1, os.getcwd())

import json
from utils import run_command, generate_document_id, get_document_signature, \
    secp256k1_pubkey_to_address

def generate_did_document(key_pair, algo="ed25519", bech32prefix="hid", is_uuid=False):
    base_document = {
        "id": "",
        "controller": [],
        "verificationMethod": [],
        "authentication": [],
    }

    did_id = generate_document_id("did", key_pair, algo, is_uuid)

    # Form the DID Document
    vm_type = ""
    if algo == "ed25519":
        vm_type = "Ed25519VerificationKey2020"
    elif algo == "secp256k1":
        vm_type = "EcdsaSecp256k1VerificationKey2019"
    elif algo == "recover-eth":
        vm_type = "EcdsaSecp256k1RecoveryMethod2020"
    else:
        raise Exception("unknown signing algorithm: " + key_pair)

    verification_method = {
        "id": "",
        "type": "",
        "controller": "",
        "blockchainAccountId": "",
        "publicKeyMultibase": ""
    }

    if algo == "recover-eth":
        verification_method["blockchainAccountId"] = "eip155:1:" + key_pair["ethereum_address"]
    elif algo == "secp256k1":

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
    
def generate_schema_document(key_pair, schema_author, vm, signature=None, algo="ed25519"):
    base_schema_doc = {
        "type": "https://schema.org/Person",
        "modelVersion": "v1.0",
        "id": "",
        "name": "Person",
        "author": "",
        "authored": "2022-08-16T10:22:12Z",
        "schema": {
            "description":"Person Schema",
            "type":"object",
            "properties":"{givenName:{type:string},gender:{type:string},email:{type:text},address:{type:text}}",
            "required":["givenName","address"],
        }
    }
    
    proof_type = ""
    if algo == "ed25519":
        proof_type = "Ed25519Signature2020"
    elif algo == "secp256k1":
        proof_type = "EcdsaSecp256k1Signature2019"
    elif algo == "recover-eth":
        proof_type = "EcdsaSecp256k1RecoverySignature2020"
    else:
        raise Exception("Invalid signing algo: " + algo)

    base_schema_proof = {
        "type": proof_type,
        "created": "2022-08-16T10:22:12Z",
        "verificationMethod": "",
        "proofValue": "",
        "proofPurpose": "assertion"
    }

    schema_id = generate_document_id("schema", algo=algo)
    base_schema_doc["id"] = schema_id
    base_schema_doc["author"] = schema_author
    base_schema_proof["verificationMethod"] = vm

    # Form Signature
    if not signature:
        signature = get_document_signature(base_schema_doc, "schema", key_pair, algo)
        
    base_schema_proof["proofValue"] = signature

    return base_schema_doc, base_schema_proof

def generate_cred_status_document(key_pair, cred_author, vm, signature=None, algo="ed25519"):
    base_cred_status_doc = {
        "claim": {
                "id": "",
                "currentStatus": "Live",
                "statusReason": "Credential Active"
        },
        "issuer": "did:hid:devnet:z3861habXtUFLNuu6J7m5p8VPsoBMduYbYeUxfx9CnWZR",
        "issuanceDate": "2022-08-16T09:37:12Z",
        "expirationDate": "2023-08-16T09:40:12Z",
        "credentialHash": "f35c3a4e3f1b8ba54ee3cf59d3de91b8b357f707fdb72a46473b65b46f92f80b"
    }
    
    proof_type = ""
    if algo == "ed25519":
        proof_type = "Ed25519Signature2020"
    elif algo == "secp256k1":
        proof_type = "EcdsaSecp256k1Signature2019"
    elif algo == "recover-eth":
        proof_type = "EcdsaSecp256k1RecoverySignature2020"
    else:
        raise Exception("Invalid signing algo: " + algo)
    
    base_cred_status_proof = {
        "type": proof_type,
        "created": "2022-08-16T09:37:12Z",
        "updated": "2022-08-16T09:37:12Z",
        "verificationMethod": "",
        "proofValue": "",
        "proofPurpose": "assertion"
    }

    cred_id = generate_document_id("cred-status", algo=algo)
    base_cred_status_doc["claim"]["id"] = cred_id
    base_cred_status_doc["issuer"] = cred_author
    base_cred_status_proof["verificationMethod"] = vm

    # Form Signature
    if not signature:
        signature = get_document_signature(base_cred_status_doc, "cred-status", key_pair, algo)

    base_cred_status_proof["proofValue"] = signature

    return base_cred_status_doc, base_cred_status_proof