package litearray

import (
	"fmt"
	"math"
	"sort"
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

// LogArrays calculates the logarithm of each element in the dividend array with respect to the base array and supports optional rounding to a specified precision.
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

// SqrtArrays calculates the square root of the sum of multiple arrays element-wise and supports optional rounding to a specified precision.
func SqrtArrays(precision int, arrays ...[]float64) ([]float64, error) {
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

	// Check for negative values in the result
	for i, value := range result {
		if value < 0 {
			return nil, fmt.Errorf("cannot calculate square root of a negative value at index %d: %f", i, value)
		}
	}

	// Take the square root of the summed result
	for i := range result {
		result[i] = math.Sqrt(result[i])
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

// AbsArrays calculates the absolute value of the sum of multiple arrays element-wise and supports optional rounding to a specified precision.
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

// MeanArrays calculates the mean of multiple arrays element-wise and supports optional rounding to a specified precision.
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

// MedianArrays calculates the median of multiple arrays element-wise and supports optional rounding to a specified precision.
func MedianArrays(precision int, arrays ...[]float64) ([]float64, error) {
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

	// Divide each element by the number of arrays to calculate the median
	for i := range result {
		result[i] /= float64(len(arrays))
	}

	// Sort the result slice to find the median
	sort.Float64s(result)
	mid := len(result) / 2

	// Calculate the median
	var median float64
	if len(result)%2 == 0 {
		median = (result[mid-1] + result[mid]) / 2
	} else {
		median = result[mid]
	}

	// Apply rounding if precision is non-negative
	if precision >= 0 {
		factor := math.Pow(10, float64(precision))
		median = math.Round(median*factor) / factor
	}

	// Return the median as a single-element slice
	return []float64{median}, nil
}

// ModeArray calculates the mode(s) of a slice of integers.
func ModeArray(nums []float64) []float64 {
	counts := make(map[float64]float64)
	maxCount := 0.00000
	for _, num := range nums {
		counts[num]++
		if counts[num] > maxCount {
			maxCount = counts[num]
		}
	}

	var modes []float64
	for num, count := range counts {
		if count == maxCount {
			modes = append(modes, num)
		}
	}
	return modes
}

// ModeMultipleArrays calculates the mode(s) across multiple arrays of integers.
func ModeMultipleArrays(precision int, arrays ...[]float64) ([]float64, error) {
	// Check if any arrays are provided
	if len(arrays) == 0 {
		return nil, fmt.Errorf("no arrays provided")
	}

	// Check if all arrays are empty
	isAllEmpty := true
	for _, arr := range arrays {
		if len(arr) > 0 {
			isAllEmpty = false
			break
		}
	}
	if isAllEmpty {
		return nil, fmt.Errorf("all arrays are empty")
	}

	// Combine all arrays into a single slice
	combined := []float64{}
	for _, arr := range arrays {
		combined = append(combined, arr...)
	}

	// Use a map to count occurrences of each number
	counts := make(map[float64]int)
	maxCount := 0
	for _, num := range combined {
		counts[num]++
		if counts[num] > maxCount {
			maxCount = counts[num]
		}
	}

	// Find all numbers with the maximum count
	var modes []float64
	for num, count := range counts {
		if count == maxCount {
			modes = append(modes, num)
		}
	}

	// Apply rounding to the modes if precision is non-negative
	if precision >= 0 {
		factor := math.Pow(10, float64(precision))
		for i := range modes {
			modes[i] = math.Round(modes[i]*factor) / factor
		}
	}

	// Validate precision range
	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	return modes, nil
}

// VarianceArrays calculates the variance of multiple arrays element-wise and supports optional rounding to a specified precision.
func VarianceArrays(precision int, arrays ...[]float64) ([]float64, error) {
	// Ensure at least two arrays are provided
	if len(arrays) < 2 {
		return nil, fmt.Errorf("at least two arrays are required")
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

	// Calculate the mean for each element
	for i := range result {
		result[i] /= float64(len(arrays))
	}

	// Calculate the variance for each element
	variance := make([]float64, length)
	for _, array := range arrays {
		for i := range array {
			diff := array[i] - result[i]
			variance[i] += diff * diff
		}
	}

	for i := range variance {
		variance[i] /= float64(len(arrays))
	}

	if precision >= 0 {
		for i := range variance {
			factor := math.Pow(10, float64(precision)) // e.g., 10^2 for two decimal places
			variance[i] = math.Round(variance[i]*factor) / factor
		}
	}

	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	return variance, nil
}

// StandardDeviationArrays calculates the standard deviation of multiple arrays element-wise and supports optional rounding to a specified precision.
func StandardDeviationArrays(precision int, arrays ...[]float64) ([]float64, error) {
	// Calculate the variance using the VarianceArrays function
	variance, err := VarianceArrays(precision, arrays...)
	if err != nil {
		return nil, err
	}

	// Take the square root of each variance value to calculate the standard deviation
	stdDev := make([]float64, len(variance))
	for i, v := range variance {
		stdDev[i] = math.Sqrt(v)
	}

	// Apply rounding if precision is non-negative
	if precision >= 0 {
		for i := range stdDev {
			factor := math.Pow(10, float64(precision)) // e.g., 10^2 for two decimal places
			stdDev[i] = math.Round(stdDev[i]*factor) / factor
		}
	}

	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	return stdDev, nil
}

// MinArrays finds the minimum value in each element across multiple arrays.
func MinArrays(precision int, arrays ...[]float64) ([]float64, error) {
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

	// Find the minimum value in each element
	minValues := make([]float64, length)
	for i := range result {
		minValues[i] = math.Inf(1)
	}

	for _, array := range arrays {
		for i := range array {
			if array[i] < minValues[i] {
				minValues[i] = array[i]
			}
		}
	}

	if precision >= 0 {
		for i := range minValues {
			factor := math.Pow(10, float64(precision)) // e.g., 10^2 for two decimal places
			minValues[i] = math.Round(minValues[i]*factor) / factor
		}
	}

	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	return minValues, nil
}

// MaxArrays finds the maximum value in each element across multiple arrays.
func MaxArrays(precision int, arrays ...[]float64) ([]float64, error) {
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

	// Find the maximum value in each element
	maxValues := make([]float64, length)
	for i := range result {
		maxValues[i] = math.Inf(-1)
	}

	for _, array := range arrays {
		for i := range array {
			if array[i] > maxValues[i] {
				maxValues[i] = array[i]
			}
		}
	}

	if precision >= 0 {
		for i := range maxValues {
			factor := math.Pow(10, float64(precision)) // e.g., 10^2 for two decimal places
			maxValues[i] = math.Round(maxValues[i]*factor) / factor
		}
	}

	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	return maxValues, nil
}

// RangeArrays calculates the range (max - min) of multiple arrays element-wise and supports optional rounding to a specified precision.
func RangeArrays(precision int, arrays ...[]float64) ([]float64, error) {
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

	// Find the minimum and maximum values in each element
	minValues := make([]float64, length)
	maxValues := make([]float64, length)
	for i := range result {
		minValues[i] = math.Inf(1)
		maxValues[i] = math.Inf(-1)
	}

	for _, array := range arrays {
		for i := range array {
			if array[i] < minValues[i] {
				minValues[i] = array[i]
			}
			if array[i] > maxValues[i] {
				maxValues[i] = array[i]
			}
		}
	}
	// Calculate the range (max - min)
	rangeValues := make([]float64, length)
	for i := range result {
		rangeValues[i] = maxValues[i] - minValues[i]
	}

	if precision >= 0 {
		for i := range rangeValues {
			factor := math.Pow(10, float64(precision)) // e.g., 10^2 for two decimal places
			rangeValues[i] = math.Round(rangeValues[i]*factor) / factor
		}
	}

	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	return rangeValues, nil
}

// PercentileArrays calculates the percentile of multiple arrays element-wise and supports optional rounding to a specified precision.
func PercentileArrays(precision int, percentile float64, arrays ...[]float64) ([]float64, error) {
	if percentile < 0 || percentile > 100 {
		return nil, fmt.Errorf("percentile must be between 0 and 100")
	}

	if precision > 10 {
		return nil, fmt.Errorf("precision too high; must be between -1 and 10")
	}

	// Check that all arrays are the same length
	if len(arrays) == 0 {
		return nil, fmt.Errorf("no arrays provided")
	}

	length := len(arrays[0])
	for _, array := range arrays {
		if len(array) != length {
			return nil, fmt.Errorf("all arrays must be of the same length")
		}
	}

	// Check that all arrays are non-empty
	for i, arr := range arrays {
		if len(arr) == 0 {
			return nil, fmt.Errorf("array at index %d cannot be empty", i)
		}
	}

	// Calculate the percentile for each array
	results := make([]float64, len(arrays))
	for i, arr := range arrays {
		sort.Float64s(arr)
		index := (percentile / 100) * float64(len(arr)-1)
		intIndex := int(index)
		fractionalPart := index - float64(intIndex)

		var result float64
		if intIndex+1 < len(arr) {
			result = arr[intIndex] + fractionalPart*(arr[intIndex+1]-arr[intIndex])
		} else {
			result = arr[intIndex]
		}

		// Apply rounding if precision is non-negative
		if precision >= 0 {
			factor := math.Pow(10, float64(precision))
			result = math.Round(result*factor) / factor
		}

		results[i] = result
	}

	return results, nil
}

// TransposeMatrix transposes a 2D matrix and supports optional rounding to a specified precision.
func TransposeMatrix(precision int, matrix [][]float64) ([][]float64, error) {
	// Validate precision
	if precision < -1 || precision > 10 {
		return nil, fmt.Errorf("precision must be between -1 and 10")
	}

	// Check if the matrix is empty
	if len(matrix) == 0 {
		return nil, fmt.Errorf("matrix cannot be empty")
	}

	// Check if the matrix is nil
	if matrix == nil {
		return nil, fmt.Errorf("matrix cannot be nil")
	}

	// Check if the matrix is a valid 2D slice
	width := len(matrix[0])
	for _, row := range matrix {
		if len(row) != width {
			return nil, fmt.Errorf("all rows in the matrix must have the same length")
		}
	}

	// Get the dimensions of the matrix
	length := len(matrix)

	// Create a new matrix with transposed dimensions
	result := make([][]float64, width)
	for i := range result {
		result[i] = make([]float64, length)
	}

	// Fill the transposed matrix
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			result[j][i] = matrix[i][j]
		}
	}

	// Apply rounding if precision is non-negative
	if precision >= 0 {
		for i := range result {
			for j := range result[i] {
				factor := math.Pow(10, float64(precision)) // e.g., 10^2 for two decimal places
				result[i][j] = math.Round(result[i][j]*factor) / factor
			}
		}
	}

	// Return the transposed matrix
	return result, nil
}

// DeterminantMatrix calculates the determinant of a square matrix using recursion.
func DeterminantMatrix(matrix [][]float64) (float64, error) {
	// Check if the matrix is square
	if len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return 0, fmt.Errorf("matrix must be square and non-empty")
	}

	// Base case for 2x2 matrix
	if len(matrix) == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0], nil
	}

	// Recursive case for larger matrices
	det := 0.0
	for i := range matrix {
		subMatrix := make([][]float64, len(matrix)-1)
		for j := range subMatrix {
			subMatrix[j] = make([]float64, len(matrix)-1)
			for k := range subMatrix[j] {
				if k < i {
					subMatrix[j][k] = matrix[j+1][k]
				} else {
					subMatrix[j][k] = matrix[j+1][k+1]
				}
			}
		}
		sign := 1.0
		if i%2 != 0 {
			sign = -1.0
		}
		subDet, err := DeterminantMatrix(subMatrix)
		if err != nil {
			return 0, err
		}
		det += sign * matrix[0][i] * subDet
	}

	return det, nil
}
