#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source $CONFIG_FILE
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

./lncli-create-with-seed.expect 2-signer $LND_2_SIGNER_WALLET_PASSWORD "$LND_2_SIGNER_WALLET_SEED"

sleep 3

# https://github.com/lightningnetwork/lnd/blob/master/docs/remote-signing.md
./lncli-2-signer.sh wallet accounts list > "$LND_2_SIGNER_ACCOUNTS_FILE"
./lncli-2-signer.sh bakemacaroon --save_to "$LND_2_SIGNER_CUSTOM_MACAROON_FILE" \
                message:write signer:generate address:read onchain:write

