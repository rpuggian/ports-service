package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rpuggian/ports-service/api/domain/contracts"
	"github.com/rpuggian/ports-service/api/domain/handlers/request"
	"github.com/rpuggian/ports-service/api/domain/handlers/response"
	port "github.com/rpuggian/ports-service/proto"
	"io"
	"log"
)

type portService struct {
	grpcClient contracts.PortServiceClient
}

func NewPortService(grpcClient contracts.PortServiceClient) *portService {
	return &portService{grpcClient: grpcClient}
}

func (p *portService) FindByID(ctx context.Context, id string) (*response.FindPortResponse, error) {
	protoPort, err := p.grpcClient.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return response.NewFindPortResponse(protoPort), nil
}

func (p *portService) UploadPortFile(ctx context.Context, fileStream io.ReadCloser) (*port.CreateResponse, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	dec := json.NewDecoder(fileStream)

	// create the stream
	stream, err := p.grpcClient.StreamCreate(ctx)
	if err != nil {
		log.Println("error when create stream: ", err)
		return nil, err
	}

	//read the first token
	_, err = dec.Token()
	if err != nil {
		log.Println("error when get first token: ", err)
		return nil, err
	}

	for dec.More() {

		//read key value token
		key, err := dec.Token()
		if err != nil {
			log.Println("error when get token key: ", err)
			return nil, err
		}

		// read port entity
		port := &request.CreatePortRequest{}
		if err := dec.Decode(port); err != nil {
			log.Println("error when decode file: ", err)
			return nil, err
		}

		//send the file by stream
		portDomain := request.ParseUploadPortRequestToProtoPortArray(fmt.Sprintf("%v", key), port)
		err = p.grpcClient.StreamSendPortFile(stream, portDomain)
		if err != nil {
			log.Println("error send stream file: ", err)
			return nil, err
		}
	}

	//close stream
	streamResponse, err := p.grpcClient.StreamClose(stream)
	if err != nil {
		log.Println("error close stream: ", err)
		return nil, err
	}

	return streamResponse, nil
}
