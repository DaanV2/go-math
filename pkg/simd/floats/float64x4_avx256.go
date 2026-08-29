//go:build simd_avx256

package simdfloats

import "math"

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x Float64x4) Scale(y Float64x4) Float64x4 {
	var result [float64_x4_len]float64
	var vx [float64_x4_len]float64
	var vy [float64_x4_len]float64

	x.Store(vx[:])
	y.Store(vy[:])

	for i := range vx {
		result[i] = vx[i] * math.Pow(2, vy[i])
	}

	return NewFloat64x4(result[:])
}
