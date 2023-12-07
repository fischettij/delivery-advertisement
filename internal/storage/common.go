package storage

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"

	"go.uber.org/zap"
)

type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}

type Establishment struct {
	ID                 string  `db:"establishment_id"`
	Latitude           float64 `db:"latitude"`
	Longitude          float64 `db:"longitude"`
	AvailabilityRadios float64 `db:"availability_radius"`
}

// IsInRangeToRadius return true if the latitude and longitude is in range of establishment.
// To calculate short distances it is preferable to use the Euclidean formula instead of the Haversine.
// The curvature of the earth for short distances does not justify the extra computing time
func IsInRangeToRadius(establishment *Establishment, latitude, longitude float64) bool {
	dLat := latitude - establishment.Latitude
	dLon := longitude - establishment.Longitude

	distance := math.Sqrt(dLat*dLat+dLon*dLon) * 111.32

	return distance <= establishment.AvailabilityRadios
}

func validateHeaders(headers []string) error {
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

func loadFromFile(logger Logger, path string, insertFunction func(establishment *Establishment) error) error {
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
	err = validateHeaders(headers)
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
			logger.Error(fmt.Sprintf("error on record: %s error parsing latitude: %s", record, err.Error()))
			continue
		}
		longitude, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			logger.Error(fmt.Sprintf("error on record: %s error parsing longitude: %s", record, err.Error()))
			continue
		}
		availabilityRadios, err := strconv.ParseFloat(record[3], 32)
		if err != nil {
			logger.Error(fmt.Sprintf("error on record: %s error parsing availability radios: %s", record, err.Error()))
			continue
		}

		establishment := &Establishment{
			ID:                 record[0],
			Latitude:           latitude,
			Longitude:          longitude,
			AvailabilityRadios: availabilityRadios,
		}
		err = insertFunction(establishment)
		if err != nil {
			return err
		}
	}
	return nil
}

/*


func insertSort(list *[]Establishment, element Establishment) *[]string {

	"""
	Using divide and conquer
	"""
	left := 0
	right := len(list) - 1
	for left <= right{
		mid := (left + right) / 2
		if list[mid].Latitude < element.Latitude

	}

	if lst[mid][2] == element[2]:
	lst.insert(mid, element)
	return
	elif lst[mid][2] < element[2]:
	left = mid + 1
	else:
	right = mid - 1
	lst.insert(left, element)

	return nil
}
*/
