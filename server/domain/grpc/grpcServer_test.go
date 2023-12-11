package grpc

import (
	"context"
	"errors"
	proto "github.com/rpuggian/ports-service/proto"
	protomocks "github.com/rpuggian/ports-service/proto/mocks"
	"github.com/rpuggian/ports-service/server/domain/contracts"
	"github.com/rpuggian/ports-service/server/domain/contracts/mocks"
	"io"
	"reflect"

	"github.com/stretchr/testify/mock"
	"testing"
)

func TestPortGRPCServer_Create(t *testing.T) {
	type fields struct {
		UnimplementedPortServiceServer proto.UnimplementedPortServiceServer
		Service                        contracts.PortService
	}
	type args struct {
		stream proto.PortService_CreateServer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Should return ok when stream run without errors",
			fields: fields{
				UnimplementedPortServiceServer: proto.UnimplementedPortServiceServer{},
				Service: func() contracts.PortService {
					service := &mocks.PortService{}
					service.On("Store", mock.Anything, &proto.Port{Id: "GRPC_MOCK"}).
						Return(nil)
					return service
				}(),
			},
			args: args{
				stream: func() proto.PortService_CreateServer {
					stream := &protomocks.PortService_CreateServer{}
					stream.On("Context").
						Return(context.Background())
					stream.On("Recv").
						Return(&proto.Port{Id: "GRPC_MOCK"}, nil).Once()
					stream.On("Recv").
						Return(nil, io.EOF)
					stream.On("SendAndClose", mock.Anything).
						Return(nil)
					return stream
				}(),
			},
			wantErr: false,
		},
		{
			name: "Should return error when Recv function return error",
			fields: fields{
				UnimplementedPortServiceServer: proto.UnimplementedPortServiceServer{},
				Service: func() contracts.PortService {
					return nil
				}(),
			},
			args: args{
				stream: func() proto.PortService_CreateServer {
					stream := &protomocks.PortService_CreateServer{}
					stream.On("Context").
						Return(context.Background())
					stream.On("Recv").
						Return(&proto.Port{Id: "GRPC_MOCK"}, errors.New(""))
					return stream
				}(),
			},
			wantErr: true,
		},
		{
			name: "Should return error when Store function fails",
			fields: fields{
				UnimplementedPortServiceServer: proto.UnimplementedPortServiceServer{},
				Service: func() contracts.PortService {
					service := &mocks.PortService{}
					service.On("Store", mock.Anything, &proto.Port{Id: "GRPC_MOCK"}).
						Return(errors.New("store_error"))
					return service
				}(),
			},
			args: args{
				stream: func() proto.PortService_CreateServer {
					stream := &protomocks.PortService_CreateServer{}
					stream.On("Context").
						Return(context.Background())
					stream.On("Recv").
						Return(&proto.Port{Id: "GRPC_MOCK"}, nil).Once()
					stream.On("Recv").
						Return(nil, io.EOF)
					stream.On("SendAndClose", mock.Anything).
						Return(nil)
					return stream
				}(),
			},
			wantErr: true,
		},
		{
			name: "Should return error when SendAndClose function fails",
			fields: fields{
				UnimplementedPortServiceServer: proto.UnimplementedPortServiceServer{},
				Service: func() contracts.PortService {
					service := &mocks.PortService{}
					service.On("Store", mock.Anything, &proto.Port{Id: "GRPC_MOCK"}).
						Return(nil)
					return service
				}(),
			},
			args: args{
				stream: func() proto.PortService_CreateServer {
					stream := &protomocks.PortService_CreateServer{}
					stream.On("Context").
						Return(context.Background())
					stream.On("Recv").
						Return(&proto.Port{Id: "GRPC_MOCK"}, nil).Once()
					stream.On("Recv").
						Return(nil, io.EOF)
					stream.On("SendAndClose", mock.Anything).
						Return(errors.New("sendandclose_error"))
					return stream
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PortGRPCServer{
				UnimplementedPortServiceServer: tt.fields.UnimplementedPortServiceServer,
				Service:                        tt.fields.Service,
			}
			if err := s.Create(tt.args.stream); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPortGRPCServer_FindByID(t *testing.T) {
	type fields struct {
		UnimplementedPortServiceServer proto.UnimplementedPortServiceServer
		Service                        contracts.PortService
	}
	type args struct {
		ctx context.Context
		req *proto.FindByIDRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.Port
		wantErr bool
	}{
		{
			name: "Should return port when find by id",
			fields: fields{
				UnimplementedPortServiceServer: proto.UnimplementedPortServiceServer{},
				Service: func() contracts.PortService {
					service := &mocks.PortService{}
					service.On("Find", mock.Anything, "MOCKID").
						Return(&proto.Port{
							Id:          "MOCKID",
							Name:        "MOCKNAME",
							City:        "MOCKCITY",
							Country:     "MOCKCOUNTRY",
							Alias:       []string{"MOCKALIAS"},
							Regions:     []string{"MOCKREGION"},
							Coordinates: []float64{1, 2},
							Province:    "MOCKPROVINCE",
							Timezone:    "MOCKTIMEZONE",
							Unlocs:      []string{"MOCKUNLOCS"},
							Code:        "MOCKCODE",
						}, nil)
					return service
				}(),
			},
			args: args{
				ctx: context.Background(),
				req: &proto.FindByIDRequest{
					Id: "MOCKID",
				},
			},
			want: &proto.Port{
				Id:          "MOCKID",
				Name:        "MOCKNAME",
				City:        "MOCKCITY",
				Country:     "MOCKCOUNTRY",
				Alias:       []string{"MOCKALIAS"},
				Regions:     []string{"MOCKREGION"},
				Coordinates: []float64{1, 2},
				Province:    "MOCKPROVINCE",
				Timezone:    "MOCKTIMEZONE",
				Unlocs:      []string{"MOCKUNLOCS"},
				Code:        "MOCKCODE",
			},
			wantErr: false,
		},
		{
			name: "Should return error when find by id fails",
			fields: fields{
				UnimplementedPortServiceServer: proto.UnimplementedPortServiceServer{},
				Service: func() contracts.PortService {
					service := &mocks.PortService{}
					service.On("Find", mock.Anything, "MOCKID").
						Return(nil, errors.New("find_error"))
					return service
				}(),
			},
			args: args{
				ctx: context.Background(),
				req: &proto.FindByIDRequest{
					Id: "MOCKID",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PortGRPCServer{
				UnimplementedPortServiceServer: tt.fields.UnimplementedPortServiceServer,
				Service:                        tt.fields.Service,
			}
			got, err := s.FindByID(tt.args.ctx, tt.args.req)
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
