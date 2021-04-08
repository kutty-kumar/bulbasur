package svc

import (
	"bulbasur/pkg/entity"
	"bulbasur/pkg/pb"
	"context"
	"github.com/kutty-kumar/db_commons/model"
)

type AuthTokenSvc struct {
	db_commons.BaseSvc
}

func (ats *AuthTokenSvc) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	authToken := entity.AuthToken{}
}

func (ats *AuthTokenSvc) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	panic("implement me")
}

func (ats *AuthTokenSvc) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	panic("implement me")
}



