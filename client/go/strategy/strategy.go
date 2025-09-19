package strategy

import (
	grpcClient "github.com/KOPs-ai/proto/client/go/common"
	strategyClient "github.com/KOPs-ai/proto/pb/go/strategy"
)

var SVC strategyClient.StrategyServiceClient

func NewStrategyServiceClient(grpc_server_address string) error {
	conn, err := grpcClient.NewGRPCClientConn(grpc_server_address)
	if err != nil {
		return err
	}
	SVC = strategyClient.NewStrategyServiceClient(conn)
	return nil
}
