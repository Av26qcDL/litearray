package litearray

import (
	"testing" // Import the testing package
)

func TestAddArrays(t *testing.T) {
	// Test case 1: Regular arrays, precision = 2
	arr1 := []float64{1.12345, 2.98765, 3.14159}
	arr2 := []float64{4.87654, 5.01235, 6.85841}
	expected := []float64{6.00, 8.00, 10.00}

	result, err := AddArrays(2, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Adjusted Test Case 2: Precision = -1 (no rounding)
	arr1 = []float64{1.12345, 2.98765, 3.14159}
	arr2 = []float64{4.87654, 5.01235, 6.85841}
	expected = []float64{6.0, 8.0, 10.0} // Updated for clarity

	result, err = AddArrays(-1, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Mismatched array lengths
	arr3 := []float64{1.0}
	_, err = AddArrays(2, arr1, arr3)
	if err == nil {
		t.Error("Expected an error for mismatched array lengths, got none")
	}

	// Test case 4: Fewer than two arrays
	_, err = AddArrays(2, arr1)
	if err == nil {
		t.Error("Expected an error for fewer than two arrays, got none")
	}
}

func TestSubtractArrays(t *testing.T) {
	// Test case 1: Regular arrays, precision = 2
	arr1 := []float64{9.11111, 8.11111, 7.11111}
	arr2 := []float64{8.11111, 7.11111, 6.11111}
	expected := []float64{1.00, 1.00, 1.00}

	result, err := SubtractArrays(2, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Adjusted Test Case 2: Precision = -1 (no rounding)
	arr1 = []float64{9.11111, 8.11111, 7.11111}
	arr2 = []float64{8.11111, 7.11111, 6.11111}
	expected = []float64{1.00, 1.00, 1.00}

	result, err = SubtractArrays(-1, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Mismatched array lengths
	arr3 := []float64{1.0}
	_, err = SubtractArrays(2, arr1, arr3)
	if err == nil {
		t.Error("Expected an error for mismatched array lengths, got none")
	}

	// Test case 4: Fewer than two arrays
	_, err = SubtractArrays(2, arr1)
	if err == nil {
		t.Error("Expected an error for fewer than two arrays, got none")
	}
}

func TestMultiplyArrays(t *testing.T) {
	// Test case 1: Regular arrays, precision = 2
	arr1 := []float64{10.00000, 8.00000, 7.00000}
	arr2 := []float64{9.00000, 7.00000, 6.00000}
	expected := []float64{90.00, 56.00, 42.00}

	result, err := MultiplyArrays(2, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Adjusted Test Case 2: Precision = -1 (no rounding)
	arr1 = []float64{10.00000, 8.00000, 7.00000}
	arr2 = []float64{9.00000, 7.00000, 6.00000}
	expected = []float64{90.00, 56.00, 42.00}

	result, err = MultiplyArrays(-1, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Mismatched array lengths
	arr3 := []float64{1.0}
	_, err = MultiplyArrays(2, arr1, arr3)
	if err == nil {
		t.Error("Expected an error for mismatched array lengths, got none")
	}

	// Test case 4: Fewer than two arrays
	_, err = MultiplyArrays(2, arr1)
	if err == nil {
		t.Error("Expected an error for fewer than two arrays, got none")
	}
}

func TestDivideArrays(t *testing.T) {
	// Test case 1: Regular arrays, precision = 2
	arr1 := []float64{10.00000, 8.00000, 7.00000}
	arr2 := []float64{9.00000, 7.00000, 6.00000}
	expected := []float64{1.11, 1.14, 1.17}

	result, err := DivideArrays(2, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Adjusted Test Case 2: Precision = -1 (no rounding)
	arr1 = []float64{10.00000, 8.00000, 7.00000}
	arr2 = []float64{9.00000, 7.00000, 6.00000}
	expected = []float64{1.1111111111111112, 1.1428571428571428, 1.1666666666666667}

	result, err = DivideArrays(-1, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Mismatched array lengths
	arr3 := []float64{1.0}
	_, err = DivideArrays(2, arr1, arr3)
	if err == nil {
		t.Error("Expected an error for mismatched array lengths, got none")
	}

	// Test case 4: Fewer than two arrays
	_, err = DivideArrays(2, arr1)
	if err == nil {
		t.Error("Expected an error for fewer than two arrays, got none")
	}
}

func TestPowerArrays(t *testing.T) {
	// Test case 1: Regular arrays, precision = 2
	arr1 := []float64{2.00000, 3.00000, 4.00000}
	arr2 := []float64{3.00000, 2.00000, 1.00000}
	expected := []float64{8.00, 9.00, 4.00}

	result, err := PowerArrays(2, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Adjusted Test Case 2: Precision = -1 (no rounding)
	arr1 = []float64{2.00000, 3.00000, 4.00000}
	arr2 = []float64{3.00000, 2.00000, 1.00000}
	expected = []float64{8.0, 9.0, 4.0} // Updated for clarity

	result, err = PowerArrays(-1, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Mismatched array lengths
	arr3 := []float64{1.0}
	_, err = PowerArrays(2, arr1, arr3)
	if err == nil {
		t.Error("Expected an error for mismatched array lengths, got none")
	}

	// Test case 4: Fewer than two arrays
	_, err = PowerArrays(2, arr1, nil)
	if err == nil {
		t.Error("Expected an error for fewer than two arrays, got none")
	}
}

func TestModuloArrays(t *testing.T) {
	// Test case 1: Regular arrays, precision = 2
	arr1 := []float64{10.00000, 8.00000, 7.00000}
	arr2 := []float64{9.00000, 7.00000, 6.00000}
	expected := []float64{1.00, 1.00, 1.00}

	result, err := ModuloArrays(2, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Adjusted Test Case 2: Precision = -1 (no rounding)
	arr1 = []float64{10.00000, 8.00000, 7.00000}
	arr2 = []float64{9.00000, 7.00000, 6.00000}
	expected = []float64{1.0, 1.0, 1.0} // Updated for clarity

	result, err = ModuloArrays(-1, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Mismatched array lengths
	arr3 := []float64{1.0}
	_, err = ModuloArrays(2, arr1, arr3)
	if err == nil {
		t.Error("Expected an error for mismatched array lengths, got none")
	}

	// Test case 4: Fewer than two arrays
	_, err = ModuloArrays(2, arr1, nil)
	if err == nil {
		t.Error("Expected an error for fewer than two arrays, got none")
	}
}

func TestLogArrays(t *testing.T) {
	// Test case 1: Regular arrays, precision = 2
	arr1 := []float64{2.00000, 10.00000, 5.00000}
	arr2 := []float64{8.00000, 100.00000, 25.00000}
	expected := []float64{3.00, 2.00, 2.00}

	result, err := LogArrays(2, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Adjusted Test Case 2: Precision = -1 (no rounding)
	arr1 = []float64{2.00000, 10.00000, 5.00000}
	arr2 = []float64{8.00000, 100.00000, 25.00000}
	expected = []float64{3.0, 2.0, 2.0} // Updated for clarity

	result, err = LogArrays(-1, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Mismatched array lengths
	arr3 := []float64{1.0}
	_, err = LogArrays(2, arr1, arr3)
	if err == nil {
		t.Error("Expected an error for mismatched array lengths, got none")
	}

	// Test case 4: Fewer than two arrays
	_, err = LogArrays(2, arr1, nil)
	if err == nil {
		t.Error("Expected an error for fewer than two arrays, got none")
	}
}

func TestSqrtArrays(t *testing.T) {
	// Test case 1: Regular array, precision = 2
	arr1 := []float64{4.00000, 9.00000, 16.00000}
	expected := []float64{2.00, 3.00, 4.00}

	result, err := SqrtArrays(2, arr1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Precision = -1 (no rounding)
	arr1 = []float64{4.00000, 9.00000, 16.00000}
	expected = []float64{2.0, 3.0, 4.0}

	result, err = SqrtArrays(-1, arr1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Negative values in the array
	arr2 := []float64{-1.0, 4.0, 9.0}
	_, err = SqrtArrays(2, arr2)
	if err == nil {
		t.Error("Expected an error for negative values, got none")
	}

	// Test case 4: Empty array
	arr3 := []float64{}
	_, err = SqrtArrays(2, arr3)
	if err == nil {
		t.Error("Expected an error for an empty array, got none")
	}
}

// compareSlices checks if two float64 slices are equal within a given tolerance.
func compareSlices(a, b []float64, tolerance float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if diff := a[i] - b[i]; diff < -tolerance || diff > tolerance {
			return false
		}
	}
	return true
}
