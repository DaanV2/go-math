package simdint64

const int64_x2_len = 2

func (x Int64x2) ToSlice() []int64 {
	result := make([]int64, int64_x2_len)
	x.Store(result)

	return result
}

func (x Int64x2) Len() int {
	return int64_x2_len
}

// NewInt64x2Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Int64x2], next 2 elements in the next [Int64x2] etc.
// If not multiple of 2, the last few numbers are put into the last lowest [Int64x2], with the rest padded to 0
func NewInt64x2Slice(data []int64) []Int64x2 {
	l := len(data) / int64_x2_len
	if (len(data) % int64_x2_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Int64x2, 0, l)

	for len(data) > 0 {
		v := NewInt64x2(data)
		result = append(result, v)

		if len(data) >= int64_x2_len {
			data = data[int64_x2_len:]
		} else {
			data = nil
		}
	}

	return result
}
