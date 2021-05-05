package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"time"

	auth_domain "github.com/kutty-kumar/bulbasur/pkg/domain/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

const (
	accessTokenExpiryTimeUnit  = time.Hour
	refreshTokenExpiryTimeUnit = time.Hour
)

type AuthHelper struct {
}

func (ah *AuthHelper) GenerateAccessRefreshKeyPair(userId string) (map[string]string, error) {
	accessToken, err := ah.createToken(userId, time.Now().Add(time.Duration(viper.GetInt("jwt_config.access_token_expiry_duration_in_hours"))*accessTokenExpiryTimeUnit))
	if err != nil {
		return nil, err
	}
	refreshToken, err := ah.createToken(userId, time.Now().Add(time.Duration(viper.GetInt("jwt_config.refresh_token_expiry_duration_in_hours"))*refreshTokenExpiryTimeUnit))
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

func (ah *AuthHelper) createToken(userId string, expirationTime time.Time) (string, error) {
	var err error
	claims := &auth_domain.Claims{
		EntityId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(viper.GetString("jwt_config.secret_key")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (ah *AuthHelper) ValidateTokenExpiry(token string) (*auth_domain.Claims, bool) {
	claims := &auth_domain.Claims{}
	at, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt_config.secret_key")), nil
	})
	if err != nil || !at.Valid {
		return nil, false
	}
	return claims, true
}

func (ah *AuthHelper) EncryptAES(text string) (string, error) {
	var iv = []byte(viper.GetString("jwt_config.cipher_text"))
	fmt.Println(string(iv))
	block, err := aes.NewCipher([]byte(viper.GetString("jwt_config.secret_key")))
	if err != nil {
		return "", err
	}
	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return text, nil
}
