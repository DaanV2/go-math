//go:build simd_avx512

package simdfloats

import "simd/archsimd"

type Float64x8 struct {
	data archsimd.Float64x8
}

func NewFloat64x8(data []float64) Float64x8 {
	var result Float64x8

	result.data, _ = archsimd.LoadFloat64x8Part(data)

	return result
}

func (v Float64x8) Store(receiver []float64) {
	v.data.StorePart(receiver)
}

func (v Float64x8) Add(other Float64x8) Float64x8 {
	return Float64x8{
		data: v.data.Add(other.data),
	}
}

// Sub performs a fused: x - y.
func (x Float64x8) Sub(y Float64x8) Float64x8 {
	return Float64x8{
		data: x.data.Sub(y.data),
	}
}

// Mul performs a fused: x * y.
func (x Float64x8) Mul(y Float64x8) Float64x8 {
	return Float64x8{
		data: x.data.Mul(y.data),
	}
}

// MulAdd performs a fused: (x * y) + z.
func (x Float64x8) MulAdd(y Float64x8, z Float64x8) Float64x8 {
	return Float64x8{
		data: x.data.MulAdd(y.data, z.data),
	}
}
