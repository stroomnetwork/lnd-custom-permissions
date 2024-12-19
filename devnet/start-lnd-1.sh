#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source "$CONFIG_FILE"
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

mkdir -p "$LND_1_LNDDIR"

"$LND_EXE_PATH" \
    --lnddir="${LND_1_LNDDIR}" \
	--externalip=127.0.0.1 \
	--alias="${LND_1_ALIAS}" \
	--color="${LND_1_COLOR}" \
	--bitcoin.active \
	--bitcoin.regtest \
	--debuglevel="${LND_1_DEBUGLEVEL}" \
	--listen="${LND_1_LISTEN}" \
	--rpclisten="${LND_1_RPCLISTEN}" \
	--restlisten="${LND_1_RESTLISTEN}" \
	--bitcoin.basefee="${LND_1_BITCOIN_BASEFEE}" \
	--bitcoin.feerate="${LND_1_BITCOIN_FEERATE}" \
	--bitcoin.node=bitcoind \
	--bitcoind.rpchost="127.0.0.1:${BITCOIND_1_RPCPORT}" \
	--bitcoind.rpcuser="${BITCOIND_1_RPCUSER}" \
	--bitcoind.rpcpass="${BITCOIND_1_RPCPASSWORD}" \
	--bitcoind.zmqpubrawblock="${BITCOIND_1_ZMQPUBRAWBLOCK_ADDR}" \
	--bitcoind.zmqpubrawtx="${BITCOIND_1_ZMQPUBRAWTX_ADDR}" \
	--rpcmiddleware.enable \
    "$@" \
    &

echo $! > $LND_1_PID_FILE
wait