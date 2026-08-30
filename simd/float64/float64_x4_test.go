package simdfloat64_test

import (
	"fmt"
	"testing"

	simdfloat64 "github.com/daanv2/go-math/simd/float64"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Math ops
func Test_Float64x4_Ops(t *testing.T) {
	dataA := []float64{1, 2, 3, 4}
	dataB := []float64{5, 6, 7, 8}
	dataC := []float64{9, 10, 11, 12}
	vec1 := simdfloat64.NewFloat64x4(dataA)
	vec2 := simdfloat64.NewFloat64x4(dataB)
	vec3 := simdfloat64.NewFloat64x4(dataC)

	t.Run("Add", func(t *testing.T) {
		result := vec1.Add(vec2)

		assert.Equal(t, []float64{6, 8, 10, 12}, result.ToSlice())
	})

	t.Run("Abs", func(t *testing.T) {
		vec := simdfloat64.NewFloat64x4([]float64{1, -2, -3, 4})
		result := vec.Abs()

		assert.Equal(t, []float64{1, 2, 3, 4}, result.ToSlice())
	})

	t.Run("Div", func(t *testing.T) {
		result := vec1.Div(vec2)

		assert.Equal(t, []float64{0.2, 0.3333333333333333, 0.42857142857142855, 0.5}, result.ToSlice())
	})

	t.Run("Mul", func(t *testing.T) {
		result := vec1.Mul(vec2)

		assert.Equal(t, []float64{5, 12, 21, 32}, result.ToSlice())
	})

	t.Run("MulAdd", func(t *testing.T) {
		result := vec1.MulAdd(vec2, vec3)

		assert.Equal(t, []float64{14, 22, 32, 44}, result.ToSlice())
	})

	t.Run("Max", func(t *testing.T) {
		dataA := []float64{1, 10, 3, 12}
		dataB := []float64{9, 2, 11, 4}

		vec1 := simdfloat64.NewFloat64x4(dataA)
		vec2 := simdfloat64.NewFloat64x4(dataB)

		result := vec1.Max(vec2)

		assert.Equal(t, []float64{9, 10, 11, 12}, result.ToSlice())
	})

	t.Run("Min", func(t *testing.T) {
		dataA := []float64{1, 10, 3, 12}
		dataB := []float64{9, 2, 11, 4}

		vec1 := simdfloat64.NewFloat64x4(dataA)
		vec2 := simdfloat64.NewFloat64x4(dataB)

		result := vec1.Min(vec2)

		assert.Equal(t, []float64{1, 2, 3, 4}, result.ToSlice())
	})

	t.Run("Sub", func(t *testing.T) {
		result := vec1.Sub(vec2)

		assert.Equal(t, []float64{-4, -4, -4, -4}, result.ToSlice())
	})

	t.Run("Scale", func(t *testing.T) {
		exponents := simdfloat64.NewFloat64x4([]float64{1, 0, 2, 1})

		result := vec1.Scale(exponents)

		assert.Equal(t, []float64{2, 2, 12, 8}, result.ToSlice())
	})
}

// Loads
func Test_Float64x4_Load(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}

	assert.Equal(t, simdfloat64.NewFloat64x4(data).ToSlice(), data[:4])
	assert.Equal(t, []float64{1, 0, 0, 0}, simdfloat64.NewFloat64x4(data[:1]).ToSlice())
	assert.Equal(t, []float64{1, 2, 3, 0}, simdfloat64.NewFloat64x4(data[:3]).ToSlice())
}

func Test_Float64x4Slice(t *testing.T) {
	data := []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9}

	result := simdfloat64.NewFloat64x4Slice(data)

	require.Len(t, result, 3)
	require.Equal(t, []float64{1, 2, 3, 4}, result[0].ToSlice())
	require.Equal(t, []float64{5, 6, 7, 8}, result[1].ToSlice())
	require.Equal(t, []float64{9, 0, 0, 0}, result[2].ToSlice())
}

// Examples

// Default example
func ExampleFloat64x4() {
	data := []float64{1, 2, 3, 4}

	v1 := simdfloat64.NewFloat64x4(data)

	fmt.Println(v1.ToSlice())

	// Output:
	// [1 2 3 4]
}
