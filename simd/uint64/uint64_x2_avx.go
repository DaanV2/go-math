//go:build simd_avx256 || simd_avx512

package simduint64

import "simd/archsimd"

type Uint64x2 struct {
	data archsimd.Uint64x2
}

func NewUint64x2(data []uint64) (result Uint64x2) {
	result.data, _ = archsimd.LoadUint64x2Part(data)

	return result
}

// NewUint64x2Boardcast returns a Uint64x2 with every lane set to value.
func NewUint64x2Boardcast(value uint64) Uint64x2 {
	return Uint64x2{
		data: archsimd.BroadcastUint64x2(value),
	}
}

func (x Uint64x2) Store(receiver []uint64) {
	if len(receiver) >= uint64_x2_len {
		x.data.Store(receiver[:uint64_x2_len])
	} else {
		x.data.StorePart(receiver)
	}
}

// Add performs a fused: x + y.
func (x Uint64x2) Add(y Uint64x2) Uint64x2 {
	return Uint64x2{
		data: x.data.Add(y.data),
	}
}

// Sub performs a fused: x - y.
func (x Uint64x2) Sub(y Uint64x2) Uint64x2 {
	return Uint64x2{
		data: x.data.Sub(y.data),
	}
}

// And performs a bitwise AND: x & y.
func (x Uint64x2) And(y Uint64x2) Uint64x2 {
	return Uint64x2{
		data: x.data.And(y.data),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Uint64x2) AndNot(y Uint64x2) Uint64x2 {
	return Uint64x2{
		data: x.data.AndNot(y.data),
	}
}

// Or performs a bitwise OR: x | y.
func (x Uint64x2) Or(y Uint64x2) Uint64x2 {
	return Uint64x2{
		data: x.data.Or(y.data),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Uint64x2) Xor(y Uint64x2) Uint64x2 {
	return Uint64x2{
		data: x.data.Xor(y.data),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Uint64x2) Not() Uint64x2 {
	return Uint64x2{
		data: x.data.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Uint64x2) ShiftLeft(count uint) Uint64x2 {
	return Uint64x2{
		data: x.data.ShiftAllLeft(uint64(count)),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (logical).
func (x Uint64x2) ShiftRight(count uint) Uint64x2 {
	return Uint64x2{
		data: x.data.ShiftAllRight(uint64(count)),
	}
}
