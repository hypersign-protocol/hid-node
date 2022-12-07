#!/bin/bash

# Start all three nodes
tmux new -s node1 -d vid-noded start --home=$HOME/.vid-node/node1
tmux new -s node2 -d vid-noded start --home=$HOME/.vid-node/node2
tmux new -s node3 -d vid-noded start --home=$HOME/.vid-node/node3

