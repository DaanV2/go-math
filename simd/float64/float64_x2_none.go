//go:build simd_none || (!simd_avx512 && !simd_avx256)

package simdfloat64

import "math"

type Float64x2 struct {
	data [float64_x2_len]float64
}

func NewFloat64x2(data []float64) Float64x2 {
	var result Float64x2
	copy(result.data[:], data)

	return result
}

// NewFloat64x2Boardcast returns a Float64x2 with every lane set to value.
func NewFloat64x2Boardcast(value float64) Float64x2 {
	buf := [float64_x2_len]float64{value, value}

	return NewFloat64x2(buf[:])
}

func (v Float64x2) Store(receiver []float64) {
	copy(receiver, v.data[:])
}

// Abs returns the absolute values of the elements of x
func (x Float64x2) Abs() Float64x2 {
	var result Float64x2

	for i := range x.data {
		result.data[i] = math.Abs(x.data[i])
	}

	return result
}

// Add performs a fused: x + y.
func (x Float64x2) Add(y Float64x2) Float64x2 {
	var result Float64x2

	for i := range x.data {
		result.data[i] = x.data[i] + y.data[i]
	}

	return result
}

// Div performs a fused: x / y.
func (x Float64x2) Div(y Float64x2) Float64x2 {
	var result Float64x2

	for i := range x.data {
		result.data[i] = x.data[i] / y.data[i]
	}

	return result
}

// Mul performs a fused: x * y.
func (x Float64x2) Mul(y Float64x2) Float64x2 {
	var result Float64x2

	for i := range x.data {
		result.data[i] = x.data[i] * y.data[i]
	}

	return result
}

// MulAdd performs a fused: (x * y) + z.
func (x Float64x2) MulAdd(y, z Float64x2) Float64x2 {
	var result Float64x2

	for i := range x.data {
		result.data[i] = (x.data[i] * y.data[i]) + z.data[i]
	}

	return result
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Float64x2) Max(y Float64x2) Float64x2 {
	var result Float64x2

	for i := range x.data {
		result.data[i] = max(x.data[i], y.data[i])
	}

	return result
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Float64x2) Min(y Float64x2) Float64x2 {
	var result Float64x2

	for i := range x.data {
		result.data[i] = min(x.data[i], y.data[i])
	}

	return result
}

// Neg returns the negation of the elements of x
func (x Float64x2) Neg() Float64x2 {
	var result Float64x2

	for i := range x.data {
		result.data[i] = x.data[i] * -1
	}

	return result
}

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x Float64x2) Scale(y Float64x2) Float64x2 {
	var result Float64x2

	for i := range x.data {
		result.data[i] = x.data[i] * math.Pow(2, y.data[i])
	}

	return result
}

// Sub performs a fused: x - y.
func (x Float64x2) Sub(y Float64x2) Float64x2 {
	var result Float64x2

	for i := range x.data {
		result.data[i] = x.data[i] - y.data[i]
	}

	return result
}

// Sqrt computes the square root of each element.
func (x Float64x2) Sqrt() Float64x2 {
	var result Float64x2

	for i := range x.data {
		result.data[i] = math.Sqrt(x.data[i])
	}

	return result
}
