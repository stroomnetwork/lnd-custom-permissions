#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source "$CONFIG_FILE"
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

# https://docs.lightning.engineering/lightning-network-tools/lnd/macaroons
./lncli-1.sh \
	constrainmacaroon \
	--custom_caveat_name rpc-interceptor-caveat \
	--custom_caveat_condition read \
	"$LND_1_ADMIN_MACAROON_FILE" \
	"$LND_1_RPC_INTERCEPTOR_MACAROON_FILE"
