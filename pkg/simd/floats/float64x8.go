package simdfloats

const float64_x8_len = 8

func (x Float64x8) ToSlice() []float64 {
	result := make([]float64, float64_x8_len)
	x.Store(result)

	return result
}

func (x Float64x8) Len() int {
	return float64_x8_len
}

// NewFloat64x8Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Float64x8], next float64_x8_len elements in the next [Float64x8] etc.
// If not multiple of float64_x8_len, the last few numbers are put into the last lowest [Float64x8], with the rest padded to 0
func NewFloat64x8Slice(data []float64) []Float64x8 {
	l := len(data) / float64_x8_len
	if (len(data) % float64_x8_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Float64x8, 0, l)

	for len(data) > 0 {
		v := NewFloat64x8(data)
		result = append(result, v)

		if len(data) >= float64_x8_len {
			data = data[float64_x8_len:]
		} else {
			data = nil
		}
	}

	return result
}
