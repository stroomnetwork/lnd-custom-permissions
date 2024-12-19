package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type LndConfig struct {
	Host string `json:"host"`
}
type Config struct {
	Lnd      LndConfig `json:"lnd"`
	MaxValue int       `json:"max_value"`
}

// TestEncodeDecode tests that decoding encoded data returns the original data.
func TestEncodeDecode(t *testing.T) {
	cfg := &Config{
		Lnd: LndConfig{
			Host: "localhost:10009",
		},
		MaxValue: 100,
	}

	encoded, err := EncodeV1(cfg)
	if err != nil {
		t.Fatalf("Error encoding data: %v", err)
	}

	decodedCfg := new(Config)
	err = DecodeV1(encoded, decodedCfg)
	if err != nil {
		t.Fatalf("Error decoding data: %v", err)
	}

	assert.Equal(t, cfg, decodedCfg)
}

// TestEncode tests encoding data to base64 encoded JSON.
func TestEncode(t *testing.T) {
	cfg := &Config{
		Lnd: LndConfig{
			Host: "localhost:10009",
		},
		MaxValue: 100,
	}

	encoded, err := EncodeV1(cfg)
	if err != nil {
		t.Fatalf("Error encoding data: %v", err)
	}

	assert.Equal(t, "1-eyJsbmQiOnsiaG9zdCI6ImxvY2FsaG9zdDoxMDAwOSJ9LCJtYXhfdmFsdWUiOjEwMH0=", encoded)
}

func TestDecode(t *testing.T) {
	decodedCfg := new(Config)
	err := DecodeV1("1-eyJsbmQiOnsiaG9zdCI6ImxvY2FsaG9zdDoxMDAwOSJ9LCJtYXhfdmFsdWUiOjEwMH0=", decodedCfg)
	if err != nil {
		t.Fatalf("Error decoding data: %v", err)
	}

	cfg := &Config{
		Lnd: LndConfig{
			Host: "localhost:10009",
		},
		MaxValue: 100,
	}

	assert.Equal(t, cfg, decodedCfg)
}

func TestDecodeError(t *testing.T) {
	decodedCfg := new(Config)
	err := DecodeV1("2-eyJsbmQiOnsiaG9zdCI6ImxvY2FsaG9zdDoxMDAwOSJ9LCJtYXhfdmFsdWUiOjEwMH0=", decodedCfg)
	assert.EqualError(t, err, "unsupported version(only version 1 is supported): 2")

	err = DecodeV1("Y2FsaG9zdDoxMDAwOSJ9LCJtYXhfdmFsdWUiOjEwMH0=", decodedCfg)
	assert.EqualError(t, err, "version separator '-' not found in data: Y2FsaG9zdDoxMDAwOSJ9LCJtYXhfdmFsdWUiOjEwMH0=")

	err = DecodeV1("1-ey", decodedCfg)
	assert.EqualError(t, err, "failed to decode base64: illegal base64 data at input byte 0")
}
