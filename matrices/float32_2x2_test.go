package matrices_test

import (
	"testing"

	"github.com/daanv2/go-math/matrices"
	"github.com/stretchr/testify/assert"
)

// Math ops
func Test_Float32_2x2_Ops(t *testing.T) {
	dataA := []float32{
		1, 2,
		3, 4,
	}
	dataB := []float32{
		17, 18,
		19, 20,
	}
	dataC := []float32{
		33, 34,
		37, 38,
	}
	mat1 := matrices.NewFloat32_2x2(dataA)
	mat2 := matrices.NewFloat32_2x2(dataB)
	mat3 := matrices.NewFloat32_2x2(dataC)

	t.Run("Add", func(t *testing.T) {
		result := mat1.Add(&mat2)

		assert.Equal(t, []float32{
			18, 20,
			22, 24,
		}, result.ToSlice())
	})

	t.Run("Abs", func(t *testing.T) {
		mat := matrices.NewFloat32_2x2([]float32{
			1, -2, -3, 4,
			-5, 6, -7, 8,
			9, -10, 11, -12,
			-13, 14, 15, -16,
		})
		result := mat.Abs()

		assert.Equal(t, []float32{
			1, 2,
			3, 4,
		}, result.ToSlice())
	})

	t.Run("Div", func(t *testing.T) {
		result := mat1.Div(&mat2)

		assert.Equal(t, []float32{
			0.058823529411732705, 0.1111111111111111,
			0.15789473684210525, 0.2,
		}, result.ToSlice())
	})

	t.Run("Mul", func(t *testing.T) {
		result := mat1.Mul(&mat2)

		assert.Equal(t, []float32{
			17, 36,
			57, 80,
		}, result.ToSlice())
	})

	t.Run("MulAdd", func(t *testing.T) {
		result := mat1.MulAdd(&mat2, &mat3)

		assert.Equal(t, []float32{
			50, 70, 94, 118,
		}, result.ToSlice())
	})

	t.Run("Max", func(t *testing.T) {
		dataA := []float32{
			1, 18,
			3, 20,
		}
		dataB := []float32{
			17, 2,
			19, 4,
		}

		vec1 := matrices.NewFloat32_2x2(dataA)
		vec2 := matrices.NewFloat32_2x2(dataB)

		result := vec1.Max(&vec2)

		assert.Equal(t, []float32{
			17, 18,
			19, 20,
		}, result.ToSlice())
	})

	t.Run("Min", func(t *testing.T) {
		dataA := []float32{
			1, 18, 3, 20,
			21, 6, 23, 8,
			9, 26, 11, 28,
			29, 14, 31, 16,
		}
		dataB := []float32{
			17, 2, 19, 4,
			5, 22, 7, 24,
			25, 10, 27, 12,
			13, 30, 15, 32,
		}

		vec1 := matrices.NewFloat32_2x2(dataA)
		vec2 := matrices.NewFloat32_2x2(dataB)

		result := vec1.Min(&vec2)

		assert.Equal(t, []float32{
			1, 2,
			3, 4,
		}, result.ToSlice())
	})

	t.Run("Sub", func(t *testing.T) {
		result := mat1.Sub(&mat2)

		assert.Equal(t, []float32{
			-16, -16,
			-16, -16,
		}, result.ToSlice())
	})

	t.Run("Scale", func(t *testing.T) {
		exponents := matrices.NewFloat32_2x2([]float32{
			1, 0, -1, 2,
			0, 1, 2, -1,
			-1, 2, 1, 0,
			2, -1, 0, 1,
		})

		result := mat1.Scale(&exponents)

		assert.Equal(t, []float32{
			2, 2,
			1.5, 16,
		}, result.ToSlice())
	})
}

// Loads
func Test_Float32_2x2_Load(t *testing.T) {
	data := []float32{
		1, 2,
		3, 4,
	}
	mat := matrices.NewFloat32_2x2(data)

	assert.Equal(t, data, mat.ToSlice())

	mat = matrices.NewFloat32_2x2(data[:1])

	assert.Equal(t, []float32{
		1, 0,
		0, 0,
	}, mat.ToSlice())

	mat = matrices.NewFloat32_2x2(data[:3])

	assert.Equal(t, []float32{
		1, 2,
		3, 0,
	}, mat.ToSlice())
}
