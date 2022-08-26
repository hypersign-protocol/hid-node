#!/bin/bash

# Setting up config files
rm -rf $HOME/.osmosisd

# Make directories for osmosis config
mkdir $HOME/.osmosisd

# Init node
osmosisd init --chain-id=osmosischain node1 --home=$HOME/.osmosisd

# Create key for the node
osmosisd keys add osmonode1 --keyring-backend=test --home=$HOME/.osmosisd

# change staking denom to uosmo
cat $HOME/.osmosisd/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="uosmo"' > $HOME/.osmosisd/config/tmp_genesis.json && mv $HOME/.osmosisd/config/tmp_genesis.json $HOME/.osmosisd/config/genesis.json

# update crisis variable to uosmo
cat $HOME/.osmosisd/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="uosmo"' > $HOME/.osmosisd/config/tmp_genesis.json && mv $HOME/.osmosisd/config/tmp_genesis.json $HOME/.osmosisd/config/genesis.json

# update gov genesis
cat $HOME/.osmosisd/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uosmo"' > $HOME/.osmosisd/config/tmp_genesis.json && mv $HOME/.osmosisd/config/tmp_genesis.json $HOME/.osmosisd/config/genesis.json
cat $HOME/.osmosisd/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="50s"' > $HOME/.osmosisd/config/tmp_genesis.json && mv $HOME/.osmosisd/config/tmp_genesis.json $HOME/.osmosisd/config/genesis.json

# update mint genesis
cat $HOME/.osmosisd/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="uosmo"' > $HOME/.osmosisd/config/tmp_genesis.json && mv $HOME/.osmosisd/config/tmp_genesis.json $HOME/.osmosisd/config/genesis.json

# update txfees genesis
cat $HOME/.osmosisd/config/genesis.json | jq '.app_state["txfees"]["basedenom"]="uosmo"' > $HOME/.osmosisd/config/tmp_genesis.json && mv $HOME/.osmosisd/config/tmp_genesis.json $HOME/.osmosisd/config/genesis.json

# update epochs genesis
cat $HOME/.osmosisd/config/genesis.json | jq '.app_state["epochs"]["epochs"][1]["duration"]="60s"' > $HOME/.osmosisd/config/tmp_genesis.json && mv $HOME/.osmosisd/config/tmp_genesis.json $HOME/.osmosisd/config/genesis.json

# create validator node with tokens
osmosisd add-genesis-account $(osmosisd keys show osmonode1 -a --keyring-backend=test --home=$HOME/.osmosisd) 100000000000uosmo --home=$HOME/.osmosisd
osmosisd gentx osmonode1 500000000uosmo --keyring-backend=test --home=$HOME/.osmosisd --chain-id=osmosischain
osmosisd collect-gentxs --home=$HOME/.osmosisd

# change app.toml values
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:2317|g' $HOME/.osmosisd/config/app.toml
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9086|g' $HOME/.osmosisd/config/app.toml
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9087|g' $HOME/.osmosisd/config/app.toml
sed -i -E '104s/enable = false/enable = true/' $HOME/.osmosisd/config/app.toml
sed -i -E '107s/swagger = false/swagger = true/' $HOME/.osmosisd/config/app.toml

# change config.toml values
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:36658|g' $HOME/.osmosisd/config/config.toml
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:36657|g' $HOME/.osmosisd/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:36656|g' $HOME/.osmosisd/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.osmosisd/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' $HOME/.osmosisd/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' $HOME/.osmosisd/config/config.toml

osmosisd config chain-id osmosischain
osmosisd config keyring-backend test
