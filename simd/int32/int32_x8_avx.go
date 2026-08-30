//go:build simd_avx256 || simd_avx512

package simdint32

import "simd/archsimd"

type Int32x8 struct {
	data archsimd.Int32x8
}

func NewInt32x8(data []int32) Int32x8 {
	var result Int32x8

	result.data, _ = archsimd.LoadInt32x8Part(data)

	return result
}

func (x Int32x8) Store(receiver []int32) {
	if len(receiver) >= int32_x8_len {
		x.data.Store(receiver[:int32_x8_len])
	} else {
		x.data.StorePart(receiver)
	}
}

// Abs returns the absolute values of the elements of x
func (x Int32x8) Abs() Int32x8 {
	return Int32x8{
		data: x.data.Abs(),
	}
}

// Add performs a fused: x + y.
func (x Int32x8) Add(y Int32x8) Int32x8 {
	return Int32x8{
		data: x.data.Add(y.data),
	}
}

// Mul performs a fused: x * y.
func (x Int32x8) Mul(y Int32x8) Int32x8 {
	return Int32x8{
		data: x.data.Mul(y.data),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Int32x8) Max(y Int32x8) Int32x8 {
	return Int32x8{
		data: x.data.Max(y.data),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Int32x8) Min(y Int32x8) Int32x8 {
	return Int32x8{
		data: x.data.Min(y.data),
	}
}

// Neg returns the negation of the elements of x
func (x Int32x8) Neg() Int32x8 {
	return Int32x8{
		data: x.data.Neg(),
	}
}

// Sub performs a fused: x - y.
func (x Int32x8) Sub(y Int32x8) Int32x8 {
	return Int32x8{
		data: x.data.Sub(y.data),
	}
}

// And performs a bitwise AND: x & y.
func (x Int32x8) And(y Int32x8) Int32x8 {
	return Int32x8{
		data: x.data.And(y.data),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Int32x8) AndNot(y Int32x8) Int32x8 {
	return Int32x8{
		data: x.data.AndNot(y.data),
	}
}

// Or performs a bitwise OR: x | y.
func (x Int32x8) Or(y Int32x8) Int32x8 {
	return Int32x8{
		data: x.data.Or(y.data),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Int32x8) Xor(y Int32x8) Int32x8 {
	return Int32x8{
		data: x.data.Xor(y.data),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Int32x8) Not() Int32x8 {
	return Int32x8{
		data: x.data.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Int32x8) ShiftLeft(count uint) Int32x8 {
	return Int32x8{
		data: x.data.ShiftAllLeft(uint64(count)),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (arithmetic).
func (x Int32x8) ShiftRight(count uint) Int32x8 {
	return Int32x8{
		data: x.data.ShiftAllRight(uint64(count)),
	}
}
