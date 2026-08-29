package simdfloats

// NewFloat64x8Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Float64x8], next 8 elements in the next [Float64x8] etc.
// If not multiple of 8, the last few numbers are put into the last lowest [Float64x8], with the rest padded to 0
func NewFloat64x8Slice(data []float64) []Float64x8 {
	l := len(data) / 8
	if (len(data) % 8) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Float64x8, 0, l)

	for len(data) > 8 {
		v := NewFloat64x8(data)
		result = append(result, v)
		data = data[8:]
	}

	return result
}
