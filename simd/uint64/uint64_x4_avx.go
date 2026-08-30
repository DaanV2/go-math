//go:build simd_avx256 || simd_avx512

package simduint64

import "simd/archsimd"

type Uint64x4 struct {
	data archsimd.Uint64x4
}

func NewUint64x4(data []uint64) Uint64x4 {
	var result Uint64x4

	result.data, _ = archsimd.LoadUint64x4Part(data)

	return result
}

// NewUint64x4Boardcast returns a Uint64x4 with every lane set to value.
func NewUint64x4Boardcast(value uint64) Uint64x4 {
	return Uint64x4{
		data: archsimd.BroadcastUint64x4(value),
	}
}

func (x Uint64x4) Store(receiver []uint64) {
	if len(receiver) >= uint64_x4_len {
		x.data.Store(receiver[:uint64_x4_len])
	} else {
		x.data.StorePart(receiver)
	}
}

// Add performs a fused: x + y.
func (x Uint64x4) Add(y Uint64x4) Uint64x4 {
	return Uint64x4{
		data: x.data.Add(y.data),
	}
}

// Sub performs a fused: x - y.
func (x Uint64x4) Sub(y Uint64x4) Uint64x4 {
	return Uint64x4{
		data: x.data.Sub(y.data),
	}
}

// And performs a bitwise AND: x & y.
func (x Uint64x4) And(y Uint64x4) Uint64x4 {
	return Uint64x4{
		data: x.data.And(y.data),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Uint64x4) AndNot(y Uint64x4) Uint64x4 {
	return Uint64x4{
		data: x.data.AndNot(y.data),
	}
}

// Or performs a bitwise OR: x | y.
func (x Uint64x4) Or(y Uint64x4) Uint64x4 {
	return Uint64x4{
		data: x.data.Or(y.data),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Uint64x4) Xor(y Uint64x4) Uint64x4 {
	return Uint64x4{
		data: x.data.Xor(y.data),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Uint64x4) Not() Uint64x4 {
	return Uint64x4{
		data: x.data.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Uint64x4) ShiftLeft(count uint) Uint64x4 {
	return Uint64x4{
		data: x.data.ShiftAllLeft(uint64(count)),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (logical).
func (x Uint64x4) ShiftRight(count uint) Uint64x4 {
	return Uint64x4{
		data: x.data.ShiftAllRight(uint64(count)),
	}
}
