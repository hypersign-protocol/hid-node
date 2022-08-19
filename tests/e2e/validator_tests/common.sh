#!/bin/bash

# This script sets up a validator in the pre-genesis 

init_node() {
    # $1 validator_name | $2 validator_wallet_name | $3 rpc port | $4 p2p port | $5 grpc | $6 grpc web | $7 peers

    HOME_DIR="$HOME/.hid-node/$1"
    # Setting up config files
    rm -rf ${HOME_DIR}

    # Make directories for hid-node config
    mkdir ${HOME_DIR}

    # Init node
    hid-noded init $1 --chain-id=hidnode --home=${HOME_DIR}

    # Set the hid-node config
    hid-noded config chain-id hidnode --home=${HOME_DIR}
    hid-noded config keyring-backend test --home=${HOME_DIR}

    # Create key for the node
    hid-noded keys add $2 --keyring-backend=test --home=${HOME_DIR}

    # change staking denom to uhid
    cat ${HOME_DIR}/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="uhid"' > ${HOME_DIR}/config/tmp_genesis.json && mv ${HOME_DIR}/config/tmp_genesis.json ${HOME_DIR}/config/genesis.json

    # update crisis variable to uhid
    cat ${HOME_DIR}/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="uhid"' > ${HOME_DIR}/config/tmp_genesis.json && mv ${HOME_DIR}/config/tmp_genesis.json ${HOME_DIR}/config/genesis.json

    # update gov genesis
    cat ${HOME_DIR}/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uhid"' > ${HOME_DIR}/config/tmp_genesis.json && mv ${HOME_DIR}/config/tmp_genesis.json ${HOME_DIR}/config/genesis.json

    # update ssi genesis
    cat ${HOME_DIR}/config/genesis.json | jq '.app_state["ssi"]["did_method"]="hs"' > ${HOME_DIR}/config/tmp_genesis.json && mv ${HOME_DIR}/config/tmp_genesis.json ${HOME_DIR}/config/genesis.json
    cat ${HOME_DIR}/config/genesis.json | jq '.app_state["ssi"]["did_namespace"]="devnet"' > ${HOME_DIR}/config/tmp_genesis.json && mv ${HOME_DIR}/config/tmp_genesis.json ${HOME_DIR}/config/genesis.json

    # update slashing genesis
    cat ${HOME_DIR}/config/genesis.json | jq '.app_state["slashing"]["params"]["signed_blocks_window"]="10"' > ${HOME_DIR}/config/tmp_genesis.json && mv ${HOME_DIR}/config/tmp_genesis.json ${HOME_DIR}/config/genesis.json
    cat ${HOME_DIR}/config/genesis.json | jq '.app_state["slashing"]["params"]["downtime_jail_duration"]="300s"' > ${HOME_DIR}/config/tmp_genesis.json && mv ${HOME_DIR}/config/tmp_genesis.json ${HOME_DIR}/config/genesis.json
    cat ${HOME_DIR}/config/genesis.json | jq '.app_state["slashing"]["params"]["slash_fraction_downtime"]="0.500000000000000000"' > ${HOME_DIR}/config/tmp_genesis.json && mv ${HOME_DIR}/config/tmp_genesis.json ${HOME_DIR}/config/genesis.json

    # update mint genesis
    cat ${HOME_DIR}/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="uhid"' > ${HOME_DIR}/config/tmp_genesis.json && mv ${HOME_DIR}/config/tmp_genesis.json ${HOME_DIR}/config/genesis.json

    # change rpc and p2p ports
    sed -i -E 's|tcp://127.0.0.1:26658||g' ${HOME_DIR}/config/config.toml
    sed -i -E "s|tcp://127.0.0.1:26657|tcp://127.0.0.1:$3|g" ${HOME_DIR}/config/config.toml
    sed -i -E "s|tcp://0.0.0.0:26656|tcp://0.0.0.0:$4|g" ${HOME_DIR}/config/config.toml

    # change app.toml values
    sed -i -E "s|0.0.0.0:9090|0.0.0.0:$5|g" ${HOME_DIR}/config/app.toml
    sed -i -E "s|0.0.0.0:9091|0.0.0.0:$6|g" ${HOME_DIR}/config/app.toml

    # change config.toml values
    sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' ${HOME_DIR}/config/config.toml
    sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' ${HOME_DIR}/config/config.toml
    sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' ${HOME_DIR}/config/config.toml
}

pre_genesis_validator() {
    # create validator node with tokens
    HOME_DIR="$HOME/.hid-node/$1"
    hid-noded add-genesis-account $(hid-noded keys show $2 -a --keyring-backend=test --home=${HOME_DIR}) 5000000000000uhid --home=${HOME_DIR}
    hid-noded gentx $2 1000000000000uhid --keyring-backend=test --home=${HOME_DIR} --chain-id=hidnode --min-self-delegation="500000000000"
    hid-noded collect-gentxs --home=${HOME_DIR}
}

post_genesis_validator() {
    # $1 - moniker/val1 dir | $2 moniker/val2 dir | $3 validator1_wallet_name | $4 validator2_wallet_name | $5 val1_rpc | $6 val2_rpc | $7 val2_wallet_name
    HOME_DIR_VAL1="$HOME/.hid-node/$1"
    HOME_DIR_VAL2="$HOME/.hid-node/$2"

    # Send some money to the second validator
    TRANSFER_AMOUNT="2000000000000uhid"
    STAKE_AMOUNT="480000000000uhid"
    MIN_DELEGATION_AMOUNT="50000000000"

    echo
    echo "Sending money from validator 1 to second full node"
    echo
    hid-noded tx bank send $3 $4 ${TRANSFER_AMOUNT} --keyring-backend=test --home ${HOME_DIR_VAL1} --node $5 --chain-id hidnode --yes
    sleep 7

    echo
    echo "Promoting full node to validator node"
    echo
    hid-noded tx staking create-validator \
    --from $7 \
    --amount ${STAKE_AMOUNT} \
    --pubkey "$(hid-noded tendermint show-validator --home=${HOME_DIR_VAL2})" \
    --chain-id hidnode \
    --moniker=$2 \
    --commission-max-change-rate=0.01 \
    --commission-max-rate=1.0 \
    --commission-rate=0.1 \
    --min-self-delegation=${MIN_DELEGATION_AMOUNT} \
    --node=$6 \
    --chain-id=hidnode \
    --keyring-backend=test \
    --home=${HOME_DIR_VAL2} \
    --yes


}

run_chain() {
    HOME_DIR="$HOME/.hid-node/$1"
    tmux new -s $2 -d hid-noded start  --home=${HOME_DIR}
}