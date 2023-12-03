package storage

import "math"

type Establishment struct {
	ID                 string
	Latitude           float64
	Longitude          float64
	AvailabilityRadios float64
}

// IsInRangeToRadius return true if the latitude and longitude is in range of establishment.
// To calculate short distances it is preferable to use the Euclidean formula instead of the Haversine.
// The curvature of the earth for short distances does not justify the extra computing time
func IsInRangeToRadius(establishment Establishment, latitude, longitude float64) bool {
	dLat := latitude - establishment.Latitude
	dLon := longitude - establishment.Longitude

	distance := math.Sqrt(dLat*dLat+dLon*dLon) * 111.32

	return distance <= establishment.AvailabilityRadios
}
