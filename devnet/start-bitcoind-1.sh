#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source $CONFIG_FILE
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

mkdir -p $BITCOIND_1_DATADIR

bitcoind \
 -server \
 -txindex \
 -chain=regtest \
 -zmqpubrawblock=$BITCOIND_1_ZMQPUBRAWBLOCK_ADDR \
 -zmqpubrawtx=$BITCOIND_1_ZMQPUBRAWTX_ADDR \
 -rpcallowip=0.0.0.0/0 \
 -rpcbind=$BITCOIND_1_RPCBIND_ADDR \
 -rpcport=18443 \
 -rpcpassword=$BITCOIND_1_RPCPASSWORD \
 -rpcuser=$BITCOIND_1_RPCUSER \
 -datadir=$BITCOIND_1_DATADIR \
 -fallbackfee="$BITCOIND_1_FALLBACKFEE" \
 "$@" \
 &

echo $! > $BITCOIND_1_PID_FILE
wait