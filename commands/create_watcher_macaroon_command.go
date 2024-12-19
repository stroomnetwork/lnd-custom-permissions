package commands

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/macaroons"
	"github.com/stroomnetwork/lnd-custom-permissions/tools/lnd_tools"
	"gopkg.in/macaroon.v2"
)

type CreateWatcherMacaroonCommand struct {
	// Lnd is the configuration for connecting to the LND node.
	Lnd *lnd_tools.LndGRPCConnectConfig

	// MacaroonDuration is the duration for which the macaroon will be valid.
	MacaroonDuration time.Duration

	// OutputMacaroonPath is the path to the output macaroon file.
	OutputMacaroonPath string
}

func ExecuteCreateWatcherMacaroonCommand(cmd *CreateWatcherMacaroonCommand) error {
	conn, err := lnd_tools.ConnectToLndGRPC(cmd.Lnd)
	if err != nil {
		return fmt.Errorf("cannot connect to LND node: %w", err)
	}

	// For a RPC call restrictions, macaroons have the form:
	// Entity: "uri"
	// Action: "rpc_method"
	// TODO(mkl): explain why these calls are needed.
	permissions := []*lnrpc.MacaroonPermission{
		{
			Entity: "uri",
			Action: "/lnrpc.Lightning/GetInfo",
		},
		{
			Entity: "uri",
			Action: "/lnrpc.Lightning/RegisterRPCMiddleware",
		},
	}

	fmt.Println("Operations:", permissions)

	// TODO(mkl): should we add AllowExternalPermissions and RootKeyId?
	bakeReq := &lnrpc.BakeMacaroonRequest{
		Permissions:              permissions,
		RootKeyId:                0,
		AllowExternalPermissions: false,
	}
	bakeResp, err := conn.BakeMacaroon(context.TODO(), bakeReq)
	if err != nil {
		return fmt.Errorf("cannot bake macaroon: %w", err)
	}

	fmt.Println("Baked macaroon:", bakeResp)

	macaroonBytes, err := hex.DecodeString(bakeResp.Macaroon)
	if err != nil {
		return fmt.Errorf("cannot decode macaroon from hex: %w", err)
	}

	// We need to add some additional caveats to the macaroon.
	// Thus we unmarshal it first.
	// Then we add caveats.
	// Then we marshal it back.

	mac := &macaroon.Macaroon{}
	if err := mac.UnmarshalBinary(macaroonBytes); err != nil {
		return fmt.Errorf("cannot unmarshal macaroon: %w", err)
	}

	macConstraints := []macaroons.Constraint{
		macaroons.TimeoutConstraint(int64(cmd.MacaroonDuration.Seconds())),
	}

	constrainedMac, err := macaroons.AddConstraints(mac, macConstraints...)
	if err != nil {
		return fmt.Errorf("error adding constraints to macaroon: %w", err)
	}

	constrainedMacBytes, err := constrainedMac.MarshalBinary()
	if err != nil {
		return fmt.Errorf("cannot marshal constrained macaroon: %w", err)
	}

	if err := os.WriteFile(cmd.OutputMacaroonPath, constrainedMacBytes, 0600); err != nil {
		return fmt.Errorf("cannot write resulting macaroon to file %v: %w", cmd.OutputMacaroonPath, err)
	}

	return nil
}
