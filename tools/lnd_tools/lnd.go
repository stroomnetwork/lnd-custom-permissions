package lnd_tools

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/macaroons"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"
	"gopkg.in/macaroon.v2"
)

// LndGRPCConnectConfig is a struct that holds the configuration options
// for connection to the GRPC interface of an LND node.
type LndGRPCConnectConfig struct {
	// CertPath is the path to the TLS certificate for the GRPC interface LND node.
	CertPath string

	// MacaroonPath is the path to the macaroon file for the GRPC interface LND node.
	MacaroonPath string

	// RpcAddress is the address of the LND node's GRPC interface.
	RpcAddress string
}

func ConnectToLndGRPC(cfg *LndGRPCConnectConfig) (lnrpc.LightningClient, error) {
	lndMacaroonBytes, err := os.ReadFile(cfg.MacaroonPath)
	if err != nil {
		return nil, fmt.Errorf("error reading LND macaroon from file %v: %w", cfg.MacaroonPath, err)
	}

	lndTlsCreds, err := credentials.NewClientTLSFromFile(cfg.CertPath, "")
	if err != nil {
		return nil, fmt.Errorf("cannot get node TLS credentials from file %v: %w", cfg.CertPath, err)
	}

	mac := &macaroon.Macaroon{}
	if err = mac.UnmarshalBinary(lndMacaroonBytes); err != nil {
		return nil, fmt.Errorf("cannot unmarshal macaroon from binary form: %w", err)
	}

	macCredentials, err := macaroons.NewMacaroonCredential(mac)
	if err != nil {
		return nil, fmt.Errorf("cannot get macaroon credentials: %w", err)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(lndTlsCreds),
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(macCredentials),
	}

	conn, err := grpc.Dial(cfg.RpcAddress, opts...)
	if err != nil {
		return nil, fmt.Errorf("cannot dial to lnd GRPC API %v: %w", cfg.RpcAddress, err)
	}

	client := lnrpc.NewLightningClient(conn)

	return client, nil
}

// It is copy-paste from lnd/cmd/lncli/commands.go
var customDataPattern = regexp.MustCompile(
	`"custom_channel_data":\s*"([0-9a-f]+)"`,
)

// It is copy-paste from lnd/cmd/lncli/commands.go
// replaceCustomData replaces the custom channel data hex string with the
// decoded custom channel data in the JSON response.
func replaceCustomData(jsonBytes []byte) []byte {
	// If there's nothing to replace, return the original JSON.
	if !customDataPattern.Match(jsonBytes) {
		return jsonBytes
	}

	replacedBytes := customDataPattern.ReplaceAllFunc(
		jsonBytes, func(match []byte) []byte {
			encoded := customDataPattern.FindStringSubmatch(
				string(match),
			)[1]
			decoded, err := hex.DecodeString(encoded)
			if err != nil {
				return match
			}

			return []byte("\"custom_channel_data\":" +
				string(decoded))
		},
	)

	var buf bytes.Buffer
	err := json.Indent(&buf, replacedBytes, "", "    ")
	if err != nil {
		// If we can't indent the JSON, it likely means the replacement
		// data wasn't correct, so we return the original JSON.
		return jsonBytes
	}

	return buf.Bytes()
}

// It is copy-paste from lnd/cmd/lncli/commands.go
// TODO(mkl): write doc. Why it is needed?
func PrintRespJSON(resp proto.Message) {
	jsonBytes, err := lnrpc.ProtoJSONMarshalOpts.Marshal(resp)
	if err != nil {
		fmt.Println("unable to decode response: ", err)
		return
	}

	jsonBytesReplaced := replaceCustomData(jsonBytes)

	fmt.Printf("%s\n", jsonBytesReplaced)
}

// TODO(mkl): add SprintRespJSON function.
