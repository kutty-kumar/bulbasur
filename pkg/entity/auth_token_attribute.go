package entity

import (
	"bulbasur/pkg/pb"
	"database/sql"
	"github.com/kutty-kumar/db_commons/model"
)

type AuthTokenAttribute struct {
	db_commons.BaseDomain
	Key string
	value string
	AuthTokenID string
}

func (a *AuthTokenAttribute) GetName() db_commons.DomainName {
	return "authTokenAttributes"
}

func (a *AuthTokenAttribute) ToDto() interface{} {
	return pb.AuthTokenAttributeDto{
		AuthTokenId: a.AuthTokenID,
		Value: a.value,
		Key: a.Key,
	}
}

func (a *AuthTokenAttribute) FillProperties(dto interface{}) db_commons.Base {
	atAttrDto := dto.(*pb.AuthTokenAttributeDto)
	if atAttrDto.Key != ""{
		a.Key = atAttrDto.Key
	}
	if atAttrDto.Value != ""{
		a.value = atAttrDto.Value
	}
	return a
}

func (a *AuthTokenAttribute) Merge(other interface{}) {
	otherEntity := other.(*AuthTokenAttribute)
	if otherEntity.Key != ""{
		a.Key = otherEntity.Key
	}
	if otherEntity.value != ""{
		a.value = otherEntity.value
	}
}

func (a *AuthTokenAttribute) FromSqlRow(rows *sql.Rows) (db_commons.Base, error) {
	err := rows.Scan(&a.Id, &a.ExternalId, &a.Status, &a.Key, &a.value, &a.AuthTokenID)
	return a, err
}

func (a *AuthTokenAttribute) SetExternalId(externalId string) {
	a.ExternalId = externalId
}



