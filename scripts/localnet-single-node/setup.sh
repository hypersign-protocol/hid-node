#!/bin/bash

# Use the already built binary or use a different one
BINARY=$1
if [ -z ${BINARY} ]; then
	BINARY=vid-noded
fi

# Check if the binary is installed
${BINARY} &> /dev/null

RET_VAL=$?
if [ ${RET_VAL} -ne 0 ]; then
    echo "${BINARY} binary not found"
    exit 1
fi

# Setting up config files
rm -rf $HOME/.vid-node/

# Make directories for vid-node config
mkdir $HOME/.vid-node

# Init node
vid-noded init --chain-id=vidnode node1 --home=$HOME/.vid-node

# Change vid-node config
vid-noded configure min-gas-prices 0uvid

# Create key for the node
vid-noded keys add node1 --keyring-backend=test --home=$HOME/.vid-node

# change staking denom to uvid
cat $HOME/.vid-node/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="uvid"' > $HOME/.vid-node/config/tmp_genesis.json && mv $HOME/.vid-node/config/tmp_genesis.json $HOME/.vid-node/config/genesis.json

# update crisis variable to uvid
cat $HOME/.vid-node/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="uvid"' > $HOME/.vid-node/config/tmp_genesis.json && mv $HOME/.vid-node/config/tmp_genesis.json $HOME/.vid-node/config/genesis.json

# update gov genesis
cat $HOME/.vid-node/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uvid"' > $HOME/.vid-node/config/tmp_genesis.json && mv $HOME/.vid-node/config/tmp_genesis.json $HOME/.vid-node/config/genesis.json
cat $HOME/.vid-node/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="50s"' > $HOME/.vid-node/config/tmp_genesis.json && mv $HOME/.vid-node/config/tmp_genesis.json $HOME/.vid-node/config/genesis.json

# update ssi genesis
cat $HOME/.vid-node/config/genesis.json | jq '.app_state["ssi"]["chain_namespace"]="devnet"' > $HOME/.vid-node/config/tmp_genesis.json && mv $HOME/.vid-node/config/tmp_genesis.json $HOME/.vid-node/config/genesis.json

# update mint genesis
cat $HOME/.vid-node/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="uvid"' > $HOME/.vid-node/config/tmp_genesis.json && mv $HOME/.vid-node/config/tmp_genesis.json $HOME/.vid-node/config/genesis.json

# create validator node with tokens
vid-noded add-genesis-account $(vid-noded keys show node1 -a --keyring-backend=test --home=$HOME/.vid-node) 500000000000000000uvid --home=$HOME/.vid-node --keyring-backend test
vid-noded gentx node1 50000000000000000uvid --keyring-backend=test --home=$HOME/.vid-node --chain-id=vidnode
vid-noded collect-gentxs --home=$HOME/.vid-node

# change app.toml values
sed -i -E '104s/enable = false/enable = true/' $HOME/.vid-node/config/app.toml
sed -i -E '107s/swagger = false/swagger = true/' $HOME/.vid-node/config/app.toml


# change config.toml values
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.vid-node/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' $HOME/.vid-node/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' $HOME/.vid-node/config/config.toml
