//go:build !goexperiment.simd && (!simd_avx512 || !simd_avx256)

package xruntime

func AVX512() bool {
	return false
}

func AVX256() bool {
	return false
}
