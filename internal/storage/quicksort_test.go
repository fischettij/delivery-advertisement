package storage

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := random.Intn(100-10) + 10
	array1 := make([]*Establishment, size)
	array2 := make([]*Establishment, size)

	for i := range array1 {
		latitude := random.Float64()
		array1[i] = &Establishment{Latitude: latitude}
		array2[i] = &Establishment{Latitude: latitude}
	}

	low := 0
	high := len(array1) - 1
	quickSort(array1, low, high)

	// Sort array 2 with native function
	sort.SliceStable(array2, func(i, j int) bool {
		return array2[i].Latitude < array2[j].Latitude
	})

	for i := range array1 {
		if array1[i].Latitude != array2[i].Latitude {
			t.Fail()
		}
	}
}
