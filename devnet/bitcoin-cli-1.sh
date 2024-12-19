#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source $CONFIG_FILE
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

bitcoin-cli \
        -rpcuser=$BITCOIND_1_RPCUSER \
        -rpcpassword=$BITCOIND_1_RPCPASSWORD \
        -chain=regtest \
        -rpcport=$BITCOIND_1_RPCPORT \
        "$@"