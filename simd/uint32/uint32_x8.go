package simduint32

const uint32_x8_len = 8

func (x Uint32x8) ToSlice() []uint32 {
	result := make([]uint32, uint32_x8_len)
	x.Store(result)

	return result
}

func (x Uint32x8) Len() int {
	return uint32_x8_len
}

// NewUint32x8Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Uint32x8], next 8 elements in the next [Uint32x8] etc.
// If not multiple of 8, the last few numbers are put into the last lowest [Uint32x8], with the rest padded to 0
func NewUint32x8Slice(data []uint32) []Uint32x8 {
	l := len(data) / uint32_x8_len
	if (len(data) % uint32_x8_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Uint32x8, 0, l)

	for len(data) > 0 {
		v := NewUint32x8(data)
		result = append(result, v)

		if len(data) >= uint32_x8_len {
			data = data[uint32_x8_len:]
		} else {
			data = nil
		}
	}

	return result
}
