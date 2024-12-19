package commands

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/copier"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/macaroons"
	"github.com/stroomnetwork/lnd-custom-permissions/tools"
	"github.com/stroomnetwork/lnd-custom-permissions/tools/lnd_tools"
	"github.com/stroomnetwork/lnd-custom-permissions/tools/token"
	"github.com/stroomnetwork/lnd-custom-permissions/verification"
	"gopkg.in/macaroon.v2"
)

// TODO(mkl): specify root key ID for the macaroon. What is it?

type CreateCustomMacaroonCommand struct {
	// Lnd is the configuration for connecting to the LND node.
	Lnd *lnd_tools.LndGRPCConnectConfig

	// CustomCaveatName is the name of the custom caveat to be added to the macaroon.
	// Additional verification info is stored in the caveat.
	CustomCaveatName string

	// MacaroonDuration is the duration for which the macaroon will be valid.
	MacaroonDuration time.Duration

	// Rules is the set of rules for approving or rejecting requests.
	Rules *verification.ApprovalRules

	// OutputMacaroonPath is the path to the output macaroon file.
	OutputMacaroonPath string
}

func ExecuteCustomMacaroonCommand(cmd *CreateCustomMacaroonCommand) error {
	conn, err := lnd_tools.ConnectToLndGRPC(cmd.Lnd)
	if err != nil {
		return fmt.Errorf("cannot connect to LND node: %w", err)
	}

	getInfoResp, err := conn.GetInfo(context.TODO(), &lnrpc.GetInfoRequest{})
	if err != nil {
		return fmt.Errorf("cannot get info from LND node: %w", err)
	}
	fmt.Println("LND node info:", getInfoResp)

	// We modify the rules (e.g. replacing "self" with the actual pubkey).
	// So we make a copy of the rules to not modify the original.
	rules := new(verification.ApprovalRules)
	if err := copier.Copy(rules, cmd.Rules); err != nil {
		return fmt.Errorf("internal error. cannot copy rules: %w", err)
	}
	tools.ReplaceAll(rules.SendPaymentV2.AllowedDestinations, "self", getInfoResp.IdentityPubkey)

	spew.Dump(rules)
	encodedRules, err := token.EncodeV1(rules)
	if err != nil {
		return fmt.Errorf("internal error. cannot encode rules: %w", err)
	}
	fmt.Println("Encoded rules:", encodedRules)

	// For a RPC call restrictions, macaroons have the form:
	// Entity: "uri"
	// Action: "rpc_method"
	operations := make([]*lnrpc.MacaroonPermission, 0, len(verification.ApprovedOperations))
	for op := range verification.ApprovedOperations {
		operations = append(operations, &lnrpc.MacaroonPermission{
			Entity: "uri",
			Action: op,
		})
	}

	fmt.Println("Operations:", operations)

	// TODO(mkl): should we add AllowExternalPermissions and RootKeyId?
	bakeReq := &lnrpc.BakeMacaroonRequest{
		Permissions:              operations,
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
		macaroons.CustomConstraint(cmd.CustomCaveatName, encodedRules),
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
