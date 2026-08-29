package simdfloats

func (x Float64x8) ToSlice() []float64 {
	var result [8]float64

	x.Store(result[:])

	return result[:]
}

// NewFloat64x8Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Float64x8], next 8 elements in the next [Float64x8] etc.
// If not multiple of 8, the last few numbers are put into the last lowest [Float64x8], with the rest padded to 0
func NewFloat64x8Slice(data []float64) []Float64x8 {
	l := len(data) / 8
	if (len(data) % 8) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Float64x8, 0, l)

	for len(data) > 0 {
		v := NewFloat64x8(data)
		result = append(result, v)

		if len(data) >= 8 {
			data = data[8:]
		} else {
			data = nil
		}
	}

	return result
}
