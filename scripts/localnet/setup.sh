#!/bin/bash
rm -rf $HOME/.vid-node/


# Make directories for vid-node config
mkdir $HOME/.vid-node
mkdir $HOME/.vid-node/node1
mkdir $HOME/.vid-node/node2
mkdir $HOME/.vid-node/node3

# init all three nodes
vid-noded init --chain-id=vidnode node1 --home=$HOME/.vid-node/node1
vid-noded init --chain-id=vidnode node2 --home=$HOME/.vid-node/node2
vid-noded init --chain-id=vidnode node3 --home=$HOME/.vid-node/node3

# Change vid-node minimum gas prices
vid-noded configure min-gas-prices 0uvid --home=$HOME/.vid-node/node1
vid-noded configure min-gas-prices 0uvid --home=$HOME/.vid-node/node2
vid-noded configure min-gas-prices 0uvid --home=$HOME/.vid-node/node3

# create keys for all three nodes
vid-noded keys add node1 --keyring-backend=test --home=$HOME/.vid-node/node1
vid-noded keys add node2 --keyring-backend=test --home=$HOME/.vid-node/node2
vid-noded keys add node3 --keyring-backend=test --home=$HOME/.vid-node/node3

# change staking denom to uvid
cat $HOME/.vid-node/node1/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="uvid"' > $HOME/.vid-node/node1/config/tmp_genesis.json && mv $HOME/.vid-node/node1/config/tmp_genesis.json $HOME/.vid-node/node1/config/genesis.json

# create validator node with tokens to transfer to the three other nodes
vid-noded add-genesis-account $(vid-noded keys show node1 -a --keyring-backend=test --home=$HOME/.vid-node/node1) 5000000000000000000uvid --home=$HOME/.vid-node/node1
vid-noded gentx node1 5000000000000000000uvid --keyring-backend=test --home=$HOME/.vid-node/node1 --chain-id=vidnode
vid-noded collect-gentxs --home=$HOME/.vid-node/node1

# update crisis variable to uvid
cat $HOME/.vid-node/node1/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="uvid"' > $HOME/.vid-node/node1/config/tmp_genesis.json && mv $HOME/.vid-node/node1/config/tmp_genesis.json $HOME/.vid-node/node1/config/genesis.json

# udpate gov genesis
cat $HOME/.vid-node/node1/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uvid"' > $HOME/.vid-node/node1/config/tmp_genesis.json && mv $HOME/.vid-node/node1/config/tmp_genesis.json $HOME/.vid-node/node1/config/genesis.json

# update mint genesis
cat $HOME/.vid-node/node1/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="uvid"' > $HOME/.vid-node/node1/config/tmp_genesis.json && mv $HOME/.vid-node/node1/config/tmp_genesis.json $HOME/.vid-node/node1/config/genesis.json

#update ssi genesis
cat $HOME/.vid-node/node1/config/genesis.json | jq '.app_state["ssi"]["chain_namespace"]="devnet"' > $HOME/.vid-node/node1/config/tmp_genesis.json && mv $HOME/.vid-node/node1/config/tmp_genesis.json $HOME/.vid-node/node1/config/genesis.json

# change app.toml values

sed -i -E '104s/enable = false/enable = true/' $HOME/.vid-node/node1/config/app.toml
sed -i -E '107s/swagger = false/swagger = true/' $HOME/.vid-node/node1/config/app.toml

# node2
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:2317|g' $HOME/.vid-node/node2/config/app.toml
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9088|g' $HOME/.vid-node/node2/config/app.toml
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9089|g' $HOME/.vid-node/node2/config/app.toml
sed -i -E '104s/enable = false/enable = true/' $HOME/.vid-node/node2/config/app.toml
sed -i -E '107s/swagger = false/swagger = true/' $HOME/.vid-node/node2/config/app.toml

# node3
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:3317|g' $HOME/.vid-node/node3/config/app.toml
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9086|g' $HOME/.vid-node/node3/config/app.toml
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9087|g' $HOME/.vid-node/node3/config/app.toml
sed -i -E '104s/enable = false/enable = true/' $HOME/.vid-node/node3/config/app.toml
sed -i -E '107s/swagger = false/swagger = true/' $HOME/.vid-node/node3/config/app.toml

# change config.toml values

# node1
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.vid-node/node1/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' $HOME/.vid-node/node1/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' $HOME/.vid-node/node1/config/config.toml

# node2
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:36658|g' $HOME/.vid-node/node2/config/config.toml
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:36657|g' $HOME/.vid-node/node2/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:36656|g' $HOME/.vid-node/node2/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.vid-node/node2/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' $HOME/.vid-node/node2/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' $HOME/.vid-node/node2/config/config.toml

# node3
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:46658|g' $HOME/.vid-node/node3/config/config.toml
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:46657|g' $HOME/.vid-node/node3/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:46656|g' $HOME/.vid-node/node3/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.vid-node/node3/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' $HOME/.vid-node/node3/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' $HOME/.vid-node/node3/config/config.toml

# copy node1 genesis file to node2 and node3
cp $HOME/.vid-node/node1/config/genesis.json $HOME/.vid-node/node2/config/genesis.json
cp $HOME/.vid-node/node1/config/genesis.json $HOME/.vid-node/node3/config/genesis.json


# Copy tendermint node id of node1 to persistent peers of node2 and node3
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$(vid-noded tendermint show-node-id --home=$HOME/.vid-node/node1)@127.0.0.1:26656\"|g" $HOME/.vid-node/node2/config/config.toml
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$(vid-noded tendermint show-node-id --home=$HOME/.vid-node/node1)@127.0.0.1:26656\"|g" $HOME/.vid-node/node3/config/config.toml



