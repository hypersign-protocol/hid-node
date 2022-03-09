#!/bin/bash

echo "Setting up Chain."

# Setting up config files
rm -rf /usr/local/app/node1


# Make directories for hid-node config
mkdir /usr/local/app/node1

# Init node
hid-noded init --chain-id=hidnode node1 --home=/usr/local/app/node1 &> /dev/null

# Create key for the node
hid-noded keys add node1 --keyring-backend=test --home=/usr/local/app/node1 &> /dev/null

# change staking denom to uhid
cat /usr/local/app/node1/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="uhid"' > /usr/local/app/node1/config/tmp_genesis.json && mv /usr/local/app/node1/config/tmp_genesis.json /usr/local/app/node1/config/genesis.json

# create validator node with tokens
hid-noded add-genesis-account $(hid-noded keys show node1 -a --keyring-backend=test --home=/usr/local/app/node1) 100000000000uhid,100000000000stake --home=/usr/local/app/node1
hid-noded gentx node1 500000000uhid --keyring-backend=test --home=/usr/local/app/node1 --chain-id=hidnode
hid-noded collect-gentxs --home=/usr/local/app/node1

# update crisis variable to uhid
cat /usr/local/app/node1/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="uhid"' > /usr/local/app/node1/config/tmp_genesis.json && mv /usr/local/app/node1/config/tmp_genesis.json /usr/local/app/node1/config/genesis.json

# udpate gov genesis
cat /usr/local/app/node1/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uhid"' > /usr/local/app/node1/config/tmp_genesis.json && mv /usr/local/app/node1/config/tmp_genesis.json /usr/local/app/node1/config/genesis.json

# update mint genesis
cat /usr/local/app/node1/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="uhid"' > /usr/local/app/node1/config/tmp_genesis.json && mv /usr/local/app/node1/config/tmp_genesis.json /usr/local/app/node1/config/genesis.json

# change app.toml values
sed -i -E '104s/enable = false/enable = true/' /usr/local/app/node1/config/app.toml
sed -i -E '107s/swagger = false/swagger = true/' /usr/local/app/node1/config/app.toml


# change config.toml values
sed -i -E 's|tcp://127.0.0.1:26658|tcp://0.0.0.0:26658|g' /usr/local/app/node1/config/config.toml
sed -i -E 's|tcp://127.0.0.1:26657|tcp://0.0.0.0:26657|g' /usr/local/app/node1/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' /usr/local/app/node1/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' /usr/local/app/node1/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' /usr/local/app/node1/config/config.toml

echo "Chain Setup is done."