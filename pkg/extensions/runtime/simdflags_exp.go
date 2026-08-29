//go:build goexperiment.simd && (!simd_avx512 || !simd_avx256)

package xruntime

import "simd/archsimd"

func AVX512() bool {
	return archsimd.X86.AVX512()
}

func AVX256() bool {
	return archsimd.X86.AVX2() || archsimd.X86.AVX()
}
