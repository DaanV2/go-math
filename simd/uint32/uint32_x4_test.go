package simduint32_test

import (
	"fmt"
	"testing"

	simduint32 "github.com/daanv2/go-math/simd/uint32"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Math ops
func Test_Uint32x4_Ops(t *testing.T) {
	dataA := []uint32{1, 2, 3, 4}
	dataB := []uint32{5, 6, 7, 8}
	vec1 := simduint32.NewUint32x4(dataA)
	vec2 := simduint32.NewUint32x4(dataB)

	t.Run("Add", func(t *testing.T) {
		result := vec1.Add(vec2)

		assert.Equal(t, []uint32{6, 8, 10, 12}, result.ToSlice())
	})

	t.Run("Mul", func(t *testing.T) {
		result := vec1.Mul(vec2)

		assert.Equal(t, []uint32{5, 12, 21, 32}, result.ToSlice())
	})

	t.Run("Max", func(t *testing.T) {
		vec1 := simduint32.NewUint32x4([]uint32{1, 10, 3, 12})
		vec2 := simduint32.NewUint32x4([]uint32{9, 2, 11, 4})

		result := vec1.Max(vec2)

		assert.Equal(t, []uint32{9, 10, 11, 12}, result.ToSlice())
	})

	t.Run("Min", func(t *testing.T) {
		vec1 := simduint32.NewUint32x4([]uint32{1, 10, 3, 12})
		vec2 := simduint32.NewUint32x4([]uint32{9, 2, 11, 4})

		result := vec1.Min(vec2)

		assert.Equal(t, []uint32{1, 2, 3, 4}, result.ToSlice())
	})

	t.Run("Sub", func(t *testing.T) {
		vec1 := simduint32.NewUint32x4([]uint32{9, 10, 11, 12})
		vec2 := simduint32.NewUint32x4([]uint32{1, 2, 3, 4})

		result := vec1.Sub(vec2)

		assert.Equal(t, []uint32{8, 8, 8, 8}, result.ToSlice())
	})
}

// Bitwise ops
func Test_Uint32x4_Bitwise(t *testing.T) {
	vecA := simduint32.NewUint32x4([]uint32{12, 10, 255, 0})
	vecB := simduint32.NewUint32x4([]uint32{10, 6, 15, 5})

	t.Run("And", func(t *testing.T) {
		result := vecA.And(vecB)

		assert.Equal(t, []uint32{8, 2, 15, 0}, result.ToSlice())
	})

	t.Run("AndNot", func(t *testing.T) {
		result := vecA.AndNot(vecB)

		assert.Equal(t, []uint32{4, 8, 240, 0}, result.ToSlice())
	})

	t.Run("Or", func(t *testing.T) {
		result := vecA.Or(vecB)

		assert.Equal(t, []uint32{14, 14, 255, 5}, result.ToSlice())
	})

	t.Run("Xor", func(t *testing.T) {
		result := vecA.Xor(vecB)

		assert.Equal(t, []uint32{6, 12, 240, 5}, result.ToSlice())
	})

	t.Run("Not", func(t *testing.T) {
		result := vecA.Not()

		assert.Equal(t, []uint32{4294967283, 4294967285, 4294967040, 4294967295}, result.ToSlice())
	})

	t.Run("ShiftLeft", func(t *testing.T) {
		vec := simduint32.NewUint32x4([]uint32{1, 2, 3, 4})
		result := vec.ShiftLeft(2)

		assert.Equal(t, []uint32{4, 8, 12, 16}, result.ToSlice())
	})

	t.Run("ShiftRight", func(t *testing.T) {
		vec := simduint32.NewUint32x4([]uint32{8, 255, 16, 3})
		result := vec.ShiftRight(1)

		assert.Equal(t, []uint32{4, 127, 8, 1}, result.ToSlice())
	})
}

// Loads
func Test_Uint32x4_Load(t *testing.T) {
	data := []uint32{1, 2, 3, 4, 5}

	assert.Equal(t, simduint32.NewUint32x4(data).ToSlice(), data[:4])
	assert.Equal(t, []uint32{1, 0, 0, 0}, simduint32.NewUint32x4(data[:1]).ToSlice())
	assert.Equal(t, []uint32{1, 2, 3, 0}, simduint32.NewUint32x4(data[:3]).ToSlice())
}

func Test_Uint32x4Slice(t *testing.T) {
	data := []uint32{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9}

	result := simduint32.NewUint32x4Slice(data)

	require.Len(t, result, 3)
	require.Equal(t, []uint32{1, 2, 3, 4}, result[0].ToSlice())
	require.Equal(t, []uint32{5, 6, 7, 8}, result[1].ToSlice())
	require.Equal(t, []uint32{9, 0, 0, 0}, result[2].ToSlice())
}

// Examples

// Default example
func ExampleUint32x4() {
	data := []uint32{1, 2, 3, 4}

	v1 := simduint32.NewUint32x4(data)

	fmt.Println(v1.ToSlice())

	// Output:
	// [1 2 3 4]
}
