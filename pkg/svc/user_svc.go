package svc

import (
	"github.com/kutty-kumar/ho_oh/pikachu_v1"
)

type UserSvc interface {
	GetUserByExternalId(userId string)(*pikachu_v1.GetUserByIdResponse, error)
}
