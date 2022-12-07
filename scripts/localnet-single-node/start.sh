#!/bin/bash

# Start all three nodes
tmux new -s node1 -d vid-noded start --home=$HOME/.vid-node/node1
