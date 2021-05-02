package config

type ServerConfig struct {
	Address           string
	Port              string
	GatewayEnable     bool
	GatewayAddress    string
	GatewayURL        string
	GatewayPort       string
	InternalEnable    bool
	InternalAddress   string
	InternalPort      string
	InternalHealth    string
	InternalReadiness string
}
