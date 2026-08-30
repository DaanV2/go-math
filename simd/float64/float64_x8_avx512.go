//go:build simd_avx512

package simdfloat64

import "simd/archsimd"

type Float64x8 struct {
	data archsimd.Float64x8
}

func NewFloat64x8(data []float64) Float64x8 {
	var result Float64x8

	result.data, _ = archsimd.LoadFloat64x8Part(data)

	return result
}

func NewFloat64x8Boardcast(value float64) Float64x8 {
	return Float64x8{
		data: archsimd.BroadcastFloat64x8(value),
	}
}

func (v Float64x8) Store(receiver []float64) {
	v.data.StorePart(receiver)
}

// Abs returns the absolute values of the elements of x
func (x Float64x8) Abs() Float64x8 {
	return Float64x8{
		data: x.data.Abs(),
	}
}

// Add performs a fused: x + y.
func (v Float64x8) Add(other Float64x8) Float64x8 {
	return Float64x8{
		data: v.data.Add(other.data),
	}
}

// Div performs a fused: x / y.
func (x Float64x8) Div(y Float64x8) Float64x8 {
	return Float64x8{
		data: x.data.Div(y.data),
	}
}

// Mul performs a fused: x * y.
func (x Float64x8) Mul(y Float64x8) Float64x8 {
	return Float64x8{
		data: x.data.Mul(y.data),
	}
}

// MulAdd performs a fused: (x * y) + z.
func (x Float64x8) MulAdd(y, z Float64x8) Float64x8 {
	return Float64x8{
		data: x.data.MulAdd(y.data, z.data),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Float64x8) Max(y Float64x8) Float64x8 {
	return Float64x8{
		data: x.data.Max(y.data),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Float64x8) Min(y Float64x8) Float64x8 {
	return Float64x8{
		data: x.data.Min(y.data),
	}
}

// Neg returns the negation of the elements of x
func (x Float64x8) Neg() Float64x8 {
	return Float64x8{
		data: x.data.Neg(),
	}
}

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x Float64x8) Scale(y Float64x8) Float64x8 {
	return Float64x8{
		data: x.data.Scale(y.data),
	}
}

// Sub performs a fused: x - y.
func (x Float64x8) Sub(y Float64x8) Float64x8 {
	return Float64x8{
		data: x.data.Sub(y.data),
	}
}

// Sqrt computes the square root of each element.
func (x Float64x8) Sqrt() Float64x8 {
	return Float64x8{
		data: x.data.Sqrt(),
	}
}
