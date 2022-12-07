#!/bin/bash

# Setup chain
echo "Setting up chain"
../../scripts/localnet-single-node/setup.sh
echo "Setup done"
echo ""

# Run the chain
echo "Running vid-node"
echo ""
tmux new -s vidnode -d vid-noded start
sleep 5
if [[ -n $(vid-noded status) ]]; then
  echo "vid-noded daemon is now running"
  echo ""
else
  echo "vid-noded daemon failed to start, exiting...."
  exit 1
fi