package thirdweb

import (
	grpcClient "github.com/KOPs-ai/proto/client/go/common"
	"github.com/KOPs-ai/proto/pb/go/thirdweb"
)

var SVC thirdweb.ThirdwebServiceClient

func NewThirdwebServiceClient(grpc_server_address string) error {
	conn, err := grpcClient.NewGRPCClientConn(grpc_server_address)
	if err != nil {
		return err
	}
	SVC = thirdweb.NewThirdwebServiceClient(conn)
	return nil
}
