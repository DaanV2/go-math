package floats64

import "simd/archsimd"

const MATRIX_2X2_LEN = 2 * 2

type Matrix2x2 [MATRIX_2X2_LEN]float64

func NewMatrix2x2(data [MATRIX_2X2_LEN]float64) Matrix2x2 {
	return Matrix2x2(data)
}

func (m Matrix2x2) Add(other Matrix2x2) Matrix2x2 {
	var result = m

	archsimd.LoadFloat64x4(result[0:4]).Add(archsimd.LoadFloat64x4(other[0:4])).Store(result[0:4])

	return Matrix2x2(result)
}
