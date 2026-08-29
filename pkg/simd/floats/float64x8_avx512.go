//go:build simd_avx512

package simdfloats

import "simd/archsimd"

type Float64x8 struct {
	data archsimd.Float64x8
}
