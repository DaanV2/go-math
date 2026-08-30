package simdfloat32

const float32_x8_len = 8

func (x Float32x8) ToSlice() []float32 {
	result := make([]float32, float32_x8_len)
	x.Store(result)

	return result
}

func (x Float32x8) Len() int {
	return float32_x8_len
}

// NewFloat32x8Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Float32x8], next 4 elements in the next [Float32x8] etc.
// If not multiple of 4, the last few numbers are put into the last lowest [Float32x8], with the rest padded to 0
func NewFloat32x8Slice(data []float32) []Float32x8 {
	l := len(data) / float32_x8_len
	if (len(data) % float32_x8_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Float32x8, 0, l)

	for len(data) > 0 {
		v := NewFloat32x8(data)
		result = append(result, v)

		if len(data) >= float32_x8_len {
			data = data[float32_x8_len:]
		} else {
			data = nil
		}
	}

	return result
}
