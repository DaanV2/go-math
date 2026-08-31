//go:build simd_avx256

package simdint64

import "simd/archsimd"

type Int64x8 struct {
	data0 archsimd.Int64x4
	data1 archsimd.Int64x4
}

func NewInt64x8(data []int64) (result Int64x8) {
	var n int
	result.data0, n = archsimd.LoadInt64x4Part(data)
	if n == 4 { // Read atleast 4 points, so there should be more
		result.data1, _ = archsimd.LoadInt64x4Part(data[(int64_x8_len / 2):])
	}

	return result
}

// NewInt64x8Boardcast returns an Int64x8 with every lane set to value.
func NewInt64x8Boardcast(value int64) Int64x8 {
	return Int64x8{
		data0: archsimd.BroadcastInt64x4(value),
		data1: archsimd.BroadcastInt64x4(value),
	}
}

func (x Int64x8) Store(receiver []int64) {
	switch {
	case len(receiver) == int64_x8_len:
		x.data0.Store(receiver[:(int64_x8_len / 2)])
		x.data1.Store(receiver[(int64_x8_len / 2):])
	case len(receiver) > (int64_x8_len / 2):
		x.data0.Store(receiver[:(int64_x8_len / 2)])
		_ = x.data1.StorePart(receiver[(int64_x8_len / 2):])
	default:
		_ = x.data0.StorePart(receiver)
	}
}

// Add performs a fused: x + y.
func (x Int64x8) Add(y Int64x8) Int64x8 {
	return Int64x8{
		data0: x.data0.Add(y.data0),
		data1: x.data1.Add(y.data1),
	}
}

// Neg returns the negation of the elements of x
func (x Int64x8) Neg() Int64x8 {
	return Int64x8{
		data0: x.data0.Neg(),
		data1: x.data1.Neg(),
	}
}

// Sub performs a fused: x - y.
func (x Int64x8) Sub(y Int64x8) Int64x8 {
	return Int64x8{
		data0: x.data0.Sub(y.data0),
		data1: x.data1.Sub(y.data1),
	}
}

// And performs a bitwise AND: x & y.
func (x Int64x8) And(y Int64x8) Int64x8 {
	return Int64x8{
		data0: x.data0.And(y.data0),
		data1: x.data1.And(y.data1),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Int64x8) AndNot(y Int64x8) Int64x8 {
	return Int64x8{
		data0: x.data0.AndNot(y.data0),
		data1: x.data1.AndNot(y.data1),
	}
}

// Or performs a bitwise OR: x | y.
func (x Int64x8) Or(y Int64x8) Int64x8 {
	return Int64x8{
		data0: x.data0.Or(y.data0),
		data1: x.data1.Or(y.data1),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Int64x8) Xor(y Int64x8) Int64x8 {
	return Int64x8{
		data0: x.data0.Xor(y.data0),
		data1: x.data1.Xor(y.data1),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Int64x8) Not() Int64x8 {
	return Int64x8{
		data0: x.data0.Not(),
		data1: x.data1.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Int64x8) ShiftLeft(count uint) Int64x8 {
	return Int64x8{
		data0: x.data0.ShiftAllLeft(uint64(count)),
		data1: x.data1.ShiftAllLeft(uint64(count)),
	}
}

// The 64-bit variants of the operations below are AVX-512 only, so on AVX2 each
// half is emulated in-register (see mul64x4 and the compare/blend sequences).

// Abs returns the absolute values of the elements of x
func (x Int64x8) Abs() Int64x8 {
	neg0, neg1 := x.data0.Neg(), x.data1.Neg()
	return Int64x8{
		data0: x.data0.IfElse(x.data0.Greater(neg0), neg0),
		data1: x.data1.IfElse(x.data1.Greater(neg1), neg1),
	}
}

// Mul performs a fused: x * y.
func (x Int64x8) Mul(y Int64x8) Int64x8 {
	return Int64x8{
		data0: mul64x4(x.data0.AsUint64x4(), y.data0.AsUint64x4()).AsInt64x4(),
		data1: mul64x4(x.data1.AsUint64x4(), y.data1.AsUint64x4()).AsInt64x4(),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Int64x8) Max(y Int64x8) Int64x8 {
	return Int64x8{
		data0: x.data0.IfElse(x.data0.Greater(y.data0), y.data0),
		data1: x.data1.IfElse(x.data1.Greater(y.data1), y.data1),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Int64x8) Min(y Int64x8) Int64x8 {
	return Int64x8{
		data0: x.data0.IfElse(x.data0.Less(y.data0), y.data0),
		data1: x.data1.IfElse(x.data1.Less(y.data1), y.data1),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (arithmetic).
func (x Int64x8) ShiftRight(count uint) Int64x8 {
	c := count & 63
	if c == 0 {
		return x
	}

	fill := archsimd.BroadcastInt64x4(int64(-1) << (64 - c)) // top c bits set
	zero := archsimd.BroadcastInt64x4(0)

	logical0 := x.data0.AsUint64x4().ShiftAllRight(uint64(c)).AsInt64x4()
	logical1 := x.data1.AsUint64x4().ShiftAllRight(uint64(c)).AsInt64x4()

	return Int64x8{
		data0: logical0.Or(fill).IfElse(x.data0.Less(zero), logical0),
		data1: logical1.Or(fill).IfElse(x.data1.Less(zero), logical1),
	}
}
