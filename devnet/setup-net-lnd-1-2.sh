#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source $CONFIG_FILE
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

./setup-net-lnd-1.sh

./lncli-create-with-seed.expect 2 $LND_2_WALLET_PASSWORD "$LND_2_WALLET_SEED"

sleep 2

# Send some funds to the LND-1 wallet
LND_2_BTC_ADDR=$(./lncli-2.sh newaddress p2tr | jq -r ".address")
./bitcoin-cli-1.sh sendtoaddress "$LND_2_BTC_ADDR" 2
./generate-blocks-1.sh 3
sleep 1
./lncli-2.sh walletbalance
