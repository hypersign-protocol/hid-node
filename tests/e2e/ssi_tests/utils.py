import subprocess
import json
import uuid

def run_command(cmd_string):
    if type(cmd_string) != str:
        raise Exception("input parameter should be a string")
    
    cmd_run = subprocess.run(cmd_string, stdout=subprocess.PIPE, stderr=subprocess.PIPE, shell=True)
    cmd_output = cmd_run.stdout.decode('utf-8')[:-1]

    # if cmd_run.returncode != 0:
    #     raise Exception(f"Command {cmd_string} returned error. Check the log: \n {cmd_output}")
    
    return cmd_output, cmd_run.returncode

def run_blockchain_command(cmd_string: str, transaction_name: str = None, expect_failure: bool = False, stateless_err: bool = False):
    if not expect_failure:
        try:
            tx_out, _ = run_command(cmd_string)
            tx_out_json = json.loads(tx_out)
            if tx_out_json["code"] != 0:
                error_log = tx_out_json["raw_log"]
                raise Exception(f"Transaction failed: Log -> {error_log}")
            print(f"{transaction_name} : transaction was successful\n")
        except Exception as e:
            print(f"{transaction_name} : Error while executing transaction command\n")
            raise(e)
    else:
        try:
            tx_out, return_code = run_command(cmd_string)
            if return_code != 0 and not stateless_err:
                raise Exception(f"tx command {cmd_string} failed")
            else:
                if not stateless_err:
                    tx_out_json = json.loads(tx_out)
                    tx_status_code = tx_out_json["code"]
                    tx_status_log = tx_out_json["raw_log"]
                    if tx_out_json["code"] == 0:
                        raise Exception(f"{transaction_name} transaction was expected to fail, but it didn't")
                    print(f"{transaction_name} : transaction failed as expected with error code {tx_status_code}. Log: {tx_status_log}\n")
                else:
                    print(f"{transaction_name} : transaction failed as expected during stateless validation\n")
        except Exception as e:
            print(f"{transaction_name} : Error while executing transaction command\n")
            raise(e)

def generate_key_pair(algo="Ed25519Signature2020"):
    cmd = ""
    if algo == "Ed25519Signature2020":
        cmd = "hid-noded debug ed25519 random"
    elif algo == "secp256k1":
        cmd = "hid-noded debug secp256k1 random"
    elif algo == "EcdsaSecp256k1RecoverySignature2020":
        cmd = "hid-noded debug secp256k1 eth-hex-random"
    elif algo == "bbs":
        cmd = "hid-noded debug bbs random"
    elif algo == "bjj":
        cmd = "hid-noded debug bjj random"
    else:
        raise Exception(algo + " is not a supported signing algorithm")
    result_str, _ = run_command(cmd)
    kp = json.loads(result_str)
    return kp

def add_keyAgreeemnt_pubKeyMultibase(verification_method, type):
    if verification_method["type"] != "Ed25519VerificationKey2020":
        raise Exception("verification method " + verification_method["id"] + " must be of type Ed25519VerificationKey2020")
    
    if type == "X25519KeyAgreementKey2020":
        verification_method["type"] = "X25519KeyAgreementKey2020"
    elif type == "X25519KeyAgreementKeyEIP5630":
        verification_method["type"] = "X25519KeyAgreementKeyEIP5630"
    else:
        raise Exception("invalid key agreement type " + type)

    return verification_method

def generate_document_id(doc_type: str, kp: dict = None, algo: str = "Ed25519Signature2020", is_uuid: bool =False):
    id = ""
    if not kp:
        kp = generate_key_pair(algo)
    
    if is_uuid:
        method_specific_id = str(uuid.uuid4())
    else:
        if algo in ["EcdsaSecp256k1RecoverySignature2020"]:
            method_specific_id = kp["ethereum_address"]
        else:
            method_specific_id = kp["pub_key_multibase"]
        
    if method_specific_id == None:
        raise Exception("Public key is empty")
    
    if doc_type == "did":
        id = "did:hid:devnet:" + method_specific_id
    elif doc_type == "schema":
        id = "sch:hid:devnet:" + method_specific_id + ":1.0"
    elif doc_type == "cred-status":
        id = "vc:hid:devnet:" + method_specific_id
    else:
        raise Exception("The argument `doc_type` only accepts the values: did, schema and cred-status")
    
    return id

def is_blockchain_active(rpc_port):
    import socket
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        assert s.connect_ex(('localhost', rpc_port)) == 0, f"hid-noded is not running"

def get_document_signature(doc: dict, doc_type: str, key_pair: dict, algo: str = "ed25519"):
    if algo in ["EcdsaSecp256k1RecoverySignature2020", "bjj"]:
        private_key = key_pair["priv_key_hex"]
    else:
        private_key = key_pair["priv_key_base_64"]

    if doc_type == "cred-status":
        doc_cmd = "cred-status-doc"
    elif doc_type == "schema":
        doc_cmd = "schema-doc"
    else:
        raise Exception("Invalid value for doc_type param: " + doc_type)
    
    cmd_string = f"hid-noded debug sign-ssi-doc {doc_cmd} '{json.dumps(doc)}' {private_key} {algo}"
    signature, _ = run_command(cmd_string)

    if signature == "":
        raise Exception(f"Signature came empty while running command: {cmd_string}")
    return signature

def secp256k1_pubkey_to_address(pub_key, prefix):
    cmd_string = f"hid-noded debug secp256k1 bech32-addr {pub_key} {prefix}"
    addr, _ = run_command(cmd_string)
    return addr
