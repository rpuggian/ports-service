package services

import (
	"context"
	"errors"
	proto "github.com/rpuggian/ports-service/proto"
	"github.com/rpuggian/ports-service/server/domain/contracts"
	"github.com/rpuggian/ports-service/server/domain/contracts/mocks"
	"github.com/rpuggian/ports-service/server/domain/entity"
	"reflect"
	"testing"
)

func Test_portService_Find(t *testing.T) {
	type fields struct {
		repository contracts.PortRepository
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.Port
		wantErr bool
	}{
		{
			name: "Should return port when found on repository",
			fields: fields{
				repository: func() contracts.PortRepository {
					m := &mocks.PortRepository{}
					m.On("Find", context.Background(), "1").
						Return(&entity.Port{ID: "1"}, nil)
					return m
				}(),
			},
			args: args{
				ctx: context.Background(),
				id:  "1",
			},
			want:    &proto.Port{Id: "1"},
			wantErr: false,
		},
		{
			name: "Should return error when find fails",
			fields: fields{
				repository: func() contracts.PortRepository {
					m := &mocks.PortRepository{}
					m.On("Find", context.Background(), "1").
						Return(nil, errors.New("find error"))
					return m
				}(),
			},
			args: args{
				ctx: context.Background(),
				id:  "1",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &portService{
				repository: tt.fields.repository,
			}
			got, err := p.Find(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_portService_Store(t *testing.T) {
	type fields struct {
		repository contracts.PortRepository
	}
	type args struct {
		ctx  context.Context
		port *proto.Port
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Should store port on repository",
			fields: fields{
				repository: func() contracts.PortRepository {
					m := &mocks.PortRepository{}
					m.On("Store", context.Background(), &entity.Port{ID: "1"}).
						Return(nil)
					return m
				}(),
			},
			args: args{
				ctx:  context.Background(),
				port: &proto.Port{Id: "1"},
			},
			wantErr: false,
		},
		{
			name: "Should retrun error when store fails",
			fields: fields{
				repository: func() contracts.PortRepository {
					m := &mocks.PortRepository{}
					m.On("Store", context.Background(), &entity.Port{ID: "1"}).
						Return(errors.New("store error"))
					return m
				}(),
			},
			args: args{
				ctx:  context.Background(),
				port: &proto.Port{Id: "1"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &portService{
				repository: tt.fields.repository,
			}
			if err := p.Store(tt.args.ctx, tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
