package simdfloat32_test

import (
	"fmt"
	"testing"

	simdfloat32 "github.com/daanv2/go-math/simd/float32"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Math ops
func Test_Float32x8_Ops(t *testing.T) {
	dataA := []float32{1, 2, 3, 4, 5, 6, 7, 8}
	dataB := []float32{9, 10, 11, 12, 13, 14, 15, 16}
	dataC := []float32{17, 18, 19, 20, 21, 22, 23, 24}
	vec1 := simdfloat32.NewFloat32x8(dataA)
	vec2 := simdfloat32.NewFloat32x8(dataB)
	vec3 := simdfloat32.NewFloat32x8(dataC)

	t.Run("Add", func(t *testing.T) {
		result := vec1.Add(vec2)

		assert.Equal(t, []float32{10, 12, 14, 16, 18, 20, 22, 24}, result.ToSlice())
	})

	t.Run("Abs", func(t *testing.T) {
		vec := simdfloat32.NewFloat32x8([]float32{1, -2, -3, 4, -5, 6, -7, 8})
		result := vec.Abs()

		assert.Equal(t, []float32{1, 2, 3, 4, 5, 6, 7, 8}, result.ToSlice())
	})

	t.Run("Div", func(t *testing.T) {
		result := vec1.Div(vec2)

		assert.Equal(t, []float32{
			0.1111111111111111,
			0.2,
			0.2727272727272727,
			0.3333333333333333,
			0.38461538461538432,
			0.42857142857142855,
			0.4666666666666667,
			0.5}, result.ToSlice())
	})

	t.Run("Mul", func(t *testing.T) {
		result := vec1.Mul(vec2)

		assert.Equal(t, []float32{9, 20, 33, 48, 65, 84, 105, 128}, result.ToSlice())
	})

	t.Run("Neg", func(t *testing.T) {
		result := vec1.Neg()

		assert.Equal(t, []float32{-1, -2, -3, -4, -5, -6, -7, -8}, result.ToSlice())
	})

	t.Run("MulAdd", func(t *testing.T) {
		result := vec1.MulAdd(vec2, vec3)

		assert.Equal(t, []float32{26, 38, 52, 68, 86, 106, 128, 152}, result.ToSlice())
	})

	t.Run("Max", func(t *testing.T) {
		dataA := []float32{1, 10, 3, 12, 5, 14, 7, 16}
		dataB := []float32{9, 2, 11, 4, 13, 6, 15, 8}

		vec1 := simdfloat32.NewFloat32x8(dataA)
		vec2 := simdfloat32.NewFloat32x8(dataB)

		result := vec1.Max(vec2)

		assert.Equal(t, []float32{9, 10, 11, 12, 13, 14, 15, 16}, result.ToSlice())
	})

	t.Run("Min", func(t *testing.T) {
		dataA := []float32{1, 10, 3, 12, 5, 14, 7, 16}
		dataB := []float32{9, 2, 11, 4, 13, 6, 15, 8}

		vec1 := simdfloat32.NewFloat32x8(dataA)
		vec2 := simdfloat32.NewFloat32x8(dataB)

		result := vec1.Min(vec2)

		assert.Equal(t, []float32{1, 2, 3, 4, 5, 6, 7, 8}, result.ToSlice())
	})

	t.Run("Sub", func(t *testing.T) {
		result := vec1.Sub(vec2)

		assert.Equal(t, []float32{-8, -8, -8, -8, -8, -8, -8, -8}, result.ToSlice())
	})

	t.Run("Scale", func(t *testing.T) {
		exponents := simdfloat32.NewFloat32x8([]float32{1, 0, 2, 1, 0, 3, 1, 0})

		result := vec1.Scale(exponents)

		assert.Equal(t, []float32{2, 2, 12, 8, 5, 48, 14, 8}, result.ToSlice())
	})

	t.Run("Sqrt", func(t *testing.T) {
		result := vec1.Sqrt()

		assert.Equal(t, []float32{1, 1.4142135, 1.7320508, 2, 2.236068, 2.4494898, 2.6457512, 2.828427}, result.ToSlice())
	})
}

// Loads
func Test_Float32x8_Load(t *testing.T) {
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9}

	assert.Equal(t, simdfloat32.NewFloat32x8(data).ToSlice(), data[:8])
	assert.Equal(t, []float32{1, 0, 0, 0, 0, 0, 0, 0}, simdfloat32.NewFloat32x8(data[:1]).ToSlice())
	assert.Equal(t, []float32{1, 2, 3, 4, 5, 6, 7, 0}, simdfloat32.NewFloat32x8(data[:7]).ToSlice())
}

func Test_Float32x8Slice(t *testing.T) {
	data := []float32{
		1, 2, 3, 4, 5, 6, 7, 8,
		9, 10, 11, 12, 13, 14, 15, 16,
		17}

	result := simdfloat32.NewFloat32x8Slice(data)

	require.Len(t, result, 3)
	require.Equal(t, []float32{1, 2, 3, 4, 5, 6, 7, 8}, result[0].ToSlice())
	require.Equal(t, []float32{9, 10, 11, 12, 13, 14, 15, 16}, result[1].ToSlice())
	require.Equal(t, []float32{17, 0, 0, 0, 0, 0, 0, 0}, result[2].ToSlice())
}

// Examples

// Default example
func ExampleFloat32x8() {
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9}

	v1 := simdfloat32.NewFloat32x8(data)

	fmt.Println(v1.ToSlice())

	// Output:
	// [1 2 3 4 5 6 7 8]
}
