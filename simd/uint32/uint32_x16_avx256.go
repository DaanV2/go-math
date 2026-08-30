//go:build simd_avx256

package simduint32

import "simd/archsimd"

type Uint32x16 struct {
	data0 archsimd.Uint32x8
	data1 archsimd.Uint32x8
}

func NewUint32x16(data []uint32) Uint32x16 {
	var result Uint32x16

	var n int
	result.data0, n = archsimd.LoadUint32x8Part(data)
	if n == 8 { // Read atleast 8 points, so there should be more
		result.data1, _ = archsimd.LoadUint32x8Part(data[(uint32_x16_len / 2):])
	}

	return result
}

func (x Uint32x16) Store(receiver []uint32) {
	switch {
	case len(receiver) == uint32_x16_len:
		x.data0.Store(receiver[:(uint32_x16_len / 2)])
		x.data1.Store(receiver[(uint32_x16_len / 2):])
	case len(receiver) > (uint32_x16_len / 2):
		x.data0.Store(receiver[:(uint32_x16_len / 2)])
		_ = x.data1.StorePart(receiver[(uint32_x16_len / 2):])
	default:
		_ = x.data0.StorePart(receiver)
	}
}

// Add performs a fused: x + y.
func (x Uint32x16) Add(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data0: x.data0.Add(y.data0),
		data1: x.data1.Add(y.data1),
	}
}

// Mul performs a fused: x * y.
func (x Uint32x16) Mul(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data0: x.data0.Mul(y.data0),
		data1: x.data1.Mul(y.data1),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Uint32x16) Max(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data0: x.data0.Max(y.data0),
		data1: x.data1.Max(y.data1),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Uint32x16) Min(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data0: x.data0.Min(y.data0),
		data1: x.data1.Min(y.data1),
	}
}

// Sub performs a fused: x - y.
func (x Uint32x16) Sub(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data0: x.data0.Sub(y.data0),
		data1: x.data1.Sub(y.data1),
	}
}

// And performs a bitwise AND: x & y.
func (x Uint32x16) And(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data0: x.data0.And(y.data0),
		data1: x.data1.And(y.data1),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Uint32x16) AndNot(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data0: x.data0.AndNot(y.data0),
		data1: x.data1.AndNot(y.data1),
	}
}

// Or performs a bitwise OR: x | y.
func (x Uint32x16) Or(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data0: x.data0.Or(y.data0),
		data1: x.data1.Or(y.data1),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Uint32x16) Xor(y Uint32x16) Uint32x16 {
	return Uint32x16{
		data0: x.data0.Xor(y.data0),
		data1: x.data1.Xor(y.data1),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Uint32x16) Not() Uint32x16 {
	return Uint32x16{
		data0: x.data0.Not(),
		data1: x.data1.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Uint32x16) ShiftLeft(count uint) Uint32x16 {
	return Uint32x16{
		data0: x.data0.ShiftAllLeft(uint64(count)),
		data1: x.data1.ShiftAllLeft(uint64(count)),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (logical).
func (x Uint32x16) ShiftRight(count uint) Uint32x16 {
	return Uint32x16{
		data0: x.data0.ShiftAllRight(uint64(count)),
		data1: x.data1.ShiftAllRight(uint64(count)),
	}
}
