package config

type DatabaseConfig struct {
	HostName     string
	Port         uint64
	DatabaseName string
	UserName     string
	Password     string
	Dsn          string
	Type         string
}
