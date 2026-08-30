package simduint64_test

import (
	"fmt"
	"testing"

	simduint64 "github.com/daanv2/go-math/simd/uint64"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Math ops
func Test_Uint64x4_Ops(t *testing.T) {
	dataA := []uint64{1, 2, 3, 4}
	dataB := []uint64{5, 6, 7, 8}
	vec1 := simduint64.NewUint64x4(dataA)
	vec2 := simduint64.NewUint64x4(dataB)

	t.Run("Add", func(t *testing.T) {
		result := vec1.Add(vec2)

		assert.Equal(t, []uint64{6, 8, 10, 12}, result.ToSlice())
	})

	t.Run("Mul", func(t *testing.T) {
		result := vec1.Mul(vec2)

		assert.Equal(t, []uint64{5, 12, 21, 32}, result.ToSlice())
	})

	t.Run("Max", func(t *testing.T) {
		vec1 := simduint64.NewUint64x4([]uint64{1, 10, 3, 12})
		vec2 := simduint64.NewUint64x4([]uint64{9, 2, 11, 4})

		result := vec1.Max(vec2)

		assert.Equal(t, []uint64{9, 10, 11, 12}, result.ToSlice())
	})

	t.Run("Min", func(t *testing.T) {
		vec1 := simduint64.NewUint64x4([]uint64{1, 10, 3, 12})
		vec2 := simduint64.NewUint64x4([]uint64{9, 2, 11, 4})

		result := vec1.Min(vec2)

		assert.Equal(t, []uint64{1, 2, 3, 4}, result.ToSlice())
	})

	t.Run("Sub", func(t *testing.T) {
		vec1 := simduint64.NewUint64x4([]uint64{9, 10, 11, 12})
		vec2 := simduint64.NewUint64x4([]uint64{1, 2, 3, 4})

		result := vec1.Sub(vec2)

		assert.Equal(t, []uint64{8, 8, 8, 8}, result.ToSlice())
	})
}

// Bitwise ops
func Test_Uint64x4_Bitwise(t *testing.T) {
	vecA := simduint64.NewUint64x4([]uint64{12, 10, 255, 0})
	vecB := simduint64.NewUint64x4([]uint64{10, 6, 15, 5})

	t.Run("And", func(t *testing.T) {
		result := vecA.And(vecB)

		assert.Equal(t, []uint64{8, 2, 15, 0}, result.ToSlice())
	})

	t.Run("AndNot", func(t *testing.T) {
		result := vecA.AndNot(vecB)

		assert.Equal(t, []uint64{4, 8, 240, 0}, result.ToSlice())
	})

	t.Run("Or", func(t *testing.T) {
		result := vecA.Or(vecB)

		assert.Equal(t, []uint64{14, 14, 255, 5}, result.ToSlice())
	})

	t.Run("Xor", func(t *testing.T) {
		result := vecA.Xor(vecB)

		assert.Equal(t, []uint64{6, 12, 240, 5}, result.ToSlice())
	})

	t.Run("Not", func(t *testing.T) {
		result := vecA.Not()

		assert.Equal(t, []uint64{18446744073709551603, 18446744073709551605, 18446744073709551360, 18446744073709551615}, result.ToSlice())
	})

	t.Run("ShiftLeft", func(t *testing.T) {
		vec := simduint64.NewUint64x4([]uint64{1, 2, 3, 4})
		result := vec.ShiftLeft(2)

		assert.Equal(t, []uint64{4, 8, 12, 16}, result.ToSlice())
	})

	t.Run("ShiftRight", func(t *testing.T) {
		vec := simduint64.NewUint64x4([]uint64{8, 255, 16, 3})
		result := vec.ShiftRight(1)

		assert.Equal(t, []uint64{4, 127, 8, 1}, result.ToSlice())
	})
}

// Loads
func Test_Uint64x4_Load(t *testing.T) {
	data := []uint64{1, 2, 3, 4, 5}

	assert.Equal(t, simduint64.NewUint64x4(data).ToSlice(), data[:4])
	assert.Equal(t, []uint64{1, 0, 0, 0}, simduint64.NewUint64x4(data[:1]).ToSlice())
	assert.Equal(t, []uint64{1, 2, 3, 0}, simduint64.NewUint64x4(data[:3]).ToSlice())
}

func Test_Uint64x4Slice(t *testing.T) {
	data := []uint64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9}

	result := simduint64.NewUint64x4Slice(data)

	require.Len(t, result, 3)
	require.Equal(t, []uint64{1, 2, 3, 4}, result[0].ToSlice())
	require.Equal(t, []uint64{5, 6, 7, 8}, result[1].ToSlice())
	require.Equal(t, []uint64{9, 0, 0, 0}, result[2].ToSlice())
}

// Examples

// Default example
func ExampleUint64x4() {
	data := []uint64{1, 2, 3, 4}

	v1 := simduint64.NewUint64x4(data)

	fmt.Println(v1.ToSlice())

	// Output:
	// [1 2 3 4]
}
