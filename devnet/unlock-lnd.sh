#!/usr/bin/env bash

CONFIG_FILE="./config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source "$CONFIG_FILE"
else
    echo "Config file $CONFIG_FILE does not exist. Exiting."
    exit 1
fi

# Name of the LND to unlock. E.g 0, 1-signer, ...
LND_NAME=$1

# Convert the LND name to uppercase and replace - with _
LND_NAME_FOR_VAR=$(echo $LND_NAME | tr '[:lower:]' '[:upper:]' | sed 's/-/_/g')

# Name of variable that holds the password for the LND
PASSWORD_VAR_NAME="LND_${LND_NAME_FOR_VAR}_WALLET_PASSWORD"

# Password for the LND
PASSWORD="${!PASSWORD_VAR_NAME}"

./lncli-unlock.expect $LND_NAME "$PASSWORD"