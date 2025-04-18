package litearray

import (
	"math/cmplx"
	"testing"
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
	arr1 := []float64{4.00000, 8.00000, 56.00000}
	arr2 := []float64{12.00000, 1.00000, 8.00000}
	expected := []float64{4.00, 3.00, 8.00}

	result, err := SqrtArrays(2, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Precision = -1 (no rounding)
	arr1 = []float64{4.00000, 8.00000, 56.00000}
	arr2 = []float64{12.00000, 1.00000, 8.00000}
	expected = []float64{4.00, 3.00, 8.00}

	result, err = SqrtArrays(-1, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Negative values in the array
	arr3 := []float64{-1.0, 4.0, 9.0}
	_, err = SqrtArrays(2, arr3)
	if err == nil {
		t.Error("Expected an error for negative values, got none")
	}

	// Test case 4: Empty array
	arr4 := []float64{}
	_, err = SqrtArrays(2, arr4)
	if err == nil {
		t.Error("Expected an error for an empty array, got none")
	}
}

func TestAbsArrays(t *testing.T) {
	// Test case 1: Regular array, precision = 2
	arr1 := []float64{-1.00000, -2.00000, 3.00000}
	arr2 := []float64{-2.00000, -3.00000, 3.00000}
	expected := []float64{3.00, 5.00, 6.00}

	result, err := AbsArrays(2, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Precision = -1 (no rounding)
	arr1 = []float64{-1.00000, -2.00000, 3.00000}
	arr2 = []float64{-2.00000, -3.00000, 3.00000}
	expected = []float64{3.00, 5.00, 6.00}

	result, err = AbsArrays(-1, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Empty array
	arr3 := []float64{}
	_, err = AbsArrays(2, arr3)
	if err == nil {
		t.Error("Expected an error for an empty array, got none")
	}
}

func TestMeanArrays(t *testing.T) {
	// Test case 1: Regular array, precision = 2
	arr1 := []float64{1.00000, 2.00000, 3.00000}
	arr2 := []float64{4.00000, 5.00000, 6.00000}
	expected := []float64{2.50, 3.50, 4.50}

	result, err := MeanArrays(2, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Precision = -1 (no rounding)
	arr1 = []float64{1.00000, 2.00000, 3.00000}
	arr2 = []float64{4.00000, 5.00000, 6.00000}
	expected = []float64{2.50, 3.50, 4.50}

	result, err = MeanArrays(-1, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Empty array
	arr3 := []float64{}
	_, err = MeanArrays(2, arr3)
	if err == nil {
		t.Error("Expected an error for an empty array, got none")
	}

	// Test case 4: Mismatched array lengths
	arr4 := []float64{1.0}
	_, err = MeanArrays(2, arr1, arr4)
	if err == nil {
		t.Error("Expected an error for mismatched array lengths, got none")
	}
}

func TestMedianArrays(t *testing.T) {
	// Test case 1: Regular array, precision = 2
	arr1 := []float64{1.00000, 2.00000, 3.00000}
	arr2 := []float64{4.00000, 5.00000, 6.00000}
	expected := []float64{3.50}

	result, err := MedianArrays(2, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Precision = -1 (no rounding)
	arr1 = []float64{1.00000, 2.00000, 3.00000}
	arr2 = []float64{4.00000, 5.00000, 6.00000}
	expected = []float64{3.50}

	result, err = MedianArrays(-1, arr1, arr2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Empty array
	arr3 := []float64{}
	_, err = MedianArrays(2, arr3)
	if err == nil {
		t.Error("Expected an error for an empty array, got none")
	}

	// Test case 4: Mismatched array lengths
	arr4 := []float64{1.0}
	_, err = MedianArrays(2, arr1, arr4)
	if err == nil {
		t.Error("Expected an error for mismatched array lengths, got none")
	}
}

func TestModeMultipleArrays(t *testing.T) {
	// Test case 1: Regular array, precision = 2
	arr1 := [][]float64{{1.00000, 2.00000, 3.00000}, {4.00000, 2.00000, 2.00000}, {8.00000, 2.00000, 2.00000}}
	expected := []float64{2.00}

	// Correctly handle the two return values
	result, err := ModeMultipleArrays(2, arr1...)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Precision = -1 (no rounding)
	arr1 = [][]float64{{1.00000, 2.00000, 3.00000}, {4.00000, 2.00000, 2.00000}, {8.00000, 2.00000, 2.00000}}
	expected = []float64{2.00}

	result, err = ModeMultipleArrays(-1, arr1...)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Empty array
	arr3 := []float64{}
	_, err = ModeMultipleArrays(2, arr3)
	if err == nil {
		t.Error("Expected an error for an empty array, got none")
	}

	// Test case 4: Mismatched array lengths
	arr4 := []float64{1.0}
	_, err = LogArrays(2, arr3, arr4)
	if err == nil {
		t.Error("Expected an error for mismatched array lengths, got none")
	}
}

func TestVarianceArrays(t *testing.T) {
	// Test case 1: Regular array, precision = 2
	arrays := [][]float64{{1.00000, 2.00000, 3.00000}, {4.00000, 5.00000, 6.00000}, {7.00000, 8.00000, 9.00000}}
	expected := []float64{6.00, 6.00, 6.00}

	result, err := VarianceArrays(2, arrays...)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Precision = -1 (no rounding)
	arrays = [][]float64{{1.00000, 2.00000, 3.00000}, {4.00000, 5.00000, 6.00000}, {7.00000, 8.00000, 9.00000}}
	expected = []float64{6.00, 6.00, 6.00} // Updated for clarity

	result, err = VarianceArrays(-1, arrays...)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Empty array
	arr3 := []float64{}
	_, err = VarianceArrays(2, arr3)
	if err == nil {
		t.Error("Expected an error for an empty array, got none")
	}

	// Test case 4: Mismatched array lengths
	arr4 := []float64{1.0}
	_, err = VarianceArrays(2, arrays[0], arr4)
	if err == nil {
		t.Error("Expected an error for mismatched array lengths, got none")
	}
}

func TestStandardDeviationArrays(t *testing.T) {
	// Test case 1: Regular array, precision = 2
	arrays := [][]float64{{1.00000, 2.00000, 3.00000}, {4.00000, 5.00000, 6.00000}, {7.00000, 8.00000, 9.00000}}
	expected := []float64{2.45, 2.45, 2.45}

	result, err := StandardDeviationArrays(2, arrays...)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Precision = -1 (no rounding)
	arrays = [][]float64{{1.00000, 2.00000, 3.00000}, {4.00000, 5.00000, 6.00000}, {7.00000, 8.00000, 9.00000}}
	expected = []float64{2.449489742783178, 2.449489742783178, 2.449489742783178} // Updated for clarity

	result, err = StandardDeviationArrays(-1, arrays...)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Empty array
	arr3 := []float64{}
	_, err = StandardDeviationArrays(2, arr3)
	if err == nil {
		t.Error("Expected an error for an empty array, got none")
	}

	// Test case 4: Mismatched array lengths
	arr4 := []float64{1.0}
	_, err = StandardDeviationArrays(2, arrays[0], arr4)
	if err == nil {
		t.Error("Expected an error for mismatched array lengths, got none")
	}
}

func TestMinArrays(t *testing.T) {
	// Test case 1: Regular array, precision = 2
	arrays := [][]float64{{1.00000, 2.00000, 3.00000}, {4.00000, 5.00000, 6.00000}, {7.00000, 8.00000, 9.00000}}
	expected := []float64{1.00, 2.00, 3.00}

	result, err := MinArrays(2, arrays...)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Precision = -1 (no rounding)
	arrays = [][]float64{{1.00000, 2.00000, 3.00000}, {4.00000, 5.00000, 6.00000}, {7.00000, 8.00000, 9.00000}}
	expected = []float64{1.0, 2.0, 3.0} // Updated for clarity

	result, err = MinArrays(-1, arrays...)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Empty array
	arr3 := []float64{}
	_, err = MinArrays(2, arr3)
	if err == nil {
		t.Error("Expected an error for an empty array, got none")
	}

	// Test case 4: Mismatched array lengths
	arr4 := []float64{1.0}
	_, err = MinArrays(2, arrays[0], arr4)
	if err == nil {
		t.Error("Expected an error for mismatched array lengths, got none")
	}
}
func TestMaxArrays(t *testing.T) {
	// Test case 1: Regular array, precision = 2
	arrays := [][]float64{{1.00000, 2.00000, 3.00000}, {4.00000, 5.00000, 6.00000}, {7.00000, 8.00000, 9.00000}}
	expected := []float64{7.00, 8.00, 9.00}

	result, err := MaxArrays(2, arrays...)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Precision = -1 (no rounding)
	arrays = [][]float64{{1.00000, 2.00000, 3.00000}, {4.00000, 5.00000, 6.00000}, {7.00000, 8.00000, 9.00000}}
	expected = []float64{7.0, 8.0, 9.0} // Updated for clarity

	result, err = MaxArrays(-1, arrays...)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Empty array
	arr3 := []float64{}
	_, err = MaxArrays(2, arr3)
	if err == nil {
		t.Error("Expected an error for an empty array, got none")
	}

	// Test case 4: Mismatched array lengths
	arr4 := []float64{1.0}
	_, err = MaxArrays(2, arrays[0], arr4)
	if err == nil {
		t.Error("Expected an error for mismatched array lengths, got none")
	}
}

func TestRangeArrays(t *testing.T) {
	// Test case 1: Regular array, precision = 2
	arrays := [][]float64{{1.50000, 2.30000, 3.70000}, {0.90000, 2.80000, 3.10000}, {1.20000, 1.90000, 4.00000}}
	expected := []float64{0.60, 0.90, 0.90}

	result, err := RangeArrays(2, arrays...)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Precision = -1 (no rounding)
	arrays = [][]float64{{1.50000, 2.30000, 3.70000}, {0.90000, 2.80000, 3.10000}, {1.20000, 1.90000, 4.00000}}
	expected = []float64{0.60, 0.90, 0.90} // Updated for clarity

	result, err = RangeArrays(-1, arrays...)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Empty array
	arr3 := []float64{}
	_, err = RangeArrays(2, arr3)
	if err == nil {
		t.Error("Expected an error for an empty array, got none")
	}

	// Test case 4: Mismatched array lengths
	arr4 := []float64{1.0}
	_, err = RangeArrays(2, arrays[0], arr4)
	if err == nil {
		t.Error("Expected an error for mismatched array lengths, got none")
	}
}

func TestPercentileArrays(t *testing.T) {
	// Test case 1: Regular array, precision = 2
	arrays := [][]float64{{1.00000, 2.00000, 3.00000}, {4.00000, 5.00000, 6.00000}, {7.00000, 8.00000, 9.00000}}
	expected := []float64{2.00, 5.00, 8.00}

	result, err := PercentileArrays(2, 50, arrays...)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices with a standard tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Precision = -1 (no rounding)
	arrays = [][]float64{{1.00000, 2.00000, 3.00000}, {4.00000, 5.00000, 6.00000}, {7.00000, 8.00000, 9.00000}}
	expected = []float64{2.00, 5.00, 8.00} // Updated for clarity

	result, err = PercentileArrays(-1, 50, arrays...)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Use compareSlices again with tolerance
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Empty array
	arr3 := []float64{}
	_, err = PercentileArrays(2, 50, arr3)
	if err == nil {
		t.Error("Expected an error for an empty array, got none")
	}

	// Test case 4: Mismatched array lengths
	arr4 := []float64{1.0}
	arr5 := []float64{1.0, 2.0}
	_, err = PercentileArrays(2, 50, arr4, arr5)
	if err == nil {
		t.Error("Expected an error for mismatched array lengths, got none")
	}
}

func TestTransposeMatrix(t *testing.T) {
	// Test case 1: Regular matrix
	matrix := [][]float64{{1.10000, 2.20000, 3.30000}, {4.40000, 5.50000, 6.60000}}
	expected := [][]float64{{1.10000, 4.40000}, {2.20000, 5.50000}, {3.30000, 6.60000}}

	result, err := TransposeMatrix(2, matrix)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !compareSlices(result[0], expected[0], 0.0001) || !compareSlices(result[1], expected[1], 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Empty matrix
	matrix = [][]float64{}
	_, err = TransposeMatrix(2, matrix)
	if err == nil || err.Error() != "matrix cannot be empty" {
		t.Errorf("Expected error 'matrix cannot be empty', got %v", err)
	}
}

func TestTransposeMatrix_InvalidInputs(t *testing.T) {
	// Test case 1: Nil matrix
	var matrix [][]float64
	_, err := TransposeMatrix(2, matrix)
	if err == nil {
		t.Error("Expected an error for nil matrix, got none")
	}

	// Test case 2: Matrix with mismatched row lengths
	matrix = [][]float64{{1.0, 2.0}, {3.0}}
	_, err = TransposeMatrix(2, matrix)
	if err == nil {
		t.Error("Expected an error for mismatched row lengths, got none")
	}

	// Test case 3: Precision out of range
	matrix = [][]float64{{1.0, 2.0}, {3.0, 4.0}}
	_, err = TransposeMatrix(11, matrix) // Precision > 10
	if err == nil {
		t.Error("Expected an error for precision out of range, got none")
	}
}

func TestDeterminantMatrix(t *testing.T) {
	// Test case 1: Regular matrix
	matrix := [][]float64{{2, 3, 1}, {4, 5, 6}, {7, 8, 9}}
	expected := 9.0

	result, err := DeterminantMatrix(matrix)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Empty matrix
	matrix = [][]float64{}
	_, err = DeterminantMatrix(matrix)
	if err == nil || err.Error() != "matrix must be square and non-empty" {
		t.Errorf("Expected error 'matrix must be square and non-empty', got %v", err)
	}
}

func TestInversionMatrix(t *testing.T) {
	// Test case 1: Regular matrix
	matrix := [][]float64{{2, 3}, {4, 5}}
	expected := [][]float64{{-2.5, 1.5}, {2, -1}}

	result, err := InversionMatrix(matrix)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !compareSlices(result[0], expected[0], 0.0001) || !compareSlices(result[1], expected[1], 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Empty matrix
	matrix = [][]float64{}
	_, err = InversionMatrix(matrix)
	if err == nil || err.Error() != "matrix must be square and non-empty" {
		t.Errorf("Expected error 'matrix must be square and non-empty', got %v", err)
	}
}

func TestEigenvalues2x2(t *testing.T) {
	// Test case 1: Regular 2x2 matrix
	matrix := [][]float64{{4, 2}, {1, 3}}
	expected := []float64{5.0, 2.0}

	result, err := Eigenvalues2x2(matrix)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !compareSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Empty matrix
	matrix = [][]float64{}
	_, err = Eigenvalues2x2(matrix)
	if err == nil || err.Error() != "matrix must be 2x2" {
		t.Errorf("Expected error 'matrix must be 2x2', got %v", err)
	}
}

func TestEigenvalues3x3AndHigher(t *testing.T) {
	// Test case 1: Regular 3x3 matrix
	matrix := [][]float64{{4, 2, 1}, {1, 3, 2}, {2, 1, 5}}
	expected := []complex128{
		complex(7.14092334, 0),
		complex(2.42953833, 0.63170371),
		complex(2.42953833, -0.63170371),
	}

	result, err := Eigenvalues3x3AndHigher(matrix)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !compareComplexSlices(result, expected, 0.0001) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Empty matrix
	matrix = [][]float64{}
	_, err = Eigenvalues3x3AndHigher(matrix)
	if err == nil || err.Error() != "matrix must be square and non-empty" {
		t.Errorf("Expected error 'matrix must be square and non-empty', got %v", err)
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

// compareComplexSlices compares two slices of complex128 for equality within a tolerance.
func compareComplexSlices(a, b []complex128, tolerance float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if cmplx.Abs(a[i]-b[i]) > tolerance { // Use cmplx.Abs to compare complex numbers
			return false
		}
	}
	return true
}
