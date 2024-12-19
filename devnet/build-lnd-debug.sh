#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source $CONFIG_FILE
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

mkdir -p "$BIN_DIR"
# We need absolute paths because we do cd later.
LND_EXE_ABSOLUTE_PATH=$(readlink -f $LND_EXE_PATH)
LNCLI_EXE_ABSOLUTE_PATH=$(readlink -f $LNCLI_EXE_PATH)

cd $LND_SRC_DIR
# By default a lot of RPC is disabled in dev build of lnd.
# To enable RPC we need to add `with-rpc=1`.
make with-rpc=1 build

cp lnd-debug $LND_EXE_ABSOLUTE_PATH
cp lncli-debug $LNCLI_EXE_ABSOLUTE_PATH
