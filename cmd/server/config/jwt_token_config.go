package config

import "time"

type JwtAuthConfig struct {
	SecretKey                  string
	AccessTokenExpiryDuration  time.Duration
	AccessTokenExpiryTimeUnit  time.Duration
	AefreshTokenExpiryDuration time.Duration
	AefreshTokenExpiryTimeUnit time.Duration
}
