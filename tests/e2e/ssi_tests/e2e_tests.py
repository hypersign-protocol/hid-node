import os
import sys
sys.path.insert(1, os.getcwd())
import time

from utils import run_blockchain_command, generate_key_pair, secp256k1_pubkey_to_address
from generate_doc import generate_did_document, generate_schema_document, generate_cred_status_document
from transactions import form_did_create_tx_multisig, form_did_update_tx_multisig, \
      query_did, form_create_schema_tx, form_did_deactivate_tx_multisig, form_create_cred_status_tx
from constants import DEFAULT_BLOCKCHAIN_ACCOUNT_NAME

# TC - I : Create DID scenarios
def create_did_test():
    print("\n--- Create DID Test ---\n") 

    print("1. FAIL: Alice has a registered DID Document where Alice is the controller. Bob tries to register their DID Document by keeping both Alice and Bob as controllers, and by sending only his signature.\n")
    kp_alice = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_alice)
    did_doc_alice = did_doc_string["id"]
    signPair_alice = {
        "kp": kp_alice,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering Alice's DID with Id: {did_doc_alice}")

    kp_bob = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_bob)
    did_doc_bob = did_doc_string["id"]
    did_doc_string["controller"] = [did_doc_alice, did_doc_bob]
    signPair_bob = {
        "kp": kp_bob,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_bob)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering Bob's DID with Id: {did_doc_bob}", True)

    print("2. PASS: Alice has a registered DID Document where Alice is the controller. Bob tries to register their DID Document by keeping both Alice and Bob as controllers, and by sending only both Alice's and Bob's signatures.\n")
    signers = []
    signers.append(signPair_bob)
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_bob}")
   

    print("3. PASS: Alice has a registered DID Document where Alice is the controller. She tries to create an organization DID, in which Alice is the only controller and it's verification method field is empty.\n")
    kp_org = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_org)
    did_doc_org = did_doc_string["id"]
    did_doc_string["controller"] = [did_doc_alice]
    did_doc_string["verificationMethod"] = []
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_org}")

    #Alice creates a DID where the controller only has Alice's DID and the verfication method has two ETH wallets added. Signature of all hot wallets are passed
    kp_hot_wallet_1 = generate_key_pair("recover-eth")
    kp_hot_wallet_2 = generate_key_pair("recover-eth")
    kp_org = generate_key_pair()

    did_doc_string = generate_did_document(kp_org)
    did_doc_canon = did_doc_string["id"]
    did_doc_string["controller"] = [did_doc_alice]
    did_doc_string["verificationMethod"] = [
        {
            "id": "did:hid:devnet:" + kp_hot_wallet_1["ethereum_address"] + "#k1",
            "controller": did_doc_alice,
            "type": "EcdsaSecp256k1RecoveryMethod2020",
            "blockchainAccountId": "eip155:1:" + kp_hot_wallet_1["ethereum_address"],
        },
        {
            "id": "did:hid:devnet:" + kp_hot_wallet_2["ethereum_address"] + "#k2",
            "controller": did_doc_alice,
            "type": "EcdsaSecp256k1RecoveryMethod2020",
            "blockchainAccountId": "eip155:1:" + kp_hot_wallet_2["ethereum_address"],
        },
    ]
    
    signPair_hotWallet1 = {
        "kp": kp_hot_wallet_1,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": "recover-eth"
    }
    signPair_hotWallet2 = {
        "kp": kp_hot_wallet_2,
        "verificationMethodId": did_doc_string["verificationMethod"][1]["id"],
        "signing_algo": "recover-eth"
    }

    print("4. FAIL: Alice has a registered DID Document where Alice is the controller. Alice tries to register an Org DID Document where Alice is the sole controller, and there are two verification Methods, of type EcdsaSecp256k1RecoveryMethod2020, and Alice is the controller for each one of them. Signature is provided by only one of the VMs.\n")
    signers = []
    signers.append(signPair_hotWallet2)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_canon}", True)

    print("5. PASS: Alice has a registered DID Document where Alice is the controller. Alice tries to register an Org DID Document where Alice is the sole controller, and there are two verification Methods, of type EcdsaSecp256k1RecoveryMethod2020, and Alice is the controller for each one of them. Signature is provided by both VMs.\n")
    signers = []
    signers.append(signPair_hotWallet1)
    signers.append(signPair_hotWallet2)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_canon}")
    
    # Alice creates a DID where they keep the VM of their friend Eve in the verificationMethod list of the document
    print("6. FAIL: Alice creates an Org DID where Alice is the controller, and she adds a verification method of her friend Eve. Only Alice sends the singature.\n")
    kp_eve = generate_key_pair("secp256k1")
    did_doc_string = generate_did_document(kp_eve, algo="secp256k1")
    did_doc_eve = did_doc_string["id"]
    did_doc_eve_vms = did_doc_string["verificationMethod"]
    signers = []
    signPair_eve = {
        "kp": kp_eve,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": "secp256k1"
    }
    signers.append(signPair_eve)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering Eve's DID with Id: {did_doc_eve}")

    kp_random = generate_key_pair()
    did_doc_string = generate_did_document(kp_random)
    did_doc_string["controller"] = [did_doc_alice]
    did_doc_string["verificationMethod"] = [
        did_doc_eve_vms[0]
    ]
    signers = []
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID by passing only Alice's Signature", True)

    print("7. PASS: Alice creates an Org DID where Alice is the controller, and she adds a verification method of her friend Eve. Both Alice and Eve send their singatures.\n")
    signers = []
    signers.append(signPair_alice)
    signers.append(signPair_eve)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID by passing both Alice's and Eve's Signature")

    print("8. FAIL: Alice tries to register a DID Document with duplicate publicKeyMultibase of type Ed25519VerificationKey2020 \n")
    kp_alice = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_alice)
    did_doc_string_2 = generate_did_document(kp_alice)
    
    did_doc_string_vm_1 = did_doc_string["verificationMethod"][0]
    did_doc_string_vm_2 = did_doc_string_2["verificationMethod"][0] 
    did_doc_string_vm_2["id"] = did_doc_string_vm_2["id"] + "new"
    
    did_doc_string["verificationMethod"] = [
        did_doc_string_vm_1,
        did_doc_string_vm_2
    ]

    did_doc_alice = did_doc_string["id"]
    signPair_alice = {
        "kp": kp_alice,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering Alice's DID with Id: {did_doc_alice}", True, True)

    print("--- Test Completed ---\n")

# TC - II : Update DID scenarios
def update_did_test():
    print("\n--- Update DID Test ---\n") 

    print("1. FAIL: Alice creates an Org DID where alice is the controller, and Bob's VM is added to its VM List only. Bob attempts to update Org DID by sending his signature.\n")
    
    # Register Alice's DID
    kp_alice = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_alice)
    did_doc_alice = did_doc_string["id"]
    did_doc_alice_vm = did_doc_string["verificationMethod"][0]
    signPair_alice = {
        "kp": kp_alice,
        "verificationMethodId": did_doc_alice_vm["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering of Alice's DID with Id: {did_doc_alice}")

    # Register Bob's DID
    kp_bob = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_bob)
    did_doc_bob = did_doc_string["id"]
    did_doc_bob_vm = did_doc_string["verificationMethod"][0]
    signPair_bob = {
        "kp": kp_bob,
        "verificationMethodId": did_doc_bob_vm["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_bob)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering of Bob's DID with Id: {did_doc_bob}")

    # Alice creates Organization DID with itself being the only controller and Bob's VM being added to VM List
    kp_org = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_org)
    did_doc_org = did_doc_string["id"]
    did_doc_string["controller"] = [did_doc_alice]
    did_doc_string["verificationMethod"] = [did_doc_bob_vm]
    signers.append(signPair_alice)
    signers.append(signPair_bob)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering of Org DID with Id: {did_doc_org}")

    # Bob (who is not the controller) attempts to make changes in Org DID
    signers = []
    did_doc_string["context"] = ["gmm", "gmm2"]
    signers.append(signPair_bob)
    update_tx_cmd = form_did_update_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Bob (non-controller) attempts to update Org DID with Id: {did_doc_org}", True)

    # Alice (who is the controller) attempts to make changes in Org DID
    print("2. PASS: Alice creates an Org DID where alice is the controller, and Bob's VM is added to its VM List only. Alice attempts to update Org DID by sending her signature.\n")
    signers = []
    did_doc_string["context"] = ["gmm", "gmm2"]
    signers.append(signPair_alice)
    update_tx_cmd = form_did_update_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Alice (controller) attempts to update Org DID with Id: {did_doc_org}")

    did_doc_string_org = did_doc_string

    print("3. FAIL: Alice attempts to add George as controller of Org ID. Only George's Signature is sent.\n")
    signers = []
    kp_george = generate_key_pair()
    did_doc_string = generate_did_document(kp_george)
    did_doc_george = did_doc_string["id"]
    did_doc_george_vm = did_doc_string["verificationMethod"][0]
    signPair_george = {
        "kp": kp_george,
        "verificationMethodId": did_doc_george_vm["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_george)
    update_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Registering George's DID with Id: {did_doc_george}")

    # Addition of George's DID to controller by only Alice's siganature
    signers = []
    did_doc_string_org["controller"] = [did_doc_alice, did_doc_george]
    signers.append(signPair_alice)
    update_tx_cmd = form_did_update_tx_multisig(did_doc_string_org, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Adding George's DID as controller with Alice's signature only", True) 

    # Addition of George's DID to controller by only George's siganature
    print("4. FAIL: Alice attempts to add George as controller of Org ID. Only Alice's Signature is sent.\n")
    signers = []
    did_doc_string_org["controller"] = [did_doc_alice, did_doc_george]
    signers.append(signPair_george)
    update_tx_cmd = form_did_update_tx_multisig(did_doc_string_org, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Adding George's DID as controller with George's signature only", True)

    # Addition of George's DID to controller by George's and Alice's siganature
    print("5. PASS: Alice attempts to add George as controller of Org ID. Both Alice's and George's Signatures are sent.\n")
    signers = []
    did_doc_string_org["controller"] = [did_doc_alice, did_doc_george]
    signers.append(signPair_alice)
    signers.append(signPair_george)
    update_tx_cmd = form_did_update_tx_multisig(did_doc_string_org, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Adding George's DID as controller with George's and Alice's signature")

    # Removal of George's controller by Alice's signature
    print("6. PASS: Alice attempts to remove George as controller of Org ID. Only Alice's Signature is sent.\n")
    signers = []
    did_doc_string_org["controller"] = [did_doc_alice]
    signers.append(signPair_alice)
    update_tx_cmd = form_did_update_tx_multisig(did_doc_string_org, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Removal of George's controller by Alice's signature")


    # Addition of George's controller and removal of Alice's controller at same time
    print("7. PASS: Addition of George as a controller and simultaneous removal of Alice as a controller. Both alice's and george's signature are passed.\n")
    signers = []
    did_doc_string_org["controller"] = [did_doc_george]
    signers.append(signPair_alice)
    signers.append(signPair_george)
    update_tx_cmd = form_did_update_tx_multisig(did_doc_string_org, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Removal of Alice's controller and addition of Bob's controller")

    #Alice creates a DID where the controller only has Alice's DID and the verfication method has two ETH wallets added. Signature of all hot wallets are passed
    print("8. FAIL: Alice has already created two didDocs, each representing a wallet. Now, she creates a DID where she is the controller, and the VMs from two DIDDocs are just added in the VM List. Each of these VMs have different controllers. One of the VMs attempts to update DIDDoc.\n")
    kp_hot_wallet_1 = generate_key_pair("recover-eth")
    signers = []
    did_doc_string = generate_did_document(kp_hot_wallet_1, algo="recover-eth")
    did_doc_hw1 = did_doc_string["id"]
    did_doc_hw1_vm = did_doc_string["verificationMethod"][0]
    signPair_hw1 = {
        "kp": kp_hot_wallet_1,
        "verificationMethodId": did_doc_hw1_vm["id"],
        "signing_algo": "recover-eth"
    }
    signers.append(signPair_hw1)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering Hot Wallet 1 with Id: {did_doc_hw1}")

    kp_hot_wallet_2 = generate_key_pair("recover-eth")
    signers = []
    did_doc_string = generate_did_document(kp_hot_wallet_2, algo="recover-eth")
    did_doc_hw2 = did_doc_string["id"]
    did_doc_hw2_vm = did_doc_string["verificationMethod"][0]
    signPair_hw2 = {
        "kp": kp_hot_wallet_2,
        "verificationMethodId": did_doc_hw2_vm["id"],
        "signing_algo": "recover-eth"
    }
    signers.append(signPair_hw2)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering Hot Wallet 2 with Id: {did_doc_hw2}")

    kp_org = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_org)
    did_doc_org = did_doc_string["id"]
    did_doc_string["controller"] = [did_doc_alice]
    did_doc_string["verificationMethod"] = [
        did_doc_hw1_vm,
        did_doc_hw2_vm
    ]
    signers.append(signPair_alice)
    signers.append(signPair_hw1)
    signers.append(signPair_hw2)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering of Org DID with Id: {did_doc_org}")

    print("Hot-Wallet 1 attemps to update Org DID. It is expected to fail")
    did_doc_string["context"] = ["exempler.org"]
    signers = []
    signers.append(signPair_hw1)
    update_tx_cmd = form_did_update_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Hot Wallet 1 attempts to update Tx", True)

    print("9. PASS: Alice has already created two didDocs, each representing a wallet. Now, she creates a DID where she is the controller, and the VMs from two DIDDocs are just added in the VM List. Each of these VMs have different controllers. Alice attempts to update DIDDoc.\n")
    signers = []
    signers.append(signPair_alice)
    update_tx_cmd = form_did_update_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Alice (controller) attempts to update Tx")

    print("10. FAIL: Alice tries to update her DID Document without changing anything\n")
    did_doc_string = query_did(did_doc_alice)["didDocument"]
    signers = []
    signers.append(signPair_alice)
    update_tx_cmd = form_did_update_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Alice attempts update without any change Tx", True)

    # Register Alice's DID
    print("11. PASS: Jenny creates herself a DID with empty Controller list. She then attempts to update the DIDDoc by changing the context field and passes her signature only.\n")
    kp_jenny = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_jenny)
    did_doc_string["controller"] = []
    did_doc_jenny = did_doc_string["id"]
    did_doc_jenny_vm = did_doc_string["verificationMethod"][0]
    signPair_jenny = {
        "kp": kp_jenny,
        "verificationMethodId": did_doc_jenny_vm["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_jenny)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering of Jenny's DID with Id: {did_doc_jenny}")

    signers = []
    did_doc_string["context"] = ["yo"]
    signers.append(signPair_jenny)
    update_tx_cmd = form_did_update_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Jenny (controller) attempts to update Tx")

    print("12. FAIL: Jenny creates a DID. She then attempts to update the DIDDoc by adding a new Verification method. She passes signature only for old VM\n")
    kp_jenny = generate_key_pair()
    kp_jenny_2 = generate_key_pair()

    signers = []
    did_doc_string = generate_did_document(kp_jenny)
    did_doc_jenny = did_doc_string["id"]
    did_doc_jenny_vm = did_doc_string["verificationMethod"][0]
    signPair_jenny = {
        "kp": kp_jenny,
        "verificationMethodId": did_doc_jenny_vm["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_jenny)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering of Jenny's DID with Id: {did_doc_jenny}")

    signers = []
    did_doc_string_alice = generate_did_document(kp_jenny_2)
    new_vm = did_doc_string_alice["verificationMethod"][0]
    new_vm_id = did_doc_string["verificationMethod"][0]["id"] + "news"
    new_vm["id"] = new_vm_id
    new_vm["controller"] = did_doc_string["id"]

    did_doc_string["verificationMethod"] = [
        did_doc_string["verificationMethod"][0],
        new_vm,
    ]
    signPair_jenny_1 = {
        "kp": kp_jenny,
        "verificationMethodId": did_doc_jenny_vm["id"],
        "signing_algo": "ed25519"
    }
    signPair_jenny_2 = {
        "kp": kp_jenny_2,
        "verificationMethodId": new_vm_id,
        "signing_algo": "ed25519"
    }
    signers.append(signPair_jenny_1)
    update_tx_cmd = form_did_update_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Jenny attempts to update the DIDDoc with only old VM's signature", True)

    print("13. PASS: Jenny attempts to update the same didDoc by passing signatures for both old and new verification methods\n")

    signers = []
    signers.append(signPair_jenny_1)
    signers.append(signPair_jenny_2)
    update_tx_cmd = form_did_update_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Jenny attempts to update the DIDDoc with both new and old VM's signature")
    
    print("14. PASS: Jenny removes the inital verification method she had added. She passes only one signature corresponding to the lastest VM")

    did_doc_string["verificationMethod"] = [
        new_vm,
    ]
    signers = []
    signers.append(signPair_jenny_1)
    update_tx_cmd = form_did_update_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Jenny attempts to remove Verification Method using signature corresponding to other verification method")

    print("--- Test Completed ---\n")

def deactivate_did():
    print("\n--- Deactivate DID Test ---\n")

    print("1. PASS: Alice creates an Org DID with herself and Bob being the Controller. Alice attempts to deactivate it \n")
    
    # Register Alice's DID
    kp_alice = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_alice)
    did_doc_alice = did_doc_string["id"]
    did_doc_alice_vm = did_doc_string["verificationMethod"][0]
    signPair_alice = {
        "kp": kp_alice,
        "verificationMethodId": did_doc_alice_vm["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering of Alice's DID with Id: {did_doc_alice}")

    # Register Bob's DID
    kp_bob = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_bob)
    did_doc_bob = did_doc_string["id"]
    did_doc_bob_vm = did_doc_string["verificationMethod"][0]
    signPair_bob = {
        "kp": kp_bob,
        "verificationMethodId": did_doc_bob_vm["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_bob)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering of Bob's DID with Id: {did_doc_bob}")

    # Alice creates Organization DID with itself being the only controller and Bob's VM being added to VM List
    kp_org = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_org)
    did_doc_org = did_doc_string["id"]
    did_doc_string["controller"] = [did_doc_alice, did_doc_bob]
    did_doc_string["verificationMethod"] = []
    signers.append(signPair_alice)
    signers.append(signPair_bob)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering of Org DID with Id: {did_doc_org}")

    # Deactivate DID
    signers = []
    signers.append(signPair_alice)
    deactivate_tx_cmd = form_did_deactivate_tx_multisig(did_doc_org, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(deactivate_tx_cmd, f"Deactivation of Org's DID with Id: {did_doc_org}")

    print("2. PASS: Mike creates a DID for himself, but the controller list is empty. Mike attempts to deactivate it \n")
    
    # Register Alice's DID
    kp_mike = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_mike)
    did_doc_string["controller"] = []
    did_doc_mike = did_doc_string["id"]
    did_doc_mike_vm = did_doc_string["verificationMethod"][0]
    signPair_mike = {
        "kp": kp_mike,
        "verificationMethodId": did_doc_mike_vm["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_mike)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering of Mike's DID with Id: {did_doc_mike}")


    # Deactivate DID
    signers = []
    signers.append(signPair_mike)
    deactivate_tx_cmd = form_did_deactivate_tx_multisig(did_doc_mike, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(deactivate_tx_cmd, f"Deactivation of Mike's DID with Id: {did_doc_mike}")

    print("--- Test Completed ---\n") 

def schema_test():
    print("\n--- Schema Test ---\n")

    print("1. FAIL: Alice creates a DID with herself being the controller, and then deactivates it. She attempts to registers a schema using one of her VMs\n")
    
    # Register Alice's DID
    kp_alice = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_alice)
    did_doc_alice = did_doc_string["id"]
    did_doc_alice_vm = did_doc_string["verificationMethod"][0]
    signPair_alice = {
        "kp": kp_alice,
        "verificationMethodId": did_doc_alice_vm["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering of Alice's DID with Id: {did_doc_alice}")

    # Deactiving Alice's DID
    signers = []
    signers.append(signPair_alice)
    deactivate_tx_cmd = form_did_deactivate_tx_multisig(did_doc_alice, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(deactivate_tx_cmd, f"Deactivation of Alice's DID with Id: {did_doc_alice}")

    # Register Schema from one of alice's VM Id
    schema_doc, schema_proof = generate_schema_document(
        kp_alice, 
        did_doc_alice, 
        did_doc_alice_vm["id"]
    )
    create_schema_cmd = form_create_schema_tx(
        schema_doc, 
        schema_proof, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    schema_doc_id = schema_doc["id"]
    run_blockchain_command(create_schema_cmd, f"Registering Schema with Id: {schema_doc_id}", True)

    print("2. PASS: Bob creates a DID with herself being the controller. He attempts to registers a schema using one of her VMs\n")
    
    # Register Alice's DID
    kp_bob = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_bob)
    did_doc_bob = did_doc_string["id"]
    did_doc_bob_vm = did_doc_string["verificationMethod"][0]
    signPair_bob = {
        "kp": kp_bob,
        "verificationMethodId": did_doc_bob_vm["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_bob)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering of Bob's DID with Id: {did_doc_bob}")


    # Register Schema from one of alice's VM Id
    schema_doc, schema_proof = generate_schema_document(
        kp_bob, 
        did_doc_bob, 
        did_doc_bob_vm["id"]
    )
    create_schema_cmd = form_create_schema_tx(
        schema_doc, 
        schema_proof, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    schema_doc_id = schema_doc["id"]
    run_blockchain_command(create_schema_cmd, f"Registering Schema with Id: {schema_doc_id}")

    print("--- Test Completed ---\n")

def credential_status_test():
    print("\n--- Credential Status Test ---\n")

    print("1. FAIL: Alice creates a DID with herself being the controller, and then deactivates it. She attempts to registers a credential status using one of her VMs\n")
    
    # Register Alice's DID
    kp_alice = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_alice)
    did_doc_alice = did_doc_string["id"]
    did_doc_alice_vm = did_doc_string["verificationMethod"][0]
    signPair_alice = {
        "kp": kp_alice,
        "verificationMethodId": did_doc_alice_vm["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering of Alice's DID with Id: {did_doc_alice}")

    # Deactiving Alice's DID
    signers = []
    signers.append(signPair_alice)
    deactivate_tx_cmd = form_did_deactivate_tx_multisig(did_doc_alice, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(deactivate_tx_cmd, f"Deactivation of Alice's DID with Id: {did_doc_alice}")

    # Register Credential Status from one of alice's VM Id
    cred_doc, cred_proof = generate_cred_status_document(
        kp_alice,
        did_doc_alice,
        did_doc_alice_vm["id"]
    )
    register_cred_status_cmd = form_create_cred_status_tx(
        cred_doc,
        cred_proof,
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    cred_id = cred_doc["claim"]["id"]
    run_blockchain_command(register_cred_status_cmd, f"Registering Credential status with Id: {cred_id}", True)

    print("2. PASS: Bob creates a DID with herself being the controller. He attempts to registers a credential status using one of his VMs\n")
    
    # Register Alice's DID
    kp_bob = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_bob)
    did_doc_bob = did_doc_string["id"]
    did_doc_bob_vm = did_doc_string["verificationMethod"][0]
    signPair_bob = {
        "kp": kp_bob,
        "verificationMethodId": did_doc_bob_vm["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_bob)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering of Bob's DID with Id: {did_doc_bob}")


    # Register Credential Status from one of alice's VM Id
    cred_doc, cred_proof = generate_cred_status_document(
        kp_bob,
        did_doc_bob,
        did_doc_bob_vm["id"]
    )
    register_cred_status_cmd = form_create_cred_status_tx(
        cred_doc,
        cred_proof,
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    cred_id = cred_doc["claim"]["id"]
    run_blockchain_command(register_cred_status_cmd, f"Registering Credential status with Id: {cred_id}")

    print("--- Test Completed ---\n")

def caip10_ethereum_support_test():
    print("\n--- CAIP-10 Test: Ethereum Chains ---\n")
    kp_algo = "recover-eth"

    # Invalid blockchain Account Ids
    invalid_blockchain_account_ids = [
        "abc345566",
        "eip:1:0x1234",
        "eip155",
        "eip155:1:",
        "eip155::",
        "eip155:1",
        "eip155:::::23",
        "eip155::0x1234567"
        "eip155:1000231432:0x23",
        "eip155:jagrat:0x23"
    ]

    for invalid_blockchain_id in invalid_blockchain_account_ids:
        print("Registering a DID Document with an invalid blockchainAccountId:", invalid_blockchain_id)
        kp = generate_key_pair(algo=kp_algo)
        
        did_doc_string = generate_did_document(kp, kp_algo)
        did_doc_string["verificationMethod"][0]["blockchainAccountId"] = invalid_blockchain_id
        signers = []
        signPair = {
            "kp": kp,
            "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
            "signing_algo": kp_algo
        }
        signers.append(signPair)
        create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
        run_blockchain_command(create_tx_cmd, f"Registering a DID Document with an invalid blockchainAccountId: {invalid_blockchain_id}", True, True)
    
    print("Registering a DID with a VM of type EcdsaSecp256k1SignatureRecovery2020 having publicKeyMultibase attribute populated")
    kp = generate_key_pair(algo=kp_algo)
    did_doc_string = generate_did_document(kp, kp_algo)
    did_doc_id = did_doc_string["id"]
    did_doc_string["verificationMethod"][0]["publicKeyMultibase"] = "zrxxgf1f9xPYTraixqi9tipLta61hp4VJWQUUW5pmwcVz"
    signers = []
    signPair = {
        "kp": kp,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, "Registering a DID with a VM of type EcdsaSecp256k1SignatureRecovery2020 having publicKeyMultibase attribute populated", True, True)

    # Test for valid blockchainAccountId
    kp = generate_key_pair(algo=kp_algo)
    did_doc_string = generate_did_document(kp, kp_algo)
    did_doc_id = did_doc_string["id"]
    print(
        "Registering a DID Document with a valid blockchainAccountId:", 
        did_doc_string["verificationMethod"][0]["blockchainAccountId"]
    )
    signers = []
    signPair = {
        "kp": kp,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_id}")

    print("--- Test Completed ---\n")

def caip10_cosmos_support_test():
    print("\n--- CAIP-10 Test: Cosmos Chains ---\n")

    kp_algo = "secp256k1"

    # Invalid blockchain Account Ids
    invalid_blockchain_account_ids = [
        "abc345566",
        "cos:1:0x1234",
        "cosmos",
        "cosmos:1",
        "cosmos:jagrat",
        "cosmos:1:",
        "cosmos::",
        "cosmos:1",
        "cosmos:::::23",
        "cosmos::0x1234567"
    ]
    print("1. FAIL: Registering a DID Document with an invalid blockchainAccountIds.\n")
    for invalid_blockchain_id in invalid_blockchain_account_ids:
        print("Registering a DID Document with an invalid blockchainAccountId:", invalid_blockchain_id)
        kp = generate_key_pair(algo=kp_algo)
        
        did_doc_string = generate_did_document(kp, kp_algo)
        did_doc_string["verificationMethod"][0]["blockchainAccountId"] = invalid_blockchain_id
        
        signers = []
        signPair = {
            "kp": kp,
            "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
            "signing_algo": kp_algo
        }
        signers.append(signPair)

        create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
        run_blockchain_command(create_tx_cmd, f"Registering a DID Document with an invalid blockchainAccountId: {invalid_blockchain_id}", True, True)
    
    print("2. PASS: Registering a DID with a VM of type EcdsaSecp256k1VerificationKey2019 having both publicKeyMultibase and blockchainAccountId attributes populated")
    kp = generate_key_pair(algo=kp_algo)
    did_doc_string = generate_did_document(kp, kp_algo)
    did_doc_id = did_doc_string["id"]
    signers = []
    signPair = {
        "kp": kp,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering the DID with DID Id {did_doc_id}")

    print("3. FAIL: Registering a DID with invalid chain-id in blockchainAccountId")
    kp = generate_key_pair(algo=kp_algo)
    did_doc_string = generate_did_document(kp, kp_algo)
    did_doc_id = did_doc_string["id"]
    blockchainAccountId = secp256k1_pubkey_to_address(kp["pub_key_base_64"], "hid")
    did_doc_string["verificationMethod"][0]["blockchainAccountId"] = "cosmos:hidnode02:" + blockchainAccountId

    signers = []
    signPair = {
        "kp": kp,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_id}", True, True)

    print("4. PASS: Registering a DID with two VM of type EcdsaSecp256k1VerificationKey2019 with duplicate publicKeyMultibase and different blockchain account Id")
    kp = generate_key_pair(algo=kp_algo)

    did_doc_string_1 = generate_did_document(kp, kp_algo)
    did_doc_string_2 = generate_did_document(kp, kp_algo, "osmo")
    did_doc_string_2_vm = did_doc_string_2["verificationMethod"][0]
    did_doc_string_2_vm["id"] = did_doc_string_2_vm["id"] + "new"
    did_doc_string_2_vm["controller"] = did_doc_string_1["verificationMethod"][0]["controller"]

    did_doc_string_1["verificationMethod"] = [
        did_doc_string_1["verificationMethod"][0],
        did_doc_string_2_vm,
    ]
    did_doc_id = did_doc_string_1["id"]
    
    signers = []
    signPair1 = {
        "kp": kp,
        "verificationMethodId": did_doc_string_1["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signPair2 = {
        "kp": kp,
        "verificationMethodId": did_doc_string_2["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair1)
    signers.append(signPair2)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string_1, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering the DID with DID Id {did_doc_id}")

    print("5. PASS: Registering a DID with two VM of type EcdsaSecp256k1VerificationKey2019 with duplicate publicKeyMultibase but one of them is without a blockchain account id")
    kp = generate_key_pair(algo=kp_algo)

    did_doc_string_1 = generate_did_document(kp, kp_algo)
    did_doc_string_2 = generate_did_document(kp, kp_algo)
    did_doc_string_2_vm = did_doc_string_2["verificationMethod"][0]
    did_doc_string_2_vm["id"] = did_doc_string_2_vm["id"] + "new"

    #Remove blockchainAccountIds
    did_doc_string_1["verificationMethod"][0]["blockchainAccountId"] = ""

    did_doc_string_1["verificationMethod"] = [
        did_doc_string_1["verificationMethod"][0],
        did_doc_string_2_vm,
    ]
    did_doc_id = did_doc_string_1["id"]
    
    signers = []
    signPair1 = {
        "kp": kp,
        "verificationMethodId": did_doc_string_1["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signPair2 = {
        "kp": kp,
        "verificationMethodId": did_doc_string_2["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair1)
    signers.append(signPair2)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string_1, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering the DID with DID Id {did_doc_id}")

    print("6. FAIL: Registering a DID with two VM of type EcdsaSecp256k1VerificationKey2019 with duplicate publicKeyMultibase and duplicate blockchainAccountId")
    kp = generate_key_pair(algo=kp_algo)

    did_doc_string_1 = generate_did_document(kp, kp_algo)
    did_doc_string_2 = generate_did_document(kp, kp_algo)

    did_doc_vm1 = did_doc_string_1["verificationMethod"][0]
    did_doc_vm2 = did_doc_string_2["verificationMethod"][0]

    # Change vm id
    did_doc_vm1["id"] = did_doc_vm1["id"] + "news"
    did_doc_vm2["id"] = did_doc_vm1["id"] + "2" 

    
    did_doc_string_1["verificationMethod"] = [
        did_doc_vm1,
        did_doc_vm2
    ]
    did_doc_id = did_doc_string_1["id"]
    
    signers = []
    signPair1 = {
        "kp": kp,
        "verificationMethodId": did_doc_string_1["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signPair2 = {
        "kp": kp,
        "verificationMethodId": did_doc_string_1["verificationMethod"][1]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair1)
    signers.append(signPair2)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string_1, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering the DID with DID Id {did_doc_id}", True, True)

    print("7. FAIL: Registering a DID with two VM of type EcdsaSecp256k1VerificationKey2019 with duplicate publicKeyMultibase and no blockchainAccountId in either of them")
    kp = generate_key_pair(algo=kp_algo)

    did_doc_string_1 = generate_did_document(kp, kp_algo)
    did_doc_string_2 = generate_did_document(kp, kp_algo)

    did_doc_vm1 = did_doc_string_1["verificationMethod"][0]
    did_doc_vm2 = did_doc_string_2["verificationMethod"][0]

    # Change vm id
    did_doc_vm1["id"] = did_doc_vm1["id"] + "news"
    did_doc_vm2["id"] = did_doc_vm1["id"] + "2" 

    # Remove blockchainAccountId
    did_doc_vm1["blockchainAccountId"] = ""
    did_doc_vm2["blockchainAccountId"] = ""
    
    did_doc_string_1["verificationMethod"] = [
        did_doc_vm1,
        did_doc_vm2
    ]
    did_doc_id = did_doc_string_1["id"]
    
    signers = []
    signPair1 = {
        "kp": kp,
        "verificationMethodId": did_doc_string_1["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signPair2 = {
        "kp": kp,
        "verificationMethodId": did_doc_vm2["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair1)
    signers.append(signPair2)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string_1, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering the DID with DID Id {did_doc_id}", True, True)

    print("--- Test Completed ---\n")

def vm_type_test():
    print("\n--- Verification Method Types Test ---\n")

    # Ed25519VerificationKey2020
    print("1. FAIL: Registering DID Document with a verification method of type Ed25519VerificationKey2020. Both publicKeyMultibase and blockchainAccountId are passed.")
    kp_alice = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_alice)
    did_doc_alice = did_doc_string["id"]
    did_doc_string["verificationMethod"][0]["blockchainAccountId"] = "solana:4sGjMW1sUnHzSxGspuhpqLDx6wiyjNtZ:7S3P4HxJpyyigGzodYwHtCxZyUQe9JiBMHyRWXArAaKv"
    signPair_alice = {
        "kp": kp_alice,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_alice}", True, True)

    print("2. FAIL: Registering DID Document with a verification method of type Ed25519VerificationKey2020. Only blockchainAccountId is passed.")
    kp_alice = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_alice)
    did_doc_alice = did_doc_string["id"]
    did_doc_string["verificationMethod"][0]["publicKeyMultibase"] = ""
    did_doc_string["verificationMethod"][0]["blockchainAccountId"] = "solana:4sGjMW1sUnHzSxGspuhpqLDx6wiyjNtZ:7S3P4HxJpyyigGzodYwHtCxZyUQe9JiBMHyRWXArAaKv"
    signPair_alice = {
        "kp": kp_alice,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_alice}", True, True)

    print("3. PASS: Registering DID Document with a verification method of type Ed25519VerificationKey2020. Only publicKeyMultibase is passed.")
    kp_alice = generate_key_pair()
    signers = []
    did_doc_string = generate_did_document(kp_alice)
    did_doc_alice = did_doc_string["id"]
    signPair_alice = {
        "kp": kp_alice,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": "ed25519"
    }
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_alice}")

    # EcdsaSecp256k1VerificationKey2019
    print("4. PASS: Registering DID Document with a verification method of type EcdsaSecp256k1VerificationKey2019. Only publicKeyMultibase is passed.")
    kp_algo = "secp256k1"
    kp = generate_key_pair(algo=kp_algo)
    did_doc_string = generate_did_document(kp, kp_algo, is_uuid=True, bech32prefix="")
    did_doc_id = did_doc_string["id"]
    did_doc_string["verificationMethod"][0]["blockchainAccountId"] = ""
    signers = []
    signPair = {
        "kp": kp,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_id}")

    print("5. PASS: Registering DID Document with a verification method of type EcdsaSecp256k1VerificationKey2019. Both publicKeyMultibase and blockchainAccountId are passed.")
    kp_algo = "secp256k1"
    kp = generate_key_pair(algo=kp_algo)
    did_doc_string = generate_did_document(kp, kp_algo)
    did_doc_id = did_doc_string["id"]
    signers = []
    signPair = {
        "kp": kp,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_id}")

    print("6. FAIL: Registering DID Document with a verification method of type EcdsaSecp256k1VerificationKey2019. Only blockchainAccountId is passed.")
    kp_algo = "secp256k1"
    kp = generate_key_pair(algo=kp_algo)
    did_doc_string = generate_did_document(kp, kp_algo)
    did_doc_id = did_doc_string["id"]
    did_doc_string["verificationMethod"][0]["publicKeyMultibase"] = ""
    signers = []
    signPair = {
        "kp": kp,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_id}", True, True)

    # EcdsaSecp256k1RecoveryMethod2020
    print("7. FAIL: Registering DID Document with a verification method of type EcdsaSecp256k1RecoveryMethod2020. Only publicKeyMultibase is passed.")
    kp_algo = "recover-eth"
    kp = generate_key_pair(algo=kp_algo)
    did_doc_string = generate_did_document(kp, kp_algo)
    did_doc_id = did_doc_string["id"]
    did_doc_string["verificationMethod"][0]["publicKeyMultibase"] = "z22XxPVrzTr24zUBGfZiZCm9bwj3RkHKLmx9ENYBxmY57o"
    did_doc_string["verificationMethod"][0]["blockchainAccountId"] = ""
    signers = []
    signPair = {
        "kp": kp,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_id}", True, True)

    print("8. FAIL: Registering DID Document with a verification method of type EcdsaSecp256k1RecoveryMethod2020. Both publicKeyMultibase and blockchainAccountId is passed.")
    kp_algo = "recover-eth"
    kp = generate_key_pair(algo=kp_algo)
    did_doc_string = generate_did_document(kp, kp_algo)
    did_doc_id = did_doc_string["id"]
    did_doc_string["verificationMethod"][0]["publicKeyMultibase"] = "z22XxPVrzTr24zUBGfZiZCm9bwj3RkHKLmx9ENYBxmY57o"
    signers = []
    signPair = {
        "kp": kp,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_id}", True, True)

    print("9. PASS: Registering DID Document with a verification method of type EcdsaSecp256k1RecoveryMethod2020. Only blockchainAccountId is passed.")
    kp_algo = "recover-eth"
    kp = generate_key_pair(algo=kp_algo)
    did_doc_string = generate_did_document(kp, kp_algo)
    did_doc_id = did_doc_string["id"]
    signers = []
    signPair = {
        "kp": kp,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_id}")

    print("--- Test Completed ---\n")
    
def method_specific_id_test():
    print("\n--- Method Specific ID Tests ---\n")

    print("1. PASS: Registering a DID Document where the user provides a blockchain address in MSI that they own")

    kp_algo = "secp256k1"
    kp_alice = generate_key_pair(algo=kp_algo)
    signers = []
    did_doc_string = generate_did_document(kp_alice, algo=kp_algo)
    did_doc_alice = did_doc_string["id"]
    did_doc_alice_vm = did_doc_string["verificationMethod"]
    signPair_alice = {
        "kp": kp_alice,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering Alice's DID with Id: {did_doc_alice}")

    print("2. FAIL: Registering a DID Document where the user provides a blockchain address in MSI that they don't own")
    
    kp_algo = "secp256k1"
    kp_bob = generate_key_pair(algo=kp_algo)
    signers = []
    did_doc_string = generate_did_document(kp_bob, algo=kp_algo)
    did_doc_string["controller"] = [did_doc_alice]
    did_doc_string["verificationMethod"] = did_doc_alice_vm
    
    did_doc_bob = did_doc_string["id"]
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering Bob's DID with Id: {did_doc_bob}", True)

    print("3. PASS: Registering a DID Document where the user provides a multibase encoded public key in MSI that they own")

    kp_algo = "ed25519"
    kp_alice = generate_key_pair(algo=kp_algo)
    signers = []
    did_doc_string = generate_did_document(kp_alice, algo=kp_algo)
    did_doc_alice = did_doc_string["id"]
    did_doc_alice_vm = did_doc_string["verificationMethod"]
    signPair_alice = {
        "kp": kp_alice,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering Alice's DID with Id: {did_doc_alice}")

    print("4. PASS: Registering a DID Document where the user provides a multibase encoded public key in MSI that they don't own")
    
    kp_algo = "ed25519"
    kp_bob = generate_key_pair(algo=kp_algo)
    signers = []
    did_doc_string = generate_did_document(kp_bob, algo=kp_algo)
    did_doc_string["controller"] = [did_doc_alice]
    did_doc_string["verificationMethod"] = did_doc_alice_vm
    
    did_doc_bob = did_doc_string["id"]
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering Bob's DID with Id: {did_doc_bob}")

    print("5. FAIL: Attempt to Register Invalid DID Documents with invalid DID Id")

    did_id_list = [
        "did:hid:1:",
        "did:hid:1",
        "did:hid:devnet",
        "did:hid:devnet:",
        "did:hid:devnet:zHiii",
        "did:hid:devnet:asa54qf",
        "did:hid:devnet:asa54qf|sds",
        "did:hid:devnet:-asa54-qfsds",
        "did:hid:devnet:.com",
    ]

    for did_id in did_id_list:
        did_doc_string["id"] = did_id
        create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
        run_blockchain_command(create_tx_cmd, f"Registering Invalid DID with Id: {did_id}", True, True)        

    print("6. PASS: Alice tries to update their DID Document by removing the Verification Method associated with the method specific id (CAIP-10 Blockchain Address)")

    kp_algo = "secp256k1"
    kp_alice = generate_key_pair(algo=kp_algo)
    signers = []
    did_doc_string = generate_did_document(kp_alice, algo=kp_algo)
    did_doc_alice = did_doc_string["id"]
    did_doc_alice_vm = did_doc_string["verificationMethod"]
    signPair_alice = {
        "kp": kp_alice,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering Alice's DID with Id: {did_doc_alice}")
    
    did_doc_string["verificationMethod"] = []
    update_tx_cmd = form_did_update_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Removal of Verification Method associated with method specific id")

    print("7. PASS: Alice tries to update their DID Document by removing the Verification Method associated with the method specific id (Multibase Encoded PublicKey)")

    kp_algo = "ed25519"
    kp_alice = generate_key_pair(algo=kp_algo)
    signers = []
    did_doc_string = generate_did_document(kp_alice, algo=kp_algo)
    did_doc_alice = did_doc_string["id"]
    did_doc_alice_vm = did_doc_string["verificationMethod"]
    signPair_alice = {
        "kp": kp_alice,
        "verificationMethodId": did_doc_string["verificationMethod"][0]["id"],
        "signing_algo": kp_algo
    }
    signers.append(signPair_alice)
    create_tx_cmd = form_did_create_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering Alice's DID with Id: {did_doc_alice}")

    did_doc_string["verificationMethod"] = []
    update_tx_cmd = form_did_update_tx_multisig(did_doc_string, signers, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(update_tx_cmd, f"Removal of Verification Method associated with method specific id")

    print("--- Test Completed ---\n")