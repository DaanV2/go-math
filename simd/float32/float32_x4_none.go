//go:build simd_none || (!simd_avx512 && !simd_avx256)

package simdfloat32

import "math"

type Float32x4 struct {
	data [float32_x4_len]float32
}

func NewFloat32x4(data []float32) Float32x4 {
	var result Float32x4
	copy(result.data[:], data)

	return result
}

func NewFloat32x4Boardcast(value float32) Float32x4 {
	buf := [float32_x4_len]float32{value, value, value, value}

	return NewFloat32x4(buf[:])
}

func (v Float32x4) Store(receiver []float32) {
	copy(receiver, v.data[:])
}

// Abs returns the absolute values of the elements of x
func (x Float32x4) Abs() Float32x4 {
	var result Float32x4

	for i := range x.data {
		result.data[i] = float32(math.Abs(float64(x.data[i])))
	}

	return result
}

// Add performs a fused: x + y.
func (x Float32x4) Add(y Float32x4) Float32x4 {
	var result Float32x4

	for i := range x.data {
		result.data[i] = x.data[i] + y.data[i]
	}

	return result
}

// Div performs a fused: x / y.
func (x Float32x4) Div(y Float32x4) Float32x4 {
	var result Float32x4

	for i := range x.data {
		result.data[i] = x.data[i] / y.data[i]
	}

	return result
}

// Mul performs a fused: x * y.
func (x Float32x4) Mul(y Float32x4) Float32x4 {
	var result Float32x4

	for i := range x.data {
		result.data[i] = x.data[i] * y.data[i]
	}

	return result
}

// MulAdd performs a fused: (x * y) + z.
func (x Float32x4) MulAdd(y, z Float32x4) Float32x4 {
	var result Float32x4

	for i := range x.data {
		result.data[i] = (x.data[i] * y.data[i]) + z.data[i]
	}

	return result
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Float32x4) Max(y Float32x4) Float32x4 {
	var result Float32x4

	for i := range x.data {
		result.data[i] = max(x.data[i], y.data[i])
	}

	return result
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Float32x4) Min(y Float32x4) Float32x4 {
	var result Float32x4

	for i := range x.data {
		result.data[i] = min(x.data[i], y.data[i])
	}

	return result
}

// Neg returns the negation of the elements of x
func (x Float32x4) Neg() Float32x4 {
	var result Float32x4

	for i := range x.data {
		result.data[i] = x.data[i] * -1
	}

	return result
}

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x Float32x4) Scale(y Float32x4) Float32x4 {
	var result Float32x4

	for i := range x.data {
		result.data[i] = x.data[i] * float32(math.Pow(2, float64(y.data[i])))
	}

	return result
}

// Sub performs a fused: x - y.
func (x Float32x4) Sub(y Float32x4) Float32x4 {
	var result Float32x4

	for i := range x.data {
		result.data[i] = x.data[i] - y.data[i]
	}

	return result
}

// Sqrt computes the square root of each element.
func (x Float32x4) Sqrt() Float32x4 {
	var result Float32x4

	for i := range x.data {
		result.data[i] = float32(math.Sqrt(float64(x.data[i])))
	}

	return result
}
