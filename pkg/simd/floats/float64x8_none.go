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

func (v Float64x8) Store(receiver []float64) {
	copy(receiver[:], v.data[:])
}

func (x Float64x8) Add(y Float64x8) Float64x8 {
	var result Float64x8

	for i := range x.data {
		result.data[i] = x.data[i] + y.data[i]
	}

	return result
}

// Sub performs a fused: x - y.
func (x Float64x8) Sub(y Float64x8) Float64x8 {
	var result Float64x8

	for i := range x.data {
		result.data[i] = x.data[i] - y.data[i]
	}

	return result
}

// Mul performs a fused: x * y.
func (x Float64x8) Mul(y Float64x8) Float64x8 {
	var result Float64x8

	for i := range x.data {
		result.data[i] = x.data[i] * y.data[i]
	}

	return result
}

// MulAdd performs a fused: (x * y) + z.
func (x Float64x8) MulAdd(y Float64x8, z Float64x8) Float64x8 {
	var result Float64x8

	for i := range x.data {
		result.data[i] = (x.data[i] * y.data[i]) + z.data[i]
	}

	return result
}
