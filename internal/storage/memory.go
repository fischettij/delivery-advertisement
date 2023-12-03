package storage

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	"os"
	"strconv"
)

var (
	ErrUnexpectedHeaders = errors.New("unexpected headers")

	headersFormat = []string{"id", "latitude", "longitude", "availability_radius", "open_hour", "close_hour", "rating"}
)

type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}

type Memory struct {
	logger         Logger
	establishments []Establishment
}

func NewMemoryStorage(logger Logger) (*Memory, error) {
	if logger == nil {
		return nil, errors.New("logger cannot be nil")
	}
	return &Memory{
		logger:         logger,
		establishments: []Establishment{},
	}, nil
}

func (m *Memory) LoadFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	headers, err := reader.Read()
	if err != nil {
		return fmt.Errorf("error reading CSV headers: %w", err)
	}
	err = m.validateHeaders(headers)
	if err != nil {
		return err
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break // End of File
		}
		if err != nil {
			return fmt.Errorf("error reading CSV record: %w", err)
		}

		latitude, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			m.logger.Error(fmt.Sprintf("error on record: %s error parsing latitude: %s", record, err.Error()))
			continue
		}
		longitude, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			m.logger.Error(fmt.Sprintf("error on record: %s error parsing longitude: %s", record, err.Error()))
			continue
		}
		availabilityRadios, err := strconv.ParseFloat(record[3], 32)
		if err != nil {
			m.logger.Error(fmt.Sprintf("error on record: %s error parsing availability radios: %s", record, err.Error()))
			continue
		}

		establishment := Establishment{
			ID:                 record[0],
			Latitude:           latitude,
			Longitude:          longitude,
			AvailabilityRadios: availabilityRadios,
		}

		m.establishments = append(m.establishments, establishment)
	}
	return nil
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

func (m *Memory) validateHeaders(headers []string) error {
	if len(headers) < 4 {
		return ErrUnexpectedHeaders
	}

	for i := 0; i <= 3; i++ {
		if headers[i] != headersFormat[i] {
			return ErrUnexpectedHeaders
		}
	}
	return nil
}
