#!/bin/bash

# Download the binary if it doesn't exist
hid-noded &> /dev/null

RET_VAL=$?
if [ ${RET_VAL} -ne 0 ]; then
    echo "hid-noded binary not found"
    exit 1
fi

# Setting up config files
rm -rf $HOME/.hid-node/


# Make directories for hid-node config
mkdir $HOME/.hid-node

# Init node
hid-noded init --chain-id=hidnode node1 --home=$HOME/.hid-node

# Create key for the node
hid-noded keys add node1 --keyring-backend=test --home=$HOME/.hid-node

# change staking denom to uhid
cat $HOME/.hid-node/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="uhid"' > $HOME/.hid-node/config/tmp_genesis.json && mv $HOME/.hid-node/config/tmp_genesis.json $HOME/.hid-node/config/genesis.json

# create validator node with tokens
hid-noded add-genesis-account $(hid-noded keys show node1 -a --keyring-backend=test --home=$HOME/.hid-node) 100000000000uhid,100000000000stake --home=$HOME/.hid-node
hid-noded gentx node1 500000000uhid --keyring-backend=test --home=$HOME/.hid-node --chain-id=hidnode
hid-noded collect-gentxs --home=$HOME/.hid-node

# update crisis variable to uhid
cat $HOME/.hid-node/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="uhid"' > $HOME/.hid-node/config/tmp_genesis.json && mv $HOME/.hid-node/config/tmp_genesis.json $HOME/.hid-node/config/genesis.json

# udpate gov genesis
cat $HOME/.hid-node/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uhid"' > $HOME/.hid-node/config/tmp_genesis.json && mv $HOME/.hid-node/config/tmp_genesis.json $HOME/.hid-node/config/genesis.json

# update mint genesis
cat $HOME/.hid-node/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="uhid"' > $HOME/.hid-node/config/tmp_genesis.json && mv $HOME/.hid-node/config/tmp_genesis.json $HOME/.hid-node/config/genesis.json

# change app.toml values
sed -i -E '104s/enable = false/enable = true/' $HOME/.hid-node/config/app.toml
sed -i -E '107s/swagger = false/swagger = true/' $HOME/.hid-node/config/app.toml


# change config.toml values
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.hid-node/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' $HOME/.hid-node/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' $HOME/.hid-node/config/config.toml

hid-noded config chain-id hidnode
hid-noded config keyring-backend test
