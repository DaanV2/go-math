package simdfloats_test

import (
	"fmt"
	"testing"

	simdfloats "github.com/daanv2/go-math/pkg/simd/floats"
	"github.com/stretchr/testify/assert"
)

func Test_Float64x8_Add(t *testing.T) {
	dataA := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	dataB := []float64{9, 10, 11, 12, 13, 14, 15, 16}

	vec1 := simdfloats.NewFloat64x8(dataA)
	vec2 := simdfloats.NewFloat64x8(dataB)

	result := vec1.Add(vec2)

	assert.Equal(t, []float64{10, 12, 14, 16, 18, 20, 22, 24}, result.ToSlice())
}

func Test_Float64x8_Load(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	assert.Equal(t, simdfloats.NewFloat64x8(data).ToSlice(), data[:8])
	assert.Equal(t, []float64{1, 0, 0, 0, 0, 0, 0, 0}, simdfloats.NewFloat64x8(data[:1]).ToSlice())
	assert.Equal(t, []float64{1, 2, 3, 4, 5, 6, 7, 0}, simdfloats.NewFloat64x8(data[:7]).ToSlice())
}

func ExampleFloat64x8() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	v1 := simdfloats.NewFloat64x8(data)

	fmt.Println(v1.ToSlice())

	// Output:
	// [1 2 3 4 5 6 7 8]
}
