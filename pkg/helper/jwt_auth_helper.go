package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"time"

	auth_domain "bulbasur/pkg/domain/auth"

	"github.com/dgrijalva/jwt-go"
)

// todo: Move these to config
// init cipher block in cmd
const (
	accessSecretKey            = "test-secret-key"
	accessTokenExpiryDuration  = 4
	accessTokenExpiryTimeUnit  = time.Hour
	refreshTokenExpiryDuration = 8760
	refreshTokenExpiryTimeUnit = time.Hour
	aesCipherKey               = "aes-test-only-secret-key"
)

type AuthHelper struct {
}

func (ah *AuthHelper) GenerateAccessRefreshKeyPair(userId string) (map[string]string, error) {
	accessToken, err := ah.createToken(userId, time.Now().Add(accessTokenExpiryDuration*accessTokenExpiryTimeUnit))
	if err != nil {
		return nil, err
	}
	refreshToken, err := ah.createToken(userId, time.Now().Add(refreshTokenExpiryDuration*refreshTokenExpiryTimeUnit))
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
	token, err := at.SignedString([]byte(accessSecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (ah *AuthHelper) ValidateTokenExpiry(token string) (*auth_domain.Claims, bool) {
	claims := &auth_domain.Claims{}
	at, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecretKey), nil
	})
	if err != nil || !at.Valid {
		return nil, false
	}
	return claims, true
}

func (ah *AuthHelper) EncryptAES(text string) (string, error) {
	textInBytes := []byte(text)
	block, err := aes.NewCipher([]byte(aesCipherKey))
	if err != nil {
		return "", err
	}
	b := base64.StdEncoding.EncodeToString(textInBytes)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return text, nil
}
