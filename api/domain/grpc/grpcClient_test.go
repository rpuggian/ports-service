package grpc

import (
	"context"
	"errors"
	protoport "github.com/rpuggian/ports-service/proto"
	"github.com/rpuggian/ports-service/proto/mocks"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestPortServiceClient_FindByID(t *testing.T) {
	type fields struct {
		client protoport.PortServiceClient
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *protoport.Port
		wantErr bool
	}{
		{
			name: "Should return data when port exists",
			fields: fields{
				client: func() protoport.PortServiceClient {
					client := &mocks.PortServiceClient{}
					client.On("FindByID", mock.Anything, &protoport.FindByIDRequest{
						Id: "1234",
					}).Return(&protoport.Port{
						Id:   "1234",
						Name: "MockName",
					}, nil)
					return client
				}(),
			},
			args: args{
				ctx: context.Background(),
				id:  "1234",
			},
			want:    &protoport.Port{Id: "1234", Name: "MockName"},
			wantErr: false,
		},
		{
			name: "Should return error when find request fails",
			fields: fields{
				client: func() protoport.PortServiceClient {
					client := &mocks.PortServiceClient{}
					client.On("FindByID", mock.Anything, &protoport.FindByIDRequest{
						Id: "1234",
					}).Return(nil, errors.New("mock error"))
					return client
				}(),
			},
			args: args{
				ctx: context.Background(),
				id:  "1234",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PortServiceClient{
				client: tt.fields.client,
			}
			got, err := p.FindByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPortServiceClient_StreamCreate(t *testing.T) {
	type fields struct {
		client protoport.PortServiceClient
	}
	type args struct {
		ctx   context.Context
		ports []*protoport.Port
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *protoport.CreateResponse
		wantErr bool
	}{
		{
			name: "Should stream data whithout error",
			fields: fields{
				client: func() protoport.PortServiceClient {
					stream := &mocks.PortService_CreateClient{}
					stream.On("Send", &protoport.Port{Id: "1234", Name: "MockName"}).
						Return(nil)
					stream.On("CloseAndRecv").
						Return(&protoport.CreateResponse{Total: 1}, nil)

					client := &mocks.PortServiceClient{}
					client.On("Create", mock.Anything).
						Return(stream, nil)
					return client
				}(),
			},
			args: args{
				ctx:   context.Background(),
				ports: []*protoport.Port{{Id: "1234", Name: "MockName"}},
			},
			want:    &protoport.CreateResponse{Total: 1},
			wantErr: false,
		},
		{
			name: "Should return error when Create stream fails",
			fields: fields{
				client: func() protoport.PortServiceClient {
					client := &mocks.PortServiceClient{}
					client.On("Create", mock.Anything).
						Return(nil, errors.New("create error"))
					return client
				}(),
			},
			args: args{
				ctx:   context.Background(),
				ports: []*protoport.Port{{Id: "1234", Name: "MockName"}},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Should return error when stream Send fails",
			fields: fields{
				client: func() protoport.PortServiceClient {
					stream := &mocks.PortService_CreateClient{}
					stream.On("Send", &protoport.Port{Id: "1234", Name: "MockName"}).
						Return(errors.New("send error"))

					client := &mocks.PortServiceClient{}
					client.On("Create", mock.Anything).
						Return(stream, nil)
					return client
				}(),
			},
			args: args{
				ctx:   context.Background(),
				ports: []*protoport.Port{{Id: "1234", Name: "MockName"}},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Should return error when stream CloseAndRecv fails",
			fields: fields{
				client: func() protoport.PortServiceClient {
					stream := &mocks.PortService_CreateClient{}
					stream.On("Send", &protoport.Port{Id: "1234", Name: "MockName"}).
						Return(nil)
					stream.On("CloseAndRecv").
						Return(nil, errors.New("close and receive error"))

					client := &mocks.PortServiceClient{}
					client.On("Create", mock.Anything).
						Return(stream, nil)
					return client
				}(),
			},
			args: args{
				ctx:   context.Background(),
				ports: []*protoport.Port{{Id: "1234", Name: "MockName"}},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PortServiceClient{
				client: tt.fields.client,
			}

			got, err := p.StreamCreate(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("StreamCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StreamCreate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
