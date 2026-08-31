//go:build simd_avx512

package simdint32

import "simd/archsimd"

type Int32x16 struct {
	data archsimd.Int32x16
}

func NewInt32x16(data []int32) (result Int32x16) {
	result.data, _ = archsimd.LoadInt32x16Part(data)

	return result
}

// NewInt32x16Boardcast returns an Int32x16 with every lane set to value.
func NewInt32x16Boardcast(value int32) Int32x16 {
	return Int32x16{
		data: archsimd.BroadcastInt32x16(value),
	}
}

func (v Int32x16) Store(receiver []int32) {
	v.data.StorePart(receiver)
}

// Abs returns the absolute values of the elements of x
func (x Int32x16) Abs() Int32x16 {
	return Int32x16{
		data: x.data.Abs(),
	}
}

// Add performs a fused: x + y.
func (x Int32x16) Add(y Int32x16) Int32x16 {
	return Int32x16{
		data: x.data.Add(y.data),
	}
}

// Mul performs a fused: x * y.
func (x Int32x16) Mul(y Int32x16) Int32x16 {
	return Int32x16{
		data: x.data.Mul(y.data),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Int32x16) Max(y Int32x16) Int32x16 {
	return Int32x16{
		data: x.data.Max(y.data),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Int32x16) Min(y Int32x16) Int32x16 {
	return Int32x16{
		data: x.data.Min(y.data),
	}
}

// Neg returns the negation of the elements of x
func (x Int32x16) Neg() Int32x16 {
	return Int32x16{
		data: x.data.Neg(),
	}
}

// Sub performs a fused: x - y.
func (x Int32x16) Sub(y Int32x16) Int32x16 {
	return Int32x16{
		data: x.data.Sub(y.data),
	}
}

// And performs a bitwise AND: x & y.
func (x Int32x16) And(y Int32x16) Int32x16 {
	return Int32x16{
		data: x.data.And(y.data),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Int32x16) AndNot(y Int32x16) Int32x16 {
	return Int32x16{
		data: x.data.AndNot(y.data),
	}
}

// Or performs a bitwise OR: x | y.
func (x Int32x16) Or(y Int32x16) Int32x16 {
	return Int32x16{
		data: x.data.Or(y.data),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Int32x16) Xor(y Int32x16) Int32x16 {
	return Int32x16{
		data: x.data.Xor(y.data),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Int32x16) Not() Int32x16 {
	return Int32x16{
		data: x.data.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Int32x16) ShiftLeft(count uint) Int32x16 {
	return Int32x16{
		data: x.data.ShiftAllLeft(uint64(count)),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (arithmetic).
func (x Int32x16) ShiftRight(count uint) Int32x16 {
	return Int32x16{
		data: x.data.ShiftAllRight(uint64(count)),
	}
}
