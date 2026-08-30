//go:build simd_avx512

package simdfloat64

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x Float64x2) Scale(y Float64x2) Float64x2 {
	return Float64x2{
		data: x.data.Scale(y.data),
	}
}
