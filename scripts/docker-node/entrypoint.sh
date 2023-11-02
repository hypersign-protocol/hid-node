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
# rm -rf /root/.hid-node/

# Make directories for hid-node config
mkdir /root/.hid-node

# Init node
hid-noded init --chain-id=hidnode node1 --home=/root/.hid-node

# Change hid-node config
hid-noded configure min-gas-prices 0uhid

# Create key for the node
if [ -n "$MNEMONIC" ]; then
  echo "$MNEMONIC" | hid-noded keys add node1 --keyring-backend=test --recover --home=/root/.hid-node
else
  hid-noded keys add node1 --keyring-backend=test --home=/root/.hid-node
fi

# change staking denom to uhid
cat /root/.hid-node/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="uhid"' > /root/.hid-node/config/tmp_genesis.json && mv /root/.hid-node/config/tmp_genesis.json /root/.hid-node/config/genesis.json

# update crisis variable to uhid
cat /root/.hid-node/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="uhid"' > /root/.hid-node/config/tmp_genesis.json && mv /root/.hid-node/config/tmp_genesis.json /root/.hid-node/config/genesis.json

# update gov genesis
cat /root/.hid-node/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uhid"' > /root/.hid-node/config/tmp_genesis.json && mv /root/.hid-node/config/tmp_genesis.json /root/.hid-node/config/genesis.json
cat /root/.hid-node/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="500s"' > /root/.hid-node/config/tmp_genesis.json && mv /root/.hid-node/config/tmp_genesis.json /root/.hid-node/config/genesis.json

# update ssi genesis
cat /root/.hid-node/config/genesis.json | jq '.app_state["ssi"]["chain_namespace"]="testnet"' > /root/.hid-node/config/tmp_genesis.json && mv /root/.hid-node/config/tmp_genesis.json /root/.hid-node/config/genesis.json

# update mint genesis
cat /root/.hid-node/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="uhid"' > /root/.hid-node/config/tmp_genesis.json && mv /root/.hid-node/config/tmp_genesis.json /root/.hid-node/config/genesis.json

# create validator node with tokens
hid-noded add-genesis-account $(hid-noded keys show node1 -a --keyring-backend=test --home=/root/.hid-node) 110000000000uhid --home=/root/.hid-node
hid-noded gentx node1 100000000000uhid --keyring-backend=test --home=/root/.hid-node --chain-id=hidnode
hid-noded collect-gentxs --home=/root/.hid-node

# change app.toml values
sed -i -E '112s/enable = false/enable = true/' /root/.hid-node/config/app.toml
sed -i -E '115s/swagger = false/swagger = true/' /root/.hid-node/config/app.toml
sed -i -E '133s/enabled-unsafe-cors = false/enabled-unsafe-cors = true/' /root/.hid-node/config/app.toml

# change config.toml values
sed -i -E 's|tcp://127.0.0.1:26657|tcp://0.0.0.0:26657|g' /root/.hid-node/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' /root/.hid-node/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' /root/.hid-node/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' /root/.hid-node/config/config.toml

# Run hid-node
hid-noded start
