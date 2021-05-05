package external

import (
	"context"

	"github.com/kutty-kumar/ho_oh/snorlax_v1"
)

type UserSvc interface {
	GetUserByEmailPassword(ctx context.Context, email string, password string) (snorlax_v1.UserDto, error)
}