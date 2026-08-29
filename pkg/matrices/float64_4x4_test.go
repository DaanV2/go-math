package matrices_test

import (
	"testing"

	"github.com/daanv2/go-math/pkg/matrices"
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
		16 + 1, 16 + 2, 16 + 3, 16 + 4,
		16 + 5, 16 + 6, 16 + 7, 16 + 8,
		16 + 9, 16 + 10, 16 + 11, 16 + 12,
		16 + 13, 16 + 14, 16 + 15, 16 + 16,
	}
	dataC := []float64{
		32 + 1, 32 + 2, 32 + 3, 32 + 4,
		32 + 5, 32 + 6, 32 + 7, 32 + 8,
		32 + 9, 32 + 10, 32 + 11, 32 + 12,
		32 + 13, 32 + 14, 32 + 15, 32 + 16,
	}
	mat1 := matrices.NewFloat64_4x4(dataA)
	mat2 := matrices.NewFloat64_4x4(dataB)
	mat3 := matrices.NewFloat64_4x4(dataC)

	t.Run("Add", func(t *testing.T) {
		result := mat1.Add(&mat2)

		assert.Equal(t, []float64{10, 12, 14, 16, 18, 20, 22, 24}, result.ToSlice())
	})

	t.Run("Abs", func(t *testing.T) {
		mat := matrices.NewFloat64_4x4([]float64{1, -2, -3, 4, -5, 6, -7, 8})
		result := mat.Abs()

		assert.Equal(t, []float64{1, 2, 3, 4, 5, 6, 7, 8}, result.ToSlice())
	})

	t.Run("Div", func(t *testing.T) {
		result := mat1.Div(&mat2)

		assert.Equal(t, []float64{
			0.1111111111111111,
			0.2,
			0.2727272727272727,
			0.3333333333333333,
			0.38461538461538464,
			0.42857142857142855,
			0.4666666666666667,
			0.5}, result.ToSlice())
	})

	t.Run("Mul", func(t *testing.T) {
		result := mat1.Mul(&mat2)

		assert.Equal(t, []float64{9, 20, 33, 48, 65, 84, 105, 128}, result.ToSlice())
	})

	t.Run("MulAdd", func(t *testing.T) {
		result := mat1.MulAdd(&mat2, &mat3)

		assert.Equal(t, []float64{26, 38, 52, 68, 86, 106, 128, 152}, result.ToSlice())
	})

	t.Run("Max", func(t *testing.T) {
		dataA := []float64{1, 10, 3, 12, 5, 14, 7, 16}
		dataB := []float64{9, 2, 11, 4, 13, 6, 15, 8}

		vec1 := matrices.NewFloat64_4x4(dataA)
		vec2 := matrices.NewFloat64_4x4(dataB)

		result := vec1.Max(&vec2)

		assert.Equal(t, []float64{9, 10, 11, 12, 13, 14, 15, 16}, result.ToSlice())
	})

	t.Run("Min", func(t *testing.T) {
		dataA := []float64{1, 10, 3, 12, 5, 14, 7, 16}
		dataB := []float64{9, 2, 11, 4, 13, 6, 15, 8}

		vec1 := matrices.NewFloat64_4x4(dataA)
		vec2 := matrices.NewFloat64_4x4(dataB)

		result := vec1.Min(&vec2)

		assert.Equal(t, []float64{1, 2, 3, 4, 5, 6, 7, 8}, result.ToSlice())
	})

	t.Run("Sub", func(t *testing.T) {
		result := mat1.Sub(&mat2)

		assert.Equal(t, []float64{-8, -8, -8, -8, -8, -8, -8, -8}, result.ToSlice())
	})

	t.Run("Scale", func(t *testing.T) {
		exponents := matrices.NewFloat64_4x4([]float64{1, 0, 2, 1, 0, 3, 1, 0})

		result := mat1.Scale(&exponents)

		assert.Equal(t, []float64{2, 2, 12, 8, 5, 48, 14, 8}, result.ToSlice())
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

	assert.Equal(t, mat.ToSlice(), data[:8])

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
