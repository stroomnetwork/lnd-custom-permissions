package verification

import "fmt"

// TODO(mkl): add explanation. Where to get all these arguments?
type Verifier interface {
	Verify(methodFullURI string, serializedReq []byte, rules *ApprovalRules) error
	Description() string
	DescriptionWithData(rules *ApprovalRules) string
}

// TODO(mkl): add explanation. Where these URIs come from?
var (
	ApprovedOperations = map[string]Verifier{
		"/lnrpc.Lightning/ConnectPeer":           &ApproveAllVerifier{},
		"/lnrpc.Lightning/DisconnectPeer":        &ApproveAllVerifier{},
		"/lnrpc.Lightning/ListPeers":             &ApproveAllVerifier{},
		"/lnrpc.Lightning/SubscribePeerEvents":   &ApproveAllVerifier{},
		"/lnrpc.Lightning/GetInfo":               &ApproveAllVerifier{},
		"/lnrpc.Lightning/GetRecoveryInfo":       &ApproveAllVerifier{},
		"/lnrpc.Lightning/DescribeGraph":         &ApproveAllVerifier{},
		"/lnrpc.Lightning/GetNodeMetrics":        &ApproveAllVerifier{},
		"/lnrpc.Lightning/GetChanInfo":           &ApproveAllVerifier{},
		"/lnrpc.Lightning/GetNodeInfo":           &ApproveAllVerifier{},
		"/lnrpc.Lightning/QueryRoutes":           &ApproveAllVerifier{},
		"/lnrpc.Lightning/GetNetworkInfo":        &ApproveAllVerifier{},
		"/lnrpc.Lightning/SubscribeChannelGraph": &ApproveAllVerifier{},
		"/lnrpc.Lightning/DebugLevel":            &ApproveAllVerifier{},
		"/lnrpc.Lightning/FeeReport":             &ApproveAllVerifier{},
		"/lnrpc.Lightning/ListPermissions":       &ApproveAllVerifier{},
		"/lnrpc.Lightning/AddInvoice":            &AddInvoiceVerifier{},
		"/lnrpc.Lightning/LookupInvoice":         &ApproveAllVerifier{},
		"/lnrpc.Lightning/ListInvoices":          &ApproveAllVerifier{},
		"/lnrpc.Lightning/OpenChannel":           &OpenChannelVerifier{},
		"/lnrpc.Lightning/BatchOpenChannel":      &ApproveAllVerifier{},
		"/lnrpc.Lightning/CloseChannel":          &CloseChannelVerifier{},
		"/lnrpc.Lightning/ListChannels":          &ApproveAllVerifier{},
		"/routerrpc.Router/SendToRouteV2":        &SendToRouteV2Verifier{},
		"/lnrpc.Lightning/UpdateChannelPolicy":   &ApproveAllVerifier{},
		"/lnrpc.Lightning/ForwardingHistory":     &ApproveAllVerifier{},
		"/lnrpc.Lightning/ListAliases":           &ApproveAllVerifier{},
	}
)

func VerifyRPC(methodFullURI string, serializedReq []byte, rules *ApprovalRules) error {
	verifier, ok := ApprovedOperations[methodFullURI]
	if !ok {
		return fmt.Errorf("method is not allowed: %v", methodFullURI)
	}
	err := verifier.Verify(methodFullURI, serializedReq, rules)
	if err != nil {
		return fmt.Errorf("verification failed: %v", err)
	}
	return nil
}
