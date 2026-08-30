package simdint64

const int64_x4_len = 4

func (x Int64x4) ToSlice() []int64 {
	result := make([]int64, int64_x4_len)
	x.Store(result)

	return result
}

func (x Int64x4) Len() int {
	return int64_x4_len
}

// NewInt64x4Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Int64x4], next 4 elements in the next [Int64x4] etc.
// If not multiple of 4, the last few numbers are put into the last lowest [Int64x4], with the rest padded to 0
func NewInt64x4Slice(data []int64) []Int64x4 {
	l := len(data) / int64_x4_len
	if (len(data) % int64_x4_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Int64x4, 0, l)

	for len(data) > 0 {
		v := NewInt64x4(data)
		result = append(result, v)

		if len(data) >= int64_x4_len {
			data = data[int64_x4_len:]
		} else {
			data = nil
		}
	}

	return result
}
