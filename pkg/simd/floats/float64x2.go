package simdfloats

const float64_x2_len = 2

func (x Float64x2) ToSlice() []float64 {
	result := make([]float64, float64_x2_len)
	x.Store(result)

	return result
}

func (x Float64x2) Len() int {
	return float64_x2_len
}

// NewFloat64x2Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Float64x2], next 4 elements in the next [Float64x2] etc.
// If not multiple of 4, the last few numbers are put into the last lowest [Float64x2], with the rest padded to 0
func NewFloat64x2Slice(data []float64) []Float64x2 {
	l := len(data) / float64_x2_len
	if (len(data) % float64_x2_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Float64x2, 0, l)

	for len(data) > 0 {
		v := NewFloat64x2(data)
		result = append(result, v)

		if len(data) >= float64_x2_len {
			data = data[float64_x2_len:]
		} else {
			data = nil
		}
	}

	return result
}
