#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source $CONFIG_FILE
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

mkdir -p $TMP_DIR
./bitcoin-cli-1.sh createwallet $BITCOIND_1_WALLET_NAME

BITCOIND_1_WALLET_ADDR_1=$(./bitcoin-cli-1.sh getnewaddress)
echo $BITCOIND_1_WALLET_ADDR_1 > $BITCOIND_1_WALLET_ADDR_1_FILE

./bitcoin-cli-1.sh generatetoaddress $NUMBER_OF_INITIAL_BLOCKS $BITCOIND_1_WALLET_ADDR_1

./bitcoin-cli-1.sh getbalance