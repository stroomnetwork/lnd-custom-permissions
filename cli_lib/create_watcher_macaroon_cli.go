package cli_lib

import (
	"context"

	"github.com/stroomnetwork/lnd-custom-permissions/commands"
	cli "github.com/urfave/cli/v3"
)

func NewCreateWatcherMacaroonCli(out *interface{}) *cli.Command {
	flags := make([]cli.Flag, 0)
	flags = append(
		flags,
		NewLndGRPCConnectCliFlags()...,
	)
	flags = append(
		flags,
		&cli.StringFlag{
			Name:  "save-to",
			Usage: "Path `OUTPUT_MACAROON_PATH` to the output macaroon file. ",
			Value: "custom.macaroon",
		},

		&cli.DurationFlag{
			Name:  "macaroon.duration",
			Usage: "Duration `DURATION` for which the macaroon will be valid.",
			Value: DefaultMacaroonDuration,
		},
	)
	cmd := &cli.Command{
		Name:        "create-watcher-macaroon",
		Description: "Create a custom macaroon with very limited permissions for watcher.",
		Flags:       flags,
		Action: func(ctx context.Context, cliCmd *cli.Command) error {
			lndCfg := ExtractLndGRPCConnectConfig(cliCmd)
			*out = &commands.CreateWatcherMacaroonCommand{
				Lnd:                lndCfg,
				MacaroonDuration:   cliCmd.Duration("macaroon.duration"),
				OutputMacaroonPath: cliCmd.String("save-to"),
			}
			return nil
		},
	}
	return cmd
}
