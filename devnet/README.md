# Introduction
These scripts should only be launched from this dir. They assume specific position of other parts, thus they work when launched from this dir only.

## Initial configuration

To work properly lnd should be compiled in debug mode and with all RPCs enabled.

Steps for building lnd:

1. Get lnd source. https://github.com/lightningnetwork/lnd/ Clone it somewhere. 
2. Create `.env` file and put path to lnd source in variable `LND_SRC_DIR` in that file. File `.env.example` can be uses as an example.
3. Run
```sh
$ ./build-lnd-debug.sh
```  


## Explanation of each file/directory
| File/directory | Description |
| File Name                   | Description                       |
| --------------------------- | --------------------------------- |
| `.env`                      | File with the configuration variables. It is not included in git so it should be created. `.env.example` can be used as an example |
| `.env.example` |  Example of `.env` file |
| `README.md`                 | This file. Contains descriptions of different scripts and usage scernarios. |
| `attach-dlv-lnd-1.sh`       | Attaches delve `dlv` Go debugger to LND-1 process(`lnd-1`) in a headless mode. Allows connecting VSCode for debugging LND-1. |
| `attach-dlv-lnd-2.sh`       | Attaches delve `dlv` Go debugger to LND-2 process(`lnd-2`) in a headless mode. Allows connecting VSCode for debugging LND-2. |
| `attach-dlv-lnd-2-signer.sh`| Attaches delve `dlv` Go debugger to LND-2 signer process(`lnd-2-signer`) in a headless mode. Allows connecting VSCode for debugging LND-2 signer. |
| `attach-dlv-lnd-2-watcher.sh`| Attaches delve `dlv` Go debugger to LND-2 watcher process(`lnd-2-watcher`) in a headless mode. Allows connecting VSCode for debugging LND-2 watcher. |
| `bitcoin-cli-1.sh`          | bitcoin-cli for `bitcoind-1`. Command line client for `bitcoind`. |
| `bin`                       | Directory with debug binaries of `lnd` and `lncli`. |
| `build-lnd-debug.sh`        | Builds `lnd` and `lncli` for debug and copy executables to bin directory. |
| `clean.sh`                  | Cleans devenv files. Currently deletes `tmp` directory. |
| `config.sh`                 | File with configuration parameters. |
| `generate-blocks-1.sh`      | Generates blocks using bitcoind. Usage `./generate-blocks-1.sh <number-of-blocks>`. Blocks are mined to `bitcoind-1` internal wallet. |
| `lncli-1.sh`                | lncli for LND-1(`lnd-1`). |
| `lncli-2.sh`                | lncli for LND-2(`lnd-2`). |
| `lncli-2-signer.sh`         | lncli for LND-2 signer(`lnd-2-signer`). |
| `lncli-2-watcher.sh`        | lncli for LND-2 watcher(`lnd-2-watcher`). |
| `lncli-create-watchonly.expect` | Script for creating watch-only wallet in LND. Usage `./lncli-create-watchonly.expect <lnd-name> <password> <JSON file with exported keys>`. |
| `lncli-create-with-seed.expect` | Script for creating a wallet from seed in LND. Usage `./lncli-create-with-seed.expect <lnd-name> <password> <seed>`. |
| `lncli-unlock.expect`       | Script for unlocking LND. Usage `./lncli-unlock.expect <lnd-name> <password>`. |
| `setup-bitcoind-1.sh`                  | Initial setup for `bitcoind-1`. Create an address for `bitcoind-1` wallet. Mine some blocks to this address. |
| `setup-lnd-2-signer.sh`     | Setup default configuration for LND-2 signer. It will create default wallet from the seed and export pubkeys. |
| `setup-lnd-2-watcher.sh`    | Setup LND-2 watcher. It will create watch-only wallet and import pubkeys. |
| `start-bitcoind-1.sh`       | Starts `bitcoind-1`. |
| `start-lnd-1.sh`            | Starts `lnd-1`. RPC middleware is enabled. https://docs.lightning.engineering/lightning-network-tools/lnd/rpc-middleware-interceptor |
| `start-lnd-2.sh`            | Starts `lnd-2`. |
| `start-lnd-2-signer.sh`     | Starts `lnd-2-signer`. Signer should be launched before watcher because watcher requires pubkeys from the signer and tries to connect to it. |
| `start-lnd-2-watcher.sh`    | Starts `lnd-2-watcher`. |
| `tmp`                       | Directory with devenv files/directories. Like bitcoind, lnd data directories. |
| `unlock-lnd.sh`             | Script for unlocking LND. Usage `./unlock-lnd.sh <lnd-name> <password>`. |


