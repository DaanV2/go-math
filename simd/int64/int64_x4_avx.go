//go:build simd_avx256 || simd_avx512

package simdint64

import "simd/archsimd"

type Int64x4 struct {
	data archsimd.Int64x4
}

func NewInt64x4(data []int64) (result Int64x4) {
	result.data, _ = archsimd.LoadInt64x4Part(data)

	return result
}

// NewInt64x4Boardcast returns an Int64x4 with every lane set to value.
func NewInt64x4Boardcast(value int64) Int64x4 {
	return Int64x4{
		data: archsimd.BroadcastInt64x4(value),
	}
}

func (x Int64x4) Store(receiver []int64) {
	if len(receiver) >= int64_x4_len {
		x.data.Store(receiver[:int64_x4_len])
	} else {
		x.data.StorePart(receiver)
	}
}

// Add performs a fused: x + y.
func (x Int64x4) Add(y Int64x4) Int64x4 {
	return Int64x4{
		data: x.data.Add(y.data),
	}
}

// Neg returns the negation of the elements of x
func (x Int64x4) Neg() Int64x4 {
	return Int64x4{
		data: x.data.Neg(),
	}
}

// Sub performs a fused: x - y.
func (x Int64x4) Sub(y Int64x4) Int64x4 {
	return Int64x4{
		data: x.data.Sub(y.data),
	}
}

// And performs a bitwise AND: x & y.
func (x Int64x4) And(y Int64x4) Int64x4 {
	return Int64x4{
		data: x.data.And(y.data),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Int64x4) AndNot(y Int64x4) Int64x4 {
	return Int64x4{
		data: x.data.AndNot(y.data),
	}
}

// Or performs a bitwise OR: x | y.
func (x Int64x4) Or(y Int64x4) Int64x4 {
	return Int64x4{
		data: x.data.Or(y.data),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Int64x4) Xor(y Int64x4) Int64x4 {
	return Int64x4{
		data: x.data.Xor(y.data),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Int64x4) Not() Int64x4 {
	return Int64x4{
		data: x.data.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Int64x4) ShiftLeft(count uint) Int64x4 {
	return Int64x4{
		data: x.data.ShiftAllLeft(uint64(count)),
	}
}
