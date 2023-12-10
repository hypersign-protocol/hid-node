#!/bin/bash

set -e pipefail

# Set the binary name
BINARY=hid-noded

# Check if the binary is installed
${BINARY} &> /dev/null

CHAINID=hidnode
CHAIN_NAMESAPCE=devnet

RET_VAL=$?
if [ ${RET_VAL} -ne 0 ]; then
    echo "hid-noded binary is not installed in your system."
    exit 1
fi

# Setting up config files
rm -rf $HOME/.hid-node/

# Make directories for hid-node config
mkdir $HOME/.hid-node

# Init node
hid-noded init --chain-id=$CHAINID node1 --home=$HOME/.hid-node

# Create key for the node or recover existing key from mnemonic
if [[ -z "$MNEMONIC" ]]; then
    hid-noded keys add node1 --keyring-backend=test --home=$HOME/.hid-node
else
    echo $MNEMONIC | hid-noded keys add node1 --keyring-backend=test --home=$HOME/.hid-node --recover
fi

# change staking denom to uhid
cat $HOME/.hid-node/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="uhid"' > $HOME/.hid-node/config/tmp_genesis.json && mv $HOME/.hid-node/config/tmp_genesis.json $HOME/.hid-node/config/genesis.json

# update crisis variable to uhid
cat $HOME/.hid-node/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="uhid"' > $HOME/.hid-node/config/tmp_genesis.json && mv $HOME/.hid-node/config/tmp_genesis.json $HOME/.hid-node/config/genesis.json

# update gov genesis
cat $HOME/.hid-node/config/genesis.json | jq '.app_state["gov"]["params"]["min_deposit"][0]["denom"]="uhid"' > $HOME/.hid-node/config/tmp_genesis.json && mv $HOME/.hid-node/config/tmp_genesis.json $HOME/.hid-node/config/genesis.json
cat $HOME/.hid-node/config/genesis.json | jq '.app_state["gov"]["params"]["voting_period"]="50s"' > $HOME/.hid-node/config/tmp_genesis.json && mv $HOME/.hid-node/config/tmp_genesis.json $HOME/.hid-node/config/genesis.json

# update ssi genesis
cat $HOME/.hid-node/config/genesis.json | jq '.app_state["ssi"]["chainNamespace"]="'$CHAIN_NAMESAPCE'"' > $HOME/.hid-node/config/tmp_genesis.json && mv $HOME/.hid-node/config/tmp_genesis.json $HOME/.hid-node/config/genesis.json

# update mint genesis
cat $HOME/.hid-node/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="uhid"' > $HOME/.hid-node/config/tmp_genesis.json && mv $HOME/.hid-node/config/tmp_genesis.json $HOME/.hid-node/config/genesis.json

# create validator node with tokens
hid-noded add-genesis-account $(hid-noded keys show node1 -a --keyring-backend=test --home=$HOME/.hid-node) 500000000000000000uhid --home=$HOME/.hid-node --keyring-backend test
hid-noded gentx node1 50000000000000000uhid --keyring-backend=test --home=$HOME/.hid-node --chain-id=$CHAINID
hid-noded collect-gentxs --home=$HOME/.hid-node

# change app.toml values
sed -i -E '119s/enable = false/enable = true/' $HOME/.hid-node/config/app.toml
sed -i -E '122s/swagger = false/swagger = true/' $HOME/.hid-node/config/app.toml
sed -i -E 's|tcp://localhost:1317|tcp://0.0.0.0:1317|g' $HOME/.hid-node/config/app.toml
sed -i -E 's|localhost:9090|0.0.0.0:9090|g' $HOME/.hid-node/config/app.toml
sed -i -E '140s/enabled-unsafe-cors = false/enabled-unsafe-cors = true/' $HOME/.hid-node/config/app.toml

# change config.toml values
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.hid-node/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' $HOME/.hid-node/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' $HOME/.hid-node/config/config.toml

echo -e "\nConfiguration set up is done, you are ready to run hid-noded now!"

echo -e "\nPlease note the important chain configurations below:"

echo -e "\nRPC server address: http://localhost:26657"
echo -e "API server address: http://localhost:1317"
echo -e "DID Namespace: $CHAIN_NAMESAPCE"