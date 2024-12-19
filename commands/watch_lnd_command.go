package commands

import (
	"context"
	"fmt"

	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/stroomnetwork/lnd-custom-permissions/tools/lnd_tools"
	"github.com/stroomnetwork/lnd-custom-permissions/tools/token"
	"github.com/stroomnetwork/lnd-custom-permissions/verification"
)

type WatchLndCommand struct {
	// Lnd is the configuration for connecting to the LND node.
	Lnd *lnd_tools.LndGRPCConnectConfig

	// TODO(mkl): explain better.
	// CustomCaveatName is the name of the custom caveat for watch RPC call.
	// Additional verification info is stored in the caveat.
	CustomCaveatName string

	// MiddlewareName is the name of middleware name.
	// TODO(mkl): where it is used? Can it be arbitrary string.
	MiddlewareName string
}

func ExecuteWatchLndCommand(cmd *WatchLndCommand) error {
	conn, err := lnd_tools.ConnectToLndGRPC(cmd.Lnd)
	if err != nil {
		return fmt.Errorf("cannot connect to LND node: %w", err)
	}

	getInfoResp, err := conn.GetInfo(context.TODO(), &lnrpc.GetInfoRequest{})
	if err != nil {
		return fmt.Errorf("cannot get info from LND node: %w", err)
	}
	lnd_tools.PrintRespJSON(getInfoResp)

	// TODO(mkl): use global context(which should cancel on os signals).
	middlewareStream, err := conn.RegisterRPCMiddleware(context.Background())
	if err != nil {
		return fmt.Errorf("cannot register RPC middleware: %w", err)
	}
	fmt.Println("RPC middleware stream: ", middlewareStream)

	// This message should be sent to the server immediately after the stream is opened.
	// Because the server will close the stream if it does not receive this message after a certain period of time.
	registrationMsg := &lnrpc.RPCMiddlewareResponse{
		MiddlewareMessage: &lnrpc.RPCMiddlewareResponse_Register{
			Register: &lnrpc.MiddlewareRegistration{
				MiddlewareName:           cmd.MiddlewareName,
				CustomMacaroonCaveatName: cmd.CustomCaveatName,
			},
		},
	}

	if err := middlewareStream.Send(registrationMsg); err != nil {
		return fmt.Errorf("cannot send registration message to the server: %w", err)
	}

	for {
		middlewareReq, err := middlewareStream.Recv()
		if err != nil {
			return fmt.Errorf("cannot receive middleware request: %w", err)
		}
		lnd_tools.PrintRespJSON(middlewareReq)

		finalErr := ""

		fmt.Printf("Intercept type: %T\n", middlewareReq.InterceptType)

	switch_label:
		switch req := middlewareReq.InterceptType.(type) {
		case *lnrpc.RPCMiddlewareRequest_Request:

			var rules verification.ApprovalRules
			err := token.DecodeV1(middlewareReq.CustomCaveatCondition, &rules)
			if err != nil {
				fmt.Printf("cannot decode custom caveat: %v\n", err)
				break switch_label
			}

			if err := verification.VerifyRPC(req.Request.MethodFullUri, req.Request.Serialized, &rules); err != nil {
				finalErr = fmt.Sprintf("interceptor rejected this request: %v", err)
			} else {
				finalErr = ""
			}
		default:
			finalErr = ""
		}

		middleWareResp := &lnrpc.RPCMiddlewareResponse{
			RefMsgId: middlewareReq.MsgId,
			MiddlewareMessage: &lnrpc.RPCMiddlewareResponse_Feedback{
				Feedback: &lnrpc.InterceptFeedback{
					Error: finalErr,
				},
			},
		}
		err = middlewareStream.Send(middleWareResp)
		if err != nil {
			return fmt.Errorf("cannot send middleware response: %w", err)
		}
	}
}
