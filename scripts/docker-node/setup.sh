#!/bin/bash

echo "Setting up Chain."

# Init node
hid-noded init --chain-id=hidnode node1  &> /dev/null

# Create key for the node
hid-noded keys add node1 --keyring-backend=test  &> /dev/null

# change staking denom to uhid
cat /root/.hid-node/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="uhid"' > /root/.hid-node/config/tmp_genesis.json && mv /root/.hid-node/config/tmp_genesis.json /root/.hid-node/config/genesis.json

# create validator node with tokens
hid-noded add-genesis-account $(hid-noded keys show node1 -a --keyring-backend=test ) 100000000000uhid,100000000000stake 
hid-noded gentx node1 500000000uhid --keyring-backend=test  --chain-id=hidnode
hid-noded collect-gentxs 

# update crisis variable to uhid
cat /root/.hid-node/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="uhid"' > /root/.hid-node/config/tmp_genesis.json && mv /root/.hid-node/config/tmp_genesis.json /root/.hid-node/config/genesis.json

# udpate gov genesis
cat /root/.hid-node/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uhid"' > /root/.hid-node/config/tmp_genesis.json && mv /root/.hid-node/config/tmp_genesis.json /root/.hid-node/config/genesis.json

# update mint genesis
cat /root/.hid-node/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="uhid"' > /root/.hid-node/config/tmp_genesis.json && mv /root/.hid-node/config/tmp_genesis.json /root/.hid-node/config/genesis.json

# update ssi genesis
cat /root/.hid-node/config/genesis.json | jq '.app_state["ssi"]["did_method"]="hs"' > /root/.hid-node/config/tmp_genesis.json && mv /root/.hid-node/config/tmp_genesis.json /root/.hid-node/config/genesis.json
cat /root/.hid-node/config/genesis.json | jq '.app_state["ssi"]["did_namespace"]="devnet"' > /root/.hid-node/config/tmp_genesis.json && mv /root/.hid-node/config/tmp_genesis.json /root/.hid-node/config/genesis.json

# change app.toml values
sed -i -E '104s/enable = false/enable = true/' /root/.hid-node/config/app.toml
sed -i -E '107s/swagger = false/swagger = true/' /root/.hid-node/config/app.toml


# change config.toml values
sed -i -E 's|tcp://127.0.0.1:26658|tcp://0.0.0.0:26658|g' /root/.hid-node/config/config.toml
sed -i -E 's|tcp://127.0.0.1:26657|tcp://0.0.0.0:26657|g' /root/.hid-node/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' /root/.hid-node/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' /root/.hid-node/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' /root/.hid-node/config/config.toml

echo "Chain Setup is done."