package svc

import (
	"bulbasur/pkg/domain/entity"
	"bulbasur/pkg/repo"
	"context"
	"errors"
	"github.com/kutty-kumar/charminder/pkg/util"
	"github.com/kutty-kumar/ho_oh/bulbasur_v1"
	"github.com/kutty-kumar/ho_oh/core_v1"
	"github.com/spf13/viper"
)

type AuthTokenSvc struct {
	refreshTokenRepo repo.RefreshTokenRepo
	userSvc          BaseUserSvc
}

func NewAuthTokenSvc(refreshTokenRepo repo.RefreshTokenRepo, userSvc BaseUserSvc) AuthTokenSvc {
	return AuthTokenSvc{
		refreshTokenRepo: refreshTokenRepo,
		userSvc:          userSvc,
	}
}

func (ats *AuthTokenSvc) Login(ctx context.Context, req *bulbasur_v1.LoginRequest) (*bulbasur_v1.LoginResponse, error) {
	var resp bulbasur_v1.LoginResponse
	user, err := ats.userSvc.GetUserByEmailPassword(ctx, req.Email, req.Password)
	if err != nil {
		return nil, errors.New("user not found")
	}
	keyPair, err := util.GenerateAccessRefreshKeyPair(viper.GetString("jwt_config.access_token_expiry"), viper.GetString("jwt_config.refresh_token_expiry"), viper.GetString("jwt_config.secret_key"), user.ExternalId)
	if err != nil {
		return nil, errors.New("error in generating tokens")
	}
	encodedRefreshToken, err := util.EncryptAES(viper.GetString("jwt_config.cipher_key"), keyPair["refresh_token"])
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
	err, _ = ats.refreshTokenRepo.Create(ctx, &refreshToken)
	return &resp, nil
}

func (ats *AuthTokenSvc) Logout(ctx context.Context, req *bulbasur_v1.LogoutRequest) (*bulbasur_v1.LogoutResponse, error) {
	var resp bulbasur_v1.LogoutResponse
	encodedRefreshToken, err := util.EncryptAES(viper.GetString("jwt_config.cipher_key"), req.RefreshToken)
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
	claims, valid := util.ValidateTokenExpiry(viper.GetString("jwt_config.secret_key"), req.RefreshToken)
	if !valid {
		return nil, errors.New("refresh token expired")
	}
	encodedRefreshToken, err := util.EncryptAES(viper.GetString("jwt_config.cipher_key"), req.RefreshToken)
	if err != nil {
		return nil, errors.New("login required")
	}
	if count, err := ats.refreshTokenRepo.GetCountByEntityIdToken(ctx, claims.UserId, encodedRefreshToken); err != nil || count < 1 {
		return nil, errors.New("login required")
	}
	keyPair, err := util.GenerateAccessRefreshKeyPair(viper.GetString("jwt_config.access_token_expiry"), viper.GetString("jwt_config.refresh_token_expiry"), viper.GetString("jwt_config.secret_key"), claims.UserId)
	if err != nil {
		return nil, errors.New("error in generating tokens")
	}
	resp.Response = &bulbasur_v1.AuthTokenDto{
		Token:  keyPair["access_token"],
		Status: core_v1.Status_active,
	}
	return &resp, nil
}
