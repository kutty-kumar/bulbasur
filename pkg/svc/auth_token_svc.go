package svc

import (
	"bulbasur/pkg/entity"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/kutty-kumar/db_commons/model"
	"github.com/kutty-kumar/ho_oh/bulbasur_v1"
	"github.com/kutty-kumar/ho_oh/core_v1"
)

type AuthTokenSvc struct {
	db_commons.BaseSvc
	UserService UserSvc
}

func (ats *AuthTokenSvc) Login(ctx context.Context, req *bulbasur_v1.LoginRequest) (*bulbasur_v1.LoginResponse, error) {
	authToken := entity.AuthToken{}
	userResponse, err := ats.UserService.GetUserByExternalId(req.UserId)
	if err != nil || userResponse == nil {
		return nil, err
	}
	if userResponse.Response.Status != core_v1.Status_active {
		return nil, errors.New("inactive user cannot login")
	}
	authToken.Token = uuid.NewUUID()
}

func (ats *AuthTokenSvc) Logout(ctx context.Context, req *bulbasur_v1.LogoutRequest) (*bulbasur_v1.LogoutResponse, error) {
	panic("implement me")
}

func (ats *AuthTokenSvc) RefreshToken(ctx context.Context, req *bulbasur_v1.RefreshTokenRequest) (*bulbasur_v1.RefreshTokenResponse, error) {
	panic("implement me")
}



