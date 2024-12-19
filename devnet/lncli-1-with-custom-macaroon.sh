#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source "$CONFIG_FILE"
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

"$LNCLI_EXE_PATH" \
	--rpcserver $LND_1_RPCLISTEN \
	--lnddir $LND_1_LNDDIR \
	--network regtest \
	--macaroonpath  "$LND1_CUSTOM_MACAROON_FILE" \
	"$@"
