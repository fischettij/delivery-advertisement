package storage

// Implementation based on https://github.com/TannerGabriel/learning-go/blob/988b5a29a29c72b4e6a760964bdecbe0d39b0fae/algorithms/sorting/QuickSort/README.md
func quickSort(arr []*Establishment, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)

		// Recursively sort elements before partition and after partition
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []*Establishment, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j].Latitude < pivot.Latitude {
			i++

			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
