package entity

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"fmt"

	"github.com/kutty-kumar/charminder/pkg"
	charminder "github.com/kutty-kumar/charminder/pkg"
	"github.com/kutty-kumar/ho_oh/bulbasur_v1"
	"github.com/kutty-kumar/ho_oh/core_v1"
)

type RefreshToken struct {
	charminder.BaseDomain
	Token    string `gorm:"unique,varchar(512)"`
	EntityId string `gorm:"index"`
}

func (r *RefreshToken) ToBytes() (*bytes.Buffer, error) {
	var rBytes bytes.Buffer
	enc := gob.NewEncoder(&rBytes)
	err := enc.Encode(*r)
	return &rBytes, err
}

func (r *RefreshToken) ToJson() (string, error) {
	rBytes, err := json.Marshal(*r)
	if err != nil {
		return "{}", err
	}
	return string(rBytes), nil
}

func (r *RefreshToken) String() string {
	return fmt.Sprintf("{\"refresh_token\": %v}", r.Token)
}

func (r *RefreshToken) GetName() pkg.DomainName {
	return "refresh_token"
}

func (r *RefreshToken) ToDto() interface{} {
	return bulbasur_v1.AuthTokenDto{
		Status:   core_v1.Status(r.Status),
		Token:    r.Token,
		EntityId: r.EntityId,
	}
}

func (r *RefreshToken) FillProperties(dto interface{}) charminder.Base {
	authDto := dto.(bulbasur_v1.AuthTokenDto)
	r.Status = int(authDto.Status)
	r.Token = authDto.RefreshToken
	r.EntityId = authDto.EntityId
	return r
}

func (r *RefreshToken) Merge(other interface{}) {
}

func (r *RefreshToken) FromSqlRow(rows *sql.Rows) (charminder.Base, error) {
	err := rows.Scan(&r.Id, &r.ExternalId, &r.EntityId, &r.Token)
	return r, err
}

func (r *RefreshToken) SetExternalId(externalId string) {
	r.ExternalId = externalId
}
