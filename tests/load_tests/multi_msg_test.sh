#!/bin/bash

echo "-------- Starting Test For Multiple Messages in a Tx ---------"
echo ""

SOURCE_ADDRESS=$(vid-noded keys show node1 -a)
DESTINATION_ADDRESS="vid17h4zk54wlayhfla7pdxxfqej3ra0rx5j6305hg"
BROADCAST=""

check_tx() {
    if [ ${1} -ne 0 ]; then
        echo "Transaction failed"
        echo ""
        echo "Code: ${1}"
        echo ""
        echo "Log: ${2}"
        rm -rf signed_tx.json
    else
        echo "Transaction has been broadcasted"
        echo ""
        rm -rf signed_tx.json
    fi
}

sign_and_broadcast() {
    # Signing the Tx
    echo "Signing the batch txs"
    echo ""
    vid-noded tx sign ./multiple_bank_msgs.json --from ${1} --chain-id vidnode > signed_tx.json
    echo ""
    echo "Transaction Signed"

    # Broadcasting the Tx
    echo "Broadcasting Batch Tx"
    echo ""
    BROADCAST=$(vid-noded tx broadcast ./signed_tx.json --broadcast-mode block --output json)
}

# Generate Tx
rm -rf ./multiple_bank_msgs.json
vid-noded tx bank send ${SOURCE_ADDRESS} ${DESTINATION_ADDRESS} 10uvid --chain-id vidnode --generate-only > multiple_bank_msgs.json

# Replicating Valid Transactions
echo "Replicating Valid Transactions"
echo ""
python3 add_multiple_msgs.py

sign_and_broadcast ${SOURCE_ADDRESS}
CODE=$(echo ${BROADCAST} | jq '.code')
TX_LOG=$(echo ${BROADCAST} | jq '.raw_log')

echo "Check the Transaction Status: It is expected to pass"
check_tx ${CODE} ${TX_LOG}

# Making a transaction invalid
echo "Making a transaction invalid"
python3 add_multiple_msgs.py "include_invalid_tx"
sign_and_broadcast ${SOURCE_ADDRESS}
CODE=$(echo ${BROADCAST} | jq '.code')
TX_LOG=$(echo ${BROADCAST} | jq '.raw_log')

echo "Checking the Transaction Status: It is expected to fail"
check_tx ${CODE} ${TX_LOG}

echo "----------------- Test Completed --------------------"