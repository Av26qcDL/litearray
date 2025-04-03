package litearray

import (
	"fmt"
	"math"
)

// AddArrays adds multiple arrays element-wise and supports optional rounding to a specified precision.

func AddArrays(precision int, arrays ...[]float64) ([]float64, error) {
	// Ensure at least two arrays are provided
	if len(arrays) < 2 {
		return nil, fmt.Errorf("at least two arrays are required to perform addition")
	}

	// Check that all arrays are the same length
	length := len(arrays[0])
	for _, array := range arrays {
		if len(array) != length {
			return nil, fmt.Errorf("all arrays must be of the same length")
		}
	}

	// Create a result slice initialized to zero
	result := make([]float64, length)

	// Loop through the arrays and perform the addition
	for _, array := range arrays {
		for i := range array {
			result[i] += array[i]
		}
	}

	// Apply rounding if precision is non-negative
	if precision >= 0 {
		for i := range result {
			factor := math.Pow(10, float64(precision)) // e.g., 10^2 for two decimal places
			fmt.Printf("Before Rounding: result[%d] = %f\n", i, result[i])
			result[i] = math.Round(result[i]*factor) / factor
			fmt.Printf("After Rounding: result[%d] = %f\n", i, result[i])
		}
	}

	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	return result, nil
}

// SubtractArrays subtracts multiple arrays element-wise and supports optional rounding to a specified precision.

func SubtractArrays(precision int, arrays ...[]float64) ([]float64, error) {
	// Ensure at least two arrays are provided
	if len(arrays) < 2 {
		return nil, fmt.Errorf("at least two arrays are required to perform subtraction")
	}

	// Check that all arrays are the same length
	length := len(arrays[0])
	for _, array := range arrays {
		if len(array) != length {
			return nil, fmt.Errorf("all arrays must be of the same length")
		}
	}

	// Create a result slice initialized to zero
	// This assumes that the first array is not nil and has the same length as the others
	result := make([]float64, length)
	// Initialize result to the first array
	copy(result, arrays[0])
	if arrays[0] == nil {
		return nil, fmt.Errorf("the first array cannot be nil")
	}
	if length == 0 {
		return nil, fmt.Errorf("the first array cannot be empty")
	}

	// Loop through the remaining arrays and perform the subtraction
	for _, array := range arrays[1:] { // Start from the second array
		for i := range array {
			result[i] -= array[i]
		}
	}

	// Apply rounding if precision is non-negative
	if precision >= 0 {
		for i := range result {
			factor := math.Pow(10, float64(precision)) // e.g., 10^2 for two decimal places
			fmt.Printf("Before Rounding: result[%d] = %f\n", i, result[i])
			result[i] = math.Round(result[i]*factor) / factor
			fmt.Printf("After Rounding: result[%d] = %f\n", i, result[i])
		}
	}

	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	return result, nil
}

// MultiplyArrays multiplies multiple arrays element-wise and supports optional rounding to a specified precision.

func MultiplyArrays(precision int, arrays ...[]float64) ([]float64, error) {
	// Ensure at least two arrays are provided
	if len(arrays) < 2 {
		return nil, fmt.Errorf("at least two arrays are required to perform subtraction")
	}

	// Check that all arrays are the same length
	length := len(arrays[0])
	for _, array := range arrays {
		if len(array) != length {
			return nil, fmt.Errorf("all arrays must be of the same length")
		}
	}

	// Create a result slice initialized to zero
	result := make([]float64, length)
	// Initialize result to the first array
	copy(result, arrays[0])
	if arrays[0] == nil {
		return nil, fmt.Errorf("the first array cannot be nil")
	}
	if length == 0 {
		return nil, fmt.Errorf("the first array cannot be empty")
	}

	// Loop through the arrays and perform the multiplication
	for _, array := range arrays[1:] { // Start from the second array
		for i := range array {
			result[i] *= array[i]
		}
	}

	// Apply rounding if precision is non-negative
	if precision >= 0 {
		for i := range result {
			factor := math.Pow(10, float64(precision)) // e.g., 10^2 for two decimal places
			fmt.Printf("Before Rounding: result[%d] = %f\n", i, result[i])
			result[i] = math.Round(result[i]*factor) / factor
			fmt.Printf("After Rounding: result[%d] = %f\n", i, result[i])
		}
	}

	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	return result, nil
}

