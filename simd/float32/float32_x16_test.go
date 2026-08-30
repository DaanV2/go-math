package simdfloat32_test

import (
	"fmt"
	"testing"

	simdfloat32 "github.com/daanv2/go-math/simd/float32"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Math ops
func Test_Float32x16_Ops(t *testing.T) {
	dataA := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	dataB := []float32{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	dataC := []float32{17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	vec1 := simdfloat32.NewFloat32x16(dataA)
	vec2 := simdfloat32.NewFloat32x16(dataB)
	vec3 := simdfloat32.NewFloat32x16(dataC)

	t.Run("Add", func(t *testing.T) {
		result := vec1.Add(vec2)

		assert.Equal(t, []float32{10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40}, result.ToSlice())
	})

	t.Run("Abs", func(t *testing.T) {
		vec := simdfloat32.NewFloat32x16([]float32{1, -2, -3, 4, -5, 6, -7, 8, -9, 10, -11, 12, -13, 14, -15, 16})
		result := vec.Abs()

		assert.Equal(t, []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, result.ToSlice())
	})

	t.Run("Div", func(t *testing.T) {
		result := vec1.Div(vec2)

		expected := []float32{0.11111111, 0.2, 0.27272728, 0.33333334, 0.3846154, 0.42857143, 0.46666667, 0.5, 0.5294118, 0.5555556, 0.57894737, 0.6, 0.61904764, 0.6363636, 0.65217394, 0.6666667}

		assert.Equal(t, expected, result.ToSlice())
	})

	t.Run("Mul", func(t *testing.T) {
		result := vec1.Mul(vec2)

		assert.Equal(t, []float32{9, 20, 33, 48, 65, 84, 105, 128, 153, 180, 209, 240, 273, 308, 345, 384}, result.ToSlice())
	})

	t.Run("Neg", func(t *testing.T) {
		result := vec1.Neg()

		assert.Equal(t, []float32{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -11, -12, -13, -14, -15, -16}, result.ToSlice())
	})

	t.Run("MulAdd", func(t *testing.T) {
		result := vec1.MulAdd(vec2, vec3)

		assert.Equal(t, []float32{26, 38, 52, 68, 86, 106, 128, 152, 178, 206, 236, 268, 302, 338, 376, 416}, result.ToSlice())
	})

	t.Run("Max", func(t *testing.T) {
		dataA := []float32{9, 2, 11, 4, 13, 6, 15, 8, 17, 10, 19, 12, 21, 14, 23, 16}
		dataB := []float32{1, 10, 3, 12, 5, 14, 7, 16, 9, 18, 11, 20, 13, 22, 15, 24}

		vec1 := simdfloat32.NewFloat32x16(dataA)
		vec2 := simdfloat32.NewFloat32x16(dataB)

		result := vec1.Max(vec2)

		assert.Equal(t, []float32{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}, result.ToSlice())
	})

	t.Run("Min", func(t *testing.T) {
		dataA := []float32{9, 2, 11, 4, 13, 6, 15, 8, 17, 10, 19, 12, 21, 14, 23, 16}
		dataB := []float32{1, 10, 3, 12, 5, 14, 7, 16, 9, 18, 11, 20, 13, 22, 15, 24}

		vec1 := simdfloat32.NewFloat32x16(dataA)
		vec2 := simdfloat32.NewFloat32x16(dataB)

		result := vec1.Min(vec2)

		assert.Equal(t, []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, result.ToSlice())
	})

	t.Run("Sub", func(t *testing.T) {
		result := vec1.Sub(vec2)

		assert.Equal(t, []float32{-8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8}, result.ToSlice())
	})

	t.Run("Scale", func(t *testing.T) {
		exponents := simdfloat32.NewFloat32x16([]float32{1, 0, 2, 1, 0, 3, 1, 0, 1, 0, 2, 1, 0, 3, 1, 0})

		result := vec1.Scale(exponents)

		assert.Equal(t, []float32{2, 2, 12, 8, 5, 48, 14, 8, 18, 10, 44, 24, 13, 112, 30, 16}, result.ToSlice())
	})

	t.Run("Sqrt", func(t *testing.T) {
		result := vec1.Sqrt()

		expected := []float32{1, 1.4142135, 1.7320508, 2, 2.236068, 2.4494898, 2.6457512, 2.828427, 3, 3.1622777, 3.3166249, 3.4641016, 3.6055512, 3.7416575, 3.8729835, 4}

		assert.Equal(t, expected, result.ToSlice())
	})
}

// Loads
func Test_Float32x16_Load(t *testing.T) {
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	assert.Equal(t, simdfloat32.NewFloat32x16(data).ToSlice(), data[:16])
	assert.Equal(t, []float32{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, simdfloat32.NewFloat32x16(data[:1]).ToSlice())
	assert.Equal(t, []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0}, simdfloat32.NewFloat32x16(data[:15]).ToSlice())
}

func Test_Float32x16Slice(t *testing.T) {
	data := []float32{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
		25}

	result := simdfloat32.NewFloat32x16Slice(data)

	require.Len(t, result, 3)
	assert.Equal(t, []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, result[0].ToSlice())
	assert.Equal(t, []float32{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}, result[1].ToSlice())
	assert.Equal(t, []float32{25, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, result[2].ToSlice())
}

// Examples

// Default example
func ExampleFloat32x16() {
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	v1 := simdfloat32.NewFloat32x16(data)

	fmt.Println(v1.ToSlice())

	// Output:
	// [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]
}
