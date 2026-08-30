//go:build simd_avx256

package simdfloat64

import "math"

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x Float64x2) Scale(y Float64x2) Float64x2 {
	var result [float64_x2_len]float64
	var vx [float64_x2_len]float64
	var vy [float64_x2_len]float64

	x.Store(vx[:])
	y.Store(vy[:])

	for i := range vx {
		result[i] = vx[i] * math.Pow(2, vy[i])
	}

	return NewFloat64x2(result[:])
}
