package simduint32

const uint32_x4_len = 4

func (x Uint32x4) ToSlice() []uint32 {
	result := make([]uint32, uint32_x4_len)
	x.Store(result)

	return result
}

func (x Uint32x4) Len() int {
	return uint32_x4_len
}

// NewUint32x4Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Uint32x4], next 4 elements in the next [Uint32x4] etc.
// If not multiple of 4, the last few numbers are put into the last lowest [Uint32x4], with the rest padded to 0
func NewUint32x4Slice(data []uint32) []Uint32x4 {
	l := len(data) / uint32_x4_len
	if (len(data) % uint32_x4_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Uint32x4, 0, l)

	for len(data) > 0 {
		v := NewUint32x4(data)
		result = append(result, v)

		if len(data) >= uint32_x4_len {
			data = data[uint32_x4_len:]
		} else {
			data = nil
		}
	}

	return result
}
