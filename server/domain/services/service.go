package services

import (
	"context"
	proto "github.com/rpuggian/ports-service/proto"
	"github.com/rpuggian/ports-service/server/domain/contracts"
	"github.com/rpuggian/ports-service/server/domain/parsers"
)

type portService struct {
	repository contracts.PortRepository
}

func NewPortService(repo contracts.PortRepository) *portService {
	return &portService{
		repository: repo,
	}
}

func (p *portService) Find(ctx context.Context, id string) (*proto.Port, error) {
	protoPort, err := p.repository.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return parsers.ParsePortEntityToProtoResponse(protoPort), nil
}

func (p *portService) Store(ctx context.Context, port *proto.Port) error {
	portEntity := parsers.ParsePortProtoRequestToEntity(port)
	return p.repository.Store(ctx, portEntity)
}
