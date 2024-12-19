package cli_lib

import (
	"context"

	"github.com/stroomnetwork/lnd-custom-permissions/commands"
	cli "github.com/urfave/cli/v3"
)

func NewWatchLndCli(out *interface{}) *cli.Command {
	flags := make([]cli.Flag, 0)
	flags = append(
		flags,
		NewLndGRPCConnectCliFlags()...,
	)
	flags = append(
		flags,
		&cli.StringFlag{
			Name:  "lnd.middleware-name",
			Usage: "Name of middleware `MIDDLEWARE_NAME` to be registered in the LND node.",
			Value: DefaultMiddlewareName,
		},
		&cli.StringFlag{
			Name:  "macaroon.custom-caveat-name",
			Usage: "Name `CAVEAT_NAME` of the custom caveat to be added to the macaroon.",
			Value: DefaultCustomCaveatName,
		},
	)
	cmd := &cli.Command{
		Name:        "watch-lnd",
		Description: "Watch LND node for incoming RPC calls. Approve/Reject requests with custom caveats.",
		Flags:       flags,
		Action: func(ctx context.Context, cliCmd *cli.Command) error {
			lndCfg := ExtractLndGRPCConnectConfig(cliCmd)
			*out = &commands.WatchLndCommand{
				Lnd:              lndCfg,
				CustomCaveatName: cliCmd.String("macaroon.custom-caveat-name"),
				MiddlewareName:   cliCmd.String("lnd.middleware-name"),
			}
			return nil
		},
	}

	return cmd
}
