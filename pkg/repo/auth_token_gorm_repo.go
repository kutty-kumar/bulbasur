package repo

import (
	"bulbasur/pkg/entity"
	"github.com/kutty-kumar/db_commons/model"
)

type AuthTokenGORMRepo struct {
	db_commons.BaseRepository
}

func (a *AuthTokenGORMRepo) findByEntityIdAndToken(entityId string, token string) (*entity.AuthToken, error) {
	authToken := entity.AuthToken{}
	if err := a.GetDb().Where("entity_id = ? AND token = ?").Find(&authToken).Error; err != nil {
		return nil, err
	}
	return &authToken, nil
}

func (a *AuthTokenGORMRepo) logout(userId string, token string) error {
	authToken, err := a.findByEntityIdAndToken(userId, token)
	if err != nil {
		return err
	}
	authToken.Status = 2
	err, _ = a.Update(authToken.ExternalId, authToken)
	if err != nil {
		return err
	}
	return nil
}
