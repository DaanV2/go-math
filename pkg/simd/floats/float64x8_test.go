package simdfloats_test

import (
	"fmt"
	"testing"

	simdfloats "github.com/daanv2/go-math/pkg/simd/floats"
	"github.com/stretchr/testify/assert"
)

func Test_Float64x8_Add(t *testing.T) {

}

func Test_Float64x8_Load(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	assert.Equal(t, simdfloats.NewFloat64x8(data).ToSlice(), data[:8])
	assert.Equal(t, simdfloats.NewFloat64x8(data[:1]).ToSlice(), data[:1])
	assert.Equal(t, simdfloats.NewFloat64x8(data[:7]).ToSlice(), data[:7])
}

func ExampleFloat64x8() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	v1 := simdfloats.NewFloat64x8(data)

	fmt.Println(v1.ToSlice())

	// Output:
	// [1 2 3 4 5 6 7 8]
}
