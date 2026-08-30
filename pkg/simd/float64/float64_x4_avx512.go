//go:build simd_avx512

package simdfloat64

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x Float64x4) Scale(y Float64x4) Float64x4 {
	return Float64x4{
		data: x.data.Scale(y.data),
	}
}
