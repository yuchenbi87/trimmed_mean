package trimmedmean

import (
	"sort"
)

// TrimmedMean computes the trimmed mean of a slice of float64 values.
// The trim parameter represents the proportion of values to remove from each end of the sorted slice.
// If only one trim value is provided, trimming is assumed to be symmetric.
func TrimmedMean(data []float64, trim ...float64) float64 {
	if len(data) == 0 {
		return 0
	}

	// Handle trimming proportion
	var lowerTrim, upperTrim float64
	if len(trim) == 1 {
		lowerTrim = trim[0]
		upperTrim = trim[0]
	} else if len(trim) == 2 {
		lowerTrim = trim[0]
		upperTrim = trim[1]
	} else {
		lowerTrim = 0.0
		upperTrim = 0.0
	}

	// Sort the data slice in ascending order
	sort.Float64s(data)

	// Calculate the number of elements to trim from each end
	lowerTrimCount := int(lowerTrim * float64(len(data)))
	upperTrimCount := int(upperTrim * float64(len(data)))

	// Ensure trimCount is valid
	if lowerTrimCount+upperTrimCount >= len(data) {
		return 0 // Not enough data left after trimming
	}

	// Compute the trimmed mean
	trimmedData := data[lowerTrimCount : len(data)-upperTrimCount]
	sum := 0.0
	for _, v := range trimmedData {
		sum += v
	}

	return sum / float64(len(trimmedData))
}
