//go:build simd_avx256

package simdfloat32

import "math"

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x Float32x4) Scale(y Float32x4) Float32x4 {
	var result [float32_x4_len]float32
	var vx [float32_x4_len]float32
	var vy [float32_x4_len]float32

	x.Store(vx[:])
	y.Store(vy[:])

	for i := range vx {
		result[i] = vx[i] * float32(math.Pow(2, float64(vy[i])))
	}

	return NewFloat32x4(result[:])
}
