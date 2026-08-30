package simdint64

const int64_x8_len = 8

func (x Int64x8) ToSlice() []int64 {
	result := make([]int64, int64_x8_len)
	x.Store(result)

	return result
}

func (x Int64x8) Len() int {
	return int64_x8_len
}

// NewInt64x8Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Int64x8], next int64_x8_len elements in the next [Int64x8] etc.
// If not multiple of int64_x8_len, the last few numbers are put into the last lowest [Int64x8], with the rest padded to 0
func NewInt64x8Slice(data []int64) []Int64x8 {
	l := len(data) / int64_x8_len
	if (len(data) % int64_x8_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Int64x8, 0, l)

	for len(data) > 0 {
		v := NewInt64x8(data)
		result = append(result, v)

		if len(data) >= int64_x8_len {
			data = data[int64_x8_len:]
		} else {
			data = nil
		}
	}

	return result
}
