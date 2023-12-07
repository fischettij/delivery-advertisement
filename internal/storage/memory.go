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
	m.establishments = []*Establishment{}
	m.logger.Info("in memory population from file started")

	// Create new slice
	err := loadFromFile(m.logger, path, func(establishment *Establishment) error {
		m.establishments = append(m.establishments, establishment)
		return nil
	})
	quickSort(m.establishments, 0, len(m.establishments)-1)
	m.logger.Info("in memory population from file finished")
	return err
}

func (m *Memory) DeliveryServicesNearLocation(ctx context.Context, latitude, longitude float64) ([]string, error) {
	founded := []string{}

	leftLimit := latitude - 0.2
	rightLimit := latitude + 0.2
	// Search the first element in range latitude - 0.2 (20 left of the point)
	startIndex := binarySearch(m.establishments, 0, len(m.establishments)-1, leftLimit)

	// Search until the right limit
	for i := startIndex; i < len(m.establishments) && m.establishments[i].Latitude <= rightLimit; i++ {
		if IsInRangeToRadius(m.establishments[i], latitude, longitude) {
			founded = append(founded, m.establishments[i].ID)
		}
	}

	return founded, nil
}

func binarySearch(slice []*Establishment, low, high int, target float64) int {
	for low <= high {
		mid := low + (high-low)/2

		if slice[mid].Latitude > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}
