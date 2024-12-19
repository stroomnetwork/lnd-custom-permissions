package cli_lib

import "github.com/urfave/cli/v3"

func NewMainCli(out *interface{}) *cli.Command {
	cmd := &cli.Command{
		Name:        "lnd-custom-permissions",
		Description: "A CLI tool for creating and using LND macaroons with custom permissions.",
		Usage:       "lnd-custom-permissions [command]",
		Commands: []*cli.Command{
			NewCreateCustomMacaroonCli(out),
			NewCreateWatcherMacaroonCli(out),
			NewWatchLndCli(out),
		},
	}
	return cmd
}
