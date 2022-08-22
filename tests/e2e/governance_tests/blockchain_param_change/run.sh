#!/bin/bash

set -eu pipefail

echo "-------------------- Starting the tests for Param Change Governance Proposals --------------------"
echo ""
# Running the node
./init_node.sh

echo ""
echo "================== Submitting a proposal and voting YES on it =================="
echo ""

./proposal_vote_yes.sh
sleep 4

echo ""
echo "========================== Test Completed Successfully ========================="
echo ""

echo ""
echo "================== Submitting a proposal and voting NO on it =================="
echo ""

./proposal_vote_no.sh
sleep 1

# Stop the chain
kill $(lsof -t -i:26657)

echo ""
echo "========================== Test Completed Successfully ========================="
echo ""
echo "-------------- Tests for Param Change Governance Proposals Completed Successfully ----------------"