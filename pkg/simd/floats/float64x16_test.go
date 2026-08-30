package simdfloats_test

import (
	"fmt"
	"testing"

	simdfloats "github.com/daanv2/go-math/pkg/simd/floats"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Math ops
func Test_Float64x16_Ops(t *testing.T) {
	dataA := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	dataB := []float64{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	dataC := []float64{17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	vec1 := simdfloats.NewFloat64x16(dataA)
	vec2 := simdfloats.NewFloat64x16(dataB)
	vec3 := simdfloats.NewFloat64x16(dataC)

	t.Run("Add", func(t *testing.T) {
		result := vec1.Add(vec2)

		assert.Equal(t, []float64{10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40}, result.ToSlice())
	})

	t.Run("Abs", func(t *testing.T) {
		vec := simdfloats.NewFloat64x16([]float64{1, -2, -3, 4, -5, 6, -7, 8, -9, 10, -11, 12, -13, 14, -15, 16})
		result := vec.Abs()

		assert.Equal(t, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, result.ToSlice())
	})

	t.Run("Div", func(t *testing.T) {
		result := vec1.Div(vec2)

		expected := []float64{0.1111111111111111, 0.2, 0.2727272727272727, 0.3333333333333333, 0.38461538461538464, 0.42857142857142855, 0.4666666666666667, 0.5, 0.5294117647058824, 0.5555555555555556, 0.5789473684210527, 0.6, 0.6190476190476191, 0.6363636363636364, 0.6521739130434783, 0.6666666666666666}

		assert.Equal(t, expected, result.ToSlice())
	})

	t.Run("Mul", func(t *testing.T) {
		result := vec1.Mul(vec2)

		assert.Equal(t, []float64{9, 20, 33, 48, 65, 84, 105, 128, 153, 180, 209, 240, 273, 308, 345, 384}, result.ToSlice())
	})

	t.Run("Neg", func(t *testing.T) {
		result := vec1.Neg()

		assert.Equal(t, []float64{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -11, -12, -13, -14, -15, -16}, result.ToSlice())
	})

	t.Run("MulAdd", func(t *testing.T) {
		result := vec1.MulAdd(vec2, vec3)

		assert.Equal(t, []float64{26, 38, 52, 68, 86, 106, 128, 152, 178, 206, 236, 268, 302, 338, 376, 416}, result.ToSlice())
	})

	t.Run("Max", func(t *testing.T) {
		dataA := []float64{9, 2, 11, 4, 13, 6, 15, 8, 17, 10, 19, 12, 21, 14, 23, 16}
		dataB := []float64{1, 10, 3, 12, 5, 14, 7, 16, 9, 18, 11, 20, 13, 22, 15, 24}

		vec1 := simdfloats.NewFloat64x16(dataA)
		vec2 := simdfloats.NewFloat64x16(dataB)

		result := vec1.Max(vec2)

		assert.Equal(t, []float64{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}, result.ToSlice())
	})

	t.Run("Min", func(t *testing.T) {
		dataA := []float64{9, 2, 11, 4, 13, 6, 15, 8, 17, 10, 19, 12, 21, 14, 23, 16}
		dataB := []float64{1, 10, 3, 12, 5, 14, 7, 16, 9, 18, 11, 20, 13, 22, 15, 24}

		vec1 := simdfloats.NewFloat64x16(dataA)
		vec2 := simdfloats.NewFloat64x16(dataB)

		result := vec1.Min(vec2)

		assert.Equal(t, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, result.ToSlice())
	})

	t.Run("Sub", func(t *testing.T) {
		result := vec1.Sub(vec2)

		assert.Equal(t, []float64{-8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8}, result.ToSlice())
	})

	t.Run("Scale", func(t *testing.T) {
		exponents := simdfloats.NewFloat64x16([]float64{1, 0, 2, 1, 0, 3, 1, 0, 1, 0, 2, 1, 0, 3, 1, 0})

		result := vec1.Scale(exponents)

		assert.Equal(t, []float64{2, 2, 12, 8, 5, 48, 14, 8, 18, 10, 44, 24, 13, 112, 30, 16}, result.ToSlice())
	})

	t.Run("Sqrt", func(t *testing.T) {
		result := vec1.Sqrt()

		expected := []float64{1, 1.4142135623730951, 1.7320508075688772, 2, 2.23606797749979, 2.449489742783178, 2.6457513110645907, 2.8284271247461903, 3, 3.1622776601683795, 3.3166247903554, 3.4641016151377544, 3.605551275463989, 3.7416573867739413, 3.872983346207417, 4}

		assert.Equal(t, expected, result.ToSlice())
	})
}

// Loads
func Test_Float64x16_Load(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	assert.Equal(t, simdfloats.NewFloat64x16(data).ToSlice(), data[:16])
	assert.Equal(t, []float64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, simdfloats.NewFloat64x16(data[:1]).ToSlice())
	assert.Equal(t, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0}, simdfloats.NewFloat64x16(data[:15]).ToSlice())
}

func Test_Float64x16Slice(t *testing.T) {
	data := []float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
		25}

	result := simdfloats.NewFloat64x16Slice(data)

	require.Len(t, result, 3)
	assert.Equal(t, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, result[0].ToSlice())
	assert.Equal(t, []float64{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}, result[1].ToSlice())
	assert.Equal(t, []float64{25, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, result[2].ToSlice())
}

// Examples

// Default example
func ExampleFloat64x16() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	v1 := simdfloats.NewFloat64x16(data)

	fmt.Println(v1.ToSlice())

	// Output:
	// [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]
}
