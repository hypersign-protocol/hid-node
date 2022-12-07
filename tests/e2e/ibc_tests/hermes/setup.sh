#!/bin/bash


# Copy config.toml to hermes config directory 
HERMES_HOME="$HOME/.hermes"
rm -rf ${HERMES_HOME}

mkdir ${HERMES_HOME}
cp ./hermes/config.toml ${HERMES_HOME}

echo "Add vidnode ibc relayer key"
hermes keys add --key-file ./hermes/test_keys/ibc_relayer_vidnode.json --chain vidnode

echo "Add osmosis ibc relayer key"
hermes keys add --key-file ./hermes/test_keys/ibc_relayer_osmosis.json --chain osmosischain

# Provide some tokens to relayer accounts ( $1 - vid-node relayer ; $2 - osmosis relayer )
vid-noded tx bank send $1 vid18t0uj2t9us7ufny0pdk94jvjt9mjtj9p72uzuq 1000000uvid --broadcast-mode block  --keyring-backend test --chain-id vidnode --yes
osmosisd tx bank send $2 osmo15w294mm9jm68ty5edw6l9wdr0nx8eswyg5fr66 1000000uosmo --broadcast-mode block --node tcp://localhost:36657 --yes

echo ""
echo "Create hermes channel"
echo y | hermes create channel --a-chain vidnode --b-chain osmosischain --a-port transfer --b-port transfer --new-client-connection
echo ""
echo "Channel Created"