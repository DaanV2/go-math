//go:build simd_none

package simdfloats

type Float64x8 struct {
	data [8]float64
}

func NewFloat64x8(data []float64) Float64x8 {
	var result Float64x8
	copy(result.data[:], data)

	return result
}

func (v Float64x8) ToSlice() []float64 {
	return v.data[:]
}

func (v Float64x8) Add(other Float64x8) Float64x8 {
	var result Float64x8

	for i := range v.data {
		result.data[i] = v.data[i] + other.data[i]
	}

	return result
}
