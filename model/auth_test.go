package model

import (
	"testing"

	"github.com/go-redis/redis"
)

func TestAddInvalidJWT(t *testing.T) {
	type args struct {
		jwtString string
		client    *redis.Client
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddInvalidJWT(tt.args.jwtString, tt.args.client); (err != nil) != tt.wantErr {
				t.Errorf("AddInvalidJWT() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsJWTExist(t *testing.T) {
	type args struct {
		tokenString string
		client      *redis.Client
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsJWTExist(tt.args.tokenString, tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsJWTExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsJWTExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
