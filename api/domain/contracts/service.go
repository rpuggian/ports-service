package contracts

import (
	"context"
	"github.com/rpuggian/ports-service/api/domain/handlers/response"
	port "github.com/rpuggian/ports-service/proto"
	"io"
)

type PortService interface {
	FindByID(ctx context.Context, id string) (*response.FindPortResponse, error)
	UploadPortFile(ctx context.Context, fileStream io.ReadCloser) (*port.CreateResponse, error)
}
