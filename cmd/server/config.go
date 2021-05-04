package main

import "time"

type DatabaseConfig struct {
	HostName     string `json:"host_name"`
	Port         uint64 `json:"port"`
	DatabaseName string `json:"database_name"`
	UserName     string `json:"user_name"`
	Password     string `json:"password"`
	Dsn          string `json:"dsn"`
	Type         string `json:"type"`
}

type HeartBeatConfig struct {
	KeepAliveTime    uint64 `json:"keep_alive_time"`
	KeepAliveTimeOut uint64 `json:"keep_alive_timeout"`
}

type ServerConfig struct {
	Address           string `json:"address"`
	Port              string `json:"port"`
	GatewayEnable     bool   `json:"gateway_enable"`
	GatewayAddress    string `json:"gateway_address"`
	GatewayURL        string `json:"gateway_url"`
	GatewayPort       string `json:"gateway_port"`
	InternalEnable    bool   `json:"internal_enable"`
	InternalAddress   string `json:"internal_address"`
	InternalPort      string `json:"internal_port"`
	InternalHealth    string `json:"internal_health"`
	InternalReadiness string `json:"internal_readiness"`
	SwaggerPath       string `json:"swagger_path"`
}

type JwtConfig struct {
	SecretKey                  string        `json:"secret_key"`
	CipherKey                  string        `json:"cipher_key"`
	AccessTokenExpiryDuration  time.Duration `json:"access_token_expiry_duration"`
	AccessTokenExpiryTimeUnit  time.Duration `json:"access_token_expiry_time_unit"`
	RefreshTokenExpiryDuration time.Duration `json:"refresh_token_expiry_duration"`
	RefreshTokenExpiryTimeUnit time.Duration `json:"refresh_token_expiry_time_unit"`
}

type LoggingConfig struct {
	LogLevel string `json:"log_level"`
}

type UserServiceConfig struct {
	ServerAddress string `json:"server_address"`
	ServerPort    uint64 `json:"server_port"`
}

type Config struct {
	ServerConfig      ServerConfig      `json:"server_config"`
	DatabaseConfig    DatabaseConfig    `json:"database_config"`
	HeartBeatConfig   HeartBeatConfig   `json:"heartbeat_config"`
	JwtConfig         JwtConfig         `json:"jwt_config"`
	LoggingConfig     LoggingConfig     `json:"logging_config"`
	UserServiceConfig UserServiceConfig `json:"user_svc_config"`
}
