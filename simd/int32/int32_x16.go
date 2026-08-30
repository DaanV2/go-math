package simdint32

const int32_x16_len = 16

func (x Int32x16) ToSlice() []int32 {
	result := make([]int32, int32_x16_len)
	x.Store(result)

	return result
}

func (x Int32x16) Len() int {
	return int32_x16_len
}

// NewInt32x16Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Int32x16], next int32_x16_len elements in the next [Int32x16] etc.
// If not multiple of int32_x16_len, the last few numbers are put into the last lowest [Int32x16], with the rest padded to 0
func NewInt32x16Slice(data []int32) []Int32x16 {
	l := len(data) / int32_x16_len
	if (len(data) % int32_x16_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Int32x16, 0, l)

	for len(data) > 0 {
		v := NewInt32x16(data)
		result = append(result, v)

		if len(data) >= int32_x16_len {
			data = data[int32_x16_len:]
		} else {
			data = nil
		}
	}

	return result
}
