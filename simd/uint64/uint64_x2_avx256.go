//go:build simd_avx256

package simduint64

import "simd/archsimd"

// On AVX2 the 64-bit VPMULLQ/VPMINUQ/VPMAXUQ instructions are unavailable (they
// were only added in AVX-512), so the operations below are emulated in-register
// with AVX2 widening-multiply and compare/blend sequences.

// mul64x2 computes the low 64 bits of a*b for each lane using 32-bit partial
// products, since AVX2 has no packed 64-bit multiply.
func mul64x2(a, b archsimd.Uint64x2) archsimd.Uint64x2 {
	a32 := a.AsUint32x4()
	b32 := b.AsUint32x4()
	aHi := a.ShiftAllRight(32).AsUint32x4()
	bHi := b.ShiftAllRight(32).AsUint32x4()

	lo := a32.MulWidenEven(b32)                               // aLo * bLo
	cross := aHi.MulWidenEven(b32).Add(a32.MulWidenEven(bHi)) // aHi*bLo + aLo*bHi
	return lo.Add(cross.ShiftAllLeft(32))
}

// Mul performs a fused: x * y.
func (x Uint64x2) Mul(y Uint64x2) Uint64x2 {
	return Uint64x2{
		data: mul64x2(x.data, y.data),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Uint64x2) Max(y Uint64x2) Uint64x2 {
	return Uint64x2{
		data: x.data.IfElse(x.data.Greater(y.data), y.data),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Uint64x2) Min(y Uint64x2) Uint64x2 {
	return Uint64x2{
		data: x.data.IfElse(x.data.Less(y.data), y.data),
	}
}
