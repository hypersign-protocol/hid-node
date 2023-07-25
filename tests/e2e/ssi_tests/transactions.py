import os
import sys

sys.path.insert(1, os.getcwd())
from utils import run_command

import json
from utils import run_command

COMMON_CREATE_DID_TX_COMMAND_FLAGS = "--chain-id hidnode --output json --broadcast-mode block --fees 4000uhid  --keyring-backend test --yes"
COMMON_UPDATE_DID_TX_COMMAND_FLAGS = "--chain-id hidnode --output json --broadcast-mode block --fees 1000uhid  --keyring-backend test --yes"
COMMON_DEACTIVATE_DID_COMMAND_FLAGS = "--chain-id hidnode --output json --broadcast-mode block --fees 1000uhid  --keyring-backend test --yes"
COMMON_CREATE_SCHEMA_COMMAND_FLAGS = "--chain-id hidnode --output json --broadcast-mode block --fees 2000uhid  --keyring-backend test --yes"
COMMON_REGISTER_CREDENTIAL_STATUS_COMMAND_FLAGS = "--chain-id hidnode --output json --broadcast-mode block --fees 2000uhid  --keyring-backend test --yes"

def form_did_create_tx_multisig(diddoc, signPairs, blockchain_account):
    signPairStr = ""
    
    for signPair in signPairs:
        private_key = ""
        vmId = signPair["verificationMethodId"]
        signAlgo = signPair["signing_algo"]

        if signAlgo == "recover-eth":
            private_key = signPair["kp"]["priv_key_hex"]
        else:
            private_key = signPair["kp"]["priv_key_base_64"]

        signPairStr += f"{vmId} {private_key} {signAlgo} "

    cmd_string = f"hid-noded tx ssi create-did '{json.dumps(diddoc)}' {signPairStr} --from {blockchain_account} " + COMMON_CREATE_DID_TX_COMMAND_FLAGS
    return cmd_string

def form_did_update_tx_multisig(diddoc, signPairs, blockchain_account):
    signPairStr = ""
    
    for signPair in signPairs:
        private_key = ""
        vmId = signPair["verificationMethodId"]
        signAlgo = signPair["signing_algo"]

        if signAlgo == "recover-eth":
            private_key = signPair["kp"]["priv_key_hex"]
        else:
            private_key = signPair["kp"]["priv_key_base_64"]

        signPairStr += f"{vmId} {private_key} {signAlgo} "

    version_id = query_did(diddoc["id"])["didDocumentMetadata"]["versionId"]
    cmd_string = f"hid-noded tx ssi update-did '{json.dumps(diddoc)}' '{version_id}' {signPairStr} --from {blockchain_account} " + COMMON_UPDATE_DID_TX_COMMAND_FLAGS
    return cmd_string

def form_did_deactivate_tx_multisig(didId, signPairs, blockchain_account):
    signPairStr = ""
    
    for signPair in signPairs:
        private_key = ""
        vmId = signPair["verificationMethodId"]
        signAlgo = signPair["signing_algo"]

        if signAlgo == "recover-eth":
            private_key = signPair["kp"]["priv_key_hex"]
        else:
            private_key = signPair["kp"]["priv_key_base_64"]

        signPairStr += f"{vmId} {private_key} {signAlgo} "

    version_id = query_did(didId)["didDocumentMetadata"]["versionId"]
    cmd_string = f"hid-noded tx ssi deactivate-did '{didId}' '{version_id}' {signPairStr} --from {blockchain_account} " + COMMON_DEACTIVATE_DID_COMMAND_FLAGS
    return cmd_string

def form_did_create_tx(did_doc, kp, blockchain_account, verificationMethodId=None, signing_algo="ed25519"):
    if signing_algo == "recover-eth":
        private_key = kp["priv_key_hex"]
    else:
        private_key = kp["priv_key_base_64"]
    
    if not verificationMethodId:
        verificationMethodId = did_doc["authentication"][0]
    
    cmd_string = f"hid-noded tx ssi create-did '{json.dumps(did_doc)}' {verificationMethodId} {private_key} {signing_algo} --from {blockchain_account} " + COMMON_CREATE_DID_TX_COMMAND_FLAGS
    return cmd_string

def form_create_schema_tx(schema_msg, schema_proof, blockchain_account):
    cmd_string = f"hid-noded tx ssi create-schema '{json.dumps(schema_msg)}' '{json.dumps(schema_proof)}' --from {blockchain_account} " + COMMON_CREATE_SCHEMA_COMMAND_FLAGS
    return cmd_string

def form_create_cred_status_tx(cred_msg, cred_proof, blockchain_account):
    cmd_string = f"hid-noded tx ssi register-credential-status '{json.dumps(cred_msg)}' '{json.dumps(cred_proof)}' --from {blockchain_account} " + COMMON_REGISTER_CREDENTIAL_STATUS_COMMAND_FLAGS
    return cmd_string

def form_did_update_tx(did_doc, kp, blockchain_account, verificationMethodId=None, signing_algo="ed25519"):
    if signing_algo in ["recover-eth"]:
        private_key = kp["priv_key_hex"]
    else:
        private_key = kp["priv_key_base_64"]

    if not verificationMethodId:
        verificationMethodId = did_doc["authentication"][0]
    version_id = query_did(did_doc["id"])["didDocumentMetadata"]["versionId"]
    cmd_string = f"hid-noded tx ssi update-did '{json.dumps(did_doc)}' '{version_id}' '{verificationMethodId}' {private_key} {signing_algo} --from {blockchain_account} " + COMMON_UPDATE_DID_TX_COMMAND_FLAGS
    return cmd_string

def form_did_deactivate_tx(did_doc_id, kp, blockchain_account, verificationMethodId=None, signing_algo="ed25519"):
    if signing_algo in ["recover-eth"]:
        private_key = kp["priv_key_hex"]
    else:
        private_key = kp["priv_key_base_64"]
    
    if not verificationMethodId:
        verificationMethodId = query_did(did_doc_id)["didDocument"]["authentication"][0]
    version_id = query_did(did_doc_id)["didDocumentMetadata"]["versionId"]
    cmd_string = f"hid-noded tx ssi deactivate-did '{did_doc_id}' '{version_id}' '{verificationMethodId}' {private_key} {signing_algo} --from {blockchain_account} " + COMMON_DEACTIVATE_DID_COMMAND_FLAGS
    return cmd_string

def query_did(did_id):
    cmd_string = f"hid-noded q ssi did {did_id} --output json"
    did_doc, _ = run_command(cmd_string)
    return json.loads(did_doc)