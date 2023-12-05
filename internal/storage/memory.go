package storage

import (
	"context"
	"errors"
)

var (
	ErrUnexpectedHeaders = errors.New("unexpected headers")

	headersFormat = []string{"id", "latitude", "longitude", "availability_radius", "open_hour", "close_hour", "rating"}
)

type Memory struct {
	logger         Logger
	establishments []*Establishment
}

func NewMemoryStorage(logger Logger) (*Memory, error) {
	if logger == nil {
		return nil, errors.New("logger cannot be nil")
	}
	return &Memory{
		logger:         logger,
		establishments: []*Establishment{},
	}, nil
}

func (m *Memory) LoadFromFile(path string) error {
	err := loadFromFile(m.logger, path, func(establishment *Establishment) error {
		m.establishments = append(m.establishments, establishment)
		return nil
	})

	return err
}

func (m *Memory) DeliveryServicesNearLocation(ctx context.Context, latitude, longitude float64) ([]string, error) {
	founded := []string{}

	for _, establishment := range m.establishments {
		if IsInRangeToRadius(establishment, latitude, longitude) {
			founded = append(founded, establishment.ID)
		}
	}
	return founded, nil
}
