//go:build simd_avx256 || simd_avx512

package simdint64

import "simd/archsimd"

type Int64x2 struct {
	data archsimd.Int64x2
}

func NewInt64x2(data []int64) Int64x2 {
	var result Int64x2

	result.data, _ = archsimd.LoadInt64x2Part(data)

	return result
}

func (x Int64x2) Store(receiver []int64) {
	if len(receiver) >= int64_x2_len {
		x.data.Store(receiver[:int64_x2_len])
	} else {
		x.data.StorePart(receiver)
	}
}

// Add performs a fused: x + y.
func (x Int64x2) Add(y Int64x2) Int64x2 {
	return Int64x2{
		data: x.data.Add(y.data),
	}
}

// Neg returns the negation of the elements of x
func (x Int64x2) Neg() Int64x2 {
	return Int64x2{
		data: x.data.Neg(),
	}
}

// Sub performs a fused: x - y.
func (x Int64x2) Sub(y Int64x2) Int64x2 {
	return Int64x2{
		data: x.data.Sub(y.data),
	}
}

// And performs a bitwise AND: x & y.
func (x Int64x2) And(y Int64x2) Int64x2 {
	return Int64x2{
		data: x.data.And(y.data),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Int64x2) AndNot(y Int64x2) Int64x2 {
	return Int64x2{
		data: x.data.AndNot(y.data),
	}
}

// Or performs a bitwise OR: x | y.
func (x Int64x2) Or(y Int64x2) Int64x2 {
	return Int64x2{
		data: x.data.Or(y.data),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Int64x2) Xor(y Int64x2) Int64x2 {
	return Int64x2{
		data: x.data.Xor(y.data),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Int64x2) Not() Int64x2 {
	return Int64x2{
		data: x.data.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Int64x2) ShiftLeft(count uint) Int64x2 {
	return Int64x2{
		data: x.data.ShiftAllLeft(uint64(count)),
	}
}
