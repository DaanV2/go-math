package simdfloats_test

import (
	"fmt"
	"testing"

	simdfloats "github.com/daanv2/go-math/pkg/simd/floats"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Math ops
func Test_Float64x8_Add(t *testing.T) {
	dataA := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	dataB := []float64{9, 10, 11, 12, 13, 14, 15, 16}

	vec1 := simdfloats.NewFloat64x8(dataA)
	vec2 := simdfloats.NewFloat64x8(dataB)

	result := vec1.Add(vec2)

	assert.Equal(t, []float64{10, 12, 14, 16, 18, 20, 22, 24}, result.ToSlice())
}

func Test_Float64x8_Sub(t *testing.T) {
	dataA := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	dataB := []float64{9, 10, 11, 12, 13, 14, 15, 16}

	vec1 := simdfloats.NewFloat64x8(dataA)
	vec2 := simdfloats.NewFloat64x8(dataB)

	result := vec1.Sub(vec2)

	assert.Equal(t, []float64{-8, -8, -8, -8, -8, -8, -8, -8}, result.ToSlice())
}

func Test_Float64x8_Mul(t *testing.T) {
	dataA := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	dataB := []float64{9, 10, 11, 12, 13, 14, 15, 16}

	vec1 := simdfloats.NewFloat64x8(dataA)
	vec2 := simdfloats.NewFloat64x8(dataB)

	result := vec1.Mul(vec2)

	assert.Equal(t, []float64{9, 20, 33, 48, 65, 84, 105, 128}, result.ToSlice())
}

func Test_Float64x8_MulAdd(t *testing.T) {
	dataA := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	dataB := []float64{9, 10, 11, 12, 13, 14, 15, 16}

	vec1 := simdfloats.NewFloat64x8(dataA)
	vec2 := simdfloats.NewFloat64x8(dataB)

	result := vec1.Add(vec2)

	assert.Equal(t, []float64{10, 12, 14, 16, 18, 20, 22, 24}, result.ToSlice())
}

// Loads
func Test_Float64x8_Load(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	assert.Equal(t, simdfloats.NewFloat64x8(data).ToSlice(), data[:8])
	assert.Equal(t, []float64{1, 0, 0, 0, 0, 0, 0, 0}, simdfloats.NewFloat64x8(data[:1]).ToSlice())
	assert.Equal(t, []float64{1, 2, 3, 4, 5, 6, 7, 0}, simdfloats.NewFloat64x8(data[:7]).ToSlice())
}

func Test_Float64x8Slice(t *testing.T) {
	data := []float64{
		1, 2, 3, 4, 5, 6, 7, 8,
		9, 10, 11, 12, 13, 14, 15, 16,
		17}

	result := simdfloats.NewFloat64x8Slice(data)

	require.Len(t, result, 3)
	require.Equal(t, []float64{1, 2, 3, 4, 5, 6, 7, 8}, result[0].ToSlice())
	require.Equal(t, []float64{9, 10, 11, 12, 13, 14, 15, 16}, result[1].ToSlice())
	require.Equal(t, []float64{17, 0, 0, 0, 0, 0, 0, 0}, result[2].ToSlice())
}

// Examples

// Default example
func ExampleFloat64x8() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	v1 := simdfloats.NewFloat64x8(data)

	fmt.Println(v1.ToSlice())

	// Output:
	// [1 2 3 4 5 6 7 8]
}
