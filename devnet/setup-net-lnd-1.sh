#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source $CONFIG_FILE
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

./lncli-create-with-seed.expect 1 $LND_1_WALLET_PASSWORD "$LND_1_WALLET_SEED"

sleep 2

# Send some funds to the LND-1 wallet
LND_1_BTC_ADDR=$(./lncli-1.sh newaddress p2tr | jq -r ".address")
./bitcoin-cli-1.sh sendtoaddress "$LND_1_BTC_ADDR" 1
./generate-blocks-1.sh 3
sleep 1
./lncli-1.sh walletbalance
