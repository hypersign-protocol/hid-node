import json
import sys

if __name__=='__main__':
    bank_tx = None
    messages = []
    num_txs = 10
    data = None

    tx_filepath = "./multiple_bank_msgs.json"
    with open(tx_filepath, "r") as read_file:
        data = json.load(read_file)
        tx_body = data["body"]
        
        messages = tx_body["messages"]
        bank_tx = messages[0]

    for i in range(1, num_txs):
        messages.append(bank_tx)
    

    # Check for include_invalid_tx cli argument
    if len(sys.argv)==2 and sys.argv[1] == "include_invalid_tx":
        messages[1]["amount"][0]["amount"] = "100000000000000000000"

    data["body"]["messages"] = messages

    with open(tx_filepath, "w") as write_file:
        json_data = json.dumps(data)
        write_file.write(json_data)

    print("Multiple transactions have been added\n")
