package simdfloats

const float32_x4_len = 4

func (x Float32x4) ToSlice() []float32 {
	result := make([]float32, float32_x4_len)
	x.Store(result)

	return result
}

func (x Float32x4) Len() int {
	return float32_x4_len
}

// NewFloat32x4Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Float32x4], next 4 elements in the next [Float32x4] etc.
// If not multiple of 4, the last few numbers are put into the last lowest [Float32x4], with the rest padded to 0
func NewFloat32x4Slice(data []float32) []Float32x4 {
	l := len(data) / float32_x4_len
	if (len(data) % float32_x4_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Float32x4, 0, l)

	for len(data) > 0 {
		v := NewFloat32x4(data)
		result = append(result, v)

		if len(data) >= float32_x4_len {
			data = data[float32_x4_len:]
		} else {
			data = nil
		}
	}

	return result
}
