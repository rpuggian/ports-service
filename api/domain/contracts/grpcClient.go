package contracts

import (
	"context"
	proto "github.com/rpuggian/ports-service/proto"
)

type PortServiceClient interface {
	FindByID(ctx context.Context, id string) (*proto.Port, error)
	StreamCreate(ctx context.Context) (proto.PortService_CreateClient, error)
	StreamSendPortFile(stream proto.PortService_CreateClient, port *proto.Port) error
	StreamClose(stream proto.PortService_CreateClient) (*proto.CreateResponse, error)
}
