package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stroomnetwork/lnd-custom-permissions/cli_lib"
	"github.com/stroomnetwork/lnd-custom-permissions/commands"
)

func main() {
	var rez interface{}
	// TODO(mkl): fix name of the command and description.
	cmd := cli_lib.NewMainCli(&rez)

	if err := cmd.Run(context.TODO(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if err := commands.ExecuteCommand(rez); err != nil {
		fmt.Fprintf(os.Stderr, "cannot execute command: %v\n", err)
		os.Exit(1)
	}
}
