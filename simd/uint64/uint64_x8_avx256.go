//go:build simd_avx256

package simduint64

import "simd/archsimd"

type Uint64x8 struct {
	data0 archsimd.Uint64x4
	data1 archsimd.Uint64x4
}

func NewUint64x8(data []uint64) (result Uint64x8) {
	var n int
	result.data0, n = archsimd.LoadUint64x4Part(data)
	if n == 4 { // Read atleast 4 points, so there should be more
		result.data1, _ = archsimd.LoadUint64x4Part(data[(uint64_x8_len / 2):])
	}

	return result
}

// NewUint64x8Boardcast returns a Uint64x8 with every lane set to value.
func NewUint64x8Boardcast(value uint64) Uint64x8 {
	return Uint64x8{
		data0: archsimd.BroadcastUint64x4(value),
		data1: archsimd.BroadcastUint64x4(value),
	}
}

func (x Uint64x8) Store(receiver []uint64) {
	switch {
	case len(receiver) == uint64_x8_len:
		x.data0.Store(receiver[:(uint64_x8_len / 2)])
		x.data1.Store(receiver[(uint64_x8_len / 2):])
	case len(receiver) > (uint64_x8_len / 2):
		x.data0.Store(receiver[:(uint64_x8_len / 2)])
		_ = x.data1.StorePart(receiver[(uint64_x8_len / 2):])
	default:
		_ = x.data0.StorePart(receiver)
	}
}

// Add performs a fused: x + y.
func (x Uint64x8) Add(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data0: x.data0.Add(y.data0),
		data1: x.data1.Add(y.data1),
	}
}

// Sub performs a fused: x - y.
func (x Uint64x8) Sub(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data0: x.data0.Sub(y.data0),
		data1: x.data1.Sub(y.data1),
	}
}

// And performs a bitwise AND: x & y.
func (x Uint64x8) And(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data0: x.data0.And(y.data0),
		data1: x.data1.And(y.data1),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Uint64x8) AndNot(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data0: x.data0.AndNot(y.data0),
		data1: x.data1.AndNot(y.data1),
	}
}

// Or performs a bitwise OR: x | y.
func (x Uint64x8) Or(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data0: x.data0.Or(y.data0),
		data1: x.data1.Or(y.data1),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Uint64x8) Xor(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data0: x.data0.Xor(y.data0),
		data1: x.data1.Xor(y.data1),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Uint64x8) Not() Uint64x8 {
	return Uint64x8{
		data0: x.data0.Not(),
		data1: x.data1.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Uint64x8) ShiftLeft(count uint) Uint64x8 {
	return Uint64x8{
		data0: x.data0.ShiftAllLeft(uint64(count)),
		data1: x.data1.ShiftAllLeft(uint64(count)),
	}
}

// The 64-bit variants of the operations below (VPMULLQ, VPMINUQ, VPMAXUQ) are
// AVX-512 only, so on AVX2 each half is emulated in-register (see mul64x4 and
// the compare/blend sequences).

// Mul performs a fused: x * y.
func (x Uint64x8) Mul(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data0: mul64x4(x.data0, y.data0),
		data1: mul64x4(x.data1, y.data1),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Uint64x8) Max(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data0: x.data0.IfElse(x.data0.Greater(y.data0), y.data0),
		data1: x.data1.IfElse(x.data1.Greater(y.data1), y.data1),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Uint64x8) Min(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data0: x.data0.IfElse(x.data0.Less(y.data0), y.data0),
		data1: x.data1.IfElse(x.data1.Less(y.data1), y.data1),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (logical).
func (x Uint64x8) ShiftRight(count uint) Uint64x8 {
	return Uint64x8{
		data0: x.data0.ShiftAllRight(uint64(count)),
		data1: x.data1.ShiftAllRight(uint64(count)),
	}
}
