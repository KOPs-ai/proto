package strategyActivity

import (
	grpcClient "github.com/KOPs-ai/proto/client/go/common"
	"github.com/KOPs-ai/proto/pb/go/strategyActivity"
)

var SVC strategyActivity.StrategyActivityServiceClient

func NewStrategyActivityServiceClient(grpc_server_address string) error {
	conn, err := grpcClient.NewGRPCClientConn(grpc_server_address)
	if err != nil {
		return err
	}
	SVC = strategyActivity.NewStrategyActivityServiceClient(conn)
	return nil
}
