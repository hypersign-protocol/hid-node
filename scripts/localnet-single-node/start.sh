#!/bin/bash

# Start all three nodes
tmux new -s node1 -d hid-noded start --home=$HOME/.hid-node/node1
