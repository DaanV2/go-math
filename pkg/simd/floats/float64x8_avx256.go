//go:build simd_avx256

package simdfloats

import "simd/archsimd"

type Float64x8 struct {
	data0 archsimd.Float64x4
	data1 archsimd.Float64x4
}

func NewFloat64x8(data []float64) Float64x8 {
	var result Float64x8

	var n int
	result.data0, n = archsimd.LoadFloat64x4Part(data)
	if n == 4 { // Read atleast 4 points, so there should be more
		result.data1, _ = archsimd.LoadFloat64x4Part(data[4:])
	}

	return result
}

func (x Float64x8) Store(receiver []float64) {
	if len(receiver) == 8 {
		x.data0.Store(receiver[:4])
		x.data1.Store(receiver[4:])
	} else if len(receiver) > 4 {
		x.data0.Store(receiver[:4])
		_ = x.data1.StorePart(receiver[4:])
	} else {
		_ = x.data0.StorePart(receiver)
	}
}

// Add performs a fused: x + y.
func (x Float64x8) Add(y Float64x8) Float64x8 {
	return Float64x8{
		data0: x.data0.Add(y.data0),
		data1: x.data1.Add(y.data1),
	}
}

// Sub performs a fused: x - y.
func (x Float64x8) Sub(y Float64x8) Float64x8 {
	return Float64x8{
		data0: x.data0.Sub(y.data0),
		data1: x.data1.Sub(y.data1),
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
func (x Float64x8) MulAdd(y Float64x8, z Float64x8) Float64x8 {
	return Float64x8{
		data0: x.data0.MulAdd(y.data0, z.data0),
		data1: x.data1.MulAdd(y.data1, z.data1),
	}
}
