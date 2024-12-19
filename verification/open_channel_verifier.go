package verification

import (
	"fmt"

	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/protobuf/proto"
)

type OpenChannelVerifier struct{}

func (v *OpenChannelVerifier) Verify(methodFullURI string, serializedReq []byte, rules *ApprovalRules) error {
	openChannelReq := new(lnrpc.OpenChannelRequest)
	if err := proto.Unmarshal(serializedReq, openChannelReq); err != nil {
		return fmt.Errorf("cannot unmarshal OpenChannelRequest from OpenChannel request: %w", err)
	}
	if openChannelReq.PushSat > rules.OpenChannel.MaxPushAmountSat {
		return fmt.Errorf("push amount in satoshis %v is greater than the allowed maximum %v", openChannelReq.PushSat, rules.OpenChannel.MaxPushAmountSat)
	}
	return nil
}

func (v *OpenChannelVerifier) Description() string {
	return "Allows OpenChannel RPC call with push_amount_sat <= MAX_PUSH_SAT"
}

func (v *OpenChannelVerifier) DescriptionWithData(rules *ApprovalRules) string {
	return fmt.Sprintf("Allows OpenChannel RPC call with push_amount_sat <= %v sat", rules.OpenChannel.MaxPushAmountSat)
}

var _ Verifier = (*OpenChannelVerifier)(nil)
