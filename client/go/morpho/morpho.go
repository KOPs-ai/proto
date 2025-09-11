package morpho

import (
	grpcClient "github.com/KOPs-ai/proto/client/go/common"
	"github.com/KOPs-ai/proto/pb/go/morpho"
)

var SVC morpho.MorphoServiceClient

func NewMorphoServiceClient(grpc_server_address string) error {
	conn, err := grpcClient.NewGRPCClientConn(grpc_server_address)
	if err != nil {
		return err
	}
	SVC = morpho.NewMorphoServiceClient(conn)
	return nil
}
