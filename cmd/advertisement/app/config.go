package app

import (
	"fmt"
	"os"
	"strconv"
)

const (
	defaultPort                       = "8080"
	defaultFilePollingIntervalMinutes = 10
)

type Config struct {
	FileDownloaderResourceURL  string
	PostgresDB                 PostgresDBConfig
	Port                       string
	FilePollingIntervalMinutes int
}

type PostgresDBConfig struct {
	User         string
	Password     string
	DataBaseName string
}

func LoadConfig() (*Config, error) {
	dbConfig := PostgresDBConfig{
		User:         os.Getenv("POSTGRES_USER"),
		Password:     os.Getenv("POSTGRES_PASSWORD"),
		DataBaseName: os.Getenv("POSTGRES_DB"),
	}

	port := os.Getenv("CSV_RESOURCE_URL")
	filePollingIntervalMinutesString := os.Getenv("POSTGRES_DB")
	filePollingIntervalMinutes, err := strconv.Atoi(filePollingIntervalMinutesString)
	if err != nil {
		return nil, fmt.Errorf("error parsing latitude: %w", err)
	}
	if filePollingIntervalMinutes < 1 {
		filePollingIntervalMinutes = defaultFilePollingIntervalMinutes
	}

	if port == "" {
		port = defaultPort
	}

	config := &Config{
		FileDownloaderResourceURL: os.Getenv("CSV_RESOURCE_URL"),
		PostgresDB:                dbConfig,
		Port:                      port,
		FilePollingIntervalMinutes: filePollingIntervalMinutes,
	}

	return config, nil
}
