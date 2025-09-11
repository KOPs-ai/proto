package hypurrfi

import (
	grpcClient "github.com/KOPs-ai/proto/client/go/common"
	"github.com/KOPs-ai/proto/pb/go/hypurrfi"
)

var SVC hypurrfi.HypurrfiServiceClient

func NewHypurrfiServiceClient(grpc_server_address string) error {
	conn, err := grpcClient.NewGRPCClientConn(grpc_server_address)
	if err != nil {
		return err
	}
	SVC = hypurrfi.NewHypurrfiServiceClient(conn)
	return nil
}
