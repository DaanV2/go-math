package floats64

import (
	"simd/archsimd"
)

const MATRIX_4X4_LEN = 4 * 4

type Matrix4x4 struct {
	row0 archsimd.Float64x4
	row1 archsimd.Float64x4
	row2 archsimd.Float64x4
	row3 archsimd.Float64x4
}

func NewMatrix4x4(data [MATRIX_4X4_LEN]float64) Matrix4x4 {
	return Matrix4x4{
		row0: archsimd.LoadFloat64x4(data[0:4]),
		row1: archsimd.LoadFloat64x4(data[4:8]),
		row2: archsimd.LoadFloat64x4(data[8:12]),
		row3: archsimd.LoadFloat64x4(data[12:16]),
	}
}

func (m Matrix4x4) Add(other Matrix4x4) Matrix4x4 {
	return Matrix4x4{
		row0: m.row0.Add(other.row0),
		row1: m.row1.Add(other.row1),
		row2: m.row2.Add(other.row2),
		row3: m.row3.Add(other.row3),
	}
}
