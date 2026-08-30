package matrices

import simdfloat32 "github.com/daanv2/go-math/simd/float32"

type Float32_2x2 struct {
	data simdfloat32.Float32x4
}

func NewFloat32_2x2(data []float32) Float32_2x2 {
	var result Float32_2x2
	result.data = simdfloat32.NewFloat32x4(data)

	return result
}

func (m *Float32_2x2) Width() int  { return 2 }
func (m *Float32_2x2) Height() int { return 2 }
func (m *Float32_2x2) Len() int    { return 2 * 2 }

func (m *Float32_2x2) ToSlice() []float32 {
	result := make([]float32, m.Len())

	m.Store(result)

	return result
}

func (m *Float32_2x2) Store(receiver []float32) {
	m.data.Store(receiver)
}

// Abs performs: c_ij = abs(x_ij)
func (x *Float32_2x2) Abs() Float32_2x2 {
	return Float32_2x2{
		data: x.data.Abs(),
	}
}

// Add performs: c_ij = x_ij + y_ij
func (x *Float32_2x2) Add(y *Float32_2x2) Float32_2x2 {
	return Float32_2x2{
		data: x.data.Add(y.data),
	}
}

// Div performs: c_ij = x_ij / y_ij
func (x *Float32_2x2) Div(y *Float32_2x2) Float32_2x2 {
	return Float32_2x2{
		data: x.data.Div(y.data),
	}
}

// Mul performs: c_ij = x_ij * y_ij
func (x *Float32_2x2) Mul(y *Float32_2x2) Float32_2x2 {
	return Float32_2x2{
		data: x.data.Mul(y.data),
	}
}

// Mul performs: c_ij = (x_ij * y_ij) + z_ij
func (x *Float32_2x2) MulAdd(y, z *Float32_2x2) Float32_2x2 {
	return Float32_2x2{
		data: x.data.MulAdd(y.data, z.data),
	}
}

// Max performs: c_ij = max(x_ij, y_ij)
func (x *Float32_2x2) Max(y *Float32_2x2) Float32_2x2 {
	return Float32_2x2{
		data: x.data.Max(y.data),
	}
}

// Min performs: c_ij = min(x_ij, y_ij)
func (x *Float32_2x2) Min(y *Float32_2x2) Float32_2x2 {
	return Float32_2x2{
		data: x.data.Min(y.data),
	}
}

// Sub performs: c_ij = x_ij - y_ij
func (x *Float32_2x2) Sub(y *Float32_2x2) Float32_2x2 {
	return Float32_2x2{
		data: x.data.Sub(y.data),
	}
}

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x *Float32_2x2) Scale(y *Float32_2x2) Float32_2x2 {
	return Float32_2x2{
		data: x.data.Scale(y.data),
	}
}
