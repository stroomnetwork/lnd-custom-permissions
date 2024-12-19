package cli_lib

import (
	"context"

	"github.com/stroomnetwork/lnd-custom-permissions/commands"
	"github.com/stroomnetwork/lnd-custom-permissions/verification"
	cli "github.com/urfave/cli/v3"
)

func NewCreateCustomMacaroonCli(out *interface{}) *cli.Command {
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
		&cli.StringFlag{
			Name:  "macaroon.custom-caveat-name",
			Usage: "Name `CAVEAT_NAME` of the custom caveat to be added to the macaroon.",
			Value: DefaultCustomCaveatName,
		},

		&cli.DurationFlag{
			Name:  "macaroon.duration",
			Usage: "Duration `DURATION` for which the macaroon will be valid.",
			Value: DefaultMacaroonDuration,
		},

		&cli.IntFlag{
			Name:  "rules.add-invoice.max-value-sat",
			Usage: "Maximum value `MAX_VALUE_SAT` of the invoice in sat. ",
			Value: Default_AddInvoiceRule_MaxValueSat,
		},

		&cli.StringSliceFlag{
			Name:  "rules.add-invoice.allowed-payment-addresses",
			Usage: "List of allowed payment addresses `PAYMENT_ADDRESS`.",
			Value: Default_AddInvoiceRule_AllowedPaymentAddresses,
		},

		&cli.IntFlag{
			Name:  "rules.open-channel.max-push-amount-sat",
			Usage: "Maximum amount `MAX_PUSH_AMOUNT_SAT` of satoshis that can be pushed to the channel.",
			Value: Default_OpenChannelRule_MaxPushAmountSat,
		},

		&cli.StringSliceFlag{
			Name:  "rules.close-channel.allowed-delivery-address",
			Usage: "List of allowed delivery addresses `DELIVERY_ADDRESS`.",
			Value: Default_CloseChannelRule_AllowedDeliveryAddresses,
		},

		&cli.StringSliceFlag{
			Name:  "rules.send-payment-v2.allowed-destination",
			Usage: "List of allowed destinations `DESTINATION`.",
			Value: Default_SendPaymentV2Rule_AllowedDestinations,
		},
	)
	description := "Create a custom macaroon with limitations for some operations.\n"
	description += "The following operations are allowed: \n"
	for uri, ver := range verification.ApprovedOperations {
		description += "\n" + uri + " :\n"
		description += "  " + ver.Description() + "\n"
	}
	cmd := &cli.Command{
		Name:        "create-custom-macaroon",
		Description: description,
		Flags:       flags,
		Action: func(ctx context.Context, cliCmd *cli.Command) error {
			lndCfg := ExtractLndGRPCConnectConfig(cliCmd)
			rules := &verification.ApprovalRules{
				AddInvoice: &verification.AddInvoiceRule{
					MaxValueSat:             int64(cliCmd.Int("rules.add-invoice.max-value-sat")),
					AllowedPaymentAddresses: cliCmd.StringSlice("rules.add-invoice.allowed-payment-address"),
				},
				OpenChannel: &verification.OpenChannelRule{
					MaxPushAmountSat: int64(cliCmd.Int("rules.open-channel.max-push-amount-sat")),
				},
				CloseChannel: &verification.CloseChannelRule{
					AllowedDeliveryAddresses: cliCmd.StringSlice("rules.close-channel.allowed-delivery-address"),
				},
				SendPaymentV2: &verification.SendPaymentV2Rule{
					AllowedDestinations: cliCmd.StringSlice("rules.send-payment-v2.allowed-destination"),
				},
			}
			*out = &commands.CreateCustomMacaroonCommand{
				Lnd:                lndCfg,
				CustomCaveatName:   cliCmd.String("macaroon.custom-caveat-name"),
				MacaroonDuration:   cliCmd.Duration("macaroon.duration"),
				Rules:              rules,
				OutputMacaroonPath: cliCmd.String("save-to"),
			}
			return nil
		},
	}
	return cmd
}
