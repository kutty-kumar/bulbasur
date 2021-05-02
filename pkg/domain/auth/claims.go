package auth

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	EntityId string `json:"entity_id"`
	jwt.StandardClaims
}
