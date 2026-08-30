package simdint32

const int32_x8_len = 8

func (x Int32x8) ToSlice() []int32 {
	result := make([]int32, int32_x8_len)
	x.Store(result)

	return result
}

func (x Int32x8) Len() int {
	return int32_x8_len
}

// NewInt32x8Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Int32x8], next 8 elements in the next [Int32x8] etc.
// If not multiple of 8, the last few numbers are put into the last lowest [Int32x8], with the rest padded to 0
func NewInt32x8Slice(data []int32) []Int32x8 {
	l := len(data) / int32_x8_len
	if (len(data) % int32_x8_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Int32x8, 0, l)

	for len(data) > 0 {
		v := NewInt32x8(data)
		result = append(result, v)

		if len(data) >= int32_x8_len {
			data = data[int32_x8_len:]
		} else {
			data = nil
		}
	}

	return result
}
