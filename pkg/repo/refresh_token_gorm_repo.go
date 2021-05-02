package repo

import (
	"bulbasur/pkg/domain/entity"
	"context"

	charminder "github.com/kutty-kumar/charminder/pkg"
	"github.com/kutty-kumar/ho_oh/core_v1"
)

type RefreshTokenGORMRepo struct {
	charminder.BaseDao
}

func NewRefreshTokenGORMRepo(dao charminder.BaseDao) RefreshTokenGORMRepo {
	return RefreshTokenGORMRepo{
		dao,
	}
}

func (rtgr *RefreshTokenGORMRepo) Logout(ctx context.Context, entityId string) error {
	if err := rtgr.GetDb().Model(entity.RefreshToken{}).Where("entity_id = ?", entityId).Update("status", core_v1.Status_inactive).Error; err != nil {
		return err
	}
	return nil
}

func (rtgr *RefreshTokenGORMRepo) GetCountByEntityIdToken(ctx context.Context, entityId string, token string) (int64, error) {
	var count int64
	err := rtgr.GetDb().Model(entity.RefreshToken{}).Where("entity_id = ? and token = ? and status = ?", entityId, token, core_v1.Status_active).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
