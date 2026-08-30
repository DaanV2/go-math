// nolint:gocritic // TODO Something for later

package simduint64

const uint64_x16_len = 16

type Uint64x16 struct {
	data0 Uint64x8
	data1 Uint64x8
}

func NewUint64x16(data []uint64) Uint64x16 {
	var result Uint64x16

	result.data0 = NewUint64x8(data)
	if len(data) > 8 { // Read atleast 8 points, so there should be more
		result.data1 = NewUint64x8(data[(uint64_x16_len / 2):])
	}

	return result
}

func NewUint64x16Boardcast(value uint64) Uint64x16 {
	return Uint64x16{
		data0: NewUint64x8Boardcast(value),
		data1: NewUint64x8Boardcast(value),
	}
}

// NewUint64x16Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Uint64x16], next uint64_x16_len elements in the next [Uint64x16] etc.
// If not multiple of uint64_x16_len, the last few numbers are put into the last lowest [Uint64x16], with the rest padded to 0
func NewUint64x16Slice(data []uint64) []Uint64x16 {
	l := len(data) / uint64_x16_len
	if (len(data) % uint64_x16_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Uint64x16, 0, l)

	for len(data) > 0 {
		v := NewUint64x16(data)
		result = append(result, v)

		if len(data) >= uint64_x16_len {
			data = data[uint64_x16_len:]
		} else {
			data = nil
		}
	}

	return result
}

func (x Uint64x16) ToSlice() []uint64 {
	result := make([]uint64, uint64_x16_len)
	x.Store(result)

	return result
}

func (x Uint64x16) Len() int {
	return uint64_x16_len
}

func (x Uint64x16) Store(receiver []uint64) {
	switch {
	case len(receiver) == uint64_x16_len:
		x.data0.Store(receiver[:(uint64_x16_len / 2)])
		x.data1.Store(receiver[(uint64_x16_len / 2):])
	case len(receiver) > (uint64_x16_len / 2):
		x.data0.Store(receiver[:(uint64_x16_len / 2)])
		x.data1.Store(receiver[(uint64_x16_len / 2):])
	default:
		x.data0.Store(receiver)
	}
}

// Add performs a fused: x + y.
func (x Uint64x16) Add(y Uint64x16) Uint64x16 {
	return Uint64x16{
		data0: x.data0.Add(y.data0),
		data1: x.data1.Add(y.data1),
	}
}

// Mul performs a fused: x * y.
func (x Uint64x16) Mul(y Uint64x16) Uint64x16 {
	return Uint64x16{
		data0: x.data0.Mul(y.data0),
		data1: x.data1.Mul(y.data1),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Uint64x16) Max(y Uint64x16) Uint64x16 {
	return Uint64x16{
		data0: x.data0.Max(y.data0),
		data1: x.data1.Max(y.data1),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Uint64x16) Min(y Uint64x16) Uint64x16 {
	return Uint64x16{
		data0: x.data0.Min(y.data0),
		data1: x.data1.Min(y.data1),
	}
}

// Sub performs a fused: x - y.
func (x Uint64x16) Sub(y Uint64x16) Uint64x16 {
	return Uint64x16{
		data0: x.data0.Sub(y.data0),
		data1: x.data1.Sub(y.data1),
	}
}

// And performs a bitwise AND: x & y.
func (x Uint64x16) And(y Uint64x16) Uint64x16 {
	return Uint64x16{
		data0: x.data0.And(y.data0),
		data1: x.data1.And(y.data1),
	}
}

// AndNot performs a bitwise AND NOT: x &^ y.
func (x Uint64x16) AndNot(y Uint64x16) Uint64x16 {
	return Uint64x16{
		data0: x.data0.AndNot(y.data0),
		data1: x.data1.AndNot(y.data1),
	}
}

// Or performs a bitwise OR: x | y.
func (x Uint64x16) Or(y Uint64x16) Uint64x16 {
	return Uint64x16{
		data0: x.data0.Or(y.data0),
		data1: x.data1.Or(y.data1),
	}
}

// Xor performs a bitwise XOR: x ^ y.
func (x Uint64x16) Xor(y Uint64x16) Uint64x16 {
	return Uint64x16{
		data0: x.data0.Xor(y.data0),
		data1: x.data1.Xor(y.data1),
	}
}

// Not performs a bitwise NOT: ^x.
func (x Uint64x16) Not() Uint64x16 {
	return Uint64x16{
		data0: x.data0.Not(),
		data1: x.data1.Not(),
	}
}

// ShiftLeft shifts each element of x left by count bits: x << count.
func (x Uint64x16) ShiftLeft(count uint) Uint64x16 {
	return Uint64x16{
		data0: x.data0.ShiftLeft(count),
		data1: x.data1.ShiftLeft(count),
	}
}

// ShiftRight shifts each element of x right by count bits: x >> count (logical).
func (x Uint64x16) ShiftRight(count uint) Uint64x16 {
	return Uint64x16{
		data0: x.data0.ShiftRight(count),
		data1: x.data1.ShiftRight(count),
	}
}
