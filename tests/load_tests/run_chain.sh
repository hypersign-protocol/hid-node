#!/bin/bash

# Setup chain
echo "Setting up chain"
../../scripts/localnet-single-node/setup.sh
echo "Setup done"
echo ""

# Run the chain
echo "Running hid-node"
echo ""
tmux new -s hidnode -d hid-noded start
sleep 5
if [[ -n $(hid-noded status) ]]; then
  echo "hid-noded daemon is now running"
  echo ""
else
  echo "hid-noded daemon failed to start, exiting...."
  exit 1
fi