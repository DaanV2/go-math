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

func (v Float64x8) Store(receiver []float64) {
	if len(receiver) == 8 {
		v.data0.Store(receiver[:4])
		v.data1.Store(receiver[4:])
	} else if len(receiver) > 4 {
		v.data0.Store(receiver[:4])
		_ = v.data1.StorePart(receiver[4:])
	} else {
		_ = v.data0.StorePart(receiver)
	}
}

func (v Float64x8) Add(other Float64x8) Float64x8 {
	return Float64x8{
		data0: v.data0.Add(other.data0),
		data1: v.data1.Add(other.data1),
	}
}
