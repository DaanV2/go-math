//go:build simd_none || (!simd_avx512 && !simd_avx256)

package simduint32

type Uint32x16 struct {
	data [uint32_x16_len]uint32
}

func NewUint32x16(data []uint32) Uint32x16 {
	var result Uint32x16
	copy(result.data[:], data)

	return result
}

func (v Uint32x16) Store(receiver []uint32) {
	copy(receiver, v.data[:])
}

// Add performs a fused: x + y.
func (x Uint32x16) Add(y Uint32x16) Uint32x16 {
	var result Uint32x16

	for i := range x.data {
		result.data[i] = x.data[i] + y.data[i]
	}

	return result
}

// Mul performs a fused: x * y.
func (x Uint32x16) Mul(y Uint32x16) Uint32x16 {
	var result Uint32x16

	for i := range x.data {
		result.data[i] = x.data[i] * y.data[i]
	}

	return result
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Uint32x16) Max(y Uint32x16) Uint32x16 {
	var result Uint32x16

	for i := range x.data {
		result.data[i] = max(x.data[i], y.data[i])
	}

	return result
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Uint32x16) Min(y Uint32x16) Uint32x16 {
	var result Uint32x16

	for i := range x.data {
		result.data[i] = min(x.data[i], y.data[i])
	}

	return result
}

// Sub performs a fused: x - y.
func (x Uint32x16) Sub(y Uint32x16) Uint32x16 {
	var result Uint32x16

	for i := range x.data {
		result.data[i] = x.data[i] - y.data[i]
	}

	return result
}

// And performs a bitwise AND: x & y.
func (x Uint32x16) And(y Uint32x16) Uint32x16 {
	var result Uint32x16

	for i := range x.data {
		result.data[i] = x.data[i] & y.data[i]
	}

	return result
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Uint32x16) AndNot(y Uint32x16) Uint32x16 {
	var result Uint32x16

	for i := range x.data {
		result.data[i] = x.data[i] &^ y.data[i]
	}

	return result
}

// Or performs a bitwise OR: x | y.
func (x Uint32x16) Or(y Uint32x16) Uint32x16 {
	var result Uint32x16

	for i := range x.data {
		result.data[i] = x.data[i] | y.data[i]
	}

	return result
}

// Xor performs a bitwise XOR: x ^ y.
func (x Uint32x16) Xor(y Uint32x16) Uint32x16 {
	var result Uint32x16

	for i := range x.data {
		result.data[i] = x.data[i] ^ y.data[i]
	}

	return result
}

// Not performs a bitwise NOT: ^x.
func (x Uint32x16) Not() Uint32x16 {
	var result Uint32x16

	for i := range x.data {
		result.data[i] = ^x.data[i]
	}

	return result
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Uint32x16) ShiftLeft(count uint) Uint32x16 {
	var result Uint32x16

	for i := range x.data {
		result.data[i] = x.data[i] << count
	}

	return result
}

// ShiftRight shifts each element of x right by count bits: x >> count (logical).
func (x Uint32x16) ShiftRight(count uint) Uint32x16 {
	var result Uint32x16

	for i := range x.data {
		result.data[i] = x.data[i] >> count
	}

	return result
}
