// nolint:gocritic // TODO Something for later

package simdint64

const int64_x16_len = 16

type Int64x16 struct {
	data0 Int64x8
	data1 Int64x8
}

func NewInt64x16(data []int64) Int64x16 {
	var result Int64x16

	result.data0 = NewInt64x8(data)
	if len(data) > 8 { // Read atleast 8 points, so there should be more
		result.data1 = NewInt64x8(data[(int64_x16_len / 2):])
	}

	return result
}

// NewInt64x16Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Int64x16], next int64_x16_len elements in the next [Int64x16] etc.
// If not multiple of int64_x16_len, the last few numbers are put into the last lowest [Int64x16], with the rest padded to 0
func NewInt64x16Slice(data []int64) []Int64x16 {
	l := len(data) / int64_x16_len
	if (len(data) % int64_x16_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Int64x16, 0, l)

	for len(data) > 0 {
		v := NewInt64x16(data)
		result = append(result, v)

		if len(data) >= int64_x16_len {
			data = data[int64_x16_len:]
		} else {
			data = nil
		}
	}

	return result
}

func (x Int64x16) ToSlice() []int64 {
	result := make([]int64, int64_x16_len)
	x.Store(result)

	return result
}

func (x Int64x16) Len() int {
	return int64_x16_len
}

func (x Int64x16) Store(receiver []int64) {
	switch {
	case len(receiver) == int64_x16_len:
		x.data0.Store(receiver[:(int64_x16_len / 2)])
		x.data1.Store(receiver[(int64_x16_len / 2):])
	case len(receiver) > (int64_x16_len / 2):
		x.data0.Store(receiver[:(int64_x16_len / 2)])
		x.data1.Store(receiver[(int64_x16_len / 2):])
	default:
		x.data0.Store(receiver)
	}
}

// Abs returns the absolute values of the elements of x
func (x Int64x16) Abs() Int64x16 {
	return Int64x16{
		data0: x.data0.Abs(),
		data1: x.data1.Abs(),
	}
}

// Add performs a fused: x + y.
func (x Int64x16) Add(y Int64x16) Int64x16 {
	return Int64x16{
		data0: x.data0.Add(y.data0),
		data1: x.data1.Add(y.data1),
	}
}

// Mul performs a fused: x * y.
func (x Int64x16) Mul(y Int64x16) Int64x16 {
	return Int64x16{
		data0: x.data0.Mul(y.data0),
		data1: x.data1.Mul(y.data1),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Int64x16) Max(y Int64x16) Int64x16 {
	return Int64x16{
		data0: x.data0.Max(y.data0),
		data1: x.data1.Max(y.data1),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Int64x16) Min(y Int64x16) Int64x16 {
	return Int64x16{
		data0: x.data0.Min(y.data0),
		data1: x.data1.Min(y.data1),
	}
}

// Neg returns the negation of the elements of x
func (x Int64x16) Neg() Int64x16 {
	return Int64x16{
		data0: x.data0.Neg(),
		data1: x.data1.Neg(),
	}
}

// Sub performs a fused: x - y.
func (x Int64x16) Sub(y Int64x16) Int64x16 {
	return Int64x16{
		data0: x.data0.Sub(y.data0),
		data1: x.data1.Sub(y.data1),
	}
}

// And performs a bitwise AND: x & y.
func (x Int64x16) And(y Int64x16) Int64x16 {
	return Int64x16{
		data0: x.data0.And(y.data0),
		data1: x.data1.And(y.data1),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Int64x16) AndNot(y Int64x16) Int64x16 {
	return Int64x16{
		data0: x.data0.AndNot(y.data0),
		data1: x.data1.AndNot(y.data1),
	}
}

// Or performs a bitwise OR: x | y.
func (x Int64x16) Or(y Int64x16) Int64x16 {
	return Int64x16{
		data0: x.data0.Or(y.data0),
		data1: x.data1.Or(y.data1),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Int64x16) Xor(y Int64x16) Int64x16 {
	return Int64x16{
		data0: x.data0.Xor(y.data0),
		data1: x.data1.Xor(y.data1),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Int64x16) Not() Int64x16 {
	return Int64x16{
		data0: x.data0.Not(),
		data1: x.data1.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Int64x16) ShiftLeft(count uint) Int64x16 {
	return Int64x16{
		data0: x.data0.ShiftLeft(count),
		data1: x.data1.ShiftLeft(count),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (arithmetic).
func (x Int64x16) ShiftRight(count uint) Int64x16 {
	return Int64x16{
		data0: x.data0.ShiftRight(count),
		data1: x.data1.ShiftRight(count),
	}
}
