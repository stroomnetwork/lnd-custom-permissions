package verification

// TODO(mkl): specify root key ID for the macaroon. What is it?

// AddInvoiceRule is a rule for AddInvoice LND RPC.
type AddInvoiceRule struct {
	// MaxValueSat is the maximum value of the invoice in sat.
	MaxValueSat int64

	// AllowedPaymentAddresses is a list of allowed payment addresses.
	AllowedPaymentAddresses []string
}

// OpenChannelRule is a rule for OpenChannel LND RPC.
type OpenChannelRule struct {
	// MaxPushAmountSat is the maximum amount of satoshis that can be pushed to the channel.
	MaxPushAmountSat int64
}

// SendPaymentV2Rule is a rule for SendPaymentV2 LND RPC.
type SendPaymentV2Rule struct {
	// AllowedDestinations is a list of allowed destinations.
	AllowedDestinations []string
}

// CloseChannelRule is a rule for CloseChannel LND RPC.
type CloseChannelRule struct {
	// AllowedDeliveryAddresses is a list of allowed delivery addresses.
	AllowedDeliveryAddresses []string
}

// ApprovalRules is a set of rules for approving or rejecting a request.
type ApprovalRules struct {
	// AddInvoice is the rule for AddInvoice LND RPC.
	AddInvoice *AddInvoiceRule

	// OpenChannel is the rule for OpenChannel LND RPC.
	OpenChannel *OpenChannelRule

	// CloseChannel is the rule for CloseChannel LND RPC.
	CloseChannel *CloseChannelRule

	// SendPaymentV2 is the rule for SendPaymentV2 LND RPC.
	SendPaymentV2 *SendPaymentV2Rule
}
