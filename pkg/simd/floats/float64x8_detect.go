//go:build simd_detect || (!simd_avx256 && !simd_avx512 && !simd_none)

package simdfloats

import (
	"simd/archsimd"

	xruntime "github.com/daanv2/go-math/pkg/extensions/runtime"
)

type Float64x8 struct {
	data [8]float64
}

func NewFloat64x8(data []float64) Float64x8 {
	var result Float64x8
	copy(result.data[:], data)

	return result
}

func (v Float64x8) Store(receiver []float64) {
	copy(receiver, v.data[:])
}

func (v Float64x8) Add(other Float64x8) Float64x8 {
	var result Float64x8

	switch {
	case xruntime.AVX512():
		v1 := archsimd.LoadFloat64x8(v.data[:])
		v2 := archsimd.LoadFloat64x8(other.data[:])
		v1.Add(v2).Store(result.data[:])
	case xruntime.AVX256():
		// First 4
		v1 := archsimd.LoadFloat64x4(v.data[:4])
		v2 := archsimd.LoadFloat64x4(other.data[:4])
		v1.Add(v2).Store(result.data[:4])

		// Last 4
		v1 = archsimd.LoadFloat64x4(v.data[4:])
		v2 = archsimd.LoadFloat64x4(other.data[4:])
		v1.Add(v2).Store(result.data[4:])
	default:
		for i := range v.data {
			result.data[i] = v.data[i] + other.data[i]
		}
	}

	return result
}
