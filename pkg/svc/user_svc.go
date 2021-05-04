package svc

import (
	"github.com/google/uuid"
	"github.com/kutty-kumar/ho_oh/pikachu_v1"
)

type BaseUserSvc interface {
	GetUserByEmailPassword(email string, password string) (pikachu_v1.UserDto, error)
}

func NewUserSvc(client pikachu_v1.UserServiceClient) BaseUserSvc {
	return &UserSvc{
		client,
	}
}

type UserSvc struct {
	pikachu_v1.UserServiceClient
}

func (us *UserSvc) GetUserByEmailPassword(email string, password string) (pikachu_v1.UserDto, error) {
	var user pikachu_v1.UserDto
	user = pikachu_v1.UserDto{
		ExternalId: uuid.New().String(),
		FirstName:  "firstName",
		LastName:   "lastName",
	}
	return user, nil
}
