package deliverys

import (
	"context"
	"errors"
)

type Storage interface {
	DeliveryServicesNearLocation(ctx context.Context, latitude, longitude float64) ([]string, error)
	LoadFromFile(path string) error
}

type FileDownloader interface {
	DownloadFile(fileName string) error
}

type Manager struct {
	storage    Storage
	downloader FileDownloader
}

func NewManager(storage Storage, downloader FileDownloader) (*Manager, error) {
	if storage == nil {
		return nil, errors.New("storage can not be nil")
	}

	if downloader == nil {
		return nil, errors.New("downloader can not be nil")
	}

	return &Manager{
		storage:    storage,
		downloader: downloader,
	}, nil
}

func (m *Manager) Start() error {
	fileName := "database.csv"
	err := m.downloader.DownloadFile(fileName)
	if err != nil {
		return err
	}
	err = m.storage.LoadFromFile(fileName)
	return err
}

func (m *Manager) DeliveryServicesNearLocation(ctx context.Context, latitude, longitude float64) ([]string, error) {
	return m.storage.DeliveryServicesNearLocation(ctx, latitude, longitude)
}
