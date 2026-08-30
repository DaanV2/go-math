//go:build simd_none || (!simd_avx512 && !simd_avx256)

package simdint32

type Int32x8 struct {
	data [int32_x8_len]int32
}

func NewInt32x8(data []int32) Int32x8 {
	var result Int32x8
	copy(result.data[:], data)

	return result
}

// NewInt32x8Boardcast returns an Int32x8 with every lane set to value.
func NewInt32x8Boardcast(value int32) Int32x8 {
	buf := [int32_x8_len]int32{value, value, value, value, value, value, value, value}

	return NewInt32x8(buf[:])
}

func (v Int32x8) Store(receiver []int32) {
	copy(receiver, v.data[:])
}

// Abs returns the absolute values of the elements of x
func (x Int32x8) Abs() Int32x8 {
	var result Int32x8

	for i := range x.data {
		if x.data[i] < 0 {
			result.data[i] = -x.data[i]
		} else {
			result.data[i] = x.data[i]
		}
	}

	return result
}

// Add performs a fused: x + y.
func (x Int32x8) Add(y Int32x8) Int32x8 {
	var result Int32x8

	for i := range x.data {
		result.data[i] = x.data[i] + y.data[i]
	}

	return result
}

// Mul performs a fused: x * y.
func (x Int32x8) Mul(y Int32x8) Int32x8 {
	var result Int32x8

	for i := range x.data {
		result.data[i] = x.data[i] * y.data[i]
	}

	return result
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Int32x8) Max(y Int32x8) Int32x8 {
	var result Int32x8

	for i := range x.data {
		result.data[i] = max(x.data[i], y.data[i])
	}

	return result
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Int32x8) Min(y Int32x8) Int32x8 {
	var result Int32x8

	for i := range x.data {
		result.data[i] = min(x.data[i], y.data[i])
	}

	return result
}

// Neg returns the negation of the elements of x
func (x Int32x8) Neg() Int32x8 {
	var result Int32x8

	for i := range x.data {
		result.data[i] = x.data[i] * -1
	}

	return result
}

// Sub performs a fused: x - y.
func (x Int32x8) Sub(y Int32x8) Int32x8 {
	var result Int32x8

	for i := range x.data {
		result.data[i] = x.data[i] - y.data[i]
	}

	return result
}

// And performs a bitwise AND: x & y.
func (x Int32x8) And(y Int32x8) Int32x8 {
	var result Int32x8

	for i := range x.data {
		result.data[i] = x.data[i] & y.data[i]
	}

	return result
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Int32x8) AndNot(y Int32x8) Int32x8 {
	var result Int32x8

	for i := range x.data {
		result.data[i] = x.data[i] &^ y.data[i]
	}

	return result
}

// Or performs a bitwise OR: x | y.
func (x Int32x8) Or(y Int32x8) Int32x8 {
	var result Int32x8

	for i := range x.data {
		result.data[i] = x.data[i] | y.data[i]
	}

	return result
}

// Xor performs a bitwise XOR: x ^ y.
func (x Int32x8) Xor(y Int32x8) Int32x8 {
	var result Int32x8

	for i := range x.data {
		result.data[i] = x.data[i] ^ y.data[i]
	}

	return result
}

// Not performs a bitwise NOT: ^x.
func (x Int32x8) Not() Int32x8 {
	var result Int32x8

	for i := range x.data {
		result.data[i] = ^x.data[i]
	}

	return result
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Int32x8) ShiftLeft(count uint) Int32x8 {
	var result Int32x8

	for i := range x.data {
		result.data[i] = x.data[i] << count
	}

	return result
}

// ShiftRight shifts each element of x right by count bits: x >> count (arithmetic).
func (x Int32x8) ShiftRight(count uint) Int32x8 {
	var result Int32x8

	for i := range x.data {
		result.data[i] = x.data[i] >> count
	}

	return result
}
