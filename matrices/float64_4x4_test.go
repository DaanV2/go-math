package matrices_test

import (
	"testing"

	"github.com/daanv2/go-math/matrices"
	"github.com/stretchr/testify/assert"
)

// Math ops
func Test_Float64_4x4_Ops(t *testing.T) {
	dataA := []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}
	dataB := []float64{
		17, 18, 19, 20,
		21, 22, 23, 24,
		25, 26, 27, 28,
		29, 30, 31, 32,
	}
	dataC := []float64{
		33, 34, 35, 36,
		37, 38, 39, 40,
		41, 42, 43, 44,
		46, 47, 48, 49,
	}
	mat1 := matrices.NewFloat64_4x4(dataA)
	mat2 := matrices.NewFloat64_4x4(dataB)
	mat3 := matrices.NewFloat64_4x4(dataC)

	t.Run("Add", func(t *testing.T) {
		result := mat1.Add(&mat2)

		assert.Equal(t, []float64{
			18, 20, 22, 24,
			26, 28, 30, 32,
			34, 36, 38, 40,
			42, 44, 46, 48,
		}, result.ToSlice())
	})

	t.Run("Abs", func(t *testing.T) {
		mat := matrices.NewFloat64_4x4([]float64{
			1, -2, -3, 4,
			-5, 6, -7, 8,
			9, -10, 11, -12,
			-13, 14, 15, -16,
		})
		result := mat.Abs()

		assert.Equal(t, []float64{
			1, 2, 3, 4,
			5, 6, 7, 8,
			9, 10, 11, 12,
			13, 14, 15, 16,
		}, result.ToSlice())
	})

	t.Run("Div", func(t *testing.T) {
		result := mat1.Div(&mat2)

		assert.Equal(t, []float64{
			0.058823529411764705, 0.1111111111111111, 0.15789473684210525, 0.2,
			0.23809523809523808, 0.2727272727272727, 0.30434782608695654, 0.3333333333333333,
			0.36, 0.38461538461538464, 0.4074074074074074, 0.42857142857142855,
			0.4482758620689655, 0.4666666666666667, 0.4838709677419355, 0.5,
		}, result.ToSlice())
	})

	t.Run("Mul", func(t *testing.T) {
		result := mat1.Mul(&mat2)

		assert.Equal(t, []float64{
			17, 36, 57, 80,
			105, 132, 161, 192,
			225, 260, 297, 336,
			377, 420, 465, 512,
		}, result.ToSlice())
	})

	t.Run("MulAdd", func(t *testing.T) {
		result := mat1.MulAdd(&mat2, &mat3)

		assert.Equal(t, []float64{
			50, 70, 92, 116,
			142, 170, 200, 232,
			266, 302, 340, 380,
			423, 467, 513, 561,
		}, result.ToSlice())
	})

	t.Run("Max", func(t *testing.T) {
		dataA := []float64{
			1, 18, 3, 20,
			21, 6, 23, 8,
			9, 26, 11, 28,
			29, 14, 31, 16,
		}
		dataB := []float64{
			17, 2, 19, 4,
			5, 22, 7, 24,
			25, 10, 27, 12,
			13, 30, 15, 32,
		}

		vec1 := matrices.NewFloat64_4x4(dataA)
		vec2 := matrices.NewFloat64_4x4(dataB)

		result := vec1.Max(&vec2)

		assert.Equal(t, []float64{
			17, 18, 19, 20,
			21, 22, 23, 24,
			25, 26, 27, 28,
			29, 30, 31, 32,
		}, result.ToSlice())
	})

	t.Run("Min", func(t *testing.T) {
		dataA := []float64{
			1, 18, 3, 20,
			21, 6, 23, 8,
			9, 26, 11, 28,
			29, 14, 31, 16,
		}
		dataB := []float64{
			17, 2, 19, 4,
			5, 22, 7, 24,
			25, 10, 27, 12,
			13, 30, 15, 32,
		}

		vec1 := matrices.NewFloat64_4x4(dataA)
		vec2 := matrices.NewFloat64_4x4(dataB)

		result := vec1.Min(&vec2)

		assert.Equal(t, []float64{
			1, 2, 3, 4,
			5, 6, 7, 8,
			9, 10, 11, 12,
			13, 14, 15, 16,
		}, result.ToSlice())
	})

	t.Run("Sub", func(t *testing.T) {
		result := mat1.Sub(&mat2)

		assert.Equal(t, []float64{
			-16, -16, -16, -16,
			-16, -16, -16, -16,
			-16, -16, -16, -16,
			-16, -16, -16, -16,
		}, result.ToSlice())
	})

	t.Run("Scale", func(t *testing.T) {
		exponents := matrices.NewFloat64_4x4([]float64{
			1, 0, -1, 2,
			0, 1, 2, -1,
			-1, 2, 1, 0,
			2, -1, 0, 1,
		})

		result := mat1.Scale(&exponents)

		assert.Equal(t, []float64{
			2, 2, 1.5, 16,
			5, 12, 28, 4,
			4.5, 40, 22, 12,
			52, 7, 15, 32,
		}, result.ToSlice())
	})
}

// Loads
func Test_Float64_4x4_Load(t *testing.T) {
	data := []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}
	mat := matrices.NewFloat64_4x4(data)

	assert.Equal(t, data, mat.ToSlice())

	mat = matrices.NewFloat64_4x4(data[:1])

	assert.Equal(t, []float64{
		1, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	}, mat.ToSlice())

	mat = matrices.NewFloat64_4x4(data[:15])

	assert.Equal(t, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 0,
	}, mat.ToSlice())
}
