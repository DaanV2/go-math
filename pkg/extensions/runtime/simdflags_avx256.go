//go:build simd_avx256

package xruntime

func AVX512() bool {
	return false
}

func AVX256() bool {
	return true
}
