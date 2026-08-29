package floats64

import (
	"simd/archsimd"
)

const VECTOR_4_LEN = 4

type Vector4 struct {
	archsimd.Float64x4
}

func NewVector4(data [VECTOR_4_LEN]float64) Vector4 {
	var result Vector4

	result.Float64x4 = archsimd.LoadFloat64x4(data[:])

	return result
}

func (v Vector4) Elem1() float64 { return v.GetLo().GetElem(0) }
func (v Vector4) Elem2() float64 { return v.GetLo().GetElem(1) }
func (v Vector4) Elem3() float64 { return v.GetHi().GetElem(2) }
func (v Vector4) Elem4() float64 { return v.GetHi().GetElem(3) }

func (v Vector4) Add(other Vector4) Vector4 {
	return Vector4{v.Float64x4.Add(other.Float64x4)}
}
