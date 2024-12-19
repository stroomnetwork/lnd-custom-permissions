#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source $CONFIG_FILE
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

BITCOIND_1_WALLET_ADDR_1=$(cat $BITCOIND_1_WALLET_ADDR_1_FILE)
./bitcoin-cli-1.sh generatetoaddress $1 $BITCOIND_1_WALLET_ADDR_1
