#!/bin/bash

# Run vid-Node Chain
echo "Setting up vid-node chain"
echo ""
../../../scripts/localnet-single-node/setup.sh
echo ""
echo "Setup Done"
echo ""

echo "Running vid-node"
echo ""
tmux new -s vidnode -d vid-noded start
sleep 5
vid-noded status &> /dev/null
RET_VAL=$?
if [ ${RET_VAL} -eq 0 ]; then
  echo "vid-node daemon is now running"
  echo ""
else
  echo "vid-node daemon failed to start, exiting...."
  exit 1
fi

# Run Osmosis Chain
echo "Setting up Osmosis Chain"
echo ""
./osmosis/osmosis_setup.sh
echo ""
echo "Setup Done"
echo ""

echo "Running osmosis"
echo ""
tmux new -s osmosisnode -d osmosisd start
sleep 5
osmosisd status &> /dev/null
RET_VAL=$?
if [ ${RET_VAL} -eq 0 ]; then
  echo "osmosisd daemon is now running"
  echo ""
else
  echo "osmosisd daemon failed to start, exiting...."
  exit 1
fi

# Run Hermes Relayer
echo "Setting up hermes relayer"
vid_NODE_VALIDATOR_WALLET=$(vid-noded keys show node1 -a --keyring-backend test)
OSMOSIS_VALIDATOR_WALLET=$(osmosisd keys show osmonode1 -a)
./hermes/setup.sh ${vid_NODE_VALIDATOR_WALLET} ${OSMOSIS_VALIDATOR_WALLET}
echo ""
sleep 3
echo "Starting hermes relayer"
echo ""
tmux new -s hermesrelayer -d hermes start
sleep 2
echo "Hermes has been started"
echo ""

echo "Transferring tokens from vid Node to Osmosis"
echo ""
IBC_TRANSFER_RESULT=$(vid-noded tx ibc-transfer transfer transfer channel-0 ${OSMOSIS_VALIDATOR_WALLET} 1234uvid --broadcast-mode block --from ${vid_NODE_VALIDATOR_WALLET} --output json --keyring-backend test --chain-id vidnode --yes)

CODE=$(echo ${IBC_TRANSFER_RESULT} | jq '.code')
TXHASH=$(echo ${IBC_TRANSFER_RESULT} | jq '.txhash')
if [ ${CODE} -eq 0 ]; then
  echo "vid Token is transferred successfully through IBC. Tx Hash: ${TXHASH}"
  echo ""
else
  echo "vid Token did not went through IBC. Tx Hash: ${TXHASH}"
  exit 1
fi

echo "Stopping vid-node chain"
echo ""
kill -9 $(lsof -t -i:26657)
echo "Stopping osmosis chain"
echo ""
kill -9 $(lsof -t -i:36657)
echo "Stopping Hermes"
tmux kill-server
