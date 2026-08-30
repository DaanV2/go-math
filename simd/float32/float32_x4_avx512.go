//go:build simd_avx512

package simdfloat32

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x Float32x4) Scale(y Float32x4) Float32x4 {
	return Float32x4{
		data: x.data.Scale(y.data),
	}
}
