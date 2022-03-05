#!/bin/bash

# create second validator
sleep 7
hid-noded tx staking create-validator --amount=500000000uhid --from=node2 --pubkey=$(hid-noded tendermint show-validator --home=$HOME/.hid-node/node2) --moniker="node2" --chain-id="hidnode" --commission-rate="0.1" --commission-max-rate="0.2" --commission-max-change-rate="0.05" --min-self-delegation="5000" --keyring-backend=test --home=$HOME/.hid-node/node2 --yes
# create third validator
sleep 7
hid-noded tx staking create-validator --amount=400000000uhid --from=node3 --pubkey=$(hid-noded tendermint show-validator --home=$HOME/.hid-node/node3) --moniker="node3" --chain-id="hidnode" --commission-rate="0.1" --commission-max-rate="0.2" --commission-max-change-rate="0.05" --min-self-delegation="4000" --keyring-backend=test --home=$HOME/.hid-node/node3 --yes


