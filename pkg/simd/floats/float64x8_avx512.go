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
