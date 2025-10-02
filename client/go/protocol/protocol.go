package protocol

import (
	grpcClient "github.com/KOPs-ai/proto/client/go/common"
	protocolClient "github.com/KOPs-ai/proto/pb/go/protocol"
)

var SVC protocolClient.ProtocolServiceClient

func NewProtocolServiceClient(grpc_server_address string) error {
	conn, err := grpcClient.NewGRPCClientConn(grpc_server_address)
	if err != nil {
		return err
	}
	SVC = protocolClient.NewProtocolServiceClient(conn)
	return nil
}
