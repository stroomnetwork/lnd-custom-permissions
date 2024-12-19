package cli_lib

import (
	"github.com/stroomnetwork/lnd-custom-permissions/tools/lnd_tools"
	cli "github.com/urfave/cli/v3"
)

func NewLndGRPCConnectCliFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "lnd.cert-path",
			Usage: "Path `TLS_CERT_PATH` to the TLS certificate for the GRPC interface of LND node.",
			Value: DefaultLndCertPath,
		},
		&cli.StringFlag{
			Name:  "lnd.macaroon-path",
			Usage: "Path `MACAROON_PATH` to the macaroon file for the GRPC interface of LND node.",
			Value: DefaultLndMacaroonPath,
		},
		&cli.StringFlag{
			Name:  "lnd.rpc-address",
			Usage: "Address `LND_GRPC_API_ADDR` of the LND node's GRPC interface.",
			Value: DefaultLndRpcAddress,
		},
	}
}

func ExtractLndGRPCConnectConfig(cliCmd *cli.Command) *lnd_tools.LndGRPCConnectConfig {
	return &lnd_tools.LndGRPCConnectConfig{
		CertPath:     cliCmd.String("lnd.cert-path"),
		MacaroonPath: cliCmd.String("lnd.macaroon-path"),
		RpcAddress:   cliCmd.String("lnd.rpc-address"),
	}
}
