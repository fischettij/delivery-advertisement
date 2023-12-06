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
	Port                       string
	FilePollingIntervalMinutes int
}

func LoadConfig() (*Config, error) {
	port := os.Getenv("PORT")

	filePollingIntervalMinutesString := os.Getenv("FILE_POLLING_INTERVAL_MINUTES")
	filePollingIntervalMinutes, err := strconv.Atoi(filePollingIntervalMinutesString)
	if err != nil {
		return nil, fmt.Errorf("error parsing FILE_POLLING_INTERVAL_MINUTES: %w", err)
	}
	if filePollingIntervalMinutes < 1 {
		filePollingIntervalMinutes = defaultFilePollingIntervalMinutes
	}

	if port == "" {
		port = defaultPort
	}

	config := &Config{
		FileDownloaderResourceURL:  os.Getenv("CSV_RESOURCE_URL"),
		Port:                       port,
		FilePollingIntervalMinutes: filePollingIntervalMinutes,
	}

	return config, nil
}
