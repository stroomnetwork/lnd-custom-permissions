package verification

import (
	"fmt"

	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/stroomnetwork/lnd-custom-permissions/tools/lnd_tools"
	"google.golang.org/protobuf/proto"
)

type SendCoinsVerifier struct{}

func (s *SendCoinsVerifier) Verify(methodFullURI string, serializedReq []byte, rules *ApprovalRules) error {
	fmt.Println("Intercepted SendCoins request.")
	sendCoinsReq := new(lnrpc.SendCoinsRequest)
	if err := proto.Unmarshal(serializedReq, sendCoinsReq); err != nil {
		return fmt.Errorf("cannot unmarshal SendCoins request: %w", err)
	}
	lnd_tools.PrintRespJSON(sendCoinsReq)
	threshold := 1000000
	if sendCoinsReq.Amount <= int64(threshold) {
		fmt.Printf("APPROVING. Requested amount %d is less than threshold %d\n", sendCoinsReq.Amount, threshold)
		return nil
	} else {

		fmt.Printf("REJECTING. Requested amount %d is greater than threshold %d\n", sendCoinsReq.Amount, threshold)
		return fmt.Errorf("interceptor rejected this request. Amount is greater than threshold. amount: %d, threshold: %d", sendCoinsReq.Amount, threshold)
	}
}

func (s *SendCoinsVerifier) Description() string {
	return "Approving SendCoins requests with value lower than some amount."
}

func (s *SendCoinsVerifier) DescriptionWithData(rules *ApprovalRules) string {
	return "Approving SendCoins requests with value lower than 1000000 sat."
}

var _ Verifier = (*SendCoinsVerifier)(nil)
