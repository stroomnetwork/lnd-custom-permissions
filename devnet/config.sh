# Read .env file and set variables.
# Only LND_SRC_DIR variable is currently required to compile lnd(debug version with all rpc enabled) and lncli.
if [ -e ".env" ]; then
    source .env
else
    echo "The file .env does not exist. Please create it. Put the LND_SRC_DIR variable in it."
    exit 1
fi

if [ -z "$LND_SRC_DIR" ]; then
    echo "The variable LND_SRC_DIR is empty. Please set it in the .env file."
    exit 1
fi

# ************************** BEGIN GENERAL VARIABLES ************************

# Directory where all devenv files are stored. Like data dirs of bitcoind, lnd, ...
TMP_DIR=./tmp

# Directory where all debug binaries are stored. Currently only lnd and lncli.
BIN_DIR=./bin


# Path to the lnd debug binary.
LND_EXE_PATH="${BIN_DIR}/lnd-debug"

# Path to the lncli debug binary.
LNCLI_EXE_PATH="${BIN_DIR}/lncli-debug"

# Number of blocks to mine during bitcoin setup.
# Some blocks should be mined to enable CSV, CLTV, taproot, ...
NUMBER_OF_INITIAL_BLOCKS=1000

# ************************** END GENERAL VARIABLES ************************

# ************************** BEGIN BITCOIND-1 VARIABLES *******************

# ZeroMQ address for publishing raw blocks.
# It is used by lnd to listen for new blocks.
BITCOIND_1_ZMQPUBRAWBLOCK_ADDR=tcp://127.0.0.1:28332

# ZeroMQ address for publishing raw transactions.
# It is used by lnd to listen for new transactions.
BITCOIND_1_ZMQPUBRAWTX_ADDR=tcp://127.0.0.1:28333

# RPC address of the bitcoind.
# Listen for JSON-RPC connections.
BITCOIND_1_RPCBIND_ADDR=127.0.0.1

# Listen for JSON-RPC connections on <port>
BITCOIND_1_RPCPORT=18443

# Password for JSON-RPC connections.
BITCOIND_1_RPCPASSWORD=some-pass

# Username for JSON-RPC connections.
BITCOIND_1_RPCUSER=some-user

# Directory where the bitcoind data is stored.
BITCOIND_1_DATADIR=${TMP_DIR}/bitcoind-1-data

# Name of the bitcoind wallet.
BITCOIND_1_WALLET_NAME=bitcoind-1-wallet

# File to store the address of the bitcoind wallet.
# It is used during mining to specify address for mining rewards.
BITCOIND_1_WALLET_ADDR_1_FILE=${TMP_DIR}/bitcoind-1-wallet-addr-1

# File to store the pid of the bitcoind process.
BITCOIND_1_PID_FILE=${TMP_DIR}/bitcoind-1.pid

# Value of the fallback fee in BTC/kvB.
# It is `-fallbackfee=<amt>` option of bitcoind.
# It is used because fee estimation fails on regtest 
# because there are no transactions initially.
# This values was taken from mainnet. 
# https://bitinfocharts.com/comparison/bitcoin-transactionfees.html#3y
# Note: this option expects the value in BTC/kVB 
# and it does not accept in scientific notation("0.000142", not "1.42e-4").
BITCOIND_1_FALLBACKFEE="0.000142"

# ************************** END BITCOIND-1 VARIABLES *******************


# ************************** BEGIN LND-1 VARIABLES **********************

# Directory where the lnd data is stored.
LND_1_LNDDIR=${TMP_DIR}/lnd-1

# Address of the lnd  gRPC RPC.
LND_1_RPCLISTEN=127.0.0.1:10009

# Address of the lnd REST server.
LND_1_RESTLISTEN=127.0.0.1:10007

# Peer address of the lnd.
LND_1_LISTEN=127.0.0.1:10008

# Alias of the lnd node.
LND_1_ALIAS=LND-1

# Color of the lnd node.
LND_1_COLOR="#FF0000"

# Debug level of the lnd node.
LND_1_DEBUGLEVEL=debug

# The base fee in millisatoshi we will charge for forwarding
# payments on our channels
LND_1_BITCOIN_BASEFEE=1000

# The fee rate used when forwarding payments on our channels. 
# The total fee charged is basefee + (amount * feerate / 1000000), 
# where amount is the forwarded amount
LND_1_BITCOIN_FEERATE=1

# File to store the pid of the lnd process.
LND_1_PID_FILE=${TMP_DIR}/lnd-1.pid

# Address for dlv headless server to listen on.
# It is used to debug lnd with delve.
LND_1_DLV_LISTEN=":10005"

