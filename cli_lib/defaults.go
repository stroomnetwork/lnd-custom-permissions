package cli_lib

import "time"

// TODO(mkl): specify root key ID for the macaroon. What is it?

const (
	DefaultLndCertPath      = "tls.cert"
	DefaultLndMacaroonPath  = "admin.macaroon"
	DefaultLndRpcAddress    = "localhost:10009"
	DefaultCustomCaveatName = "interceptor-custom-caveat"
	DefaultMiddlewareName   = "interceptor-middleware"
	DefaultMacaroonDuration = 365 * 24 * time.Hour
)

var (
	Default_AddInvoiceRule_MaxValueSat                int64 = 0
	Default_AddInvoiceRule_AllowedPaymentAddresses          = []string{}
	Default_OpenChannelRule_MaxPushAmountSat          int64 = 0
	Default_CloseChannelRule_AllowedDeliveryAddresses       = []string{}
	Default_SendPaymentV2Rule_AllowedDestinations           = []string{}
)
