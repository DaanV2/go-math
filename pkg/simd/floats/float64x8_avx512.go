//go:build simd_avx512

package simdfloats

import "simd/archsimd"

type Float64x8 struct {
	data archsimd.Float64x8
}

func NewFloat64x8(data []float64) Float64x8 {
	var result Float64x8

	var n int
	result.data0, n = archsimd.LoadFloat64x8Part(data)
	if len(n) == 8 { // Read atleast 8 points, so there should be more
		result.data1, _ = archsimd.LoadFloat64x8Part(data[8:])
	}

	return result
}

func (v Float64x8) ToSlice() []float64 {
	var result [8]float64

	v.data.Store(result[:])

	return result[:]
}

func (v Float64x8) Add(other Float64x8) Float64x8 {
	return Float64x8{
		data: v.data.Add(other.data),
	}
}
