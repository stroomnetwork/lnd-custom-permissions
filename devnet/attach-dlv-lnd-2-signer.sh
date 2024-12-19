#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source "$CONFIG_FILE"
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

dlv \
	attach $(cat $LND_2_SIGNER_PID_FILE) \
	--headless \
	--listen=$LND_2_SIGNER_DLV_LISTEN \
	--accept-multiclient \
	--continue
