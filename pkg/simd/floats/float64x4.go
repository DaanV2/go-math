package simdfloats

const float64_x4_len = 4

func (x Float64x4) ToSlice() []float64 {
	result := make([]float64, float64_x4_len)
	x.Store(result)

	return result
}

func (x Float64x4) Len() int {
	return float64_x4_len
}

// NewFloat64x4Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Float64x4], next 4 elements in the next [Float64x4] etc.
// If not multiple of 4, the last few numbers are put into the last lowest [Float64x4], with the rest padded to 0
func NewFloat64x4Slice(data []float64) []Float64x4 {
	l := len(data) / float64_x4_len
	if (len(data) % float64_x4_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Float64x4, 0, l)

	for len(data) > 0 {
		v := NewFloat64x4(data)
		result = append(result, v)

		if len(data) >= float64_x4_len {
			data = data[float64_x4_len:]
		} else {
			data = nil
		}
	}

	return result
}
