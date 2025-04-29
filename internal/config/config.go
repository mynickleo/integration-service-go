package config

import "integration_service/internal/utils"

type Config struct {
	Port     string
	MongoURI string
}

func Load() (*Config, error) {
	return &Config{
		Port:     utils.GetEnv("PORT", "8080"),
		MongoURI: utils.GetEnv("MONGO_URI", "mongodb://localhost:27017"),
	}, nil
}
