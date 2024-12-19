#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source $CONFIG_FILE
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

./lncli-create-watchonly.expect 2-watcher $LND_2_WATCHER_WALLET_PASSWORD "$LND_2_SIGNER_ACCOUNTS_FILE"

