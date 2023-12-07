package deliverys

import (
	"context"
	"errors"
	"time"
)

type Storage interface {
	DeliveryServicesNearLocation(ctx context.Context, latitude, longitude float64) ([]string, error)
	LoadFromFile(path string) error
}

type FileDownloader interface {
	DownloadFile(fileName string) (string, error)
	RemoveFile(name string) error
}

type Manager struct {
	storage             Storage
	downloader          FileDownloader
	csvVersion          string
	filePollingInterval time.Duration
}

func NewManager(storage Storage, downloader FileDownloader, filePollingInterval time.Duration) (*Manager, error) {
	if storage == nil {
		return nil, errors.New("storage can not be nil")
	}

	if downloader == nil {
		return nil, errors.New("downloader can not be nil")
	}

	return &Manager{
		storage:             storage,
		downloader:          downloader,
		filePollingInterval: filePollingInterval,
	}, nil
}

// Start Every 10 minutes it validates if there is a new version of the file and updates it if necessary
func (m *Manager) Start(done chan<- error) {
	go func() {
		for {
			fileName := "database.csv"
			md5, err := m.downloader.DownloadFile(fileName)
			if err != nil {
				done <- err
				return
			}
			if md5 != m.csvVersion {
				m.csvVersion = md5
				err = m.storage.LoadFromFile(fileName)
				if err != nil {
					done <- err
					return
				}
			}
			err = m.downloader.RemoveFile(fileName)
			if err != nil {
				done <- err
				return
			}
			time.Sleep(m.filePollingInterval)
		}
	}()
}

func (m *Manager) DeliveryServicesNearLocation(ctx context.Context, latitude, longitude float64) ([]string, error) {
	return m.storage.DeliveryServicesNearLocation(ctx, latitude, longitude)
}
