#!/bin/bash

echo "====================== Tests for Validator Started ======================="

echo "======= Intialising and the chain with a validator and a full node ======="

./initiate_chain.sh

echo "=========================================================================="

echo "======= Promoting the full node to validator ======="

sleep 5

./promote_validator.sh

echo "===================================================="

echo "======= Jailing the newly promoted validator ======="

sleep 5

./jailed_validator_check.sh

echo "===================================================="

echo "Stopping any vid-node instance"

sleep 2

kill $(lsof -t -i:36657)

echo "=============== Test Completed ==============="
