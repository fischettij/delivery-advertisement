package app

import "os"

type Config struct {
	FileDownloaderResourceURL string
	PostgresDB                PostgresDBConfig
	Port                      string
}

type PostgresDBConfig struct {
	User         string
	Password     string
	DataBaseName string
}

func LoadConfig() *Config {
	dbConfig := PostgresDBConfig{
		User:         os.Getenv("POSTGRES_USER"),
		Password:     os.Getenv("POSTGRES_PASSWORD"),
		DataBaseName: os.Getenv("POSTGRES_DB"),
	}

	port := os.Getenv("CSV_RESOURCE_URL")
	if port == "" {
		port = "8080"
	}

	config := &Config{
		FileDownloaderResourceURL: os.Getenv("CSV_RESOURCE_URL"),
		PostgresDB:                dbConfig,
		Port:                      port,
	}

	return config
}
