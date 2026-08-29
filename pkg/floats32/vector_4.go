package floats32

import (
	"simd/archsimd"
)

type Vector4 struct {
	archsimd.Float32x4
}

func (v Vector4) Add(other Vector4) Vector4 {
	return Vector4{v.Float32x4.Add(other.Float32x4)}
}
