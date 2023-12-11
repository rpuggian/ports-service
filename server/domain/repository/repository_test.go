package repository

import (
	"context"
	"errors"
	"github.com/rpuggian/ports-service/server/domain/entity"
	"github.com/rpuggian/ports-service/server/infra/redis"
	"github.com/rpuggian/ports-service/server/infra/redis/mocks"
	"reflect"
	"testing"
)

func TestPortRepository_Find(t *testing.T) {
	type fields struct {
		client redis.RedisClient
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Port
		wantErr bool
	}{
		{
			name: "Should return ok if redis client returns no error",
			fields: fields{
				client: func() redis.RedisClient {
					client := &mocks.RedisClient{}
					client.On("Find", "1", &entity.Port{}).
						Return(nil)
					return client
				}(),
			},
			args: args{
				ctx: context.Background(),
				id:  "1",
			},
			want:    &entity.Port{},
			wantErr: false,
		},
		{
			name: "Should return error if redis client returns error",
			fields: fields{
				client: func() redis.RedisClient {
					client := &mocks.RedisClient{}
					client.On("Find", "1", &entity.Port{}).
						Return(errors.New("redis_error"))
					return client
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
			r := &PortRepository{
				client: tt.fields.client,
			}
			got, err := r.Find(tt.args.ctx, tt.args.id)
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

func TestPortRepository_Store(t *testing.T) {
	type fields struct {
		client redis.RedisClient
	}
	type args struct {
		ctx    context.Context
		domain *entity.Port
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Should return ok if redis client returns no error",
			fields: fields{
				client: func() redis.RedisClient {
					client := &mocks.RedisClient{}
					client.On("Store", "1", &entity.Port{ID: "1"}, 0).
						Return(nil)
					return client
				}(),
			},
			args: args{
				ctx:    context.Background(),
				domain: &entity.Port{ID: "1"},
			},
			wantErr: false,
		},
		{
			name: "Should return error if redis client returns error",
			fields: fields{
				client: func() redis.RedisClient {
					client := &mocks.RedisClient{}
					client.On("Store", "1", &entity.Port{ID: "1"}, 0).
						Return(errors.New("redis_error"))
					return client
				}(),
			},
			args: args{
				ctx:    context.Background(),
				domain: &entity.Port{ID: "1"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PortRepository{
				client: tt.fields.client,
			}
			if err := r.Store(tt.args.ctx, tt.args.domain); (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
