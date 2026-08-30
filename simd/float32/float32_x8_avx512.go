//go:build simd_avx512

package simdfloat32

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x Float32x8) Scale(y Float32x8) Float32x8 {
	return Float32x8{
		data: x.data.Scale(y.data),
	}
}
