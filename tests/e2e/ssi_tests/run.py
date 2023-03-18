import sys
import contextlib
import os

from e2e_tests import *
from utils import is_blockchain_active

def generate_report(func):
    print("Generating reports in artifacts directory...")
    try:
        report_dir = os.getcwd() + "/artifacts"
        report_name = "e2e_ssi_module_test_report.txt"
        report_path = report_dir + "/" + report_name

        if not os.path.exists(report_dir):
            os.makedirs(report_dir)

        with open(report_path, 'w') as f, contextlib.redirect_stdout(f):
            func()
        print("Test report is generated.")
    except Exception as e:
        print("Test report generation failed\n", e)

def run_all_tests():
    print("============= 🔧️ Running all x/ssi e2e tests ============== \n")
    
    # create_did_test()
    update_did_test()
    # schema_test()
    # deactivate_did()
    # credential_status_test()
    # caip10_ethereum_support_test()
    # caip10_cosmos_support_test()
    # vm_type_test()
    #run_something()

    print("============= 😃️ All test cases completed successfully ============== \n")

if __name__=='__main__':
    # Assert if blockchain is currently running
    is_blockchain_active(rpc_port=26657)

    try:
        if len(sys.argv) > 1 and sys.argv[1] == "generate":
            generate_report(run_all_tests)
        else:
            run_all_tests()

    except Exception as e:
        raise(e)
