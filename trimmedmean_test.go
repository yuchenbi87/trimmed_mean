package trimmedmean

import (
	"testing"
)

func TestTrimmedMean(t *testing.T) {
	data := []float64{10, 8, 12, 10, 9, 7, 25, 23, 11, 9, 18, 19, 15, 11, 12, 14, 23, 19}
	result := TrimmedMean(data, 0.1)
	expected := 13.9375
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	result = TrimmedMean(data, 0.1, 0.2)
	expected = 12.642857142857142
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
