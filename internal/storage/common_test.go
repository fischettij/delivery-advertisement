package storage_test

import (
	"github.com/fischettij/delivery-advertisement/internal/storage"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsInRangeToRadius(t *testing.T) {
	type args struct {
		establishment *storage.Establishment
		latitude      float64
		longitude     float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "given_a_latitude_and_longitude_when_establishment_is_in_range_then_return_true",
			args: args{
				establishment: &storage.Establishment{
					Latitude:           51.194253600000003,
					Longitude:          6.455508,
					AvailabilityRadios: 5,
				},
				latitude:  51.19485472542295,
				longitude: 6.4503560526581385,
			},
			want: true,
		},
		{
			name: "given_a_latitude_and_longitude_when_establishment_is_out_of_range_then_return_true",
			args: args{
				establishment: &storage.Establishment{
					Latitude:           51.194253600000003,
					Longitude:          6.455508,
					AvailabilityRadios: 5,
				},
				latitude:  52.19485472542295,
				longitude: 6.4503560526581385,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, storage.IsInRangeToRadius(tt.args.establishment, tt.args.latitude, tt.args.longitude))
		})
	}
}