// DivideArrays divides multiple arrays element-wise and supports optional rounding to a specified precision.

func DivideArrays(precision int, arrays ...[]float64) ([]float64, error) {
	// Ensure at least two arrays are provided
	if len(arrays) < 2 {
		return nil, fmt.Errorf("at least two arrays are required to perform division")
	}

	// Check that all arrays are the same length
	length := len(arrays[0])
	for _, array := range arrays {
		if len(array) != length {
			return nil, fmt.Errorf("all arrays must be of the same length")
		}
	}

	// Create a result slice initialized to zero
	result := make([]float64, length)
	// Initialize result to the first array
	copy(result, arrays[0])
	if arrays[0] == nil {
		return nil, fmt.Errorf("the first array cannot be nil")
	}
	if length == 0 {
		return nil, fmt.Errorf("the first array cannot be empty")
	}

	// Loop through the remaining arrays and perform the division
	for _, array := range arrays[1:] { // Start from the second array
		for i := range array {
			if array[i] == 0 {
				return nil, fmt.Errorf("division by zero at index %d", i)
			}
			result[i] /= array[i]
		}
	}

	// Apply rounding if precision is non-negative
	if precision >= 0 {
		for i := range result {
			factor := math.Pow(10, float64(precision)) // e.g., 10^2 for two decimal places
			fmt.Printf("Before Rounding: result[%d] = %f\n", i, result[i])
			result[i] = math.Round(result[i]*factor) / factor
			fmt.Printf("After Rounding: result[%d] = %f\n", i, result[i])
		}
	}

	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	return result, nil
}

// PowerArrays raises each element of the base array to the corresponding element of the exponent array and supports optional rounding to a specified precision.

func PowerArrays(precision int, base []float64, exponent []float64) ([]float64, error) {
	// Ensure at least two arrays are provided
	if len(base) == 0 || len(exponent) == 0 {
		return nil, fmt.Errorf("both base and exponent arrays must be provided")
	}

	// Check that all arrays are the same length
	if len(base) != len(exponent) {
		return nil, fmt.Errorf("base and exponent arrays must be of the same length")
	}

	// Create a result slice initialized to zero
	result := make([]float64, len(base))

	// Loop through the arrays and perform the power operation
	for i := range base {
		result[i] = math.Pow(base[i], exponent[i])
	}

	// Apply rounding if precision is non-negative
	if precision >= 0 {
		for i := range result {
			factor := math.Pow(10, float64(precision)) // e.g., 10^2 for two decimal places
			fmt.Printf("Before Rounding: result[%d] = %f\n", i, result[i])
			result[i] = math.Round(result[i]*factor) / factor
			fmt.Printf("After Rounding: result[%d] = %f\n", i, result[i])
		}
	}

	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	return result, nil
}

// ModuloArrays calculates the modulo of two arrays element-wise and supports optional rounding to a specified precision.
func ModuloArrays(precision int, dividend []float64, divisor []float64) ([]float64, error) {
	// Ensure at least two arrays are provided
	if len(dividend) == 0 || len(divisor) == 0 {
		return nil, fmt.Errorf("both dividend and divisor arrays must be provided")
	}

	// Check that all arrays are the same length
	if len(dividend) != len(divisor) {
		return nil, fmt.Errorf("dividend and divisor arrays must be of the same length")
	}

	// Create a result slice initialized to zero
	result := make([]float64, len(dividend))

	// Loop through the arrays and perform the modulo operation
	for i := range dividend {
		if divisor[i] == 0 {
			return nil, fmt.Errorf("division by zero at index %d", i)
		}
		result[i] = math.Mod(dividend[i], divisor[i])
	}

	// Apply rounding if precision is non-negative
	if precision >= 0 {
		for i := range result {
			factor := math.Pow(10, float64(precision)) // e.g., 10^2 for two decimal places
			fmt.Printf("Before Rounding: result[%d] = %f\n", i, result[i])
			result[i] = math.Round(result[i]*factor) / factor
			fmt.Printf("After Rounding: result[%d] = %f\n", i, result[i])
		}
	}

	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	return result, nil
}

