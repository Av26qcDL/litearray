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
