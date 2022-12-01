import subprocess
import json

def run_command(cmd_string):
    if type(cmd_string) != str:
        raise Exception("input parameter should be a string")
    
    cmd_run = subprocess.run(cmd_string, stdout=subprocess.PIPE, shell=True)
    cmd_output = cmd_run.stdout.decode('utf-8')[:-1]

    if cmd_run.returncode != 0:
        raise Exception(f"Command {cmd_string} returned error. Check the log: \n {cmd_output}")
    
    return cmd_output

def run_blockchain_command(cmd_string: str, transaction_name: str = None, expect_failure: bool = False):
    if not expect_failure:
        try:
            tx_out = run_command(cmd_string)
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
            tx_out = run_command(cmd_string)
            tx_out_json = json.loads(tx_out)
            tx_status_code = tx_out_json["code"]
            tx_status_log = tx_out_json["raw_log"]
            if tx_out_json["code"] == 0:
                raise Exception(f"{transaction_name} transaction was expected to fail, but it didn't")
            print(f"{transaction_name} : transaction failed as expected with error code {tx_status_code}. Log: {tx_status_log}\n")
        except Exception as e:
            print(f"{transaction_name} : Error while executing transaction command\n")
            raise(e)

def generate_key_pair(algo="ed25519"):
    cmd = ""
    if algo == "ed25519":
        cmd = "hid-noded debug ed25519 random"
    else:
        raise Exception(algo + " is not a supported signing algorithm")
    result_str = run_command(cmd)
    kp = json.loads(result_str)
    return kp

def generate_document_id(doc_type: str, kp: dict = None):
    id = ""
    if not kp:
        kp = generate_key_pair()
    
    public_key_multibase = kp["pub_key_multibase"]
    
    if public_key_multibase == None:
        raise Exception("Public key is empty")
    
    if doc_type == "did":
        id = "did:hid:devnet:" + public_key_multibase
    elif doc_type == "schema":
        id = "sch:hid:devnet:" + public_key_multibase + ":1.0"
    elif doc_type == "cred-status":
        id = "vc:hid:devnet:" + public_key_multibase
    else:
        raise Exception("The argument `doc_type` only accepts the values: did, schema and cred-status")
    
    return id

def is_blockchain_active(rpc_port):
    import socket
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        assert s.connect_ex(('localhost', rpc_port)) == 0, f"hid-noded is not running"

def get_document_signature(doc: dict, doc_type: str, key_pair: dict):
    private_key = key_pair["priv_key_base_64"]

    if doc_type == "cred-status":
        doc_cmd = "cred-status-doc"
    elif doc_type == "schema":
        doc_cmd = "schema-doc"
    else:
        raise Exception("Invalid value for doc_type param: " + doc_type)
    
    cmd_string = f"hid-noded debug ed25519 sign-ssi-doc {doc_cmd} '{json.dumps(doc)}' {private_key}"
    signature = run_command(cmd_string)
    return signature