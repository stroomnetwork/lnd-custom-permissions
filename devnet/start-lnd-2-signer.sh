#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source "$CONFIG_FILE"
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

mkdir -p "$LND_2_SIGNER_LNDDIR"

"$LND_EXE_PATH" \
	--nolisten \
	--nobootstrap \
	--lnddir="${LND_2_SIGNER_LNDDIR}" \
	--alias="${LND_2_SIGNER_ALIAS}" \
	--color="${LND_2_SIGNER_COLOR}" \
	--bitcoin.active \
	--bitcoin.regtest \
	--debuglevel="${LND_2_SIGNER_DEBUGLEVEL}" \
	--rpclisten="${LND_2_SIGNER_RPCLISTEN}" \
	--restlisten="${LND_2_SIGNER_RESTLISTEN}" \
	--bitcoin.basefee="${LND_2_SIGNER_BITCOIN_BASEFEE}" \
	--bitcoin.feerate="${LND_2_SIGNER_BITCOIN_FEERATE}" \
	--bitcoin.node=nochainbackend \
    "$@" \
    &

echo $! > $LND_2_SIGNER_PID_FILE
wait