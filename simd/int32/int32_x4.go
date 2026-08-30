package simdint32

const int32_x4_len = 4

func (x Int32x4) ToSlice() []int32 {
	result := make([]int32, int32_x4_len)
	x.Store(result)

	return result
}

func (x Int32x4) Len() int {
	return int32_x4_len
}

// NewInt32x4Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Int32x4], next 4 elements in the next [Int32x4] etc.
// If not multiple of 4, the last few numbers are put into the last lowest [Int32x4], with the rest padded to 0
func NewInt32x4Slice(data []int32) []Int32x4 {
	l := len(data) / int32_x4_len
	if (len(data) % int32_x4_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Int32x4, 0, l)

	for len(data) > 0 {
		v := NewInt32x4(data)
		result = append(result, v)

		if len(data) >= int32_x4_len {
			data = data[int32_x4_len:]
		} else {
			data = nil
		}
	}

	return result
}
