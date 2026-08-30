//go:build simd_none || (!simd_avx512 && !simd_avx256)

package simduint64

type Uint64x4 struct {
	data [uint64_x4_len]uint64
}

func NewUint64x4(data []uint64) Uint64x4 {
	var result Uint64x4
	copy(result.data[:], data)

	return result
}

// NewUint64x4Boardcast returns a Uint64x4 with every lane set to value.
func NewUint64x4Boardcast(value uint64) Uint64x4 {
	buf := [uint64_x4_len]uint64{value, value, value, value}

	return NewUint64x4(buf[:])
}

func (v Uint64x4) Store(receiver []uint64) {
	copy(receiver, v.data[:])
}

// Add performs a fused: x + y.
func (x Uint64x4) Add(y Uint64x4) Uint64x4 {
	var result Uint64x4

	for i := range x.data {
		result.data[i] = x.data[i] + y.data[i]
	}

	return result
}

// Mul performs a fused: x * y.
func (x Uint64x4) Mul(y Uint64x4) Uint64x4 {
	var result Uint64x4

	for i := range x.data {
		result.data[i] = x.data[i] * y.data[i]
	}

	return result
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Uint64x4) Max(y Uint64x4) Uint64x4 {
	var result Uint64x4

	for i := range x.data {
		result.data[i] = max(x.data[i], y.data[i])
	}

	return result
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Uint64x4) Min(y Uint64x4) Uint64x4 {
	var result Uint64x4

	for i := range x.data {
		result.data[i] = min(x.data[i], y.data[i])
	}

	return result
}

// Sub performs a fused: x - y.
func (x Uint64x4) Sub(y Uint64x4) Uint64x4 {
	var result Uint64x4

	for i := range x.data {
		result.data[i] = x.data[i] - y.data[i]
	}

	return result
}

// And performs a bitwise AND: x & y.
func (x Uint64x4) And(y Uint64x4) Uint64x4 {
	var result Uint64x4

	for i := range x.data {
		result.data[i] = x.data[i] & y.data[i]
	}

	return result
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Uint64x4) AndNot(y Uint64x4) Uint64x4 {
	var result Uint64x4

	for i := range x.data {
		result.data[i] = x.data[i] &^ y.data[i]
	}

	return result
}

// Or performs a bitwise OR: x | y.
func (x Uint64x4) Or(y Uint64x4) Uint64x4 {
	var result Uint64x4

	for i := range x.data {
		result.data[i] = x.data[i] | y.data[i]
	}

	return result
}

// Xor performs a bitwise XOR: x ^ y.
func (x Uint64x4) Xor(y Uint64x4) Uint64x4 {
	var result Uint64x4

	for i := range x.data {
		result.data[i] = x.data[i] ^ y.data[i]
	}

	return result
}

// Not performs a bitwise NOT: ^x.
func (x Uint64x4) Not() Uint64x4 {
	var result Uint64x4

	for i := range x.data {
		result.data[i] = ^x.data[i]
	}

	return result
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Uint64x4) ShiftLeft(count uint) Uint64x4 {
	var result Uint64x4

	for i := range x.data {
		result.data[i] = x.data[i] << count
	}

	return result
}

// ShiftRight shifts each element of x right by count bits: x >> count (logical).
func (x Uint64x4) ShiftRight(count uint) Uint64x4 {
	var result Uint64x4

	for i := range x.data {
		result.data[i] = x.data[i] >> count
	}

	return result
}
