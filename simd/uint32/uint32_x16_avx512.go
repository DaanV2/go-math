//go:build simd_avx512

package simduint32

import "simd/archsimd"

type Uint32x16 struct {
	data archsimd.Uint32x16
}

func NewUint32x16(data []uint32) (result Uint32x16) {
	result.data, _ = archsimd.LoadUint32x16Part(data)

	return result
}

// NewUint32x16Boardcast returns a Uint32x16 with every lane set to value.
func NewUint32x16Boardcast(value uint32) Uint32x16 {
	return Uint32x16{
		data: archsimd.BroadcastUint32x16(value),
	}
}

func (v Uint32x16) Store(receiver []uint32) {
	v.data.StorePart(receiver)
}

// Add performs a fused: x + y.
func (x Uint32x16) Add(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data: x.data.Add(y.data),
	}
}

// Mul performs a fused: x * y.
func (x Uint32x16) Mul(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data: x.data.Mul(y.data),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Uint32x16) Max(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data: x.data.Max(y.data),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Uint32x16) Min(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data: x.data.Min(y.data),
	}
}

// Sub performs a fused: x - y.
func (x Uint32x16) Sub(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data: x.data.Sub(y.data),
	}
}

// And performs a bitwise AND: x & y.
func (x Uint32x16) And(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data: x.data.And(y.data),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Uint32x16) AndNot(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data: x.data.AndNot(y.data),
	}
}

// Or performs a bitwise OR: x | y.
func (x Uint32x16) Or(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data: x.data.Or(y.data),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Uint32x16) Xor(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data: x.data.Xor(y.data),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Uint32x16) Not() Uint32x16 {
	return Uint32x16{
		data: x.data.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Uint32x16) ShiftLeft(count uint) Uint32x16 {
	return Uint32x16{
		data: x.data.ShiftAllLeft(uint64(count)),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (logical).
func (x Uint32x16) ShiftRight(count uint) Uint32x16 {
	return Uint32x16{
		data: x.data.ShiftAllRight(uint64(count)),
	}
}
