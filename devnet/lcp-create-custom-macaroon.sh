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
        "create-custom-macaroon" \
        --lnd.cert-path "$LND_1_TLS_CERT_FILE" \
        --lnd.macaroon-path "$LND_1_ADMIN_MACAROON_FILE" \
        --lnd.rpc-address "$LND_1_RPCLISTEN" \
        --macaroon.custom-caveat-name "rpc-interceptor-caveat" \
        --rules.add-invoice.max-value-sat  200000 \
        --rules.open-channel.max-push-amount-sat 0 \
        --rules.send-payment-v2.allowed-destination "self" \
        --save-to "$LND1_CUSTOM_MACAROON_FILE"
