package simduint64

const uint64_x4_len = 4

func (x Uint64x4) ToSlice() []uint64 {
	result := make([]uint64, uint64_x4_len)
	x.Store(result)

	return result
}

func (x Uint64x4) Len() int {
	return uint64_x4_len
}

// NewUint64x4Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Uint64x4], next 4 elements in the next [Uint64x4] etc.
// If not multiple of 4, the last few numbers are put into the last lowest [Uint64x4], with the rest padded to 0
func NewUint64x4Slice(data []uint64) []Uint64x4 {
	l := len(data) / uint64_x4_len
	if (len(data) % uint64_x4_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Uint64x4, 0, l)

	for len(data) > 0 {
		v := NewUint64x4(data)
		result = append(result, v)

		if len(data) >= uint64_x4_len {
			data = data[uint64_x4_len:]
		} else {
			data = nil
		}
	}

	return result
}
