//go:build simd_avx256 || simd_avx512

package simduint32

import "simd/archsimd"

type Uint32x8 struct {
	data archsimd.Uint32x8
}

func NewUint32x8(data []uint32) (result Uint32x8) {
	result.data, _ = archsimd.LoadUint32x8Part(data)

	return result
}

// NewUint32x8Boardcast returns a Uint32x8 with every lane set to value.
func NewUint32x8Boardcast(value uint32) Uint32x8 {
	return Uint32x8{
		data: archsimd.BroadcastUint32x8(value),
	}
}

func (x Uint32x8) Store(receiver []uint32) {
	if len(receiver) >= uint32_x8_len {
		x.data.Store(receiver[:uint32_x8_len])
	} else {
		x.data.StorePart(receiver)
	}
}

// Add performs a fused: x + y.
func (x Uint32x8) Add(y Uint32x8) Uint32x8 {
	return Uint32x8{
		data: x.data.Add(y.data),
	}
}

// Mul performs a fused: x * y.
func (x Uint32x8) Mul(y Uint32x8) Uint32x8 {
	return Uint32x8{
		data: x.data.Mul(y.data),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Uint32x8) Max(y Uint32x8) Uint32x8 {
	return Uint32x8{
		data: x.data.Max(y.data),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Uint32x8) Min(y Uint32x8) Uint32x8 {
	return Uint32x8{
		data: x.data.Min(y.data),
	}
}

// Sub performs a fused: x - y.
func (x Uint32x8) Sub(y Uint32x8) Uint32x8 {
	return Uint32x8{
		data: x.data.Sub(y.data),
	}
}

// And performs a bitwise AND: x & y.
func (x Uint32x8) And(y Uint32x8) Uint32x8 {
	return Uint32x8{
		data: x.data.And(y.data),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Uint32x8) AndNot(y Uint32x8) Uint32x8 {
	return Uint32x8{
		data: x.data.AndNot(y.data),
	}
}

// Or performs a bitwise OR: x | y.
func (x Uint32x8) Or(y Uint32x8) Uint32x8 {
	return Uint32x8{
		data: x.data.Or(y.data),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Uint32x8) Xor(y Uint32x8) Uint32x8 {
	return Uint32x8{
		data: x.data.Xor(y.data),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Uint32x8) Not() Uint32x8 {
	return Uint32x8{
		data: x.data.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Uint32x8) ShiftLeft(count uint) Uint32x8 {
	return Uint32x8{
		data: x.data.ShiftAllLeft(uint64(count)),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (logical).
func (x Uint32x8) ShiftRight(count uint) Uint32x8 {
	return Uint32x8{
		data: x.data.ShiftAllRight(uint64(count)),
	}
}
