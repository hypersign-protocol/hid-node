#!/bin/bash

# Submit the proposal
echo "Submitting the proposal"
echo ""
vid-noded tx gov submit-proposal community-pool-spend ./community-pool-spend-proposal.json --from node1 --yes
sleep 7
echo ""
echo "Proposal is submitted"
echo ""

# Vote for the proposal
echo "Voting yes for the proposal"
echo ""
vid-noded tx gov vote 1 yes --from node1 --yes
sleep 7
echo ""
echo "Vote given"
echo ""

echo "Waiting for the voting period to end...."
echo ""
sleep 65

# Check if the recipient has recived the grant
ACTUAL_GRANTS_CREDIT=$(vid-noded q bank balances $1 --output json | jq '.balances[0]["amount"]')
EXPECTED_GRANTS_CREDIT='"'68'"'

if [ $ACTUAL_GRANTS_CREDIT != $EXPECTED_GRANTS_CREDIT ]; then
    echo "Recipient did not recieve the expected grants amount: $EXPECTED_GRANTS_CREDIT, their wallet balance: $ACTUAL_GRANTS_CREDIT"
    exit 1
else
    echo "Recipient has recieved the grants: $ACTUAL_GRANTS_CREDIT"
fi
