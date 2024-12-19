#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source "$CONFIG_FILE"
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

go run \
    ../cmd/lnd-custom-permissions/main.go \
        "watch-lnd" \
        --lnd.cert-path "$LND_1_TLS_CERT_FILE" \
        --lnd.macaroon-path "$LND_1_RPC_WATCHER_MACAROON_FILE" \
        --lnd.rpc-address "$LND_1_RPCLISTEN" \
        --lnd.middleware-name "watcher-middleware" \
        --macaroon.custom-caveat-name "rpc-interceptor-caveat"
