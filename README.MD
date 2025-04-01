This is a library for lightweight array operations 

## Rounding Behavior
This library handles floating-point operations with a small tolerance (default: 0.0001). Be aware that results may be approximated due to rounding when combining arrays or performing math operations. 

Example:
```go
arr1 := []float64{1.12345, 2.98765}
arr2 := []float64{4.87654, 5.01235}
result, _ := AddArrays(2, arr1, arr2) 
// result => []float64{6.0, 8.0}