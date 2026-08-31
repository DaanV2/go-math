//go:build simd_none || (!simd_avx512 && !simd_avx256)

package simduint32

type Uint32x4 struct {
	data [uint32_x4_len]uint32
}

func NewUint32x4(data []uint32) (result Uint32x4) {
	copy(result.data[:], data)

	return result
}

// NewUint32x4Boardcast returns a Uint32x4 with every lane set to value.
func NewUint32x4Boardcast(value uint32) Uint32x4 {
	buf := [uint32_x4_len]uint32{value, value, value, value}

	return NewUint32x4(buf[:])
}

func (v Uint32x4) Store(receiver []uint32) {
	copy(receiver, v.data[:])
}

// Add performs a fused: x + y.
func (x Uint32x4) Add(y Uint32x4) (result Uint32x4) {
	for i := range x.data {
		result.data[i] = x.data[i] + y.data[i]
	}

	return result
}

// Mul performs a fused: x * y.
func (x Uint32x4) Mul(y Uint32x4) (result Uint32x4) {
	for i := range x.data {
		result.data[i] = x.data[i] * y.data[i]
	}

	return result
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Uint32x4) Max(y Uint32x4) (result Uint32x4) {
	for i := range x.data {
		result.data[i] = max(x.data[i], y.data[i])
	}

	return result
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Uint32x4) Min(y Uint32x4) (result Uint32x4) {
	for i := range x.data {
		result.data[i] = min(x.data[i], y.data[i])
	}

	return result
}

// Sub performs a fused: x - y.
func (x Uint32x4) Sub(y Uint32x4) (result Uint32x4) {
	for i := range x.data {
		result.data[i] = x.data[i] - y.data[i]
	}

	return result
}

// And performs a bitwise AND: x & y.
func (x Uint32x4) And(y Uint32x4) (result Uint32x4) {
	for i := range x.data {
		result.data[i] = x.data[i] & y.data[i]
	}

	return result
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Uint32x4) AndNot(y Uint32x4) (result Uint32x4) {
	for i := range x.data {
		result.data[i] = x.data[i] &^ y.data[i]
	}

	return result
}

// Or performs a bitwise OR: x | y.
func (x Uint32x4) Or(y Uint32x4) (result Uint32x4) {
	for i := range x.data {
		result.data[i] = x.data[i] | y.data[i]
	}

	return result
}

// Xor performs a bitwise XOR: x ^ y.
func (x Uint32x4) Xor(y Uint32x4) (result Uint32x4) {
	for i := range x.data {
		result.data[i] = x.data[i] ^ y.data[i]
	}

	return result
}

// Not performs a bitwise NOT: ^x.
func (x Uint32x4) Not() (result Uint32x4) {
	for i := range x.data {
		result.data[i] = ^x.data[i]
	}

	return result
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Uint32x4) ShiftLeft(count uint) (result Uint32x4) {
	for i := range x.data {
		result.data[i] = x.data[i] << count
	}

	return result
}

// ShiftRight shifts each element of x right by count bits: x >> count (logical).
func (x Uint32x4) ShiftRight(count uint) (result Uint32x4) {
	for i := range x.data {
		result.data[i] = x.data[i] >> count
	}

	return result
}
