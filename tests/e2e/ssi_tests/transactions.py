import os
import sys

sys.path.insert(1, os.getcwd())
from utils import run_command

import json
from utils import run_command, get_document_signature

COMMON_CREATE_DID_TX_COMMAND_FLAGS = "--chain-id hidnode --output json --broadcast-mode block --fees 4000uhid  --keyring-backend test --yes"
COMMON_UPDATE_DID_TX_COMMAND_FLAGS = "--chain-id hidnode --output json --broadcast-mode block --fees 1000uhid  --keyring-backend test --yes"
COMMON_DEACTIVATE_DID_COMMAND_FLAGS = "--chain-id hidnode --output json --broadcast-mode block --fees 1000uhid  --keyring-backend test --yes"
COMMON_CREATE_SCHEMA_COMMAND_FLAGS = "--chain-id hidnode --output json --broadcast-mode block --fees 2000uhid  --keyring-backend test --yes"
COMMON_UPDATE_SCHEMA_COMMAND_FLAGS = "--chain-id hidnode --output json --broadcast-mode block --fees 2000uhid  --keyring-backend test --yes"
COMMON_REGISTER_CREDENTIAL_STATUS_COMMAND_FLAGS = "--chain-id hidnode --output json --broadcast-mode block --fees 2000uhid  --keyring-backend test --yes"
COMMON_UPDATE_CREDENTIAL_STATUS_COMMAND_FLAGS = "--chain-id hidnode --output json --broadcast-mode block --fees 2000uhid  --keyring-backend test --yes"

def form_did_create_tx_multisig(diddoc, signPairs, blockchain_account):
    proofsStr = ""
    
    for signPair in signPairs:
        vmId = signPair["verificationMethodId"]
        signAlgo = signPair["signing_algo"]

        base_diddoc_proof = {
            "type": signAlgo,
            "created": "2022-08-16T10:22:12Z",
            "verificationMethod": vmId,
            "proofPurpose": "assertionMethod",
            "proofValue": "",
        }

        signature = get_document_signature(diddoc, "did", signPair["kp"], signAlgo, base_diddoc_proof)
        base_diddoc_proof["proofValue"] = signature
        proofsStr += f"'{json.dumps(base_diddoc_proof)}' "

    cmd_string = f"hid-noded tx ssi register-did '{json.dumps(diddoc)}' {proofsStr} --from {blockchain_account} " + COMMON_CREATE_DID_TX_COMMAND_FLAGS
    return cmd_string

def form_did_update_tx_multisig(diddoc, signPairs, blockchain_account):
    proofsStr = ""
    
    for signPair in signPairs:
        vmId = signPair["verificationMethodId"]
        signAlgo = signPair["signing_algo"]

        base_diddoc_proof = {
            "type": signAlgo,
            "created": "2022-08-16T10:22:12Z",
            "verificationMethod": vmId,
            "proofPurpose": "assertionMethod",
            "proofValue": "",
        }

        signature = get_document_signature(diddoc, "did", signPair["kp"], signAlgo, base_diddoc_proof)
        base_diddoc_proof["proofValue"] = signature
        proofsStr += f"'{json.dumps(base_diddoc_proof)}' "

    version_id = query_did(diddoc["id"])["didDocumentMetadata"]["versionId"]
    cmd_string = f"hid-noded tx ssi update-did '{json.dumps(diddoc)}' '{version_id}' {proofsStr} --from {blockchain_account} " + COMMON_UPDATE_DID_TX_COMMAND_FLAGS
    return cmd_string

def form_did_deactivate_tx_multisig(didId, signPairs, blockchain_account):
    proofsStr = ""
    didDocState = query_did(didId) 
    diddoc = didDocState["didDocument"]
    version_id = didDocState["didDocumentMetadata"]["versionId"]
    
    for signPair in signPairs:
        vmId = signPair["verificationMethodId"]
        signAlgo = signPair["signing_algo"]

        base_diddoc_proof = {
            "type": signAlgo,
            "created": "2022-08-16T10:22:12Z",
            "verificationMethod": vmId,
            "proofPurpose": "assertionMethod",
            "proofValue": "",
        }

        signature = get_document_signature(diddoc, "did", signPair["kp"], signAlgo, base_diddoc_proof)
        base_diddoc_proof["proofValue"] = signature
        proofsStr += f"'{json.dumps(base_diddoc_proof)}' "
    
    cmd_string = f"hid-noded tx ssi deactivate-did '{didId}' '{version_id}' {proofsStr} --from {blockchain_account} " + COMMON_DEACTIVATE_DID_COMMAND_FLAGS
    return cmd_string

def form_create_schema_tx(schema_msg, schema_proof, blockchain_account):
    cmd_string = f"hid-noded tx ssi create-schema '{json.dumps(schema_msg)}' '{json.dumps(schema_proof)}' --from {blockchain_account} " + COMMON_CREATE_SCHEMA_COMMAND_FLAGS
    return cmd_string

def form_update_schema_tx(schema_msg, schema_proof, blockchain_account):
    cmd_string = f"hid-noded tx ssi update-schema '{json.dumps(schema_msg)}' '{json.dumps(schema_proof)}' --from {blockchain_account} " + COMMON_UPDATE_SCHEMA_COMMAND_FLAGS
    return cmd_string

def form_create_cred_status_tx(cred_msg, cred_proof, blockchain_account):
    cmd_string = f"hid-noded tx ssi register-credential-status '{json.dumps(cred_msg)}' '{json.dumps(cred_proof)}' --from {blockchain_account} " + COMMON_REGISTER_CREDENTIAL_STATUS_COMMAND_FLAGS
    return cmd_string

def form_update_cred_status_tx(cred_msg, cred_proof, blockchain_account):
    cmd_string = f"hid-noded tx ssi update-credential-status '{json.dumps(cred_msg)}' '{json.dumps(cred_proof)}' --from {blockchain_account} " + COMMON_UPDATE_CREDENTIAL_STATUS_COMMAND_FLAGS
    return cmd_string

def query_did(did_id):
    cmd_string = f"hid-noded q ssi did {did_id} --output json"
    did_doc, _ = run_command(cmd_string)

    if did_doc == "":
        raise Exception(f"unable to fetch DID Document {did_id} from server")

    return json.loads(did_doc)