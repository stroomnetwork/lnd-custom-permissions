#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source "$CONFIG_FILE"
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

mkdir -p "$LND_2_WATCHER_LNDDIR"

"$LND_EXE_PATH" \
	--lnddir="${LND_2_WATCHER_LNDDIR}" \
	--externalip=127.0.0.1 \
	--alias="${LND_2_WATCHER_ALIAS}" \
	--color="${LND_2_WATCHER_COLOR}" \
	--bitcoin.active \
	--bitcoin.regtest \
	--debuglevel="${LND_2_WATCHER_DEBUGLEVEL}" \
	--listen="${LND_2_WATCHER_LISTEN}" \
	--restlisten="${LND_2_WATCHER_RESTLISTEN}" \
	--rpclisten="${LND_2_WATCHER_RPCLISTEN}" \
	--bitcoin.basefee="${LND_2_WATCHER_BITCOIN_BASEFEE}" \
	--bitcoin.feerate="${LND_2_WATCHER_BITCOIN_FEERATE}" \
	--bitcoin.node=bitcoind \
	--bitcoind.rpchost="127.0.0.1:${BITCOIND_1_RPCPORT}" \
	--bitcoind.rpcuser="${BITCOIND_1_RPCUSER}" \
	--bitcoind.rpcpass="${BITCOIND_1_RPCPASSWORD}" \
	--bitcoind.zmqpubrawblock="${BITCOIND_1_ZMQPUBRAWBLOCK_ADDR}" \
	--bitcoind.zmqpubrawtx="${BITCOIND_1_ZMQPUBRAWTX_ADDR}" \
	--remotesigner.enable \
	--remotesigner.rpchost="$LND_2_SIGNER_RPCLISTEN" \
	--remotesigner.tlscertpath="$LND_2_SIGNER_TLS_CERT_FILE" \
	--remotesigner.macaroonpath="$LND_2_SIGNER_CUSTOM_MACAROON_FILE" \
	"$@" \
	&

echo $! > $LND_2_WATCHER_PID_FILE
wait