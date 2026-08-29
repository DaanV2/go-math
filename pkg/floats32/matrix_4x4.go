package floats32

import "simd/archsimd"

type Matrix4x4 struct {
	data archsimd.Float32x16
}
