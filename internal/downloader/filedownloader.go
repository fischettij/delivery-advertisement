package downloader

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
)

type Downloader struct {
	url    string
	client *resty.Client
}

func NewDownloader(resourceURL string, client *resty.Client) (*Downloader, error) {
	if resourceURL == "" {
		return nil, errors.New("resource url cannot be empty string")
	}
	if client == nil {
		return nil, errors.New("resty client can not be nil")
	}
	return &Downloader{
		url:    resourceURL,
		client: client,
	}, nil
}

func (d *Downloader) DownloadFile(fileName string) (string, error) {
	response, err := d.client.R().Get(d.url)
	if err != nil {
		return "", err
	}

	if response.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("failed to download CSV. status: %s", response.Status())
	}

	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(response.Body())
	if err != nil {
		return "", err
	}

	hasher := md5.New()
	hasher.Write(response.Body())
	hash := hex.EncodeToString(hasher.Sum(nil))

	return hash, nil
}

func (d *Downloader) RemoveFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return fmt.Errorf("error removing CSV file: %w", err)
	}
	return nil
}
