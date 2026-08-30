package simduint64

const uint64_x2_len = 2

func (x Uint64x2) ToSlice() []uint64 {
	result := make([]uint64, uint64_x2_len)
	x.Store(result)

	return result
}

func (x Uint64x2) Len() int {
	return uint64_x2_len
}

// NewUint64x2Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Uint64x2], next 2 elements in the next [Uint64x2] etc.
// If not multiple of 2, the last few numbers are put into the last lowest [Uint64x2], with the rest padded to 0
func NewUint64x2Slice(data []uint64) []Uint64x2 {
	l := len(data) / uint64_x2_len
	if (len(data) % uint64_x2_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Uint64x2, 0, l)

	for len(data) > 0 {
		v := NewUint64x2(data)
		result = append(result, v)

		if len(data) >= uint64_x2_len {
			data = data[uint64_x2_len:]
		} else {
			data = nil
		}
	}

	return result
}
