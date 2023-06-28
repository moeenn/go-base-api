package config

type ServerConfig struct {
	Host string
	Port string
}

func loadServerConfig() *ServerConfig {
	return &ServerConfig{
		Host: "0.0.0.0",
		Port: ":5000",
	}
}
