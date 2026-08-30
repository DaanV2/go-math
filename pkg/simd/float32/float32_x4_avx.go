//go:build simd_avx256 || simd_avx512

package simdfloat32

import "simd/archsimd"

type Float32x4 struct {
	data archsimd.Float32x4
}

func NewFloat32x4(data []float32) Float32x4 {
	var result Float32x4

	result.data, _ = archsimd.LoadFloat32x4Part(data)

	return result
}

func (x Float32x4) Store(receiver []float32) {
	if len(receiver) >= float32_x4_len {
		x.data.Store(receiver[:float32_x4_len])
	} else {
		x.data.StorePart(receiver)
	}
}

// Abs returns the absolute values of the elements of x
func (x Float32x4) Abs() Float32x4 {
	return Float32x4{
		data: x.data.Abs(),
	}
}

// Add performs a fused: x + y.
func (x Float32x4) Add(y Float32x4) Float32x4 {
	return Float32x4{
		data: x.data.Add(y.data),
	}
}

// Div performs a fused: x / y.
func (x Float32x4) Div(y Float32x4) Float32x4 {
	return Float32x4{
		data: x.data.Div(y.data),
	}
}

// Mul performs a fused: x * y.
func (x Float32x4) Mul(y Float32x4) Float32x4 {
	return Float32x4{
		data: x.data.Mul(y.data),
	}
}

// MulAdd performs a fused: (x * y) + z.
func (x Float32x4) MulAdd(y, z Float32x4) Float32x4 {
	return Float32x4{
		data: x.data.MulAdd(y.data, z.data),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Float32x4) Max(y Float32x4) Float32x4 {
	return Float32x4{
		data: x.data.Max(y.data),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Float32x4) Min(y Float32x4) Float32x4 {
	return Float32x4{
		data: x.data.Min(y.data),
	}
}

// Neg returns the negation of the elements of x
func (x Float32x4) Neg() Float32x4 {
	return Float32x4{
		data: x.data.Neg(),
	}
}

// Sub performs a fused: x - y.
func (x Float32x4) Sub(y Float32x4) Float32x4 {
	return Float32x4{
		data: x.data.Sub(y.data),
	}
}

// Sqrt computes the square root of each element.
func (x Float32x4) Sqrt() Float32x4 {
	return Float32x4{
		data: x.data.Sqrt(),
	}
}
