#!/bin/bash

# send uvid from first node to second node
sleep 7
vid-noded tx bank send node1 $(vid-noded keys show node2 -a --keyring-backend=test --home=$HOME/.vid-node/node2) 500000000uvid --keyring-backend=test --home=$HOME/.vid-node/node1 --chain-id=vidnode --yes
sleep 7
vid-noded tx bank send node1 $(vid-noded keys show node3 -a --keyring-backend=test --home=$HOME/.vid-node/node3) 400000000uvid --keyring-backend=test --home=$HOME/.vid-node/node1 --chain-id=vidnode --yes
