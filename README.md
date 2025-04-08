# LiteArray

LiteArray is a lightweight Go library for performing advanced mathematical operations on arrays and matrices. It provides a wide range of functions for element-wise operations, statistical calculations, and matrix manipulations, with support for optional rounding to a specified precision. 

## Why LiteArray?

While Go has some built-in capabilities for basic array operations, it lacks the comprehensive mathematical array functionality found in languages like Python (NumPy) or MATLAB. LiteArray aims to bridge this gap by providing:

- **Simplicity**: A straightforward API that's easy to learn and use
- **Precision Control**: Unlike many libraries, LiteArray gives you explicit control over rounding precision
- **Comprehensive Error Handling**: Clear, actionable error messages that help developers quickly identify issues
- **Lightweight**: Minimal dependencies and efficient implementation
- **Go-Idiomatic**: Follows Go conventions and patterns, making it feel natural to Go developers

LiteArray was created to make array and matrix operations in Go as accessible and powerful as they are in data science-focused languages, while maintaining Go's performance advantages.

## Features

- **Array Operations**: Add, subtract, multiply, divide, and perform other element-wise operations on arrays.
- **Statistical Functions**: Calculate mean, median, mode, variance, standard deviation, and percentiles.
- **Matrix Operations**: Transpose matrices, compute determinants, invert matrices, and calculate eigenvalues.
- **Precision Control**: Specify rounding precision for floating-point results.
- **Error Handling**: Comprehensive error handling for invalid inputs, mismatched array lengths, and edge cases.

## Installation

To use LiteArray in your project, add it to your `go.mod` file:

```bash
go get github.com/Av26qcDL/litearray
```

## Usage

### Array Operations

Perform element-wise addition of arrays with optional rounding:

```go
package main

import (
    "fmt"
    "github.com/Av26qcDL/litearray"
)

func main() {
    arr1 := []float64{1.12345, 2.98765}
    arr2 := []float64{4.87654, 5.01235}
    result, err := litearray.AddArrays(2, arr1, arr2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result) // Output: [6.0, 8.0]
}
```

### Statistical Functions

Calculate the mean of multiple arrays:

```go
arr1 := []float64{1.0, 2.0, 3.0}
arr2 := []float64{4.0, 5.0, 6.0}
mean, _ := litearray.MeanArrays(2, arr1, arr2)
// mean => []float64{2.50, 3.50, 4.50}
```

### Matrix Operations

Transpose a matrix:

```go
matrix := [][]float64{
    {1.0, 2.0, 3.0},
    {4.0, 5.0, 6.0},
}
transposed, _ := litearray.TransposeMatrix(2, matrix)
// transposed => [][]float64{{1.0, 4.0}, {2.0, 5.0}, {3.0, 6.0}}
```

## Rounding Behavior

LiteArray handles floating-point operations with a small tolerance (default: 0.0001). Be aware that results may be approximated due to rounding when combining arrays or performing mathematical operations.

Example:

```go
arr1 := []float64{1.12345, 2.98765}
arr2 := []float64{4.87654, 5.01235}
result, _ := litearray.AddArrays(2, arr1, arr2)
// result => []float64{6.0, 8.0}
```

## Error Handling

LiteArray provides detailed error messages for invalid inputs, such as:

- Mismatched array lengths
- Division by zero
- Invalid precision values (must be between -1 and 10)
- Empty or nil arrays/matrices

## Dependencies

LiteArray uses the [Gonum](https://gonum.org/) library for advanced matrix operations.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.