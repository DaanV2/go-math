package simd

import "simd/archsimd"

func AVX512() bool {
	return archsimd.X86.AVX512()
}

func AVX256() bool {
	return archsimd.X86.AVX2() || archsimd.X86.AVX()
}
