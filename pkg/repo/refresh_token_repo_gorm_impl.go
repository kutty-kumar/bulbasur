package repo

import (
	"context"

	"github.com/kutty-kumar/bulbasur/pkg/domain/entity"

	charminder "github.com/kutty-kumar/charminder/pkg"
	"github.com/kutty-kumar/ho_oh/core_v1"
)

type RefreshTokenRepoGormImpl struct {
	charminder.BaseDao
}

func NewRefreshTokenRepoGormImpl(dao charminder.BaseDao) RefreshTokenRepoGormImpl {
	return RefreshTokenRepoGormImpl{
		dao,
	}
}

func (rtgr *RefreshTokenRepoGormImpl) Logout(ctx context.Context, token string) error {
	if err := rtgr.GetDb().Model(entity.RefreshToken{}).Where("token = ?", token).Update("status", core_v1.Status_inactive).Error; err != nil {
		return err
	}
	return nil
}

func (rtgr *RefreshTokenRepoGormImpl) GetCountByEntityIdToken(ctx context.Context, entityId string, token string) (int64, error) {
	var count int64
	err := rtgr.GetDb().Model(entity.RefreshToken{}).Where("entity_id = ? and token = ? and status = ?", entityId, token, core_v1.Status_active).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
