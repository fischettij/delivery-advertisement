package app

import "os"

type Config struct {
	FileDownloaderResourceURL string
}

func LoadConfig() *Config {
	resourceURL := os.Getenv("CSV_RESOURCE_URL")
	config := &Config{
		FileDownloaderResourceURL: resourceURL,
	}
	return config
}