# Path to the lnd TLS certificate file.
# Calculated value. Changing it will not change anything.
LND_1_TLS_CERT_FILE=${LND_1_LNDDIR}/tls.cert

# Path to the lnd TLS key file.
# Calculated value. Changing it will not change anything.
LND_1_TLS_KEY_FILE=${LND_1_LNDDIR}/tls.key

# Path to the lnd admin macaroon file.
# Calculated value. Changing it will not change anything.
LND_1_ADMIN_MACAROON_FILE=${LND_1_LNDDIR}/data/chain/bitcoin/regtest/admin.macaroon

# Path to the lnd macaroon for rpc interceptor.
# This macaroon has custom caveat thus allowing to intercept rpc calls(with this macaroon).
# Calculated value. Changing it will not change anything.
LND_1_RPC_INTERCEPTOR_MACAROON_FILE=${LND_1_LNDDIR}/data/chain/bitcoin/regtest/rpcinterceptor.macaroon

# Password for the lnd wallet.
LND_1_WALLET_PASSWORD="LND-1-12345678"

# Seed for the lnd wallet.
# Constant seed allows reproducible address generations and thus
# much easier debugging.
LND_1_WALLET_SEED="abandon glove clean wild frog frozen robust despair hire helmet slide coach boost evolve mix capable region climb vote dress buddy output puppy pulse"

# Path to the file with custom macaroon for lnd.
# This value should be sync with values in ./vscode/launch.json
LND1_CUSTOM_MACAROON_FILE=${TMP_DIR}/lnd-1-custom.macaroon

# Path to the file with watcher macaroon for lnd.
# This value should be sync with values in ./vscode/launch.json
LND_1_RPC_WATCHER_MACAROON_FILE=${TMP_DIR}/lnd-1-watcher.macaroon

# ************************** END LND-1 VARIABLES **********************

# ************************** BEGIN LND-2 VARIABLES **********************

# Directory where the lnd data is stored.
LND_2_LNDDIR=${TMP_DIR}/lnd-2

# Address of the lnd  gRPC RPC.
LND_2_RPCLISTEN=127.0.0.1:11009

# Address of the lnd REST server.
LND_2_RESTLISTEN=127.0.0.1:11007

# Peer address of the lnd.
LND_2_LISTEN=127.0.0.1:11008

# Alias of the lnd node.
LND_2_ALIAS=LND-2

# Color of the lnd node.
LND_2_COLOR="#00FF00"

# Debug level of the lnd node.
LND_2_DEBUGLEVEL=debug

# The base fee in millisatoshi we will charge for forwarding
# payments on our channels
LND_2_BITCOIN_BASEFEE=1000

# The fee rate used when forwarding payments on our channels. 
# The total fee charged is basefee + (amount * feerate / 1000000), 
# where amount is the forwarded amount
LND_2_BITCOIN_FEERATE=1

# File to store the pid of the lnd process.
LND_2_PID_FILE=${TMP_DIR}/lnd-2.pid

# Address for dlv headless server to listen on.
# It is used to debug lnd with delve.
LND_2_DLV_LISTEN=":11005"

# Path to the lnd TLS certificate file.
# Calculated value. Changing it will not change anything.
LND_2_TLS_CERT_FILE=${LND_2_LNDDIR}/tls.cert

# Path to the lnd TLS key file.
# Calculated value. Changing it will not change anything.
LND_2_TLS_KEY_FILE=${LND_2_LNDDIR}/tls.key

# Path to the lnd admin macaroon file.
# Calculated value. Changing it will not change anything.
LND_2_ADMIN_MACAROON_FILE=${LND_2_LNDDIR}/data/chain/bitcoin/regtest/admin.macaroon

# Password for the lnd wallet.
LND_2_WALLET_PASSWORD="LND-2-12345678"

# Seed for the lnd wallet.
# Constant seed allows reproducible address generations and thus
# much easier debugging.
LND_2_WALLET_SEED="abandon insect crush camera steak pepper laugh float tornado south cram salmon volcano soldier sail wrap wood mango forward ghost stove hawk cotton olympic"

# ************************** END LND-2 VARIABLES **********************

# ************************** BEGIN LND-2-SIGNER VARIABLES **********************

# Directory where the lnd data is stored.
LND_2_SIGNER_LNDDIR=${TMP_DIR}/lnd-2-signer

# Address of the lnd  gRPC RPC.
LND_2_SIGNER_RPCLISTEN=127.0.0.1:11109

# Address of the lnd REST server.
LND_2_SIGNER_RESTLISTEN=127.0.0.1:11107

# Alias of the lnd node.
LND_2_SIGNER_ALIAS=LND-2-signer

# Color of the lnd node.
LND_2_SIGNER_COLOR="#00FF00"

