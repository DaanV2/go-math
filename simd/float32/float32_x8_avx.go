//go:build simd_avx256 || simd_avx512

package simdfloat32

import "simd/archsimd"

type Float32x8 struct {
	data archsimd.Float32x8
}

func NewFloat32x8(data []float32) Float32x8 {
	var result Float32x8

	result.data, _ = archsimd.LoadFloat32x8Part(data)

	return result
}

func NewFloat32x8Boardcast(value float32) Float32x8 {
	return Float32x8{
		data: archsimd.BroadcastFloat32x8(value),
	}
}

func (x Float32x8) Store(receiver []float32) {
	if len(receiver) >= float32_x8_len {
		x.data.Store(receiver[:float32_x8_len])
	} else {
		x.data.StorePart(receiver)
	}
}

// Abs returns the absolute values of the elements of x
func (x Float32x8) Abs() Float32x8 {
	return Float32x8{
		data: x.data.Abs(),
	}
}

// Add performs a fused: x + y.
func (x Float32x8) Add(y Float32x8) Float32x8 {
	return Float32x8{
		data: x.data.Add(y.data),
	}
}

// Div performs a fused: x / y.
func (x Float32x8) Div(y Float32x8) Float32x8 {
	return Float32x8{
		data: x.data.Div(y.data),
	}
}

// Mul performs a fused: x * y.
func (x Float32x8) Mul(y Float32x8) Float32x8 {
	return Float32x8{
		data: x.data.Mul(y.data),
	}
}

// MulAdd performs a fused: (x * y) + z.
func (x Float32x8) MulAdd(y, z Float32x8) Float32x8 {
	return Float32x8{
		data: x.data.MulAdd(y.data, z.data),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Float32x8) Max(y Float32x8) Float32x8 {
	return Float32x8{
		data: x.data.Max(y.data),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Float32x8) Min(y Float32x8) Float32x8 {
	return Float32x8{
		data: x.data.Min(y.data),
	}
}

// Neg returns the negation of the elements of x
func (x Float32x8) Neg() Float32x8 {
	return Float32x8{
		data: x.data.Neg(),
	}
}

// Sub performs a fused: x - y.
func (x Float32x8) Sub(y Float32x8) Float32x8 {
	return Float32x8{
		data: x.data.Sub(y.data),
	}
}

// Sqrt computes the square root of each element.
func (x Float32x8) Sqrt() Float32x8 {
	return Float32x8{
		data: x.data.Sqrt(),
	}
}
