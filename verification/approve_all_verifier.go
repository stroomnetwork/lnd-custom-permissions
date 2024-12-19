package verification

type ApproveAllVerifier struct{}

func (a *ApproveAllVerifier) Verify(methodFullURI string, serializedReq []byte, rules *ApprovalRules) error {
	return nil
}

func (a *ApproveAllVerifier) Description() string {
	return "Approve all requests."
}

func (a *ApproveAllVerifier) DescriptionWithData(rules *ApprovalRules) string {
	return "Approve all requests."
}

var _ Verifier = (*ApproveAllVerifier)(nil)