func LogArrays(precision int, base []float64, dividend []float64) ([]float64, error) {
	// Ensure at least two arrays are provided
	if len(base) == 0 || len(dividend) == 0 {
		return nil, fmt.Errorf("both base and dividend arrays must be provided")
	}

	// Check that all arrays are the same length
	if len(base) != len(dividend) {
		return nil, fmt.Errorf("base and dividend arrays must be of the same length")
	}

	// Create a result slice initialized to zero
	result := make([]float64, len(base))

	// Loop through the arrays and perform the logarithm operation
	for i := range base {
		if base[i] <= 0 || dividend[i] <= 0 {
			return nil, fmt.Errorf("logarithm undefined for non-positive values at index %d", i)
		}
		result[i] = math.Log(dividend[i]) / math.Log(base[i])
	}

	// Apply rounding if precision is non-negative
	if precision >= 0 {
		for i := range result {
			factor := math.Pow(10, float64(precision)) // e.g., 10^2 for two decimal places
			fmt.Printf("Before Rounding: result[%d] = %f\n", i, result[i])
			result[i] = math.Round(result[i]*factor) / factor
			fmt.Printf("After Rounding: result[%d] = %f\n", i, result[i])
		}
	}

	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	return result, nil
}

func SqrtArrays(precision int, arrays ...[]float64) ([]float64, error) {
	// Ensure at least one array is provided
	if len(arrays) == 0 {
		return nil, fmt.Errorf("at least one array is required to perform square root")
	}

	// Check that all arrays are the same length
	length := len(arrays[0])
	if length == 0 {
		return nil, fmt.Errorf("array cannot be empty")
	}
	for _, array := range arrays {
		if len(array) != length {
			return nil, fmt.Errorf("all arrays must be of the same length")
		}
	}

	// Create a result slice initialized to zero
	result := make([]float64, length)

	// Loop through the arrays and perform the square root operation
	for _, array := range arrays {
		for i := range array {
			if array[i] < 0 {
				return nil, fmt.Errorf("square root undefined for negative values at index %d", i)
			}
			result[i] = math.Sqrt(array[i])
		}
	}

	// Apply rounding if precision is non-negative
	if precision >= 0 {
		for i := range result {
			factor := math.Pow(10, float64(precision)) // e.g., 10^2 for two decimal places
			result[i] = math.Round(result[i]*factor) / factor
		}
	}

	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	return result, nil
}

func AbsArrays(precision int, arrays ...[]float64) ([]float64, error) {
	// Ensure at least two arrays are provided
	if len(arrays) < 2 {
		return nil, fmt.Errorf("at least two arrays are required to perform addition")
	}

	// Check that all arrays are the same length
	length := len(arrays[0])
	if length == 0 {
		return nil, fmt.Errorf("array cannot be empty")
	}
	for _, array := range arrays {
		if len(array) != length {
			return nil, fmt.Errorf("all arrays must be of the same length")
		}
	}

	// Create a result slice initialized to zero
	result := make([]float64, length)

	// Loop through the arrays and sum the elements element-wise
	for _, array := range arrays {
		for i := range array {
			result[i] += array[i]
		}
	}

	// Take the absolute value of the summed result
	for i := range result {
		result[i] = math.Abs(result[i])
	}

	// Apply rounding if precision is non-negative
	if precision >= 0 {
		for i := range result {
			factor := math.Pow(10, float64(precision)) // e.g., 10^2 for two decimal places
			result[i] = math.Round(result[i]*factor) / factor
		}
	}

	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	return result, nil
}

func MeanArrays(precision int, arrays ...[]float64) ([]float64, error) {
	// Ensure at least two arrays are provided
	if len(arrays) < 2 {
		return nil, fmt.Errorf("at least two arrays are required to perform addition")
	}

	// Check that all arrays are the same length
	length := len(arrays[0])
	if length == 0 {
		return nil, fmt.Errorf("array cannot be empty")
	}
	for _, array := range arrays {
		if len(array) != length {
			return nil, fmt.Errorf("all arrays must be of the same length")
		}
	}

	// Create a result slice initialized to zero
	result := make([]float64, length)

	// Loop through the arrays and perform the mean operation
	for _, array := range arrays {
		for i := range array {
			result[i] += array[i]
		}
	}

	for i := range result {
		result[i] /= float64(len(arrays))
	}

	// Apply rounding if precision is non-negative
	if precision >= 0 {
		for i := range result {
			factor := math.Pow(10, float64(precision)) // e.g., 10^2 for two decimal places
			result[i] = math.Round(result[i]*factor) / factor
		}
	}

	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	return result, nil
}
