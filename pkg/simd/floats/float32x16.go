package simdfloats

const float32_x16_len = 16

func (x Float32x16) ToSlice() []float32 {
	result := make([]float32, float32_x16_len)
	x.Store(result)

	return result
}

func (x Float32x16) Len() int {
	return float32_x16_len
}

// NewFloat32x16Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Float32x16], next float32_x16_len elements in the next [Float32x16] etc.
// If not multiple of float32_x16_len, the last few numbers are put into the last lowest [Float32x16], with the rest padded to 0
func NewFloat32x16Slice(data []float32) []Float32x16 {
	l := len(data) / float32_x16_len
	if (len(data) % float32_x16_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Float32x16, 0, l)

	for len(data) > 0 {
		v := NewFloat32x16(data)
		result = append(result, v)

		if len(data) >= float32_x16_len {
			data = data[float32_x16_len:]
		} else {
			data = nil
		}
	}

	return result
}
