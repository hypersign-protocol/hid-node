#!/bin/bash

# send uhid from first node to second node
sleep 7
hid-noded tx bank send node1 $(hid-noded keys show node2 -a --keyring-backend=test --home=$HOME/.hid-node/node2) 500000000uhid --keyring-backend=test --home=$HOME/.hid-node/node1 --chain-id=hidnode --yes
sleep 7
hid-noded tx bank send node1 $(hid-noded keys show node3 -a --keyring-backend=test --home=$HOME/.hid-node/node3) 400000000uhid --keyring-backend=test --home=$HOME/.hid-node/node1 --chain-id=hidnode --yes
