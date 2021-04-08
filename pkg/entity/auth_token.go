package entity

import (
	"bulbasur/pkg/pb"
	"database/sql"
	"github.com/kutty-kumar/db_commons/model"
)

type AuthToken struct {
	db_commons.BaseDomain
	Token          string `gorm:"unique"`
	ExpiryTimeUnit pb.TimeUnit
	ExpiryDuration uint64
	EntityId       string               `gorm:"index"`
	Attributes     []AuthTokenAttribute `gorm:"association_foreignkey:ExternalId;foreignkey:AuthTokenID"`
}

func (a *AuthToken) GetName() db_commons.DomainName {
	return "authTokens"
}

func (a *AuthToken) ToDto() interface{} {
	return pb.AuthTokenDto{
		Status:         pb.Status(a.Status),
		Token:          a.Token,
		EntityId:       a.EntityId,
		ExpiryTimeUnit: a.ExpiryTimeUnit,
		ExpiryDuration: a.ExpiryDuration,
	}
}

func (a *AuthToken) FillProperties(dto interface{}) db_commons.Base {
	authTokenDto := dto.(pb.AuthTokenDto)
	a.ExpiryDuration = authTokenDto.ExpiryDuration
	a.ExpiryTimeUnit = authTokenDto.ExpiryTimeUnit
	a.EntityId = authTokenDto.EntityId
	a.Token = authTokenDto.Token
	a.Status = int(authTokenDto.Status)
	return a
}

func (a *AuthToken) Merge(other interface{}) {
	authToken := other.(*AuthToken)
	if authToken.ExpiryDuration != 0 {
		a.ExpiryDuration = authToken.ExpiryDuration
	}
	if &authToken.ExpiryTimeUnit != nil {
		a.ExpiryTimeUnit = authToken.ExpiryTimeUnit
	}
}

func (a *AuthToken) FromSqlRow(rows *sql.Rows) (db_commons.Base, error) {
	err := rows.Scan(&a.Id, &a.ExternalId, &a.ExpiryTimeUnit, &a.ExpiryDuration, &a.EntityId, &a.Token)
	return a, err
}

func (a *AuthToken) SetExternalId(externalId string) {
	a.ExternalId = externalId
}
