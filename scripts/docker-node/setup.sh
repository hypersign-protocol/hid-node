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
rm -rf /root/.vid-node/

# Make directories for vid-node config
mkdir /root/.vid-node

# Init node
vid-noded init --chain-id=vidnode node1 --home=/root/.vid-node

# Change vid-node config
vid-noded configure min-gas-prices 0uvid

# Create key for the node
vid-noded keys add node1 --keyring-backend=test --home=/root/.vid-node

# change staking denom to uvid
cat /root/.vid-node/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="uvid"' > /root/.vid-node/config/tmp_genesis.json && mv /root/.vid-node/config/tmp_genesis.json /root/.vid-node/config/genesis.json

# update crisis variable to uvid
cat /root/.vid-node/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="uvid"' > /root/.vid-node/config/tmp_genesis.json && mv /root/.vid-node/config/tmp_genesis.json /root/.vid-node/config/genesis.json

# update gov genesis
cat /root/.vid-node/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uvid"' > /root/.vid-node/config/tmp_genesis.json && mv /root/.vid-node/config/tmp_genesis.json /root/.vid-node/config/genesis.json
cat /root/.vid-node/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="500s"' > /root/.vid-node/config/tmp_genesis.json && mv /root/.vid-node/config/tmp_genesis.json /root/.vid-node/config/genesis.json

# update ssi genesis
cat /root/.vid-node/config/genesis.json | jq '.app_state["ssi"]["chain_namespace"]="devnet"' > /root/.vid-node/config/tmp_genesis.json && mv /root/.vid-node/config/tmp_genesis.json /root/.vid-node/config/genesis.json

# update mint genesis
cat /root/.vid-node/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="uvid"' > /root/.vid-node/config/tmp_genesis.json && mv /root/.vid-node/config/tmp_genesis.json /root/.vid-node/config/genesis.json

# create validator node with tokens
vid-noded add-genesis-account $(vid-noded keys show node1 -a --keyring-backend=test --home=/root/.vid-node) 110000000000uvid --home=/root/.vid-node
vid-noded gentx node1 100000000000uvid --keyring-backend=test --home=/root/.vid-node --chain-id=vidnode
vid-noded collect-gentxs --home=/root/.vid-node

# change app.toml values
sed -i -E '104s/enable = false/enable = true/' /root/.vid-node/config/app.toml
sed -i -E '107s/swagger = false/swagger = true/' /root/.vid-node/config/app.toml


# change config.toml values
sed -i -E 's|tcp://127.0.0.1:26657|tcp://0.0.0.0:26657|g' /root/.vid-node/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' /root/.vid-node/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' /root/.vid-node/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' /root/.vid-node/config/config.toml
