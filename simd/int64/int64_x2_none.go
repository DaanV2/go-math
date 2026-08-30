//go:build simd_none || (!simd_avx512 && !simd_avx256)

package simdint64

type Int64x2 struct {
	data [int64_x2_len]int64
}

func NewInt64x2(data []int64) Int64x2 {
	var result Int64x2
	copy(result.data[:], data)

	return result
}

// NewInt64x2Boardcast returns an Int64x2 with every lane set to value.
func NewInt64x2Boardcast(value int64) Int64x2 {
	buf := [int64_x2_len]int64{value, value}

	return NewInt64x2(buf[:])
}

func (v Int64x2) Store(receiver []int64) {
	copy(receiver, v.data[:])
}

// Abs returns the absolute values of the elements of x
func (x Int64x2) Abs() Int64x2 {
	var result Int64x2

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
func (x Int64x2) Add(y Int64x2) Int64x2 {
	var result Int64x2

	for i := range x.data {
		result.data[i] = x.data[i] + y.data[i]
	}

	return result
}

// Mul performs a fused: x * y.
func (x Int64x2) Mul(y Int64x2) Int64x2 {
	var result Int64x2

	for i := range x.data {
		result.data[i] = x.data[i] * y.data[i]
	}

	return result
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Int64x2) Max(y Int64x2) Int64x2 {
	var result Int64x2

	for i := range x.data {
		result.data[i] = max(x.data[i], y.data[i])
	}

	return result
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Int64x2) Min(y Int64x2) Int64x2 {
	var result Int64x2

	for i := range x.data {
		result.data[i] = min(x.data[i], y.data[i])
	}

	return result
}

// Neg returns the negation of the elements of x
func (x Int64x2) Neg() Int64x2 {
	var result Int64x2

	for i := range x.data {
		result.data[i] = x.data[i] * -1
	}

	return result
}

// Sub performs a fused: x - y.
func (x Int64x2) Sub(y Int64x2) Int64x2 {
	var result Int64x2

	for i := range x.data {
		result.data[i] = x.data[i] - y.data[i]
	}

	return result
}

// And performs a bitwise AND: x & y.
func (x Int64x2) And(y Int64x2) Int64x2 {
	var result Int64x2

	for i := range x.data {
		result.data[i] = x.data[i] & y.data[i]
	}

	return result
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Int64x2) AndNot(y Int64x2) Int64x2 {
	var result Int64x2

	for i := range x.data {
		result.data[i] = x.data[i] &^ y.data[i]
	}

	return result
}

// Or performs a bitwise OR: x | y.
func (x Int64x2) Or(y Int64x2) Int64x2 {
	var result Int64x2

	for i := range x.data {
		result.data[i] = x.data[i] | y.data[i]
	}

	return result
}

// Xor performs a bitwise XOR: x ^ y.
func (x Int64x2) Xor(y Int64x2) Int64x2 {
	var result Int64x2

	for i := range x.data {
		result.data[i] = x.data[i] ^ y.data[i]
	}

	return result
}

// Not performs a bitwise NOT: ^x.
func (x Int64x2) Not() Int64x2 {
	var result Int64x2

	for i := range x.data {
		result.data[i] = ^x.data[i]
	}

	return result
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Int64x2) ShiftLeft(count uint) Int64x2 {
	var result Int64x2

	for i := range x.data {
		result.data[i] = x.data[i] << count
	}

	return result
}

// ShiftRight shifts each element of x right by count bits: x >> count (arithmetic).
func (x Int64x2) ShiftRight(count uint) Int64x2 {
	var result Int64x2

	for i := range x.data {
		result.data[i] = x.data[i] >> count
	}

	return result
}
