package svc

import (
	"bulbasur/pkg/domain/entity"
	"bulbasur/pkg/helper"
	"bulbasur/pkg/repo"
	"context"
	"errors"

	"github.com/kutty-kumar/ho_oh/bulbasur_v1"
	"github.com/kutty-kumar/ho_oh/core_v1"
)

type AuthTokenSvc struct {
	authHelper       helper.AuthHelper
	refreshTokenRepo repo.RefreshTokenRepo
	userSvc          BaseUserSvc
}

func NewAuthTokenSvc(refreshTokenRepo repo.RefreshTokenRepo, userSvc BaseUserSvc) AuthTokenSvc {
	return AuthTokenSvc{
		refreshTokenRepo: refreshTokenRepo,
		authHelper:       helper.AuthHelper{},
		userSvc:          userSvc,
	}
}

func (ats *AuthTokenSvc) Login(ctx context.Context, req *bulbasur_v1.LoginRequest) (*bulbasur_v1.LoginResponse, error) {
	var resp bulbasur_v1.LoginResponse
	user, err := ats.userSvc.GetUserByEmailPassword(req.Email, req.Password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	keyPair, err := ats.authHelper.GenerateAccessRefreshKeyPair(user.ExternalId)
	if err != nil {
		return nil, errors.New("error in generating tokens")
	}
	encodedRefreshToken, err := ats.authHelper.EncryptAES(keyPair["refresh_token"])
	if err != nil {
		return nil, errors.New("error in generating tokens")
	}
	resp.Response = &bulbasur_v1.AuthTokenDto{
		Token:        keyPair["access_token"],
		Status:       core_v1.Status_active,
		RefreshToken: keyPair["refresh_token"],
		EntityId:     user.ExternalId,
	}
	refreshToken := entity.RefreshToken{}
	refreshToken.FillProperties(*resp.Response)
	refreshToken.Token = encodedRefreshToken
	ats.refreshTokenRepo.Create(ctx, &refreshToken)
	return &resp, nil
}

func (ats *AuthTokenSvc) Logout(ctx context.Context, req *bulbasur_v1.LogoutRequest) (*bulbasur_v1.LogoutResponse, error) {
	var resp bulbasur_v1.LogoutResponse
	encodedRefreshToken, err := ats.authHelper.EncryptAES(req.RefreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}
	err = ats.refreshTokenRepo.Logout(ctx, encodedRefreshToken)
	if err != nil {
		return nil, errors.New("error in logging out user")
	}
	resp.Successful = true
	return &resp, nil
}

func (ats *AuthTokenSvc) RefreshToken(ctx context.Context, req *bulbasur_v1.RefreshTokenRequest) (*bulbasur_v1.RefreshTokenResponse, error) {
	var resp bulbasur_v1.RefreshTokenResponse
	claims, valid := ats.authHelper.ValidateTokenExpiry(req.RefreshToken)
	if !valid {
		return nil, errors.New("refresh token expired")
	}
	encodedRefreshToken, err := ats.authHelper.EncryptAES(req.RefreshToken)
	if err != nil {
		return nil, errors.New("login required")
	}
	if count, err := ats.refreshTokenRepo.GetCountByEntityIdToken(ctx, claims.EntityId, encodedRefreshToken); err != nil || count < 1 {
		return nil, errors.New("login required")
	}
	keyPair, err := ats.authHelper.GenerateAccessRefreshKeyPair(claims.EntityId)
	if err != nil {
		return nil, errors.New("error in generating tokens")
	}
	resp.Response = &bulbasur_v1.AuthTokenDto{
		Token:  keyPair["access_token"],
		Status: core_v1.Status_active,
	}
	return &resp, nil
}
