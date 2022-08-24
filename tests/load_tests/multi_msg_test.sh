#!/bin/bash

SOURCE_ADDRESS=$(hid-noded keys show node1 -a)

echo "Signing the batch txs"
echo ""
hid-noded tx sign ./sample_bank_tx.json --from ${SOURCE_ADDRESS} --chain-id hidnode > signed_tx.json
echo ""
echo "Transaction Signed"

echo "Broadcasting Batch Tx"
echo ""
hid-noded tx broadcast ./signed_tx.json --broadcast-mode block
echo ""
echo "Batch Tx Broadcasted"

rm -rf signed_tx.json