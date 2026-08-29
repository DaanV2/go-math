//go:build simd_avx256

package simdfloats

import "simd/archsimd"

type Float64x8 struct {
	data0 archsimd.Float64x4
	data1 archsimd.Float64x4
}
