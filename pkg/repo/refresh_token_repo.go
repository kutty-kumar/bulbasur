package repo

import (
	"context"

	charminder "github.com/kutty-kumar/charminder/pkg"
)

type RefreshTokenRepo interface {
	charminder.BaseRepository
	Logout(ctx context.Context, token string) error
	GetCountByEntityIdToken(ctx context.Context, entityId string, token string) (int64, error)
}
