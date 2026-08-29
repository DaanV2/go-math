//go:build simd_avx256 || simd_avx512

package simdfloats

import "simd/archsimd"

type Float64x4 struct {
	data archsimd.Float64x4
}

func NewFloat64x4(data []float64) Float64x4 {
	var result Float64x4

	var n int
	result.data, n = archsimd.LoadFloat64x4Part(data)
	if n == 4 { // Read atleast 4 points, so there should be more
		result.data1, _ = archsimd.LoadFloat64x4Part(data[4:])
	}

	return result
}

func (x Float64x4) Store(receiver []float64) {
	switch {
	case len(receiver) == float64_x4_len:
		x.data.Store(receiver[:(float64_x4_len / 2)])
		x.data1.Store(receiver[(float64_x4_len / 2):])
	case len(receiver) > (float64_x4_len / 2):
		x.data.Store(receiver[:(float64_x4_len / 2)])
		_ = x.data1.StorePart(receiver[(float64_x4_len / 2):])
	default:
		_ = x.data.StorePart(receiver)
	}
}

// Abs returns the absolute values of the elements of x
func (x Float64x4) Abs() Float64x4 {
	return Float64x4{
		data:  x.data.Abs(),
		data1: x.data1.Abs(),
	}
}

// Add performs a fused: x + y.
func (x Float64x4) Add(y Float64x4) Float64x4 {
	return Float64x4{
		data:  x.data.Add(y.data),
		data1: x.data1.Add(y.data1),
	}
}

// Div performs a fused: x / y.
func (x Float64x4) Div(y Float64x4) Float64x4 {
	return Float64x4{
		data:  x.data.Div(y.data),
		data1: x.data1.Div(y.data1),
	}
}

// Mul performs a fused: x * y.
func (x Float64x4) Mul(y Float64x4) Float64x4 {
	return Float64x4{
		data:  x.data.Mul(y.data),
		data1: x.data1.Mul(y.data1),
	}
}

// MulAdd performs a fused: (x * y) + z.
func (x Float64x4) MulAdd(y, z Float64x4) Float64x4 {
	return Float64x4{
		data:  x.data.MulAdd(y.data, z.data),
		data1: x.data1.MulAdd(y.data1, z.data1),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Float64x4) Max(y Float64x4) Float64x4 {
	return Float64x4{
		data:  x.data.Max(y.data),
		data1: x.data1.Max(y.data1),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Float64x4) Min(y Float64x4) Float64x4 {
	return Float64x4{
		data:  x.data.Min(y.data),
		data1: x.data1.Min(y.data1),
	}
}

// Neg returns the negation of the elements of x
func (x Float64x4) Neg() Float64x4 {
	return Float64x4{
		data:  x.data.Neg(),
		data1: x.data1.Neg(),
	}
}

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x Float64x4) Scale(y Float64x4) Float64x4 {
	return Float64x4{
		data:  x.data.Scale(y.data),
		data1: x.data1.Scale(y.data1),
	}
}

// Sub performs a fused: x - y.
func (x Float64x4) Sub(y Float64x4) Float64x4 {
	return Float64x4{
		data:  x.data.Sub(y.data),
		data1: x.data1.Sub(y.data1),
	}
}

// Sqrt computes the square root of each element.
func (x Float64x4) Sqrt() Float64x4 {
	return Float64x4{
		data:  x.data.Sqrt(),
		data1: x.data1.Sqrt(),
	}
}
