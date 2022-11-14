from e2e_tests import *
from utils import is_blockchain_active

def run_all_tests():
    print("Running all SSI Related E2E tests\n")
    
    simple_ssi_flow()
    controller_creates_schema_cred_status()
    invalid_case_controller_creates_schema_cred_status()
    non_controller_did_trying_to_update_diddoc()
    controller_did_trying_to_update_diddoc()
    multiple_controllers_with_one_signer()
    deactivated_did_should_not_create_ssi_elements()

if __name__=='__main__':
    # Assert if blockchain is currently running
    is_blockchain_active(rpc_port=26657)

    try:
        # If you want to run a handful of tests,
        # comment out run_all() function and
        # mention the test functions here
        run_all_tests()
    except Exception as e:
        raise(e)
