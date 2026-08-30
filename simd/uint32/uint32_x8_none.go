//go:build simd_none || (!simd_avx512 && !simd_avx256)

package simduint32

type Uint32x8 struct {
	data [uint32_x8_len]uint32
}

func NewUint32x8(data []uint32) Uint32x8 {
	var result Uint32x8
	copy(result.data[:], data)

	return result
}

func NewUint32x8Boardcast(value uint32) Uint32x8 {
	buf := [uint32_x8_len]uint32{value, value, value, value, value, value, value, value}

	return NewUint32x8(buf[:])
}

func (v Uint32x8) Store(receiver []uint32) {
	copy(receiver, v.data[:])
}

// Add performs a fused: x + y.
func (x Uint32x8) Add(y Uint32x8) Uint32x8 {
	var result Uint32x8

	for i := range x.data {
		result.data[i] = x.data[i] + y.data[i]
	}

	return result
}

// Mul performs a fused: x * y.
func (x Uint32x8) Mul(y Uint32x8) Uint32x8 {
	var result Uint32x8

	for i := range x.data {
		result.data[i] = x.data[i] * y.data[i]
	}

	return result
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Uint32x8) Max(y Uint32x8) Uint32x8 {
	var result Uint32x8

	for i := range x.data {
		result.data[i] = max(x.data[i], y.data[i])
	}

	return result
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Uint32x8) Min(y Uint32x8) Uint32x8 {
	var result Uint32x8

	for i := range x.data {
		result.data[i] = min(x.data[i], y.data[i])
	}

	return result
}

// Sub performs a fused: x - y.
func (x Uint32x8) Sub(y Uint32x8) Uint32x8 {
	var result Uint32x8

	for i := range x.data {
		result.data[i] = x.data[i] - y.data[i]
	}

	return result
}

// And performs a bitwise AND: x & y.
func (x Uint32x8) And(y Uint32x8) Uint32x8 {
	var result Uint32x8

	for i := range x.data {
		result.data[i] = x.data[i] & y.data[i]
	}

	return result
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Uint32x8) AndNot(y Uint32x8) Uint32x8 {
	var result Uint32x8

	for i := range x.data {
		result.data[i] = x.data[i] &^ y.data[i]
	}

	return result
}

// Or performs a bitwise OR: x | y.
func (x Uint32x8) Or(y Uint32x8) Uint32x8 {
	var result Uint32x8

	for i := range x.data {
		result.data[i] = x.data[i] | y.data[i]
	}

	return result
}

// Xor performs a bitwise XOR: x ^ y.
func (x Uint32x8) Xor(y Uint32x8) Uint32x8 {
	var result Uint32x8

	for i := range x.data {
		result.data[i] = x.data[i] ^ y.data[i]
	}

	return result
}

// Not performs a bitwise NOT: ^x.
func (x Uint32x8) Not() Uint32x8 {
	var result Uint32x8

	for i := range x.data {
		result.data[i] = ^x.data[i]
	}

	return result
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Uint32x8) ShiftLeft(count uint) Uint32x8 {
	var result Uint32x8

	for i := range x.data {
		result.data[i] = x.data[i] << count
	}

	return result
}

// ShiftRight shifts each element of x right by count bits: x >> count (logical).
func (x Uint32x8) ShiftRight(count uint) Uint32x8 {
	var result Uint32x8

	for i := range x.data {
		result.data[i] = x.data[i] >> count
	}

	return result
}
