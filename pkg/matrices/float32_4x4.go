package matrices

import simdfloats "github.com/daanv2/go-math/pkg/simd/floats"

type Float32_4x4 struct {
	data0 simdfloats.Float32x8
	data1 simdfloats.Float32x8
}

func NewFloat32_4x4(data []float32) Float32_4x4 {
	var result Float32_4x4

	if len(data) > 8 {
		result.data0 = simdfloats.NewFloat32x8(data[:8])
		result.data1 = simdfloats.NewFloat32x8(data[8:])
	} else {
		result.data0 = simdfloats.NewFloat32x8(data)
	}

	return result
}

func (m *Float32_4x4) Width() int  { return 4 }
func (m *Float32_4x4) Height() int { return 4 }
func (m *Float32_4x4) Len() int    { return 4 * 4 }

func (m *Float32_4x4) ToSlice() []float32 {
	result := make([]float32, m.Len())

	m.Store(result)

	return result
}

func (m *Float32_4x4) Store(receiver []float32) {
	m.data0.Store(receiver)
	if len(receiver) > 8 {
		m.data1.Store(receiver[8:])
	}
}

// Abs performs: c_ij = abs(x_ij)
func (x *Float32_4x4) Abs() Float32_4x4 {
	return Float32_4x4{
		data0: x.data0.Abs(),
		data1: x.data1.Abs(),
	}
}

// Add performs: c_ij = x_ij + y_ij
func (x *Float32_4x4) Add(y *Float32_4x4) Float32_4x4 {
	return Float32_4x4{
		data0: x.data0.Add(y.data0),
		data1: x.data1.Add(y.data1),
	}
}

// Div performs: c_ij = x_ij / y_ij
func (x *Float32_4x4) Div(y *Float32_4x4) Float32_4x4 {
	return Float32_4x4{
		data0: x.data0.Div(y.data0),
		data1: x.data1.Div(y.data1),
	}
}

// Mul performs: c_ij = x_ij * y_ij
func (x *Float32_4x4) Mul(y *Float32_4x4) Float32_4x4 {
	return Float32_4x4{
		data0: x.data0.Mul(y.data0),
		data1: x.data1.Mul(y.data1),
	}
}

// Mul performs: c_ij = (x_ij * y_ij) + z_ij
func (x *Float32_4x4) MulAdd(y, z *Float32_4x4) Float32_4x4 {
	return Float32_4x4{
		data0: x.data0.MulAdd(y.data0, z.data0),
		data1: x.data1.MulAdd(y.data1, z.data1),
	}
}

// Max performs: c_ij = max(x_ij, y_ij)
func (x *Float32_4x4) Max(y *Float32_4x4) Float32_4x4 {
	return Float32_4x4{
		data0: x.data0.Max(y.data0),
		data1: x.data1.Max(y.data1),
	}
}

// Min performs: c_ij = min(x_ij, y_ij)
func (x *Float32_4x4) Min(y *Float32_4x4) Float32_4x4 {
	return Float32_4x4{
		data0: x.data0.Min(y.data0),
		data1: x.data1.Min(y.data1),
	}
}

// Sub performs: c_ij = x_ij - y_ij
func (x *Float32_4x4) Sub(y *Float32_4x4) Float32_4x4 {
	return Float32_4x4{
		data0: x.data0.Sub(y.data0),
		data1: x.data1.Sub(y.data1),
	}
}

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x *Float32_4x4) Scale(y *Float32_4x4) Float32_4x4 {
	return Float32_4x4{
		data0: x.data0.Scale(y.data0),
		data1: x.data1.Scale(y.data1),
	}
}
