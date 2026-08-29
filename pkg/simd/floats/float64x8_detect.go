//go:build simd_detect || (!simd_avx256 && !simd_avx512 && !simd_none)

package simdfloats

type Float64x8 struct {
	data [8]float64
}
