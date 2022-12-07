#!/bin/bash

# Setup single node
echo "Setting up single node"
echo ""
../../../../scripts/localnet-single-node/setup.sh
echo ""
echo "Single Node Setup done"
echo ""

# Running the node
echo "Starting the node..."
echo ""
tmux new -s node1 -d vid-noded start
sleep 6
if [[ -n $(vid-noded status) ]]; then
  echo "vid-noded daemon is now running"
else
  echo "vid-noded daemon failed to start, exiting...."
  exit 1
fi
