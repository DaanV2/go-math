//go:build simd_none || (!simd_avx512 && !simd_avx256)

package simdint64

type Int64x8 struct {
	data [int64_x8_len]int64
}

func NewInt64x8(data []int64) (result Int64x8) {

	copy(result.data[:], data)

	return
}

// NewInt64x8Boardcast returns an Int64x8 with every lane set to value.
func NewInt64x8Boardcast(value int64) Int64x8 {
	buf := [int64_x8_len]int64{value, value, value, value, value, value, value, value}

	return NewInt64x8(buf[:])
}

func (v Int64x8) Store(receiver []int64) {
	copy(receiver, v.data[:])
}

// Abs returns the absolute values of the elements of x
func (x Int64x8) Abs() (result Int64x8) {
	for i := range &x.data {
		if x.data[i] < 0 {
			result.data[i] = -x.data[i]
		} else {
			result.data[i] = x.data[i]
		}
	}

	return
}

// Add performs a fused: x + y.
func (x Int64x8) Add(y Int64x8) (result Int64x8) {
	for i := range x.data {
		result.data[i] = x.data[i] + y.data[i]
	}

	return
}

// Mul performs a fused: x * y.
func (x Int64x8) Mul(y Int64x8) (result Int64x8) {

	for i := range x.data {
		result.data[i] = x.data[i] * y.data[i]
	}

	return
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Int64x8) Max(y Int64x8) (result Int64x8) {

	for i := range x.data {
		result.data[i] = max(x.data[i], y.data[i])
	}

	return
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Int64x8) Min(y Int64x8) (result Int64x8) {

	for i := range x.data {
		result.data[i] = min(x.data[i], y.data[i])
	}

	return
}

// Neg returns the negation of the elements of x
func (x Int64x8) Neg() (result Int64x8) {

	for i := range x.data {
		result.data[i] = x.data[i] * -1
	}

	return
}

// Sub performs a fused: x - y.
func (x Int64x8) Sub(y Int64x8) (result Int64x8) {

	for i := range x.data {
		result.data[i] = x.data[i] - y.data[i]
	}

	return
}

// And performs a bitwise AND: x & y.
func (x Int64x8) And(y Int64x8) (result Int64x8) {

	for i := range x.data {
		result.data[i] = x.data[i] & y.data[i]
	}

	return
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Int64x8) AndNot(y Int64x8) (result Int64x8) {

	for i := range x.data {
		result.data[i] = x.data[i] &^ y.data[i]
	}

	return
}

// Or performs a bitwise OR: x | y.
func (x Int64x8) Or(y Int64x8) (result Int64x8) {

	for i := range x.data {
		result.data[i] = x.data[i] | y.data[i]
	}

	return
}

// Xor performs a bitwise XOR: x ^ y.
func (x Int64x8) Xor(y Int64x8) (result Int64x8) {

	for i := range x.data {
		result.data[i] = x.data[i] ^ y.data[i]
	}

	return
}

// Not performs a bitwise NOT: ^x.
func (x Int64x8) Not() (result Int64x8) {

	for i := range x.data {
		result.data[i] = ^x.data[i]
	}

	return
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Int64x8) ShiftLeft(count uint) (result Int64x8) {

	for i := range x.data {
		result.data[i] = x.data[i] << count
	}

	return
}

// ShiftRight shifts each element of x right by count bits: x >> count (arithmetic).
func (x Int64x8) ShiftRight(count uint) (result Int64x8) {

	for i := range x.data {
		result.data[i] = x.data[i] >> count
	}

	return
}