## Usage scenario
### Debugging LND-1

1. Clean from previous devenv.
```shell
$ ./clean.sh
```

2. Start bitcoind.
```shell
$ ./start-bitcoind-1.sh
```

3. Setup bitcoin.
```shell
$ ./setup-bitcoind-1.sh
```

4. Start lnd-1.
```shell
$ ./start-lnd-1.sh
```

5. Setup LND-1 (create wallet, send some money to it).
```shell
$ ./setup-net-lnd-1.sh
```

6. Attach debugger to lnd-1.
```shell
$ ./attach-dlv-lnd-1.sh
```

7. In VSCode attach to headless debugger.

### Launching Remote Signer from Scratch

1. Clean from previous devenv.
```shell
(terminal-1)$ ./clean.sh
```

2. Start bitcoind.
```shell
(terminal-1)$ ./start-bitcoind-1.sh
```

3. Setup bitcoin.
```shell
(terminal-2)$ ./setup-bitcoind-1.sh
```

4. Start lnd-2-signer. Signer should be launched before watcher before watcher because watcher requires pubkeys from the signer and tries to connect to it.
```shell
(terminal-2)$ ./start-lnd-2-signer.sh
```

5. Setup default configuration for signer. It will create default wallet and export keys.
```shell
(terminal-3)$ ./setup-lnd-2-signer.sh
```

6. Start lnd-2-watcher
```shell
(terminal-3)$ ./start-lnd-2-watcher.sh
```

7. Setup watcher. It will create watch-only wallet and import pubkeys.
```shell
(terminal-4)$ ./setup-lnd-2-watcher.sh
```

8. Check that everything works. Generate a new address.
```shell
(terminal-4)$ ./lncli-2-watcher.sh newaddress p2tr
```

9. To debug in VSCode lnd-2-watcher:
```shell
(terminal-4)$ ./attach-dlv-lnd-2-watcher.sh
```
and connect VSCode to remote debugger using the  corresponding profile (or simply connect to port 11205)

10. To debug in VSCode lnd-2-signer:
```shell
(terminal-4)$ ./attach-dlv-lnd-2-signer.sh
```
and connect VSCode to remote debugger using the  corresponding profile (or simply connect to port 11105)



## Used ports
| Port | Usage | Variable |
| ---- | ----- | -------- |
| 10005 | Delve debugger listening port. This debugger is attached to LND-1 | `LND_1_DLV_LISTEN` |
| 10007 | LND-1 REST RPC | `LND_2_RESTLISTEN` |
| 10008 | LND-1 peer listening | `LND_1_LISTEN` |
| 10009 | LND-1 gRPC RPC | `LND_1_RPCLISTEN` |
| 11005 | Delve debugger listening port. This debugger is attached to LND-2 | `LND_2_DLV_LISTEN` |
| 11007 | LND-2 REST RPC | `LND_2_RESTLISTEN` |
| 11008 | LND-2 peer listening | `LND_2_LISTEN` |
| 11009 | LND-2 gRPC RPC | `LND_2_RPCLISTEN` |
| 11105 | Delve debugger listening port. This debugger is attached to LND-2-SIGNER | `LND_2_SIGNER_DLV_LISTEN` |
| 11107 | LND-2-SIGNER REST RPC | `LND_2_SIGNER_RESTLISTEN` |
| 11109 | LND-2-SIGNER gRPC RPC | `LND_2_SIGNER_RPCLISTEN` |
| 11205 | Delve debugger listening port. This debugger is attached to LND-2-WATCHER | `LND_2_WATCHER_DLV_LISTEN` |
| 11207 | LND-2-WATCHER REST RPC | `LND_2_WATCHER_RESTLISTEN` |
| 11208 | LND-2-WATCHER peer listening | `LND_2_WATCHER_LISTEN` |
| 11209 | LND-2-WATCHER gRPC RPC | `LND_2_WATCHER_RPCLISTEN` |
| 18443 | bitcoind-1 JSON-RPC port | `BITCOIND_1_RPCPORT` |
| 28332 | bitcoind-1 ZeroMQ address for publishing raw blocks | `BITCOIND_1_ZMQPUBRAWBLOCK_ADDR` |
| 28333 | bitcoind-1 ZeroMQ address for publishing raw transactions. | `BITCOIND_1_ZMQPUBRAWTX_ADDR` |
