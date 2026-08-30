package simdint64_test

import (
	"fmt"
	"testing"

	simdint64 "github.com/daanv2/go-math/simd/int64"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Math ops
func Test_Int64x2_Ops(t *testing.T) {
	dataA := []int64{1, 2}
	dataB := []int64{5, 6}
	vec1 := simdint64.NewInt64x2(dataA)
	vec2 := simdint64.NewInt64x2(dataB)

	t.Run("Add", func(t *testing.T) {
		result := vec1.Add(vec2)

		assert.Equal(t, []int64{6, 8}, result.ToSlice())
	})

	t.Run("Abs", func(t *testing.T) {
		vec := simdint64.NewInt64x2([]int64{-1, 2})
		result := vec.Abs()

		assert.Equal(t, []int64{1, 2}, result.ToSlice())
	})

	t.Run("Mul", func(t *testing.T) {
		result := vec1.Mul(vec2)

		assert.Equal(t, []int64{5, 12}, result.ToSlice())
	})

	t.Run("Max", func(t *testing.T) {
		vec1 := simdint64.NewInt64x2([]int64{1, 10})
		vec2 := simdint64.NewInt64x2([]int64{9, 2})

		result := vec1.Max(vec2)

		assert.Equal(t, []int64{9, 10}, result.ToSlice())
	})

	t.Run("Min", func(t *testing.T) {
		vec1 := simdint64.NewInt64x2([]int64{1, 10})
		vec2 := simdint64.NewInt64x2([]int64{9, 2})

		result := vec1.Min(vec2)

		assert.Equal(t, []int64{1, 2}, result.ToSlice())
	})

	t.Run("Neg", func(t *testing.T) {
		result := vec1.Neg()

		assert.Equal(t, []int64{-1, -2}, result.ToSlice())
	})

	t.Run("Sub", func(t *testing.T) {
		result := vec1.Sub(vec2)

		assert.Equal(t, []int64{-4, -4}, result.ToSlice())
	})
}

// Bitwise ops
func Test_Int64x2_Bitwise(t *testing.T) {
	vecA := simdint64.NewInt64x2([]int64{12, 255})
	vecB := simdint64.NewInt64x2([]int64{10, 15})

	t.Run("And", func(t *testing.T) {
		result := vecA.And(vecB)

		assert.Equal(t, []int64{8, 15}, result.ToSlice())
	})

	t.Run("AndNot", func(t *testing.T) {
		result := vecA.AndNot(vecB)

		assert.Equal(t, []int64{4, 240}, result.ToSlice())
	})

	t.Run("Or", func(t *testing.T) {
		result := vecA.Or(vecB)

		assert.Equal(t, []int64{14, 255}, result.ToSlice())
	})

	t.Run("Xor", func(t *testing.T) {
		result := vecA.Xor(vecB)

		assert.Equal(t, []int64{6, 240}, result.ToSlice())
	})

	t.Run("Not", func(t *testing.T) {
		result := vecA.Not()

		assert.Equal(t, []int64{-13, -256}, result.ToSlice())
	})

	t.Run("ShiftLeft", func(t *testing.T) {
		vec := simdint64.NewInt64x2([]int64{1, 2})
		result := vec.ShiftLeft(2)

		assert.Equal(t, []int64{4, 8}, result.ToSlice())
	})

	t.Run("ShiftRight", func(t *testing.T) {
		vec := simdint64.NewInt64x2([]int64{8, -8})
		result := vec.ShiftRight(1)

		assert.Equal(t, []int64{4, -4}, result.ToSlice())
	})
}

// Loads
func Test_Int64x2_Load(t *testing.T) {
	data := []int64{1, 2, 3}

	assert.Equal(t, simdint64.NewInt64x2(data).ToSlice(), data[:2])
	assert.Equal(t, []int64{1, 0}, simdint64.NewInt64x2(data[:1]).ToSlice())
}

func Test_Int64x2Slice(t *testing.T) {
	data := []int64{
		1, 2,
		3, 4,
		5}

	result := simdint64.NewInt64x2Slice(data)

	require.Len(t, result, 3)
	require.Equal(t, []int64{1, 2}, result[0].ToSlice())
	require.Equal(t, []int64{3, 4}, result[1].ToSlice())
	require.Equal(t, []int64{5, 0}, result[2].ToSlice())
}

// Examples

// Default example
func ExampleInt64x2() {
	data := []int64{1, 2}

	v1 := simdint64.NewInt64x2(data)

	fmt.Println(v1.ToSlice())

	// Output:
	// [1 2]
}
