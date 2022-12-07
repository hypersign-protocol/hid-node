#!/bin/bash

# create second validator
sleep 7
vid-noded tx staking create-validator --amount=500000000uvid --from=node2 --pubkey=$(vid-noded tendermint show-validator --home=$HOME/.vid-node/node2) --moniker="node2" --chain-id="vidnode" --commission-rate="0.1" --commission-max-rate="0.2" --commission-max-change-rate="0.05" --min-self-delegation="5000" --keyring-backend=test --home=$HOME/.vid-node/node2 --yes
# create third validator
sleep 7
vid-noded tx staking create-validator --amount=400000000uvid --from=node3 --pubkey=$(vid-noded tendermint show-validator --home=$HOME/.vid-node/node3) --moniker="node3" --chain-id="vidnode" --commission-rate="0.1" --commission-max-rate="0.2" --commission-max-change-rate="0.05" --min-self-delegation="4000" --keyring-backend=test --home=$HOME/.vid-node/node3 --yes


