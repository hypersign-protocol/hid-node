#!/bin/bash
#!/bin/bash

set -eu pipefail

echo "-------------------- Starting the tests for Community Pool Spend Governance Proposals --------------------"
echo ""
# Running the node
../init_node.sh

echo ""
echo "Creating a wallet for the recipient"
echo ""
hid-noded keys add recipient
RECIPIENT_ADDRESS=$(hid-noded keys show recipient -a)
GRANTS=$(cat ./community-pool-spend-proposal.json | jq ".amount")
cat ./community-pool-spend-proposal.json | jq ".recipient=\"${RECIPIENT_ADDRESS}\"" > ./tmp-community-pool-spend-proposal.json && mv ./tmp-community-pool-spend-proposal.json ./community-pool-spend-proposal.json

echo ""
echo "================== Submitting a proposal and voting YES on it =================="
echo ""

./proposal_vote_yes.sh ${RECIPIENT_ADDRESS} ${GRANTS}

echo ""
echo "========================== Test Completed Successfully ========================="
echo ""

echo ""
echo "================== Submitting a proposal and voting NO on it =================="
echo ""

./proposal_vote_no.sh ${RECIPIENT_ADDRESS}
sleep 1

# Stop the chain
kill $(lsof -t -i:26657)

echo ""
echo "========================== Test Completed Successfully ========================="
echo ""
echo "-------------- Tests for Community Pool Spend Governance Proposals Completed Successfully ----------------"