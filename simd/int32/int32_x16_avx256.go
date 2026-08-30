//go:build simd_avx256

package simdint32

import "simd/archsimd"

type Int32x16 struct {
	data0 archsimd.Int32x8
	data1 archsimd.Int32x8
}

func NewInt32x16(data []int32) Int32x16 {
	var result Int32x16

	var n int
	result.data0, n = archsimd.LoadInt32x8Part(data)
	if n == 8 { // Read atleast 8 points, so there should be more
		result.data1, _ = archsimd.LoadInt32x8Part(data[(int32_x16_len / 2):])
	}

	return result
}

func NewInt32x16Boardcast(value int32) Int32x16 {
	return Int32x16{
		data0: archsimd.BroadcastInt32x8(value),
		data1: archsimd.BroadcastInt32x8(value),
	}
}

func (x Int32x16) Store(receiver []int32) {
	switch {
	case len(receiver) == int32_x16_len:
		x.data0.Store(receiver[:(int32_x16_len / 2)])
		x.data1.Store(receiver[(int32_x16_len / 2):])
	case len(receiver) > (int32_x16_len / 2):
		x.data0.Store(receiver[:(int32_x16_len / 2)])
		_ = x.data1.StorePart(receiver[(int32_x16_len / 2):])
	default:
		_ = x.data0.StorePart(receiver)
	}
}

// Abs returns the absolute values of the elements of x
func (x Int32x16) Abs() Int32x16 {
	return Int32x16{
		data0: x.data0.Abs(),
		data1: x.data1.Abs(),
	}
}

// Add performs a fused: x + y.
func (x Int32x16) Add(y Int32x16) Int32x16 {
	return Int32x16{
		data0: x.data0.Add(y.data0),
		data1: x.data1.Add(y.data1),
	}
}

// Mul performs a fused: x * y.
func (x Int32x16) Mul(y Int32x16) Int32x16 {
	return Int32x16{
		data0: x.data0.Mul(y.data0),
		data1: x.data1.Mul(y.data1),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Int32x16) Max(y Int32x16) Int32x16 {
	return Int32x16{
		data0: x.data0.Max(y.data0),
		data1: x.data1.Max(y.data1),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Int32x16) Min(y Int32x16) Int32x16 {
	return Int32x16{
		data0: x.data0.Min(y.data0),
		data1: x.data1.Min(y.data1),
	}
}

// Neg returns the negation of the elements of x
func (x Int32x16) Neg() Int32x16 {
	return Int32x16{
		data0: x.data0.Neg(),
		data1: x.data1.Neg(),
	}
}

// Sub performs a fused: x - y.
func (x Int32x16) Sub(y Int32x16) Int32x16 {
	return Int32x16{
		data0: x.data0.Sub(y.data0),
		data1: x.data1.Sub(y.data1),
	}
}

// And performs a bitwise AND: x & y.
func (x Int32x16) And(y Int32x16) Int32x16 {
	return Int32x16{
		data0: x.data0.And(y.data0),
		data1: x.data1.And(y.data1),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Int32x16) AndNot(y Int32x16) Int32x16 {
	return Int32x16{
		data0: x.data0.AndNot(y.data0),
		data1: x.data1.AndNot(y.data1),
	}
}

// Or performs a bitwise OR: x | y.
func (x Int32x16) Or(y Int32x16) Int32x16 {
	return Int32x16{
		data0: x.data0.Or(y.data0),
		data1: x.data1.Or(y.data1),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Int32x16) Xor(y Int32x16) Int32x16 {
	return Int32x16{
		data0: x.data0.Xor(y.data0),
		data1: x.data1.Xor(y.data1),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Int32x16) Not() Int32x16 {
	return Int32x16{
		data0: x.data0.Not(),
		data1: x.data1.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Int32x16) ShiftLeft(count uint) Int32x16 {
	return Int32x16{
		data0: x.data0.ShiftAllLeft(uint64(count)),
		data1: x.data1.ShiftAllLeft(uint64(count)),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (arithmetic).
func (x Int32x16) ShiftRight(count uint) Int32x16 {
	return Int32x16{
		data0: x.data0.ShiftAllRight(uint64(count)),
		data1: x.data1.ShiftAllRight(uint64(count)),
	}
}
