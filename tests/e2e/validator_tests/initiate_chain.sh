#!/bin/bash

# This script will run a non-validator full node

. common.sh

# Init chain config
echo "Setting up Validator 1 and Validator 2 Config"

VALIDATOR_1_NAME="val1"
VALIDATOR_1_WALLET="validator1"

VALIDATOR_2_NAME="val2"
VALIDATOR_2_WALLET="validator2"

VAL_1_HOME_DIR=$HOME/.hid-node/$VALIDATOR_1_NAME
VAL_2_HOME_DIR=$HOME/.hid-node/$VALIDATOR_2_NAME

init_node ${VALIDATOR_1_NAME} ${VALIDATOR_1_WALLET} 36657 36656 10090 10091

sleep 2

init_node ${VALIDATOR_2_NAME} ${VALIDATOR_2_WALLET} 46657 46656 11090 11091

sleep 2

echo "Validator config is set"

# Run the chain

echo "Running the chain (pre genesis stage)"

pre_genesis_validator ${VALIDATOR_1_NAME} ${VALIDATOR_1_WALLET}

run_chain ${VALIDATOR_1_NAME} ${VALIDATOR_1_WALLET} 36657 36656 10090 10091

echo "Validator 1 is now a validator"

echo "Starting the second full node validator"

sleep 5

# copy genesis
cp ${VAL_1_HOME_DIR}/config/genesis.json ${VAL_2_HOME_DIR}/config/genesis.json

# add peers
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$(hid-noded tendermint show-node-id --home=${VAL_1_HOME_DIR})@127.0.0.1:36656\"|g" ${VAL_2_HOME_DIR}/config/config.toml

run_chain ${VALIDATOR_2_NAME} ${VALIDATOR_2_WALLET} 46657 46656 11090 11091

echo "Second Full Node is now running"