//go:build simd_avx512

package simdint64

import "simd/archsimd"

type Int64x8 struct {
	data archsimd.Int64x8
}

func NewInt64x8(data []int64) Int64x8 {
	var result Int64x8

	result.data, _ = archsimd.LoadInt64x8Part(data)

	return result
}

func (v Int64x8) Store(receiver []int64) {
	v.data.StorePart(receiver)
}

// Abs returns the absolute values of the elements of x
func (x Int64x8) Abs() Int64x8 {
	return Int64x8{
		data: x.data.Abs(),
	}
}

// Add performs a fused: x + y.
func (x Int64x8) Add(y Int64x8) Int64x8 {
	return Int64x8{
		data: x.data.Add(y.data),
	}
}

// Mul performs a fused: x * y.
func (x Int64x8) Mul(y Int64x8) Int64x8 {
	return Int64x8{
		data: x.data.Mul(y.data),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Int64x8) Max(y Int64x8) Int64x8 {
	return Int64x8{
		data: x.data.Max(y.data),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Int64x8) Min(y Int64x8) Int64x8 {
	return Int64x8{
		data: x.data.Min(y.data),
	}
}

// Neg returns the negation of the elements of x
func (x Int64x8) Neg() Int64x8 {
	return Int64x8{
		data: x.data.Neg(),
	}
}

// Sub performs a fused: x - y.
func (x Int64x8) Sub(y Int64x8) Int64x8 {
	return Int64x8{
		data: x.data.Sub(y.data),
	}
}

// And performs a bitwise AND: x & y.
func (x Int64x8) And(y Int64x8) Int64x8 {
	return Int64x8{
		data: x.data.And(y.data),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Int64x8) AndNot(y Int64x8) Int64x8 {
	return Int64x8{
		data: x.data.AndNot(y.data),
	}
}

// Or performs a bitwise OR: x | y.
func (x Int64x8) Or(y Int64x8) Int64x8 {
	return Int64x8{
		data: x.data.Or(y.data),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Int64x8) Xor(y Int64x8) Int64x8 {
	return Int64x8{
		data: x.data.Xor(y.data),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Int64x8) Not() Int64x8 {
	return Int64x8{
		data: x.data.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Int64x8) ShiftLeft(count uint) Int64x8 {
	return Int64x8{
		data: x.data.ShiftAllLeft(uint64(count)),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (arithmetic).
func (x Int64x8) ShiftRight(count uint) Int64x8 {
	return Int64x8{
		data: x.data.ShiftAllRight(uint64(count)),
	}
}
