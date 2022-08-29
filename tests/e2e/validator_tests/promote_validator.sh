#!/bin/bash

# This script will run a non-validator full node
. common.sh

VALIDATOR_1_NAME="val1"
VALIDATOR_1_WALLET="validator1"

VALIDATOR_2_NAME="val2"
VALIDATOR_2_WALLET="validator2"

VAL_1_HOME_DIR=$HOME/.hid-node/$VALIDATOR_1_NAME
VAL_2_HOME_DIR=$HOME/.hid-node/$VALIDATOR_2_NAME

VALIDATOR_1_RPC="tcp://127.0.0.1:36657"
VALIDATOR_2_RPC="tcp://127.0.0.1:46657"

VALIDATOR_1_WALLET_ADDRESS=$(hid-noded keys show ${VALIDATOR_1_WALLET} -a --keyring-backend=test --home=${VAL_1_HOME_DIR})
VALIDATOR_2_WALLET_ADDRESS=$(hid-noded keys show ${VALIDATOR_2_WALLET} -a --keyring-backend=test --home=${VAL_2_HOME_DIR})

post_genesis_validator ${VALIDATOR_1_NAME} ${VALIDATOR_2_NAME} ${VALIDATOR_1_WALLET_ADDRESS} ${VALIDATOR_2_WALLET_ADDRESS} ${VALIDATOR_1_RPC} ${VALIDATOR_2_RPC} ${VALIDATOR_2_WALLET}
