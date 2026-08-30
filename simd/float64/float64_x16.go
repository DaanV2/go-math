// nolint:gocritic // TODO Something for later

package simdfloat64

import (
	"math"
)

const float64_x16_len = 16

type Float64x16 struct {
	data0 Float64x8
	data1 Float64x8
}

func NewFloat64x16(data []float64) Float64x16 {
	var result Float64x16

	result.data0 = NewFloat64x8(data)
	if len(data) > 8 { // Read atleast 4 points, so there should be more
		result.data1 = NewFloat64x8(data[(float64_x16_len / 2):])
	}

	return result
}

// NewFloat64x16Boardcast returns a Float64x16 with every lane set to value.
func NewFloat64x16Boardcast(value float64) Float64x16 {
	return Float64x16{
		data0: NewFloat64x8Boardcast(value),
		data1: NewFloat64x8Boardcast(value),
	}
}

// NewFloat64x16Slice takes the given data and transfer them in a simd layout,
// element n0, n1 ... n7 are taken a stored in the first [Float64x16], next float64_x16_len elements in the next [Float64x16] etc.
// If not multiple of float64_x16_len, the last few numbers are put into the last lowest [Float64x16], with the rest padded to 0
func NewFloat64x16Slice(data []float64) []Float64x16 {
	l := len(data) / float64_x16_len
	if (len(data) % float64_x16_len) != 0 { // Do we need more room for padding
		l += 1
	}

	result := make([]Float64x16, 0, l)

	for len(data) > 0 {
		v := NewFloat64x16(data)
		result = append(result, v)

		if len(data) >= float64_x16_len {
			data = data[float64_x16_len:]
		} else {
			data = nil
		}
	}

	return result
}

func (x Float64x16) ToSlice() []float64 {
	result := make([]float64, float64_x16_len)
	x.Store(result)

	return result
}

func (x Float64x16) Len() int {
	return float64_x16_len
}

func (x Float64x16) Store(receiver []float64) {
	switch {
	case len(receiver) == float64_x16_len:
		x.data0.Store(receiver[:(float64_x16_len / 2)])
		x.data1.Store(receiver[(float64_x16_len / 2):])
	case len(receiver) > (float64_x16_len / 2):
		x.data0.Store(receiver[:(float64_x16_len / 2)])
		x.data1.Store(receiver[(float64_x16_len / 2):])
	default:
		x.data0.Store(receiver)
	}
}

// Abs returns the absolute values of the elements of x
func (x Float64x16) Abs() Float64x16 {
	return Float64x16{
		data0: x.data0.Abs(),
		data1: x.data1.Abs(),
	}
}

// Add performs a fused: x + y.
func (x Float64x16) Add(y Float64x16) Float64x16 {
	return Float64x16{
		data0: x.data0.Add(y.data0),
		data1: x.data1.Add(y.data1),
	}
}

// Div performs a fused: x / y.
func (x Float64x16) Div(y Float64x16) Float64x16 {
	return Float64x16{
		data0: x.data0.Div(y.data0),
		data1: x.data1.Div(y.data1),
	}
}

// Mul performs a fused: x * y.
func (x Float64x16) Mul(y Float64x16) Float64x16 {
	return Float64x16{
		data0: x.data0.Mul(y.data0),
		data1: x.data1.Mul(y.data1),
	}
}

// MulAdd performs a fused: (x * y) + z.
func (x Float64x16) MulAdd(y, z Float64x16) Float64x16 {
	return Float64x16{
		data0: x.data0.MulAdd(y.data0, z.data0),
		data1: x.data1.MulAdd(y.data1, z.data1),
	}
}

// Max computes the maximum of each pair of corresponding elements in x and y.
func (x Float64x16) Max(y Float64x16) Float64x16 {
	return Float64x16{
		data0: x.data0.Max(y.data0),
		data1: x.data1.Max(y.data1),
	}
}

// Min computes the minimum of each pair of corresponding elements in x and y.
func (x Float64x16) Min(y Float64x16) Float64x16 {
	return Float64x16{
		data0: x.data0.Min(y.data0),
		data1: x.data1.Min(y.data1),
	}
}

// Neg returns the negation of the elements of x
func (x Float64x16) Neg() Float64x16 {
	return Float64x16{
		data0: x.data0.Neg(),
		data1: x.data1.Neg(),
	}
}

// Scale multiplies each element of x by 2 raised to the power of the floor of the corresponding element in y.
func (x Float64x16) Scale(y Float64x16) Float64x16 {
	var result [float64_x16_len]float64
	var vx [float64_x16_len]float64
	var vy [float64_x16_len]float64

	x.Store(vx[:])
	y.Store(vy[:])

	for i := range vx {
		result[i] = vx[i] * math.Pow(2, vy[i])
	}

	return NewFloat64x16(result[:])
}

// Sub performs a fused: x - y.
func (x Float64x16) Sub(y Float64x16) Float64x16 {
	return Float64x16{
		data0: x.data0.Sub(y.data0),
		data1: x.data1.Sub(y.data1),
	}
}

// Sqrt computes the square root of each element.
func (x Float64x16) Sqrt() Float64x16 {
	return Float64x16{
		data0: x.data0.Sqrt(),
		data1: x.data1.Sqrt(),
	}
}
