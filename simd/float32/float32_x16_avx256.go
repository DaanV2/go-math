//go:build simd_avx256

package simdfloat32

import (
	"math"
	"simd/archsimd"
)

type Float32x16 struct {
	data0 archsimd.Float32x8
	data1 archsimd.Float32x8
}

func NewFloat32x16(data []float32) Float32x16 {
	var result Float32x16

	var n int
	result.data0, n = archsimd.LoadFloat32x8Part(data)
	if n == 8 { // Read atleast 8 points, so there should be more
		result.data1, _ = archsimd.LoadFloat32x8Part(data[(float32_x16_len / 2):])
	}

	return result
}

func NewFloat32x16Boardcast(value float32) Float32x16 {
	return Float32x16{
		data0: archsimd.BroadcastFloat32x8(value),
		data1: archsimd.BroadcastFloat32x8(value),
	}
}

func (x Float32x16) Store(receiver []float32) {
	switch {
	case len(receiver) == float32_x16_len:
		x.data0.Store(receiver[:(float32_x16_len / 2)])
		x.data1.Store(receiver[(float32_x16_len / 2):])
	case len(receiver) > (float32_x16_len / 2):
		x.data0.Store(receiver[:(float32_x16_len / 2)])
		_ = x.data1.StorePart(receiver[(float32_x16_len / 2):])
	default:
		_ = x.data0.StorePart(receiver)
	}
}

// Abs returns the absolute values of the elements of x
func (x Float32x16) Abs() Float32x16 {
	return Float32x16{
		data0: x.data0.Abs(),
		data1: x.data1.Abs(),
	}
}

// Add performs a fused: x + y.
func (x Float32x16) Add(y Float32x16) Float32x16 {
	return Float32x16{
		data0: x.data0.Add(y.data0),
		data1: x.data1.Add(y.data1),
	}
}

// Div performs a fused: x / y.
func (x Float32x16) Div(y Float32x16) Float32x16 {
	return Float32x16{
		data0: x.data0.Div(y.data0),
		data1: x.data1.Div(y.data1),
	}
}

// Mul performs a fused: x * y.
func (x Float32x16) Mul(y Float32x16) Float32x16 {
	return Float32x16{
		data0: x.data0.Mul(y.data0),
		data1: x.data1.Mul(y.data1),
	}
}

// MulAdd performs a fused: (x * y) + z.
func (x Float32x16) MulAdd(y, z Float32x16) Float32x16 {
	return Float32x16{
		data0: x.data0.MulAdd(y.data0, z.data0),
		data1: x.data1.MulAdd(y.data1, z.data1),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Float32x16) Max(y Float32x16) Float32x16 {
	return Float32x16{
		data0: x.data0.Max(y.data0),
		data1: x.data1.Max(y.data1),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Float32x16) Min(y Float32x16) Float32x16 {
	return Float32x16{
		data0: x.data0.Min(y.data0),
		data1: x.data1.Min(y.data1),
	}
}

// Neg returns the negation of the elements of x
func (x Float32x16) Neg() Float32x16 {
	return Float32x16{
		data0: x.data0.Neg(),
		data1: x.data1.Neg(),
	}
}

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x Float32x16) Scale(y Float32x16) Float32x16 {
	var result [float32_x16_len]float32
	var vx [float32_x16_len]float32
	var vy [float32_x16_len]float32

	x.Store(vx[:])
	y.Store(vy[:])

	for i := range vx {
		result[i] = vx[i] * float32(math.Pow(2, float64(vy[i])))
	}

	return NewFloat32x16(result[:])
}

// Sub performs a fused: x - y.
func (x Float32x16) Sub(y Float32x16) Float32x16 {
	return Float32x16{
		data0: x.data0.Sub(y.data0),
		data1: x.data1.Sub(y.data1),
	}
}

// Sqrt computes the square root of each element.
func (x Float32x16) Sqrt() Float32x16 {
	return Float32x16{
		data0: x.data0.Sqrt(),
		data1: x.data1.Sqrt(),
	}
}
