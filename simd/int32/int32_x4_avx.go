//go:build simd_avx256 || simd_avx512

package simdint32

import "simd/archsimd"

type Int32x4 struct {
	data archsimd.Int32x4
}

func NewInt32x4(data []int32) Int32x4 {
	var result Int32x4

	result.data, _ = archsimd.LoadInt32x4Part(data)

	return result
}

// NewInt32x4Boardcast returns an Int32x4 with every lane set to value.
func NewInt32x4Boardcast(value int32) Int32x4 {
	return Int32x4{
		data: archsimd.BroadcastInt32x4(value),
	}
}

func (x Int32x4) Store(receiver []int32) {
	if len(receiver) >= int32_x4_len {
		x.data.Store(receiver[:int32_x4_len])
	} else {
		x.data.StorePart(receiver)
	}
}

// Abs returns the absolute values of the elements of x
func (x Int32x4) Abs() Int32x4 {
	return Int32x4{
		data: x.data.Abs(),
	}
}

// Add performs a fused: x + y.
func (x Int32x4) Add(y Int32x4) Int32x4 {
	return Int32x4{
		data: x.data.Add(y.data),
	}
}

// Mul performs a fused: x * y.
func (x Int32x4) Mul(y Int32x4) Int32x4 {
	return Int32x4{
		data: x.data.Mul(y.data),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Int32x4) Max(y Int32x4) Int32x4 {
	return Int32x4{
		data: x.data.Max(y.data),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Int32x4) Min(y Int32x4) Int32x4 {
	return Int32x4{
		data: x.data.Min(y.data),
	}
}

// Neg returns the negation of the elements of x
func (x Int32x4) Neg() Int32x4 {
	return Int32x4{
		data: x.data.Neg(),
	}
}

// Sub performs a fused: x - y.
func (x Int32x4) Sub(y Int32x4) Int32x4 {
	return Int32x4{
		data: x.data.Sub(y.data),
	}
}

// And performs a bitwise AND: x & y.
func (x Int32x4) And(y Int32x4) Int32x4 {
	return Int32x4{
		data: x.data.And(y.data),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Int32x4) AndNot(y Int32x4) Int32x4 {
	return Int32x4{
		data: x.data.AndNot(y.data),
	}
}

// Or performs a bitwise OR: x | y.
func (x Int32x4) Or(y Int32x4) Int32x4 {
	return Int32x4{
		data: x.data.Or(y.data),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Int32x4) Xor(y Int32x4) Int32x4 {
	return Int32x4{
		data: x.data.Xor(y.data),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Int32x4) Not() Int32x4 {
	return Int32x4{
		data: x.data.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Int32x4) ShiftLeft(count uint) Int32x4 {
	return Int32x4{
		data: x.data.ShiftAllLeft(uint64(count)),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (arithmetic).
func (x Int32x4) ShiftRight(count uint) Int32x4 {
	return Int32x4{
		data: x.data.ShiftAllRight(uint64(count)),
	}
}
