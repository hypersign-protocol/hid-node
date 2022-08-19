#!/bin/bash

SECOND_VALIDATOR_RPC_PORT="46657"

kill $(lsof -t -i:${SECOND_VALIDATOR_RPC_PORT})

echo "Second Validator is shut down"

echo "Checking for validator's jailed status after 70 seconds"

VALIDATOR_DELEGATED_SHARES="480000000000.000000000000000000"

sleep 70

IS_JAILED=$(hid-noded q staking validators --home $HOME/.hid-node/val1 --node tcp://localhost:36657 --output json | jq '.validators' | jq -c '.[] | select(.delegator_shares=="'"${VALIDATOR_DELEGATED_SHARES}"'")' | jq '.jailed')
if [ $IS_JAILED == true ]; then
    echo "Success: Validator is Jailed!"
else
    echo "Fail: Validator is not jailed within the stipulated time frame"
fi