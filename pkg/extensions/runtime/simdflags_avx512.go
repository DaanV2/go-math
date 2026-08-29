//go:build simd_avx512

package xruntime

func AVX512() bool {
	return true
}

func AVX256() bool {
	return true
}
