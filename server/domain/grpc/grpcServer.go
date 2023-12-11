package grpc

import (
	"context"
	proto "github.com/rpuggian/ports-service/proto"
	"github.com/rpuggian/ports-service/server/domain/contracts"
	"io"
	"log"
)

type PortGRPCServer struct {
	proto.UnimplementedPortServiceServer

	Service contracts.PortService
}

func NewPortGRPCServer(service contracts.PortService) *PortGRPCServer {
	return &PortGRPCServer{
		Service: service,
	}
}

func (s *PortGRPCServer) Create(stream proto.PortService_CreateServer) error {
	ctx := stream.Context()
	var total int32

	for {
		portProto, err := stream.Recv()
		if err == io.EOF {
			log.Println("Stream ended, total ports saved: ", total)
			return stream.SendAndClose(&proto.CreateResponse{
				Total: total,
			})
		}
		if err != nil {
			log.Println("Error while reading client stream: ", err)
			return err
		}

		total++
		err = s.Service.Store(ctx, portProto)
		if err != nil {
			log.Println("Error while store proto: ", err)
			return err
		}
	}
}

func (s *PortGRPCServer) FindByID(ctx context.Context, req *proto.FindByIDRequest) (*proto.Port, error) {
	port, err := s.Service.Find(ctx, req.GetId())
	if err != nil {
		log.Println("Error while find proto by id: ", err)
		return nil, err
	}
	return port, nil
}
