//go:build simd_avx256

package simdint64

import "simd/archsimd"

// On AVX2 the 64-bit VPABSQ/VPMULLQ/VPMINSQ/VPMAXSQ/VPSRAQ instructions are
// unavailable (they were only added in AVX-512), so the operations below are
// emulated in-register with AVX2 compare/blend/widening-multiply sequences.

// mul64x2 computes the low 64 bits of a*b for each lane using 32-bit partial
// products, since AVX2 has no packed 64-bit multiply. The low 64 bits are the
// same for signed and unsigned multiplication.
func mul64x2(a, b archsimd.Uint64x2) archsimd.Uint64x2 {
	a32 := a.AsUint32x4()
	b32 := b.AsUint32x4()
	aHi := a.ShiftAllRight(32).AsUint32x4()
	bHi := b.ShiftAllRight(32).AsUint32x4()

	lo := a32.MulWidenEven(b32)                               // aLo * bLo
	cross := aHi.MulWidenEven(b32).Add(a32.MulWidenEven(bHi)) // aHi*bLo + aLo*bHi
	return lo.Add(cross.ShiftAllLeft(32))
}

// Abs returns the absolute values of the elements of x
func (x Int64x2) Abs() Int64x2 {
	neg := x.data.Neg()
	return Int64x2{
		data: x.data.IfElse(x.data.Greater(neg), neg), // max(x, -x)
	}
}

// Mul performs a fused: x * y.
func (x Int64x2) Mul(y Int64x2) Int64x2 {
	return Int64x2{
		data: mul64x2(x.data.AsUint64x2(), y.data.AsUint64x2()).AsInt64x2(),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Int64x2) Max(y Int64x2) Int64x2 {
	return Int64x2{
		data: x.data.IfElse(x.data.Greater(y.data), y.data),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Int64x2) Min(y Int64x2) Int64x2 {
	return Int64x2{
		data: x.data.IfElse(x.data.Less(y.data), y.data),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (arithmetic).
func (x Int64x2) ShiftRight(count uint) Int64x2 {
	c := count & 63
	if c == 0 {
		return x
	}

	logical := x.data.AsUint64x2().ShiftAllRight(uint64(c)).AsInt64x2()
	fill := archsimd.BroadcastInt64x2(int64(-1) << (64 - c)) // top c bits set
	neg := x.data.Less(archsimd.BroadcastInt64x2(0))

	return Int64x2{
		data: logical.Or(fill).IfElse(neg, logical), // negative lanes get the sign fill
	}
}