# Debug level of the lnd node.
LND_2_SIGNER_DEBUGLEVEL=debug

# The base fee in millisatoshi we will charge for forwarding
# payments on our channels
LND_2_SIGNER_BITCOIN_BASEFEE=1000

# The fee rate used when forwarding payments on our channels. 
# The total fee charged is basefee + (amount * feerate / 1000000), 
# where amount is the forwarded amount
LND_2_SIGNER_BITCOIN_FEERATE=1

# File to store the pid of the lnd process.
LND_2_SIGNER_PID_FILE=${TMP_DIR}/lnd-2-signer.pid

# Address for dlv headless server to listen on.
# It is used to debug lnd with delve.
LND_2_SIGNER_DLV_LISTEN=":11105"

# Path to the lnd TLS certificate file.
# Calculated value. Changing it will not change anything.
LND_2_SIGNER_TLS_CERT_FILE=${LND_2_SIGNER_LNDDIR}/tls.cert

# Path to the lnd TLS key file.
# Calculated value. Changing it will not change anything.
LND_2_SIGNER_TLS_KEY_FILE=${LND_2_SIGNER_LNDDIR}/tls.key

# Path to the lnd admin macaroon file.
# Calculated value. Changing it will not change anything.
LND_2_SIGNER_ADMIN_MACAROON_FILE=${LND_SIGNER_2_LNDDIR}/data/chain/bitcoin/regtest/admin.macaroon

# Password for the lnd wallet.
LND_2_SIGNER_WALLET_PASSWORD="LND-2-signer-12345678"

# Seed for the lnd wallet.
# Constant seed allows reproducible address generations and thus
# much easier debugging.
LND_2_SIGNER_WALLET_SEED="abandon insect crush camera steak pepper laugh float tornado south cram salmon volcano soldier sail wrap wood mango forward ghost stove hawk cotton olympic"

# File with `xpub`s of the wallet.
LND_2_SIGNER_ACCOUNTS_FILE=${TMP_DIR}/lnd-2-signer-accounts.json

# A custom macaroon for the watch-only node, 
# with the minimum required permissions on the signer instance:
LND_2_SIGNER_CUSTOM_MACAROON_FILE=${TMP_DIR}/lnd-2-signer-custom.macaroon

# ************************** END LND-2-SIGNER VARIABLES **********************

# ************************** BEGIN LND-2-WATCHER VARIABLES **********************

# Directory where the lnd data is stored.
LND_2_WATCHER_LNDDIR=${TMP_DIR}/lnd-2-watcher

# Address of the lnd  gRPC RPC.
LND_2_WATCHER_RPCLISTEN=127.0.0.1:11209

# Address of the lnd REST server.
LND_2_WATCHER_RESTLISTEN=127.0.0.1:11207

# Peer address of the lnd.
LND_2_WATCHER_LISTEN=127.0.0.1:11208

# Alias of the lnd node.
LND_2_WATCHER_ALIAS=LND-2

# Color of the lnd node.
LND_2_WATCHER_COLOR="#00FF00"

# Debug level of the lnd node.
LND_2_WATCHER_DEBUGLEVEL=debug

# The base fee in millisatoshi we will charge for forwarding
# payments on our channels
LND_2_WATCHER_BITCOIN_BASEFEE=1000

# The fee rate used when forwarding payments on our channels. 
# The total fee charged is basefee + (amount * feerate / 1000000), 
# where amount is the forwarded amount
LND_2_WATCHER_BITCOIN_FEERATE=1

# File to store the pid of the lnd process.
LND_2_WATCHER_PID_FILE=${TMP_DIR}/lnd-2-watcher.pid

# Address for dlv headless server to listen on.
# It is used to debug lnd with delve.
LND_2_WATCHER_DLV_LISTEN=":11205"

# Path to the lnd TLS certificate file.
# Calculated value. Changing it will not change anything.
LND_2_WATCHER_TLS_CERT_FILE=${LND_2_WATCHER_LNDDIR}/tls.cert

# Path to the lnd TLS key file.
# Calculated value. Changing it will not change anything.
LND_2_WATCHER_TLS_KEY_FILE=${LND_2_WATCHER_LNDDIR}/tls.key

# Path to the lnd admin macaroon file.
# Calculated value. Changing it will not change anything.
LND_2_WATCHER_ADMIN_MACAROON_FILE=${LND_WATCHER_2_LNDDIR}/data/chain/bitcoin/regtest/admin.macaroon

# Password for the lnd wallet.
LND_2_WATCHER_WALLET_PASSWORD="LND-2-watcher-12345678"

# ************************** END LND-2-WATCHER VARIABLES **********************