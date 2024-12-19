package verification

import (
	"fmt"
	"slices"

	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/protobuf/proto"
)

type SendToRouteV2Verifier struct{}

func (v *SendToRouteV2Verifier) Verify(methodFullURI string, serializedReq []byte, rules *ApprovalRules) error {
	sendToRouteReq := new(lnrpc.SendToRouteRequest)
	if err := proto.Unmarshal(serializedReq, sendToRouteReq); err != nil {
		return fmt.Errorf("cannot unmarshal SendToRouteRequest from SendToRouteV2 request: %w", err)
	}

	hops := sendToRouteReq.Route.GetHops()
	if len(hops) == 0 {
		return fmt.Errorf("no hops in the route")
	}

	lastHop := hops[len(hops)-1]
	if !slices.Contains(rules.SendPaymentV2.AllowedDestinations, lastHop.PubKey) {
		return fmt.Errorf("destination %v is not in the list of allowed destinations: %v", lastHop.PubKey, rules.SendPaymentV2.AllowedDestinations)
	}

	return nil
}

func (v *SendToRouteV2Verifier) Description() string {
	return "Allows SendToRoute RPC call with destination(last hop in the path) in the list of allowed destinations."
}

func (v *SendToRouteV2Verifier) DescriptionWithData(rules *ApprovalRules) string {
	return fmt.Sprintf("Allows SendToRoute RPC call with destination(last hop in the path) in the list of allowed destinations: %v", rules.SendPaymentV2.AllowedDestinations)
}

var _ Verifier = (*SendToRouteV2Verifier)(nil)
