//go:build simd_avx512

package simdfloat32

import "simd/archsimd"

type Float32x16 struct {
	data archsimd.Float32x16
}

func NewFloat32x16(data []float32) Float32x16 {
	var result Float32x16

	result.data, _ = archsimd.LoadFloat32x16Part(data)

	return result
}

func NewFloat32x16Boardcast(value float32) Float32x16 {
	return Float32x16{
		data: archsimd.BroadcastFloat32x16(value),
	}
}

func (v Float32x16) Store(receiver []float32) {
	v.data.StorePart(receiver)
}

// Abs returns the absolute values of the elements of x
func (x Float32x16) Abs() Float32x16 {
	return Float32x16{
		data: x.data.Abs(),
	}
}

// Add performs a fused: x + y.
func (v Float32x16) Add(other Float32x16) Float32x16 {
	return Float32x16{
		data: v.data.Add(other.data),
	}
}

// Div performs a fused: x / y.
func (x Float32x16) Div(y Float32x16) Float32x16 {
	return Float32x16{
		data: x.data.Div(y.data),
	}
}

// Mul performs a fused: x * y.
func (x Float32x16) Mul(y Float32x16) Float32x16 {
	return Float32x16{
		data: x.data.Mul(y.data),
	}
}

// MulAdd performs a fused: (x * y) + z.
func (x Float32x16) MulAdd(y, z Float32x16) Float32x16 {
	return Float32x16{
		data: x.data.MulAdd(y.data, z.data),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Float32x16) Max(y Float32x16) Float32x16 {
	return Float32x16{
		data: x.data.Max(y.data),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Float32x16) Min(y Float32x16) Float32x16 {
	return Float32x16{
		data: x.data.Min(y.data),
	}
}

// Neg returns the negation of the elements of x
func (x Float32x16) Neg() Float32x16 {
	return Float32x16{
		data: x.data.Neg(),
	}
}

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x Float32x16) Scale(y Float32x16) Float32x16 {
	return Float32x16{
		data: x.data.Scale(y.data),
	}
}

// Sub performs a fused: x - y.
func (x Float32x16) Sub(y Float32x16) Float32x16 {
	return Float32x16{
		data: x.data.Sub(y.data),
	}
}

// Sqrt computes the square root of each element.
func (x Float32x16) Sqrt() Float32x16 {
	return Float32x16{
		data: x.data.Sqrt(),
	}
}
