#!/bin/bash

set -eu pipefail

echo "------------- Running test for Node upgrade --------------"
echo ""
# Node Parameters
OLD_HIDNODE_BINARY=./binaries/hid-noded-old
NEW_HIDNODE_BINARY=./binaries/hid-noded-new
BLOCK_PERIOD=5

# Governance Parameters
PROPOSAL_NAME="test"
PROPOSAL_TITLE="Test Upgrade"
PROPOSAL_DESCRIPTION="E2E Test for Node Upgrade"
UPGRADE_HEIGHT=25
DEPOSIT="10000000uhid"


node_run_check() {
    if [[ -n $($1 status) ]]; then
        echo "hid-noded daemon is now running"
        echo ""
    else
        echo "hid-noded daemon failed to start, exiting...."
        kill $(lsof -t -i:26657)
        exit 1
    fi
}
# Running the node
echo "Setting up cosmovisor with old binary"
echo ""
../../../../scripts/localnet-single-node/setup.sh ${OLD_HIDNODE_BINARY}

# Create Cosmovisor directory
export DAEMON_NAME=hid-noded
export DAEMON_HOME=$HOME/.hid-node
mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin
cp ${OLD_HIDNODE_BINARY} $DAEMON_HOME/cosmovisor/genesis/bin/hid-noded

# Run Cosmovisor
echo "Cosmovisor Setup done"
echo ""
echo "Starting the node..."
echo ""
tmux new -s cosmovisor-node -d cosmovisor run start
sleep 6
node_run_check ${OLD_HIDNODE_BINARY}
if [[ -n $(${OLD_HIDNODE_BINARY} status) ]]; then
    echo "hid-noded daemon is now running"
    echo ""
else
    echo "hid-noded daemon failed to start, exiting...."
    kill $(lsof -t -i:26657)
    exit 1
fi

echo "Submitting Governance Proposal for Upgrade"
echo ""
${OLD_HIDNODE_BINARY} tx gov submit-proposal software-upgrade "${PROPOSAL_NAME}" --title "${PROPOSAL_TITLE}" --description "${PROPOSAL_DESCRIPTION}" --from node1 --upgrade-height "${UPGRADE_HEIGHT}" --deposit "${DEPOSIT}" --chain-id hidnode --broadcast-mode block --yes
echo ""
echo "Proposal Submitted"
echo ""

echo "Voting Yes on the upgrade proposal"
echo ""
${OLD_HIDNODE_BINARY} tx gov vote 1 yes --from node1 --chain-id hidnode --broadcast-mode block --yes
echo ""
echo "Vote Given"
echo ""

echo "Moving the latest binary to Cosmovisor directory"
echo ""
mkdir -p $HOME/.hid-node/cosmovisor/upgrades/test/bin
cp ${NEW_HIDNODE_BINARY} $HOME/.hid-node/cosmovisor/upgrades/test/bin/hid-noded

LATEST_BLOCK_HEIGHT=$(hid-noded status | jq '.SyncInfo["latest_block_height"]' | bc)
UPGRADE_CHECK_WAIT=$(($((${UPGRADE_HEIGHT} - ${LATEST_BLOCK_HEIGHT} + 5)) * ${BLOCK_PERIOD})) 

echo "Waiting for ${UPGRADE_CHECK_WAIT} seconds to check if the node is upgraded and running.."
echo ""
sleep ${UPGRADE_CHECK_WAIT}

# Check if the node is running or halted
if [[ -n $(${NEW_HIDNODE_BINARY} status) ]]; then
    echo "hid-noded daemon is now running"
    echo ""
else
    echo "hid-noded daemon failed to start, exiting...."
    kill $(lsof -t -i:26657)
    exit 1
fi
echo "Checking now if the node daemon has been upgraded"
echo ""
LATEST_BLOCK_HEIGHT=$(hid-noded status | jq '.SyncInfo["latest_block_height"]' | bc)
if [ ${LATEST_BLOCK_HEIGHT} -gt ${UPGRADE_HEIGHT} ]; then
    echo "Node has been upgraded"
    echo "------------- Test Completed Successfully --------------"
    echo ""
else
    echo "Node hasn't crossed the upgrade height: ${UPGRADE_HEIGHT}. Current Height: ${LATEST_BLOCK_HEIGHT}"
    kill -9 $(lsof -t -i:26657)
    exit 1
fi



sleep 4
# Stop the node
kill -9 $(lsof -t -i:26657)