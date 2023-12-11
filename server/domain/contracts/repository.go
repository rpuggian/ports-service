package contracts

import (
	"context"
	"github.com/rpuggian/ports-service/server/domain/entity"
)

type PortRepository interface {
	Find(ctx context.Context, id string) (*entity.Port, error)
	Store(ctx context.Context, domain *entity.Port) error
}
