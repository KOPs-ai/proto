package user

import (
	grpcClient "github.com/KOPs-ai/proto/client/go/common"
	"github.com/KOPs-ai/proto/pb/go/user"
)

var SVC user.UserServiceClient

func NewUserServiceClient(grpc_server_address string) error {
	conn, err := grpcClient.NewGRPCClientConn(grpc_server_address)
	if err != nil {
		return err
	}
	SVC = user.NewUserServiceClient(conn)
	return nil
}
