package token

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// EncodeV1 encodes data to base64 encoded JSON
// It is prefixed with "1-" to indicate version 1.
func EncodeV1(data interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal data to JSON: %w", err)
	}
	base64Encoded := base64.StdEncoding.EncodeToString(b)
	return "1-" + base64Encoded, nil
}

func DecodeV1(s string, data interface{}) error {
	verStr, encodedStr, found := strings.Cut(s, "-")
	if !found {
		return fmt.Errorf("version separator '-' not found in data: %s", s)
	}
	if verStr != "1" {
		return fmt.Errorf("unsupported version(only version 1 is supported): %s", verStr)
	}
	b, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		return fmt.Errorf("failed to decode base64: %w", err)
	}
	if err := json.Unmarshal(b, data); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return nil
}
