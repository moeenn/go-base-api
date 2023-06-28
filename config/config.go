package config

type Config struct {
	ServerConfig   *ServerConfig
	DatabaseConfig *DatabaseConfig
	AuthConfig     *AuthConfig
}

func Load() (*Config, error) {
	config := &Config{}

	databaseConfig, err := loadDatabaseConfig()
	if err != nil {
		return config, err
	}

	authConfig, err := loadAuthConfig()
	if err != nil {
		return config, err
	}

	config.ServerConfig = loadServerConfig()
	config.DatabaseConfig = databaseConfig
	config.AuthConfig = authConfig

	return config, nil
}
