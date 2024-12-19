package verification

import (
	"fmt"

	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/protobuf/proto"
)

type AddInvoiceVerifier struct{}

func (v *AddInvoiceVerifier) Verify(methodFullURI string, serializedReq []byte, rules *ApprovalRules) error {
	invoiceReq := new(lnrpc.Invoice)
	if err := proto.Unmarshal(serializedReq, invoiceReq); err != nil {
		return fmt.Errorf("cannot unmarshal Invoice from AddInvoice request: %w", err)
	}
	if invoiceReq.Value > rules.AddInvoice.MaxValueSat {
		return fmt.Errorf("invoice value in satoshis %v is greater than the allowed maximum %v", invoiceReq.Value, rules.AddInvoice.MaxValueSat)
	}
	if invoiceReq.ValueMsat > rules.AddInvoice.MaxValueSat*1000 {
		return fmt.Errorf("invoice value in millisatoshis %v is greater than the allowed maximum %v", invoiceReq.ValueMsat, rules.AddInvoice.MaxValueSat*1000)
	}

	return nil
}

func (v *AddInvoiceVerifier) Description() string {
	return "Allows AddInvoice RPC call with value<=MAX_INVOICE_VALUE and destination in the list of allowed destinations."
}

func (v *AddInvoiceVerifier) DescriptionWithData(rules *ApprovalRules) string {
	return fmt.Sprintf("Allows AddInvoice RPC call with value<=%v sat and destination in the list of allowed destinations: %v", rules.AddInvoice.MaxValueSat, rules.AddInvoice.AllowedPaymentAddresses)
}

var _ Verifier = (*AddInvoiceVerifier)(nil)
