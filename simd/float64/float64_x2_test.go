package simdfloat64_test

import (
	"fmt"
	"testing"

	simdfloat64 "github.com/daanv2/go-math/simd/float64"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Math ops
func Test_Float64x2_Ops(t *testing.T) {
	dataA := []float64{1, 2}
	dataB := []float64{5, 6}
	dataC := []float64{9, 10}
	vec1 := simdfloat64.NewFloat64x2(dataA)
	vec2 := simdfloat64.NewFloat64x2(dataB)
	vec3 := simdfloat64.NewFloat64x2(dataC)

	t.Run("Add", func(t *testing.T) {
		result := vec1.Add(vec2)

		assert.Equal(t, []float64{6, 8}, result.ToSlice())
	})

	t.Run("Abs", func(t *testing.T) {
		vec := simdfloat64.NewFloat64x2([]float64{1, -2})
		result := vec.Abs()

		assert.Equal(t, []float64{1, 2}, result.ToSlice())
	})

	t.Run("Div", func(t *testing.T) {
		result := vec1.Div(vec2)

		assert.Equal(t, []float64{0.2, 0.3333333333333333}, result.ToSlice())
	})

	t.Run("Mul", func(t *testing.T) {
		result := vec1.Mul(vec2)

		assert.Equal(t, []float64{5, 12}, result.ToSlice())
	})

	t.Run("MulAdd", func(t *testing.T) {
		result := vec1.MulAdd(vec2, vec3)

		assert.Equal(t, []float64{14, 22}, result.ToSlice())
	})

	t.Run("Max", func(t *testing.T) {
		dataA := []float64{1, 10}
		dataB := []float64{9, 2}

		vec1 := simdfloat64.NewFloat64x2(dataA)
		vec2 := simdfloat64.NewFloat64x2(dataB)

		result := vec1.Max(vec2)

		assert.Equal(t, []float64{9, 10}, result.ToSlice())
	})

	t.Run("Min", func(t *testing.T) {
		dataA := []float64{1, 10}
		dataB := []float64{9, 2}

		vec1 := simdfloat64.NewFloat64x2(dataA)
		vec2 := simdfloat64.NewFloat64x2(dataB)

		result := vec1.Min(vec2)

		assert.Equal(t, []float64{1, 2}, result.ToSlice())
	})

	t.Run("Sub", func(t *testing.T) {
		result := vec1.Sub(vec2)

		assert.Equal(t, []float64{-4, -4}, result.ToSlice())
	})

	t.Run("Scale", func(t *testing.T) {
		exponents := simdfloat64.NewFloat64x2([]float64{1, 0})

		result := vec1.Scale(exponents)

		assert.Equal(t, []float64{2, 2}, result.ToSlice())
	})
}

// Loads
func Test_Float64x2_Load(t *testing.T) {
	data := []float64{1, 2, 3}

	assert.Equal(t, simdfloat64.NewFloat64x2(data).ToSlice(), data[:2])
	assert.Equal(t, []float64{1, 0}, simdfloat64.NewFloat64x2(data[:1]).ToSlice())
}

func Test_Float64x2Slice(t *testing.T) {
	data := []float64{
		1, 2,
		5, 6,
		9}

	result := simdfloat64.NewFloat64x2Slice(data)

	require.Len(t, result, 3)
	require.Equal(t, []float64{1, 2}, result[0].ToSlice())
	require.Equal(t, []float64{5, 6}, result[1].ToSlice())
	require.Equal(t, []float64{9, 0}, result[2].ToSlice())
}

// Examples

// Default example
func ExampleFloat64x2() {
	data := []float64{1, 2}

	v1 := simdfloat64.NewFloat64x2(data)

	fmt.Println(v1.ToSlice())

	// Output:
	// [1 2]
}
