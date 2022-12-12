import os
import sys
sys.path.insert(1, os.getcwd())

import json

from utils import run_blockchain_command, generate_key_pair
from generate_doc import generate_did_document, \
    generate_schema_document, generate_cred_status_document
from transactions import form_create_schema_tx, form_did_create_tx, \
    form_create_cred_status_tx, form_did_deactivate_tx, form_did_update_tx, query_did
from constants import DEFAULT_BLOCKCHAIN_ACCOUNT_NAME

def deactivated_did_should_not_create_ssi_elements():
    print("\n---- Test: Deactivated DID attempting to register Schema and Credential Status Docs ----\n")
    print("In this workflow, a deactivated DID attempts to register Schema Doc and Credential Status Doc")
    print("It is expected to fail\n")

    print("Registering a DID Document")
    kp_did = generate_key_pair()
    did_doc = generate_did_document(kp_did)
    create_did_tx_cmd = form_did_create_tx(did_doc, kp_did, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    did_doc_id = did_doc["id"]
    run_blockchain_command(create_did_tx_cmd, f"Registering DID with Id: {did_doc_id}")

    print("Deactivating the Registered DID Document")
    deactivate_did_tx_cmd = form_did_deactivate_tx(
        did_doc_id, 
        kp_did, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    run_blockchain_command(deactivate_did_tx_cmd, f"Deactivating DID with Id: {did_doc_id}")

    print("Deactivated DID attempts to register Schema Document")
    schema_doc, schema_proof = generate_schema_document(
        kp_did, 
        did_doc_id, 
        did_doc["authentication"][0]
    )
    create_schema_cmd = form_create_schema_tx(
        schema_doc, 
        schema_proof, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    schema_doc_id = schema_doc["id"]
    run_blockchain_command(create_schema_cmd, f"Registering Schema with Id: {schema_doc_id} with {did_doc_id} being the author", True)

    print("Deactivated DID attempts to register Credential Status Document")
    cred_doc, cred_proof = generate_cred_status_document(
        kp_did,
        did_doc_id,
        did_doc["authentication"][0]
    )
    register_cred_status_cmd = form_create_cred_status_tx(
        cred_doc,
        cred_proof,
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    cred_id = cred_doc["claim"]["id"]
    run_blockchain_command(register_cred_status_cmd, f"Registering Credential status with Id: {cred_id} with {did_doc_id} being the issuer", True)
    print("\n----Test Completed: Deactivated DID attempting to register Schema and Credential Status Docs------\n")

def multiple_controllers_with_one_signer():
    print("\n--- Test: Multiple Controllers with One Signer ---\n")
    print("In this worflow, two DIDs are added to another DID's Controller Group. One of them attempts to make change in the parent DID Document")

    print("Registering Parent DID Document")
    kp_parent_did = generate_key_pair()
    parent_did_doc = generate_did_document(kp_parent_did)
    create_did_tx_cmd = form_did_create_tx(parent_did_doc, kp_parent_did, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    parent_did_doc_id = parent_did_doc["id"]
    run_blockchain_command(create_did_tx_cmd, f"Registering DID with Id: {parent_did_doc_id}")

    print("Registering 1st Controller DID Document")
    kp_controller_1_did = generate_key_pair()
    controller_1_did_doc = generate_did_document(kp_controller_1_did)
    create_did_tx_cmd = form_did_create_tx(controller_1_did_doc, kp_controller_1_did, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    controller_1_did_doc_id = controller_1_did_doc["id"]
    run_blockchain_command(create_did_tx_cmd, f"Registering DID with Id: {controller_1_did_doc_id}")

    print("Registering 2nd Controller DID Document")
    kp_controller_2_did = generate_key_pair(algo="secp256k1")
    controller_2_did_doc = generate_did_document(kp_controller_2_did, algo="secp256k1")
    create_did_tx_cmd = form_did_create_tx(controller_2_did_doc, kp_controller_2_did, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME, signing_algo="secp256k1")
    controller_2_did_doc_id = controller_2_did_doc["id"]
    run_blockchain_command(create_did_tx_cmd, f"Registering DID with Id: {controller_2_did_doc_id}")

    print("Adding 1st and 2nd DID Controllers to Parent DID's Controller group")
    parent_did_doc = query_did(parent_did_doc_id)["didDocument"]
    parent_did_doc["controller"] = [controller_1_did_doc_id, controller_2_did_doc_id]
    update_did_tx_cmd = form_did_update_tx(
        parent_did_doc, 
        kp_parent_did, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    run_blockchain_command(update_did_tx_cmd, f"Adding 1st and 2nd Controller DIDs to Parent DID's Control Group")

    print("2nd Controller trying to make a change in Parent DID")
    parent_did_doc = query_did(parent_did_doc_id)["didDocument"]
    parent_did_doc["context"] = ["websitebycontroller2.com"]
    update_did_tx_cmd = form_did_update_tx(
        parent_did_doc, 
        kp_controller_2_did, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME,
        controller_2_did_doc["authentication"][0],
        "secp256k1"
    )
    run_blockchain_command(update_did_tx_cmd, f"2nd Controller trying to update Parent DID")

    print("Parent DID trying to update it's DID")
    parent_did_doc = query_did(parent_did_doc_id)["didDocument"]
    parent_did_doc["context"] = ["websitebyparent.com"]
    update_did_tx_cmd = form_did_update_tx(
        parent_did_doc, 
        kp_parent_did, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    run_blockchain_command(update_did_tx_cmd, f"Parent DID trying to update it's own DID Document", True)
    print("\n----Test Completed: Multiple Controllers with One Signer-------\n")

def controller_did_trying_to_update_diddoc():
    print("\n--- Test: Controller DID attempts to change Canon DID ---\n")
    print("In this workflow, Controller DID attempts to update its Parent DID Document")

    print("Registering the canon DID Document")
    kp_canon_did = generate_key_pair()
    canon_did_doc = generate_did_document(kp_canon_did)
    canon_did_doc["controller"] = [canon_did_doc["id"]]
    create_did_tx_cmd = form_did_create_tx(canon_did_doc, kp_canon_did, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    canon_did_doc_id = canon_did_doc["id"]
    run_blockchain_command(create_did_tx_cmd, f"Registering DID with Id: {canon_did_doc_id}")
    
    print("Registering the controller DID Document")
    kp_controller_did = generate_key_pair()
    controller_did_doc = generate_did_document(kp_controller_did)
    create_did_tx_cmd = form_did_create_tx(
        controller_did_doc, 
        kp_controller_did, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    controller_did_doc_id = controller_did_doc["id"]
    run_blockchain_command(create_did_tx_cmd, f"Registering DID with Id: {controller_did_doc_id}")
    
    print("Adding the registered controller DID to Canon DID Document's controller group")
    canon_did_doc = query_did(canon_did_doc_id)["didDocument"]
    canon_did_doc["controller"] = [canon_did_doc_id, controller_did_doc_id]
    update_did_tx_cmd = form_did_update_tx(
        canon_did_doc, 
        kp_canon_did, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    run_blockchain_command(update_did_tx_cmd, f"Adding Controller DID to Canon DID's Control Group")

    print("Controller DID trying to update canon DID Document")
    canon_did_doc = query_did(canon_did_doc_id)["didDocument"]
    canon_did_doc["context"] = ["somenewwebsite.com"]
    update_did_tx_cmd = form_did_update_tx(
        canon_did_doc, 
        kp_controller_did, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME,
        controller_did_doc["authentication"][0]
    )
    run_blockchain_command(update_did_tx_cmd, f"Attempt by controller to make changes in Canon DID")
    print("\n------ Test Completed ---------\n")

def non_controller_did_trying_to_update_diddoc():
    print("\n--- Test: Non Controller DID attempts to change Canon DID ---\n")
    print("In this workflow, Non-Controller DID attempts to update a DID Document")
    print("This is expected to fail\n")

    print("Registering the canon DID Document")
    kp_canon_did = generate_key_pair()
    canon_did_doc = generate_did_document(kp_canon_did)
    create_did_tx_cmd = form_did_create_tx(canon_did_doc, kp_canon_did, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    canon_did_doc_id = canon_did_doc["id"]
    run_blockchain_command(create_did_tx_cmd, f"Registering DID with Id: {canon_did_doc_id}")

    print("Registering the non controller DID Document")
    kp_non_controller_did = generate_key_pair()
    non_controller_did_doc = generate_did_document(kp_non_controller_did)
    create_did_tx_cmd = form_did_create_tx(
        non_controller_did_doc, 
        kp_non_controller_did, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    non_controller_did_doc_id = non_controller_did_doc["id"]
    run_blockchain_command(create_did_tx_cmd, f"Registering DID with Id: {non_controller_did_doc_id}")
    
    print("Non Controller DID trying to update canon DID Document")
    canon_did_doc["context"] = canon_did_doc["context"].append("newwebsite.com")
    update_did_tx_cmd = form_did_update_tx(
        canon_did_doc, 
        kp_non_controller_did, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME,
        non_controller_did_doc["authentication"][0]
    )
    run_blockchain_command(update_did_tx_cmd, f"Attempt by non-controller to make changes in Canon DID", True)
    print("\n------ Test Completed ---------\n")

def simple_ssi_flow():
    print("\n--- Test: Simple SSI Worflow ---\n")
    print("In this workflow, a DID document is registered, following which a credential schema and a credential status document is registered\n")
    
    kp = generate_key_pair()

    print("Registering a DID Document")
    did_doc_string = generate_did_document(kp)
    did_doc_id = did_doc_string["id"]
    create_tx_cmd = form_did_create_tx(did_doc_string, kp, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_id}")

    print("Registering a Schema Document")
    schema_doc, schema_proof = generate_schema_document(
        kp, 
        did_doc_id, 
        did_doc_string["authentication"][0]
    )
    create_schema_cmd = form_create_schema_tx(
        schema_doc, 
        schema_proof, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    schema_doc_id = schema_doc["id"]
    run_blockchain_command(create_schema_cmd, f"Registering Schema with Id: {schema_doc_id}")

    print("Registering a Credential Status Document")
    cred_doc, cred_proof = generate_cred_status_document(
        kp,
        did_doc_id,
        did_doc_string["authentication"][0]
    )
    register_cred_status_cmd = form_create_cred_status_tx(
        cred_doc,
        cred_proof,
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    cred_id = cred_doc["claim"]["id"]
    run_blockchain_command(register_cred_status_cmd, f"Registering Credential status with Id: {cred_id}")
    print("\n------ Test Completed ---------\n")

def controller_creates_schema_cred_status():
    print("--- Test: Schema and Credential Status document registration by Controllers\n")
    print("In this workflow, a DID document registered with another DID Id in its controller group. The controller is expected to register schema and credential status")
    
    print("Registering DID for an Employee")
    employee_kp = generate_key_pair()
    employee_did = generate_did_document(employee_kp)
    employee_did_id = employee_did["id"]
    create_employee_did_tx = form_did_create_tx(employee_did, employee_kp, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_employee_did_tx, f"Registering Employee DID Document with ID {employee_did_id}")

    print("Registering DID for an Organization")
    org_kp = generate_key_pair()
    org_did = generate_did_document(org_kp)
    org_did["controller"] = [employee_did_id]
    org_did_id = org_did["id"]
    create_org_did_tx = form_did_create_tx(org_did, employee_kp, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME, employee_did["authentication"][0])
    run_blockchain_command(create_org_did_tx, f"Registering Organisation DID Document with ID {org_did_id}")

    print("Employee registering a Schema on behalf of Organization's DID")
    schema_doc, schema_proof = generate_schema_document(
        employee_kp, 
        org_did_id, 
        employee_did["authentication"][0]
    )
    create_schema_cmd = form_create_schema_tx(
        schema_doc, 
        schema_proof, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    schema_doc_id = schema_doc["id"]
    schema_author = schema_doc["author"]
    run_blockchain_command(create_schema_cmd, f"Registering Schema with Id: {schema_doc_id} with {schema_author} being the author")

    print("Employee registering a Credential Status Document on behalf of Organization's DID")
    cred_doc, cred_proof = generate_cred_status_document(
        employee_kp, 
        org_did_id, 
        employee_did["authentication"][0]
    )
    register_cred_status_cmd = form_create_cred_status_tx(
        cred_doc,
        cred_proof, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    cred_id = cred_doc["claim"]["id"]
    cred_author = cred_doc["issuer"]
    run_blockchain_command(register_cred_status_cmd, f"Registering credential status with Id: {cred_id} and {cred_author} being the author")
    print("\n------ Test Completed ---------\n")

def invalid_case_controller_creates_schema_cred_status():
    print("--- Test: Invalid Schema and Credential Status document registration by Non Controllers\n")
    print("In this workflow, a DID document registered with another DID Id in its controller group. In this case, if the canon DID tries to create Schema or Credential Document, it should fail.\n")
    
    print("Registering DID for an Employee")
    employee_kp = generate_key_pair()
    employee_did = generate_did_document(employee_kp)
    employee_did_id = employee_did["id"]
    create_employee_did_tx = form_did_create_tx(employee_did, employee_kp, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME)
    run_blockchain_command(create_employee_did_tx, f"Registering Employee DID Document with ID {employee_did_id}")

    print("Registering DID for an Organization")
    org_kp = generate_key_pair()
    org_did = generate_did_document(org_kp)
    org_did["controller"] = [employee_did_id]
    org_did_id = org_did["id"]
    create_org_did_tx = form_did_create_tx(org_did, employee_kp, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME, employee_did["authentication"][0])
    run_blockchain_command(create_org_did_tx, f"Registering Organisation DID Document with ID {org_did_id}")

    print("\nAttempting to register Schema using Organization's cryptographic material")
    print("This is expected to fail")
    schema_doc, schema_proof = generate_schema_document(
        org_kp, 
        org_did_id,
        org_did["authentication"][0]
    )
    create_schema_cmd = form_create_schema_tx(
        schema_doc, 
        schema_proof, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    schema_doc_id = schema_doc["id"]
    schema_author = schema_doc["author"]
    run_blockchain_command(create_schema_cmd, f"Registering Schema with Id: {schema_doc_id} with {schema_author} being the author\n", True)

    print("\nAttempting to register Credential Status using Organization's cryptographic material")
    print("This is expected to fail")
    cred_doc, cred_proof = generate_cred_status_document(
        org_kp, 
        org_did_id, 
        org_did["authentication"][0]
    )
    register_cred_status_cmd = form_create_cred_status_tx(
        cred_doc,
        cred_proof, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    cred_id = cred_doc["claim"]["id"]
    cred_author = cred_doc["issuer"]
    run_blockchain_command(register_cred_status_cmd, f"Registering credential status with Id: {cred_id} with {cred_author} being the issuer\n", True)

    print("\n-------Test Completed------\n")

def did_operations_using_secp256k1():
    print("\n--- Test: DID Operations using Secp256k1 Key Pair ---\n")

    kp_algo = "secp256k1"
    kp = generate_key_pair(algo=kp_algo)

    print("Registering a DID Document")
    did_doc_string = generate_did_document(kp, kp_algo)
    did_doc_id = did_doc_string["id"]
    create_tx_cmd = form_did_create_tx(did_doc_string, kp, DEFAULT_BLOCKCHAIN_ACCOUNT_NAME, signing_algo=kp_algo)
    run_blockchain_command(create_tx_cmd, f"Registering DID with Id: {did_doc_id}")

    print("Registering a Schema Document")
    schema_doc, schema_proof = generate_schema_document(
        kp, 
        did_doc_id, 
        did_doc_string["authentication"][0],
        algo = kp_algo
    )
    create_schema_cmd = form_create_schema_tx(
        schema_doc, 
        schema_proof, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    schema_doc_id = schema_doc["id"]
    run_blockchain_command(create_schema_cmd, f"Registering Schema with Id: {schema_doc_id}")

    print("Registering a Credential Status Document")
    cred_doc, cred_proof = generate_cred_status_document(
        kp,
        did_doc_id,
        did_doc_string["authentication"][0],
        algo = kp_algo
    )
    register_cred_status_cmd = form_create_cred_status_tx(
        cred_doc,
        cred_proof,
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME
    )
    cred_id = cred_doc["claim"]["id"]
    run_blockchain_command(register_cred_status_cmd, f"Registering Credential status with Id: {cred_id}")

    print("Updating the Registered DID Document")
    registered_did_doc = query_did(did_doc_id)["didDocument"]
    registered_did_doc["context"] = registered_did_doc["context"].append("newwebsite.com")
    update_did_tx_cmd = form_did_update_tx(
        registered_did_doc, 
        kp, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME,
        signing_algo=kp_algo
    )
    run_blockchain_command(update_did_tx_cmd, f"Updating DID Document with Id: {did_doc_id}")

    print("Deactivating the Registered DID Document")
    deactivate_did_tx_cmd = form_did_deactivate_tx(
        did_doc_id, 
        kp, 
        DEFAULT_BLOCKCHAIN_ACCOUNT_NAME,
        signing_algo=kp_algo
    )
    run_blockchain_command(deactivate_did_tx_cmd, f"Deactivating DID Document with Id: {did_doc_id}")

    print("\n-------Test Completed------\n")