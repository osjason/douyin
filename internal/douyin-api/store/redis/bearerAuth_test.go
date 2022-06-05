package redis

import (
	"context"
	"douyin-app/pkg/model"
	"github.com/go-redis/redis/v8"
	"reflect"
	"testing"
)

func Test_newAuthRedisStore(t *testing.T) {
	type args struct {
		rdStore *redisStore
	}
	tests := []struct {
		name string
		args args
		want *BearerAuthStore
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newBearerAuthStore(tt.args.rdStore); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newBearerAuthStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authStore_SetJWTToken(t *testing.T) {
	type fields struct {
		rd *redis.Client
	}
	type args struct {
		ctx       context.Context
		jwtClaims *model.JWTClaims
		jwtToken  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := BearerAuthStore{
				redisClient: tt.fields.rd,
			}
			if err := a.SetJWTToken(tt.args.ctx, tt.args.jwtClaims, tt.args.jwtToken); (err != nil) != tt.wantErr {
				t.Errorf("SetJWTToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authStore_GetJWTToken(t *testing.T) {
	type fields struct {
		rd *redis.Client
	}
	type args struct {
		ctx       context.Context
		jwtClaims *model.JWTClaims
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := BearerAuthStore{
				redisClient: tt.fields.rd,
			}
			got, err := a.GetJWTToken(tt.args.ctx, tt.args.jwtClaims)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJWTToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetJWTToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authStore_RenewJWTToken(t *testing.T) {
	type fields struct {
		rd *redis.Client
	}
	type args struct {
		ctx       context.Context
		jwtClaims *model.JWTClaims
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := BearerAuthStore{
				redisClient: tt.fields.rd,
			}
			if err := a.RenewJWTToken(tt.args.ctx, tt.args.jwtClaims); (err != nil) != tt.wantErr {
				t.Errorf("RenewJWTToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authStore_DeleteJWTToken(t *testing.T) {
	type fields struct {
		rd *redis.Client
	}
	type args struct {
		ctx       context.Context
		jwtClaims *model.JWTClaims
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := BearerAuthStore{
				redisClient: tt.fields.rd,
			}
			if err := a.DeleteJWTToken(tt.args.ctx, tt.args.jwtClaims); (err != nil) != tt.wantErr {
				t.Errorf("DeleteJWTToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
