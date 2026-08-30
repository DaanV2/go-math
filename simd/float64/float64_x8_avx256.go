//go:build simd_avx256

package simdfloat64

import (
	"math"
	"simd/archsimd"
)

type Float64x8 struct {
	data0 archsimd.Float64x4
	data1 archsimd.Float64x4
}

func NewFloat64x8(data []float64) Float64x8 {
	var result Float64x8

	var n int
	result.data0, n = archsimd.LoadFloat64x4Part(data)
	if n == 4 { // Read atleast 4 points, so there should be more
		result.data1, _ = archsimd.LoadFloat64x4Part(data[(float64_x8_len / 2):])
	}

	return result
}

// NewFloat64x8Boardcast returns a Float64x8 with every lane set to value.
func NewFloat64x8Boardcast(value float64) Float64x8 {
	return Float64x8{
		data0: archsimd.BroadcastFloat64x4(value),
		data1: archsimd.BroadcastFloat64x4(value),
	}
}

func (x Float64x8) Store(receiver []float64) {
	switch {
	case len(receiver) == float64_x8_len:
		x.data0.Store(receiver[:(float64_x8_len / 2)])
		x.data1.Store(receiver[(float64_x8_len / 2):])
	case len(receiver) > (float64_x8_len / 2):
		x.data0.Store(receiver[:(float64_x8_len / 2)])
		_ = x.data1.StorePart(receiver[(float64_x8_len / 2):])
	default:
		_ = x.data0.StorePart(receiver)
	}
}

// Abs returns the absolute values of the elements of x
func (x Float64x8) Abs() Float64x8 {
	return Float64x8{
		data0: x.data0.Abs(),
		data1: x.data1.Abs(),
	}
}

// Add performs a fused: x + y.
func (x Float64x8) Add(y Float64x8) Float64x8 {
	return Float64x8{
		data0: x.data0.Add(y.data0),
		data1: x.data1.Add(y.data1),
	}
}

// Div performs a fused: x / y.
func (x Float64x8) Div(y Float64x8) Float64x8 {
	return Float64x8{
		data0: x.data0.Div(y.data0),
		data1: x.data1.Div(y.data1),
	}
}

// Mul performs a fused: x * y.
func (x Float64x8) Mul(y Float64x8) Float64x8 {
	return Float64x8{
		data0: x.data0.Mul(y.data0),
		data1: x.data1.Mul(y.data1),
	}
}

// MulAdd performs a fused: (x * y) + z.
func (x Float64x8) MulAdd(y, z Float64x8) Float64x8 {
	return Float64x8{
		data0: x.data0.MulAdd(y.data0, z.data0),
		data1: x.data1.MulAdd(y.data1, z.data1),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Float64x8) Max(y Float64x8) Float64x8 {
	return Float64x8{
		data0: x.data0.Max(y.data0),
		data1: x.data1.Max(y.data1),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Float64x8) Min(y Float64x8) Float64x8 {
	return Float64x8{
		data0: x.data0.Min(y.data0),
		data1: x.data1.Min(y.data1),
	}
}

// Neg returns the negation of the elements of x
func (x Float64x8) Neg() Float64x8 {
	return Float64x8{
		data0: x.data0.Neg(),
		data1: x.data1.Neg(),
	}
}

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x Float64x8) Scale(y Float64x8) Float64x8 {
	var result [float64_x8_len]float64
	var vx [float64_x8_len]float64
	var vy [float64_x8_len]float64

	x.Store(vx[:])
	y.Store(vy[:])

	for i := range vx {
		result[i] = vx[i] * math.Pow(2, vy[i])
	}

	return NewFloat64x8(result[:])
}

// Sub performs a fused: x - y.
func (x Float64x8) Sub(y Float64x8) Float64x8 {
	return Float64x8{
		data0: x.data0.Sub(y.data0),
		data1: x.data1.Sub(y.data1),
	}
}

// Sqrt computes the square root of each element.
func (x Float64x8) Sqrt() Float64x8 {
	return Float64x8{
		data0: x.data0.Sqrt(),
		data1: x.data1.Sqrt(),
	}
}
