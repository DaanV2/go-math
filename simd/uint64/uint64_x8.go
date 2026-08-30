package simduint64

const uint64_x8_len = 8

func (x Uint64x8) ToSlice() []uint64 {
	result := make([]uint64, uint64_x8_len)
	x.Store(result)

	return result
}

func (x Uint64x8) Len() int {
	return uint64_x8_len
}

// NewUint64x8Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Uint64x8], next uint64_x8_len elements in the next [Uint64x8] etc.
// If not multiple of uint64_x8_len, the last few numbers are put into the last lowest [Uint64x8], with the rest padded to 0
func NewUint64x8Slice(data []uint64) []Uint64x8 {
	l := len(data) / uint64_x8_len
	if (len(data) % uint64_x8_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Uint64x8, 0, l)

	for len(data) > 0 {
		v := NewUint64x8(data)
		result = append(result, v)

		if len(data) >= uint64_x8_len {
			data = data[uint64_x8_len:]
		} else {
			data = nil
		}
	}

	return result
}
