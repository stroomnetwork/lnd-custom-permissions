package verification

import (
	"fmt"
	"slices"

	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/protobuf/proto"
)

type CloseChannelVerifier struct{}

func (v *CloseChannelVerifier) Verify(methodFullURI string, serializedReq []byte, rules *ApprovalRules) error {
	closeChannelReq := new(lnrpc.CloseChannelRequest)

	if err := proto.Unmarshal(serializedReq, closeChannelReq); err != nil {
		return fmt.Errorf("cannot unmarshal CloseChannelRequest from CloseChannel request: %w", err)
	}

	if closeChannelReq.DeliveryAddress != "" && !slices.Contains(rules.CloseChannel.AllowedDeliveryAddresses, closeChannelReq.DeliveryAddress) {
		return fmt.Errorf("delivery address %v is not in the list of allowed addresses: %v", closeChannelReq.DeliveryAddress, rules.CloseChannel.AllowedDeliveryAddresses)
	}

	return nil
}

func (v *CloseChannelVerifier) Description() string {
	return "Allows CloseChannel RPC call with delivery address in the list of allowed addresses."
}

func (v *CloseChannelVerifier) DescriptionWithData(rules *ApprovalRules) string {
	return fmt.Sprintf("Allows CloseChannel RPC call with delivery address in the list of allowed addresses: %v", rules.CloseChannel.AllowedDeliveryAddresses)
}

var _ Verifier = (*CloseChannelVerifier)(nil)
