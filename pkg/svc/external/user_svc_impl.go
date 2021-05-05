package external

import (
	"context"

	"github.com/kutty-kumar/ho_oh/snorlax_v1"
)

type UserSvcImpl struct {
	snorlax_v1.UserServiceClient
}

func NewUserSvcImpl(client snorlax_v1.UserServiceClient) UserSvc {
	return &UserSvcImpl{
		client,
	}
}

func (us *UserSvcImpl) GetUserByEmailPassword(ctx context.Context, email string, password string) (snorlax_v1.UserDto, error) {
	var request snorlax_v1.GetUserByEmailAndPasswordRequest
	request.Email = email
	request.Password = password
	resp, err := us.GetUserByEmailAndPassword(ctx, &request)
	if err != nil {
		return snorlax_v1.UserDto{}, err
	}
	return *resp.Response, nil
}
