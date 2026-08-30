package simduint32

const uint32_x16_len = 16

func (x Uint32x16) ToSlice() []uint32 {
	result := make([]uint32, uint32_x16_len)
	x.Store(result)

	return result
}

func (x Uint32x16) Len() int {
	return uint32_x16_len
}

// NewUint32x16Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Uint32x16], next uint32_x16_len elements in the next [Uint32x16] etc.
// If not multiple of uint32_x16_len, the last few numbers are put into the last lowest [Uint32x16], with the rest padded to 0
func NewUint32x16Slice(data []uint32) []Uint32x16 {
	l := len(data) / uint32_x16_len
	if (len(data) % uint32_x16_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Uint32x16, 0, l)

	for len(data) > 0 {
		v := NewUint32x16(data)
		result = append(result, v)

		if len(data) >= uint32_x16_len {
			data = data[uint32_x16_len:]
		} else {
			data = nil
		}
	}

	return result
}
