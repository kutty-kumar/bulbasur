package svc

import (
	"context"
	"github.com/kutty-kumar/ho_oh/pikachu_v1"
)

type BaseUserSvc interface {
	GetUserByEmailPassword(ctx context.Context, email string, password string) (*pikachu_v1.UserDto, error)
}

func NewUserSvc(client pikachu_v1.UserServiceClient) BaseUserSvc {
	return &UserSvc{
		client,
	}
}

type UserSvc struct {
	pikachu_v1.UserServiceClient
}

func (us *UserSvc) GetUserByEmailPassword(ctx context.Context, email string, password string) (*pikachu_v1.UserDto, error) {
	resp, err := us.UserServiceClient.GetUserByEmailAndPassword(ctx, &pikachu_v1.GetUserByEmailAndPasswordRequest{Email: email, Password: password})
	if err != nil {
		return nil, err
	}
	return resp.Response, nil
}
