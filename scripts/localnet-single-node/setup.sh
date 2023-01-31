#!/bin/bash

# Use the already built binary or use a different one
BINARY=$1
if [ -z ${BINARY} ]; then
	BINARY=hid-noded
fi

# Check if the binary is installed
${BINARY} &> /dev/null

RET_VAL=$?
if [ ${RET_VAL} -ne 0 ]; then
    echo "${BINARY} binary not found"
    exit 1
fi

# Setting up config files
rm -rf $HOME/.hid-node/

# Make directories for hid-node config
mkdir $HOME/.hid-node

# Init node
hid-noded init --chain-id=hidnode node1 --home=$HOME/.hid-node

# Change hid-node config
hid-noded configure min-gas-prices 0uhid

# Create key for the node
hid-noded keys add node1 --keyring-backend=test --home=$HOME/.hid-node

# change staking denom to uhid
cat $HOME/.hid-node/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="uhid"' > $HOME/.hid-node/config/tmp_genesis.json && mv $HOME/.hid-node/config/tmp_genesis.json $HOME/.hid-node/config/genesis.json

# update crisis variable to uhid
cat $HOME/.hid-node/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="uhid"' > $HOME/.hid-node/config/tmp_genesis.json && mv $HOME/.hid-node/config/tmp_genesis.json $HOME/.hid-node/config/genesis.json

# update gov genesis
cat $HOME/.hid-node/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uhid"' > $HOME/.hid-node/config/tmp_genesis.json && mv $HOME/.hid-node/config/tmp_genesis.json $HOME/.hid-node/config/genesis.json
cat $HOME/.hid-node/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="50s"' > $HOME/.hid-node/config/tmp_genesis.json && mv $HOME/.hid-node/config/tmp_genesis.json $HOME/.hid-node/config/genesis.json

# update ssi genesis
cat $HOME/.hid-node/config/genesis.json | jq '.app_state["ssi"]["chain_namespace"]="devnet"' > $HOME/.hid-node/config/tmp_genesis.json && mv $HOME/.hid-node/config/tmp_genesis.json $HOME/.hid-node/config/genesis.json

# update mint genesis
cat $HOME/.hid-node/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="uhid"' > $HOME/.hid-node/config/tmp_genesis.json && mv $HOME/.hid-node/config/tmp_genesis.json $HOME/.hid-node/config/genesis.json

# create validator node with tokens
hid-noded add-genesis-account $(hid-noded keys show node1 -a --keyring-backend=test --home=$HOME/.hid-node) 500000000000000000uhid --home=$HOME/.hid-node --keyring-backend test
hid-noded gentx node1 50000000000000000uhid --keyring-backend=test --home=$HOME/.hid-node --chain-id=hidnode
hid-noded collect-gentxs --home=$HOME/.hid-node

# change app.toml values
sed -i -E '112s/enable = false/enable = true/' $HOME/.hid-node/config/app.toml
sed -i -E '115s/swagger = false/swagger = true/' $HOME/.hid-node/config/app.toml
sed -i -E '133s/enabled-unsafe-cors = false/enabled-unsafe-cors = true/' $HOME/.hid-node/config/app.toml


# change config.toml values
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.hid-node/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' $HOME/.hid-node/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' $HOME/.hid-node/config/config.toml
