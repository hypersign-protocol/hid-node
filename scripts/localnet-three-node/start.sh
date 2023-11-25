#!/bin/bash

# Start all three nodes
tmux new -s node1 -d hid-noded start --home=$HOME/.hid-node/node1
tmux new -s node2 -d hid-noded start --home=$HOME/.hid-node/node2
tmux new -s node3 -d hid-noded start --home=$HOME/.hid-node/node3

