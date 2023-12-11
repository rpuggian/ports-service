package grpc

import (
	"context"
	"fmt"
	proto "github.com/rpuggian/ports-service/proto"
	"google.golang.org/grpc"
)

type PortServiceClient struct {
	client proto.PortServiceClient
}

func NewPortServiceClient(conn grpc.ClientConnInterface) PortServiceClient {
	protoClient := proto.NewPortServiceClient(conn)

	return PortServiceClient{
		client: protoClient,
	}
}

func (p PortServiceClient) FindByID(ctx context.Context, id string) (*proto.Port, error) {
	request := &proto.FindByIDRequest{
		Id: id,
	}

	response, err := p.client.FindByID(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("find by id: %w", err)
	}

	return response, nil
}

func (p PortServiceClient) StreamCreate(ctx context.Context) (proto.PortService_CreateClient, error) {
	return p.client.Create(ctx)
}

func (p PortServiceClient) StreamSendPortFile(stream proto.PortService_CreateClient, port *proto.Port) error {
	return stream.Send(port)
}

func (p PortServiceClient) StreamClose(stream proto.PortService_CreateClient) (*proto.CreateResponse, error) {
	return stream.CloseAndRecv()
}
