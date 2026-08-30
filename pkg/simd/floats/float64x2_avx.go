//go:build simd_avx256 || simd_avx512

package simdfloats

import "simd/archsimd"

type Float64x2 struct {
	data archsimd.Float64x2
}

func NewFloat64x2(data []float64) Float64x2 {
	var result Float64x2

	result.data, _ = archsimd.LoadFloat64x2Part(data)

	return result
}

func (x Float64x2) Store(receiver []float64) {
	if len(receiver) >= float64_x2_len {
		x.data.Store(receiver[:float64_x2_len])
	} else {
		x.data.StorePart(receiver)
	}
}

// Abs returns the absolute values of the elements of x
func (x Float64x2) Abs() Float64x2 {
	return Float64x2{
		data: x.data.Abs(),
	}
}

// Add performs a fused: x + y.
func (x Float64x2) Add(y Float64x2) Float64x2 {
	return Float64x2{
		data: x.data.Add(y.data),
	}
}

// Div performs a fused: x / y.
func (x Float64x2) Div(y Float64x2) Float64x2 {
	return Float64x2{
		data: x.data.Div(y.data),
	}
}

// Mul performs a fused: x * y.
func (x Float64x2) Mul(y Float64x2) Float64x2 {
	return Float64x2{
		data: x.data.Mul(y.data),
	}
}

// MulAdd performs a fused: (x * y) + z.
func (x Float64x2) MulAdd(y, z Float64x2) Float64x2 {
	return Float64x2{
		data: x.data.MulAdd(y.data, z.data),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Float64x2) Max(y Float64x2) Float64x2 {
	return Float64x2{
		data: x.data.Max(y.data),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Float64x2) Min(y Float64x2) Float64x2 {
	return Float64x2{
		data: x.data.Min(y.data),
	}
}

// Neg returns the negation of the elements of x
func (x Float64x2) Neg() Float64x2 {
	return Float64x2{
		data: x.data.Neg(),
	}
}

// Sub performs a fused: x - y.
func (x Float64x2) Sub(y Float64x2) Float64x2 {
	return Float64x2{
		data: x.data.Sub(y.data),
	}
}

// Sqrt computes the square root of each element.
func (x Float64x2) Sqrt() Float64x2 {
	return Float64x2{
		data: x.data.Sqrt(),
	}
}
