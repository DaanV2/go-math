//go:build simd_avx512

package simduint64

import "simd/archsimd"

type Uint64x8 struct {
	data archsimd.Uint64x8
}

func NewUint64x8(data []uint64) Uint64x8 {
	var result Uint64x8

	result.data, _ = archsimd.LoadUint64x8Part(data)

	return result
}

func NewUint64x8Boardcast(value uint64) Uint64x8 {
	return Uint64x8{
		data: archsimd.BroadcastUint64x8(value),
	}
}

func (v Uint64x8) Store(receiver []uint64) {
	v.data.StorePart(receiver)
}

// Add performs a fused: x + y.
func (x Uint64x8) Add(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data: x.data.Add(y.data),
	}
}

// Mul performs a fused: x * y.
func (x Uint64x8) Mul(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data: x.data.Mul(y.data),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Uint64x8) Max(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data: x.data.Max(y.data),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Uint64x8) Min(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data: x.data.Min(y.data),
	}
}

// Sub performs a fused: x - y.
func (x Uint64x8) Sub(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data: x.data.Sub(y.data),
	}
}

// And performs a bitwise AND: x & y.
func (x Uint64x8) And(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data: x.data.And(y.data),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Uint64x8) AndNot(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data: x.data.AndNot(y.data),
	}
}

// Or performs a bitwise OR: x | y.
func (x Uint64x8) Or(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data: x.data.Or(y.data),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Uint64x8) Xor(y Uint64x8) Uint64x8 {
	return Uint64x8{
		data: x.data.Xor(y.data),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Uint64x8) Not() Uint64x8 {
	return Uint64x8{
		data: x.data.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Uint64x8) ShiftLeft(count uint) Uint64x8 {
	return Uint64x8{
		data: x.data.ShiftAllLeft(uint64(count)),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (logical).
func (x Uint64x8) ShiftRight(count uint) Uint64x8 {
	return Uint64x8{
		data: x.data.ShiftAllRight(uint64(count)),
	}
}
