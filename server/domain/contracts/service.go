package contracts

import (
	"context"
	proto "github.com/rpuggian/ports-service/proto"
)

type PortService interface {
	Find(ctx context.Context, id string) (*proto.Port, error)
	Store(ctx context.Context, port *proto.Port) error
}
